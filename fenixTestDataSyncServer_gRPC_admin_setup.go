package main

import (
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// Set up and start Backend gRPC-server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) InitGrpcAdminServer() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "0c729512-91a2-4401-aee2-1bfbaac3720f",
	}).Debug("Incoming gRPC 'InitGrpcAdminServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "784d0575-b325-4edd-86cf-eeeb80c662a8",
	}).Debug("Outgoing gRPC 'InitGrpcAdminServer'")

	var err error

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "cd06dd40-5b41-45f1-a39e-435a31ef67db",
	}).Info("Fenix TestData Sync Server (Admin) tries to start")

	localServerEngineLocalPort = localServerEngineLocalPort
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                                "56610866-03c0-4c51-8947-8c6440b0f740",
		"localServerEngineLocalAdminPort: ": localServerEngineLocalAdminPort,
	}).Info("Start listening on:")
	gRPCAdminLis, err = net.Listen("tcp", ":"+strconv.Itoa(localServerEngineLocalAdminPort))

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "5f2ad52c-cfab-43b0-bb96-4ef0876bc544",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":                                "66d6ae5a-e392-4c13-8c28-d62392e342be",
			"localServerEngineLocalAdminPort: ": localServerEngineLocalAdminPort,
		}).Info("Success in listening on port:")

	}

	// Creates a new gRPC server
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "5b3590a9-3842-4325-ba16-cb27dc621d47",
	}).Info("Starting Fenix TestData Sync gRPC Admin Server")
	registerfenixTestDataSyncAdminServerServer = grpc.NewServer()
	fenixTestDataSyncServerGrpcApi.RegisterFenixTestDataGrpcAdminServicesServer(registerfenixTestDataSyncAdminServerServer, &FenixTestDataGrpcServicesAdminServer{})

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                           "97f74966-71e6-4c90-a172-be0cbbc9da65",
		"localServerEngineLocalPort: ": localServerEngineLocalAdminPort,
	}).Info("Fenix TestData Sync gRPC Server - started")
	registerfenixTestDataSyncAdminServerServer.Serve(gRPCAdminLis)

}

// Stop Backend gRPC-Adminserver
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) StopGrpcAdminServer() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "84e70e96-bd41-4de5-9d16-fb1de95db5e3",
	}).Debug("Incoming gRPC 'StopGrpcAdminServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "60783ab3-45af-4f8c-967a-a8b6258c1c9d",
	}).Debug("Outgoing gRPC 'StopGrpcAdminServer'")

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{}).Info("Gracefully stop for: 'Fenix TestData Sync gRPC Admin Server'")
	registerfenixTestDataSyncServerServer.GracefulStop()

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"localServerEngineLocalAdminPort: ": localServerEngineLocalAdminPort,
	}).Info("Close net.Listing")
	_ = gRPCAdminLis.Close()

}
