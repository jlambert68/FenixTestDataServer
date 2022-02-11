package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

// ****************************************************************************************************************
// Load data from CloudDB into memory structures, to speed up stuff
//
// All TestTDataClients in CloudDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllClientsFromCloudDB(testDataClients *[]cloudDBTestDataClientStruct) (err error, memCloudDBAllClientsMap cloudDBClientsMapType) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "16af90a4-aa07-4d8b-921a-a47c04811a9b",
	}).Debug("Entering: loadAllClientsFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "e9659490-9ba7-437b-a235-88d8369ebf36",
		}).Debug("Exiting: loadAllClientsFromCloudDB()")
	}()

	/* Example
	SELECT c.*
	FROM public.clients c
	WHERE c.activated = true
	AND
	      c.replaced_by_new_version = false
	AND
	      c.client_areatyp_id = 1; // Clients used for 'TestData'

	client_uuid             uuid      not null
	client_name             varchar   not null,
	domain_uuid             uuid      not null
	description             varchar,
	activated               boolean   not null,
	deleted                 boolean   not null,
	update_timestamp        timestamp not null,
	replaced_by_new_version boolean   not null,
	client_id               integer   not null,
	client_version          integer   not null,
	client_areatyp_id       integer   not null

	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT c.client_uuid, c.client_name, c.domain_uuid, c.description "
	sqlToExecute = sqlToExecute + "FROM clients c"
	sqlToExecute = sqlToExecute + "WHERE c.activated = true "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "c.deleted = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "c.replaced_by_new_version = false "
	sqlToExecute = sqlToExecute + "AND "
	sqlToExecute = sqlToExecute + "c.client_areatyp_id = 1 " // Clients used for 'TestData'

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
	var testDataClient cloudDBTestDataClientStruct
	memCloudDBAllClientsMap = make(map[memDBClientUuidType]cloudDBTestDataClientMapStruct)

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataClient.clientUuid,
			&testDataClient.clientName,
			&testDataClient.domainUuid,
			&testDataClient.description)

		if err != nil {
			return err, nil
		}

		*testDataClients = append(*testDataClients, testDataClient)
		memCloudDBAllClientsMap[memDBClientUuidType(testDataClient.clientUuid)] = cloudDBTestDataClientMapStruct{
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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataHeaderFilterValuesForClientFromCloudDB(clientUuid string, testDataHeaderFilterValues *[]cloudDBTestDataHeaderFilterValuesStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "097f9f31-f29b-4c4a-aadb-0d4120429cf5",
	}).Debug("Entering: loadAllTestDataHeaderFilterValuesForClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "cfb2f17f-579d-4d98-b91b-ba3c01a32771",
		}).Debug("Exiting: loadAllTestDataHeaderFilterValuesForClientFromCloudDB()")
	}()

	/* Example

	SELECT tdhfv.*
	FROM public.testdata_header_filtervalues tdhfv
	WHERE tdhfv.client_uuid::text = '45a217d1-55ed-4531-a801-779e566d75cb';

	header_item_hash          varchar   not null
	header_filter_value       varchar   not null,
	client_uuid               uuid      not null
	header_filter_value_order integer   not null,
	header_filter_values_hash varchar   not null,
	updated_timestamp         timestamp not null,
	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tdhfv.* "
	sqlToExecute = sqlToExecute + "FROM public.testdata_header_filtervalues tdhfv "
	sqlToExecute = sqlToExecute + "WHERE tdhfv.client_uuid::text = '" + clientUuid + "';"

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
	var testDataHeaderFilterValue cloudDBTestDataHeaderFilterValuesStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataHeaderFilterValue.headerItemHash,
			&testDataHeaderFilterValue.headerFilterValue,
			&testDataHeaderFilterValue.clientUuid,
			&testDataHeaderFilterValue.headerFilterValueOrder,
			&testDataHeaderFilterValue.headerFilterValuesHash,
			&testDataHeaderFilterValue.updatedTimeStamp)

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataHeaderItemsForClientFromCloudDB(clientUuid string, testDataHeaderItems *[]cloudDBTestDataHeaderItemStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "a2951e5e-d7d0-4240-88a9-5dc570f2bbe9",
	}).Debug("Entering: loadAllTestDataHeaderItemsForClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "70e1c4da-5001-4199-adf4-bbc2576ccdab",
		}).Debug("Exiting: loadAllTestDataHeaderItemsForClientFromCloudDB()")
	}()

	/* Example

	SELECT tdhi.*
	FROM public.testdata_header_items tdhi
	WHERE tdhi.client_uuid::text = '45a217d1-55ed-4531-a801-779e566d75cb';

	client_uuid              uuid      not null
	updated_timestamp        timestamp not null,
	header_item_hash         varchar   not null
	header_label             varchar   not null,
	should_be_used_in_filter boolean   not null,
	is_mandatory_in_filter   boolean   not null,
	filter_selection_type    integer   not null
	header_column_order      integer   not null,
	header_items_hash        varchar   not null

	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tdhi.* "
	sqlToExecute = sqlToExecute + "FROM public.testdata_header_items tdhi "
	sqlToExecute = sqlToExecute + "WHERE tdhi.client_uuid::text = '" + clientUuid + "';"

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
	var testDataHeaderItem cloudDBTestDataHeaderItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataHeaderItem.clientUuid,
			&testDataHeaderItem.updatedTimeStamp,
			&testDataHeaderItem.headerItemHash,
			&testDataHeaderItem.headerLabel,
			&testDataHeaderItem.shouldBeUsedInFilter,
			&testDataHeaderItem.isMandatoryInFilter,
			&testDataHeaderItem.filterSelectionType,
			&testDataHeaderItem.headerColumnOrder,
			&testDataHeaderItem.headerItemsHash)

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataMerkleHashesForClientFromCloudDB(clientUuid string, testDataMerkleHashs *[]cloudDBTestDataMerkleHashStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "53461f88-c773-477e-b459-cfb93a1c3eaa",
	}).Debug("Entering: loadAllTestDataMerkleHashesForClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "851cbecf-4084-4e38-b922-eab2a4b526d1",
		}).Debug("Exiting: loadAllTestDataMerkleHashesForClientFromCloudDB()")
	}()

	/* Example

		SELECT tdmh.*
		FROM public.testdata_merklehashes tdmh
		WHERE tdmh.client_uuid::text = '45a217d1-55ed-4531-a801-779e566d75cb';

	   client_uuid            uuid      not null
	   updated_timestamp      timestamp not null,
	   merklehash             varchar   not null
	   merkle_filterpath      varchar   not null,
	   merkle_filterpath_hash varchar   not null

	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tdmh.* "
	sqlToExecute = sqlToExecute + "FROM public.testdata_merklehashes tdmh "
	sqlToExecute = sqlToExecute + "WHERE tdmh.client_uuid::text = '" + clientUuid + "';"

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
	var testDataMerkleHash cloudDBTestDataMerkleHashStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataMerkleHash.clientUuid,
			&testDataMerkleHash.updatedTimeStamp,
			&testDataMerkleHash.merkleHash,
			&testDataMerkleHash.merkleFilterPath,
			&testDataMerkleHash.merkleFilterPathHash)

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataMerkleTreesForClientFromCloudDB(clientUuid string, testDataMerkleTrees *[]cloudDBTestDataMerkleTreeStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "6e0a6f4c-cc54-4aff-94f1-243aee6141ae",
	}).Debug("Entering: loadAllTestDataMerkleTreesForClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "a6d6b40f-c30f-4ec1-b5d2-f94b0234471c",
		}).Debug("Exiting: loadAllTestDataMerkleTreesForClientFromCloudDB()")
	}()

	/* Example

	SELECT tdmt.*
	FROM public.testdata_merkletrees tdmt
	WHERE tdmt.client_uuid::text = '45a217d1-55ed-4531-a801-779e566d75cb';

	    client_uuid       uuid      not null
	    node_level        integer   not null,
	    node_name         varchar   not null,
	    node_path         varchar   not null,
	    node_hash         varchar   not null,
	    node_child_hash   varchar   not null
	    updated_timestamp timestamp not null,
	    merkleHash      varchar   not null
	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tdmt.* "
	sqlToExecute = sqlToExecute + "public.testdata_merkletrees tdmt "
	sqlToExecute = sqlToExecute + "WHERE tdmt.client_uuid::text = '" + clientUuid + "';"

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
	var testDataMerkleTree cloudDBTestDataMerkleTreeStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataMerkleTree.clientUuid,
			&testDataMerkleTree.nodeLevel,
			&testDataMerkleTree.nodeName,
			&testDataMerkleTree.nodePath,
			&testDataMerkleTree.nodeHash,
			&testDataMerkleTree.nodeChildHash,
			&testDataMerkleTree.updatedTimeStamp,
			&testDataMerkleTree.merkleHash)

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadAllTestDataRowItemsForClientFromCloudDB(clientUuid string, testDataRowItems *[]cloudDBTestDataRowItemStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "61b8b021-9568-463e-b867-ac1ddb10584d",
	}).Debug("Entering: loadAllTestDataRowItemsForClientFromCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "78a97c41-a098-4122-88d2-01ed4b6c4844",
		}).Debug("Exiting: loadAllTestDataRowItemsForClientFromCloudDB()")
	}()

	/* Example
	SELECT tdri.*
	FROM public.testdata_row_items_current tdri
	WHERE tdri.client_uuid::text = '45a217d1-55ed-4531-a801-779e566d75cb';

	client_uuid              uuid      not null
	row_hash                 varchar   not null,
	testdata_value_as_string varchar   not null,
	updated_timestamp        timestamp not null,
	leaf_node_name           varchar   not null,
	leaf_node_path           varchar   not null,
	leaf_node_hash           varchar   not null
	value_column_order       integer   not null,
	value_row_order          integer   not null,

	*/

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT tdri.* "
	sqlToExecute = sqlToExecute + "public.testdata_row_items_current tdri "
	sqlToExecute = sqlToExecute + "WHERE tdri.client_uuid::text = '" + clientUuid + "';"

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
	var testDataRowItem cloudDBTestDataRowItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(
			&testDataRowItem.clientUuid,
			&testDataRowItem.rowHash,
			&testDataRowItem.testdataValueAsString,
			&testDataRowItem.updatedTimeStamp,
			&testDataRowItem.leafNodeName,
			&testDataRowItem.leafNodePath,
			&testDataRowItem.leafNodeHash,
			&testDataRowItem.valueColumnOrder,
			&testDataRowItem.valueRowOrder)

		if err != nil {
			return err
		}
		// Add values to the object that is pointed to by variable in function
		*testDataRowItems = append(*testDataRowItems, testDataRowItem)

	}

	// No errors occurred
	return nil

}
