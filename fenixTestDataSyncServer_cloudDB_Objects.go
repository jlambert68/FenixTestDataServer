package main

/* START Cloud Sync
./cloud_sql_proxy -instances=mycloud-run-project:europe-north1:fenix-sqlserver=tcp:5432

*/

// All TestTDataClients in CloudDB
var memCloudDBAllClients []memCloudDBAllTestDataClientStruct
var memCloudDBAllClientsMap memCloudDBAllClientsMapType

type memCloudDBAllTestDataClientStruct struct {
	clientUuid  memDBClientUuidType
	clientName  string
	domainUuid  memDBDomainUuidType
	description string
}

type memCloudDBAllClientsMapType map[memDBClientUuidType]memCloudDBAllTestDataClientMapStruct
type memCloudDBAllTestDataClientMapStruct struct {
	clientName  string
	domainUuid  memDBDomainUuidType
	description string
}

// All TestDataHeaderFilterValues in CloudDB
var memCloudDBAllTestDataHeaderFilterValues []memCloudDBAllTestDataHeaderFilterValueStruct

type memCloudDBAllTestDataHeaderFilterValueStruct struct {
	headerItemHash    string
	headerFilterValue string
	clientUuid        memDBClientUuidType
	domainUuid        memDBDomainUuidType
}

// All TestDataHeaderItems in CloudDB
var memCloudDBAllTestDataHeaderItems []memCloudDBAllTestDataHeaderItemStruct

type memCloudDBAllTestDataHeaderItemStruct struct {
	clientUuid           memDBClientUuidType
	domainUuid           memDBDomainUuidType
	headerItemHash       string
	headerLabel          string
	shouldBeUsedInFilter bool
	isMandatoryInFilter  bool
	filterSelection_type int
	filterValuesHash     string
}

// All TestDataMerkleHashes in CloudDB
var memCloudDBAllTestDataMerkleHashes []memCloudDBAllTestDataMerkleHashStruct

type memCloudDBAllTestDataMerkleHashStruct struct {
	clientUuid memDBClientUuidType
	domainUuid memDBDomainUuidType
	merkleHash string
	merklePath string
}

// All TestDataMerkleTrees in CloudDB
var memCloudDBAllTestDataMerkleTrees []memCloudDBAllTestDataMerkleTreeStruct

type memCloudDBAllTestDataMerkleTreeStruct struct {
	clientUuid    memDBClientUuidType
	domainUuid    memDBDomainUuidType
	nodeLevel     int
	nodeName      string
	nodePath      string
	nodeHash      string
	nodeChildHash string
}

// All TestDataRowItems in CloudDB
var memCloudDBAllTestDataRowItems []memCloudDBAllTestDataRowItemStruct

type memCloudDBAllTestDataRowItemStruct struct {
	clientUuid            memDBClientUuidType
	domainUuid            memDBDomainUuidType
	rowHash               string
	testdataValueAsString string
	leafNodeName          string
	leafNodePath          string
}
