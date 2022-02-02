package main

import (
	"errors"
	"github.com/go-gota/gota/dataframe"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// ***** Before I implement a DB everthing will be stored in variables *****
/*
var dbCurrentMerkleHashForClient string
var dbCurrentMerkleTreeForClient dataframe.DataFrame
var dbCurrentHeaderHashsForClient string
var dbCurrentHeadersForClient []string
*/

var dbDataMap map[memDBClientUuidType]*tempDBStruct

//var dbData tempDBStruct

//type tempDBMap map[memDBClientUuidType]tempDBStruct

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
type tempDBStruct struct {
	serverData *tempDBDataStruct
	clientData *tempDBDataStruct
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

// Initiate tempDBStruct
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) initiateTempDBStruct() tempDBStruct {
	tempDBStruct_initiated := tempDBStruct{
		serverData: &tempDBDataStruct{
			merkleHash: "",
			merkleTree: dataframe.DataFrame{
				Err: nil,
			},
			headerHash: "",
			headers:    nil,
			testDataRows: dataframe.DataFrame{
				Err: nil,
			},
		},
		clientData: &tempDBDataStruct{
			merkleHash: "",
			merkleTree: dataframe.DataFrame{
				Err: nil,
			},
			headerHash: "",
			headers:    nil,
			testDataRows: dataframe.DataFrame{
				Err: nil,
			},
		},
	}

	return tempDBStruct_initiated
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForClient(testDataClientGuid string) (currentMerkleHashForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleHashForClient = tempdbData.clientData.merkleHash
	} else {
		currentMerkleHashForClient = "#VALUE IS MISSING#"
	}

	return currentMerkleHashForClient
}

// Save current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleHash
		clientData.merkleHash = merkleHashMessage.MerkleHash

	} else {

		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.merkleHash = merkleHashMessage.MerkleHash

		dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForServer(testDataClientGuid string) (currentMerkleHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleHashForServer = tempdbData.serverData.merkleHash
	} else {
		currentMerkleHashForServer = "#VALUE IS MISSING#"
	}

	return currentMerkleHashForServer
}

// Save current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// MerkleHash
	serverData.merkleHash = merkleHashMessage.MerkleHash

	return true
}

// Retrieve current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeForClient(testDataClientGuid string) (currentMerkleTreeForClient dataframe.DataFrame) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleTreeForClient = tempdbData.clientData.merkleTree
	} else {
		currentMerkleTreeForClient = dataframe.DataFrame{}
	}

	return currentMerkleTreeForClient
}

// Save current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeForClient(testDataClientGuid string, merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.merkleTree = merkleTreeDataFrameMessage

	return true
}

// Retrieve current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeForServer(testDataClientGuid string) (currentMerkleTreeForServer dataframe.DataFrame) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleTreeForServer = tempdbData.serverData.merkleTree
	} else {
		currentMerkleTreeForServer = dataframe.DataFrame{}
	}

	return currentMerkleTreeForServer
}

// Save current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeForServer(testDataClientGuid string, merkleTreeDataFrameMessage dataframe.DataFrame) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.merkleTree = merkleTreeDataFrameMessage

	return true
}

// Retrieve current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashForClient(testDataClientGuid string) (currentHeaderHashsForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeaderHashsForClient = tempdbData.clientData.headerHash
	} else {
		currentHeaderHashsForClient = "#VALUE IS MISSING#"
	}

	return currentHeaderHashsForClient
}

// Save current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderHashForClient(testDataClientGuid string, currentHeaderHashsForClient string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.headerHash = currentHeaderHashsForClient

	return true
}

// Retrieve currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeadersForClient(testDataClientGuid string) (currentHeadersForClient []string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeadersForClient = tempdbData.clientData.headers
	} else {
		currentHeadersForClient = []string{}
	}

	return currentHeadersForClient
}

// Save currentTestData-Headers for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeadersForClient(testDataClientGuid string, testDataHeaderItems []string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.headers = testDataHeaderItems

	return true
}

// Retrieve current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashForServer(testDataClientGuid string) (currentHeaderHashsForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeaderHashsForServer = tempdbData.serverData.headerHash
	} else {
		currentHeaderHashsForServer = "#VALUE IS MISSING#"
	}

	return currentHeaderHashsForServer
}

// Save current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderHashForServer(testDataClientGuid string, currentHeaderHashForServer string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.headerHash = currentHeaderHashForServer

	return true
}

// Retrieve current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeadersForServer(testDataClientGuid string) (currentHeadersForServer []string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeadersForServer = tempdbData.serverData.headers
	} else {
		currentHeadersForServer = []string{}
	}

	return currentHeadersForServer
}

// Save current TestData-Header for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeadersForServer(testDataClientGuid string, currentHeadersForServer []string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.headers = currentHeadersForServer

	return true
}

// Transfer TestDataHeader from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) moveCurrentHeaderDataFromClientToServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Move the data
	tempdbData.serverData.headerHash = tempdbData.clientData.headerHash
	tempdbData.serverData.headers = tempdbData.clientData.headers

	return true
}

// Transfer MerkleTree and TestDataRows from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) moveCurrentTestDataAndMerkleTreeFromClientToServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Move the data
	tempdbData.serverData.merkleHash = tempdbData.clientData.merkleHash
	tempdbData.serverData.merkleTree = tempdbData.clientData.merkleTree
	tempdbData.serverData.testDataRows = tempdbData.clientData.testDataRows

	return true
}

// Get current Server TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowsForServer(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		testDataRows = tempdbData.serverData.testDataRows
	} else {
		testDataRows = dataframe.DataFrame{}
	}

	return testDataRows
}

// Get current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowsForClient(testDataClientGuid string) (testDataRows dataframe.DataFrame) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		testDataRows = tempdbData.clientData.testDataRows
	} else {
		testDataRows = dataframe.DataFrame{}
	}

	return testDataRows
}

// Save current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentTestDataRowsForClient(testDataClientGuid string, testDataRows dataframe.DataFrame) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.testDataRows = testDataRows

	return true
}

// Veriy if Client exists in Memory-copy of CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) existsClientInDB(testDataClientUUID string) (bool, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "214a4c17-c2e5-42bc-9acb-8b8164efcb26",
	}).Debug("Entering: ListTestInstructionsInDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "de2427f0-1ecf-42b9-823a-c3e677ba8ef0",
		}).Debug("Exiting: ListTestInstructionsInDB()")
	}()

	_, clientExits := memCloudDBAllClientsMap[memDBClientUuidType(testDataClientUUID)]

	//Only OK if client exists in memoryDBMap
	if clientExits != true {

		myErr := errors.New("Client '" + testDataClientUUID + "' does not exit in MemoryDB")
		return false, myErr

	} else {

		return true, nil

	}

}
