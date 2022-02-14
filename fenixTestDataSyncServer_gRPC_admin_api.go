package main

import (
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// *********************************************************************
//AreYouAlive - Fenix client can check if Fenix Testdata sync server is alive with this service
func (s *FenixTestDataGrpcServicesAdminServer) AreYouAlive(_ context.Context, emptyParameter *fenixTestDataSyncServerGrpcAdminApi.EmptyParameter) (*fenixTestDataSyncServerGrpcAdminApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "92921719-3cb7-4d8d-8876-0ad8771faba8",
	}).Debug("Incoming gRPC 'AreYouAlive (Admin)'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "6d99df63-7313-4190-b478-4e3aa1a440cc",
	}).Debug("Outgoing gRPC 'AreYouAlive (Admin)'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isAdminClientUsingCorrectTestDataProtoFileVersion("666", emptyParameter.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{AckNack: true, Comments: "I'am Fenix TestDataSyncAdminServer and I'm alive, and my time is " + fenixSyncShared.GenerateDatetimeTimeStampForDB()}, nil
}

// *********************************************************************
// AllowIncomingAndOutgoingMessages - Retry to allow incoming gRPC calls and process outgoing calls
func (s *FenixTestDataGrpcServicesAdminServer) AllowOrDisallowIncomingAndOutgoingMessages(_ context.Context, allowOrDisallowIncomingAndOutgoingMessage *fenixTestDataSyncServerGrpcAdminApi.AllowOrDisallowIncomingAndOutgoingMessage) (*fenixTestDataSyncServerGrpcAdminApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "e7b5686b-6bdf-4218-ac05-e184acf7feb1",
	}).Debug("Incoming gRPC 'AllowOrDisallowIncomingAndOutgoingMessages'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "3961dda0-7af8-4746-a5b1-aa85a6fec479",
	}).Debug("Outgoing gRPC 'AllowOrDisallowIncomingAndOutgoingMessages'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isAdminClientUsingCorrectTestDataProtoFileVersion("666", allowOrDisallowIncomingAndOutgoingMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Set state depending on parameter so incoming and outgoing messages can be processed or not
	fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = allowOrDisallowIncomingAndOutgoingMessage.AllowInAndOutgoingMessages

	return &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************
// RestartFenixServerProcesses - Restart Fenix TestData Server processes
func (s *FenixTestDataGrpcServicesAdminServer) RestartFenixServerProcesses(_ context.Context, emptyParameter *fenixTestDataSyncServerGrpcAdminApi.EmptyParameter) (*fenixTestDataSyncServerGrpcAdminApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "8ecf6bfc-2ff6-4b50-b4be-faef443e0c0e",
	}).Debug("Incoming 'RestartFenixServerProcesses'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "7144c46e-c3d9-4d44-b4c7-20641c103a27",
	}).Debug("Outgoing 'RestartFenixServerProcesses'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isAdminClientUsingCorrectTestDataProtoFileVersion("666", emptyParameter.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Reload Data from cloudDB into memoryDB
	_ = fenixTestDataSyncServerObject.loadNecessaryTestDataFromCloudDB()
	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessagesAdmin()
	if returnMessage != nil {
		// Can't process ingoing or outgoing messages right now
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}
