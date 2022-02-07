package main

/* START Cloud Sync
./cloud_sql_proxy -instances=mycloud-run-project:europe-north1:fenix-sqlserver=tcp:5432

*/

import (
	"github.com/sirupsen/logrus"
)

// Load TestData from CloudDB into memDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadTestDataFromCloudDB() (err error) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "fec5c67e-4679-4e42-bcc4-fa64f46d3b59",
	}).Debug("Incoming gRPC 'loadTestDataFromCloudDB'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "b5410c3f-ba1b-4d77-b85a-050985ee26fd",
	}).Debug("Outgoing gRPC 'loadTestDataFromCloudDB'")

	// Will not process anything while 'stateProcessIncomingAndOutgoingMessage' == false
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "36fa4890-69d1-4e68-940a-915fdadd7968",
		}).Info("Will not process 'loadTestDataFromCloudDB()' while stateProcessIncomingAndOutgoingMessage == false")
		return nil
	}

	// All TestTDataClients in CloudDB
	var tempMemDBAllClients []memCloudDBAllTestDataClientStruct

	err, tempMemCloudDBAllClientsMap := fenixTestDataSyncServerObject.loadAllmemDBAllClientsFromCloudDB(&tempMemDBAllClients)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "06e04586-c8ce-4172-8391-8fdd235b15ab",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllClientsFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataHeaderFilterValues in CloudDB
	var tempMemDBAllTestDataHeaderFilterValues []memCloudDBAllTestDataHeaderFilterValueStruct
	err = fenixTestDataSyncServerObject.loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB(&tempMemDBAllTestDataHeaderFilterValues)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "3457adf0-ac33-4fdb-a23e-ce0b000bb64e",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllTestDataHeaderFilterValuesFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataHeaderItems in CloudDB
	var tempMemDBAllTestDataHeaderItems []memCloudDBAllTestDataHeaderItemStruct
	err = fenixTestDataSyncServerObject.loadAllmemDBAllTestDataHeaderItemsFromCloudDB(&tempMemDBAllTestDataHeaderItems)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "a4e56ce6-2d5b-4941-91d2-2ebf9f91f5c2",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllTestDataHeaderItemsFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataMerkleHashes in CloudDB
	var tempMemDBAllTestDataMerkleHashes []memCloudDBAllTestDataMerkleHashStruct
	err = fenixTestDataSyncServerObject.loadAllmemDBAllTestDataMerkleHashesFromCloudDB(&tempMemDBAllTestDataMerkleHashes)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "202bbd5b-19a1-4ebd-923d-0114161e6c2b",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllTestDataMerkleHashesFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataMerkleTrees in CloudDB
	var tempMemDBAllTestDataMerkleTrees []memCloudDBAllTestDataMerkleTreeStruct
	err = fenixTestDataSyncServerObject.loadAllmemDBAllTestDataMerkleTreesFromCloudDB(&tempMemDBAllTestDataMerkleTrees)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "202bbd5b-19a1-4ebd-923d-0114161e6c2b",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllTestDataMerkleTreesFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// All TestDataRowItems in CloudDB
	var tempMemDBAllTestDataRowItems []memCloudDBAllTestDataRowItemStruct
	err = fenixTestDataSyncServerObject.loadAllmemDBAllTestDataRowItemsFromCloudDB(&tempMemDBAllTestDataRowItems)
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id":    "502b250c-1565-4a1b-bdf3-eea77382c957",
			"error": err,
		}).Error("Problem when executing: 'loadAllmemDBAllTestDataRowItemsFromCloudDB()'")

		fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = false
		return err
	}

	// Could get all data from CloudDB, so now we can move the tempObjects into memDB
	memCloudDBAllClients = tempMemDBAllClients
	memCloudDBAllClientsMap = tempMemCloudDBAllClientsMap
	memCloudDBAllTestDataHeaderFilterValues = tempMemDBAllTestDataHeaderFilterValues
	memCloudDBAllTestDataHeaderItems = tempMemDBAllTestDataHeaderItems
	memCloudDBAllTestDataMerkleHashes = tempMemDBAllTestDataMerkleHashes
	memCloudDBAllTestDataMerkleTrees = tempMemDBAllTestDataMerkleTrees
	memCloudDBAllTestDataRowItems = tempMemDBAllTestDataRowItems

	// Everything was loaded, now allow in- and outgoing messages
	fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage = true

	return err

}
