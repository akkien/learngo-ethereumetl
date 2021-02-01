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
func ParseBlocksAndTransactions(blocks []uint64) {
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
		response := rpc.Call(blockReq)

		var blockRes []rpc.BlockRPCResponse
		err = json.Unmarshal(response, &blockRes)
		if err != nil {
			fmt.Println("Error parse result")
		}

		blockRPC := blockRes[0].Result
		block := model.MapBlock(blockRPC)
		fmt.Println(block.Number)

		nextBlock := block
		nextBlock.Number++

		blockList := []model.Block{block, nextBlock}

		query, values := db.GetInsertParamsBlock(blockList)
		res, err := pg.Exec(query, values...)
		if err != nil {
			panic(err)
		}
		fmt.Println("RES:", res)
	}
}
