package job

import (
	"encoding/json"
	"fmt"

	"github.com/akkien/ethereumetl/db"
	"github.com/akkien/ethereumetl/model"
	"github.com/akkien/ethereumetl/rpc"

	"github.com/jmoiron/sqlx"
)

// ParseBlocksAndTransactions parse blocks & transactions
func ParseBlocksAndTransactions(blocks []uint64,

// startBlock,
// endBlock,
// batchSize,
// batchWeb3Provider,
// maxWorkers,
// itemExporter
) {
	RopstenHTTP := "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	//RopstenHTTP := "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	connStr := "postgres://akkien:trungkien@127.0.0.1:5432/ropsten?sslmode=disable"
	pg, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} else {
		if err = pg.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("Connected")

		// Start pulling
		blockReq, err := rpc.GetBlockRequest(blocks, true)
		if err != nil {
			fmt.Println("Error generate block request")
		}
		response := rpc.Call(RopstenHTTP, blockReq)

		var blockRes []model.BlockRPCResponse
		err = json.Unmarshal(response, &blockRes)
		if err != nil {
			fmt.Println("Error parse result")
		}

		blocks, txs := model.RPCResponseToBlock(&blockRes)

		blockQuery, blockValues := db.GetInsertParamsBlock(blocks)
		res, err := pg.Exec(blockQuery, blockValues...)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted Block:", res)

		txQuery, txValues := db.GetInsertParamsTransaction(txs)
		res, err = pg.Exec(txQuery, txValues...)
		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted Transactions:", res)
	}
}
