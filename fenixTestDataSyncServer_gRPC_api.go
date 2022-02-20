package main

import (
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// *********************************************************************

// AreYouAlive :
// Fenix client can check if Fenix Testdata sync server is alive with this service
func (s *FenixTestDataGrpcServicesServer) AreYouAlive(_ context.Context, emptyParameter *fenixTestDataSyncServerGrpcApi.EmptyParameter) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
	}).Debug("Incoming gRPC 'AreYouAlive'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84isClientUsingCorrectTestDataProtoFileVersion",
	}).Debug("Outgoing gRPC 'AreYouAlive'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion("666", emptyParameter.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: "I'am Fenix TestDataSyncServer and I'm alive, and my time is " + fenixSyncShared.GenerateDatetimeTimeStampForDB()}, nil
}

// *********************************************************************

// SendMerkleHash :
// Fenix client can send TestData MerkleHash to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleHash(_ context.Context, merkleHashMessage *fenixTestDataSyncServerGrpcApi.MerkleHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "a55f9c82-1d74-44a5-8662-058b8bc9e48f",
	}).Debug("Incoming gRPC 'SendMerkleHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "27fb45fe-3266-41aa-a6af-958513977e28",
	}).Debug("Outgoing gRPC 'SendMerkleHash'")

	// Get calling client
	callingClientUuid := merkleHashMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleHashMessage.GetTestDataClientUuid(), merkleHashMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not the correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(merkleHashMessage.GetTestDataClientUuid())
	if returnMessage != nil {
		// Client in unknown
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Verify that MerkleFilterPath is hashed correctly
	returnMessage = fenixTestDataSyncServerObject.isClientsMerklePathCorrectlyHashed(callingClientUuid, merkleHashMessage.MerkleFilter, merkleHashMessage.MerkleFilterHash)
	if returnMessage != nil {
		// Wrongly hashed MerklePathNo processing of incoming messages
		return returnMessage, nil
	}

	// Compare current server- and client MerkleHash and MerklePathFilter
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientUuid)
	currentServerMerkleFilterHash := fenixTestDataSyncServerObject.getCurrentMerkleFilterPathHashForServer(callingClientUuid)

	//  if different in MerkleHash or MerkleFilterHash, then ask client for MerkleTree
	if currentServerMerkleHash != merkleHashMessage.MerkleHash || currentServerMerkleFilterHash != merkleHashMessage.MerkleFilterHash {

		// Save the MerkleHash
		_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForClient(callingClientUuid, merkleHashMessage.MerkleHash)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "d9c676ca-11a1-4007-8182-2f54834013a5",
		}).Debug("Saved the MerkleHash for Client: " + callingClientUuid)

		// Save the MerklePath
		_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathForClient(callingClientUuid, merkleHashMessage.MerkleFilter)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "d296a28c-4ad4-4dc5-9d2e-90230e55b4e2",
		}).Debug("Saved the MerkleFilterPath for Client: " + callingClientUuid)

		// Save the MerkleFilterHash
		_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForClient(callingClientUuid, merkleHashMessage.MerkleFilterHash)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "2ac5988d-28f0-4e39-aa59-729608bdfb0c",
		}).Debug("Saved the MerkleFilterPathHash for Client: " + callingClientUuid)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("There is different in MerkleHash or/and MerkleFilterHash, so ask client for MerkleTree for Client: " + callingClientUuid)

		defer fenixTestDataSyncServerObject.AskClientToSendMerkleTree(callingClientUuid)
	} else {

		// MerkleHash and MerkleFilterHash is already  in memDB, so just return that everthing is 'OK'
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e09b2a3f-82f4-43a4-8ebd-c169557714e7",
		}).Debug("MerkleHash and MerkleFilterHash is already  in memDB, so just return that everything is 'OK' for Client: " + callingClientUuid)

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************

// SendMerkleTree :
// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleTree(_ context.Context, merkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cffc25f0-b0e6-407a-942a-71fc74f831ac",
	}).Debug("Incoming gRPC 'SendMerkleTree'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "61e2c28d-b091-442a-b7f8-d2502d9547cf",
	}).Debug("Outgoing gRPC 'SendMerkleTree'")

	// Get calling client
	callingClientUuid := merkleTreeMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleTreeMessage.GetTestDataClientUuid(), merkleTreeMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(callingClientUuid)
	if returnMessage != nil {
		// Client in unknown
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

	// Verify MerkleTree
	clientsMerkleRootHash := fenixSyncShared.ExtractMerkleRootHashFromMerkleTree(merkleTreeAsDataFrame)
	recalculatedMerkleRootHash := fenixSyncShared.CalculateMerkleHashFromMerkleTree(merkleTreeAsDataFrame)

	// Something is wrong with clients hash computation
	if clientsMerkleRootHash != recalculatedMerkleRootHash {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Info("There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes")

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: false, Comments: "There is something wrong with Hash computation. Expected: '" + recalculatedMerkleRootHash + "' as MerkleRoot based on MerkleTree-nodes"}, nil
	}

	currentClientMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForClient(callingClientUuid)

	// Verify that Client Sent MerkleHash, in a previous message, matches the MerkleHash from MerkleTable
	if currentClientMerkleHash != recalculatedMerkleRootHash {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "a86f766a-c9ae-4698-abf2-b10a0f5ec475",
		}).Info("Client hasn't sent a MerkleHash message previous to MerkleTree message, Client: " + callingClientUuid)

		// Ask Client to send over the MerkleHash
		fenixTestDataSyncServerObject.AskClientToSendMerkleHash(callingClientUuid)

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: false, Comments: "Client hasn't sent a MerkleHash message previous to MerkleTree message"}, nil
	}

	// Compare current server- and client merklehash
	currentServerMerkleHash := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientUuid)

	//  if different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got
	if currentServerMerkleHash != recalculatedMerkleRootHash {

		// Convert incoming gRPC-MerkleTree into memoryDB-object for MerkleTreeNodes
		memDBMerkleTreeNodes := fenixTestDataSyncServerObject.convertGrpcMerkleTreeNodesIntoMemDBMerkleTreeNodes(recalculatedMerkleRootHash, merkleTreeMessage)

		// Save the MerkleHash and memDBMerkleTreeNodes
		_ = fenixTestDataSyncServerObject.saveCurrentMerkleTreeNodesForClient(callingClientUuid, memDBMerkleTreeNodes)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "cd30f2ae-6f79-4a0a-a8d8-a78d32dd6c71",
		}).Debug("Saved MerkleTree for Client: " + callingClientUuid)

		// If Server don't have any previous MerkleHash(and then no rows) then ask for all TestDataRows
		if currentServerMerkleHash == "" {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"id": "e011e854-7854-425f-9592-dcfc785203cf",
			}).Debug("No previous MerkleHash in DB then ask for all TestDataRow (MerkleTree was different) Client: " + callingClientUuid)

			defer fenixTestDataSyncServerObject.AskClientToSendAllTestDataRows(callingClientUuid)
		} else {
			// Remove TestDataRows that is not represented in current client MerkleTree
			fenixTestDataSyncServerObject.removeTestDataRowItemsInMemoryDBThatIsNotRepresentedInClientsMerkleTree(callingClientUuid)

			// Ask for the rows that is missing
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"id": "f0aa86c1-0179-4cb4-a891-60d55fdce84c",
			}).Debug("Different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got. Client: " + callingClientUuid)

			defer fenixTestDataSyncServerObject.AskClientToSendTestDataRows(callingClientUuid)
		}
	} else {

		// MerkleTree is the same, 'nothing to see'
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "94273c83-cfe4-415d-96bb-2890ea917116",
		}).Debug("MerkleHash from MerkleTree is the same as in memroyDB for Client: " + callingClientUuid)

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************

// SendTestDataHeaderHash :
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaderHash(_ context.Context, testDataHeaderHashMessageMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c8e72834-338c-48be-885c-f083fe7951a6",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b3576bdc-ac11-44cd-9fa2-cd7065f1253d",
	}).Debug("Outgoing gRPC 'SendTestDataHeaderHash'")

	// Get calling client
	callingClientUuid := testDataHeaderHashMessageMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderHashMessageMessage.GetTestDataClientUuid(), testDataHeaderHashMessageMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(callingClientUuid)
	if returnMessage != nil {
		// Client in unknown
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

	// Check if Server Header Hash is the same as received Client HeaderHash
	serverHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashForServer(callingClientUuid)
	if serverHeaderHash != clientHeaderHash {

		// Save the message
		_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientUuid, clientHeaderHash)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "2cdb17e4-95c8-4d56-b2cd-7f4c8829c735",
		}).Debug("Saved Header hash to DB for Client: " + callingClientUuid)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Debug("Server Header hash is not the same as Client Header Hash, Ask Client for all Headers, Client: " + callingClientUuid)

		// Ask Client for all Headers
		defer fenixTestDataSyncServerObject.AskClientToSendTestDataHeaders(callingClientUuid)

	}

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************

// SendTestDataHeaders :
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaders(_ context.Context, testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aee48999-12ad-4bb7-bc8a-96b62a8eeedf",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "ca0b58a8-6d56-4392-8751-45906670e86b",
	}).Debug("Outgoing gRPC 'SendTestDataHeaders'")

	// Get calling client
	callingClientUuid := testDataHeaderMessage.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderMessage.GetTestDataClientUuid(), testDataHeaderMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(callingClientUuid)
	if returnMessage != nil {
		// Client in unknown
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	// Verify that HeaderItemHash is correct calculated
	returnMessage = fenixTestDataSyncServerObject.verifyThatHeaderItemsHashIsCorrectCalculated(*testDataHeaderMessage)
	if returnMessage != nil {
		// HeaderItemsHash not correct calculated
		return returnMessage, nil
	}

	// Convert gRPC-message into other 'format'
	headerHash, headerItems := fenixTestDataSyncServerObject.convertgRpcHeaderMessageToStringArray(*testDataHeaderMessage)

	// Get current HeaderHash for the Client
	currentClientHeaderHash := fenixTestDataSyncServerObject.getCurrentHeaderHashForClient(callingClientUuid)

	// Check if HeaderData already is saved
	if currentClientHeaderHash != headerHash {

		// Save the message
		_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientUuid, headerHash)
		//_ = fenixTestDataSyncServerObject.saveCurrentHeadersForClient(callingClientUuid, headerItems)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "302ecd09-68a1-42df-ae63-7e4483ea62e1",
		}).Debug("Saved Header hash to DB for client: " + callingClientUuid)

		// Replace Server version of Headers with Client version of Headers
		_ = fenixTestDataSyncServerObject.moveCurrentHeaderDataFromClientToServer(callingClientUuid)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "6c90bc16-890a-402e-91b1-68fc060986c6",
		}).Debug("Moved HeaderHash from Client to Server, for client: " + callingClientUuid)

		// The Headers are change, so we must ask client to send the MerkleHash
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "ab857f2e-95dc-46d3-b065-71cdd38ea97b",
		}).Debug("The Headers are change, so we must ask client to send the MerkleHash; Client: " + callingClientUuid)

		fenixTestDataSyncServerObject.AskClientToSendMerkleHash(callingClientUuid)

	}

	// Save the Headers message --**** Not the solution I want because it saves it evan if nothing is changed ***
	_ = fenixTestDataSyncServerObject.saveCurrentHeadersForClient(callingClientUuid, headerItems)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1c99aec3-48b0-4ad2-be98-ec3dda32c4bd",
	}).Debug("Saved Headers data to DB for client: " + callingClientUuid)

	// Replace Server version of Headers with Client version of Headers
	_ = fenixTestDataSyncServerObject.moveCurrentHeaderDataFromClientToServer(callingClientUuid)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "6c90bc16-890a-402e-91b1-68fc060986c6",
	}).Debug("Moved Headers data from Client to Server, for client: " + callingClientUuid)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************

// SendTestDataRows :
// Fenix client can send TestData rows to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataRows(_ context.Context, testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2b1c8752-eb84-4c15-b8a7-22e2464e5168",
	}).Debug("Incoming gRPC 'SendTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "755e8b4f-f184-4277-ad41-e041714c2ca8",
	}).Debug("Outgoing gRPC 'SendTestDataRows'")

	// Get calling client
	callingClientUuid := testdataRowsMessages.TestDataClientUuid

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(
		testdataRowsMessages.TestDataClientUuid,
		testdataRowsMessages.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(callingClientUuid)
	if returnMessage != nil {
		// Client in unknown
		return returnMessage, nil
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return returnMessage, nil
	}

	/*
		// Convert proto-message for rows into Dataframe object
		newRowsAsDataFrame, returnMessage := fenixTestDataSyncServerObject.convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages)

		// When row-hash is wrong calculated from client then respond that back to client
		if returnMessage != nil {
			return returnMessage, nil
		}
	*/

	// Convert gRPC-RowsMessage into cloudDBTestDataRowItems-message
	cloudDBTestDataRowItemsMessage := fenixTestDataSyncServerObject.convertgRpcTestDataRowsMessageToCloudDBTestDataRowItems(
		testdataRowsMessages)

	currentTestDataHeaders := fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientUuid)

	// If there are no headers in Database then Ask client for HeaderHash
	if len(currentTestDataHeaders) == 0 {
		fenixTestDataSyncServerObject.AskClientToSendTestDataHeaderHash(callingClientUuid)
		currentTestDataHeaders = fenixTestDataSyncServerObject.getCurrentHeadersForClient(callingClientUuid)

	}

	// Concatenate with current Server data
	concatenatedTestDataRows := fenixTestDataSyncServerObject.concatenateWithCurrentServerTestData(
		callingClientUuid,
		cloudDBTestDataRowItemsMessage)

	// Get calling Client's MerkleFilterPath
	merkleFilterPath := fenixTestDataSyncServerObject.getCurrentMerkleFilterPathForClient(callingClientUuid)

	// Generate RowHashToMerkleChildNodeHashMap for TestData Rows and return map of type; "MAP[<RowHash>]=<MerkleChildNodeHash>" - 'MAP[string]string')
	rowHashToLeafNodeHashMap := fenixTestDataSyncServerObject.generateRowHashToMerkleChildNodeHashMap(
		concatenatedTestDataRows,
		fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForClient(callingClientUuid))
	//TODO debugga hit och se varför leaf node hash saknas i testdata rows
	// Add MerkleLeafNodeHashes to TestDataRowItems
	testDataRowsIncludingLeafNodeHashes := fenixTestDataSyncServerObject.addMerkleLeafNodeHashesToTestDataRowItems(
		concatenatedTestDataRows,
		rowHashToLeafNodeHashMap)

	// Convert the memoryDB-object for TestDataRows into a DataFrame
	testDataRowsIncludingLeafNodeHashesAsDataFrame, returnMessage := fenixTestDataSyncServerObject.convertCloudDBTestDataRowItemsMessageToDataFrame(
		testDataRowsIncludingLeafNodeHashes)
	if returnMessage != nil {
		// Something got wrong
		return returnMessage, nil
	}

	// Recreate MerkleHash from All testdata rows, both existing rows for Server and new from Client
	computedMerkleHash, _, _ := fenixSyncShared.CreateMerkleTreeFromDataFrame(
		testDataRowsIncludingLeafNodeHashesAsDataFrame,
		merkleFilterPath)

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
			Comments:   "MerkleRoot hash is not the same as sent to server. Got " + clientMerkleHash + ", but recalculated to " + computedMerkleHash + " from testdata",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e9b43ac4-ed89-41ef-9762-bc7406633906",
		}).Info("MerkleRoot hash is not the same as sent to server. Got " + clientMerkleHash + ", but recalculated to " + computedMerkleHash + " from testdata  Client: " + callingClientUuid)

		// respond back to client when it used wrong proto-file
		return returnMessage, nil

	} else {

		// Convert gRPC-RowsMessage into cloudDBTestDataRowItems-message
		//cloudDBTestDataRowItemMessage := fenixTestDataSyncServerObject.convertgRpcTestDataRowsMessageToCloudDBTestDataRowItems(
		//	testdataRowsMessages)

		// Save TestDataRows to MemoryDB
		_ = fenixTestDataSyncServerObject.saveCurrentTestDataRowItemsForClient(
			callingClientUuid, testDataRowsIncludingLeafNodeHashes) //cloudDBTestDataRowItemMessage)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "9aa8379e-5d5c-4eb4-9b18-d52da4291795",
		}).Debug("The TestDataRows were saved for Client: " + callingClientUuid)

		// Move Client-data to Server-data in MemoryDB for client
		success := fenixTestDataSyncServerObject.moveCurrentTestDataAndMerkleTreeFromClientToServer(callingClientUuid)

		if success == true {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"id": "c7728251-ec36-4cca-a529-1f8db67acc96",
			}).Debug("The TestDataRows were copied from Client to Server for Client: " + callingClientUuid)
		} else {
			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   "Something went wrong",
				ErrorCodes: nil,
			}

			return returnMessage, nil
		}
	}

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "78be9c02-5862-4dbb-bf3b-24a8d40aeb1d",
	}).Debug("The TestDataRows are the same as the Server already have for Client: " + callingClientUuid)

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************

// RegisterTestDataClient :
// Fenix client can register itself with the Fenix Testdata sync server
func (s *FenixTestDataGrpcServicesServer) RegisterTestDataClient(
	_ context.Context,
	testDataClientInformationMessage *fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5133b80b-6f3a-4562-9e62-1b3ceb169cc1",
	}).Debug("Incoming gRPC 'RegisterTestDataClient'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "316dcd7e-2229-4a82-b15b-0f808c2dd8aa",
	}).Debug("Outgoing gRPC 'RegisterTestDataClient'")

	// Check if Client is using correct proto files version
	returnMessage := fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataClientInformationMessage.TestDataClientUuid, testDataClientInformationMessage.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return returnMessage, nil
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(testDataClientInformationMessage.TestDataClientUuid)
	if returnMessage != nil {
		// Client in unknown
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
