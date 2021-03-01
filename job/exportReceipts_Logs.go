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

// ParseReceiptsAndLogs parse blocks & transactions
func ParseReceiptsAndLogs(
	txsHash []string,
	providerURI string,
	pg *sqlx.DB,
	batchSize int,
	maxWorkers int,
) {
	wp := workerpool.New(maxWorkers)
	t1 := time.Now()
	txBatchs := util.GenerateBatchs(txsHash, batchSize)

	for index, txsBatch := range txBatchs {
		index := index

		wp.Submit(func() {
			// Export
			blockReq, err := rpc.GetReceiptRequest(txsBatch)
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
			var receiptRes []model.ReceiptRPCResponse
			err = json.Unmarshal(response, &receiptRes)
			if err != nil {
				fmt.Println("Error parse blocks result")
			}
			receipts, logs := model.RPCResponseToReceipt(&receiptRes)

			// Load
			receiptQuery, receiptValues := db.GetInsertParamsReceipt(receipts)
			res, err := pg.Exec(receiptQuery, receiptValues...)
			if err != nil {
				fmt.Println("Error insert receipts", err)
			}
			fmt.Println(index, "Inserted Receipts:", res)

			if len(logs) > 0 {
				logQuery, logValues := db.GetInsertParamsLog(logs)
				res, err = pg.Exec(logQuery, logValues...)
				if err != nil {
					fmt.Println("Error insert logs", err)
				}
				fmt.Println(index, "Inserted Logs:", res)
			}
		})
	}
	wp.StopWait()

}
