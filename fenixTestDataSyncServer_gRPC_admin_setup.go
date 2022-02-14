package main

import (
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// InitGrpcAdminServer :
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

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                                 "56610866-03c0-4c51-8947-8c6440b0f740",
		"fenixTestDataSyncServerAdminPort: ": fenixTestDataSyncServerAdminPort,
	}).Info("Admin Server Start listening on:")
	gRPCAdminLis, err = net.Listen("tcp", ":"+strconv.Itoa(fenixTestDataSyncServerAdminPort))

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "5f2ad52c-cfab-43b0-bb96-4ef0876bc544",
			"err: ": err,
		}).Error("Fenix Server failed to listen:")
	} else {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":                                 "66d6ae5a-e392-4c13-8c28-d62392e342be",
			"fenixTestDataSyncServerAdminPort: ": fenixTestDataSyncServerAdminPort,
		}).Info("Admin Server Success in listening on port:")

	}

	// Creates a new gRPC server
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "5b3590a9-3842-4325-ba16-cb27dc621d47",
	}).Info("Starting Fenix TestData Sync gRPC Admin Server")
	registerfenixTestDataSyncAdminServerServer = grpc.NewServer()
	fenixTestDataSyncServerGrpcAdminApi.RegisterFenixTestDataGrpcAdminServicesServer(registerfenixTestDataSyncAdminServerServer, &FenixTestDataGrpcServicesAdminServer{})

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                                 "97f74966-71e6-4c90-a172-be0cbbc9da65",
		"fenixTestDataSyncServerAdminPort: ": fenixTestDataSyncServerAdminPort,
	}).Info("Fenix TestData Sync gRPC Admin Server - started")
	registerfenixTestDataSyncAdminServerServer.Serve(gRPCAdminLis)

}

// StopGrpcAdminServer :
// Stop Backend gRPC-AdminServer
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
		"fenixTestDataSyncServerAdminPort: ": fenixTestDataSyncServerAdminPort,
	}).Info("Fenix Admin Server Close net.Listing")
	_ = gRPCAdminLis.Close()

}
