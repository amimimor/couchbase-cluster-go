package main

import (
	"io/ioutil"
	"log"
	"regexp"

	"github.com/amimimor/couchbase-cluster-go"
	"github.com/docopt/docopt-go"
)

func main() {

	usage := `Sync-Gw-Config.

Usage:
  sync-gw-config rewrite --destination=<config-dest> [--etcd-servers=<server-list>] [--fleet-uri]
  sync-gw-config -h | --help

Options:
  -h --help     Show this screen.
  --etcd-servers=<server-list>  Comma separated list of etcd servers, or omit to connect to etcd running on localhost
  --fleet-uri=<URI> Fleet service URI formated to http://localhost:49153 | unix:///var/run/fleet.sock
  --destination=<config-dest> The path where the updated config should be written
`

	arguments, err := docopt.Parse(usage, nil, true, "Sync-Gw-Config", false)
	log.Printf("args: %v.  err: %v", arguments, err)

	if cbcluster.IsCommandEnabled(arguments, "rewrite") {
		if err := rewriteConfig(arguments); err != nil {
			log.Fatalf("Failed: %v", err)
		}
		return
	}

	log.Printf("Nothing to do!")

}

// does this config need to be rerwritten?  if it doesn't have
// any placeholder variables, then the answer is no.
func requiresRewrite(syncGwConfig string) bool {
	re := regexp.MustCompile(`{{.*}}`)
	placeholder := re.FindString(syncGwConfig)
	return placeholder != ""
}

func rewriteConfig(arguments map[string]interface{}) error {

	etcdServers := cbcluster.ExtractEtcdServerList(arguments)
	fleetURI, err := cbcluster.ExtractStringArg(arguments, "--fleet-uri")
	if err != nil {
		return err
	}

	dest, err := cbcluster.ExtractStringArg(arguments, "--destination")
	if err != nil {
		return err
	}

	syncGwCluster := cbcluster.NewSyncGwCluster(etcdServers, fleetURI)

	// get the sync gw config from etcd (cbcluster.KEY_SYNC_GW_CONFIG)
	syncGwConfig, err := syncGwCluster.FetchSyncGwConfig()

	if !requiresRewrite(syncGwConfig) {
		log.Printf("No placeholder variables in config, no rewrite required")
	} else {

		// get a couchbase live node
		couchbaseCluster := cbcluster.NewCouchbaseCluster(etcdServers)

		liveNodeIp, err := couchbaseCluster.FindLiveNode()

		log.Printf("LiveNodeIp: %v", liveNodeIp)

		if err != nil {
			return err
		}

		// run the sync gw config through go templating engine
		syncGwConfigBytes, err := syncGwCluster.UpdateConfig(liveNodeIp, syncGwConfig)
		if err != nil {
			return err
		}
		syncGwConfig = string(syncGwConfigBytes)

	}

	log.Printf("Sync GW config file: %v", string(syncGwConfig))

	// write the new config to the dest file
	if err := ioutil.WriteFile(dest, []byte(syncGwConfig), 0644); err != nil {
		return err
	}

	return nil

}
