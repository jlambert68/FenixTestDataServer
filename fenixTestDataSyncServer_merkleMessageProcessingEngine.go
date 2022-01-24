package main

import (
	"FenixTestDataServer/common_config"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
)

/*
// Initiate channels for "incomming gRPC-messages"
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) InitateProcessEngineChannels() {

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcMerkleTreeMessageToDataframe(merkleTreeMessage fenixTestDataSyncServerGrpcApi.MerkleTreeMessage) dataframe.DataFrame {

	var myMerkleTree []MerkleTree_struct

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	merkleTreeNodes := merkleTreeMessage.MerkleTreeNodes

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, merkleTreeNode := range merkleTreeNodes {
		myMerkleTreeRow := MerkleTree_struct{
			MerkleLevel:     int(merkleTreeNode.MerkleLevel),
			MerklePath:      merkleTreeNode.MerklePath,
			MerkleHash:      merkleTreeNode.MerkleHash,
			MerkleChildHash: merkleTreeNode.MerkleChildHash,
		}
		myMerkleTree = append(myMerkleTree, myMerkleTreeRow)

	}

	df := dataframe.LoadStructs(myMerkleTree)

	return df
}

// Convert gRPC-Header message into string and string array objects
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcHeaderMessageToStringArray(testDataHeaderMessage fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage) (headerHash string, headersItems []string) {

	// Extract  HeaderHash
	headerHash = testDataHeaderMessage.HeadersHash

	//dbCurrentMerkleTreeForClient = merkleTreeMessage.MerkleTreeNodes
	testDataHeaderItems := testDataHeaderMessage.TestDataHeaderItems

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, headerItem := range testDataHeaderItems {
		headersItems = append(headersItems, headerItem.HeaderPresentationsLabel)

	}

	return headerHash, headersItems
}

// Convert TestDataRow message into TestData dataframe object
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcTestDataRowsMessageToDataFrame(testdataRowsMessages *fenixTestDataSyncServerGrpcApi.TestdataRowsMessages) (testdataAsDataFrame dataframe.DataFrame, returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	testdataAsDataFrame = dataframe.New()

	currentTestDataClientGuid := testdataRowsMessages.TestDataClientGuid

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
				Acknack:    false,
				Comments:   "Fenix Asked for TestDataHeaders but didn't receive them i a correct way",
				ErrorCodes: errorCodes,
			}

			// leave
			return testdataAsDataFrame, returnMessage
		}
	}

	testDataRows := testdataRowsMessages.TestDataRows

	// Loop all MerkleTreeNodes and create a DataFrame for the data
	for _, testDataRow := range testDataRows {

		// Create one row, as a dataframe
		rowDataframe := dataframe.New(
			series.New([]string{"key"}, series.String, "KEY"))
		var valuesToHash []string

		for testDataItemCounter, testDataItem := range testDataRow.TestDataItems {
			rowDataframe = rowDataframe.Mutate(
				series.New([]string{testDataItem.TestDataItemValueAsString}, series.String, currentTestDataHeaders[testDataItemCounter]))

			valuesToHash = append(valuesToHash, testDataItem.TestDataItemValueAsString)
		}

		// Hash all values for row
		hashedRow := common_config.HashValues(valuesToHash)

		// Validate that Row-hash is correct calculated
		if hashedRow != testDataRow.RowHash {

			// Set Error codes to return message
			var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
			var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

			errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_ROWHASH_NOT_CORRECT_CALCULATED
			errorCodes = append(errorCodes, errorCode)

			// Create Return message
			returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
				Acknack:    false,
				Comments:   "RowsHashes seems not to be correct calculated",
				ErrorCodes: errorCodes,
			}

			// Exit function Respond back to client when hash error
			return testdataAsDataFrame, returnMessage
		}

		// Add TestDataHash to row DataFrame
		rowDataframe.Elem(0, 0).Set(hashedRow)
		//) Mutate(
		//	series.New([]string{hashedRow}, series.String, "TestDataHash"))

		// Add the row to the Dataframe for the testdata
		// When first time
		if testdataAsDataFrame.Nrow() == 0 {
			testdataAsDataFrame = rowDataframe.Copy()
		} else {
			testdataAsDataFrame = testdataAsDataFrame.OuterJoin(rowDataframe, "KEY")
		}
	}

	return testdataAsDataFrame, nil

}

// Concartenate TestDataRows as a DataFrame with the current Server TestDataRows-DataFrame
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) concartenateWithCurrentServerTestData(testDataClientGuid string, testdataDataframe dataframe.DataFrame) (allTestdataAsDataFrame dataframe.DataFrame) {

	allTestdataAsDataFrame = fenixTestDataSyncServerObject.getCurrentTestDataRowsForServer(testDataClientGuid)

	allTestdataAsDataFrame = allTestdataAsDataFrame.OuterJoin(testdataDataframe)

	return allTestdataAsDataFrame

}
