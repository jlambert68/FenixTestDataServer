package main

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) existsClinetInDB(testDataClientUUID string) (bool, error) {

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
