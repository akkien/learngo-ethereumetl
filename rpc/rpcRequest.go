package rpc

import (
	"encoding/json"
	"strconv"

	"github.com/akkien/ethereumetl/model"
)

// ReceiptRPCResponse RPC response
type ReceiptRPCResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  model.ReceiptRPC `json:"result"`
}

// BlockRPCResponse RPC response
type BlockRPCResponse struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  model.BlockRPC `json:"result"`
}

// RPC rpc call
type RPC struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// GetBlockRequest generate get block rpc request
func GetBlockRequest(blockNumbers []uint64, isIncludeTxs bool) ([]byte, error) {
	rpcs := make([]RPC, len(blockNumbers))

	for idx, block := range blockNumbers {
		blockHex := "0x" + strconv.FormatUint(block, 16)
		rpc := RPC{
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
