package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend gRPC Server
		fenixTestDataSyncServerObject.StopGrpcServer()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func FenixServerMain() {

	// Init variables
	dbDataMap = make(map[memDBClientUuidType]*tempDBStruct)

	tempDBStruct_val := &tempDBStruct{
		clientData: &tempDBDataStruct{
			merkleHash: "APAN",
		},
	}
	a, b := dbDataMap[memDBClientUuidType("Hej")]
	fmt.Println(a, b)
	dbDataMap[memDBClientUuidType("Hej")] = tempDBStruct_val
	a, b = dbDataMap[memDBClientUuidType("Hej")]
	fmt.Println(a, b)

	connectToDB()

	// Set up BackendObject
	fenixTestDataSyncServerObject = &fenixTestDataSyncServerObjectStruct{stateProcessIncomingAndOutgoingMessage: true}

	// Init logger
	fenixTestDataSyncServerObject.InitLogger("")

	// Celan up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Load Data from cloudDB into memoryDB
	_ = fenixTestDataSyncServerObject.loadTestDataFromCloudDB()

	// Start Backend gRPC-server
	fenixTestDataSyncServerObject.InitGrpcServer()

}
