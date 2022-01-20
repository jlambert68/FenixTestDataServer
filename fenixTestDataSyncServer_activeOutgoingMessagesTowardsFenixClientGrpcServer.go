package main

import (
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Set upp connection and Dial to FenixTestDataSyncServer
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) SetConnectionToFenixClientTestDataSyncServer() {

	var err error

	// Set up connection to FenixTestDataSyncServer
	remoteFenixClientTestDataSyncServerConnection, err = grpc.Dial(fenixClientTestDataSyncServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixClientTestDataSyncServer_address_to_dial": fenixClientTestDataSyncServer_address_to_dial,
			"error message": err,
		}).Error("Did not connect to FenixClientTestDataSyncServer via gRPC")
		//os.Exit(0)
	} else {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixClientTestDataSyncServer_address_to_dial": fenixClientTestDataSyncServer_address_to_dial,
		}).Info("gRPC connection OK to FenixClientTestDataSyncServer")

		// Creates a new Clients
		fenixClientTestDataSyncServerClient = fenixClientTestDataSyncServerGrpcApi.NewFenixClientTestDataGrpcServicesClient(remoteFenixClientTestDataSyncServerConnection)

	}
}

// Get the highest FenixProtoFileVersionEnumeration
func (fenixClientTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getHighestFenixProtoFileVersion() int32 {

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

// Get the highest ClientProtoFileVersionEnumeration
func (fenixClientTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) getHighestClientProtoFileVersion() int32 {

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

// Check if Calling Client is using correct proto-file version
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) isClientUsingCorrectProtoFileVersion(usedProtoFileVersion fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum) (clientUseCorrectProtoFileVersion bool, protoFileExpected fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum, protoFileUsed fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum) {

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestDataSyncServerGrpcApi.CurrentFenixTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestFenixProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = true
	}

	//protoFileExpectedDescription := protoFileExpected.String()
	//protoFileExpectedDescription := protoFileExpected.String()

	return clientUseCorrectProtoFileVersion, protoFileExpected, protoFileUsed
}

// Fenix Server asks Fenix client to register itself with the Fenix Testdata sync server
//func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct)  RegisterTestDataClient(EmptyParameter) returns (AckNackResponse) {
//}

// Fenix Server asks Fenix client to send TestData MerkleHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) AskClientToSendMerkleHash(TestDataClientGuid string) {

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByCaller: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendMerkleHash(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "6b080a23-4e06-4d16-8295-a67ba7115a56",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleHash'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "44671efb-e24d-450a-acba-006cc248d058",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleHash'")
		}
	}
}

// Fenix Server asks Fenix client to send TestData MerkleTree to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) AskClientToSendMerkleTree(TestDataClientGuid string) {

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByCaller: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendMerkleTree(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "23ed314f-f7d6-41da-8489-07dfc970ab31",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleTree'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bba9d885-6dc6-4bd1-9f48-5928e22552ec",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleTree'")
		}
	}

}

// Fenix Server asks Fenix client to send TestDataHeaderHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) AskClientToSendTestDataHeaderHash(TestDataClientGuid string) {

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByCaller: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendTestDataHeaderHash(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "ef7d59cc-88b5-447e-83f5-ac018f2320bd",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaderHash'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bac41696-c8a3-4d11-ac1c-68965c8a1572",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaderHash'")
		}
	}

}

// Fenix Server asks Fenix client to send TestDataHeaders to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) AskClientToSendTestDataHeaders(TestDataClientGuid string) {

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByCaller: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendTestDataHeaders(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "ef7d59cc-88b5-447e-83f5-ac018f2320bd",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaders'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.Acknack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bac41696-c8a3-4d11-ac1c-68965c8a1572",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaders'")
		}
	}

}

// Fenix Server asks Fenix client to  send TestData rows, based on list of MerklePaths, to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObject_struct) AskClientToSendTestDataRows(TestDataClientGuid string) {

}
