package main

import (
	"fmt"
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
	testDataRows dataframe.DataFrame //[]tempTestDataRowStruct
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

// Memory Object used as temporary storage before saving testdata to Cloud-DB
// Memory Object also used as cash and there by minimize DB-access
var memoryDB memDBTestDataDomainType

type memDBTestDataDomainType map[memDBDomainUuidType]memDBTestDataContainerType
type memDBTestDataContainerType map[memDBClientUuidType]memDBTestDataStruct

type memDBClientUuidType string
type memDBDomainUuidType string

type memDBTestDataStruct struct {
	serverData memDBDataStructureStruct
	clientData memDBDataStructureStruct
}

type memDBDataStructureStruct struct {
	merkleHash       string
	merklePath       string
	merkleTree       memDBMerkleTreeStruct
	headerItemsHash  string
	headerLabelsHash string
	headerItems      memDBHeaderItemsStruct
	testDataRows     memDBTestDataItemsStruct
}

type memDBMerkleTreeStruct struct {
	nodeLevel     string
	nodeName      string
	nodePath      string
	nodeHash      string
	nodeChildHash string
}

type memDBHeaderItemsStruct struct {
	headerItemHash             string
	headerLabel                string
	headerShouldBeUsedInFilter bool
	headerIsMandatoryInFilter  bool
	headerFilterSelectionType  HeaderFilterSelectionTypeType
	headerFilterValuesItem     memDBHeaderFilterValuesItemStruct
}

type HeaderFilterSelectionTypeType int

const (
	HEADER_IS_SINGLE_SELECT HeaderFilterSelectionTypeType = iota
	HEADER_IS_MULTI_SELECT
)

type memDBHeaderFilterValuesItemStruct struct {
	HeaderFilterValuesHash string
	HeaderFilterValues     []string
}

type memDBTestDataItemsStruct struct {
	testDataRowHash        string
	leafNodeName           string
	leafNodePath           string
	testDataValuesAsString []string
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentMerkleHashForClient(testDataClientGuid string) string {

	var currentMerkleHashForClient string

	currentMerkleHashForClient = dbData.clientData.merkleHash

	fmt.Println("dbData.clientData.merkleHash: ", dbData.clientData.merkleHash)

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

	fmt.Println("dbData.serverData.merkleHash: ", dbData.serverData.merkleHash)

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

	fmt.Println("dbData.clientData.headerHash: ", dbData.clientData.headerHash)

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

	fmt.Println("dbData.serverData.headerHash: ", dbData.serverData.headerHash)

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

// Get current Server TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentTestDataRowsForServer(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	testDataRows = dbData.serverData.testDataRows

	return testDataRows
}

// Get current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getCurrentTestDataRowsForClient(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	testDataRows = dbData.clientData.testDataRows

	return testDataRows
}

// Save current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) saveCurrentTestDataRowsForClient(testDataClientGuid string, testDataRows dataframe.DataFrame) bool {

	dbData.clientData.testDataRows = testDataRows

	return true
}
