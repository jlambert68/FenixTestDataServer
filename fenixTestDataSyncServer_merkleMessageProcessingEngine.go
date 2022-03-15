package main

import (
	//fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"strconv"
)

/*
// Initiate channels for "incoming gRPC-messages"
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) InitiateProcessEngineChannels() {

	make(TestDataClientInformationMessageChannel, fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage, 1)

	// 'MerkleHashMessage' from 'gRPC-SendMerkleHash'
	var MerkleHashMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleHashMessage

	// 'MerkleTreeMessage' from 'gRPC-SendMerkleTree'
	var MerkleTreeMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

	// 'TestDataHeaderMessage' from 'gRPC-SendTestDataHeaders'
	var TestDataHeaderMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage

	// 'MerkleTreeMessage' from 'gRPC-SendTestDataRows'
	var MerkleTreeMessageMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

}
*/

// Convert gRPC-MerkleTree message into a DataFrame object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcMerkleTreeMessageToDataframe(merkleTreeMessage fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) dataframe.DataFrame {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "3f08ad8e-97d0-4762-b2a4-073f5eea113a",
	}).Debug("Incoming gRPC 'convertgRpcMerkleTreeMessageToDataframe'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "17be0c7e-cfb9-4505-a7a0-d44542485f59",
	}).Debug("Outgoing gRPC 'convertgRpcMerkleTreeMessageToDataframe'")

	var myMerkleTree []MerkletreeStruct

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	merkleTreeNodes := merkleTreeMessage.MerkleTreeNodes

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, merkleTreeNode := range merkleTreeNodes {
		myMerkleTreeRow := MerkletreeStruct{
			MerkleLevel:     int(merkleTreeNode.NodeLevel),
			MerklePath:      merkleTreeNode.NodeName,
			MerkleHash:      merkleTreeNode.NodeHash,
			MerkleChildHash: merkleTreeNode.NodeChildHash,
		}
		myMerkleTree = append(myMerkleTree, myMerkleTreeRow)

	}

	df := dataframe.LoadStructs(myMerkleTree)

	return df
}

//TODO Remove the below and use FenixShared instead
/*
// Convert leafNodeHash and LeafNodeName message into a MerkleTree DataFrame object;
// leafNodesMessage [][]string; [[<LeafNodeHash>, <LeafNodeName], [<>, <>]]
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertLeafNodeMessagesToDataframe(leafNodesMessage [][]string) dataframe.DataFrame {
	// leafNodesMessage[n] = 'leafNode'
	// leafNode[0] = 'LeafNodeHash'
	// leafNode[1] = 'LeafNodeName'

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c0b9dd6c-2431-4b71-b476-9d71eebf6d29",
	}).Debug("Incoming gRPC 'convertLeafNodeMessagesToDataframe'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "ce67d061-777b-4cc6-9672-d0cfdf3f2c83",
	}).Debug("Outgoing gRPC 'convertLeafNodeMessagesToDataframe'")

	var myMerkleTree []MerkletreeStruct

	// Number of MerkleLevels for MerkleTree
	var numberOfMerkleLevels = 0

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, leafNode := range leafNodesMessage {

		// Get number of MerkleLevels for MerkleTree
		if numberOfMerkleLevels == 0 {
			numberOfMerkleLevels = strings.Count(leafNode[1], "/")
		}

		// Create row and add to MerkleTree
		myMerkleTreeRow := MerkletreeStruct{
			MerkleLevel:     numberOfMerkleLevels,
			MerklePath:      leafNode[1],
			MerkleHash:      leafNode[0],
			MerkleChildHash: "1", // Set '1', doesn't matter
		}
		myMerkleTree = append(myMerkleTree, myMerkleTreeRow)

	}

	df := dataframe.LoadStructs(myMerkleTree)

	return df
}

*/
// Verify that HeaderHash is correct calculated
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) verifyThatHeaderItemsHashIsCorrectCalculated(testDataHeadersMessage fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "2d11c2bb-83cc-4a01-be55-c291105fa98f",
	}).Debug("Incoming gRPC 'verifyThatHeaderItemsHashIsCorrectCalculated'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "19bca393-92e4-40ea-958f-ab97f8a1fcd1",
	}).Debug("Outgoing gRPC 'verifyThatHeaderItemsHashIsCorrectCalculated'")

	// Extract  HeaderHash
	headerHash := testDataHeadersMessage.TestDataHeaderItemsHash

	var headerItemValues []string
	testDataHeaderItems := testDataHeadersMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		testDataHeaderItemMessageHash := createTestDataHeaderItemMessageHash(headerItem)
		headerItemValues = append(headerItemValues, testDataHeaderItemMessageHash)

	}

	// Hash all 'testDataHeaderItemMessageHash' into a single hash
	reHashedHeaderItemMessageHash := fenixSyncShared.HashValues(headerItemValues, false)

	if headerHash != reHashedHeaderItemMessageHash {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_UNSPECIFIED // TODO Add error code for HeaderHash-error
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "HeaderItemsHash is not correct calculated. Expected '" + reHashedHeaderItemMessageHash + "', but got '" + headerHash + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "785424d8-2d82-4581-b8ce-4ac49650666c",
		}).Info("HeaderItemsHash is not correct calculated. Expected '" + reHashedHeaderItemMessageHash + "', but got '" + headerHash + "'")

		// Exit function Respond back to client when hash error
		return returnMessage
	}

	return nil
}

// Convert gRPC-Header message into string and string array objects
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcHeaderMessageToStringArray(testDataHeadersMessage fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (headerHash string, headerLabelHash string, headersItems []string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aa1d5eb1-7503-467f-9336-1927fa5529f2",
	}).Debug("Incoming gRPC 'convertgRpcHeaderMessageToStringArray'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "72779d91-4b28-44ad-b3b7-4beeed915cc9",
	}).Debug("Outgoing gRPC 'convertgRpcHeaderMessageToStringArray'")

	// Extract  HeaderHash
	headerHash = testDataHeadersMessage.TestDataHeaderItemsHash

	// Extract  HeaderHash
	headerLabelHash = testDataHeadersMessage.HeaderLabelsHash

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	testDataHeaderItems := testDataHeadersMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		headersItems = append(headersItems, headerItem.HeaderLabel)

	}

	return headerHash, headerLabelHash, headersItems
}

// Convert TestDataRow message into TestData dataframe object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (testdataAsDataFrame dataframe.DataFrame, returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "585ed48f-dc60-4d93-9628-fb77db8057fb",
	}).Debug("Incoming gRPC 'convertgRpcTestDataRowsMessageToDataFrame'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "8afb8c12-7dcc-431f-9255-f0866555bd62",
	}).Debug("Outgoing gRPC 'convertgRpcTestDataRowsMessageToDataFrame'")

	testdataAsDataFrame = dataframe.New()

	currentTestDataClientGuid := testdataRowsMessages.TestDataClientUuid

	currentTestDataHeaders := fenixTestDataSyncServerObject.getCurrentHeadersForClient(currentTestDataClientGuid)

	// If there are no headers in Database then Ask client for HeaderHash
	if len(currentTestDataHeaders) == 0 {
		fenixTestDataSyncServerObject.AskClientToSendTestDataHeaderHash(currentTestDataClientGuid)
		currentTestDataHeaders = fenixTestDataSyncServerObject.getCurrentHeadersForClient(currentTestDataClientGuid)

		// Validate that we got hte TestData Headers
		if len(currentTestDataHeaders) == 0 {

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_UNKNOWN_CALLER //TODO Change to correct error
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   "Fenix Asked for TestDataHeaders but didn't receive them i a correct way",
				ErrorCodes: errorCodes,
			}

			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id": "bccab03c-dda2-47bd-bbca-70f10bc052c7",
			}).Info("Fenix Asked for TestDataHeaders but didn't receive them i a correct way")

			// leave
			return testdataAsDataFrame, returnMessage
		}
	}

	// Add 'KEY' to all headers
	var testDataHeadersInDataFrame []string
	testDataHeadersInDataFrame = append(testDataHeadersInDataFrame, currentTestDataHeaders...)
	testDataHeadersInDataFrame = append(testDataHeadersInDataFrame, "TestDataHash")

	testDataRows := testdataRowsMessages.TestDataRows

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, testDataRow := range testDataRows {

		// Create one row, as a dataframe
		rowDataframe := dataframe.New()
		var valuesToHash []string

		for testDataItemCounter, testDataItem := range testDataRow.TestDataItems {

			if rowDataframe.Nrow() == 0 {
				// Create New
				rowDataframe = dataframe.New(
					series.New([]string{testDataItem.TestDataItemValueAsString}, series.String, currentTestDataHeaders[testDataItemCounter]))
			} else {
				// Add to existing
				rowDataframe = rowDataframe.Mutate(
					series.New([]string{testDataItem.TestDataItemValueAsString}, series.String, currentTestDataHeaders[testDataItemCounter]))
			}

			valuesToHash = append(valuesToHash, testDataItem.TestDataItemValueAsString)
		}

		// Create and add column for 'TestDataHash'
		testDataHashSeriesColumn := series.New([]string{"key"}, series.String, "TestDataHash")
		rowDataframe = rowDataframe.Mutate(testDataHashSeriesColumn)

		// Hash all values for row
		hashedRow := fenixSyncShared.HashValues(valuesToHash, true)

		// Validate that Row-hash is correct calculated
		if hashedRow != testDataRow.RowHash {

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_ROWHASH_NOT_CORRECT_CALCULATED
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   "RowsHashes seems not to be correct calculated.",
				ErrorCodes: errorCodes,
			}

			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id": "eff5fb2f-4fd5-4fc7-a4eb-69bee9f22d92",
			}).Info("RowsHashes seems not to be correct calculated.")

			// Exit function Respond back to client when hash error
			return testdataAsDataFrame, returnMessage
		}

		// Add TestDataHash to row DataFrame
		rowDataframe.Elem(0, rowDataframe.Ncol()-1).Set(hashedRow)
		//) Mutate(
		//	series.New([]string{hashedRow}, series.String, "TestDataHash"))

		// Add the row to the Dataframe for the testdata
		// Special handling first when first time
		if testdataAsDataFrame.Nrow() == 0 {
			testdataAsDataFrame = rowDataframe.Copy()

		} else {
			testdataAsDataFrame = testdataAsDataFrame.OuterJoin(rowDataframe, testDataHeadersInDataFrame...)
		}
	}

	return testdataAsDataFrame, nil

}

// Concatenate TestDataRows as a DataFrame with the current Server TestDataRows-DataFrame
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) concatenateWithCurrentServerTestData(testDataClientGuid string, testdataToBeAdded []cloudDBTestDataRowItemCurrentStruct) (concatenatedTestdata []cloudDBTestDataRowItemCurrentStruct) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "75e3cea1-d459-4f4f-9289-b728feee7ec0",
	}).Debug("Incoming gRPC 'concatenateWithCurrentServerTestData'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "f77a358e-50ad-4f1c-a540-6559688e173f",
	}).Debug("Outgoing gRPC 'concatenateWithCurrentServerTestData'")

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5716a1e8-11dc-4c70-9579-94d7d177689b",
	}).Debug("New TestDataRowsItems to add to existing TestDataRowsItems are '"+strconv.Itoa(len(testdataToBeAdded))+"' for Client: ", testDataClientGuid)

	// Get TestDataRowsItems from memoryDB
	testDataRowItemsForServerInMemDB := fenixTestDataSyncServerObject.getCurrentTestDataRowItemsForServer(testDataClientGuid)
	//concatenatedTestdata = fenixTestDataSyncServerObject.getCurrentTestDataRowItemsForClient(testDataClientGuid)

	if len(testDataRowItemsForServerInMemDB) == 0 {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "a98bf92e-ec7d-4968-b8fa-b72e47fef830",
		}).Debug("There are no TestDataRowsItems in memDB for so returning new TestDataRowsItems to work with, Client: ", testDataClientGuid)

		return testdataToBeAdded

	} else {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "c8385ebf-53e1-4449-8171-f0c5bc2bdd68",
		}).Debug("Existing number of TestDataRowsItems are '"+strconv.Itoa(len(testDataRowItemsForServerInMemDB))+"' for Client: ", testDataClientGuid)

		//headerKeys := testdataToBeAdded.Names()
		//concatenatedTestdata = concatenatedTestdata.Concat(testdataToBeAdded) //, headerKeys...)
		concatenatedTestdata = append(testDataRowItemsForServerInMemDB, testdataToBeAdded...)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "85138c62-28a8-4669-a701-590c29fb4de1",
		}).Debug("Concatenated number of TestDataRowsItems are '"+strconv.Itoa(len(concatenatedTestdata))+"' for Client: ", testDataClientGuid)

	}

	return concatenatedTestdata

}

// Convert TestDataRow message into TestData dataframe object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertCloudDBTestDataRowItemsMessageToDataFrame(cloudDBTestDataRowItems []cloudDBTestDataRowItemCurrentStruct) (testdataAsDataFrame dataframe.DataFrame, returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "4a22eea7-806f-4d50-9b2f-1d5449203db6",
	}).Debug("Incoming gRPC 'convertCloudDBTestDataRowItemsMessageToDataFrame'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "9768e8b2-0c0f-41cc-90e5-3f0c17bb9ed8",
	}).Debug("Outgoing gRPC 'convertCloudDBTestDataRowItemsMessageToDataFrame'")

	testdataAsDataFrame = dataframe.New()

	currentTestDataClientGuid := cloudDBTestDataRowItems[0].clientUuid

	currentTestDataHeaders := fenixTestDataSyncServerObject.getCurrentHeadersForClient(currentTestDataClientGuid)

	// If there are no headers in Database then Ask client for HeaderHash
	if len(currentTestDataHeaders) == 0 {
		fenixTestDataSyncServerObject.AskClientToSendTestDataHeaderHash(currentTestDataClientGuid)
		currentTestDataHeaders = fenixTestDataSyncServerObject.getCurrentHeadersForClient(currentTestDataClientGuid)

		// Validate that we got hte TestData Headers
		if len(currentTestDataHeaders) == 0 {

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_UNKNOWN_CALLER //TODO Change to correct error
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   "Fenix Asked for TestDataHeaders but didn't receive them i a correct way",
				ErrorCodes: errorCodes,
			}

			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id": "b20fb287-2e60-4f6f-b635-fea49f367a67",
			}).Info("Fenix Asked for TestDataHeaders but didn't receive them i a correct way")

			// leave
			return testdataAsDataFrame, returnMessage
		}
	}

	// Add 'KEY' to all headers
	var testDataHeadersInDataFrame []string
	testDataHeadersInDataFrame = append(testDataHeadersInDataFrame, currentTestDataHeaders...)
	testDataHeadersInDataFrame = append(testDataHeadersInDataFrame, "TestDataHash")

	//testDataRows := testdataRowsMessages.TestDataRows
	// Create matrix for testdata
	dataMatrix := make(map[int]map[int]string) //make(map[<row>>]map[<column>]<value>

	// Create a map for RowHashes
	testDataRowHashes := make(map[int]string) //make(map[<row>>]<rowHash>

	// Loop over all 'cloudDBTestDataRowItems' and add to matriix
	for _, cloudDBTestDataRowItem := range cloudDBTestDataRowItems {

		// Verify that datapoint doesn't exist
		_, dataPointExists := dataMatrix[cloudDBTestDataRowItem.valueRowOrder][cloudDBTestDataRowItem.valueColumnOrder]
		if dataPointExists == true {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":                                   "efa873d1-b023-48db-b948-67fe00e103d7",
				"cloudDBTestDataRowItem.valueRowOrder": cloudDBTestDataRowItem.valueRowOrder,
				"cloudDBTestDataRowItem.valueColumnOrder": cloudDBTestDataRowItem.valueRowOrder,
			}).Fatal("Datapoint should only appears once")
		}

		// Add data to matrix
		// If 'row-map" already doesn't exist then initiate it
		_, rowExists := dataMatrix[cloudDBTestDataRowItem.valueRowOrder]
		if rowExists == false {
			// Initiate row in map and add column value
			dataMatrix[cloudDBTestDataRowItem.valueRowOrder] = map[int]string{}
			dataMatrix[cloudDBTestDataRowItem.valueRowOrder][cloudDBTestDataRowItem.valueColumnOrder] = cloudDBTestDataRowItem.testdataValueAsString
		} else {
			// Row exists then just add column value
			dataMatrix[cloudDBTestDataRowItem.valueRowOrder][cloudDBTestDataRowItem.valueColumnOrder] = cloudDBTestDataRowItem.testdataValueAsString
		}

		// Only add RowHash if it not exists
		_, rowHashExists := testDataRowHashes[cloudDBTestDataRowItem.valueRowOrder]
		if rowHashExists == false {
			testDataRowHashes[cloudDBTestDataRowItem.valueRowOrder] = cloudDBTestDataRowItem.rowHash
		}
	}

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	numberOfRowsInMatrix := len(dataMatrix)
	var numberOfColumnsInMatrixRow int
	var numberOfColumnInFirstMatrixRow int

	for testDataRowCounter := 0; testDataRowCounter < numberOfRowsInMatrix; testDataRowCounter++ {

		// Extract row
		testDataRow := dataMatrix[testDataRowCounter]

		// Create one row, as a dataframe
		rowDataframe := dataframe.New()
		var valuesToHash []string

		// Get the number of columns in row
		numberOfColumnsInMatrixRow = len(testDataRow)

		// Verify that all rows have the same number of columns
		if testDataRowCounter == 0 {
			numberOfColumnInFirstMatrixRow = numberOfColumnsInMatrixRow

		} else {

			if numberOfColumnsInMatrixRow != numberOfColumnInFirstMatrixRow {
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                             "a2043be7-657a-4c94-a1a0-374243e82571",
					"numberOfColumnInFirstMatrixRow": numberOfColumnInFirstMatrixRow,
					"numberOfColumnsInMatrixRow":     numberOfColumnsInMatrixRow,
				}).Fatal("It seems that all TestDataRows doesn't have the same number o columns")
			}
		}

		// Loop over columns
		for testDataColumnCounter := 0; testDataColumnCounter < numberOfColumnsInMatrixRow; testDataColumnCounter++ {
			//		for testDataItemCounter, testDataItem := range testDataRow {

			if rowDataframe.Nrow() == 0 {
				// Create New
				rowDataframe = dataframe.New(
					series.New([]string{testDataRow[testDataColumnCounter]}, series.String, currentTestDataHeaders[testDataColumnCounter]))
			} else {
				// Add to existing
				rowDataframe = rowDataframe.Mutate(
					series.New([]string{testDataRow[testDataColumnCounter]}, series.String, currentTestDataHeaders[testDataColumnCounter]))
			}

			valuesToHash = append(valuesToHash, testDataRow[testDataColumnCounter])
		}

		// Create and add column for 'TestDataHash'
		testDataHashSeriesColumn := series.New([]string{"key"}, series.String, "TestDataHash")
		rowDataframe = rowDataframe.Mutate(testDataHashSeriesColumn)

		// Hash all values for row
		hashedRow := fenixSyncShared.HashValues(valuesToHash, true)

		// Validate that Row-hash is correct calculated
		if hashedRow != testDataRowHashes[testDataRowCounter] {

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_ROWHASH_NOT_CORRECT_CALCULATED
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				AckNack:    false,
				Comments:   "RowsHashes seems not to be correct calculated.",
				ErrorCodes: errorCodes,
			}

			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id": "9e591230-1100-4771-ae38-c98a71daf784",
			}).Info("RowsHashes seems not to be correct calculated.")

			// Exit function Respond back to client when hash error
			return testdataAsDataFrame, returnMessage
		}

		// Add TestDataHash to row DataFrame
		rowDataframe.Elem(0, rowDataframe.Ncol()-1).Set(hashedRow)
		//) Mutate(
		//	series.New([]string{hashedRow}, series.String, "TestDataHash"))

		// Add the row to the Dataframe for the testdata
		// Special handling first when first time
		if testdataAsDataFrame.Nrow() == 0 {
			testdataAsDataFrame = rowDataframe.Copy()

		} else {
			testdataAsDataFrame = testdataAsDataFrame.OuterJoin(rowDataframe, testDataHeadersInDataFrame...)
		}
	}

	return testdataAsDataFrame, nil

}

// Convert memDB-MerkleTree message into a DataFrame object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertMemDBMerkleTreeMessageToDataframe(merkleTreeMessage []cloudDBTestDataMerkleTreeStruct) dataframe.DataFrame {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "72f59758-5101-4117-8368-cb14acf4368b",
	}).Debug("Incoming gRPC 'convertmemDBMerkleTreeMessageToDataframe'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "7a910fb5-f6a2-4704-ad38-da5ac0153cf2",
	}).Debug("Outgoing gRPC 'convertmemDBMerkleTreeMessageToDataframe'")

	var myMerkleTree []MerkletreeStruct

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, merkleTreeNode := range merkleTreeMessage {
		myMerkleTreeRow := MerkletreeStruct{
			MerkleLevel:     int(merkleTreeNode.nodeLevel),
			MerklePath:      merkleTreeNode.nodeName,
			MerkleHash:      merkleTreeNode.nodeHash,
			MerkleChildHash: merkleTreeNode.nodeChildHash,
		}
		myMerkleTree = append(myMerkleTree, myMerkleTreeRow)

	}

	df := dataframe.LoadStructs(myMerkleTree)

	return df
}
