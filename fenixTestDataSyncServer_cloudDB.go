package main

/* START Cloud Sync
./cloud_sql_proxy -instances=mycloud-run-project:europe-north1:fenix-sqlserver=tcp:5432

*/

import (
	"github.com/sirupsen/logrus"
)

// Load TestData from CloudDB into memDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadNecessaryTestDataFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "fec5c67e-4679-4e42-bcc4-fa64f46d3b59",
	}).Debug("Incoming gRPC 'loadNecessaryTestDataFromCloudDB'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b5410c3f-ba1b-4d77-b85a-050985ee26fd",
	}).Debug("Outgoing gRPC 'loadNecessaryTestDataFromCloudDB'")

	// Will not process anything while 'stateProcessIncomingAndOutgoingMessage' == false
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "36fa4890-69d1-4e68-940a-915fdadd7968",
		}).Info("Will not process 'loadNecessaryTestDataFromCloudDB()' while stateProcessIncomingAndOutgoingMessage == false")
		return nil
	}

	// All TestTDataClients in CloudDB
	var tempCloudDBClients []cloudDBTestDataClientStruct
	var tempCloudDBClientsMap cloudDBClientsMapType

	tempCloudDBClientsMap, err = fenixTestDataSyncServerObject.loadAllClientsFromCloudDB(&tempCloudDBClients)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "06e04586-c8ce-4172-8391-8fdd235b15ab",
			"error": err,
		}).Error("Problem when executing: 'loadAllClientsFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// Move to from temp-variables
	cloudDBClients = tempCloudDBClients
	cloudDBClientsMap = tempCloudDBClientsMap

	return nil
}

/*
	// All TestDataHeaderFilterValues in CloudDB
	var tempMemDBAllTestDataHeaderFilterValues []cloudDBTestDataHeaderFilterValuesStruct
	err = fenixTestDataSyncServerObject.loadAllTestDataHeaderFilterValuesForClientFromCloudDB(clientUuid, &tempMemDBAllTestDataHeaderFilterValues)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "3457adf0-ac33-4fdb-a23e-ce0b000bb64e",
			"error": err,
		}).Error("Problem when executing: 'loadAllTestDataHeaderFilterValuesForClientFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataHeaderItems in CloudDB
	var tempMemDBAllTestDataHeaderItems []cloudDBTestDataHeaderItemStruct
	err = fenixTestDataSyncServerObject.loadAllTestDataHeaderItemsForClientFromCloudDB(&tempMemDBAllTestDataHeaderItems)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "a4e56ce6-2d5b-4941-91d2-2ebf9f91f5c2",
			"error": err,
		}).Error("Problem when executing: 'loadAllTestDataHeaderItemsForClientFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataMerkleHashes in CloudDB
	var tempMemDBAllTestDataMerkleHashes []cloudDBTestDataMerkleHashStruct
	err = fenixTestDataSyncServerObject.loadAllTestDataMerkleHashesForClientFromCloudDB(&tempMemDBAllTestDataMerkleTrees)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "40d40c74-f1b5-4dca-87f7-2f12eb545183",
			"error": err,
		}).Error("Problem when executing: 'loadAllTestDataMerkleHashesForClientFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataMerkleTrees in CloudDB
	var tempMemDBAllTestDataMerkleTrees []cloudDBTestDataMerkleTreeStruct
	err = fenixTestDataSyncServerObject.loadAllTestDataMerkleTreesForClientFromCloudDB(&tempMemDBAllTestDataMerkleTrees)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "202bbd5b-19a1-4ebd-923d-0114161e6c2b",
			"error": err,
		}).Error("Problem when executing: 'loadAllTestDataMerkleTreesForClientFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataRowItems in CloudDB
	var tempMemDBAllTestDataRowItems []cloudDBTestDataRowItemCurrentStruct
	err = fenixTestDataSyncServerObject.loadAllTestDataRowItemsForClientFromCloudDB(&tempMemDBAllTestDataRowItems)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "502b250c-1565-4a1b-bdf3-eea77382c957",
			"error": err,
		}).Error("Problem when executing: 'loadAllTestDataRowItemsForClientFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// Could get all data from CloudDB, so now we can move the tempObjects into memDB
	cloudDBClients = tempMemDBAllClients
	cloudDBClientsMap = tempMemCloudDBAllClientsMap
	cloudDBTestDataHeadersFilterValues = tempMemDBAllTestDataHeaderFilterValues
	cloudDBTestDataHeaderItems = tempMemDBAllTestDataHeaderItems
	cloudDBTestDataMerkleHashes = tempMemDBAllTestDataMerkleHashes
	cloudDBTestDataMerkleTrees = tempMemDBAllTestDataMerkleTrees
	cloudDBTestDataRowItems = tempMemDBAllTestDataRowItems

	// Create the working copies in 'dbDataMap-structure'  for each Client


	// Everything was loaded, now allow in- and outgoing messages


	return err

}
*/
