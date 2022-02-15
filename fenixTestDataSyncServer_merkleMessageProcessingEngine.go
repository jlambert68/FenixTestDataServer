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

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "aa1d5eb1-7503-467f-9336-1927fa5529f2",
	}).Debug("Incoming gRPC 'convertgRpcHeaderMessageToStringArray'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "72779d91-4b28-44ad-b3b7-4beeed915cc9",
	}).Debug("Outgoing gRPC 'convertgRpcHeaderMessageToStringArray'")

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

// Concartenate TestDataRows as a DataFrame with the current Server TestDataRows-DataFrame
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) concatenateWithCurrentServerTestData(testDataClientGuid string, testdataDataframe dataframe.DataFrame) (allTestdataAsDataFrame dataframe.DataFrame) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "75e3cea1-d459-4f4f-9289-b728feee7ec0",
	}).Debug("Incoming gRPC 'concatenateWithCurrentServerTestData'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "f77a358e-50ad-4f1c-a540-6559688e173f",
	}).Debug("Outgoing gRPC 'concatenateWithCurrentServerTestData'")

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "5716a1e8-11dc-4c70-9579-94d7d177689b",
	}).Debug("New rows to add to existing rows are '"+strconv.Itoa(testdataDataframe.Nrow())+"' for Client: ", testDataClientGuid)

	allTestdataAsDataFrame = fenixTestDataSyncServerObject.getCurrentTestDataRowItemsForClient(testDataClientGuid)

	if allTestdataAsDataFrame.Nrow() == 0 {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "a98bf92e-ec7d-4968-b8fa-b72e47fef830",
		}).Debug("There are no TestDataRows in memDB for so returning new Rows to work with, Client: ", testDataClientGuid)

		return testdataDataframe
	} else {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "c8385ebf-53e1-4449-8171-f0c5bc2bdd68",
		}).Debug("Existing number of rows are '"+strconv.Itoa(allTestdataAsDataFrame.Nrow())+"' for Client: ", testDataClientGuid)

		//headerKeys := testdataDataframe.Names()
		allTestdataAsDataFrame = allTestdataAsDataFrame.Concat(testdataDataframe) //, headerKeys...)

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "c8385ebf-53e1-4449-8171-f0c5bc2bdd68",
		}).Debug("Concatenated number of rows are '"+strconv.Itoa(allTestdataAsDataFrame.Nrow())+"' for Client: ", testDataClientGuid)

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
