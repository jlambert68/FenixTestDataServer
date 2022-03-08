package main

import (
	"fmt"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
)

// Used for only process cleanup once
var cleanupProcessed = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend gRPC Server
		fenixTestDataSyncServerObject.StopGrpcServer()

		// Stop backend Admin gRPC Server
		fenixTestDataSyncServerObject.StopGrpcAdminServer()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func FenixServerMain() {

	// Init variables
	dbDataMap = make(map[memDBClientUuidType]*tempDBStruct)

	tempdbstructVal := &tempDBStruct{
		clientData: &tempDBDataStruct{
			merkleHash: "APAN",
		},
	}

	a, b := dbDataMap[("Hej")]
	fmt.Println(a, b)
	dbDataMap[("Hej")] = tempdbstructVal
	a, b = dbDataMap[("Hej")]
	fmt.Println(a, b)

	fenixSyncShared.ConnectToDB()

	// Set up BackendObject
	fenixTestDataSyncServerObject = &fenixTestDataSyncServerObjectStruct{
		stateProcessIncomingAndOutgoingMessage: true,
		currentTestDataState:                   CurrenStateMerkleHash,
	}

	// Init logger
	fenixTestDataSyncServerObject.InitLogger("")

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Load Data from cloudDB into memoryDB
	_ = fenixTestDataSyncServerObject.loadNecessaryTestDataFromCloudDB()

	// Initiate client data in memoryDB

	// Start Backend gRPC-server and Admin-gRPC-server
	go fenixTestDataSyncServerObject.InitGrpcServer()
	fenixTestDataSyncServerObject.InitGrpcAdminServer()

}

// TODO slå ihop common_code till ett bibliotek
