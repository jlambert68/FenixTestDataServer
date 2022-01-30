package main

import (
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
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

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Check if Calling Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) isClientUsingCorrectTestDataProtoFileVersion(usedProtoFileVersion fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum) (returnMessage *fenixTestDataSyncServerGrpcApi.AckNackResponse) {

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

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func (fenixClientTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestFenixTestDataProtoFileVersion() int32 {

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
func (fenixClientTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) getHighestClientTestDataProtoFileVersion() int32 {

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
