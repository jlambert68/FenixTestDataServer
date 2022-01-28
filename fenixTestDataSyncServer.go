package main

import (
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

	connectToDB()

	// Set up BackendObject
	fenixTestDataSyncServerObject = &fenixTestDataSyncServerObjectStruct{}

	// Init logger
	fenixTestDataSyncServerObject.InitLogger("")

	// Celan up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Load Data from cloudDB into memoryDB
	fenixTestDataSyncServerObject.loadCloudDBIntoMemoryDB()

	// Start Backend gRPC-server
	fenixTestDataSyncServerObject.InitGrpcServer()

	// Register at QML Server
	// TODO Detta ska inte göras. Denna komponent ska vara passiv
	//fenixTestDataSyncServerObject.SendMQmlServerIpAndPortForBackendServer()
	/*
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			cleanup()
			os.Exit(0)
		}()

		for {
			fmt.Println("sleeping...for another 5 minutes")
			time.Sleep(300 * time.Second) // or runtime.Gosched() or similar per @misterbee
		}


	*/
	//Wait until user exit
	/*
		   for {
			   time.Sleep(10)
		   }
	*/
}
