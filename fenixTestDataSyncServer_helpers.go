package main

import (
	"FenixTestDataServer/common_config"
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// ********************************************************************************************************************
// Check if there is a temporary stop in processing incoming or outgoing messages
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isThereATemporaryStopInProcessingInOrOutgoingMessages() (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	// Check if there is a temporary stop in processing in- or outgoing messages
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "04df8862-3e76-45aa-9cee-9858c348e18d",
		}).Debug("There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server")

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if there is a temporary stop in processing incoming or outgoing messages
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isThereATemporaryStopInProcessingInOrOutgoingMessagesAdmin() (returnMessage *fenixTestDataSyncServerGrpcAdminApi.AckNackResponse) {

	// Check if there is a temporary stop in processing in- or outgoing messages
	if fenixTestDataSyncServerObject.stateProcessIncomingAndOutgoingMessage == false {
		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum_ERROR_TEMPORARY_STOP_IN_PROCESSING
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{
			AckNack:    false,
			Comments:   "There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "2e45c69b-a8c5-4b20-b40b-8a3e4025ed4c",
		}).Debug("There is a temporary stop in processing any ingoing or outgoing messages at Fenix TestSync-server")

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum
	var protoFileUsed fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
		}).Debug("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Admin-Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isAdminClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum) (returnMessage *fenixTestDataSyncServerGrpcAdminApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum
	var protoFileUsed fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestDataSyncServerGrpcAdminApi.CurrentFenixTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcAdminApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcAdminApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "7831224f-5d6b-4afe-98d8-17a1acd6c799",
		}).Debug("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client has Hashed the MerklePath correctly
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientsMerklePathCorrectlyHashed(callingClientUuid string, merklePath string, merklePathHash string) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

	var hashIsCorrectlyHashed bool

	// Verify that MerkleFilterPath is hashed correctly
	tempHashedMerkleFilter := common_config.HashSingleValue(merklePath)

	// Check if MerklePath is correctly hashed
	if tempHashedMerkleFilter == merklePathHash {
		hashIsCorrectlyHashed = true
	} else {
		hashIsCorrectlyHashed = false
	}

	// Generate returnMessage if wrongly hashed
	if hashIsCorrectlyHashed == false {

		// Set Error codes to return message
		var errorCodes []fenixTestDataSyncServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestDataSyncServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestDataSyncServerGrpcApi.ErrorCodesEnum_ERROR_MERKLEPATHHASH_IS_NOT_CORRECT_CALCULATED
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestDataSyncServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "MerklePathHash is not correct calculated for MerklePath='" + merklePath + "' . Expected: '" + tempHashedMerkleFilter + "', but got: '" + merklePathHash + "'",
			ErrorCodes: errorCodes,
		}

		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"id": "57f3ef6b-ae39-40b3-904b-6f3e28ac2d32",
		}).Debug("MerklePathHash is not correct calculated for MerklePath='" + merklePath + "' . Expected: '" + tempHashedMerkleFilter + "', but got: '" + merklePathHash + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestFenixTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixProtoFileVersion != -1 {
		return highestFenixProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixProtoFileVersion = maxValue

	return highestFenixProtoFileVersion
}

// ********************************************************************************************************************
// Get the highest ClientProtoFileVersionEnumeration
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestClientTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestclientProtoFileVersion' saved, if so use that one
	if highestClientProtoFileVersion != -1 {
		return highestClientProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestClientProtoFileVersion = maxValue

	return highestClientProtoFileVersion
}
