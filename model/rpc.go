package model

// ReceiptRPCResponse RPC response
type ReceiptRPCResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  ReceiptRPC `json:"result"`
}

// BlockRPCResponse RPC response
type BlockRPCResponse struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Result  BlockRPC `json:"result"`
}

// BlockNumberRPCResponse RPC response
type BlockNumberRPCResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

// RPC rpc call
type RPC struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}
