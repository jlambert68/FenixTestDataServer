package main

import (
	"FenixTestDataServer/common_config"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
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

	var myMerkleTree []MerkleTree_struct

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	merkleTreeNodes := merkleTreeMessage.MerkleTreeNodes

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, merkleTreeNode := range merkleTreeNodes {
		myMerkleTreeRow := MerkleTree_struct{
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

// Verify that HeaderHash is correct calculated
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) verifyThatHeaderItemsHashIsCorrectCalculated(testDataHeadersMessage fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// Extract  HeaderHash
	headerHash := testDataHeadersMessage.TestDataHeaderItemsHash

	var headerItemValues []string
	testDataHeaderItems := testDataHeadersMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		testDataHeaderItemMessageHash := common_config.CreateTestDataHeaderItemMessageHash(headerItem)
		headerItemValues = append(headerItemValues, testDataHeaderItemMessageHash)

	}

	// Hash all 'testDataHeaderItemMessageHash' into a single hash
	reHashedHeaderItemMessageHash := common_config.HashValues(headerItemValues, false)

	if headerHash != reHashedHeaderItemMessageHash {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_HEADERLABELHASH_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "HeaderItemsHash is not correct calculated. Expected '" + headerHash + "', but got '" + headerHash + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "785424d8-2d82-4581-b8ce-4ac49650666c",
		}).Info("HeaderItemsHash is not correct calculated. Expected '" + headerHash + "', but got '" + headerHash + "'")

		// Exit function Respond back to client when hash error
		return returnMessage
	}

	return nil
}

// Convert gRPC-Header message into string and string array objects
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcHeaderMessageToStringArray(testDataHeadersMessage fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage) (headerHash string, headersItems []string) {

	// Extract  HeaderHash
	headerHash = testDataHeadersMessage.TestDataHeaderItemsHash

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	testDataHeaderItems := testDataHeadersMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		headersItems = append(headersItems, headerItem.HeaderLabel)

	}

	return headerHash, headersItems
}

// Convert TestDataRow message into TestData dataframe object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (testdataAsDataFrame dataframe.DataFrame, returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

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
		hashedRow := common_config.HashValues(valuesToHash, true)

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

// Concartenate TestDataRows as a DataFrame with the current Server TestDataRows-DataFrame
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) concartenateWithCurrentServerTestData(testDataClientGuid string, testdataDataframe dataframe.DataFrame) (allTestdataAsDataFrame dataframe.DataFrame) {

	allTestdataAsDataFrame = fenixTestDataSyncServerObject.getCurrentTestDataRowsForServer(testDataClientGuid)

	if allTestdataAsDataFrame.Nrow() == 0 {
		return testdataDataframe
	} else {
		headerKeys := testdataDataframe.Names()
		allTestdataAsDataFrame = allTestdataAsDataFrame.OuterJoin(testdataDataframe, headerKeys...)
	}

	return allTestdataAsDataFrame

}

/*
// Convert Dataframe into JSON
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) convertDataFramIntoJSON(testdataDataframe dataframe.DataFrame) (json string) {

	var a io.Writer

	err := testdataDataframe.WriteJSON(a)
	if err != nil {
		return ""
	}

	return ""
}


*/
