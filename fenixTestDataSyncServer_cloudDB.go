package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/sirupsen/logrus"
)

// All TestTDataClients in CloudDB
var memDBAllClients []memDBAllTestDataClientStruct

type memDBAllTestDataClientStruct struct {
	clientUid   memDBClientUuidType
	clientName  string
	domainUuid  memDBDomainUuidType
	description string
}

// All TestDataHeaderFilterValues in CloudDB
var memDBAllTestDataHeaderFilterValues []memDBAllTestDataHeaderFilterValueStruct

type memDBAllTestDataHeaderFilterValueStruct struct {
	headerItemHash    string
	headerFilterValue string
	clientUuid        memDBClientUuidType
	domainUuid        memDBDomainUuidType
}

// All TestDataHeaderItems in CloudDB
var memDBAllTestDataHeaderItems []memDBAllTestDataHeaderItemStruct

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

// All TestDataMerkleHashes in CloudDB
var memDBAllTestDataMerkleHashes []memDBAllTestDataMerkleHashStruct

type memDBAllTestDataMerkleHashStruct struct {
	clientUuid memDBClientUuidType
	domainUuid memDBDomainUuidType
	merkleHash string
	merklePath string
}

// All TestDataMerkleTrees in CloudDB
var memDBAllTestDataMerkleTrees []memDBAllTestDataMerkleTreeStruct

type memDBAllTestDataMerkleTreeStruct struct {
	clientUuid    memDBClientUuidType
	domainUuid    memDBDomainUuidType
	nodeLevel     int
	nodeName      string
	nodePath      string
	nodeHash      string
	nodeChildHash string
}

// All TestDataRowItems in CloudDB
var memDBAllTestDataRowItems []memDBAllTestDataRowItemStruct

type memDBAllTestDataRowItemStruct struct {
	clientUuid            memDBClientUuidType
	domainUuid            memDBDomainUuidType
	rowHash               string
	testdataValueAsString string
	leafNodeName          string
	leafNodePath          string
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

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestTDataClients in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllClientsFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "16af90a4-aa07-4d8b-921a-a47c04811a9b",
	}).Debug("Entering: loadAllmemDBAllClientsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "e9659490-9ba7-437b-a235-88d8369ebf36",
		}).Debug("Exiting: loadAllmemDBAllClientsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\", clients.\"client_name\" , clients.\"domain_uuid\", clients.\"client_description\"  "
	sqlToExecute = sqlToExecute + "FROM clients "
	sqlToExecute = sqlToExecute + "WHERE clients.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.replaced_by_new_version = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.client_areatyp_id = 1 " // Clients used for 'TestData'

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataClient memDBAllTestDataClientStruct
	var testDataClients []memDBAllTestDataClientStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataClient.clientUid, &testDataClient.clientName, &testDataClient.domainUuid, &testDataClient.description)
		if err != nil {
			return err
		}
		testDataClients = append(testDataClients, testDataClient)

	}

	// Copy extracted data into memory object
	memDBAllClients = testDataClients

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataHeaderFilterValues in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "097f9f31-f29b-4c4a-aadb-0d4120429cf5",
	}).Debug("Entering: loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "cfb2f17f-579d-4d98-b91b-ba3c01a32771",
		}).Debug("Exiting: loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.header.filtervalues.\"header_item_hash\", testdata.header.filtervalues.\"header_filter_value\" , testdata.header.filtervalues.\"client_uuid\", testdata.header.filtervalues.\"domain_uuid\"  "
	sqlToExecute = sqlToExecute + "FROM testdata.header.filtervalues "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataHeaderFilterValue memDBAllTestDataHeaderFilterValueStruct
	var testDataHeaderFilterValues []memDBAllTestDataHeaderFilterValueStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderFilterValue.headerItemHash, &testDataHeaderFilterValue.headerFilterValue, &testDataHeaderFilterValue.clientUuid, &testDataHeaderFilterValue.domainUuid)
		if err != nil {
			return err
		}
		testDataHeaderFilterValues = append(testDataHeaderFilterValues, testDataHeaderFilterValue)

	}

	// Copy extracted data into memory object
	memDBAllTestDataHeaderFilterValues = testDataHeaderFilterValues

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataHeaderItems in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataHeaderItemsFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "a2951e5e-d7d0-4240-88a9-5dc570f2bbe9",
	}).Debug("Entering: loadAllmemDBAllTestDataHeaderItemsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "70e1c4da-5001-4199-adf4-bbc2576ccdab",
		}).Debug("Exiting: loadAllmemDBAllTestDataHeaderItemsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.header.items.\"client_uuid\", testdata.header.items.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"header_item_hash\", testdata.header.items.\"header_label\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"should_be_used_in_filter\", testdata.header.items.\"is_mandatory_in_filter\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"filter_selection_type\", testdata.header.items.\"filter_values_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata.header.items "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataHeaderItem memDBAllTestDataHeaderItemStruct
	var testDataHeaderItems []memDBAllTestDataHeaderItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderItem.clientUuid, &testDataHeaderItem.domainUuid, &testDataHeaderItem.headerItemHash,
			&testDataHeaderItem.headerLabel, &testDataHeaderItem.shouldBeUsedInFilter, &testDataHeaderItem.isMandatoryInFilter,
			&testDataHeaderItem.filterSelection_type, &testDataHeaderItem.filterValuesHash)
		if err != nil {
			return err
		}
		testDataHeaderItems = append(testDataHeaderItems, testDataHeaderItem)

	}

	// Copy extracted data into memory object
	memDBAllTestDataHeaderItems = testDataHeaderItems

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataMerkleHashes in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataMerkleHashesFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "53461f88-c773-477e-b459-cfb93a1c3eaa",
	}).Debug("Entering: loadAllmemDBAllTestDataMerkleHashesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "851cbecf-4084-4e38-b922-eab2a4b526d1",
		}).Debug("Exiting: loadAllmemDBAllTestDataMerkleHashesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.merklehashes.\"client_uuid\", testdata.merklehashes.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.merklehashes.\"merklehash\", testdata.merklehashes.\"merkle_path\" "
	sqlToExecute = sqlToExecute + "FROM testdata.merklehashes "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataMerkleHash memDBAllTestDataMerkleHashStruct
	var testDataMerkleHashs []memDBAllTestDataMerkleHashStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataMerkleHash.clientUuid, &testDataMerkleHash.domainUuid,
			&testDataMerkleHash.merkleHash, &testDataMerkleHash.merklePath)
		if err != nil {
			return err
		}
		testDataMerkleHashs = append(testDataMerkleHashs, testDataMerkleHash)

	}

	// Copy extracted data into memory object
	memDBAllTestDataMerkleHashes = testDataMerkleHashs

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataMerkleTrees in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataMerkleTreesFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "6e0a6f4c-cc54-4aff-94f1-243aee6141ae",
	}).Debug("Entering: loadAllmemDBAllTestDataMerkleTreesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "a6d6b40f-c30f-4ec1-b5d2-f94b0234471c",
		}).Debug("Exiting: loadAllmemDBAllTestDataMerkleTreesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.merkletrees.\"client_uuid\", testdata.merkletrees.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_level\", testdata.merkletrees.\"node_name\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_path\", testdata.merkletrees.\"node_hash\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_child_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata.merkletrees "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataMerkleTree memDBAllTestDataMerkleTreeStruct
	var testDataMerkleTrees []memDBAllTestDataMerkleTreeStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataMerkleTree.clientUuid, &testDataMerkleTree.domainUuid,
			&testDataMerkleTree.nodeLevel, &testDataMerkleTree.nodeName,
			&testDataMerkleTree.nodePath, &testDataMerkleTree.nodeHash,
			&testDataMerkleTree.nodeChildHash)
		if err != nil {
			return err
		}
		testDataMerkleTrees = append(testDataMerkleTrees, testDataMerkleTree)

	}

	// Copy extracted data into memory object
	memDBAllTestDataMerkleTrees = testDataMerkleTrees

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataRowItems in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataRowItemsFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "61b8b021-9568-463e-b867-ac1ddb10584d",
	}).Debug("Entering: loadAllmemDBAllTestDataRowItemsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "78a97c41-a098-4122-88d2-01ed4b6c4844",
		}).Debug("Exiting: loadAllmemDBAllTestDataRowItemsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.row.items.current.\"client_uuid\", testdata.row.items.current.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.row.items.current.\"row_hash\", testdata.row.items.current.\"testdata_value_as_string\", "
	sqlToExecute = sqlToExecute + "testdata.row.items.current.\"leaf_node_name\", testdata.row.items.current.\"leaf_node_path\" "
	sqlToExecute = sqlToExecute + "FROM testdata.row.items.current "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataRowItem memDBAllTestDataRowItemStruct
	var testDataRowItems []memDBAllTestDataRowItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataRowItem.clientUuid, &testDataRowItem.domainUuid,
			&testDataRowItem.rowHash, &testDataRowItem.testdataValueAsString,
			&testDataRowItem.leafNodeName, &testDataRowItem.leafNodePath)
		if err != nil {
			return err
		}
		testDataRowItems = append(testDataRowItems, testDataRowItem)

	}

	// Copy extracted data into memory object
	memDBAllTestDataRowItems = testDataRowItems

	// No errors occurred
	return nil

}
