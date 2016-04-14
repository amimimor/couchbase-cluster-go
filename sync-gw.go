package cbcluster

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/tleyden/go-etcd/etcd"
)

const (
	KEY_SYNC_GW_NODE_STATE = "/couchbase.com/sync-gw-node-state"
	KEY_SYNC_GW_CONFIG     = "/couchbase.com/sync-gateway/config"
)

type SyncGwCluster struct {
	etcdClient               *etcd.Client
	EtcdServers              []string
	NumNodes                 int
	ContainerTag             string
	ConfigUrl                string
	CreateBucketName         string
	CreateBucketSize         int
	CreateBucketReplicaCount int
	LocalIp                  string
	RequiresCouchbaseServer  bool
	LaunchNginxEnabled       bool
	fleetClient              *CouchbaseFleet
}

func NewSyncGwCluster(etcdServers []string, fleetURI string) *SyncGwCluster {

	s := &SyncGwCluster{}
	c := NewCouchbaseFleet(etcdServers, fleetURI)

	if len(etcdServers) > 0 {
		s.EtcdServers = etcdServers
		c.EtcdServers = etcdServers
		log.Printf("Connect to explicit etcd servers: %v", s.EtcdServers)
	} else {
		s.EtcdServers = []string{}
		c.EtcdServers = []string{}
		log.Printf("Connect to etcd on localhost")
	}

	if fleetURI != "" {
		c.FleetURI = FLEET_API_ENDPOINT
	}

	s.fleetClient = c

	s.ConnectToEtcd()
	return s

}

func (s *SyncGwCluster) ConnectToEtcd() {

	s.etcdClient = etcd.NewClient(s.EtcdServers)
	s.etcdClient.SetConsistency(etcd.STRONG_CONSISTENCY)
}

func (s *SyncGwCluster) ExtractDocOptArgs(arguments map[string]interface{}) error {

	numnodes, err := ExtractNumNodes(arguments)
	if err != nil {
		return err
	}
	s.NumNodes = numnodes

	configUrl, err := ExtractStringArg(arguments, "--config-url")
	if err != nil {
		return err
	}
	if configUrl == "" {
		return fmt.Errorf("Missing or empty config url")
	}
	s.ConfigUrl = configUrl

	createBucketName, _ := ExtractStringArg(arguments, "--create-bucket")
	if createBucketName != "" {
		s.CreateBucketName = createBucketName
	}

	createBucketSize, _ := ExtractIntArg(arguments, "--create-bucket-size")
	s.CreateBucketSize = createBucketSize
	if s.CreateBucketSize == 0 {
		s.CreateBucketSize = 512
	}

	createBucketReplicaCount, _ := ExtractIntArg(arguments, "--create-bucket-replicas")
	s.CreateBucketReplicaCount = createBucketReplicaCount
	if s.CreateBucketReplicaCount <= 0 {
		s.CreateBucketReplicaCount = 1
	}

	s.ContainerTag = ExtractDockerTagOrLatest(arguments)

	s.RequiresCouchbaseServer = !ExtractBoolArg(arguments, "--in-memory-db")

	s.LaunchNginxEnabled = ExtractBoolArg(arguments, "--launch-nginx")

	return nil
}

func (s SyncGwCluster) UpdateConfig(liveNodeIp, configTemplate string) (config []byte, err error) {

	tmpl, err := template.New("sgw_config").Parse(configTemplate)
	if err != nil {
		return nil, err
	}

	params := struct {
		COUCHBASE_SERVER_IP string
	}{
		COUCHBASE_SERVER_IP: liveNodeIp,
	}

	out := &bytes.Buffer{}

	// execute template and write to dest
	err = tmpl.Execute(out, params)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil

}

func (s SyncGwCluster) FetchSyncGwConfig() (config string, err error) {
	log.Printf("FetchSyncGwConfig()")
	configUrl, err := s.FetchSyncGwConfigUrl()
	if err != nil {
		return "", err
	}
	resp, err := http.Get(configUrl)
	if err != nil {
		return "", fmt.Errorf("Error %v getting sync gw config from %v", err, configUrl)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid status %v getting sync gw config from %v", resp.StatusCode, configUrl)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s SyncGwCluster) FetchSyncGwConfigUrl() (configUrl string, err error) {
	response, err := s.etcdClient.Get(KEY_SYNC_GW_CONFIG, false, false)
	if err != nil {
		return "", err
	}
	return response.Node.Value, nil
}

func (s SyncGwCluster) LaunchSyncGateway() error {

	log.Printf("Launching sync gw")

	// create bucket (if user asked for this)
	if err := s.createBucketIfNeeded(); err != nil {
		return err
	}

	// stash some values into etcd
	if err := s.addValuesEtcd(); err != nil {
		return err
	}

	// kick off fleet units
	if err := s.kickOffFleetUnits(); err != nil {
		return err
	}

	// kick off fleet sidekicks
	if err := s.kickOffFleetSidekickUnits(); err != nil {
		return err
	}

	// wait for all sync gw nodes to be running
	if err := s.waitForAllSyncGwNodesRunning(); err != nil {
		return err
	}

	// launch nginx (if enabled by command line arg)
	if s.LaunchNginxEnabled {
		log.Printf("Launching Nginx")
		if err := s.LaunchNginx(); err != nil {
			return err
		}
	} else {
		log.Printf("Not Launching Nginx")
	}

	log.Printf("Your sync gateway cluster has been launched successfully!")

	return nil
}

// wait for s.NumNodes to appear in etcd /couchbase.com/sgw-node-state
// and able to be reached on port 4984
func (s SyncGwCluster) waitForAllSyncGwNodesRunning() error {

	maxAttempts := 50

	worker := func() (finished bool, err error) {

		ipAddresses, err := s.syncGwIpAddresses()
		if err != nil {
			log.Printf("syncGwIpAddresses returned err: %v", err)
			return false, nil
		}

		if len(ipAddresses) < s.NumNodes {
			log.Printf("%v sync gateways running, expected %v", len(ipAddresses), s.NumNodes)
			return false, nil
		}

		err = s.checkSyncGwNodesRunning(ipAddresses)
		if err != nil {
			log.Printf("checkSyncGwNodesRunning returned err: %v", err)
			return false, nil
		}

		return true, nil

	}

	sleeper := func(numAttempts int) (bool, int) {
		if numAttempts > maxAttempts {
			return false, -1
		}
		sleepSeconds := numAttempts * 2
		return true, sleepSeconds
	}

	return RetryLoop(worker, sleeper)

}

func (s SyncGwCluster) checkSyncGwNodesRunning(nodeIpAddresses []string) error {

	for _, syncGwIpAddress := range nodeIpAddresses {

		// TODO: don't use hardcoded port.  I guess this could get pulled from
		// the sync gateway configuration
		endpointUrl := fmt.Sprintf("http://%v:4984/", syncGwIpAddress)
		log.Printf("Waiting for Sync Gw at %v to be up", endpointUrl)
		resp, err := http.Get(endpointUrl)
		if err != nil {
			return fmt.Errorf("Error %v connecting to %v", err, endpointUrl)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return fmt.Errorf("Unable to connect to %v", endpointUrl)
		}

	}

	return nil

}

func (s SyncGwCluster) syncGwIpAddresses() (nodeIpAddresses []string, err error) {

	response, err := s.etcdClient.Get(KEY_SYNC_GW_NODE_STATE, false, false)
	if err != nil {
		return nil, fmt.Errorf("Error getting key.  Err: %v", err)
	}

	node := response.Node

	if node == nil {
		log.Printf("node is nil, returning")
		return nil, nil
	}

	if len(node.Nodes) == 0 {
		log.Printf("len(node.Nodes) == 0, returning")
		return nil, nil
	}

	nodeIpAddresses = []string{}

	for _, subNode := range node.Nodes {

		// the key will be: /couchbase.com/sync-gw-node-state/172.17.8.101, but we
		// only want the last element in the path
		_, subNodeIp := path.Split(subNode.Key)

		nodeIpAddresses = append(nodeIpAddresses, subNodeIp)

	}

	return nodeIpAddresses, nil

}

func (s SyncGwCluster) LaunchSyncGatewaySidekick() error {

	if s.LocalIp == "" {
		return fmt.Errorf("You must define LocalIp before calling")
	}

	// create /couchbase.com/sync-gw-node-state/ directory
	if err := s.CreateNodeStateDirectoryKey(); err != nil {
		return err
	}

	s.EventLoop()

	return fmt.Errorf("Event loop died") // should never get here

}

func (s SyncGwCluster) CreateNodeStateDirectoryKey() error {

	// since we don't knoow how long it will be until we go
	// into the event loop, set TTL to 0 (infinite) for now.
	_, err := s.etcdClient.CreateDir(KEY_SYNC_GW_NODE_STATE, TTL_NONE)

	if err != nil {
		// expected error where someone beat us out
		if strings.Contains(err.Error(), "Key already exists") {
			return nil
		}

		// otherwise, unexpected error
		log.Printf("Unexpected error creating %v: %v", KEY_SYNC_GW_NODE_STATE, err)
		return err
	}

	return nil

}

func (s SyncGwCluster) EventLoop() {

	for {
		// update the node-state directory ttl.  we want this directory
		// to disappear in case all nodes in the cluster are down, since
		// otherwise it would just be unwanted residue.
		ttlSeconds := uint64(10)
		_, err := s.etcdClient.UpdateDir(KEY_SYNC_GW_NODE_STATE, ttlSeconds)
		if err != nil {
			msg := fmt.Sprintf("Error updating %v dir in etc with new TTL. "+
				"Ignoring error, but this could cause problems",
				KEY_SYNC_GW_NODE_STATE)
			log.Printf(msg)
		}

		if err := s.PublishNodeStateEtcd(ttlSeconds); err != nil {
			msg := fmt.Sprintf("Error publishing node state to etcd: %v. "+
				"Check if etcd is running.",
				err)
			log.Printf(msg)
		}

		// sleep for a while
		<-time.After(time.Second * time.Duration(ttlSeconds/2))

	}

}

func (s SyncGwCluster) PublishNodeStateEtcd(ttlSeconds uint64) error {

	key := path.Join(KEY_SYNC_GW_NODE_STATE, s.LocalIp)

	// TODO: don't hardcode port
	ipAndPort := fmt.Sprintf("%v:4984", s.LocalIp)
	_, err := s.etcdClient.Set(key, ipAndPort, ttlSeconds)

	return err
}

func (s SyncGwCluster) kickOffFleetUnits() error {

	fleetUnitJson, err := s.generateFleetUnitJson()
	if err != nil {
		return err
	}

	for i := 1; i < s.NumNodes+1; i++ {

		if err := s.fleetClient.launchFleetUnitN(i, "sync_gw_node", fleetUnitJson); err != nil {
			return err
		}

	}

	return nil
}

func (s SyncGwCluster) generateNodeFleetUnitFile() (string, error) {

	assetName := "data/sync_gw_node@.service.template"
	content, err := Asset(assetName)
	if err != nil {
		return "", fmt.Errorf("could not find asset: %v.  err: %v", assetName, err)
	}

	params := struct {
		CONTAINER_TAG      string
		WAIT_UNTIL_RUNNING string
	}{
		CONTAINER_TAG: s.ContainerTag,
	}

	if s.RequiresCouchbaseServer {
		params.WAIT_UNTIL_RUNNING = "wait-until-running"
		log.Printf("params.WAIT_UNTIL_RUNNING")
	} else {
		params.WAIT_UNTIL_RUNNING = "--help"
		log.Printf("params.!!WAIT_UNTIL_RUNNING")
	}

	return generateUnitFileFromTemplate(content, params)

}

func (s SyncGwCluster) generateSidekickFleetUnitFile(unitNumber string) (string, error) {

	assetName := "data/sync_gw_sidekick@.service.template"
	content, err := Asset(assetName)
	if err != nil {
		return "", fmt.Errorf("could not find asset: %v.  err: %v", assetName, err)
	}

	params := struct {
		CONTAINER_TAG string
		UNIT_NUMBER   string
	}{
		CONTAINER_TAG: s.ContainerTag,
		UNIT_NUMBER:   unitNumber,
	}

	return generateUnitFileFromTemplate(content, params)

}

func (s SyncGwCluster) generateFleetUnitJson() (string, error) {

	unitFile, err := s.generateNodeFleetUnitFile()
	if err != nil {
		return "", err
	}

	log.Printf("Sync Gw fleet unit: %v", unitFile)

	jsonBytes, err := unitFileToJson(unitFile)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), err

}

func (s SyncGwCluster) kickOffFleetSidekickUnits() error {

	for i := 1; i < s.NumNodes+1; i++ {

		fleetUnitJson, err := s.generateFleetSidekickUnitJson(i)
		if err != nil {
			return err
		}

		if err := s.fleetClient.launchFleetUnitN(i, "sync_gw_sidekick", fleetUnitJson); err != nil {
			return err
		}

	}

	return nil
}

func (s SyncGwCluster) generateFleetSidekickUnitJson(unitNumber int) (string, error) {

	unitNumberStr := fmt.Sprintf("%v", unitNumber)
	unitFile, err := s.generateSidekickFleetUnitFile(unitNumberStr)
	if err != nil {
		return "", err
	}

	log.Printf("Sync Gw sidekick fleet unit: %v", unitFile)

	jsonBytes, err := unitFileToJson(unitFile)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), err

}

func (s SyncGwCluster) addValuesEtcd() error {

	// add values to etcd
	_, err := s.etcdClient.Set(KEY_SYNC_GW_CONFIG, s.ConfigUrl, 0)
	if err != nil {
		return err
	}

	return nil

}

func (s SyncGwCluster) createBucketIfNeeded() error {

	if s.CreateBucketName == "" {
		return nil
	}

	cb := NewCouchbaseCluster(s.EtcdServers)

	if err := cb.LoadAdminCredsFromEtcd(); err != nil {
		return err
	}

	liveNodeIp, err := cb.FindLiveNode()
	if err != nil {
		return err
	}
	cb.LocalCouchbaseIp = liveNodeIp

	ramQuotaMB := fmt.Sprintf("%v", s.CreateBucketSize)
	replicaNumber := fmt.Sprintf("%v", s.CreateBucketReplicaCount)

	bucketParams := bucketParams{
		Name:          s.CreateBucketName,
		RamQuotaMB:    ramQuotaMB,
		AuthType:      "none",
		ReplicaNumber: replicaNumber,
	}

	return cb.CreateBucket(bucketParams)

}

func (s SyncGwCluster) LaunchNginx() error {

	fleetUnits := map[string]string{
		"confdata": "data/confdata.service",
		"confd":    "data/confd.service",
		"nginx":    "data/nginx.service",
	}

	for unitName, unitFilePath := range fleetUnits {
		if err := s.fleetClient.launchFleetUnitFile(unitName, unitFilePath); err != nil {
			return err
		}
	}

	return nil

}
