// Package job - Main process: pulling, processing, and saving the blockchain data
package job

import (
	"fmt"

	"github.com/akkien/ethereumetl/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ExportAll : Export all data from block
func ExportAll(
	startBlock int, endBlock int, paritionBatchSize int, batchSize int,
	providerURI string,
	connStr string,
	maxWorkers int,
) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxWorkers + 1)

	fmt.Println("Database Connected")

	/** Parse Blocks & Transactions **/
	blockPartitions := util.GeneratePatitions(startBlock, endBlock, paritionBatchSize)
	var partitionTxsHash []string

	for _, partition := range blockPartitions {
		fmt.Println("Partition", partition)
		txsHashChan := ParseBlocksAndTransactions(partition[0], partition[1], providerURI, db, batchSize, maxWorkers)

		for txsHash := range txsHashChan {
			/** Parse Receipts & Logs **/
			partitionTxsHash = append(partitionTxsHash, txsHash...)
		}
	}

	ParseReceiptsAndLogs(partitionTxsHash, providerURI, db, batchSize, maxWorkers)
}
