package main

import (
	"github.com/go-gota/gota/dataframe"
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
	merkleHash string
	//merkleTree           dataframe.DataFrame
	merkleTreeNodes                             []cloudDBTestDataMerkleTreeStruct
	merkleTreeNodesChildHashesThatNoLongerExist []string
	requestedMerkleNodeNamesFromClient          []string
	MerkleFilterPath                            string
	MerkleFilterPathHash                        string
	//headerHash                                  string
	headers []string
	//testDataRows         dataframe.DataFrame //[]tempTestDataRowStruct
	testDataRowItems            []cloudDBTestDataRowItemCurrentStruct
	testDataHeaderItemsHashes   cloudDBTestDataHeaderItemsHashesStruct
	testDataHeaderItems         []cloudDBTestDataHeaderItemStruct
	testDataHeadersFilterValues []cloudDBTestDataHeaderFilterValuesStruct
}
type tempDBStruct struct {
	serverData *tempDBDataStruct
	clientData *tempDBDataStruct
}

type MerkletreeStruct struct {
	MerkleLevel     int
	MerklePath      string
	MerkleHash      string
	MerkleChildHash string
}

// Memory Object used as temporary storage before saving testdata to Cloud-DB
// Memory Object also used as cash and there by minimize DB-access
// TODO THis object is not used, I think --> REMOVE
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
	memCloudDBAllTestDataHeaderFilterValueStruct          cloudDBTestDataHeaderFilterValuesStruct
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

/*
const (
	HEADER_IS_SINGLE_SELECT HeaderFilterSelectionTypeType = iota
	HEADER_IS_MULTI_SELECT
)
*/

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

	tempdbstructInitiated := tempDBStruct{
		serverData: &tempDBDataStruct{
			merkleHash:      "",
			merkleTreeNodes: nil,
			merkleTreeNodesChildHashesThatNoLongerExist: nil,
			requestedMerkleNodeNamesFromClient:          nil,
			MerkleFilterPath:                            "",
			MerkleFilterPathHash:                        "",
			headers:                                     nil,
			testDataRowItems:                            nil,
			testDataHeaderItemsHashes: cloudDBTestDataHeaderItemsHashesStruct{
				headerItemsHash:  "",
				clientUuid:       "",
				headerLabelsHash: "",
				updatedTimeStamp: "",
			},
			testDataHeaderItems:         nil,
			testDataHeadersFilterValues: nil,
		},
		clientData: &tempDBDataStruct{
			merkleHash:      "",
			merkleTreeNodes: nil,
			merkleTreeNodesChildHashesThatNoLongerExist: nil,
			requestedMerkleNodeNamesFromClient:          nil,
			MerkleFilterPath:                            "",
			MerkleFilterPathHash:                        "",
			headers:                                     nil,
			testDataRowItems:                            nil,
			testDataHeaderItemsHashes: cloudDBTestDataHeaderItemsHashesStruct{
				headerItemsHash:  "",
				clientUuid:       "",
				headerLabelsHash: "",
				updatedTimeStamp: "",
			},
			testDataHeaderItems:         nil,
			testDataHeadersFilterValues: nil,
		},
	}

	return tempdbstructInitiated
}

// Initiate all memoryDB for all clients found in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) initiateMemoryDBForClients() bool {

	// Loop over all loaded clients
	for _, client := range cloudDBClients {

		// Create empty struct and add it for client_uuid
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		dbDataMap[memDBClientUuidType(client.clientUuid)] = &tempDBStructInitiated

	}

	return true
}

// Retrieve current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForClient(testDataClientGuid string) (currentMerkleHashForClient string) {
	//TODO Denna är fel och går mot server istället för Client
	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data, and it must be any saved to be used as serve memoryDB-copy
	if valueExits == true && len(tempdbData.clientData.merkleHash) > 0 {
		currentMerkleHashForClient = tempdbData.clientData.merkleHash
	} else {

		// Load Client's TestDataMerkleHashes from CloudDB
		var tempMemDBAllTestDataMerkleHashes []cloudDBTestDataMerkleHashStruct
		err := fenixTestDataSyncServerObject.loadAllTestDataMerkleHashesForClientFromCloudDB(testDataClientGuid, &tempMemDBAllTestDataMerkleHashes)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "202bbd5b-19a1-4ebd-923d-0114161e6c2b",
				"error": err,
			}).Error("Problem when executing: 'loadAllTestDataMerkleHashesForClientFromCloudDB()'")

			fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
			currentMerkleHashForClient = "#VALUE IS MISSING#"

		} else {

			// Verify that a maximum of one MerkleHash-object has been retrieved
			switch len(tempMemDBAllTestDataMerkleHashes) {

			// No Saved MerkleHash in CloudDB for specified Client
			case 0:

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                 "7ee3595a-0457-4d0f-a1a3-c89b3fd34787",
					"testDataClientGuid": testDataClientGuid,
				}).Debug("No Saved MerkleHash in CloudDB for specified Client")

			// MerkleHash has previously been saved and loaded from CloudDB
			case 1:

				// Save MerkleHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleHash)

				// save MerkleFilterPathHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPathHash)

				// save MerkleFilterPathHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPathHash)

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                               "08d8c4d7-2f39-4a99-b94d-dd44f705172d",
					"testDataClientGuid":               testDataClientGuid,
					"tempMemDBAllTestDataMerkleHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Debug("MerkleHash data found in CloudDB")

			// There are more than one MerkleHash-object, which shouldn't happen
			default:
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                               "a9f37032-167d-4524-878e-814ffe27f012",
					"testDataClientGuid":               testDataClientGuid,
					"Number of MerkleHash-objecs":      len(tempMemDBAllTestDataMerkleHashes),
					"tempMemDBAllTestDataMerkleHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Fatalln("There are more than one MerkleHash-object found in CloudDB, which shouldn't happen")

			}

			currentMerkleHashForClient = tempMemDBAllTestDataMerkleHashes[0].merkleHash

		}
	}

	return currentMerkleHashForClient
}

// Save current TestData-MerkleHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForClient(testDataClientUuid string, merkleHash string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleHash
		clientData.merkleHash = merkleHash

	} else {

		// Initiate data structure for client
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.merkleHash = merkleHash

		dbDataMap[memDBClientUuidType(testDataClientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-requestedMerkleNodeNamesFromClient for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentRequestedMerkleNodeNamesFromClient(testDataClientGuid string) (currentRequestedMerkleNodeNamesFromClient []string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentRequestedMerkleNodeNamesFromClient = tempdbData.clientData.requestedMerkleNodeNamesFromClient
	} else {
		currentRequestedMerkleNodeNamesFromClient = []string{}
	}

	return currentRequestedMerkleNodeNamesFromClient
}

// Save current TestData-requestedMerkleNodeNamesFromClient for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentRequestedMerkleNodeNamesFromClient(callingClientUuid string, currentRequestedMerkleNodeNamesFromClient []string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(callingClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleFilterPathHash
		clientData.requestedMerkleNodeNamesFromClient = currentRequestedMerkleNodeNamesFromClient

	} else {

		// Initiate data structure for client
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.requestedMerkleNodeNamesFromClient = currentRequestedMerkleNodeNamesFromClient

		dbDataMap[memDBClientUuidType(callingClientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleFilterPathHash for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterPathHashForClient(testDataClientGuid string) (currentMerkleFilterPathHashForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterPathHashForClient = tempdbData.clientData.MerkleFilterPathHash
	} else {
		currentMerkleFilterPathHashForClient = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterPathHashForClient
}

// Save current TestData-MerkleFilterPathHash for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterPathHashForClient(callingClientUuid string, merkleFilterPathHash string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(callingClientUuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleFilterPathHash
		clientData.MerkleFilterPathHash = merkleFilterPathHash

	} else {

		// Initiate data structure for client
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.MerkleFilterPathHash = merkleFilterPathHash

		dbDataMap[memDBClientUuidType(callingClientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleFilterPath for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterPathForClient(testDataClientGuid string) (currentMerkleFilterPathForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterPathForClient = tempdbData.clientData.MerkleFilterPath
	} else {
		currentMerkleFilterPathForClient = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterPathForClient
}

// Save current TestData-MerkleFilterPath for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterPathForClient(testDataClientGuid string, merkleFilterPath string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	if valueExits == true {
		// Get pointer to client data
		clientData := tempdbData.clientData

		// MerkleFilterPath
		clientData.MerkleFilterPath = merkleFilterPath

	} else {

		// Initiate data structure for client
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.clientData.MerkleFilterPath = merkleFilterPath

		dbDataMap[memDBClientUuidType(testDataClientGuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleHashForServer(testDataClientGuid string) (currentMerkleHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data, and it must be any saved to be used as serve memoryDB-copy
	if valueExits == true && len(tempdbData.clientData.merkleHash) > 0 {
		currentMerkleHashForServer = tempdbData.serverData.merkleHash
	} else {

		// Load Client's TestDataMerkleHashes from CloudDB
		var tempMemDBAllTestDataMerkleHashes []cloudDBTestDataMerkleHashStruct
		err := fenixTestDataSyncServerObject.loadAllTestDataMerkleHashesForClientFromCloudDB(testDataClientGuid, &tempMemDBAllTestDataMerkleHashes)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "202bbd5b-19a1-4ebd-923d-0114161e6c2b",
				"error": err,
			}).Error("Problem when executing: 'loadAllTestDataMerkleHashesForClientFromCloudDB()'")

			fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
			currentMerkleHashForServer = "#VALUE IS MISSING#"

		} else {

			// Verify that a maximum of one MerkleHash-object has been retrieved
			switch len(tempMemDBAllTestDataMerkleHashes) {

			// No Saved MerkleHash in CloudDB for specified Client
			case 0:

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                 "7ee3595a-0457-4d0f-a1a3-c89b3fd34787",
					"testDataClientGuid": testDataClientGuid,
				}).Debug("No Saved MerkleHash in CloudDB for specified Client")

				currentMerkleHashForServer = ""

			// MerkleHash has previously been saved and loaded from CloudDB
			case 1:

				// Save MerkleHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleHash)

				// Save MerkleFilterPathHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPathHash)

				// save MerkleFilterPath in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPath)

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                               "08d8c4d7-2f39-4a99-b94d-dd44f705172d",
					"testDataClientGuid":               testDataClientGuid,
					"tempMemDBAllTestDataMerkleHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Debug("MerkleHash data found in CloudDB")

				currentMerkleHashForServer = tempMemDBAllTestDataMerkleHashes[0].merkleHash

			// There are more than one MerkleHash-object, which shouldn't happen
			default:
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                               "39bb3486-dd0c-4655-8495-cb54be2571ad",
					"testDataClientGuid":               testDataClientGuid,
					"Number of MerkleHash-objecs":      len(tempMemDBAllTestDataMerkleHashes),
					"tempMemDBAllTestDataMerkleHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Fatalln("There are more than one MerkleHash-object found in CloudDB, which shouldn't happen")

			}

		}
	}

	return currentMerkleHashForServer

}

// Save current TestData-MerkleHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleHashForServer(clientUuid string, merkleHash string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(clientUuid)]

	if valueExits == true {

		// Get pointer to server data
		serverData := tempdbData.serverData

		// MerkleHash
		serverData.merkleHash = merkleHash

	} else {

		// Initiate data structure for server
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.serverData.merkleHash = merkleHash

		dbDataMap[memDBClientUuidType(clientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleFilterPathHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterPathHashForServer(testDataClientGuid string) (currentMerkleFilterPathHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data, and it must be any saved to be used as server memoryDB-copy
	if valueExits == true && len(tempdbData.serverData.MerkleFilterPathHash) > 0 {
		currentMerkleFilterPathHashForServer = tempdbData.serverData.MerkleFilterPathHash
	} else {

		// Load Client's TestDataMerkleFilterPathHashes from CloudDB
		var tempMemDBAllTestDataMerkleHashes []cloudDBTestDataMerkleHashStruct
		err := fenixTestDataSyncServerObject.loadAllTestDataMerkleHashesForClientFromCloudDB(testDataClientGuid, &tempMemDBAllTestDataMerkleHashes)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "2c41b7df-f66c-4ebc-8ed1-33b188cca613",
				"error": err,
			}).Error("Problem when executing: 'loadAllTestDataMerkleHashesForClientFromCloudDB()'")

			fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
			currentMerkleFilterPathHashForServer = "#VALUE IS MISSING#"

		} else {

			// Verify that a maximum of one MerkleFilterPathHash-object has been retrieved
			switch len(tempMemDBAllTestDataMerkleHashes) {

			// No Saved MerkleFilterPathHash in CloudDB for specified Client
			case 0:

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                 "1115e3c8-464b-437b-87b2-98bebd1a03d8",
					"testDataClientGuid": testDataClientGuid,
				}).Debug("No Saved MerkleFilterPathHash in CloudDB for specified Client")

				currentMerkleFilterPathHashForServer = "#VALUE IS MISSING#"

			// MerkleFilterPathHash has previously been saved and loaded from CloudDB
			case 1:

				// Save MerkleFilterPathHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPath)

				// save MerkleFilterPathHash in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathHashForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPathHash)

				// save MerkleFilterPath in memDB
				_ = fenixTestDataSyncServerObject.saveCurrentMerkleFilterPathForServer(testDataClientGuid, tempMemDBAllTestDataMerkleHashes[0].merkleFilterPath)

				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                               "b02aa02c-8786-471f-9ed7-8920d8859308",
					"testDataClientGuid":               testDataClientGuid,
					"tempMemDBAllTestDataMerkleHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Debug("MerkleFilterPathHash data found in CloudDB")

				currentMerkleFilterPathHashForServer = tempMemDBAllTestDataMerkleHashes[0].merkleFilterPathHash

			// There are more than one MerkleFilterPathHash-object, which shouldn't happen
			default:
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"Id":                                    "eed7acae-6acb-40c7-bd3c-60ddab9f194d",
					"testDataClientGuid":                    testDataClientGuid,
					"Number of MerkleFilterPathHash-objecs": len(tempMemDBAllTestDataMerkleHashes),
					"tempMemDBAllTestDataMerkleFilterPathHashes": tempMemDBAllTestDataMerkleHashes[0],
				}).Fatalln("There are more than one MerkleFilterPathHash-object found in CloudDB, which shouldn't happen")

			}

		}
	}

	return currentMerkleFilterPathHashForServer

}

// Save current TestData-MerkleFilterPathHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterPathHashForServer(clientUuid string, MerkleFilterPathHash string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(clientUuid)]

	if valueExits == true {

		// Get pointer to server data
		serverData := tempdbData.serverData

		// MerkleFilterPathHash
		serverData.MerkleFilterPathHash = MerkleFilterPathHash

	} else {

		// Initiate data structure for server
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.serverData.merkleHash = MerkleFilterPathHash

		dbDataMap[memDBClientUuidType(clientUuid)] = &tempDBStructInitiated
	}

	return true
}

// Retrieve current TestData-MerkleFilterPath for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleFilterPathForServer(testDataClientGuid string) (currentMerkleFilterPathForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleFilterPathForServer = tempdbData.serverData.MerkleFilterPath
	} else {
		currentMerkleFilterPathForServer = "#VALUE IS MISSING#"
	}

	return currentMerkleFilterPathForServer
}

// Save current TestData-MerkleFilterPath for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleFilterPathForServer(clientUuid string, merkleFilterPath string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(clientUuid)]

	if valueExits == true {

		// Get pointer to server data
		serverData := tempdbData.serverData

		// MerkleFilterPath
		serverData.MerkleFilterPath = merkleFilterPath

	} else {

		// Initiate data structure for server
		tempDBStructInitiated := fenixTestDataSyncServerObject.initiateTempDBStruct()
		tempDBStructInitiated.serverData.MerkleFilterPath = merkleFilterPath

		dbDataMap[memDBClientUuidType(clientUuid)] = &tempDBStructInitiated
	}

	return true

}

// Retrieve current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeNodesForClient(testDataClientGuid string) (currentMerkleTreeNodesForClient []cloudDBTestDataMerkleTreeStruct) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentMerkleTreeNodesForClient = tempdbData.clientData.merkleTreeNodes
	} else {
		currentMerkleTreeNodesForClient = []cloudDBTestDataMerkleTreeStruct{}
	}

	return currentMerkleTreeNodesForClient
}

// Save current TestData-MerkleTree for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeNodesForClient(testDataClientGuid string, merkleTreeNodes []cloudDBTestDataMerkleTreeStruct) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.merkleTreeNodes = merkleTreeNodes

	return true
}

// Retrieve current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeNodesForServer(testDataClientGuid string) (currentMerkleTreeNodesForServer []cloudDBTestDataMerkleTreeStruct) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true && len(tempdbData.serverData.merkleTreeNodes) > 0 {
		currentMerkleTreeNodesForServer = tempdbData.serverData.merkleTreeNodes
	} else {
		currentMerkleTreeNodesForServer = []cloudDBTestDataMerkleTreeStruct{}

		// Load Client's TestDataMerkleFilterPathHashes from CloudDB
		var tempCurrentMerkleTreeNodesForServer []cloudDBTestDataMerkleTreeStruct
		err := fenixTestDataSyncServerObject.loadAllTestDataMerkleTreesForClientFromCloudDB(testDataClientGuid, &tempCurrentMerkleTreeNodesForServer)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "5d049aea-28ff-4cea-9387-cc76efe2031e",
				"error": err,
			}).Error("Problem when executing: 'loadAllTestDataMerkleTreesForClientFromCloudDB()'")

			fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
			tempCurrentMerkleTreeNodesForServer = nil

		} else {

			// Save MerkleTreeNodes in memDB
			_ = fenixTestDataSyncServerObject.saveCurrentMerkleTreeNodesForServer(testDataClientGuid, tempCurrentMerkleTreeNodesForServer)

			currentMerkleTreeNodesForServer = tempCurrentMerkleTreeNodesForServer

		}
	}

	return currentMerkleTreeNodesForServer

}

// Retrieve the LeafNodes from the MerkleTree
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentMerkleTreeLeafNodeHashesForServer(testDataClientGuid string) (merkleTreeLeafNodeHashes []string) {

	var currentMerkleTreeNodesForServer []cloudDBTestDataMerkleTreeStruct
	var hightNodeLevel = -1

	// Get the Server MerkleTree
	currentMerkleTreeNodesForServer = fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForServer(testDataClientGuid)

	// If there are no MerkleTreeNodes then return empty slice
	if len(currentMerkleTreeNodesForServer) == 0 {
		return []string{}
	}

	// Get Highest NodelLevel, which are the LeafeNodes
	for _, merkleTreeNodeItem := range currentMerkleTreeNodesForServer {
		if merkleTreeNodeItem.nodeLevel > hightNodeLevel {
			hightNodeLevel = merkleTreeNodeItem.nodeLevel
		}
	}

	for _, merkleTreeNodeItem := range currentMerkleTreeNodesForServer {

		// Only process LeafNodes
		if merkleTreeNodeItem.nodeLevel == hightNodeLevel {
			merkleTreeLeafNodeHashes = append(merkleTreeLeafNodeHashes, merkleTreeNodeItem.nodeChildHash)
		}
	}

	return merkleTreeLeafNodeHashes

}

// Save current TestData-MerkleTree for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentMerkleTreeNodesForServer(testDataClientGuid string, merkleTreeNodes []cloudDBTestDataMerkleTreeStruct) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.merkleTreeNodes = merkleTreeNodes

	return true
}

// Retrieve current TestData-HeaderHash for client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashForClient(testDataClientGuid string) (currentHeaderHashsForClient string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeaderHashsForClient = tempdbData.clientData.testDataHeaderItemsHashes.headerItemsHash
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
	clientData.testDataHeaderItemsHashes.headerItemsHash = currentHeaderHashsForClient

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderHashForServer(testDataClientGuid string) (currentHeaderHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeaderHashForServer = tempdbData.serverData.testDataHeaderItemsHashes.headerItemsHash
	} else {
		currentHeaderHashForServer = "#VALUE IS MISSING#"
	}

	return currentHeaderHashForServer
}

// Retrieve current TestData-HeaderLabelHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentHeaderLabelHashForServer(testDataClientGuid string) (currentHeaderLabelHashForServer string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		currentHeaderLabelHashForServer = tempdbData.serverData.testDataHeaderItemsHashes.headerLabelsHash
	} else {
		currentHeaderLabelHashForServer = "#VALUE IS MISSING#"
	}

	return currentHeaderLabelHashForServer
}

// Save current TestData-HeaderHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderHashForServer(testDataClientGuid string, currentHeaderHashForServer string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.testDataHeaderItemsHashes.headerItemsHash = currentHeaderHashForServer

	return true
}

// Save current TestData-HeaderLabelHash for Server
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderLabelHashForServer(testDataClientGuid string, currentHeaderLabelHashForServer string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Save the data
	serverData.testDataHeaderItemsHashes.headerLabelsHash = currentHeaderLabelHashForServer

	return true
}

// Save current TestData-HeaderLabelHash for Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentHeaderLabelHashForClient(testDataClientGuid string, currentHeaderLabelHashForClient string) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to server data
	clientData := tempdbData.clientData

	// Save the data
	clientData.testDataHeaderItemsHashes.headerLabelsHash = currentHeaderLabelHashForClient

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
	tempdbData.serverData.testDataHeaderItemsHashes = tempdbData.clientData.testDataHeaderItemsHashes
	tempdbData.serverData.headers = tempdbData.clientData.headers
	tempdbData.serverData.testDataHeadersFilterValues = tempdbData.clientData.testDataHeadersFilterValues
	tempdbData.serverData.testDataHeaderItems = tempdbData.clientData.testDataHeaderItems

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
	tempdbData.serverData.MerkleFilterPath = tempdbData.clientData.MerkleFilterPath
	tempdbData.serverData.MerkleFilterPathHash = tempdbData.clientData.MerkleFilterPathHash
	tempdbData.serverData.merkleTreeNodes = tempdbData.clientData.merkleTreeNodes
	tempdbData.serverData.testDataRowItems = tempdbData.clientData.testDataRowItems

	// Save TestData in CloudDB
	err := fenixTestDataSyncServerObject.saveMerkleHashMerkleTreeAndTestDataRowsToCloudDB(testDataClientGuid)

	if err != nil {
		return false
	}
	return true
}

// Get current Server TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowItemsForServer(testDataClientGuid string) (testDataRowItems []cloudDBTestDataRowItemCurrentStruct) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true && len(tempdbData.serverData.testDataRowItems) > 0 {
		testDataRowItems = tempdbData.serverData.testDataRowItems
	} else {

		// Load Client's TestDataMerkleHashes from CloudDB
		var tempMemDBTestDataRowItems []cloudDBTestDataRowItemCurrentStruct
		err := fenixTestDataSyncServerObject.loadAllTestDataRowItemsForClientFromCloudDB(testDataClientGuid, &tempMemDBTestDataRowItems)
		if err != nil {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id":    "bac0347f-c066-4a0f-9132-2ae1175a795c",
				"error": err,
			}).Error("Problem when executing: 'loadAllTestDataRowItemsForClientFromCloudDB()'")

			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"Id": "8b84863d-f18f-4db6-9d91-904971fec6b3",
			}).Info("Switching: 'stateProcessIncomingAndOutgoingMessage = false'")

			fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
			testDataRowItems = []cloudDBTestDataRowItemCurrentStruct{}

		} else {
			// Save RowsItems in memDB
			tempdbData.serverData.testDataRowItems = tempMemDBTestDataRowItems

			// Prepare for returning of RowsItems found in CloudDB
			testDataRowItems = tempMemDBTestDataRowItems
		}
	}

	return testDataRowItems

}

// Get current ChildNodeHashes to be removed from TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentChildNodeHashesToBeRemovedForServer(testDataClientGuid string) (leafNodeHashesThatNoLongerExist []string) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true { //&& { len(tempdbData.serverData.merkleTreeNodesChildHashesThatNoLongerExist) > 0 {
		leafNodeHashesThatNoLongerExist = tempdbData.serverData.merkleTreeNodesChildHashesThatNoLongerExist
	}
	/*else {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "3aee5180-1dad-418d-afc5-140d408477d3",
		}).Fatalln("'merkleTreeNodesChildHashesThatNoLongerExist' shouldn't be empty")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
	}


	*/
	return leafNodeHashesThatNoLongerExist

}

// Get current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getCurrentTestDataRowItemsForClient(testDataClientGuid string) (testDataRowItems []cloudDBTestDataRowItemCurrentStruct) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true && len(tempdbData.clientData.testDataRowItems) > 0 {
		testDataRowItems = tempdbData.clientData.testDataRowItems
	} else {
		testDataRowItems = fenixTestDataSyncServerObject.getCurrentTestDataRowItemsForServer(testDataClientGuid)
		//testDataRowItems = []cloudDBTestDataRowItemCurrentStruct{}
	}

	return testDataRowItems
}

// Get the Domain for the Client
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getDomainUuidForClient(testDataClientGuid string) (domainUuid memDBDomainUuidType) {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := cloudDBClientsMap[memDBClientUuidType(testDataClientGuid)]

	// Get the data
	if valueExits == true {
		domainUuid = memDBDomainUuidType(tempdbData.domainUuid)
	} else {
		domainUuid = "666"

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "5617e51e-36ce-4c71-a10d-836c2eb604ee",
		}).Fatalln("Should not happen. Existing Client in memoryDB for Domain-Client is missing")

	}

	return domainUuid
}

// Save current Client TestDataRows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveCurrentTestDataRowItemsForClient(testDataClientGuid string, testDataRowItems []cloudDBTestDataRowItemCurrentStruct) bool {

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	// Save the data
	clientData.testDataRowItems = testDataRowItems

	return true
}

/*
Using 'isClientKnownToServer' instead


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

*/

// Remove the rows that don't is represented in the Clients MerkleTree
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) removeMerkleTreeNodeItemsInMemoryDBThatIsNotRepresentedInClientsNewMerkleTree(callingClientUuid string) bool {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "2712573d-354a-4fb9-a510-39b3844fe95d",
	}).Debug("Entering: 'removeMerkleTreeNodeItemsInMemoryDBThatIsNotRepresentedInClientsNewMerkleTree'")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "0fcc0362-263a-4945-a400-6e22baf53240",
		}).Debug("Exiting: 'removeMerkleTreeNodeItemsInMemoryDBThatIsNotRepresentedInClientsNewMerkleTree'")
	}()

	// Get pointer to data for Client_UUID
	tempdbData := dbDataMap[memDBClientUuidType(callingClientUuid)]

	// Get pointer to client data
	clientData := tempdbData.clientData

	//highestLeafNodeLevel := fenixTestDataSyncServerObject.getMerkleLeafNodeLevel(clientData.merkleTreeNodes)

	// Extract the hashes for all leaf nodes from Client MerkleTree
	var clientsNewLeafNodesHashes []string
	for _, MerkleTreeNode := range clientData.merkleTreeNodes {
		clientsNewLeafNodesHashes = append(clientsNewLeafNodesHashes, MerkleTreeNode.nodeHash)
	}

	// Get pointer to server data
	serverData := tempdbData.serverData

	// Get the TestData to process
	serverMerkleTree := serverData.merkleTreeNodes

	// Filter out what should be kept and what should be removed
	var serverTestDataToKeep []cloudDBTestDataMerkleTreeStruct
	var leafNodeHashesToRemoveFromServer []string

	// Get NodeLevel for LeafNodes
	var highestNodeLevel = -1
	for _, serverMerkleTreeNodeItem := range serverMerkleTree {
		if serverMerkleTreeNodeItem.nodeLevel > highestNodeLevel {
			highestNodeLevel = serverMerkleTreeNodeItem.nodeLevel
		}
	}

	// Loop over existing ServerMerkleTree-data and process each item
	for _, serverMerkleTreeNodeItem := range serverMerkleTree {

		// Only process LeafNodes
		if serverMerkleTreeNodeItem.nodeLevel == highestNodeLevel {
			if existsValueInStringArray(serverMerkleTreeNodeItem.nodeHash, clientsNewLeafNodesHashes) == true {
				// MerkleTreeRow that should be kept
				serverTestDataToKeep = append(serverTestDataToKeep, serverMerkleTreeNodeItem)

			} else {
				// LeafNodeHashes to be removed
				leafNodeHashesToRemoveFromServer = append(leafNodeHashesToRemoveFromServer, serverMerkleTreeNodeItem.nodeHash)

			}
		}
	}

	// Save MerkleTreeRows to be kept in memoryDB
	serverData.merkleTreeNodes = serverTestDataToKeep

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                   "bfdbc46f-963a-427e-96bc-f949f76c6d62",
		"serverTestDataToKeep": serverTestDataToKeep,
	}).Debug("MerkleTreeRows to be kept in memoryDB for Client: " + callingClientUuid)

	// Save the leafNodeHashes to be removed
	serverData.merkleTreeNodesChildHashesThatNoLongerExist = leafNodeHashesToRemoveFromServer

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "6e8f00c5-881c-471c-bc02-984c73326c06",
		"serverData.merkleTreeNodesChildHashesThatNoLongerExist": serverData.merkleTreeNodesChildHashesThatNoLongerExist,
	}).Debug("LeafNodeHashes to be removed for Client: " + callingClientUuid)

	return true
}

// Clear current Server MerkleHash, MerkleTree and TestData
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentMerkleDataAndTestDataRowsForServer(testDataClientGuid string) bool {

	// Clear current Server MerkleHash
	_ = fenixTestDataSyncServerObject.clearCurrentMerkleHashForServer(testDataClientGuid)

	// Clear current Server MerkleTree
	_ = fenixTestDataSyncServerObject.clearCurrentMerkleTreeForServer(testDataClientGuid)

	// Clear current Server MerkleFilterPath
	_ = fenixTestDataSyncServerObject.clearCurrentMerkleFilterPathForServer(testDataClientGuid)

	// Clear current Server MerkleFilterPathHash
	_ = fenixTestDataSyncServerObject.clearCurrentMerkleFilterPathHashForServer(testDataClientGuid)

	// Clear current Server TestDataRowItems
	_ = fenixTestDataSyncServerObject.clearCurrentTestDataRowItemsForServer(testDataClientGuid)

	return true

}

// Clear current Server TestDataHeaderItemsHashes, TestDataHeaderItems, TestDataHeadersFilterValues andTestDataHeaderNames
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentHeaderDataForServer(testDataClientGuid string) bool {

	// Clear current Server TestDataHeaderItemsHashes
	_ = fenixTestDataSyncServerObject.clearCurrentTestDataHeaderItemsHashesForServer(testDataClientGuid)

	// Clear current Server TestDataHeaderItems
	_ = fenixTestDataSyncServerObject.clearCurrentTestDataHeaderItemsForServer(testDataClientGuid)

	// Clear current Server MerkleFilterPath
	_ = fenixTestDataSyncServerObject.clearCurrentTestDataHeadersFilterValuesForServer(testDataClientGuid)

	// Clear current Server MerkleFilterPathHash
	_ = fenixTestDataSyncServerObject.clearCurrentTestDataHeaderNamesForServer(testDataClientGuid)

	return true

}

// Clear current Server TestDataHeaders, slice with the Header Names
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentTestDataHeaderNamesForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.headers = []string{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "a928cb53-71b8-47c8-9dfe-e7f6119d5d95",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server HeaderItemsHashes
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentTestDataHeaderItemsHashesForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.testDataHeaderItemsHashes = cloudDBTestDataHeaderItemsHashesStruct{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "72cbd468-abe0-4b3d-b3cb-54d203a72e6a",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server TestDataHeaderItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentTestDataHeaderItemsForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.testDataHeaderItems = []cloudDBTestDataHeaderItemStruct{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "07df1323-6727-4093-b223-83dda9594804",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server TestDataHeadersFilterValues
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentTestDataHeadersFilterValuesForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.testDataHeadersFilterValues = []cloudDBTestDataHeaderFilterValuesStruct{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "b7f024fb-dc11-4c5b-a0a3-f3b3fda2f45c",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server MerkleHash
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentMerkleHashForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.merkleHash = ""

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "f6b19b42-fdfb-4adb-ab3f-735cc51c28c4",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server MerkleTree
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentMerkleTreeForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.merkleTreeNodes = []cloudDBTestDataMerkleTreeStruct{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "a350f7bf-b537-4494-aeec-739634b200ca",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server MerkleFilterPath
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentMerkleFilterPathForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.MerkleFilterPath = ""

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "18045c11-2f5a-4ccd-ae5d-c837aab859a8",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server MerkleFilterPathHash
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentMerkleFilterPathHashForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.MerkleFilterPathHash = ""

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "07c70742-a965-4890-824a-f28576fea5d9",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}

// Clear current Server TestDataRowItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) clearCurrentTestDataRowItemsForServer(testDataClientGuid string) bool {

	// Get pointer to data for Client_UUID
	tempdbData, valueExits := dbDataMap[memDBClientUuidType(testDataClientGuid)]

	// Validate that reference exists
	if valueExits == true {
		// Clear values
		tempdbData.serverData.testDataRowItems = []cloudDBTestDataRowItemCurrentStruct{}

	} else {
		// This should not happen
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "4c23449d-53b0-4f38-a31e-555da1487252",
		}).Fatalln("Reference to Client in memoryDB should exist for Client: ", testDataClientGuid)
	}

	return true
}
