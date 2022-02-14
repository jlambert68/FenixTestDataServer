package main

import (
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcAdminApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_admin_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

// LoggingLevel
// Logrus debug level
//const LoggingLevel = logrus.DebugLevel
//const LoggingLevel = logrus.InfoLevel
const LoggingLevel = logrus.DebugLevel // InfoLevel

type fenixTestDataSyncServerObjectStruct struct {
	logger                                 *logrus.Logger
	stateProcessIncomingAndOutgoingMessage bool
}

var fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct

// Address to Fenix TestData Server & Client, will have their values from Environment variables at startup
var (
	fenixTestDataSyncServerAddress string
	fenixTestDataSyncServerPort    int
	//fenixTestDataSyncServerAdminAddress string
	fenixTestDataSyncServerAdminPort int
	clientTestDataSyncServerAddress  string
	clientTestDataSyncServerPort     int
)

// Global connection constants
//var localServerEngineLocalPort = fenixTestDataSyncServerPort
//var localServerEngineLocalAdminPort int

var (
	registerfenixTestDataSyncServerServer      *grpc.Server
	registerfenixTestDataSyncAdminServerServer *grpc.Server
	gRPClis                                    net.Listener
	gRPCAdminLis                               net.Listener
)

var (
	// Standard gRPC Clientr for calling Fenix Client
	remoteFenixClientTestDataSyncServerConnection *grpc.ClientConn
	gRpcClientForFenixClientTestDataSyncServer    fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient

	fenixclienttestdatasyncserverAddressToDial string

	fenixClientTestDataSyncServerClient fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient
)

// FenixTestDataGrpcServicesServer :
// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixTestDataGrpcServicesServer struct {
	fenixTestDataSyncServerGrpcApi.UnimplementedFenixTestDataGrpcServicesServer
}

// FenixTestDataGrpcServicesAdminServer :
// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixTestDataGrpcServicesAdminServer struct {
	fenixTestDataSyncServerGrpcAdminApi.UnimplementedFenixTestDataGrpcAdminServicesServer
}

/*
// Channels which takes incoming gRPC-messages and pass them to 'message process engine'
// 'TestDataClientInformationMessage' from 'gRPC-RegisterTestDataClient'
var TestDataClientInformationMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataClientInformationMessage

// 'MerkleHashMessage' from 'gRPC-SendMerkleHash'
var MerkleHashMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleHashMessage

// 'MerkleTreeMessage' from 'gRPC-SendMerkleTree'
var MerkleTreeMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage

// 'TestDataHeaderMessage' from 'gRPC-SendTestDataHeaders'
var TestDataHeaderMessageChannel chan fenixTestDataSyncServerGrpcApi.TestDataHeadersMessage

// 'MerkleTreeMessage' from 'gRPC-SendTestDataRows'
var MerkleTreeMessageMessageChannel chan fenixTestDataSyncServerGrpcApi.MerkleTreeMessage
*/
var highestFenixProtoFileVersion int32 = -1
var highestClientProtoFileVersion int32 = -1
