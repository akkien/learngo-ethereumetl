package model

import (
	"github.com/akkien/ethereumetl/util"
)

// BlockRPC RPC response
type BlockRPC struct {
	Difficulty       string           `json:"difficulty"`
	ExtraData        string           `json:"extraData"`
	GasLimit         string           `json:"gasLimit"`
	GasUsed          string           `json:"gasUsed"`
	Hash             string           `json:"hash"`
	LogsBloom        string           `json:"logsBloom"`
	Miner            string           `json:"miner"`
	MixHash          string           `json:"mixHash"`
	Nonce            string           `json:"nonce"`
	Number           string           `json:"number"`
	ParentHash       string           `json:"parentHash"`
	ReceiptsRoot     string           `json:"receiptsRoot"`
	Sha3Uncles       string           `json:"sha3Uncles"`
	Size             string           `json:"size"`
	StateRoot        string           `json:"stateRoot"`
	Timestamp        string           `json:"timestamp"`
	TotalDifficulty  string           `json:"totalDifficulty"`
	Transactions     []TransactionRPC `json:"transactions"`
	TransactionsRoot string           `json:"transactionsRoot"`
	Uncles           []interface{}    `json:"uncles"`
}

// Block for PostgreSQL
type Block struct {
	Difficulty       int    `json:"difficulty" db:"difficulty"`
	ExtraData        string `json:"extraData" db:"extra_data"`
	GasLimit         int    `json:"gasLimit" db:"gas_limit"`
	GasUsed          int    `json:"gasUsed" db:"gas_used"`
	Hash             string `json:"hash" db:"hash"`
	LogsBloom        string `json:"logsBloom" db:"logs_bloom"`
	Miner            string `json:"miner" db:"miner"`
	MixHash          string `json:"mixHash" db:"mix_hash"`
	Nonce            string `json:"nonce" db:"nonce"`
	Number           int    `json:"number" db:"number"`
	ParentHash       string `json:"parentHash" db:"parent_hash"`
	ReceiptsRoot     string `json:"receiptsRoot" db:"receipts_root"`
	Sha3Uncles       string `json:"sha3Uncles" db:"sha3_uncles"`
	Size             int    `json:"size" db:"size"`
	StateRoot        string `json:"stateRoot" db:"state_root"`
	Timestamp        int    `json:"timestamp" db:"timestamp"`
	TotalDifficulty  int    `json:"totalDifficulty" db:"total_difficulty"`
	TransactionsRoot string `json:"transactionsRoot" db:"transactions_root"`
	TransactionCount int    `json:"transactionCount" db:"transaction_count"`
	CreatedTimestamp string `json:"createdTimestamp" db:"created_timestamp"`
}

func mapBlock(rpcBlock BlockRPC) Block {
	out := Block{}

	out.ExtraData = rpcBlock.ExtraData
	out.Difficulty = util.HexToDec(rpcBlock.Difficulty)
	out.ExtraData = rpcBlock.ExtraData
	out.GasLimit = util.HexToDec(rpcBlock.GasLimit)
	out.GasUsed = util.HexToDec(rpcBlock.GasUsed)
	out.Hash = rpcBlock.Hash
	out.LogsBloom = rpcBlock.LogsBloom
	out.Miner = rpcBlock.Miner
	out.MixHash = rpcBlock.MixHash
	out.Nonce = rpcBlock.Nonce
	out.Number = util.HexToDec(rpcBlock.Number)
	out.ParentHash = rpcBlock.ParentHash
	out.ReceiptsRoot = rpcBlock.ReceiptsRoot
	out.Sha3Uncles = rpcBlock.Sha3Uncles
	out.Size = util.HexToDec(rpcBlock.Size)
	out.StateRoot = rpcBlock.StateRoot
	out.Timestamp = util.HexToDec(rpcBlock.Timestamp)
	out.TotalDifficulty = util.HexToDec(rpcBlock.TotalDifficulty)
	out.TransactionsRoot = rpcBlock.TransactionsRoot
	out.TransactionCount = len(rpcBlock.Transactions)
	out.CreatedTimestamp = ""

	return out
}

// RPCResponseToBlock map rpc result to block
func RPCResponseToBlock(response *[]BlockRPCResponse) ([]Block, []Transaction) {
	blocks := make([]Block, len((*response)))
	txs := []Transaction{}
	for i, blockRes := range *response {
		blocks[i] = mapBlock(blockRes.Result)
		for _, rpcTx := range blockRes.Result.Transactions {
			txs = append(txs, mapTransaction(rpcTx))
		}
	}

	return blocks, txs
}
