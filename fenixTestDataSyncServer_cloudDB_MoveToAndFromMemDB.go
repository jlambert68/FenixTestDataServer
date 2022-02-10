package main

/*
import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/sirupsen/logrus"
)


// Load all CloudDB-data that should be held in MemoryDB
func (fenixTestDataSyncServerObject *fenixTestDataSyncServerObjectStruct) loadCloudDBIntoMemoryDB() error {

	fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
		"Id": "42648bb9-e953-4e13-9855-51d7401fa291",
	}).Debug("Entering: loadCloudDBIntoMemoryDB()")

	defer func() {
		fenixTestDataSyncServerObject.logger.WithFields(logrus.Fields{
			"Id": "87694618-5cba-411b-8d8a-cdec9c16f78e",
		}).Debug("Exiting: loadCloudDBIntoMemoryDB()")
	}()

	// Temporary memObject object
	var tempMemoryDB memoryDBStruct

	// Load Clients into MemoryDB
	allowedClients, err := fenixTestDataSyncServerObject.loadClientsFromCloudDB()
	if err != nil {
		return err
	}

	tempMemoryDB.allowedClients.memDBTestDataDomainType = allowedClients.memDBTestDataDomainType

	var (
		headerItems []memDBHeaderItemsStruct

		headerFilterValues []string
	)

	// Load Servers TestData content, that was previously sent by the clients

	// Loop over all Client&Domain combinations and retrieve Data from CloudDB
	for _, allowedClient := range cloudDBClients {
		fmt.Println(allowedClient.clientUuid, allowedClient.domainUuid)

		// Add all HeaderItems
		headerFilterValues = []string{}

		headerItem := memDBHeaderItemsStruct{
			headerItemHash:             "",
			headerLabel:                "",
			headerShouldBeUsedInFilter: false,
			headerIsMandatoryInFilter:  false,
			headerFilterSelectionType:  0,
			headerFilterValuesItem: memDBHeaderFilterValuesItemStruct{
				HeaderFilterValuesHash: "",
				HeaderFilterValues:     headerFilterValues,
			},
		}

		headerItems = []memDBHeaderItemsStruct{headerItem}

		// Add all MerkleTreeRows
		merkleTreeRow := memDBMerkleTreeRowStruct{}

		merkleTreeRows := []memDBMerkleTreeRowStruct{merkleTreeRow}

		// Build TestData object
		memDBTestDataStruct := memDBTestDataStruct{memDBDataStructureStruct{
			merkleHash: "",
			merkleFilterPath: "",
			merkleTree: memDBMerkleTreeRowsStruct{
				merkleTreeRows: merkleTreeRows,
			},
			headerItemsHash:  "",
			headerLabelsHash: "",
			headerItems:      headerItems,
			testDataRows:     []memDBTestDataItemsStruct{},
			testDataAsDataFrame: dataframe.DataFrame{
				Err: nil,
			},
		},
		}

		// Initiate and add TestData object to correct Domain and Client
		tempMemoryDB.server.memDBTestDataDomain[allowedClient.domainUuid][allowedClient.clientUuid] = memDBTestDataStruct

	}

		// Add all testdata rows
		var testDataRows []memDBTestDataItemsStruct
		for _, memCloudDBAllTestDataRowItem := range cloudDBTestDataRowItems {
				tempTestDataItems:= *tempMemoryDB.server.memDBTestDataDomain[memCloudDBAllTestDataRowItem.domainUuid][memCloudDBAllTestDataRowItem.clientUuid].
				memDBDataStructure.testDataRows
				if len(tempTestDataRows) == 0 {
					//No rows added, so create the first one
					testDataRow := memDBTestDataItemsStruct{
						testDataRowHash:        memCloudDBAllTestDataRowItem.rowHash,
						leafNodeName:           memCloudDBAllTestDataRowItem.leafNodeName,
						leafNodePath:           memCloudDBAllTestDataRowItem.leafNodePath,
						testDataValuesAsString: []string{},
					}

				}
			testDataRow := memDBTestDataItemsStruct{
				testDataRowHash:        memCloudDBAllTestDataRowItem.rowHash,
				leafNodeName:           memCloudDBAllTestDataRowItem.leafNodeName,
				leafNodePath:           memCloudDBAllTestDataRowItem.leafNodePath,
				testDataValuesAsString: append(memCloudDBAllTestDataRowItem.testdataValueAsString,
			}
			testDataRows = append(testDataRows, testDataRow)
		}

		testDataRows, err := fenixTestDataSyncServerObject.mapTestDataRowsFromCloudDBToMemDB(allowedClient, allowedDomain)
		if err != nil {
			return err
		}

		// Add all HeaderItems
		var headerFilterValue string

		headerFilterValues = []string{headerFilterValue}

		headerItem := memDBHeaderItemsStruct{
			headerItemHash:             "",
			headerLabel:                "",
			headerShouldBeUsedInFilter: false,
			headerIsMandatoryInFilter:  false,
			headerFilterSelectionType:  0,
			headerFilterValuesItem: memDBHeaderFilterValuesItemStruct{
				HeaderFilterValuesHash: "",
				HeaderFilterValues:     headerFilterValues,
			},
		}

		headerItems = []memDBHeaderItemsStruct{headerItem}

		// Add all MerkleTreeRows
		merkleTreeRow := memDBMerkleTreeRowStruct{
			nodeLevel:     "",
			nodeName:      "",
			nodePath:      "",
			nodeHash:      "",
			nodeChildHash: "",
		}

		merkleTreeRows := []memDBMerkleTreeRowStruct{merkleTreeRow}

		// Build TestData object
		memDBTestDataStruct := memDBTestDataStruct{memDBDataStructureStruct{
			merkleHash: "",
			merkleFilterPath: "",
			merkleTree: memDBMerkleTreeRowsStruct{
				merkleTreeRows: merkleTreeRows,
			},
			headerItemsHash:  "",
			headerLabelsHash: "",
			headerItems:      headerItems,
			testDataRows:     testDataRows,
			testDataAsDataFrame: dataframe.DataFrame{
				Err: nil,
			},
		},
		}

		// Add TestData object to correct Domain and Client
		tempMemoryDB.server.memDBTestDataDomain["domainID"]["ClientID"] = memDBTestDataStruct

	}

	/*

		memoryDB.allowedClients = tempMemoryDB.allowedClients
		memoryDB.


*/
/*
	return nil
}



*/
