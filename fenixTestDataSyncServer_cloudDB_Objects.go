package main

/* START Cloud Sync
./cloud_sql_proxy -instances=mycloud-run-project:europe-north1:fenix-sqlserver=tcp:5432

*/

// All TestTDataClients in CloudDB
var cloudDBClients []cloudDBTestDataClientStruct
var cloudDBClientsMap cloudDBClientsMapType

type cloudDBTestDataClientStruct struct {
	clientUuid  string
	clientName  string
	domainUuid  string
	description string
}

type cloudDBClientsMapType map[memDBClientUuidType]cloudDBTestDataClientMapStruct
type cloudDBTestDataClientMapStruct struct {
	clientName  string
	domainUuid  string
	description string
}

// All TestDataHeaderFilterValues in CloudDB
var cloudDBTestDataHeaderItemsHashes []cloudDBTestDataHeaderItemsHashesStruct

type cloudDBTestDataHeaderItemsHashesStruct struct {
	headerItemsHash  string
	clientUuid       string
	headerLabelsHash string
	updatedTimeStamp string
}

// All TestDataHeaderFilterValues in CloudDB
var cloudDBTestDataHeadersFilterValues []cloudDBTestDataHeaderFilterValuesStruct

type cloudDBTestDataHeaderFilterValuesStruct struct {
	headerItemHash         string
	headerFilterValueOrder int
	headerFilterValue      string
	clientUuid             string
	headerFilterValuesHash string
	updatedTimeStamp       string
}

// All TestDataHeaderItems in CloudDB
var cloudDBTestDataHeaderItems []cloudDBTestDataHeaderItemStruct

type cloudDBTestDataHeaderItemStruct struct {
	headerItemsHash      string
	headerItemHash       string
	clientUuid           string
	headerLabel          string
	shouldBeUsedInFilter bool
	isMandatoryInFilter  bool
	filterSelectionType  int
	headerColumnOrder    int
	updatedTimeStamp     string
}

// All TestDataMerkleHashes in CloudDB
var cloudDBTestDataMerkleHashes []cloudDBTestDataMerkleHashStruct

type cloudDBTestDataMerkleHashStruct struct {
	clientUuid           string
	merkleHash           string
	merkleFilterPath     string
	merkleFilterPathHash string
	updatedTimeStamp     string
}

// All TestDataMerkleTrees in CloudDB
var cloudDBTestDataMerkleTrees []cloudDBTestDataMerkleTreeStruct

type cloudDBTestDataMerkleTreeStruct struct {
	clientUuid       string
	merkleHash       string
	nodeLevel        int
	nodeName         string
	nodePath         string
	nodeHash         string
	nodeChildHash    string
	updatedTimeStamp string
}

// All TestDataRowItems in CloudDB
var cloudDBTestDataRowItems []cloudDBTestDataRowItemStruct

type cloudDBTestDataRowItemStruct struct {
	clientUuid            string
	rowHash               string
	testdataValueAsString string
	leafNodeName          string
	leafNodePath          string
	leafNodeHash          string
	valueColumnOrder      int
	valueRowOrder         int
	updatedTimeStamp      string
}
