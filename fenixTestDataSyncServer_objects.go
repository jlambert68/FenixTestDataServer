package main

import (
	"FenixTestDataServer/common_config"
	fenixClientTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Client/fenixClientTestDataSyncServerGrpcApi/go_grpc_api"
	fenixTestDataSyncServerGrpcApi "github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type fenixTestDataSyncServerObjectStruct struct {
	logger *logrus.Logger
}

var fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct

// Global connection constants
var localServerEngineLocalPort = common_config.FenixTestDataSyncServer_port

var (
	registerfenixTestDataSyncServerServer *grpc.Server
	lis                                   net.Listener
)

var (
	// Standard gRPC Clientr
	remoteFenixClientTestDataSyncServerConnection *grpc.ClientConn
	gRpcClientForFenixClientTestDataSyncServer    fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient

	fenixClientTestDataSyncServer_address_to_dial string = common_config.FenixClientTestDataSyncServer_address + ":" + strconv.Itoa(common_config.FenixClientTestDataSyncServer_initial_port)

	fenixClientTestDataSyncServerClient fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServicesClient
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type FenixTestDataGrpcServicesServer struct {
	fenixTestDataSyncServerGrpcApi.UnimplementedFenixTestDataGrpcServicesServer
}

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

var highestFenixProtoFileVersion int32 = -1
var highestClientProtoFileVersion int32 = -1
