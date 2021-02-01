package model

import (
	"github.com/akkien/ethereumetl/util"
)

// BlockRPC RPC response
type BlockRPC struct {
	Difficulty       string            `json:"difficulty"`
	ExtraData        string            `json:"extraData"`
	GasLimit         string            `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	Hash             string            `json:"hash"`
	LogsBloom        string            `json:"logsBloom"`
	Miner            string            `json:"miner"`
	MixHash          string            `json:"mixHash"`
	Nonce            string            `json:"nonce"`
	Number           string            `json:"number"`
	ParentHash       string            `json:"parentHash"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Sha3Uncles       string            `json:"sha3Uncles"`
	Size             string            `json:"size"`
	StateRoot        string            `json:"stateRoot"`
	Timestamp        string            `json:"timestamp"`
	TotalDifficulty  string            `json:"totalDifficulty"`
	Transactions     []*TransactionRPC `json:"transactions"`
	TransactionsRoot string            `json:"transactionsRoot"`
	Uncles           []interface{}     `json:"uncles"`
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

// MapBlock map rpc result to block
func MapBlock(in BlockRPC) (out Block) {
	out.ExtraData = in.ExtraData
	out.Difficulty = util.HexToDec(in.Difficulty)
	out.ExtraData = in.ExtraData
	out.GasLimit = util.HexToDec(in.GasLimit)
	out.GasUsed = util.HexToDec(in.GasUsed)
	out.Hash = in.Hash
	out.LogsBloom = in.LogsBloom
	out.Miner = in.Miner
	out.MixHash = in.MixHash
	out.Nonce = in.Number
	out.Number = util.HexToDec(in.Number)
	out.ParentHash = in.ParentHash
	out.ReceiptsRoot = in.ReceiptsRoot
	out.Sha3Uncles = in.Sha3Uncles
	out.Size = util.HexToDec(in.Size)
	out.StateRoot = in.StateRoot
	out.Timestamp = util.HexToDec(in.Timestamp)
	out.TotalDifficulty = util.HexToDec(in.TotalDifficulty)
	out.TransactionsRoot = in.TransactionsRoot
	out.TransactionCount = len(in.Transactions)
	out.CreatedTimestamp = ""

	return
}
