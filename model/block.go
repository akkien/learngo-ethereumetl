package model

import (
	"fmt"

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
	Difficulty       int64  `json:"difficulty" db:"difficulty"`
	ExtraData        string `json:"extraData" db:"extra_data"`
	GasLimit         int64  `json:"gasLimit" db:"gas_limit"`
	GasUsed          int64  `json:"gasUsed" db:"gas_used"`
	Hash             string `json:"hash" db:"hash"`
	LogsBloom        string `json:"logsBloom" db:"logs_bloom"`
	Miner            string `json:"miner" db:"miner"`
	MixHash          string `json:"mixHash" db:"mix_hash"`
	Nonce            string `json:"nonce" db:"nonce"`
	Number           int64  `json:"number" db:"number"`
	ParentHash       string `json:"parentHash" db:"parent_hash"`
	ReceiptsRoot     string `json:"receiptsRoot" db:"receipts_root"`
	Sha3Uncles       string `json:"sha3Uncles" db:"sha3_uncles"`
	Size             int64  `json:"size" db:"size"`
	StateRoot        string `json:"stateRoot" db:"state_root"`
	Timestamp        int64  `json:"timestamp" db:"timestamp"`
	TotalDifficulty  int64  `json:"totalDifficulty" db:"total_difficulty"`
	TransactionsRoot string `json:"transactionsRoot" db:"transactions_root"`
	TransactionCount int64  `json:"transactionCount" db:"transaction_count"`
	CreatedTimestamp string `json:"createdTimestamp" db:"created_timestamp"`
}

func mapBlock(rpcBlock BlockRPC) Block {
	out := Block{}
	var err error
	out.ExtraData = rpcBlock.ExtraData
	out.Difficulty, err = util.HexToDec(rpcBlock.Difficulty)
	if err != nil {
		fmt.Println("block Difficulty", err)
	}
	out.ExtraData = rpcBlock.ExtraData
	out.GasLimit, err = util.HexToDec(rpcBlock.GasLimit)
	out.GasUsed, err = util.HexToDec(rpcBlock.GasUsed)
	out.Hash = rpcBlock.Hash
	out.LogsBloom = rpcBlock.LogsBloom
	out.Miner = rpcBlock.Miner
	out.MixHash = rpcBlock.MixHash
	out.Nonce = rpcBlock.Nonce
	out.Number, err = util.HexToDec(rpcBlock.Number)
	out.ParentHash = rpcBlock.ParentHash
	out.ReceiptsRoot = rpcBlock.ReceiptsRoot
	out.Sha3Uncles = rpcBlock.Sha3Uncles
	out.Size, err = util.HexToDec(rpcBlock.Size)
	out.StateRoot = rpcBlock.StateRoot
	out.Timestamp, err = util.HexToDec(rpcBlock.Timestamp)
	out.TotalDifficulty, err = util.HexToDec(rpcBlock.TotalDifficulty)
	out.TransactionsRoot = rpcBlock.TransactionsRoot
	out.TransactionCount = int64(len(rpcBlock.Transactions))
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
