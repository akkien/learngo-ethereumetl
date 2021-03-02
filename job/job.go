// Package job - Main process: pulling, processing, and saving the blockchain data
package job

import (
	"fmt"
	"strconv"

	"github.com/akkien/ethereumetl/util"
	"github.com/go-redis/redis"
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

	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisCli.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	/** Parse Blocks & Transactions **/
	blockPartitions := util.GeneratePatitions(startBlock, endBlock, paritionBatchSize)
	partitionTxsKey := strconv.Itoa(startBlock) + "_" + strconv.Itoa(endBlock) + "-TXS"
	for _, partition := range blockPartitions {
		fmt.Println("Partition", partition)
		ParseBlocksAndTransactions(partition[0], partition[1], providerURI, db, batchSize, maxWorkers, redisCli, partitionTxsKey)
	}

	/** Parse Receipts & Logs **/
	txsHash, err := redisCli.LRange(partitionTxsKey, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	redisCli.Del(partitionTxsKey)
	ParseReceiptsAndLogs(txsHash, providerURI, db, batchSize, maxWorkers)

}
