// Package job - Main process: pulling, processing, and saving the blockchain data
package job

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ExportAll : Export all data from block
func ExportAll(
	//startBlock int, endBlock int, paritionBatchSize int, providerURI string, batchSize int,

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
		db.SetMaxOpenConns(maxWorkers)

		ParseBlocksAndTransactions([]uint64{11775341, 11775342})

	}
}
