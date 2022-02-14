package main

import (
	"fmt"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"os"
	"strconv"
)

func main() {
	//time.Sleep(15 * time.Second)
	FenixServerMain()
}

// Extracting all environment variables at startup
func init() {

	var err error

	// Get Environment variable to tell how this program was started
	var executionLocationForClient = fenixSyncShared.MustGetEnvironmentVariable("ExecutionLocationForClient")

	switch executionLocationForClient {
	case "LOCALHOST_NODOCKER":
		fenixSyncShared.ExecutionLocationForClient = fenixSyncShared.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		fenixSyncShared.ExecutionLocationForClient = fenixSyncShared.LocalhostDocker

	case "GCP":
		fenixSyncShared.ExecutionLocationForClient = fenixSyncShared.GCP

	default:
		fmt.Println("Unknown Execution location for Client: " + executionLocationForClient + ". Expected one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
		os.Exit(0)

	}

	// Get Environment variable to tell where Fenix TestData Sync Server is started
	var executionLocationForFenixTestDataServer = fenixSyncShared.MustGetEnvironmentVariable("ExecutionLocationForFenixTestDataServer")

	switch executionLocationForFenixTestDataServer {
	case "LOCALHOST_NODOCKER":
		fenixSyncShared.ExecutionLocationForFenixTestDataServer = fenixSyncShared.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		fenixSyncShared.ExecutionLocationForFenixTestDataServer = fenixSyncShared.LocalhostDocker

	case "GCP":
		fenixSyncShared.ExecutionLocationForFenixTestDataServer = fenixSyncShared.GCP

	default:
		fmt.Println("Unknown Execution location for Fenix TestData Syn Server: " + executionLocationForFenixTestDataServer + ". Expected one of the following: LOCALHOST_NODOCKER, LOCALHOST_DOCKER, GCP")
		os.Exit(0)

	}

	// Extract all other Environment variables
	// Address to Fenix TestData Sync server
	fenixTestDataSyncServerAddress = fenixSyncShared.MustGetEnvironmentVariable("FenixTestDataSyncServerAddress")

	// Port for Fenix TestData Sync server
	fenixTestDataSyncServerPort, err = strconv.Atoi(fenixSyncShared.MustGetEnvironmentVariable("FenixTestDataSyncServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'FenixTestDataSyncServerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Port for Fenix TestData Sync Admin server
	fenixTestDataSyncServerAdminPort, err = strconv.Atoi(fenixSyncShared.MustGetEnvironmentVariable("FenixTestDataSyncServerAdminPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'FenixTestDataSyncServerAdminPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Address to Client TestData Sync server
	clientTestDataSyncServerAddress = fenixSyncShared.MustGetEnvironmentVariable("ClientTestDataSyncServerAddress")

	// Port for Client TestData Sync server
	clientTestDataSyncServerPort, err = strconv.Atoi(fenixSyncShared.MustGetEnvironmentVariable("ClientTestDataSyncServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'clientTestDataSyncServerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Create the Dial-up string to Fenix Client TestData SyncServer
	fenixclienttestdatasyncserverAddressToDial = clientTestDataSyncServerAddress + ":" + strconv.Itoa(clientTestDataSyncServerPort)

}
