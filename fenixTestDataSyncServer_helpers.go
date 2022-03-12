package main

import (
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
)

// ********************************************************************************************************************
// Check if there is a temporary stop in processing incoming or outgoing messages
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isThereATemporaryStopInProcessingInOrOutgoingMessages() (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// Check if there is a temporary stop in processing in- or outgoing messages
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "04df8862-3e76-45aa-9cee-9858c348e18d",
		}).Debug("There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server")

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if the system is in correct TestDataState and return next expected state
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isSystemInCorrectTestDataRowsState(currentCallingUser string, expectedTestDataState executionStateTestDataTypeType) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse, nextExpectedState executionStateTestDataTypeType) {

	// Verify that the system is in correct TestDataState
	if expectedTestDataState != fenixTestDataSyncServerObject.currentTestDataState {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TESTDATASERVER_INTERNAL_STATE_NOT_MATCHING_GRPC_CALL
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Expected TestData-State for TestDataServer is " + strconv.FormatInt(int64(expectedTestDataState), 10) + " but current State is " + strconv.FormatInt(int64(fenixTestDataSyncServerObject.currentTestDataState), 10),
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e903390a-421a-4d16-976d-08a464424385",
		}).Error("Expected TestData-State for TestDataServer is " + strconv.FormatInt(int64(expectedTestDataState), 10) + " but current State is " + strconv.FormatInt(int64(fenixTestDataSyncServerObject.currentTestDataState), 10) + " for Client: " + currentCallingUser)

		return returnMessage, CurrenStateMerkleHash

	} else {

		nextExpectedState = nextTestDataStateMap[expectedTestDataState]

		return nil, nextExpectedState
	}

}

// ********************************************************************************************************************
// Change TestData-state
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) changeTestDataState(currentCallingUser string, nextTestDataState *executionStateTestDataTypeType) {

	//TODO Change from .Info to .Debug
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id":                "aada4adf-8795-4189-a720-518cf68354e5",
		"Old TestDataState": fenixTestDataSyncServerObject.currentTestDataState,
		"New TestDataState": *nextTestDataState,
	}).Info("Change 'ChangeTestDataState' for Client: " + currentCallingUser)

	fenixTestDataSyncServerObject.currentTestDataState = *nextTestDataState

}

// ********************************************************************************************************************
// Change TestDataHeader-state
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) changeTestDataHeaderState(currentCallingUser string, nextTestDataHeaderState *executionStateTestDataHeaderTypeType) {

	//TODO Change from .Info to .Debug
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id":                      "47925769-4f29-4333-93ce-f07d7e442cb7",
		"Old TestDataHeaderState": fenixTestDataSyncServerObject.currentTestDataHeaderState,
		"New TestDataHeaderState": *nextTestDataHeaderState,
	}).Info("Change 'ChangeTestDataHeaderState' for Client: " + currentCallingUser)

	fenixTestDataSyncServerObject.currentTestDataHeaderState = *nextTestDataHeaderState

}

// ********************************************************************************************************************
// Check if the system is in correct TestDataHeaderState and return next expected state
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isSystemInCorrectTestDataHeaderState(currentCallingUser string, expectedTestDataHeaderState executionStateTestDataHeaderTypeType) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse, nextExpectedState executionStateTestDataHeaderTypeType) {

	// Verify that the system is in correct TestDataState
	if expectedTestDataHeaderState != fenixTestDataSyncServerObject.currentTestDataHeaderState {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TESTDATASERVER_INTERNAL_STATE_NOT_MATCHING_GRPC_CALL
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Expected Header-State for TestDataServer is " + strconv.FormatInt(int64(expectedTestDataHeaderState), 10) + " but current State is " + strconv.FormatInt(int64(fenixTestDataSyncServerObject.currentTestDataState), 10),
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "e903390a-421a-4d16-976d-08a464424385",
		}).Info("Expected Header-State for TestDataServer is " + strconv.FormatInt(int64(expectedTestDataHeaderState), 10) + " but current State is " + strconv.FormatInt(int64(fenixTestDataSyncServerObject.currentTestDataState), 10) + " for Client: " + currentCallingUser)

		return returnMessage, CurrenStateTestDataHeaderHash

	} else {

		nextExpectedState = nextTestDataHeaderStateMap[expectedTestDataHeaderState]

		return nil, nextExpectedState
	}

}

// ********************************************************************************************************************
// Check if there is a temporary stop in processing incoming or outgoing messages
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isThereATemporaryStopInProcessingInOrOutgoingMessagesAdmin() (returnMessage *fenixTestDataSyncServerGrpcAdminApi.AckNackResponse) {

	// Check if there is a temporary stop in processing in- or outgoing messages
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{
			AckNack:    false,
			Comments:   "There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "2e45c69b-a8c5-4b20-b40b-8a3e4025ed4c",
		}).Info("There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server")

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum
	var protoFileUsed fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
		}).Debug("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Admin-Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isAdminClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum) (returnMessage *fenixTestDataSyncServerGrpcAdminApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum
	var protoFileUsed fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "7831224f-5d6b-4afe-98d8-17a1acd6c799",
		}).Debug("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client is in list with known klients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientKnownToServer(callingClientUuid string) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// Check if Client exist in CloudDB
	_, valueExits := cloudDBClientsMap[memDBClientUuidType(callingClientUuid)]

	// Check if Client exists
	if valueExits == false {
		// Client doesn't exit in CloudDB (or hasn't been loaded from cloud

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_UNKNOWN_CALLER
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Client '" + callingClientUuid + "' is unknown to FenixTestDataSyncServer",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "dc311b1e-13c2-4c76-ab43-95c50f439d35",
		}).Debug("Calling Client '" + callingClientUuid + "' is unknown to FenixTestDataSyncServer")

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client has Hashed the MerklePath correctly
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientsMerklePathCorrectlyHashed(callingClientUuid string, merklePath string, merklePathHash string) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	var hashIsCorrectlyHashed bool

	// Verify that MerkleFilterPath is hashed correctly
	tempHashedMerkleFilter := fenixSyncShared.HashSingleValue(merklePath)

	// Check if MerklePath is correctly hashed
	if tempHashedMerkleFilter == merklePathHash {
		hashIsCorrectlyHashed = true
	} else {
		hashIsCorrectlyHashed = false
	}

	// Generate returnMessage if wrongly hashed
	if hashIsCorrectlyHashed == false {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_MERKLEPATHHASH_IS_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "MerklePathHash is not correct calculated for MerklePath='" + merklePath + "' . Expected: '" + tempHashedMerkleFilter + "', but got: '" + merklePathHash + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "57f3ef6b-ae39-40b3-904b-6f3e28ac2d32",
		}).Debug("MerklePathHash is not correct calculated for MerklePath='" + merklePath + "' . Expected: '" + tempHashedMerkleFilter + "', but got: '" + merklePathHash + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestFenixTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixProtoFileVersion != -1 {
		return highestFenixProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixProtoFileVersion = maxValue

	return highestFenixProtoFileVersion
}

// ********************************************************************************************************************
// Get the highest ClientProtoFileVersionEnumeration
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestClientTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestclientProtoFileVersion' saved, if so use that one
	if highestClientProtoFileVersion != -1 {
		return highestClientProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestClientProtoFileVersion = maxValue

	return highestClientProtoFileVersion
}

// Extract Values, and create, for TestDataHeaderItemMessageHash
func createTestDataHeaderItemMessageHash(testDataHeaderItemMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderItemMessage) (testDataHeaderItemMessageHash string) {

	var valuesToHash []string
	var valueToHash string

	// Extract and add values to array
	// HeaderLabel
	valueToHash = testDataHeaderItemMessage.HeaderLabel
	valuesToHash = append(valuesToHash, valueToHash)

	// HeaderShouldBeUsedForTestDataFilter as 'true' or 'false'
	if testDataHeaderItemMessage.HeaderShouldBeUsedForTestDataFilter == false {
		valuesToHash = append(valuesToHash, "false")
	} else {
		valuesToHash = append(valuesToHash, "true")
	}

	// HeaderIsMandatoryInTestDataFilter as 'true' or 'false'
	if testDataHeaderItemMessage.HeaderIsMandatoryInTestDataFilter == false {
		valuesToHash = append(valuesToHash, "false")
	} else {
		valuesToHash = append(valuesToHash, "true")
	}

	// HeaderSelectionType
	valueToHash = testDataHeaderItemMessage.HeaderSelectionType.String()
	valuesToHash = append(valuesToHash, valueToHash)

	// HeaderFilterValues - An array thar is added
	valueToHash = testDataHeaderItemMessage.HeaderLabel
	valuesToHash = append(valuesToHash, valueToHash)

	// Hash all values in the array
	testDataHeaderItemMessageHash = fenixSyncShared.HashValues(valuesToHash, true)

	return testDataHeaderItemMessageHash
}

// ********************************************************************************************************************
// Convert gRPC MerkleTreeNodes into object used in memDB and when storing to database
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertGrpcMerkleTreeNodesIntoMemDBMerkleTreeNodes(merkleHash string, gRPCMerkleTreeMessage *fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) (memDBMerkleTreeNodes []cloudDBTestDataMerkleTreeStruct) {

	var memDBMerkleTreeNode cloudDBTestDataMerkleTreeStruct

	// Loop all gRPC MerkleNodes and convert into object used in memoryDB
	for _, gRPCMerkleTreeNode := range gRPCMerkleTreeMessage.MerkleTreeNodes {

		memDBMerkleTreeNode = cloudDBTestDataMerkleTreeStruct{
			clientUuid:       gRPCMerkleTreeMessage.TestDataClientUuid,
			merkleHash:       merkleHash,
			nodeLevel:        int(gRPCMerkleTreeNode.NodeLevel),
			nodeName:         gRPCMerkleTreeNode.NodeName,
			nodePath:         gRPCMerkleTreeNode.NodePath,
			nodeHash:         gRPCMerkleTreeNode.NodeHash,
			nodeChildHash:    gRPCMerkleTreeNode.NodeChildHash,
			updatedTimeStamp: "",
		}

		// Append to list of all nodes
		memDBMerkleTreeNodes = append(memDBMerkleTreeNodes, memDBMerkleTreeNode)

	}

	return memDBMerkleTreeNodes

}

// Extract the MerkleFilterPaths to be sent to client, to be able to receive the missing, or changed, TestDataRows
func missedPathsToRetrieveFromClient(serverCopyMerkleTree []cloudDBTestDataMerkleTreeStruct, newClientMerkleTree []cloudDBTestDataMerkleTreeStruct) (merkleNodeNamesToRequest []string) {

	var highestLeafNodeLevel = -1
	var foundChildNodeHashInServer bool
	var leafNodeHashForClient string
	var leafNodesToRequest []cloudDBTestDataMerkleTreeStruct

	// Get highest LeafNodeLevel for Server-copy
	for _, leafNode := range serverCopyMerkleTree {
		if leafNode.nodeLevel > highestLeafNodeLevel {
			highestLeafNodeLevel = leafNode.nodeLevel
		}
	}

	// Extract all LeafNodes for Client
	// Loop over all LeafNodes to find those that are not in Server-copy
	for _, clientMerkleTreeNode := range newClientMerkleTree {

		// only process node when NodeLevel ==  'highestLeafNodeLevel'
		if clientMerkleTreeNode.nodeLevel == highestLeafNodeLevel {

			leafNodeHashForClient = clientMerkleTreeNode.nodeChildHash

			// Loop over  Server MerkleNodeChildHashes to see if the ChildNodeHash exists there
			foundChildNodeHashInServer = false

			for _, serverMerkleTreeNode := range serverCopyMerkleTree {

				// only process node when NodeLevel ==  'highestLeafNodeLevel'
				if clientMerkleTreeNode.nodeLevel == highestLeafNodeLevel {

					if serverMerkleTreeNode.nodeChildHash == leafNodeHashForClient {
						foundChildNodeHashInServer = true
						break
					}
				}
			}

			// if the Clients ChildHash wasn't found among ServerMerkleTreeNodes then add it to array
			if foundChildNodeHashInServer == false {
				leafNodesToRequest = append(leafNodesToRequest, clientMerkleTreeNode)
			}
		}
	}

	// Loop over 'leafNodesToRequest' to generate array with all MerkleTreeNodeNames to be able to retrieve correct TestData-rows from Client
	for _, leafNodeToRequest := range leafNodesToRequest {
		merkleNodeNamesToRequest = append(merkleNodeNamesToRequest, leafNodeToRequest.nodeName)
	}

	return merkleNodeNamesToRequest

}

// Search array and look if value exists in the array
func existsValueInStringArray(valueToSearchFor string, arrayToSearchIn []string) (existsInArray bool) {

	existsInArray = false

	// Loop array and search for value
	for _, value := range arrayToSearchIn {
		if valueToSearchFor == value {
			// Value was found
			existsInArray = true
			break
		}
	}

	return existsInArray
}

// // Convert gRPC-RowsMessage into cloudDBTestDataRowItems-message
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcTestDataRowsMessageToCloudDBTestDataRowItems(
	gRpcTestDataRowsItemMessage *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (memDBtestDataRowItems []cloudDBTestDataRowItemCurrentStruct) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "bdea5110-1af8-4e2f-a78a-ed1b2ad15514",
	}).Debug("Incoming gRPC 'convertgRpcTestDataRowsMessageToCloudDBTestDataRowItems'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "50c6be0a-a522-4aa6-ae7e-1db5936846f1",
	}).Debug("Outgoing gRPC 'convertgRpcTestDataRowsMessageToCloudDBTestDataRowItems'")

	//var leafNodeHash string
	//var leafNodeHashColumn int
	//var rowHashColumn int

	gRpcTestDataRowsMessage := gRpcTestDataRowsItemMessage.TestDataRows
	// Loop over gRPC-TestDataRow-messages and convert into memoryDB-object
	for gRpcTestDataRow, gRpcTestDataRowMessage := range gRpcTestDataRowsMessage {

		/*
			// Find the LeafNodeHash from dataFrame containing that info
			// TODO change from DataFrame to contain this into more standard data structure
			numberOfRowsInDataFrame := dataFrameHavingLeafNodeChildHashes.Nrow()
			headersInDataFrame := dataFrameHavingLeafNodeChildHashes.Names()
			for rowColumn, headerInDataFrame := range headersInDataFrame {
				if headerInDataFrame == "LeafNodeHash" {
					leafNodeHashColumn = rowColumn
				}
				if headerInDataFrame == "TestDataHash" {
					rowHashColumn = rowColumn
				}
			}

			for rowCounterInDataFrame := 0; rowCounterInDataFrame < numberOfRowsInDataFrame; rowCounterInDataFrame++ {
				if gRpcTestDataRowMessage.RowHash == dataFrameHavingLeafNodeChildHashes.Elem(rowCounterInDataFrame, rowHashColumn).String() {
					leafNodeHash = dataFrameHavingLeafNodeChildHashes.Elem(rowCounterInDataFrame, leafNodeHashColumn).String()
					break
				}
			}
		*/

		// Loop over columns in 'testDataColumnsItem'
		testDataColumnsItem := gRpcTestDataRowMessage.TestDataItems
		for testDataColumn, columnValue := range testDataColumnsItem {

			// Extract data and populate memoryDB-object
			memDBtestDataRowItem := cloudDBTestDataRowItemCurrentStruct{
				clientUuid:            gRpcTestDataRowsItemMessage.TestDataClientUuid,
				rowHash:               gRpcTestDataRowMessage.RowHash,
				testdataValueAsString: columnValue.TestDataItemValueAsString,
				leafNodeName:          gRpcTestDataRowMessage.LeafNodeName,
				leafNodePath:          gRpcTestDataRowMessage.LeafNodePath,
				leafNodeHash:          "", //leafNodeHash,
				valueColumnOrder:      testDataColumn,
				valueRowOrder:         gRpcTestDataRow,
				updatedTimeStamp:      "",
			}

			// Add 'memDBtestDataRowItem' to array
			memDBtestDataRowItems = append(memDBtestDataRowItems, memDBtestDataRowItem)
		}
	}

	return memDBtestDataRowItems
}

// Generate relations between LeafNodeName, LeafNodeHash and TestDataRowHash
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) generateRowHashToNodeHashMap(OldServerDataRowItemsAndNewClientDataRowItems []cloudDBTestDataRowItemCurrentStruct, clientsNewMerkleTreeRowItems []cloudDBTestDataMerkleTreeStruct) (rowHashToLeafNodeHashMap map[string]string) {

	rowHashToLeafNodeHashMap = make(map[string]string)       //map[<rowHash>]<leafNodeHash>
	leafNodeNameToLeafNodeHashMap := make(map[string]string) //map[<leafNodeName>]<leafNodHash>

	var highestLeafNodeLevel = -1
	// Get highest LeafNodeLevel for 'clientsNewMerkleTreeRowItems'
	for _, leafNode := range clientsNewMerkleTreeRowItems {
		if leafNode.nodeLevel > highestLeafNodeLevel {
			highestLeafNodeLevel = leafNode.nodeLevel
		}
	}

	// Loop over MerkleTreeItems and create relation between LeafNodeName -> LeafNodeHash
	for _, merkleTreeRowItem := range clientsNewMerkleTreeRowItems {

		// Only process LeafNodes
		if merkleTreeRowItem.nodeLevel == highestLeafNodeLevel {

			_, leafNodeNameExist := rowHashToLeafNodeHashMap[merkleTreeRowItem.nodeName]
			if leafNodeNameExist == true {
				// LeafNodeName should never exist twice
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"id":                         "d63cf823-8d8d-4053-b0e2-fbc2be89d867",
					"merkleTreeRowItem.nodeName": merkleTreeRowItem.nodeName,
				}).Fatal("LeafNodeName should never exist twice")
			}

			// Add relation to map
			leafNodeNameToLeafNodeHashMap[merkleTreeRowItem.nodeName] = merkleTreeRowItem.nodeHash

		}
	}

	/*
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":                            "b0986c76-932e-43ad-89e8-78773d43cc4c",
			"leafNodeNameToLeafNodeHashMap": leafNodeNameToLeafNodeHashMap,
		}).Debug("Map for 'leafNodeNameToLeafNodeHashMap'")
	*/

	// Loop all TestDataItems and create relation RowHash -> LeafNodeHash
	for _, testDataRowItem := range OldServerDataRowItemsAndNewClientDataRowItems {

		// Only need to process for first column because all columns, in same row, have same rowHash and LeafNode
		if testDataRowItem.valueColumnOrder == 0 {

			// Check if RowHash already exist in Map
			_, rowHashExist := rowHashToLeafNodeHashMap[testDataRowItem.rowHash]
			if rowHashExist == false {

				// No RowHash exist so add RowHash and LeafNodeHash that will be found in 'leafNodeNameToLeafNodeHashMap'
				// First verify that value exist in 'leafNodeNameToLeafNodeHashMap'
				_, leafNodeNameExist := leafNodeNameToLeafNodeHashMap[testDataRowItem.leafNodeName]
				if leafNodeNameExist == true {
					rowHashToLeafNodeHashMap[testDataRowItem.rowHash] = leafNodeNameToLeafNodeHashMap[testDataRowItem.leafNodeName]

				} else {

					// We should never come here because it should exist
					fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
						"id":                           "598f7326-37cc-4bdb-80dc-997235219d83",
						"testDataRowItem.leafNodeName": testDataRowItem.leafNodeName,
					}).Fatal("LeafNodeName is missing in 'leafNodeNameToLeafNodeHashMap' which never should happen")
				}
			}
		}
	}

	/*
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":                       "d383195f-07ff-4396-bbc8-3d9a400377da",
			"rowHashToLeafNodeHashMap": rowHashToLeafNodeHashMap,
		}).Debug("Map for 'rowHashToLeafNodeHashMap'")
	*/

	return rowHashToLeafNodeHashMap
}

// Add MerkleLeafNodeHashes to TestDataRowItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) addMerkleLeafNodeHashesToTestDataRowItems(testDataRowItems []cloudDBTestDataRowItemCurrentStruct, rowHashToLeafNodeHashMap map[string]string) []cloudDBTestDataRowItemCurrentStruct {

	/*
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":               "0ca4b60a-a5a2-4edd-8716-b854878b13d6",
			"testDataRowItems": testDataRowItems,
		}).Debug("testDataRowItems before processing 'addMerkleLeafNodeHashesToTestDataRowItems'")
	*/

	// Loop all TestDataRowItems and add the NodeHash to it
	for testDataRowItemPosition, testDataRowItem := range testDataRowItems {

		// Verify that value exist
		leafNodeHash, existsRowHash := rowHashToLeafNodeHashMap[testDataRowItem.rowHash]

		if existsRowHash == false {
			// It didn't exist which should not happen
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"id":                           "598f7326-37cc-4bdb-80dc-997235219d83",
				"testDataRowItem.leafNodeName": testDataRowItem.leafNodeName,
			}).Fatal("LeafNodeName is missing in 'leafNodeNameToLeafNodeHashMap' which never should happen")

		} else {

			// RowHash existed so add the LeadNodeHash to the TestDataRowItem
			testDataRowItems[testDataRowItemPosition].leafNodeHash = leafNodeHash
		}
	}

	/*
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":               "7856b33f-27b7-4d1d-b91f-be74c01f29bf",
			"testDataRowItems": testDataRowItems,
		}).Debug("testDataRowItems after processing 'addMerkleLeafNodeHashesToTestDataRowItems'")
	*/

	return testDataRowItems
}

// Get all LeafNodeItems for Server from memoryDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) extractLeafNodeItemsFromMerkleTree(merkleTreeNodeItems []cloudDBTestDataMerkleTreeStruct, serverHasNoPreviousClientData bool) (leafNodeItems []cloudDBTestDataMerkleTreeStruct) {

	highestLeafNodeLevel := fenixTestDataSyncServerObject.getMerkleLeafNodeLevel(merkleTreeNodeItems)

	if highestLeafNodeLevel == -1 {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "ce63654f-c3ff-4236-aad9-18bf684236f2",
		}).Fatal("Didn't find any LeafNodes among MerkleTreeNodeItems. This should not not happen")
	}

	// Loop MerkleTree and extract all LeafNodes
	for _, merkleNodeItem := range merkleTreeNodeItems {
		if merkleNodeItem.nodeLevel == highestLeafNodeLevel {
			leafNodeItems = append(leafNodeItems, merkleNodeItem)
		}
	}

	return leafNodeItems
}

// Find the NodeLevel for LeafNodes in MerkleTree
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getMerkleLeafNodeLevel(merkleTreeNodeItems []cloudDBTestDataMerkleTreeStruct) (highestLeafNodeLevel int) {

	highestLeafNodeLevel = -1

	// Get highest LeafNodeLevel for 'merkleTreeNodeItems'
	for _, leafNode := range merkleTreeNodeItems {
		if leafNode.nodeLevel > highestLeafNodeLevel {
			highestLeafNodeLevel = leafNode.nodeLevel
		}
	}

	return highestLeafNodeLevel
}

// Processes two sets of MerkleTreeNodeItems, A and B, by doing 'Not(A) Intersecting B'.
// THis givs you all MerkleTreeItems that exist in B but not in A
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) notAIntersectingBOnMerkleTreeItems(merkleTreeLeafNodeItemsA []cloudDBTestDataMerkleTreeStruct, merkleTreeLeafNodeItemsB []cloudDBTestDataMerkleTreeStruct) (merkleTreeNodeItemsToBeReturned []cloudDBTestDataMerkleTreeStruct) {

	var foundNodeBInMerkleTreeA bool

	// Loop over all MerkleTreeItems for 'B'
	for _, merkleTreeNodeForB := range merkleTreeLeafNodeItemsB {

		foundNodeBInMerkleTreeA = false

		// Loop over all MerkleTreeItems for 'A'
		for _, merkleTreeNodeForA := range merkleTreeLeafNodeItemsA {

			// only process node when NodeLevel ==  'highestLeafNodeLevel'
			if merkleTreeNodeForB.nodeHash == merkleTreeNodeForA.merkleHash {
				foundNodeBInMerkleTreeA = true
				break
			}
		}

		// If NodeB was not found in MerkleTreeA, then that Node should bes saved
		if foundNodeBInMerkleTreeA == false {
			merkleTreeNodeItemsToBeReturned = append(merkleTreeNodeItemsToBeReturned, merkleTreeNodeForB)
		}
	}

	return merkleTreeNodeItemsToBeReturned
}

// Verify that when recreating NodeHashes from RowHashes they should be found among MerkleTreeLeafeNodes for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) verifyThatRowHashesMatchesLeafNodeHashes(callingClientUuid string, testDataRowItems []cloudDBTestDataRowItemCurrentStruct, merkleTreeLeafNodeItems []cloudDBTestDataMerkleTreeStruct) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// Secure that we only operate on LeafNodes
	highestLeafNodeLevel := fenixTestDataSyncServerObject.getMerkleLeafNodeLevel(merkleTreeLeafNodeItems)

	if highestLeafNodeLevel == -1 {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "72437862-186d-411a-850c-4aa076f5731b",
		}).Fatal("Didn't find any LeafNodes among MerkleTreeNodeItems. This should not not happen")
	}

	// Map to hold RowHashes for each NodeName
	rowHashMap := make(map[string][]string) //map[<NodeName>][]<RowHash>

	// Generate the map for NodeNames and RowHashes
	for _, testDataRowItem := range testDataRowItems {

		// Only process on 'column = 0' because all columns on same row have the same RowHash
		if testDataRowItem.valueColumnOrder == 0 {

			// Check if LeafNodeName already is added
			_, valueExits := rowHashMap[testDataRowItem.leafNodeName]

			// Add RowHash to array
			if valueExits == true {
				// Get existing array to add to
				arrayOfRowHashes := rowHashMap[testDataRowItem.leafNodeName]
				arrayOfRowHashes = append(arrayOfRowHashes, testDataRowItem.rowHash)

				rowHashMap[testDataRowItem.leafNodeName] = arrayOfRowHashes
			} else {
				// Create new empty array to add to
				arrayOfRowHashes := []string{testDataRowItem.rowHash}

				rowHashMap[testDataRowItem.leafNodeName] = arrayOfRowHashes
			}
		}
	}

	// Map to hold LeafNodeHashes which are calculated from RowHashes group by NodeNames
	leafNodeHashesMap := make(map[string]string) //map[<NodeName>]<LeafNodeHash>

	// Loop over NodeNames and calculate LeafNodeHashes
	for nodeName, rowHashes := range rowHashMap {

		// Hash all RowHashes
		hashedValues := fenixSyncShared.HashValues(rowHashes, false)

		// Add value to map
		leafNodeHashesMap[nodeName] = hashedValues
	}

	// Verify that each calculated LeafNodeHash exists among the MerkleTreeItems
	var numberOfMissedLeafHashes = 0
	//var numberOfNodeHashesThatDidNotHaveAnyLeafNode = 0
	//var currentNodeHashWasFound bool

	// Loop all LeafNodes with calculated LeafNodeHashes
	for nodeName, nodeHash := range leafNodeHashesMap {

		// Loop all MerkleTreeNodeItems
		for _, merkleTreeLeafNodeItem := range merkleTreeLeafNodeItems {

			//currentNodeHashWasFound = false

			if nodeName == merkleTreeLeafNodeItem.nodeName {

				// If calculate Hash is not the same as the Hash in MerkleTree then we have something 'fishy'
				if nodeHash != merkleTreeLeafNodeItem.nodeHash {
					numberOfMissedLeafHashes = numberOfMissedLeafHashes + 1
					//currentNodeHashWasFound = true

					break
				}

			}

		}

		//numberOfNodeHashesThatDidNotHaveAnyLeafNode = numberOfNodeHashesThatDidNotHaveAnyLeafNode + 1
	}

	// Verify that there are no missed hashes
	if numberOfMissedLeafHashes > 0 {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":                       "4f26d271-56e9-4b15-ae5f-b00dab679c95",
			"numberOfMissedLeafHashes": numberOfMissedLeafHashes,
		}).Error("RowsHashes from TestDataRows did not recreated expected LeafNodeHashes")

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_ROWHASH_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "When recreating missed LeafNodeHashes from RowHashes all LeafNodeHashes could not be created",
			ErrorCodes: errorCodes,
		}

		return returnMessage

	}

	return nil

}

// Verify that the requested LeafNodeNames are found among the received TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) verifyThatRequestedLeafNodeNamesAreFoundAmongReceivedTestDataRows(callingClientUuid string, cloudDBTestDataRowItemsMessage []cloudDBTestDataRowItemCurrentStruct, requestedLeafNodeNames []string) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// When 'requestedLeafNodeNames' is empty then all rows are requested and we can't do this verification
	if len(requestedLeafNodeNames) == 0 {
		return nil
	}

	var numberOfMissedLeafNodeNames int
	var currentNodeNameWasFound bool

	// Verify that each LeafNodeName exists among the TestDataRows
	numberOfMissedLeafNodeNames = 0

	// Loop all TestDataRows
	for _, testDataRowItem := range cloudDBTestDataRowItemsMessage {

		currentNodeNameWasFound = false
		// Loop all MerkleTreeNodeItems
		for _, leafNodeName := range requestedLeafNodeNames {

			if leafNodeName == testDataRowItem.leafNodeName {
				currentNodeNameWasFound = true
				break
			}

		}
		// LeafNodeName was not found so add counte
		if currentNodeNameWasFound == false {
			numberOfMissedLeafNodeNames = numberOfMissedLeafNodeNames + 1

		}

	}

	// Verify that there are no missed hashes
	if numberOfMissedLeafNodeNames > 0 {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":                          "e9a51e1b-d848-4b56-a383-621e7a1aeabc",
			"numberOfMissedLeafNodeNames": numberOfMissedLeafNodeNames,
		}).Error("Didn't found all Requested LeafNodeNames among TestDataRows")

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_ROWHASH_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Didn't found all Requested LeafNodeNames among TestDataRows",
			ErrorCodes: errorCodes,
		}

		return returnMessage

	}

	return nil

}

// Verify that saved TestDataLeafNodeHashes recreates MerkleHash
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) verifyThatMerkleLeafNodeHashesReCreatesMerkleHash(callingClientUuid string) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	var leafNodeHashes [][]string

	// Load TestDataLeafNodeHashes from CLoudDBleafNodeHash
	err := fenixTestDataSyncServerObject.loadAllTestDataLeafNodeHashesForClientFromCloudDB(callingClientUuid, &leafNodeHashes)

	if err != nil {

		// Problem executing SQL
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":  "dd5e92e7-fd09-4745-b64c-b8020b3ae04a",
			"err": err,
		}).Error("Problem executing SQL for Client: " + callingClientUuid)

		// Stop processing in- and outgoing messages
		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "6e391cda-959d-417b-92eb-7cf20d90822e",
		}).Error("Stop processing in- and outgoing messages for Client: " + callingClientUuid)

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Problem when executing SQL",
			ErrorCodes: errorCodes,
		}

		return returnMessage
	}

	// Verify that there are some rows in CLoudDB
	if len(leafNodeHashes) == 0 {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "6903f33f-78ed-4ab6-847f-681a3f32fcec",
		}).Fatal("Didn't expect zero rows when reading CloudDB for Client: " + callingClientUuid)
	}

	// Convert LeafNode-data into MerkleTree-DataFrame
	leafNodeMerkleTree := fenixSyncShared.ConvertLeafNodeMessagesToDataframe(leafNodeHashes, fenixTestDataSyncServerObject.logger)

	// XXX
	f, err := os.Create("leafNodeMerkleTree.csv")
	if err != nil {
		log.Fatal(err)
	}
	leafNodeMerkleTree.WriteCSV(f)
	f.Close()

	// Calculate MerkleHash from MerkleTree
	//TODO In FenixShare then add Client_UUID instead of "" as starting value
	reCalculatedMerkleHash := fenixSyncShared.CalculateMerkleHashFromMerkleTree(leafNodeMerkleTree)

	if reCalculatedMerkleHash == "-1" {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "a23734b2-fbdd-470d-88b8-9699a3637404",
		}).Fatal("Something went from when calculating MerkleHash for Client: " + callingClientUuid)

	}

	currentMerkleHashInMemDB := fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(callingClientUuid)

	// Verify that 'reCalculatedMerkleHash' == 'currentMerkleHashInMemDB'
	if reCalculatedMerkleHash != currentMerkleHashInMemDB {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":                       "4f26d271-56e9-4b15-ae5f-b00dab679c95",
			"reCalculatedMerkleHash":   reCalculatedMerkleHash,
			"currentMerkleHashInMemDB": currentMerkleHashInMemDB,
		}).Fatal("'reCalculatedMerkleHash' is not the same as 'currentMerkleHashInMemDB' for Client: " + currentMerkleHashInMemDB)

	}

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id":                       "fdeade26-544e-436d-bf59-55d91614230f",
		"reCalculatedMerkleHash":   reCalculatedMerkleHash,
		"currentMerkleHashInMemDB": currentMerkleHashInMemDB,
	}).Debug("'reCalculatedMerkleHash' from TestData-rows in CloudDB is the same as 'currentMerkleHashInMemDB' for Client: " + currentMerkleHashInMemDB)

	return nil

}

// *********************************************************************

// Creates a Defer function that process the function in FIFO way instead of standard defer LIFO way
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) reversedDefer(deferFunctions *[]func()) {

	// Loop all defer-functions and execute them in FIFO way
	for _, functionToCall := range *deferFunctions {
		functionToCall()
	}

	// Clear slice for functions
	deferFunctions = nil
}

// Save all components of the gRPCHeaderMessage into memoryDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderMessageDataForClient(testDataHeaderMessage *fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) bool {

	var callingClientUuid string
	var headerItemsHash string
	//var headerLabelsHash string

	// Extract the calling clients unique id
	callingClientUuid = testDataHeaderMessage.TestDataClientUuid

	// Extract the HeaderHash
	headerItemsHash = testDataHeaderMessage.HeaderLabelsHash

	// Extract the HeaderLabelsHash
	//headerLabelsHash = testDataHeaderMessage.HeaderLabelsHash

	// Extract TestDataHeaderItems
	testDataHeaderItems := testDataHeaderMessage.TestDataHeaderItems

	var cloudDBTestDataHeaderItems []cloudDBTestDataHeaderItemStruct
	var cloudDBTestDataHeaderFilterValues []cloudDBTestDataHeaderFilterValuesStruct

	// Loop all TestDataHeaderItems and split different parts and Save to memoryDB Objects
	for testDataHeaderItemCounter, testDataHeaderItem := range testDataHeaderItems {

		cloudDBTestDataHeaderItem := cloudDBTestDataHeaderItemStruct{
			headerItemsHash:      headerItemsHash,
			headerItemHash:       testDataHeaderItem.TestDataHeaderItemMessageHash,
			clientUuid:           callingClientUuid,
			headerLabel:          testDataHeaderItem.HeaderLabel,
			shouldBeUsedInFilter: testDataHeaderItem.HeaderShouldBeUsedForTestDataFilter,
			isMandatoryInFilter:  testDataHeaderItem.HeaderIsMandatoryInTestDataFilter,
			filterSelectionType:  int(testDataHeaderItem.HeaderSelectionType),
			headerColumnOrder:    testDataHeaderItemCounter,
		}

		cloudDBTestDataHeaderItems = append(cloudDBTestDataHeaderItems, cloudDBTestDataHeaderItem)

		// Extract Filter Values
		cloudDBTestDataHeaderFilterValues = []cloudDBTestDataHeaderFilterValuesStruct{}

		// Loop all Header Filter values
		for headerFilterValueCounter, headerFilterValue := range testDataHeaderItem.HeaderFilterValues {
			cloudDBTestDataHeaderFilterValue := cloudDBTestDataHeaderFilterValuesStruct{
				headerItemHash:         testDataHeaderItem.TestDataHeaderItemMessageHash,
				headerFilterValueOrder: headerFilterValueCounter,
				headerFilterValue:      headerFilterValue.HeaderFilterValuesAsString,
				clientUuid:             callingClientUuid,
			}

			cloudDBTestDataHeaderFilterValues = append(cloudDBTestDataHeaderFilterValues, cloudDBTestDataHeaderFilterValue)
		}
	}

	// Save 'cloudDBTestDataHeaderItems' to memoryDB
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderItemsToClient(callingClientUuid, cloudDBTestDataHeaderItems)

	// Save 'cloudDBTestDataHeaderFilterValues' to memoryDB
	_ = fenixTestDataSyncServerObject.saveCurrentHeaderFilterValuesToClient(callingClientUuid, cloudDBTestDataHeaderFilterValues)

	return true
}
