package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestTDataClients in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllClientsFromCloudDB(testDataClients *[]memCloudDBAllTestDataClientStruct) (err error, memCloudDBAllClientsMap memCloudDBAllClientsMapType) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "16af90a4-aa07-4d8b-921a-a47c04811a9b",
	}).Debug("Entering: loadAllmemDBAllClientsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "e9659490-9ba7-437b-a235-88d8369ebf36",
		}).Debug("Exiting: loadAllmemDBAllClientsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT clients.\"client_uuid\", clients.\"client_name\" , clients.\"domain_uuid\", clients.\"description\"  "
	sqlToExecute = sqlToExecute + "FROM clients "
	sqlToExecute = sqlToExecute + "WHERE clients.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.replaced_by_new_version = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "clients.client_areatyp_id = 1 " // Clients used for 'TestData'

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "831e74ff-e697-4228-9598-8f1dfcf24c65",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err, nil
	}

	// Variables to used when extract data from result set
	var testDataClient memCloudDBAllTestDataClientStruct
	memCloudDBAllClientsMap = make(map[memDBClientUuidType]memCloudDBAllTestDataClientMapStruct)

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataClient.clientUuid, &testDataClient.clientName, &testDataClient.domainUuid, &testDataClient.description)
		if err != nil {
			return err, nil
		}
		*testDataClients = append(*testDataClients, testDataClient)
		memCloudDBAllClientsMap[testDataClient.clientUuid] = memCloudDBAllTestDataClientMapStruct{
			clientName:  testDataClient.clientName,
			domainUuid:  testDataClient.domainUuid,
			description: testDataClient.description,
		}

	}

	// No errors occurred
	return nil, memCloudDBAllClientsMap

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataHeaderFilterValues in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB(testDataHeaderFilterValues *[]memCloudDBAllTestDataHeaderFilterValueStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "097f9f31-f29b-4c4a-aadb-0d4120429cf5",
	}).Debug("Entering: loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "cfb2f17f-579d-4d98-b91b-ba3c01a32771",
		}).Debug("Exiting: loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata_header_filtervalues.\"header_item_hash\", testdata_header_filtervalues.\"header_filter_value\" , testdata_header_filtervalues.\"client_uuid\", testdata_header_filtervalues.\"domain_uuid\"  "
	sqlToExecute = sqlToExecute + "FROM testdata_header_filtervalues "

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "66813fb3-73d1-4e99-98e7-d4161b1869e1",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err
	}

	// Variables to used when extract data from result set
	var testDataHeaderFilterValue memCloudDBAllTestDataHeaderFilterValueStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderFilterValue.headerItemHash, &testDataHeaderFilterValue.headerFilterValue, &testDataHeaderFilterValue.clientUuid, &testDataHeaderFilterValue.domainUuid)
		if err != nil {
			return err
		}
		*testDataHeaderFilterValues = append(*testDataHeaderFilterValues, testDataHeaderFilterValue)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataHeaderItems in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataHeaderItemsFromCloudDB(testDataHeaderItems *[]memCloudDBAllTestDataHeaderItemStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "a2951e5e-d7d0-4240-88a9-5dc570f2bbe9",
	}).Debug("Entering: loadAllmemDBAllTestDataHeaderItemsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "70e1c4da-5001-4199-adf4-bbc2576ccdab",
		}).Debug("Exiting: loadAllmemDBAllTestDataHeaderItemsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata_header_items.\"client_uuid\", testdata_header_items.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata_header_items.\"header_item_hash\", testdata_header_items.\"header_label\", "
	sqlToExecute = sqlToExecute + "testdata_header_items.\"should_be_used_in_filter\", testdata_header_items.\"is_mandatory_in_filter\", "
	sqlToExecute = sqlToExecute + "testdata_header_items.\"filter_selection_type\", testdata_header_items.\"filter_values_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata_header_items "

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "b425c42f-ddfa-4474-9099-40d38c2a968d",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err
	}

	// Variables to used when extract data from result set
	var testDataHeaderItem memCloudDBAllTestDataHeaderItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderItem.clientUuid, &testDataHeaderItem.domainUuid, &testDataHeaderItem.headerItemHash,
			&testDataHeaderItem.headerLabel, &testDataHeaderItem.shouldBeUsedInFilter, &testDataHeaderItem.isMandatoryInFilter,
			&testDataHeaderItem.filterSelection_type, &testDataHeaderItem.filterValuesHash)
		if err != nil {
			return err
		}
		*testDataHeaderItems = append(*testDataHeaderItems, testDataHeaderItem)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataMerkleHashes in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataMerkleHashesFromCloudDB(testDataMerkleHashs *[]memCloudDBAllTestDataMerkleHashStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "53461f88-c773-477e-b459-cfb93a1c3eaa",
	}).Debug("Entering: loadAllmemDBAllTestDataMerkleHashesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "851cbecf-4084-4e38-b922-eab2a4b526d1",
		}).Debug("Exiting: loadAllmemDBAllTestDataMerkleHashesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata_merklehashes.\"client_uuid\", testdata_merklehashes.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata_merklehashes.\"merklehash\", testdata_merklehashes.\"merkle_filterpath\", "
	sqlToExecute = sqlToExecute + "testdata_merklehashes.\"merkle_filterpath_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata_merklehashes "

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "ab4dd291-a270-498b-9eb4-13da153f7afb",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err
	}

	// Variables to used when extract data from result set
	var testDataMerkleHash memCloudDBAllTestDataMerkleHashStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataMerkleHash.clientUuid, &testDataMerkleHash.domainUuid,
			&testDataMerkleHash.merkleHash, &testDataMerkleHash.merklePath)
		if err != nil {
			return err
		}
		*testDataMerkleHashs = append(*testDataMerkleHashs, testDataMerkleHash)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataMerkleTrees in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataMerkleTreesFromCloudDB(testDataMerkleTrees *[]memCloudDBAllTestDataMerkleTreeStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "6e0a6f4c-cc54-4aff-94f1-243aee6141ae",
	}).Debug("Entering: loadAllmemDBAllTestDataMerkleTreesFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "a6d6b40f-c30f-4ec1-b5d2-f94b0234471c",
		}).Debug("Exiting: loadAllmemDBAllTestDataMerkleTreesFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata_merkletrees.\"client_uuid\", testdata_merkletrees.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata_merkletrees.\"node_level\", testdata_merkletrees.\"node_name\", "
	sqlToExecute = sqlToExecute + "testdata_merkletrees.\"node_path\", testdata_merkletrees.\"node_hash\", "
	sqlToExecute = sqlToExecute + "testdata_merkletrees.\"node_child_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata_merkletrees "

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "e60b727a-ae76-4188-b159-634d80978658",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err
	}

	// Variables to used when extract data from result set
	var testDataMerkleTree memCloudDBAllTestDataMerkleTreeStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataMerkleTree.clientUuid, &testDataMerkleTree.domainUuid,
			&testDataMerkleTree.nodeLevel, &testDataMerkleTree.nodeName,
			&testDataMerkleTree.nodePath, &testDataMerkleTree.nodeHash,
			&testDataMerkleTree.nodeChildHash)
		if err != nil {
			return err
		}
		*testDataMerkleTrees = append(*testDataMerkleTrees, testDataMerkleTree)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestDataRowItems in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllmemDBAllTestDataRowItemsFromCloudDB(testDataRowItems *[]memCloudDBAllTestDataRowItemStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "61b8b021-9568-463e-b867-ac1ddb10584d",
	}).Debug("Entering: loadAllmemDBAllTestDataRowItemsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "78a97c41-a098-4122-88d2-01ed4b6c4844",
		}).Debug("Exiting: loadAllmemDBAllTestDataRowItemsFromCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata_row_items_current.\"client_uuid\", testdata_row_items_current.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata_row_items_current.\"row_hash\", testdata_row_items_current.\"testdata_value_as_string\", "
	sqlToExecute = sqlToExecute + "testdata_row_items_current.\"leaf_node_name\", testdata_row_items_current.\"leaf_node_path\" "
	sqlToExecute = sqlToExecute + "FROM testdata_row_items_current "

	// Query DB
	rows, err := DbPool.Query(context.Background(), sqlToExecute)

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":           "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
			"Error":        err,
			"sqlToExecute": sqlToExecute,
		}).Error("Something went wrong when executing SQL")

		return err
	}

	// Variables to used when extract data from result set
	var testDataRowItem memCloudDBAllTestDataRowItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataRowItem.clientUuid, &testDataRowItem.domainUuid,
			&testDataRowItem.rowHash, &testDataRowItem.testdataValueAsString,
			&testDataRowItem.leafNodeName, &testDataRowItem.leafNodePath)
		if err != nil {
			return err
		}
		// Add values to the object that is pointed to by variable in function
		*testDataRowItems = append(*testDataRowItems, testDataRowItem)

	}

	// No errors occurred
	return nil

}
