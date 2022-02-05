package main

import (
	"FenixTestDataServer/common_config"
	"crypto/tls"
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Set upp connection and Dial to FenixTestDataSyncServer
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) SetConnectionToFenixClientTestDataSyncServer() {

	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if common_config.ExecutionLocationForFenixTestDataServer == common_config.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if common_config.ExecutionLocationForFenixTestDataServer == common_config.GCP {
		// Run on GCP
		remoteFenixClientTestDataSyncServerConnection, err = grpc.Dial(fenixClientTestDataSyncServer_address_to_dial, opts...)
	} else {
		// Run Local
		remoteFenixClientTestDataSyncServerConnection, err = grpc.Dial(fenixClientTestDataSyncServer_address_to_dial, grpc.WithInsecure())
	}

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

// Fenix Server asks Fenix client to register itself with the Fenix Testdata sync server
//func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct)  RegisterTestDataClient(EmptyParameter) returns (AckNackResponse) {
//}

// Fenix Server asks Fenix client to send TestData MerkleHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendMerkleHash(TestDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
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
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "44671efb-e24d-450a-acba-006cc248d058",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleHash'")
		}
	}
}

// Fenix Server asks Fenix client to send TestData MerkleTree to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendMerkleTree(TestDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
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
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bba9d885-6dc6-4bd1-9f48-5928e22552ec",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleTree'")
		}
	}

}

// Fenix Server asks Fenix client to send TestDataHeaderHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataHeaderHash(TestDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
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
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bac41696-c8a3-4d11-ac1c-68965c8a1572",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaderHash'")
		}
	}

}

// Fenix Server asks Fenix client to send TestDataHeaders to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataHeaders(TestDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
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
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bac41696-c8a3-4d11-ac1c-68965c8a1572",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaders'")
		}
	}

}

// Fenix Server asks Fenix client to  send TestData rows, based on list of MerklePaths, to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataRows(testDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	serverCopyMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeForServer(testDataClientGuid)
	clientsNewMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeForClient(testDataClientGuid)

	// Extract all paths to retrieve from client
	merklePathsToRetreive := common_config.MissedPathsToRetreiveFromCLient(serverCopyMerkleTree, clientsNewMerkleTree)

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	merklePathsMessage := &fenixClientTestDataSyncServerGrpcApi.MerklePathsMessage{
		MerklePath:                   merklePathsToRetreive,
		ProtoFileVersionUsedByCaller: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendTestDataRows(ctx, merklePathsMessage)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "383be247-e9b8-4b12-bb22-58e8b69d3ab4",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'SendTestDataRows'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "3cbc84af-c2d5-4302-a92a-d11bfd53bdba",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'SendTestDataRows'")
		}
	}

}

// Fenix Server asks Fenix client to  send All TestData rows to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendAllTestDataRows(TestDataClientGuid string) {

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	// Set up connection to Client-server
	fenixTestDataSyncServerObject.SetConnectionToFenixClientTestDataSyncServer()

	emptyParameter := &fenixClientTestDataSyncServerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixClientTestDataSyncServerGrpcApi.CurrentFenixClientTestDataProtoFileVersionEnum(fenixTestDataSyncServerObject.getHighestClientTestDataProtoFileVersion()),
	}

	// Do gRPC-call
	ctx := context.Background()
	returnMessage, err := fenixClientTestDataSyncServerClient.SendAllTestDataRows(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"ID":    "5196e98c-ca19-45f1-a3b8-e4efa95cc312",
			"error": err,
		}).Fatal("Problem to do gRPC-call to FenixClientTestDataSyngocServer for 'AskClientToSendTestDataHeaders'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "7105c16e-cf7d-48ca-8fbc-094d0d5b6f3f",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendTestDataHeaders'")
		}
	}

}
