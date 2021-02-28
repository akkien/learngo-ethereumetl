package job

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/akkien/ethereumetl/db"
	"github.com/akkien/ethereumetl/model"
	"github.com/akkien/ethereumetl/rpc"
	"github.com/akkien/ethereumetl/util"
	"github.com/gammazero/workerpool"

	"github.com/jmoiron/sqlx"
)

// ParseBlocksAndTransactions parse blocks & transactions
func ParseBlocksAndTransactions(
	startBlock int,
	endBlock int,
	providerURI string,
	pg *sqlx.DB,
	batchSize int,
	maxWorkers int,

) {
	wp := workerpool.New(maxWorkers)
	t1 := time.Now()
	blockBatchs := util.GeneratePatitions(startBlock, endBlock, batchSize)

	for index, blockBatch := range blockBatchs {
		start, end := blockBatch[0], blockBatch[1]
		index := index
		wp.Submit(func() {
			// Export
			blockRange := util.MakeRange(start, end)
			blockReq, err := rpc.GetBlockRequest(blockRange, true)
			if err != nil {
				fmt.Println("Error generate block request", err)
			}
			response, err := rpc.Call(providerURI, blockReq)
			if err != nil {
				fmt.Println("Error call block rpc", err)
			}

			t2 := time.Now()
			fmt.Println(index, "ExportTime", t2.Sub(t1))

			// Transfer
			var blockRes []model.BlockRPCResponse
			err = json.Unmarshal(response, &blockRes)
			if err != nil {
				fmt.Println("Error parse blocks result")
			}

			blocks, txs := model.RPCResponseToBlock(&blockRes)

			t3 := time.Now()
			fmt.Println("transferTime", index, t3.Sub(t1))

			// Load
			blockQuery, blockValues := db.GetInsertParamsBlock(blocks)
			res, err := pg.Exec(blockQuery, blockValues...)
			if err != nil {
				panic(err)
			}
			fmt.Println(index, "Inserted Blocks:", res)

			txQuery, txValues := db.GetInsertParamsTransaction(txs)
			res, err = pg.Exec(txQuery, txValues...)
			if err != nil {
				panic(err)
			}
			fmt.Println(index, "Inserted Transactions:", res)

			time.Sleep(1 * time.Second)
			t5 := time.Now()
			fmt.Println(index, "exportTxTime", t5.Sub(t1))
		})
	}
	wp.StopWait()

}
