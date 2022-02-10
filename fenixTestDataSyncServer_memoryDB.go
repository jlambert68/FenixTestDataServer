package main

import (
	"errors"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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
	merkleHash       string
	merkleTree       dataframe.DataFrame
	merkleFilter     string
	merkleFilterHash string
	headerHash       string
	headers          []string
	testDataRows     dataframe.DataFrame //[]tempTestDataRowStruct
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

	// New Structure that will overtake the above
	memCloudDBStructureForTestDataMerkleHash cloudDBTestDataMerkleHashStruct
	memCloudDBStructureForTestDataMerkleTree cloudDBTestDataMerkleTreeStruct

	memCloudDBStructureForTestDataHeaderItemsHashesStruct cloudDBTestDataHeaderItemsHashesStruct
	memCloudDBStructureForTestDataHeaderItem              []cloudDBTestDataHeaderItemStruct
	memCloudDBAllTestDataHeaderFilterValueStruct cloudDBTestDataHeaderFilterValuesStruct
	lägg                                         order för värden som det finns fler av, typ header, rowdata, filter
}
används verkligen FilterValueHash???
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
	leafNodeHash           string
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

// Retrieve current TestData-MerkleFilterHash for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterHashForClient(testDataClientGuid string) (currentMerkleFilterHashForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterHashForClient = tempdbData.clientData.merkleFilterHash
	} else {
		currentMerkleFilterHashForClient = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterHashForClient
}

// Save current TestData-MerkleFilterHash for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterHashForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleFilterHash
		clientData.merkleFilterHash = merkleHashMessage.MerkleFilterHash

	} else {

		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.merkleFilterHash = merkleHashMessage.MerkleFilterHash

		dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleFilter for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterForClient(testDataClientGuid string) (currentMerkleFilterForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterForClient = tempdbData.clientData.merkleFilter
	} else {
		currentMerkleFilterForClient = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterForClient
}

// Save current TestData-MerkleFilter for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterForClient(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleFilter
		clientData.merkleFilter = merkleHashMessage.MerkleFilter

	} else {

		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.merkleFilter = merkleHashMessage.MerkleFilter

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

// Retrieve current TestData-MerkleFilterHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterHashForServer(testDataClientGuid string) (currentMerkleFilterHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterHashForServer = tempdbData.serverData.merkleFilterHash
	} else {
		currentMerkleFilterHashForServer = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterHashForServer
}

// Save current TestData-MerkleFilterHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterHashForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// MerkleHash
	serverData.merkleFilterHash = merkleHashMessage.MerkleFilterHash

	return true
}

// Retrieve current TestData-MerkleFilter for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterForServer(testDataClientGuid string) (currentMerkleFilterForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterForServer = tempdbData.serverData.merkleFilter
	} else {
		currentMerkleFilterForServer = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterForServer
}

// Save current TestData-MerkleFilter for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterForServer(merkleHashMessage fenixTestDataSyncServerGrpcApi.MerkleHashMessage) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(merkleHashMessage.TestDataClientUuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// MerkleHashFilter
	serverData.merkleFilter = merkleHashMessage.MerkleFilter

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

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "082dff5e-33ac-4edd-94a4-37a7d0aaf8f7",
	}).Debug("Incoming gRPC 'moveCurrentHeaderDataFromClientToServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "437fb222-757a-4771-a7d7-d7eefdd81951",
	}).Debug("Outgoing gRPC 'moveCurrentHeaderDataFromClientToServer'")

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Move the data
	tempdbData.serverData.headerHash = tempdbData.clientData.headerHash
	tempdbData.serverData.headers = tempdbData.clientData.headers

	return true
}

// Transfer MerkleTree and TestDataRows from Client to Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) moveCurrentTestDataAndMerkleTreeFromClientToServer(testDataClientGuid string) bool {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "067df05e-19e8-4ac6-9c8d-ab3ef70f79d3",
	}).Debug("Incoming gRPC 'moveCurrentTestDataAndMerkleTreeFromClientToServer'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b31a193e-548d-459a-ba62-0af4982a1122",
	}).Debug("Outgoing gRPC 'moveCurrentTestDataAndMerkleTreeFromClientToServer'")

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Move the data
	tempdbData.serverData.merkleHash = tempdbData.clientData.merkleHash
	tempdbData.serverData.merkleFilter = tempdbData.clientData.merkleFilter
	tempdbData.serverData.merkleFilterHash = tempdbData.clientData.merkleFilterHash
	tempdbData.serverData.merkleTree = tempdbData.clientData.merkleTree
	tempdbData.serverData.testDataRows = tempdbData.clientData.testDataRows

	fenixTestDataSyncServerObject.testSQL(testDataClientGuid)

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

// Get the Domain for the Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getDomainUuidForClient(testDataClientGuid string) (domainUuid memDBDomainUuidType) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := cloudDBClientsMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		domainUuid = tempdbData.domainUuid
	} else {
		domainUuid = "666"

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "5617e51e-36ce-4c71-a10d-836c2eb604ee",
		}).Fatalln("Should not happen. Existing Client in memoryDB for Domain-Client is missing")

	}

	return domainUuid
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

// Verify if Client exists in Memory-copy of CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) existsClientInDB(testDataClientUUID string) (bool, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "214a4c17-c2e5-42bc-9acb-8b8164efcb26",
	}).Debug("Entering: ListTestInstructionsInDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "de2427f0-1ecf-42b9-823a-c3e677ba8ef0",
		}).Debug("Exiting: ListTestInstructionsInDB()")
	}()

	_, clientExits := cloudDBClientsMap[memDBClientUuidType(testDataClientUUID)]

	//Only OK if client exists in memoryDBMap
	if clientExits != true {

		myErr := errors.New("Client '" + testDataClientUUID + "' does not exit in MemoryDB")
		return false, myErr

	} else {

		return true, nil

	}

}

// Remove the rows that don't is represented in the Clients MerkleTree
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) removeTestDataRowsInMemoryDBThatIsNotRepresentedInClientsMerkleTree(callingClientUuid string) bool {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "2712573d-354a-4fb9-a510-39b3844fe95d",
	}).Debug("Entering: 'removeTestDataRowsInMemoryDBThatIsNotRepresentedInClientsMerkleTree'")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "0fcc0362-263a-4945-a400-6e22baf53240",
		}).Debug("Exiting: 'removeTestDataRowsInMemoryDBThatIsNotRepresentedInClientsMerkleTree'")
	}()

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(callingClientUuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Extract the hashes for all leaf nodes from MerkleTree
	leafNodesHashes := clientData.merkleTree.Col(("MerkleChildHash")).Records()

	// Get the TestData to process
	clientTestData := clientData.testDataRows

	// Filter
	//isNotInListFkn := common_config.IsNotInListFilter(leafNodesHashes)
	clientTestData = clientTestData.Filter(
		dataframe.F{
			Colname:    "LeafNodeHash",
			Comparator: series.In,       // series.CompFunc,
			Comparando: leafNodesHashes, //isNotInListFkn(),
		})

	clientData.testDataRows = clientTestData

	return true
}
