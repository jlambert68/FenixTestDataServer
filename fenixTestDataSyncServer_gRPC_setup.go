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

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b184b9c8-e9ec-4149-a28a-c34f50c0e105",
	}).Debug("Incoming gRPC 'InitGrpcServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "8aa2cf80-d8b2-49f7-a9f0-abcf14c57744",
	}).Debug("Outgoing gRPC 'InitGrpcServer'")

	var err error

	// Find first non allocated port from defined start port
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("Fenix TestData Sync Server tries to start")

	localServerEngineLocalPort = localServerEngineLocalPort
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

	}

	// Creates a new gRPC server
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
	}).Info("Starting Fenix TestData Sync gRPC Server")
	registerfenixTestDataSyncServerServer = grpc.NewServer()
	fenixTestDataSyncServerGrpcApi.RegisterFenixTestDataGrpcServicesServer(registerfenixTestDataSyncServerServer, &FenixTestDataGrpcServicesServer{})

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                           "e843ece9-b707-4c60-b1d8-14464305e68f",
		"localServerEngineLocalPort: ": localServerEngineLocalPort,
	}).Info("Fenix TestData Sync gRPC Server - started")
	registerfenixTestDataSyncServerServer.Serve(lis)

}

// Stop Backend gRPC-server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) StopGrpcServer() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "4281e3d0-453d-4fe7-b57f-543a42d0f986",
	}).Debug("Incoming gRPC 'StopGrpcServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9a33fce7-8e35-4fbd-a8c9-d7f3f6c72582",
	}).Debug("Outgoing gRPC 'StopGrpcServer'")

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{}).Info("Gracefully stop for: 'Fenix TestData Sync gRPC Server'")
	registerfenixTestDataSyncServerServer.GracefulStop()

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"localServerEngineLocalPort: ": localServerEngineLocalPort,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
