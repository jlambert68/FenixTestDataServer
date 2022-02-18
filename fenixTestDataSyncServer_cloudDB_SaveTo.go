package main

import "C"
import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
)

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestTDataClients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllClientsToCloudDB(testDataClients *[]cloudDBTestDataClientStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "60b8ec33-b847-4904-a2a2-705d89455ce3",
	}).Debug("Entering: saveAllmemDBAllClientsToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "3b7e8422-ae5a-4925-843c-f5d05828e5e1",
		}).Debug("Exiting: saveAllmemDBAllClientsToCloudDB()")
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
	rows, _ := fenixSyncShared.DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataClient cloudDBTestDataClientStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataClient.clientUuid, &testDataClient.clientName, &testDataClient.domainUuid, &testDataClient.description)
		if err != nil {
			return err
		}
		*testDataClients = append(*testDataClients, testDataClient)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestDataHeaderFilterValues
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataHeaderFilterValuesToCloudDB(testDataHeaderFilterValues *[]cloudDBTestDataHeaderFilterValuesStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "a015ff6a-7a76-47a4-b61a-dec2b3ea3f7f",
	}).Debug("Entering: saveAllmemDBAllTestDataHeaderFilterValuesToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "0e995a9e-a872-4956-890f-e2b0a87afd13",
		}).Debug("Exiting: saveAllmemDBAllTestDataHeaderFilterValuesToCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.header.filtervalues.\"header_item_hash\", testdata.header.filtervalues.\"header_filter_value\" , testdata.header.filtervalues.\"client_uuid\", testdata.header.filtervalues.\"domain_uuid\"  "
	sqlToExecute = sqlToExecute + "FROM testdata.header.filtervalues "

	// Query DB
	rows, _ := fenixSyncShared.DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataHeaderFilterValue cloudDBTestDataHeaderFilterValuesStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderFilterValue.headerItemHash, &testDataHeaderFilterValue.headerFilterValue, &testDataHeaderFilterValue.clientUuid, &testDataHeaderFilterValue.clientUuid)
		if err != nil {
			return err
		}
		*testDataHeaderFilterValues = append(*testDataHeaderFilterValues, testDataHeaderFilterValue)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestDataHeaderItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataHeaderItemsToCloudDB(testDataHeaderItems *[]cloudDBTestDataHeaderItemStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "575cbf13-51d8-4f35-9edf-b36385927f1d",
	}).Debug("Entering: saveAllmemDBAllTestDataHeaderItemsToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "2d7ad1cf-4518-4931-8a4f-f0068cb84831",
		}).Debug("Exiting: saveAllmemDBAllTestDataHeaderItemsToCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.header.items.\"client_uuid\", testdata.header.items.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"header_item_hash\", testdata.header.items.\"header_label\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"should_be_used_in_filter\", testdata.header.items.\"is_mandatory_in_filter\", "
	sqlToExecute = sqlToExecute + "testdata.header.items.\"filter_selection_type\", testdata.header.items.\"filter_values_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata.header.items "

	// Query DB
	rows, _ := fenixSyncShared.DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataHeaderItem cloudDBTestDataHeaderItemStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataHeaderItem.clientUuid, &testDataHeaderItem.headerItemsHash, &testDataHeaderItem.headerItemHash,
			&testDataHeaderItem.headerLabel, &testDataHeaderItem.shouldBeUsedInFilter, &testDataHeaderItem.isMandatoryInFilter,
			&testDataHeaderItem.filterSelectionType, &testDataHeaderItem)
		if err != nil {
			return err
		}
		*testDataHeaderItems = append(*testDataHeaderItems, testDataHeaderItem)

	}

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestDataMerkleHashes
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveTestDataMerkleHashDataFromMemDBToCloudDB(dbTransaction pgx.Tx, currentUserUuid string) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "539d7f8e-4a69-4fe1-b4dd-5d5148d1a8b6",
	}).Debug("Entering: saveTestDataMerkleHashDataFromMemDBToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "66ce8020-aec6-49ef-875e-f16bd36e7667",
		}).Debug("Exiting: saveTestDataMerkleHashDataFromMemDBToCloudDB()")
	}()

	var dataRowToBeInsertedMultiType []interface{}
	var dataRowsToBeInsertedMultiType [][]interface{}
	sqlToExecute := ""

	// Create Delete Statement for removing current MerkleTree-data for Client
	sqlToExecute = sqlToExecute + "DELETE FROM public.testdata_merkletrees "
	sqlToExecute = sqlToExecute + "WHERE client_uuid = '" + currentUserUuid + "'; "

	// Create Delete Statement for removing current MerkleHash-data for Client
	sqlToExecute = sqlToExecute + "DELETE FROM public.testdata_merklehashes "
	sqlToExecute = sqlToExecute + "WHERE client_uuid = '" + currentUserUuid + "'; "

	// Data to be inserted in the DB-table
	/*
		var dataToBeInserted = [][]string{
			{currentUserUuid,
				fenixSyncShared.GenerateDatetimeTimeStampForDB(),
				fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(currentUserUuid),
				fenixTestDataSyncServerObject.getCurrentMerkleFilterPathForServer(currentUserUuid),
				fenixTestDataSyncServerObject.getCurrentMerkleFilterPathHashForServer(currentUserUuid)},
		}
	*/
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, currentUserUuid)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixSyncShared.GenerateDatetimeTimeStampForDB())
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(currentUserUuid))
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixTestDataSyncServerObject.getCurrentMerkleFilterPathForServer(currentUserUuid))
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixTestDataSyncServerObject.getCurrentMerkleFilterPathHashForServer(currentUserUuid))

	dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)

	// Create Insert Statement for current MerkleHash-data for Client
	sqlToExecute = sqlToExecute + "INSERT INTO public.testdata_merklehashes "
	sqlToExecute = sqlToExecute + "(client_uuid, updated_timestamp, merklehash, merkle_filterpath, merkle_filterpath_hash) "
	sqlToExecute = sqlToExecute + fenixTestDataSyncServerObject.generateSQLInsertValues(dataRowsToBeInsertedMultiType)
	sqlToExecute = sqlToExecute + ";"

	// Data to be inserted in the DB-table
	dataRowsToBeInsertedMultiType = nil
	merkleTreeNodes := dbDataMap[memDBClientUuidType(currentUserUuid)].serverData.merkleTreeNodes //merkleTree

	for _, merkleTreeNode := range merkleTreeNodes {

		dataRowToBeInsertedMultiType = nil
		/*
			dataRowToBeInserted = []string{
				currentUserUuid,
				strconv.FormatInt(int64(merkleTreeNode.nodeLevel), 10),
				merkleTreeNode.nodeName,
				merkleTreeNode.nodePath,
				merkleTreeNode.nodeHash,
				merkleTreeNode.nodeChildHash,
				fenixSyncShared.GenerateDatetimeTimeStampForDB(),
				fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(currentUserUuid),
			}
		*/
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, currentUserUuid)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, merkleTreeNode.nodeLevel)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, merkleTreeNode.nodeName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, merkleTreeNode.nodePath)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, merkleTreeNode.nodeHash)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, merkleTreeNode.nodeChildHash)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixSyncShared.GenerateDatetimeTimeStampForDB())
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, fenixTestDataSyncServerObject.getCurrentMerkleHashForServer(currentUserUuid))

		dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)

	}

	// Create Insert Statement for current MerkleTree-data for Client
	sqlToExecute = sqlToExecute + "INSERT INTO public.testdata_merkletrees "
	sqlToExecute = sqlToExecute + "(client_uuid, node_level, node_name, node_filter_path, node_hash, node_child_hash, updated_timestamp, \"merkleHash\") "
	sqlToExecute = sqlToExecute + fenixTestDataSyncServerObject.generateSQLInsertValues(dataRowsToBeInsertedMultiType)
	sqlToExecute = sqlToExecute + ";"

	// Create Delete Statement for removing current TestData for Client

	// Create Insert Statement for current MerkleTree-data for Client
	//TODO Fix order of deletion and order of Insert
	// Query DB
	comandTag, err := dbTransaction.Exec(context.Background(), sqlToExecute)

	if err != nil {
		return err
	}

	// Log response from CloudDB
	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id":                       "539d7f8e-4a69-4fe1-b4dd-5d5148d1a8b6",
		"comandTag.Insert()":       comandTag.Insert(),
		"comandTag.Delete()":       comandTag.Delete(),
		"comandTag.Select()":       comandTag.Select(),
		"comandTag.Update()":       comandTag.Update(),
		"comandTag.RowsAffected()": comandTag.RowsAffected(),
		"comandTag.String()":       comandTag.String(),
	}).Debug("Return data for SQL executed in database")

	// No errors occurred
	return nil

}

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestDataMerkleTrees
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataMerkleTreesToCloudDB(testDataMerkleTrees *[]cloudDBTestDataMerkleTreeStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "2c764bc6-4bf2-420a-8164-229866365c8c",
	}).Debug("Entering: saveAllmemDBAllTestDataMerkleTreesToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "e3bebbcf-6615-4ff9-b78b-9205ec420fb6",
		}).Debug("Exiting: saveAllmemDBAllTestDataMerkleTreesToCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.merkletrees.\"client_uuid\", testdata.merkletrees.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_level\", testdata.merkletrees.\"node_name\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_path\", testdata.merkletrees.\"node_hash\", "
	sqlToExecute = sqlToExecute + "testdata.merkletrees.\"node_child_hash\" "
	sqlToExecute = sqlToExecute + "FROM testdata.merkletrees "

	// Query DB
	rows, _ := fenixSyncShared.DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataMerkleTree cloudDBTestDataMerkleTreeStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataMerkleTree.clientUuid, &testDataMerkleTree.merkleHash,
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
// Save data to CloudDB
//
// All TestDataRowItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataRowItemsToCloudDB(testDataRowItems *[]cloudDBTestDataRowItemCurrentStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "28f028a7-79d8-449d-a2c0-0fe2119c1586",
	}).Debug("Entering: saveAllmemDBAllTestDataRowItemsToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "6d805914-c0d6-4386-b4d7-2a0fba2b4858",
		}).Debug("Exiting: saveAllmemDBAllTestDataRowItemsToCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.row.items.current.\"client_uuid\", testdata.row.items.current.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.row.items.current.\"row_hash\", testdata.row.items.current.\"testdata_value_as_string\", "
	sqlToExecute = sqlToExecute + "testdata.row.items.current.\"leaf_node_name\", testdata.row.items.current.\"leaf_node_path\" "
	sqlToExecute = sqlToExecute + "FROM testdata.row.items.current "

	// Query DB
	rows, _ := fenixSyncShared.DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataRowItem cloudDBTestDataRowItemCurrentStruct

	// Extract data from DB result set
	for rows.Next() {
		err := rows.Scan(&testDataRowItem.clientUuid, &testDataRowItem.clientUuid,
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

func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveMerkleHashMerkleTreeAndTestDataRowsToCloudDB(currentClientUuid string) (err error) {

	// Begin SQL Transaction
	txn, err := fenixSyncShared.DbPool.Begin(context.Background())
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer txn.Commit(context.Background())

	// Save MerkleHash-data
	err = fenixTestDataSyncServerObject.saveTestDataMerkleHashDataFromMemDBToCloudDB(txn, currentClientUuid)
	if err != nil {

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id":    "07b91f77-db17-484f-8448-e53375df94ce",
			"error": err,
		}).Error("Couldn't Save MerkleHash-data in database for Client: ", currentClientUuid)

		// Stop process in and outgoing messages
		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = true

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "348629ad-c358-4043-81ca-ff5f73b579c5",
		}).Error("Stop process in and outgoing messages")

		// Rollback any SQL transactions
		txn.Rollback(context.Background())

		// Clear memoryDB for Server
		_ = fenixTestDataSyncServerObject.clearCurrentMerkleDataAndTestDataRowsForServer(currentClientUuid)
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "1557cc22-c291-45f7-b85b-008b38e60b0b",
		}).Error("Clearing memoryDB for Server, regarding MerkleHash, MerkleTree and TestDataRows")

		return err

	}

	return nil
}

/*
// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
*/

// Generates all "VALUES('xxx', 'yyy')..." for insert statements
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) generateSQLInsertValues(testdata [][]interface{}) (sqlInsertValuesString string) {

	sqlInsertValuesString = ""

	// Loop over both rows and values
	for _, rowValues := range testdata {
		sqlInsertValuesString = sqlInsertValuesString + "VALUES("

		for valueCounter, value := range rowValues {
			switch valueType := value.(type) {

			case int:
				sqlInsertValuesString = sqlInsertValuesString + fmt.Sprint(value)
			case string:

				sqlInsertValuesString = sqlInsertValuesString + "'" + fmt.Sprint(value) + "'"

			default:
				fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
					"id": "53539786-cbb6-418d-8752-c2e337b9e962",
				}).Fatal("Unhandled type, %valueType", valueType)
			}

			// After the last value then add ')'
			if valueCounter == len(rowValues)-1 {
				sqlInsertValuesString = sqlInsertValuesString + ") "
			} else {
				// Not last value, so Add ','
				sqlInsertValuesString = sqlInsertValuesString + ", "
			}

		}

	}

	return sqlInsertValuesString
}
