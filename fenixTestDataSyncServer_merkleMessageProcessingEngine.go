package main

import (
	"github.com/go-gota/gota/dataframe"
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
/*
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) convertgRpcTestDataRowsMessageToDataFrame(testDataHeaderMessage fenixTestDataSyncServerGrpcApi.TestDataHeaderMessage) (testdataAsDataFrame dataframe.DataFrame {

	jsonStr := `[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`
	df := dataframe.ReadJSON(strings.NewReader(jsonStr))


	return testdataAsDataFrame

}


*/
