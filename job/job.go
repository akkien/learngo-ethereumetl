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
	} else {
		if err = db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("Connected")
		db.SetMaxOpenConns(maxWorkers + 1)

		blockPartitions := util.GeneratePatitions(startBlock, endBlock, paritionBatchSize)
		for _, partition := range blockPartitions {
			fmt.Println("Partition", partition)
			ParseBlocksAndTransactions(partition[0], partition[1], providerURI, db, batchSize)
		}
	}
}
