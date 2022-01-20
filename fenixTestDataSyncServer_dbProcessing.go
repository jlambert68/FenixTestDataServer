package main

import (
	"github.com/go-gota/gota/dataframe"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
)

// ***** Before I implement a DB everthing will be stored in variables *****
/*
var dbCurrentMerkleHashForClient string
var dbCurrentMerkleTreeForClient dataframe.DataFrame
var dbCurrentHeaderHashsForClient string
var dbCurrentHeadersForClient []string
*/

var dbData tempDB

type tempTestDataRowStruct struct {
	rowHash       string
	testDataValue []string
}

type tempDBDataStruct struct {
	merkleHash   string
	merkleTree   dataframe.DataFrame
	headerHash   string
	headers      []string
	testDataRows []tempTestDataRowStruct
}
type tempDB struct {
	serverData tempDBDataStruct
	clientData tempDBDataStruct
}

type MerkleTree_struct struct {
	MerkleLevel     int
	MerklePath      string
	MerkleHash      string
	MerkleChildHash string
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleHashForClient(testDataClientGuid string) string {

	var currentMerkleHashForClient string

	currentMerkleHashForClient = dbData.clientData.merkleHash

	return currentMerkleHashForClient
}

// Save current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleHashForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.clientData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleHashForServer(testDataClientGuid string) string {

	var currentMerkleHashForServer string

	currentMerkleHashForServer = dbData.serverData.merkleHash

	return currentMerkleHashForServer
}

// Save current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleHashForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.serverData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleTreeForClient(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForClient dataframe.DataFrame

	currentMerkleTreeForClient = dbData.clientData.merkleTree

	return currentMerkleTreeForClient
}

// Save current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleTreeForClient(testDataClientGuid string, merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	dbData.clientData.merkleTree = merkleTreeDataFrameMessage

	/*
		m := jsonpb.Marshaler{}
		js, err := m.MarshalToString(&merkleTreeMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(js)
		}
	*/

	return true
}

// Retrieve current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleTreeForServer(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForServer dataframe.DataFrame

	currentMerkleTreeForServer = dbData.serverData.merkleTree

	return currentMerkleTreeForServer
}

// Save current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentMerkleTreeForServer(merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	dbData.serverData.merkleTree = merkleTreeDataFrameMessage

	return true
}

// Retrieve current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeaderHashsForClient(testDataClientGuid string) string {

	var currentHeaderHashsForClient string

	currentHeaderHashsForClient = dbData.clientData.headerHash

	return currentHeaderHashsForClient
}

// Save current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeaderHashForClient(testDataClientGuid string, currentHeaderHashsForClient string) bool {

	dbData.clientData.headerHash = currentHeaderHashsForClient

	return true
}

// Retrieve currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeadersForClient(testDataClientGuid string) []string {

	var currentHeadersForClient []string

	currentHeadersForClient = dbData.clientData.headers

	return currentHeadersForClient
}

// Save currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeadersForClient(testDataClientGuid string, testDataHeaderItems []string) bool {

	dbData.clientData.headers = testDataHeaderItems

	return true
}

// Retrieve current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeaderHashForServer(testDataClientGuid string) string {

	var currentHeaderHashsForServer string

	currentHeaderHashsForServer = dbData.serverData.headerHash

	return currentHeaderHashsForServer
}

// Save current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeaderHashForServer(currentHeaderHashForServer string) bool {

	dbData.serverData.headerHash = currentHeaderHashForServer

	return true
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentHeadersForServer() []string {

	var currentHeadersForServer []string

	currentHeadersForServer = dbData.serverData.headers

	return currentHeadersForServer
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentHeadersForServer(currentHeadersForServer []string) bool {

	dbData.serverData.headers = currentHeadersForServer

	return true
}

// Transfer TestDataHeader from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) moveCurrentHeadersFromClientToServer(testDataClientGuid string) bool {

	dbData.serverData.headerHash = dbData.clientData.headerHash
	dbData.serverData.headers = dbData.clientData.headers

	return true
}

// Transfer MerkleTree and TestDataRows from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) moveCurrentTestDataAndMerkleTreeFromClientToServer(testDataClientGuid string) bool {

	dbData.serverData.merkleHash = dbData.clientData.merkleHash
	dbData.serverData.merkleTree = dbData.clientData.merkleTree
	dbData.serverData.testDataRows = dbData.clientData.testDataRows

	return true
}
