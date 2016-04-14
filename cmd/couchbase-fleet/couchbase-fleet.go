package main

import (
	"log"

	"github.com/amimimor/couchbase-cluster-go"
)

func main() {

	usage := `Couchbase-Fleet.

Usage:
  couchbase-fleet launch-cbs --version=<cb-version> --num-nodes=<num_nodes> --userpass=<user:pass> [--edition=<edition>] [--etcd-servers=<server-list>] [--fleet-uri] [--docker-tag=<dt>] [--skip-clean-slate-check]
  couchbase-fleet stop [--all-units] [--etcd-servers=<server-list>]
  couchbase-fleet destroy [--all-units] [--etcd-servers=<server-list>]
  couchbase-fleet generate-units --version=<cb-version> --num-nodes=<num_nodes> --userpass=<user:pass> [--etcd-servers=<server-list>] [--docker-tag=<dt>] --output-dir=<output_dir>
  couchbase-fleet -h | --help

Options:
  -h --help     Show this screen.
  --version=<cb-version> Couchbase Server version (examples: latest, 3.0.3, 2.2).  The list of supported version corresponds to available tags on dockerhub: https://hub.docker.com/u/couchbase/server 
  --num-nodes=<num_nodes> number of couchbase nodes to start
  --userpass=<user:pass> the username and password as a single string, delimited by a colon (:)
  --edition=<edition> the edition to use, either "enterprise" or "community".  Defaults to "community" edition. 
  --etcd-servers=<server-list>  Comma separated list of etcd servers, or omit to connect to etcd running on localhost
  --fleet-uri=<URI> Fleet service URI formated to [http://localhost:49153 | unix:///var/run/fleet.sock]
  --docker-tag=<dt>  if present, use this docker tag for spawned containers, otherwise, default to "latest"
  --skip-clean-slate-check  if present, will skip the check that we are starting from clean state
  --output-dir=<output_dir>

`

	arguments, err := docopt.Parse(usage, nil, true, "Couchbase-Fleet", false)
	if err != nil {
		log.Fatalf("Failed to parse args: %v", err)
	}

	if cbcluster.IsCommandEnabled(arguments, "launch-cbs") {
		if err := launchCouchbaseServer(arguments); err != nil {
			log.Fatalf("Failed: %v", err)
		}
		return
	}

	if cbcluster.IsCommandEnabled(arguments, "generate-units") {
		if err := generateUnits(arguments); err != nil {
			log.Fatalf("Failed: %v", err)
		}
		return
	}

	if cbcluster.IsCommandEnabled(arguments, "stop") {
		if err := stopUnits(arguments); err != nil {
			log.Fatalf("Failed: %v", err)
		}
		return
	}

	if cbcluster.IsCommandEnabled(arguments, "destroy") {
		if err := destroyUnits(arguments); err != nil {
			log.Fatalf("Failed: %v", err)
		}
		return
	}

	log.Printf("Nothing to do!")

}

func launchCouchbaseServer(arguments map[string]interface{}) error {

	etcdServers := cbcluster.ExtractEtcdServerList(arguments)
	fleetURI, err := cbcluster.ExtractStringArg(arguments, "--fleet-uri")
	if err != nil {
		return err
	}

	couchbaseFleet := cbcluster.NewCouchbaseFleet(etcdServers, fleetURI)
	if err := couchbaseFleet.ExtractDocOptArgs(arguments); err != nil {
		return err
	}

	return couchbaseFleet.LaunchCouchbaseServer()

}

func generateUnits(arguments map[string]interface{}) error {

	etcdServers := cbcluster.ExtractEtcdServerList(arguments)
	fleetURI, err := cbcluster.ExtractStringArg(arguments, "--fleet-uri")
	if err != nil {
		return err
	}

	couchbaseFleet := cbcluster.NewCouchbaseFleet(etcdServers, fleetURI)
	if err := couchbaseFleet.ExtractDocOptArgs(arguments); err != nil {
		return err
	}

	// get the output dir from args
	outputDir, err := cbcluster.ExtractStringArg(arguments, "--output-dir")
	if err != nil {
		return err
	}

	if err := couchbaseFleet.GenerateUnits(outputDir); err != nil {
		return err
	}

	log.Printf("Unit files written to %v", outputDir)

	return nil

}

func stopUnits(arguments map[string]interface{}) error {

	etcdServers := cbcluster.ExtractEtcdServerList(arguments)
	fleetURI, err := cbcluster.ExtractStringArg(arguments, "--fleet-uri")
	if err != nil {
		return err
	}

	couchbaseFleet := cbcluster.NewCouchbaseFleet(etcdServers, fleetURI)
	allUnits := cbcluster.ExtractBoolArg(arguments, "--all-units")

	return couchbaseFleet.StopUnits(allUnits)

}

func destroyUnits(arguments map[string]interface{}) error {

	etcdServers := cbcluster.ExtractEtcdServerList(arguments)
	fleetURI, err := cbcluster.ExtractStringArg(arguments, "--fleet-uri")
	if err != nil {
		return err
	}

	couchbaseFleet := cbcluster.NewCouchbaseFleet(etcdServers, fleetURI)
	allUnits := cbcluster.ExtractBoolArg(arguments, "--all-units")

	return couchbaseFleet.DestroyUnits(allUnits)

}
