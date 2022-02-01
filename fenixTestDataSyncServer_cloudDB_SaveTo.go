package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

// ****************************************************************************************************************
// Save data to CloudDB
//
// All TestTDataClients
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllClientsToCloudDB(testDataClients *[]memCloudDBAllTestDataClientStruct) (err error) {

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
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

	// Variables to used when extract data from result set
	var testDataClient memCloudDBAllTestDataClientStruct

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
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataHeaderFilterValuesToCloudDB(testDataHeaderFilterValues *[]memCloudDBAllTestDataHeaderFilterValueStruct) (err error) {

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
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

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
// Save data to CloudDB
//
// All TestDataHeaderItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataHeaderItemsToCloudDB(testDataHeaderItems *[]memCloudDBAllTestDataHeaderItemStruct) (err error) {

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
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

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
// Save data to CloudDB
//
// All TestDataMerkleHashes
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataMerkleHashesToCloudDB(testDataMerkleHashs *[]memCloudDBAllTestDataMerkleHashStruct) (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "539d7f8e-4a69-4fe1-b4dd-5d5148d1a8b6",
	}).Debug("Entering: saveAllmemDBAllTestDataMerkleHashesToCloudDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "66ce8020-aec6-49ef-875e-f16bd36e7667",
		}).Debug("Exiting: saveAllmemDBAllTestDataMerkleHashesToCloudDB()")
	}()

	sqlToExecute := ""
	sqlToExecute = sqlToExecute + "SELECT testdata.merklehashes.\"client_uuid\", testdata.merklehashes.\"domain_uuid\", "
	sqlToExecute = sqlToExecute + "testdata.merklehashes.\"merklehash\", testdata.merklehashes.\"merkle_path\" "
	sqlToExecute = sqlToExecute + "FROM testdata.merklehashes "

	// Query DB
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

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
// Save data to CloudDB
//
// All TestDataMerkleTrees
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataMerkleTreesToCloudDB(testDataMerkleTrees *[]memCloudDBAllTestDataMerkleTreeStruct) (err error) {

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
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

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
// Save data to CloudDB
//
// All TestDataRowItems
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) saveAllmemDBAllTestDataRowItemsToCloudDB(testDataRowItems *[]memCloudDBAllTestDataRowItemStruct) (err error) {

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
	rows, _ := DbPool.Query(context.Background(), sqlToExecute)

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

func testSQL() {
	/*

	   	data := make(map[string]string)

	   	// Begin SQL Transaction
	   	txn, err := DbPool.Begin(context.Background())
	   	if err != nil {
	   		fmt.Println("err: ", err)
	   		return
	   	}
	   	defer txn.Commit(context.Background())

	   	sqlStr := "INSERT INTO test(n1, n2, n3) VALUES "
	   	vals := []interface{}{}

	   	for _, row := range data {
	   		sqlStr += "(?, ?, ?)," // Put "?" symbol equal to number of columns
	   		vals = append(vals, row["v1"], row["v2"], row["v3"]) // Put row["v{n}"] blocks equal to number of columns
	   	}

	   	//trim the last ,
	   	sqlStr = strings.TrimSuffix(sqlStr, ",")

	   	//Replacing ? with $n for postgres
	   	sqlStr = ReplaceSQL(sqlStr, "?")



	   	//prepare the statement
	   	stmt, _ := DbPool.SendBatch() Prepare(sqlStr)

	   	//format all vals at once
	   	res, _ := stmt.Exec(vals...)
	   }

	   // ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
	   func ReplaceSQL(old, searchPattern string) string {
	   	tmpCount := strings.Count(old, searchPattern)
	   	for m := 1; m <= tmpCount; m++ {
	   		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	   	}
	   	return old

	*/
}
