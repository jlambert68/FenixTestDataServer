package main

import (
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io"
	"log"
	"os"
)

// *********************************************************************

// AreYouAlive :
// Fenix client can check if Fenix Testdata sync server is alive with this service
func (s *FenixTestDataGrpcServicesServer) AreYouAlive(_ context.Context, emptyParameter *fenixTestDataSyncServerGrpcApi.EmptyParameter) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
	}).Debug("Incoming gRPC 'AreYouAlive'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "1762bfd6-d7f8-47ad-8651-1bc13b9bac32",
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

	// Set up slice of functions to be processed when leaving using ReversedDeferOrder, FIFO
	var deferFunctions []func()
	defer fenixTestDataSyncServerObject.reversedDefer(&deferFunctions)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "a55f9c82-1d74-44a5-8662-058b8bc9e48f",
	}).Debug("Incoming gRPC 'SendMerkleHash'")

	exitLogMessageDeferFunction := func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "27fb45fe-3266-41aa-a6af-958513977e28",
		}).Debug("Outgoing gRPC 'SendMerkleHash'")
	}
	deferFunctions = append(deferFunctions, exitLogMessageDeferFunction)

	// Get calling client
	callingClientUuid := merkleHashMessage.TestDataClientUuid

	// When leaving set the next allowed state
	var nextTestDataState = CurrenStateMerkleHash //Will be used if not changed, when stuff goes wrong
	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.changeTestDataState(callingClientUuid, &nextTestDataState)
	})

	// Verify that system is in expected State
	returnMessage, nextExpectedState := fenixTestDataSyncServerObject.isSystemInCorrectTestDataRowsState(callingClientUuid, CurrenStateMerkleHash)
	if returnMessage != nil {
		// System is in wrong state
		return returnMessage, nil
	}

	// Check if Client is using correct proto files version
	returnMessage = fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleHashMessage.GetTestDataClientUuid(), merkleHashMessage.ProtoFileVersionUsedByClient)
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

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id":                           "88add85d-f15e-4c89-97b3-30f291c45ed7",
		"currentServerMerkleHash":      currentServerMerkleHash,
		"merkleHashMessage.MerkleHash": merkleHashMessage.MerkleHash,
	}).Debug("Comparing incoming MerkleHash with stored MerkleHash for Client: " + callingClientUuid)

	//  if different in MerkleHash or MerkleFilterHFor LeafNodes the childHash will be calculated by using SHA256(NodeHash)ash, then ask client for MerkleTree
	if currentServerMerkleHash != merkleHashMessage.MerkleHash || currentServerMerkleFilterHash != merkleHashMessage.MerkleFilterHash {

		// Save the MerkleHash
		_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForClient(callingClientUuid, merkleHashMessage.MerkleHash)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "2987b74c-844c-4f70-8f9e-6ce0886857ab",
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

		// Add to ReversedDefer to call when leaving this function
		deferFunctions = append(deferFunctions, func() {
			fenixTestDataSyncServerObject.AskClientToSendMerkleTree(callingClientUuid)
		})
	} else {

		// MerkleHash and MerkleFilterHash is already  in memDB, so just return that everthing is 'OK'
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e09b2a3f-82f4-43a4-8ebd-c169557714e7",
		}).Debug("MerkleHash and MerkleFilterHash is already  in memDB, so just return that everything is 'OK' for Client: " + callingClientUuid)

		// No Change in State is needed
		nextTestDataState = CurrenStateMerkleHash

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

	}

	// When all processing went well set next TestDataState to be expected
	nextTestDataState = nextExpectedState

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************

// SendMerkleTree :
// Fenix client can send TestData MerkleTree to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendMerkleTree(_ context.Context, merkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	// Set up slice of functions to be processed when leaving using ReversedDeferOrder, FIFO
	var deferFunctions []func()
	defer fenixTestDataSyncServerObject.reversedDefer(&deferFunctions)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "cffc25f0-b0e6-407a-942a-71fc74f831ac",
	}).Debug("Incoming gRPC 'SendMerkleTree'")

	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "61e2c28d-b091-442a-b7f8-d2502d9547cf",
		}).Debug("Outgoing gRPC 'SendMerkleTree'")
	})

	// Get calling client
	callingClientUuid := merkleTreeMessage.TestDataClientUuid

	// When leaving set the next allowed state
	var nextTestDataState = CurrenStateMerkleHash //Will be used if not changed, when stuff goes wrong
	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.changeTestDataState(callingClientUuid, &nextTestDataState)
	})

	// Verify that system is in expected State
	returnMessage, nextExpectedState := fenixTestDataSyncServerObject.isSystemInCorrectTestDataRowsState(callingClientUuid, CurrenStateMerkleTree)
	if returnMessage != nil {
		// System is in wrong state
		return returnMessage, nil
	}

	// Check if Client is using correct proto files version
	returnMessage = fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(merkleTreeMessage.GetTestDataClientUuid(), merkleTreeMessage.ProtoFileVersionUsedByClient)
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

	// XXX
	f, err := os.Create("incomingMerkleTree.csv")
	if err != nil {
		log.Fatal(err)
	}
	merkleTreeAsDataFrame.WriteCSV(f)
	f.Close()

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

		// Load MerkleTree from CloadDB, to memDB, for server to be used later down in code
		_ = fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForServer(callingClientUuid)

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

			// Add to ReversedDefer to call when leaving this function
			deferFunctions = append(deferFunctions, func() {
				fenixTestDataSyncServerObject.AskClientToSendAllTestDataRows(callingClientUuid)
			})
		} else {
			// Remove TestDataRows that is not represented in current client MerkleTree
			fenixTestDataSyncServerObject.removeMerkleTreeNodeItemsInMemoryDBThatIsNotRepresentedInClientsNewMerkleTree(callingClientUuid)

			// Ask for the rows that is missing
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"id": "f0aa86c1-0179-4cb4-a891-60d55fdce84c",
			}).Debug("Different in MerkleHash(MerkleTree was different) then ask client for TestData-rows that the Server hasn't got. Client: " + callingClientUuid)

			// Add to ReversedDefer to call when leaving this function
			deferFunctions = append(deferFunctions, func() {
				fenixTestDataSyncServerObject.AskClientToSendTestDataRows(callingClientUuid)
			})
		}
	} else {

		// MerkleTree is the same, 'nothing to see'
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "94273c83-cfe4-415d-96bb-2890ea917116",
		}).Debug("MerkleHash from MerkleTree is the same as in memoryDB for Client: " + callingClientUuid)

		// When all processing went well set next TestDataState to be expected
		nextTestDataState = CurrenStateMerkleHash

		return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

	}

	// When all processing went well set next TestDataState to be expected
	nextTestDataState = nextExpectedState

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil
}

// *********************************************************************

// SendTestDataHeaderHash :
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaderHash(_ context.Context, testDataHeaderHashMessageMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderHashMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	// Set up slice of functions to be processed when leaving using ReversedDeferOrder, FIFO
	var deferFunctions []func()
	defer fenixTestDataSyncServerObject.reversedDefer(&deferFunctions)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c8e72834-338c-48be-885c-f083fe7951a6",
	}).Debug("Incoming gRPC 'SendTestDataHeaderHash'")

	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "6d5c99c6-8b73-4b17-9c01-474d4e12538c",
		}).Debug("Outgoing gRPC 'SendTestDataHeaderHash'")
	})

	// Get calling client
	callingClientUuid := testDataHeaderHashMessageMessage.TestDataClientUuid

	// When leaving set the next allowed state
	var nextTestDataState = CurrenStateTestDataHeaderHash //Will be used if not changed, when stuff goes wrong
	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.changeTestDataHeaderState(callingClientUuid, &nextTestDataState)
	})

	// Verify that system is in expected State
	returnMessage, nextExpectedState := fenixTestDataSyncServerObject.isSystemInCorrectTestDataHeaderState(callingClientUuid, CurrenStateTestDataHeaderHash)
	if returnMessage != nil {
		// System is in wrong state
		return returnMessage, nil
	}

	// Check if Client is using correct proto files version
	returnMessage = fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderHashMessageMessage.GetTestDataClientUuid(), testDataHeaderHashMessageMessage.ProtoFileVersionUsedByClient)
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

	// Get current Server HeaderHash
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
		}).Debug("Saved HeaderHash to memoryDB-client for Client: " + callingClientUuid)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e4efabb7-eec6-4bb6-bbc2-f3447e14c15f",
		}).Debug("Server HeaderHash is not the same as Client HeaderHash, Ask Client for all Headers, Client: " + callingClientUuid)

		// Ask Client for all Headers
		// Add to ReversedDefer to call when leaving this function
		deferFunctions = append(deferFunctions, func() {
			fenixTestDataSyncServerObject.AskClientToSendTestDataHeaders(callingClientUuid)
		})

	}

	// When all processing went well set next TestDataHeaderState to be expected
	nextTestDataState = nextExpectedState

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************

// SendTestDataHeaders :
// Fenix client can send TestDataHeaders to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataHeaders(_ context.Context, testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	// Set up slice of functions to be processed when leaving using ReversedDeferOrder, FIFO
	var deferFunctions []func()
	defer fenixTestDataSyncServerObject.reversedDefer(&deferFunctions)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aee48999-12ad-4bb7-bc8a-96b62a8eeedf",
	}).Debug("Incoming gRPC 'SendTestDataHeaders'")

	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "ee0d6762-4db2-4432-8b0e-fd5ddba74f64",
		}).Debug("Outgoing gRPC 'SendTestDataHeaders'")
	})

	// Get calling client
	callingClientUuid := testDataHeaderMessage.TestDataClientUuid

	// When leaving set the next allowed state
	var nextTestDataState = CurrenStateTestDataHeaderHash //Will be used if not changed, when stuff goes wrong
	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.changeTestDataHeaderState(callingClientUuid, &nextTestDataState)
	})

	// Verify that system is in expected State
	returnMessage, nextExpectedState := fenixTestDataSyncServerObject.isSystemInCorrectTestDataHeaderState(callingClientUuid, CurrenStateTestDataHeaders)
	if returnMessage != nil {
		// System is in wrong state
		return returnMessage, nil
	}

	// Check if Client is using correct proto files version
	returnMessage = fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(testDataHeaderMessage.GetTestDataClientUuid(), testDataHeaderMessage.ProtoFileVersionUsedByClient)
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

		// Save the HeaderHash to memoryDB
		_ = fenixTestDataSyncServerObject.saveCurrentHeaderHashForClient(callingClientUuid, headerHash)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e7107232-d312-47e8-8e39-ec74d4ef9dd5",
		}).Debug("Saved HeaderHash to memoryDB-client for client: " + callingClientUuid)

		// Save testDataHeaderMessage to memoryDB
		_ = fenixTestDataSyncServerObject.saveCurrentHeaderMessageDataForClient(*testDataHeaderMessage)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e7107232-d312-47e8-8e39-ec74d4ef9dd5",
		}).Debug("Saved HeaderHash to memoryDB-client for client: " + callingClientUuid)

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

	// When all processing went well set next TestDataHeaderState to be expected
	nextTestDataState = nextExpectedState

	return &fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""}, nil

}

// *********************************************************************

// SendTestDataRows :
// Fenix client can send TestData rows to Fenix Testdata sync server with this service
func (s *FenixTestDataGrpcServicesServer) SendTestDataRows(stream fenixTestDataSyncServerGrpcApi.FenixTestDataGrpcServices_SendTestDataRowsServer) error { //(_ context.Context, testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (*fenixTestDataSyncServerGrpcApi.AckNackResponse, error) {

	// Set up slice of functions to be processed when leaving using ReversedDeferOrder, FIFO
	var deferFunctions []func()
	defer fenixTestDataSyncServerObject.reversedDefer(&deferFunctions)

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2b1c8752-eb84-4c15-b8a7-22e2464e5168",
	}).Debug("Incoming gRPC 'SendTestDataRows'")

	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "974ff435-8653-4ee7-9aeb-fe5c8bfd1716",
		}).Debug("Outgoing gRPC 'SendTestDataRows'")
	})

	// Container to store all messages before process them
	var testdataRowsMessagesStreamContainer []*fenixTestDataSyncServerGrpcApi.TestdataRowsMessages
	var testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages

	var err error
	var returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse

	// Retrieve stream from Client
	for {
		// Receive message and add i to 'testdataRowsMessagesStreamContainer'
		testdataRowsMessages, err = stream.Recv()
		testdataRowsMessagesStreamContainer = append(testdataRowsMessagesStreamContainer, testdataRowsMessages)

		// When no more messages is received then continue
		if err == io.EOF {
			break
		}
	}

	testdataRowsMessages = testdataRowsMessagesStreamContainer[0]

	// Get calling client
	callingClientUuid := testdataRowsMessages.TestDataClientUuid

	// When leaving set the next allowed state
	var nextTestDataState = CurrenStateMerkleHash //Will be used if not changed, when stuff goes wrong
	deferFunctions = append(deferFunctions, func() {
		fenixTestDataSyncServerObject.changeTestDataState(callingClientUuid, &nextTestDataState)
	})

	// Verify that system is in expected State
	returnMessage, nextExpectedState := fenixTestDataSyncServerObject.isSystemInCorrectTestDataRowsState(callingClientUuid, CurrenStateTestData)
	if returnMessage != nil {
		// System is in wrong state
		return stream.SendAndClose(returnMessage)
	}

	// Check if Client is using correct proto files version
	returnMessage = fenixTestDataSyncServerObject.isClientUsingCorrectTestDataProtoFileVersion(
		testdataRowsMessages.TestDataClientUuid,
		testdataRowsMessages.ProtoFileVersionUsedByClient)
	if returnMessage != nil {
		// Not correct proto-file version is used
		return stream.SendAndClose(returnMessage)
	}

	// Verify that Client is known to Server
	returnMessage = fenixTestDataSyncServerObject.isClientKnownToServer(callingClientUuid)
	if returnMessage != nil {
		// Client in unknown
		return stream.SendAndClose(returnMessage)
	}

	// Check if TestData server should process incoming messages
	returnMessage = fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessage != nil {
		// No processing of incoming messages
		return stream.SendAndClose(returnMessage)
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
	// TODO Denna är bortkommenterad då ingen ihopslagning av TestDataRader från Server +  Nya Rader bör göras då
	//concatenatedTestDataRows := fenixTestDataSyncServerObject.concatenateWithCurrentServerTestData(
	//	callingClientUuid,
	//	cloudDBTestDataRowItemsMessage)

	// Generate RowHashToMerkleChildNodeHashMap for TestData Rows and return map of type; "MAP[<RowHash>]=<MerkleChildNodeHash>" - 'MAP[string]string')
	rowHashToLeafNodeHashMap := fenixTestDataSyncServerObject.generateRowHashToNodeHashMap(
		cloudDBTestDataRowItemsMessage, //concatenatedTestDataRows,
		fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForClient(callingClientUuid))

	testDataRowsIncludingLeafNodeHashes := fenixTestDataSyncServerObject.addMerkleLeafNodeHashesToTestDataRowItems(
		cloudDBTestDataRowItemsMessage, //concatenatedTestDataRows,
		rowHashToLeafNodeHashMap)

	// Convert the memoryDB-object for TestDataRows into a DataFrame
	//testDataRowsIncludingLeafNodeHashesAsDataFrame, returnMessage := fenixTestDataSyncServerObject.convertCloudDBTestDataRowItemsMessageToDataFrame(
	//	testDataRowsIncludingLeafNodeHashes)
	//if returnMessage != nil {
	// Something got wrong
	//	return stream.SendAndClose(returnMessage)
	//}

	// Get the Server MerkleTree
	//currentMerkleTreeNodesForServer := fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForServer(callingClientUuid)
	/*
		//
		1) UseClientMerkleTree
		istället
		för
		ServerMerkleTree
		och
		bara
		spara
		ner
		2) Verifiera
		att
		Inkommande
		rader
		genererar
		de
		LeadNodeHashes
		som
		efterfrågas
	*/
	// Get Server and Clients version of the MerkleTree
	serverCopyMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForServer(callingClientUuid)
	clientsNewMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForClient(callingClientUuid)

	var serverHasPreviousDataForCLient bool
	if len(serverCopyMerkleTree) == 0 {
		serverHasPreviousDataForCLient = false
	} else {
		serverHasPreviousDataForCLient = true
	}

	// Extract the LeafNodes from ServerMerkleTree
	// If Server has no previous data from client then set to nil
	var serverCopyMerkleTreeLeafNodeItems []cloudDBTestDataMerkleTreeStruct
	if serverHasPreviousDataForCLient == true {
		serverCopyMerkleTreeLeafNodeItems = fenixTestDataSyncServerObject.extractLeafNodeItemsFromMerkleTree(serverCopyMerkleTree, serverHasPreviousDataForCLient)

	} else {
		serverCopyMerkleTreeLeafNodeItems = nil
	}

	// Extract the LeafNodes from ClientMerkleTree
	clientsNewMerkleTreeLeafNodeItems := fenixTestDataSyncServerObject.extractLeafNodeItemsFromMerkleTree(clientsNewMerkleTree, true)

	// Extract LeafNodeItems that is in Client but not i Server, will use 'Not(Server) Intersecting Client'
	newMerkleTreeLeafNodeItemsFromClient := fenixTestDataSyncServerObject.notAIntersectingBOnMerkleTreeItems(serverCopyMerkleTreeLeafNodeItems, clientsNewMerkleTreeLeafNodeItems)

	// Verify that New Rows from client will generate LeafNodesHashes found in 'newMerkleTreeLeafNodeItemsFromClient'
	returnMessage = fenixTestDataSyncServerObject.verifyThatRowHashesMatchesLeafNodeHashes(callingClientUuid, cloudDBTestDataRowItemsMessage, newMerkleTreeLeafNodeItemsFromClient)
	if returnMessage != nil {
		// RowHashes didn't generate expected MerkleLeafNodeHashes
		return stream.SendAndClose(returnMessage)
	}

	requestedLeafNodeNames := fenixTestDataSyncServerObject.getCurrentRequestedMerkleNodeNamesFromClient(callingClientUuid)

	// Verify that all requested NodesNames is found among TestDataRows
	returnMessage = fenixTestDataSyncServerObject.verifyThatRequestedLeafNodeNamesAreFoundAmongReceivedTestDataRows(callingClientUuid, cloudDBTestDataRowItemsMessage, requestedLeafNodeNames)
	if returnMessage != nil {
		// RowHashes didn't generate expected MerkleLeafNodeHashes
		return stream.SendAndClose(returnMessage)
	}

	/*
			// Extract all NodeNames to request from client
			merkleNodeNamesToRequest := missedPathsToRetrieveFromClient(serverCopyMerkleTree, clientsNewMerkleTree)

			// Create LeafNodes from RowItems
			merkleTreeLeafNodesItems := fenixTestDataSyncServerObject.extractLeafNodeItemsFromMerkleTree(testDataRowsIncludingLeafNodeHashes)

			//
			3) Läs upp alla LeadNodeNames & LeafNodeHAshes för att säkerställa att [DB - "de som ska bort" + "De nya"] är samma som i Nya ClientMerkleTree
		// Convert MerkleTreeNodes into Dataframe-object
			currentMerkleTreeNodesForServerAsDataFrame := fenixTestDataSyncServerObject.convertMemDBMerkleTreeMessageToDataframe(currentMerkleTreeNodesForServer)
			f, err := os.Create("currentMerkleTreeNodesForServerAsDataFrame.csv")
			if err != nil {
				log.Fatal(err)
			}
			currentMerkleTreeNodesForServerAsDataFrame.WriteCSV(f)
			f.Close()

	*/

	// Create MerkleHash and MerkleTree from LeafNodeHashes
	//computedMerkleHash := fenixSyncShared.CalculateMerkleHashFromMerkleTree(currentMerkleTreeNodesForServerAsDataFrame)

	// Recreate MerkleHash from All testdata rows, both existing rows for Server and new from Client
	//computedMerkleHash, _, _ := fenixSyncShared.CreateMerkleTreeFromDataFrame(
	//	testDataRowsIncludingLeafNodeHashesAsDataFrame,
	//	merkleFilterPath)
	/*
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

			return stream.SendAndClose(returnMessage)

		} else {
	*/
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
			"id": "601150cd-856f-49db-acc5-779004e7701b",
		}).Debug("The TestDataRows were copied from Client to Server for Client: " + callingClientUuid)

		// Verify that LeafNodes in DB recreates MerkleHash
		returnMessage = fenixTestDataSyncServerObject.verifyThatMerkleLeafNodeHashesReCreatesMerkleHash(callingClientUuid)
		if returnMessage != nil {
			// MerkleHash not correct recreated
			return stream.SendAndClose(returnMessage)
		}
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "b68303d6-6973-452b-9b8b-adaf724e5651",
		}).Debug("Verified OK that LeafNodes in CloudDB recreated MerkleHash, Client: " + callingClientUuid)

		// When all processing went well set next TestDataState to be expected
		nextTestDataState = nextExpectedState

		return stream.SendAndClose(&fenixTestDataSyncServerGrpcApi.AckNackResponse{AckNack: true, Comments: ""})

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

		return stream.SendAndClose(returnMessage)
	}
}

//}

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
