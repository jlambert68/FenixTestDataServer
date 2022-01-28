package main

import (
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// Set up and start Backend gRPC-server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) InitGrpcServer() {

	var err error

	// Find first non allocated port from defined start port
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("Backend Server tries to start")
	for counter := 0; counter < 10; counter++ {
		localServerEngineLocalPort = localServerEngineLocalPort + counter
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":                           "ca3593b1-466b-4536-be91-5e038de178f4",
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("Start listening on:")
		lis, err = net.Listen("tcp", ":"+strconv.Itoa(localServerEngineLocalPort))

		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
				"err: ": err,
			}).Error("failed to listen:")
		} else {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":                           "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
				"localServerEngineLocalPort: ": localServerEngineLocalPort,
			}).Info("Success in listening on port:")

			break
		}
	}

	// Creates a new RegisterWorkerServer gRPC server
	//	go func() {
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
	}).Info("Starting Backend gRPC Server")
	registerfenixTestDataSyncServerServer = grpc.NewServer()
	fenixTestDataSyncServerGrpcApi.RegisterFenixTestDataGrpcServicesServer(registerfenixTestDataSyncServerServer, &FenixTestDataGrpcServicesServer{})

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                           "e843ece9-b707-4c60-b1d8-14464305e68f",
		"localServerEngineLocalPort: ": localServerEngineLocalPort,
	}).Info("registerfenixTestDataSyncServerServer for TestInstruction Backend Server started")
	registerfenixTestDataSyncServerServer.Serve(lis)
	//	}()

}

// Stop Backend gRPC-server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) StopGrpcServer() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")
	registerfenixTestDataSyncServerServer.GracefulStop()

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"localServerEngineLocalPort: ": localServerEngineLocalPort,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
