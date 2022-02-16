package main

import (
	"crypto/tls"
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// SetConnectionToFenixClientTestDataSyncServer :
// Set upp connection and Dial to FenixTestDataSyncServer
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) SetConnectionToFenixClientTestDataSyncServer() {

	var err error
	var opts []grpc.DialOption

	//When running on GCP then use credential otherwise not
	if fenixSyncShared.ExecutionLocationForFenixTestDataServer == fenixSyncShared.GCP {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}
	}

	// Set up connection to FenixTestDataSyncServer
	// When run on GCP, use credentials
	if fenixSyncShared.ExecutionLocationForFenixTestDataServer == fenixSyncShared.GCP {
		// Run on GCP
		remoteFenixClientTestDataSyncServerConnection, err = grpc.Dial(fenixclienttestdatasyncserverAddressToDial, opts...)
	} else {
		// Run Local
		remoteFenixClientTestDataSyncServerConnection, err = grpc.Dial(fenixclienttestdatasyncserverAddressToDial, grpc.WithInsecure())
	}

	if err != nil {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixClientTestDataSyncServer_address_to_dial": fenixclienttestdatasyncserverAddressToDial,
			"error message": err,
		}).Error("Did not connect to FenixClientTestDataSyncServer via gRPC")
		//os.Exit(0)
	} else {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"fenixClientTestDataSyncServer_address_to_dial": fenixclienttestdatasyncserverAddressToDial,
		}).Info("gRPC connection OK to FenixClientTestDataSyncServer")

		// Creates a new Clients
		fenixClientTestDataSyncServerClient = fenixClientTestDataSyncServerGrpcApi.NewFenixClientTestDataGrpcServicesClient(remoteFenixClientTestDataSyncServerConnection)

	}
}

// Fenix Server asks Fenix client to register itself with the Fenix Testdata sync server
//func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct)  RegisterTestDataClient(EmptyParameter) returns (AckNackResponse) {
//}

// AskClientToSendMerkleHash :
// Fenix Server asks Fenix client to send TestData MerkleHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendMerkleHash(TestDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "684cea55-6e04-4bee-952f-ffce3c362fcb",
	}).Debug("Incoming gRPC 'AskClientToSendMerkleHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "28020281-ae9d-439c-bc80-d6c227323f7b",
	}).Debug("Outgoing gRPC 'AskClientToSendMerkleHash'")

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

// AskClientToSendMerkleTree :
// Fenix Server asks Fenix client to send TestData MerkleTree to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendMerkleTree(TestDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "be6f045a-f212-4cb6-8f5a-f874a1f79c05",
	}).Debug("Incoming gRPC 'AskClientToSendMerkleTree'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "7a806598-e60d-492f-8a0b-2afd555707b3",
	}).Debug("Outgoing gRPC 'AskClientToSendMerkleTree'")

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
		}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleTree'")

		// FenixTestDataSyncServer couldn't handle gPRC call
		if returnMessage.AckNack == false {
			fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
				"ID": "bba9d885-6dc6-4bd1-9f48-5928e22552ec",
				"Message from FenixClientTestDataSyncServerObject": returnMessage.Comments,
			}).Error("Problem to do gRPC-call to FenixClientTestDataSyncServer for 'AskClientToSendMerkleTree'")
		}
	}

}

// AskClientToSendTestDataHeaderHash :
// Fenix Server asks Fenix client to send TestDataHeaderHash to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataHeaderHash(TestDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "238ece45-409b-41bd-88d6-1ef9edb77edc",
	}).Debug("Incoming gRPC 'AskClientToSendTestDataHeaderHash'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "d3ad936d-6239-4024-8005-e4e198e70026",
	}).Debug("Outgoing gRPC 'AskClientToSendTestDataHeaderHash'")

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

// AskClientToSendTestDataHeaders :
// Fenix Server asks Fenix client to send TestDataHeaders to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataHeaders(TestDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "efc0c455-fb09-467d-8263-d64254a16c79",
	}).Debug("Incoming gRPC 'AskClientToSendTestDataHeaders'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "c7685518-f8f5-4478-9c72-424181cde74e",
	}).Debug("Outgoing gRPC 'AskClientToSendTestDataHeaders'")

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

// AskClientToSendTestDataRows :
// Fenix Server asks Fenix client to  send TestData rows, based on list of MerklePaths, to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendTestDataRows(testDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "12c788ea-29a2-40d0-9807-8cf4685e67b8",
	}).Debug("Incoming gRPC 'AskClientToSendTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "4d25404b-41b2-46e2-b0e9-ab948a675c65",
	}).Debug("Outgoing gRPC 'AskClientToSendTestDataRows'")

	// Check if TestData server should process outgoing messages
	returnMessageStop := fenixTestDataSyncServerObject.isThereATemporaryStopInProcessingInOrOutgoingMessages()
	if returnMessageStop != nil {
		// Temporary stop in processing messages
		return
	}

	serverCopyMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForServer(testDataClientGuid)
	clientsNewMerkleTree := fenixTestDataSyncServerObject.getCurrentMerkleTreeNodesForClient(testDataClientGuid)

	// Extract all paths to retrieve from client
	merklePathsToRetreive := missedPathsToRetrieveFromClient(serverCopyMerkleTree, clientsNewMerkleTree)

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

// AskClientToSendAllTestDataRows :
// Fenix Server asks Fenix client to  send All TestData rows to Fenix Testdata sync server with this service
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) AskClientToSendAllTestDataRows(TestDataClientGuid string) {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "489f78ed-2ab3-4c4c-9436-63e968e1ef8b",
	}).Debug("Incoming gRPC 'AskClientToSendAllTestDataRows'")

	defer fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"id": "001fb9fa-2577-49e2-806f-f8d64f05a3d3",
	}).Debug("Outgoing gRPC 'AskClientToSendAllTestDataRows'")

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
