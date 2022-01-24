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
	}).Debug("Incoming 'AreYouAlive'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84",
	}).Debug("Outgoing 'AreYouAlive'")

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(emptyParameter.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: "I'am Fenix TestDataSyncServer and I'm alive"}, nil
}

// *********************************************************************
// Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) SendMerkleHash(ctx context.Context, merkleHashMessage *fenixTestDataSyncServerGrpcApi.MerkleHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "a55f9c82-1d74-44a5-8662-058b8bc9e48f",
	}).Debug("Incoming 'SendMerkleHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "27fb45fe-3266-41aa-a6af-958513977e28",
	}).Debug("Outgoing 'SendMerkleHash'")

	// Get calling client
	callingClientGuid := merkleHashMessage.TestDataClientGuid

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(merkleHashMessage.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForClient(*merkleHashMessage)

	// Compare current server- and client merklehash
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientGuid)

	//  if different in MerkleHash then ask client for MerkleTree
	if currentServerMerkleHash != merkleHashMessage.MerkleHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("There is different in MerkleHash so ask client for MerkleTree for Client: " + callingClientGuid)

		defer fenixTestDataSyncServerObject.AskClientToSendMerkleTree(callingClientGuid)
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleTree(ctx context.Context, merkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cffc25f0-b0e6-407a-942a-71fc74f831ac",
	}).Debug("Incoming 'SendMerkleTree'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "61e2c28d-b091-442a-b7f8-d2502d9547cf",
	}).Debug("Outgoing 'SendMerkleTree'")

	// Get calling client
	callingClientGuid := merkleTreeMessage.TestDataClientGuid

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(merkleTreeMessage.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	// Convert the merkleTree into a DataFrame object
	merkleTreeAsDataFrame := fenixTestDataSyncServerObject.convertgRpcMerkleTreeMessageToDataframe(*merkleTreeMessage)

	// Verify MerkleTree
	clientsMerkleRootHash := common_config.ExtractMerkleRootHashFromMerkleTree(merkleTreeAsDataFrame)
	recalculatedMerkleRootHash := common_config.CalculateMerkleHashFromMerkleTree(merkleTreeAsDataFrame)

	// Something is wrong with clients hash computation
	if clientsMerkleRootHash != recalculatedMerkleRootHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes")

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: false, Comments: "There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes"}, nil
	}

	// Save the MerkleTree Dataframe message
	_ = fenixTestDataSyncServerObject.saveCurrentMerkleTreeForClient(callingClientGuid, merkleTreeAsDataFrame)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
	}).Debug("Saved MerkleTree for Client: " + callingClientGuid)

	// Compare current server- and client merklehash
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientGuid)

	//  if different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got
	if currentServerMerkleHash != recalculatedMerkleRootHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e011e854-7854-425f-9592-dcfc785203cf",
		}).Debug("Different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got. Client: " + callingClientGuid)

		defer fenixTestDataSyncServerObject.AskClientToSendAllTestDataRows(callingClientGuid)
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// *********************************************************************
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaderHash(ctx context.Context, testDataHeaderHashMessageMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c8e72834-338c-48be-885c-f083fe7951a6",
	}).Debug("Incoming 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b3576bdc-ac11-44cd-9fa2-cd7065f1253d",
	}).Debug("Outgoing 'SendTestDataHeaders'")

	// Get calling client
	callingClientGuid := testDataHeaderHashMessageMessage.TestDataClientGuid

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(testDataHeaderHashMessageMessage.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	// Convert gRPC-message into other 'format'
	clientHeaderHash := testDataHeaderHashMessageMessage.HeadersHash

	// Get current client Header Hash
	currentClientHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashsForClient(callingClientGuid)

	// If Header Hash is already in DB then return OK
	if currentClientHeaderHash == clientHeaderHash {
		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientGuid, clientHeaderHash)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2cdb17e4-95c8-4d56-b2cd-7f4c8829c735",
	}).Debug("Saved Header hash to DB for client: " + callingClientGuid)

	// Check if Server Header Hash is the same as received Client HeaderHash
	serverHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashForServer(callingClientGuid)
	if serverHeaderHash != clientHeaderHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Debug("Server Header hash is not the same as Client Header Hash, Ask Client for all Headers, Client: " + callingClientGuid)

		// Ask Client for all Headers
		defer fenixTestDataSyncServerObject.AskClientToSendTestDataHeaders(callingClientGuid)

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil

}

// *********************************************************************
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaders(ctx context.Context, testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aee48999-12ad-4bb7-bc8a-96b62a8eeedf",
	}).Debug("Incoming 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "ca0b58a8-6d56-4392-8751-45906670e86b",
	}).Debug("Outgoing 'SendTestDataHeaders'")

	// Get calling client
	callingClientGuid := testDataHeaderMessage.TestDataClientGuid

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(testDataHeaderMessage.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	// Convert gRPC-message into other 'format'
	headerHash, headerItems := fenixTestDataSyncServerObject.convertgRpcHeaderMessageToStringArray(*testDataHeaderMessage)

	// Validate HeaderHash
	computedHeaderHash := common_config.HashValues(headerItems)
	if computedHeaderHash != headerHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Info("Header hash is not correct computed from Client. Expected '" + computedHeaderHash + "' as HeaderHash but got " + headerHash)

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: false, Comments: "Header hash is not correct computed. Expected '" + computedHeaderHash + "' as HeaderHash"}, nil
	}

	// Save the message
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientGuid, headerHash)
	_ = fenixTestDataSyncServerObject.saveCurrentHeadersForClient(callingClientGuid, headerItems)

	// Replace Server version of Headers with Client version of Headers
	_ = fenixTestDataSyncServerObject.moveCurrentHeadersFromClientToServer(callingClientGuid)
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "6c90bc16-890a-402e-91b1-68fc060986c6",
	}).Debug("Saved Header hash and Headers to DB for client: " + callingClientGuid)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil

}

// *********************************************************************
// Fenix client can send TestData rows to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataRows(ctx context.Context, testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2b1c8752-eb84-4c15-b8a7-22e2464e5168",
	}).Debug("Incoming 'SendTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "755e8b4f-f184-4277-ad41-e041714c2ca8",
	}).Debug("Outgoing 'SendTestDataRows'")

	// Get calling client
	callingClientGuid := testdataRowsMessages.TestDataClientGuid

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(testdataRowsMessages.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		// respond back to client when it used wrong proto-file
		return returnMessage, nil

	}

	// Convert proto-message for rows into Dataframe object
	newRowsAsDataFrame, returnMessage := fenixTestDataSyncServerObject.convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages)

	// When row-hash is wrong calculated from client then respond that back to client
	if returnMessage != nil {
		return returnMessage, nil
	}

	currentTestDataHeaders := fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientGuid)

	// If there are no headers in Database then Ask client for HeaderHash
	if len(currentTestDataHeaders) == 0 {
		fenixTestDataSyncServerObject.AskClientToSendTestDataHeaderHash(callingClientGuid)
		currentTestDataHeaders = fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientGuid)

	}

	// Concatenate with current Server data
	allRowsAsDataFrame := fenixTestDataSyncServerObject.concartenateWithCurrentServerTestData(callingClientGuid, newRowsAsDataFrame)

	// Recreate MerkleHash from All testdata rows, both existing rows for Server and new from Client
	computedMerkleHash, _ := common_config.CreateMerkleTreeFromDataFrame(allRowsAsDataFrame)

	//Compare 'computedMerkleHash' with MerkleHash from Client
	clientMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForClient(callingClientGuid)

	if computedMerkleHash != clientMerkleHash {
		if clientUseCorrectProtoFileVersion == false {
			// Computed MerkleHash is not the same as the one sent by the Client

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_MERKLEHASH_NOT_CORRECT_CALCULATED
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				Acknack:    false,
				Comments:   "MerklRoot hash is not the same as sent to server. Got " + clientMerkleHash + ", but recalculated to " + computedMerkleHash + " from testdata",
				ErrorCodes: errorCodes,
			}

			// respond back to client when it used wrong proto-file
			return returnMessage, nil

		}

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

// RegisterTestDataClient Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) RegisterTestDataClient(ctx context.Context, testDataClientInformationMessage *fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5133b80b-6f3a-4562-9e62-1b3ceb169cc1",
	}).Debug("Incoming 'RegisterTestDataClient'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "316dcd7e-2229-4a82-b15b-0f808c2dd8aa",
	}).Debug("Outgoing 'RegisterTestDataClient'")

	// Check if Client is using correct proto files version
	clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed := fenixTestDataSyncServerObject.isClientUsingCorrectProtoFileVersion(testDataClientInformationMessage.ProtoFileVersionUsedByClient)
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage := &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			Acknack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage, nil

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{Acknack: true, Comments: ""}, nil
}

/*
func (s *FenixTestDataGrpcServicesServer) mustEmbedUnimplementedFenixTestDataGrpcServicesServer() {
	//TODO implement me
	panic("implement me")
}


*/
