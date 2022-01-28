package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/sirupsen/logrus"
)

// All TestTDataClients in CloudDB
var memDBAllClients memDBAllTestDataClientsStruct

type memDBAllTestDataClientStruct struct {
	clientUid   memDBClientUuidType
	clientName  string
	domainUuid  memDBDomainUuidType
	description string
}
type memDBAllTestDataClientsStruct struct {
	memDBAllClients []memDBAllTestDataClientStruct
}

// All TestDataHeaderFilterValues in CloudDB
var memDBAllTestDataHeaderFilterValues memDBAllTestDataHeaderFilterValuesStruct

type memDBAllTestDataHeaderFilterValueStruct struct {
	headerItemHash    string
	headerFilterValue string
	clientUuid        memDBClientUuidType
	domainUuid        memDBDomainUuidType
}
type memDBAllTestDataHeaderFilterValuesStruct struct {
	memDBAllTestDataHeaderFilterValues []memDBAllTestDataHeaderFilterValueStruct
}

// All TestDataHeaderItems in CloudDB
var memDBAllTestDataHeaderItems memDBAllTestDataHeaderItemsStruct

type memDBAllTestDataHeaderItemStruct struct {
	clientUuid           memDBClientUuidType
	domainUuid           memDBDomainUuidType
	headerItemHash       string
	headerLabel          string
	shouldBeUsedInFilter bool
	isMandatoryInFilter  bool
	filterSelection_type int
	filterValuesHash     string
}
type memDBAllTestDataHeaderItemsStruct struct {
	memDBAllTestDataHeaderItems []memDBAllTestDataHeaderItemStruct
}

// All TestDataMerkleHashes in CloudDB
var memDBAllTestDataMerkleHashes memDBAllTestDataMerkleHashesStruct

type memDBAllTestDataMerkleHashStruct struct {
	clientUuid memDBClientUuidType
	domainUuid memDBDomainUuidType
	merkleHash string
	merklePath string
}
type memDBAllTestDataMerkleHashesStruct struct {
	memDBAllTestDataMerkleHash []memDBAllTestDataMerkleHashStruct
}

// All TestDataMerkleTrees in CloudDB
var memDBAllTestDataMerkleTrees memDBAllTestDataMerkleTreesStruct

type memDBAllTestDataMerkleTreeStruct struct {
	clientUuid    memDBClientUuidType
	domainUuid    memDBDomainUuidType
	nodeLevel     int
	nodeName      string
	nodePath      string
	nodeHash      string
	nodeChildHash string
}
type memDBAllTestDataMerkleTreesStruct struct {
	memDBAllTestDataMerkleTrees []memDBAllTestDataMerkleTreeStruct
}

// All TestDataRowItems in CloudDB
var memDBAllTestDataRowItems memDBAllTestDataRowItemsStruct

type memDBAllTestDataRowItemStruct struct {
	clientUuid            memDBClientUuidType
	domainUuid            memDBDomainUuidType
	rowHash               string
	testdataValueAsString string
	leafNodeName          string
	leafNodePath          string
}
type memDBAllTestDataRowItemsStruct struct {
	memDBAllTestDataRowItems []memDBAllTestDataRowItemStruct
}

//
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) existsClientInDB(testDataClientUUID string) (bool, error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "73768ffb-e069-42cf-945f-009a332a9ac6",
	}).Debug("Entering: ListTestInstructionsInDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "07e53e86-e45f-4799-a4cf-4d5c73872fbe",
		}).Debug("Exiting: ListTestInstructionsInDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\" "
	sqlToExecute = sqlToExecute + "FROM clients "
	sqlToExecute = sqlToExecute + "WHERE clients.\"client_uuid\" = '" + testDataClientUUID + "' "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.replaced_by_new_version = false "

	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	var clientUuid string
	var clientUuidRows []string

	for rows.Next() {
		err := rows.Scan(&clientUuid)
		if err != nil {
			return false, err
		}
		clientUuidRows = append(clientUuidRows, clientUuid)

	}

	//Only OK result is ONE row
	numberOfInstancesFound := len(clientUuidRows)
	if numberOfInstancesFound != 1 {

		myErr := errors.New("Found " + string(numberOfInstancesFound) + " but expected 1 instance of Client UUID, for: " + clientUuid)
		return false, myErr

	} else {

		return true, nil

	}

}

// Load all CloudDB-data that should be held in MemoryDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadCloudDBIntoMemoryDB() error {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "42648bb9-e953-4e13-9855-51d7401fa291",
	}).Debug("Entering: loadCloudDBIntoMemoryDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "87694618-5cba-411b-8d8a-cdec9c16f78e",
		}).Debug("Exiting: loadCloudDBIntoMemoryDB()")
	}()

	// Temporary memObject object
	var tempMemoryDB memoryDBStruct

	// Load Clients into MemoryDB
	allowedClients, err := fenixTestDataSyncServerObject.loadClientsFromCloudDB()
	if err != nil {
		return err
	}

	tempMemoryDB.allowedClients.memDBTestDataDomainType = allowedClients.memDBTestDataDomainType

	var (
		headerItems []memDBHeaderItemsStruct

		headerFilterValues []string
	)

	// Load Servers TestData content, that was previously sent by the clients

	// Loop over all Client&Domain combinations and retrieve Data from CloudDB
	for allowedDomain, allowedClient := range allowedClients.memDBTestDataDomainType {
		fmt.Println(allowedDomain, allowedClient)

		// Add all testdata rows
		testDataRows, err := fenixTestDataSyncServerObject.mapTestDataRowsFromCloudDBToMemDB(allowedClient, allowedDomain)
		if err != nil {
			return err
		}

		// Add all HeaderItems
		var headerFilterValue string

		headerFilterValues = []string{headerFilterValue}

		headerItem := memDBHeaderItemsStruct{
			headerItemHash:             "",
			headerLabel:                "",
			headerShouldBeUsedInFilter: false,
			headerIsMandatoryInFilter:  false,
			headerFilterSelectionType:  0,
			headerFilterValuesItem: memDBHeaderFilterValuesItemStruct{
				HeaderFilterValuesHash: "",
				HeaderFilterValues:     headerFilterValues,
			},
		}

		headerItems = []memDBHeaderItemsStruct{headerItem}

		// Add all MerkleTreeRows
		merkleTreeRow := memDBMerkleTreeRowStruct{
			nodeLevel:     "",
			nodeName:      "",
			nodePath:      "",
			nodeHash:      "",
			nodeChildHash: "",
		}

		merkleTreeRows := []memDBMerkleTreeRowStruct{merkleTreeRow}

		// Build TestData object
		memDBTestDataStruct := memDBTestDataStruct{memDBDataStructureStruct{
			merkleHash: "",
			merklePath: "",
			merkleTree: memDBMerkleTreeRowsStruct{
				merkleTreeRows: merkleTreeRows,
			},
			headerItemsHash:  "",
			headerLabelsHash: "",
			headerItems:      headerItems,
			testDataRows:     testDataRows,
			testDataAsDataFrame: dataframe.DataFrame{
				Err: nil,
			},
		},
		}

		// Add TestData object to correct Domain and Client
		tempMemoryDB.server.memDBTestDataDomain["domainID"]["ClientID"] = memDBTestDataStruct

	}

	/*

		memoryDB.allowedClients = tempMemoryDB.allowedClients
		memoryDB.


	*/
	return nil
}

// Load all Clients and their Domains from CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadClientsFromCloudDB() (allowedClients memDBAllowedClientsStruct, err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "6a407f6f-6a44-48e6-b50a-22a3775b55fe",
	}).Debug("Entering: loadClientsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "6d20a55b-cb15-4d88-be7a-dee5a9b24630",
		}).Debug("Exiting: loadClientsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\", clients.\"domain_uuid\"  "
	sqlToExecute = sqlToExecute + "FROM clients "
	sqlToExecute = sqlToExecute + "WHERE clients.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.replaced_by_new_version = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.client_areatyp_id = 1 "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var clientUuid memDBClientUuidType
	var domainUuid memDBDomainUuidType
	clientsWithDomains := make(map[memDBClientUuidType]memDBDomainUuidType)

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&clientUuid, &domainUuid)
		if err != nil {
			return memDBAllowedClientsStruct{}, err
		}
		clientsWithDomains[clientUuid] = domainUuid

	}

	allowedClients.memDBTestDataDomainType = clientsWithDomains

	return allowedClients, nil

}

// Load all Testa sent by all Clients from CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataForEachClientFromCloudDB() (allowedClients memDBAllowedClientsStruct, err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "388717a0-87f4-4d91-b00c-01ef922f89f3",
	}).Debug("Entering: loadAllTestDataForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "017921fd-fa8e-4c0f-8193-ee83121f32f6",
		}).Debug("Exiting: loadAllTestDataForEachClientFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\", clients.\"domain_uuid\"  "
	sqlToExecute = sqlToExecute + "FROM clients "
	sqlToExecute = sqlToExecute + "WHERE clients.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.replaced_by_new_version = false "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var clientUuid memDBClientUuidType
	var domainUuid memDBDomainUuidType
	clientsWithDomains := make(map[memDBClientUuidType]memDBDomainUuidType)

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&clientUuid, &domainUuid)
		if err != nil {
			return memDBAllowedClientsStruct{}, err
		}
		clientsWithDomains[clientUuid] = domainUuid

	}

	allowedClients.memDBTestDataDomainType = clientsWithDomains

	return allowedClients, nil

}

// Load all MerkleHash data, from CloudDB, previously sent by all Clients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllMerkleHashesForEachClientFromCloudDB() (allowedClients memDBAllowedClientsStruct, err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "b709e284-5414-4695-b4f9-79d33593ec47",
	}).Debug("Entering: loadAllMerkleHashesForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "89ac4da2-52b3-4108-ac77-83eb535b2003",
		}).Debug("Exiting: loadAllMerkleHashesForEachClientFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\", clients.\"domain_uuid\", clients.\"merklehash\", clients.\"merkle_path\"  "
	sqlToExecute = sqlToExecute + "FROM testdata.merklehash "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var clientUuid memDBClientUuidType
	var domainUuid memDBDomainUuidType
	var merkleHash string
	var merklePath string

	clientsWithDomains := make(map[memDBClientUuidType]memDBDomainUuidType)

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&clientUuid, &domainUuid, &merkleHash, &merklePath)
		if err != nil {
			return memDBAllowedClientsStruct{}, err
		}
		clientsWithDomains[clientUuid] = domainUuid

	}

	allowedClients.memDBTestDataDomainType = clientsWithDomains

	return allowedClients, nil

}

// Load all MerkleTree data, from CloudDB, previously sent by all Clients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllMerkleTreesForEachClientFromCloudDB() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "3333e654-38c3-4ec2-82c6-8b5dd3bdaa38",
	}).Debug("Entering: loadAllMerkleTreesForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "ba576560-4878-44d6-ac7a-3e39aa6b3fd8",
		}).Debug("Exiting: loadAllMerkleTreesForEachClientFromCloudDB()")
	}()
}

// Load all TestDataRows, from CloudDB, previously sent by all Clients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataRowForEachClientFromCloudDB() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "45e80b7a-622a-437a-9b63-6857191cd3cc",
	}).Debug("Entering: loadAllTestDataRowForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "2dcb288f-91b6-4849-91cf-6a22cfdca0eb",
		}).Debug("Exiting: loadAllTestDataRowForEachClientFromCloudDB()")
	}()
}

// Load all TestDataHeaderItems, from CloudDB, previously sent by all Clients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataHeaderItemsForEachClientFromCloudDB() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "11b27f30-097f-412d-aa43-90e9873327a7",
	}).Debug("Entering: loadAllTestDataHeaderItemsForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "70e84116-a16c-4fb3-a4c0-14abd7c85562",
		}).Debug("Exiting: loadAllTestDataHeaderItemsForEachClientFromCloudDB()")
	}()
}

// Load all TestDataHeaderItemFilterValues, from CloudDB, previously sent by all Clients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataHeaderItemFilterValuesForEachClientFromCloudDB() {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "7655381e-6cd9-4819-81ce-1bb4b1fd5674",
	}).Debug("Entering: loadAllTestDataHeaderItemFilterValuesForEachClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "e6ad1e09-d7ec-4dd6-9dbd-cbf68bab4e00",
		}).Debug("Exiting: loadAllTestDataHeaderItemFilterValuesForEachClientFromCloudDB()")
	}()
}

// Add all testdata rows
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) mapTestDataRowsFromCloudDBToMemDB(allowedDomain memDBDomainUuidType, allowedClient memDBClientUuidType) (testDataRows []memDBTestDataItemsStruct, err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "2018c1fd-92eb-4413-a765-ccef99718410",
	}).Debug("Entering: mapTestDataRowsFromCloudDBToMemDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "575294aa-9446-4d8c-8a5e-c0d3393076b2",
		}).Debug("Exiting: mapTestDataRowsFromCloudDBToMemDB()")
	}()

	testDataRow := memDBTestDataItemsStruct{
		testDataRowHash:        "",
		leafNodeName:           "",
		leafNodePath:           "",
		testDataValuesAsString: nil,
	}

	testDataRows = []memDBTestDataItemsStruct{testDataRow}

	return testDataRows, nil
}
