package rpc

import (
	"encoding/json"
	"strconv"

	"github.com/akkien/ethereumetl/model"
)

// GetBlockRequest generate get block rpc request
func GetBlockRequest(blockNumbers []int, isIncludeTxs bool) ([]byte, error) {
	rpcs := make([]model.RPC, len(blockNumbers))

	for idx, block := range blockNumbers {
		blockHex := "0x" + strconv.FormatInt(int64(block), 16)
		rpc := model.RPC{
			Jsonrpc: "2.0",
			Method:  "eth_getBlockByNumber",
			Params:  []interface{}{blockHex, isIncludeTxs},
			ID:      idx,
		}
		rpcs[idx] = rpc
	}

	data, err := json.Marshal(&rpcs)
	if err != nil {
		return nil, err
	}
	return data, nil
}
