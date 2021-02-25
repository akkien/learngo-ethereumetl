package job

import (
	"encoding/json"
	"fmt"

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
	// maxWorkers,

) {
	wp := workerpool.New(10)

	// Start pulling
	blockBatchs := util.GeneratePatitions(startBlock, endBlock, batchSize)
	for _, blockBatch := range blockBatchs {
		start, end := blockBatch[0], blockBatch[1]
		wp.Submit(func() {
			blockRange := util.MakeRange(start, end)
			blockReq, err := rpc.GetBlockRequest(blockRange, true)
			if err != nil {
				fmt.Println("Error generate block request", err)
			}
			response := rpc.Call(providerURI, blockReq)

			var blockRes []model.BlockRPCResponse
			err = json.Unmarshal(response, &blockRes)
			if err != nil {
				fmt.Println("Error parse blocks result")
			}

			blocks, txs := model.RPCResponseToBlock(&blockRes)

			blockQuery, blockValues := db.GetInsertParamsBlock(blocks)
			res, err := pg.Exec(blockQuery, blockValues...)
			if err != nil {
				panic(err)
			}
			fmt.Println("Inserted Blocks:", res)

			txQuery, txValues := db.GetInsertParamsTransaction(txs)
			res, err = pg.Exec(txQuery, txValues...)
			if err != nil {
				panic(err)
			}
			fmt.Println("Inserted Transactions:", res)
		})
	}
	wp.StopWait()
}
