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
var memoryDB memoryDBStruct

type memoryDBStruct struct {
	allowedClients memDBAllowedClientsStruct
	client         memDBClientStruct
	server         memDBServerStruct
}

type memDBClientStruct struct {
	memDBTestDataDomain memDBTestDataDomainType
}

type memDBServerStruct struct {
	memDBTestDataDomain memDBTestDataDomainType
}

type memDBTestDataDomainType map[memDBDomainUuidType]memDBTestDataContainerType
type memDBTestDataContainerType map[memDBClientUuidType]memDBTestDataStruct

type memDBClientUuidType string
type memDBDomainUuidType string

type memDBTestDataStruct struct {
	memDBDataStructure memDBDataStructureStruct
}

type memDBDataStructureStruct struct {
	merkleHash          string
	merklePath          string
	merkleTree          memDBMerkleTreeRowsStruct
	headerItemsHash     string
	headerLabelsHash    string
	headerItems         []memDBHeaderItemsStruct
	testDataRows        []memDBTestDataItemsStruct
	testDataAsDataFrame dataframe.DataFrame
}

type memDBMerkleTreeRowsStruct struct {
	merkleTreeRows []memDBMerkleTreeRowStruct
}

type memDBMerkleTreeRowStruct struct {
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

type memDBAllowedClientsStruct struct {
	memDBTestDataDomainType map[memDBClientUuidType]memDBDomainUuidType
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForClient(testDataClientGuid string) string {

	var currentMerkleHashForClient string

	currentMerkleHashForClient = dbData.clientData.merkleHash

	fmt.Println("dbData.clientData.merkleHash: ", dbData.clientData.merkleHash)

	return currentMerkleHashForClient
}

// Save current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.clientData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForServer(testDataClientGuid string) string {

	var currentMerkleHashForServer string

	currentMerkleHashForServer = dbData.serverData.merkleHash

	fmt.Println("dbData.serverData.merkleHash: ", dbData.serverData.merkleHash)

	return currentMerkleHashForServer
}

// Save current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	dbData.serverData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeForClient(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForClient dataframe.DataFrame

	currentMerkleTreeForClient = dbData.clientData.merkleTree

	return currentMerkleTreeForClient
}

// Save current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeForClient(testDataClientGuid string, merkleTreeDataFrameMessage dataframe.DataFrame) bool {

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeForServer(testDataClientGuid string) dataframe.DataFrame {

	var currentMerkleTreeForServer dataframe.DataFrame

	currentMerkleTreeForServer = dbData.serverData.merkleTree

	return currentMerkleTreeForServer
}

// Save current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeForServer(merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	dbData.serverData.merkleTree = merkleTreeDataFrameMessage

	return true
}

// Retrieve current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashsForClient(testDataClientGuid string) string {

	var currentHeaderHashsForClient string

	currentHeaderHashsForClient = dbData.clientData.headerHash

	fmt.Println("dbData.clientData.headerHash: ", dbData.clientData.headerHash)

	return currentHeaderHashsForClient
}

// Save current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderHashForClient(testDataClientGuid string, currentHeaderHashsForClient string) bool {

	dbData.clientData.headerHash = currentHeaderHashsForClient

	return true
}

// Retrieve currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeadersForClient(testDataClientGuid string) []string {

	var currentHeadersForClient []string

	currentHeadersForClient = dbData.clientData.headers

	return currentHeadersForClient
}

// Save currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeadersForClient(testDataClientGuid string, testDataHeaderItems []string) bool {

	dbData.clientData.headers = testDataHeaderItems

	return true
}

// Retrieve current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashForServer(testDataClientGuid string) string {

	var currentHeaderHashsForServer string

	currentHeaderHashsForServer = dbData.serverData.headerHash

	fmt.Println("dbData.serverData.headerHash: ", dbData.serverData.headerHash)

	return currentHeaderHashsForServer
}

// Save current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderHashForServer(currentHeaderHashForServer string) bool {

	dbData.serverData.headerHash = currentHeaderHashForServer

	return true
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeadersForServer() []string {

	var currentHeadersForServer []string

	currentHeadersForServer = dbData.serverData.headers

	return currentHeadersForServer
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeadersForServer(currentHeadersForServer []string) bool {

	dbData.serverData.headers = currentHeadersForServer

	return true
}

// Transfer TestDataHeader from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) moveCurrentHeadersFromClientToServer(testDataClientGuid string) bool {

	dbData.serverData.headerHash = dbData.clientData.headerHash
	dbData.serverData.headers = dbData.clientData.headers

	return true
}

// Transfer MerkleTree and TestDataRows from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) moveCurrentTestDataAndMerkleTreeFromClientToServer(testDataClientGuid string) bool {

	dbData.serverData.merkleHash = dbData.clientData.merkleHash
	dbData.serverData.merkleTree = dbData.clientData.merkleTree
	dbData.serverData.testDataRows = dbData.clientData.testDataRows

	return true
}

// Get current Server TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowsForServer(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	testDataRows = dbData.serverData.testDataRows

	return testDataRows
}

// Get current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowsForClient(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	testDataRows = dbData.clientData.testDataRows

	return testDataRows
}

// Save current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentTestDataRowsForClient(testDataClientGuid string, testDataRows dataframe.DataFrame) bool {

	dbData.clientData.testDataRows = testDataRows

	return true
}
