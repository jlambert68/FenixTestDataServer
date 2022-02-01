package main

import (
	"FenixTestDataServer/common_config"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// *********************************************************************
//Fenix client can check if Fenix Testdata sync server is alive with this service
func (s *FenixTestDataGrpcServicesServer) AreYouAlive(ctx context.Context, emptyParameter *fenixTestDataSyncServerGrpcApi.EmptyParameter) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
	}).Debug("Incoming gRPC 'AreYouAlive'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84isClientUsingCorrectTestDataProtoFileVersion",
	}).Debug("Incoming gRPC 'AreYouAlive'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(emptyParameter.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: "I'am Fenix TestDataSyncServer and I'm alive"}, nil
}

// *********************************************************************
// Fenix client can send TestData MerkleHash to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleHash(ctx context.Context, merkleHashMessage *fenixTestDataSyncServerGrpcApi.MerkleHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "a55f9c82-1d74-44a5-8662-058b8bc9e48f",
	}).Debug("Incoming gRPC 'SendMerkleHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "27fb45fe-3266-41aa-a6af-958513977e28",
	}).Debug("Incoming gRPC 'SendMerkleHash'")

	// Get calling client
	callingClientUuid := merkleHashMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleHashMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not the correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForClient(*merkleHashMessage)

	// Compare current server- and client MerkleHash
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientUuid)

	//  if different in MerkleHash then ask client for MerkleTree
	if currentServerMerkleHash != merkleHashMessage.MerkleHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("There is different in MerkleHash so ask client for MerkleTree for Client: " + callingClientUuid)

		defer fenixTestDataSyncServerObject.AskClientToSendMerkleTree(callingClientUuid)
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleTree(ctx context.Context, merkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cffc25f0-b0e6-407a-942a-71fc74f831ac",
	}).Debug("Incoming gRPC 'SendMerkleTree'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "61e2c28d-b091-442a-b7f8-d2502d9547cf",
	}).Debug("Incoming gRPC 'SendMerkleTree'")

	// Get calling client
	callingClientUuid := merkleTreeMessage.TestDataClientUuid

	// Verify that Client Exists in DB
	verfied, err := fenixTestDataSyncServerObject.existsClientInDB(callingClientUuid)
	if err != nil || verfied == false {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "f5eb251c-8639-496f-9101-bac8fc9300f7",
		}).Info(err.Error())

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: false, Comments: err.Error()}, nil
	}

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleTreeMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Convert the merkleTree into a DataFrame object
	merkleTreeAsDataFrame := fenixTestDataSyncServerObject.convertgRpcMerkleTreeMessageToDataframe(*merkleTreeMessage)

	//_ = fenixTestDataSyncServerObject.convertDataFramIntoJSON(merkleTreeAsDataFrame)

	// Verify MerkleTree
	clientsMerkleRootHash := common_config.ExtractMerkleRootHashFromMerkleTree(merkleTreeAsDataFrame)
	recalculatedMerkleRootHash := common_config.CalculateMerkleHashFromMerkleTree(merkleTreeAsDataFrame)

	// Something is wrong with clients hash computation
	if clientsMerkleRootHash != recalculatedMerkleRootHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes")

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: false, Comments: "There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes"}, nil
	}

	// Save the MerkleTree Dataframe message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleTreeForClient(callingClientUuid, merkleTreeAsDataFrame)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
	}).Debug("Saved MerkleTree for Client: " + callingClientUuid)

	// Compare current server- and client merklehash
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientUuid)

	//  if different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got
	if currentServerMerkleHash != recalculatedMerkleRootHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e011e854-7854-425f-9592-dcfc785203cf",
		}).Debug("Different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got. Client: " + callingClientUuid)

		defer fenixTestDataSyncServerObject.AskClientToSendAllTestDataRows(callingClientUuid)
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaderHash(ctx context.Context, testDataHeaderHashMessageMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c8e72834-338c-48be-885c-f083fe7951a6",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b3576bdc-ac11-44cd-9fa2-cd7065f1253d",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	// Get calling client
	callingClientUuid := testDataHeaderHashMessageMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderHashMessageMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Convert gRPC-message into other 'format'
	clientHeaderHash := testDataHeaderHashMessageMessage.TestDataHeaderItemsHash

	// Get current client Header Hash
	currentClientHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashForClient(callingClientUuid)

	// If Header Hash is already in DB then return OK
	if currentClientHeaderHash == clientHeaderHash {
		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientUuid, clientHeaderHash)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2cdb17e4-95c8-4d56-b2cd-7f4c8829c735",
	}).Debug("Saved Header hash to DB for client: " + callingClientUuid)

	// Check if Server Header Hash is the same as received Client HeaderHash
	serverHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashForServer(callingClientUuid)
	if serverHeaderHash != clientHeaderHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Debug("Server Header hash is not the same as Client Header Hash, Ask Client for all Headers, Client: " + callingClientUuid)

		// Ask Client for all Headers
		defer fenixTestDataSyncServerObject.AskClientToSendTestDataHeaders(callingClientUuid)

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaders(ctx context.Context, testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aee48999-12ad-4bb7-bc8a-96b62a8eeedf",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "ca0b58a8-6d56-4392-8751-45906670e86b",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	// Get calling client
	callingClientUuid := testDataHeaderMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Convert gRPC-message into other 'format'
	headerHash, headerItems := fenixTestDataSyncServerObject.convertgRpcHeaderMessageToStringArray(*testDataHeaderMessage)

	// Validate HeaderHash
	computedHeaderHash := common_config.HashValues(headerItems, true)
	if computedHeaderHash != headerHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Info("Header hash is not correct computed from Client. Expected '" + computedHeaderHash + "' as HeaderHash but got " + headerHash)

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: false, Comments: "Header hash is not correct computed. Expected '" + computedHeaderHash + "' as HeaderHash"}, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientUuid, headerHash)
	_ = fenixTestDataSyncServerObject.saveCurrentHeadersForClient(callingClientUuid, headerItems)

	// Replace Server version of Headers with Client version of Headers
	_ = fenixTestDataSyncServerObject.moveCurrentHeaderHashFromClientToServer(callingClientUuid)
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "6c90bc16-890a-402e-91b1-68fc060986c6",
	}).Debug("Saved Header hash and Headers to DB for client: " + callingClientUuid)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************
// Fenix client can send TestData rows to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataRows(ctx context.Context, testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2b1c8752-eb84-4c15-b8a7-22e2464e5168",
	}).Debug("Incoming gRPC 'SendTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "755e8b4f-f184-4277-ad41-e041714c2ca8",
	}).Debug("Incoming gRPC 'SendTestDataRows'")

	// Get calling client
	callingClientUuid := testdataRowsMessages.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testdataRowsMessages.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Convert proto-message for rows into Dataframe object
	newRowsAsDataFrame, returnMessage := fenixTestDataSyncServerObject.convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages)

	// When row-hash is wrong calculated from client then respond that back to client
	if returnMessage != nil {
		return returnMessage, nil
	}

	currentTestDataHeaders := fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientUuid)

	// If there are no headers in Database then Ask client for HeaderHash
	if len(currentTestDataHeaders) == 0 {
		fenixTestDataSyncServerObject.AskClientToSendTestDataHeaderHash(callingClientUuid)
		currentTestDataHeaders = fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientUuid)

	}

	// Concatenate with current Server data
	allRowsAsDataFrame := fenixTestDataSyncServerObject.concartenateWithCurrentServerTestData(callingClientUuid, newRowsAsDataFrame)

	// Recreate MerkleHash from All testdata rows, both existing rows for Server and new from Client
	computedMerkleHash, _ := common_config.CreateMerkleTreeFromDataFrame(allRowsAsDataFrame)

	//Compare 'computedMerkleHash' with MerkleHash from Client
	clientMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForClient(callingClientUuid)

	if computedMerkleHash != clientMerkleHash {
		// Computed MerkleHash is not the same as the one sent by the Client

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_MERKLEHASH_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "MerklRoot hash is not the same as sent to server. Got " + clientMerkleHash + ", but recalculated to " + computedMerkleHash + " from testdata",
			ErrorCodes: errorCodes,
		}

		// respond back to client when it used wrong proto-file
		return returnMessage, nil

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// RegisterTestDataClient Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) RegisterTestDataClient(ctx context.Context, testDataClientInformationMessage *fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5133b80b-6f3a-4562-9e62-1b3ceb169cc1",
	}).Debug("Incoming gRPC 'RegisterTestDataClient'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "316dcd7e-2229-4a82-b15b-0f808c2dd8aa",
	}).Debug("Incoming gRPC 'RegisterTestDataClient'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataClientInformationMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// AllowIncomingAndOutgoingMessages - Retry to allow incoming gRPC calls and process outgoing calls
func (s *FenixTestDataGrpcServicesServer) AllowOrDisallowIncomingAndOutgoingMessages(ctx context.Context, allowOrDisallowIncomingAndOutgoingMessage *fenixTestDataSyncServerGrpcApi.AllowOrDisallowIncomingAndOutgoingMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "e7b5686b-6bdf-4218-ac05-e184acf7feb1",
	}).Debug("Incoming gRPC 'AllowOrDisallowIncomingAndOutgoingMessages'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "3961dda0-7af8-4746-a5b1-aa85a6fec479",
	}).Debug("Incoming gRPC 'AllowOrDisallowIncomingAndOutgoingMessages'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(allowOrDisallowIncomingAndOutgoingMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Set state depending on parameter so incoming and outgoing messages can be processed or not
	fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = allowOrDisallowIncomingAndOutgoingMessage.AllowInAndOutgoingMessages

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// RestartFenixServerProcesses - Restart Fenix TestData Server processes
func (s *FenixTestDataGrpcServicesServer) RestartFenixServerProcesses(ctx context.Context, emptyParameter *fenixTestDataSyncServerGrpcApi.EmptyParameter) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "8ecf6bfc-2ff6-4b50-b4be-faef443e0c0e",
	}).Debug("Incoming 'RestartFenixServerProcesses'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "7144c46e-c3d9-4d44-b4c7-20641c103a27",
	}).Debug("Outgoing 'RestartFenixServerProcesses'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(emptyParameter.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Reload Data from cloudDB into memoryDB
	_ = fenixTestDataSyncServerObject.loadTestDataFromCloudDB()
	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// Can't process ingoing or outgoing messages right now
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}
