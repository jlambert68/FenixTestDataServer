package main

import (
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
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
		}).Debug("There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server")

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) generateRowHashToMerkleChildNodeHashMap(OldServerDataRowItemsAndNewClientDataRowItems []cloudDBTestDataRowItemCurrentStruct, clientsNewMerkleTreeRowItems []cloudDBTestDataMerkleTreeStruct) (rowHashToLeafNodeHashMap map[string]string) {

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
			leafNodeNameToLeafNodeHashMap[merkleTreeRowItem.nodeName] = merkleTreeRowItem.nodeChildHash

		}
	}

	// Loop all TestDataItems and create relation RowHash -> LeafNodeHash
	for _, testDataRowItem := range OldServerDataRowItemsAndNewClientDataRowItems {

		// Only need to process for first column because all columns, in same row, have same rowHash and LeafNode
		if testDataRowItem.valueColumnOrder == 0 {

			// Check if RowHah already exist in Map
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

	return rowHashToLeafNodeHashMap
}

// Add MerkleLeafNodeHashes to TestDataRowItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) addMerkleLeafNodeHashesToTestDataRowItems(testDataRowItems []cloudDBTestDataRowItemCurrentStruct, rowHashToLeafNodeHashMap map[string]string) []cloudDBTestDataRowItemCurrentStruct {

	// Loop all TestDataRowItems and add the ChildNodeHash to it
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

	return testDataRowItems
}

// Get all LeafNodeItems for Server from memoryDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) extractLeafNodeItemsFromMerkleTree(merkleTreeNodeItems []cloudDBTestDataMerkleTreeStruct) (leafNodeItems []cloudDBTestDataMerkleTreeStruct) {

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

//
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
			if valueExits == false {
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

	lägg till något så man inte missar något

	// Verify that there are no missed hashes
	if numberOfMissedLeafHashes > 0 {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "4f26d271-56e9-4b15-ae5f-b00dab679c95",
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
