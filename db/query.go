package db

import (
	"strconv"
	"strings"

	"github.com/akkien/ethereumetl/model"
)

var blockParams = []string{
	"difficulty",
	"extra_data",
	"gas_limit",
	"gas_used",
	"hash",
	"logs_bloom",
	"miner",
	"mix_hash",
	"nonce",
	"number",
	"parent_hash",
	"receipts_root",
	"sha3_uncles",
	"size",
	"state_root",
	"timestamp",
	"total_difficulty",
	"transactions_root",
	"transaction_count",
}

var transactionParams = []string{
	"block_hash",
	"block_number",
	"from_address",
	"gas",
	"gas_price",
	"hash",
	"input",
	"nonce",
	"r",
	"s",
	"to_address",
	"transaction_index",
	"v",
	"value",
}

// Create create table
func Create(table string) string {
	createBlock := `
	CREATE TABLE blocks (
		difficulty TEXT,
		extra_data TEXT,
		gas_limit BIGINT,
		gas_used BIGINT,
		hash CHAR(66),
		logs_bloom TEXT,
		miner CHAR(42),
		mix_hash CHAR(66),
		nonce TEXT,
		number BIGINT PRIMARY KEY,
		parent_hash CHAR(66),
		receipts_root CHAR(66),
		sha3_uncles CHAR(66),
		size BIGINT,
		state_root CHAR(66),
		timestamp BIGINT,
		total_difficulty TEXT,
		transactions_root CHAR(66),
		transaction_count SMALLINT,
		created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
	);`

	createTransaction := `
	CREATE TABLE transactions (
    	block_hash CHAR(66),
    	block_number BIGINT REFERENCES blocks(number),
    	from_address CHAR(42),
    	gas BIGINT,
		gas_price BIGINT,
		hash CHAR(66) PRIMARY KEY,
		input TEXT,
		nonce BIGINT,
		r TEXT,
		s TEXT,
		to_address CHAR(42),
		transaction_index SMALLINT,
		v TEXT,
		value DECIMAL(38,0),
    	created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
	);
	`
	switch table {
	case "blocks":
		return createBlock
	case "transactions":
		return createTransaction
	default:
		return ""
	}
}

// getInsertPlaceholder get table insert placeholder
func getInsertPlaceholder(noOfFields int) string {
	ret := make([]string, noOfFields)
	for i := 0; i < noOfFields; i++ {
		ret[i] = "$" + strconv.Itoa(i)
	}
	return strings.Join(ret, ",")
}

// GetInsertParamsBlock new record to database
func GetInsertParamsBlock(blocks []model.Block) (query string, values []interface{}) {
	numFields := len(blockParams)
	query = `INSERT INTO blocks (` + strings.Join(blockParams, ",") + `) VALUES `
	values = []interface{}{}
	for i, block := range blocks {
		values = append(values,
			block.Difficulty,
			block.ExtraData,
			block.GasLimit,
			block.GasUsed,
			block.Hash,
			block.LogsBloom,
			block.Miner,
			block.MixHash,
			block.Nonce,
			block.Number,
			block.ParentHash,
			block.ReceiptsRoot,
			block.Sha3Uncles,
			block.Size,
			block.StateRoot,
			block.Timestamp,
			block.TotalDifficulty,
			block.TransactionsRoot,
			block.TransactionCount,
		)
		n := i * numFields
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1] // remove the trailing comma
	return
}

// GetInsertParamsTransaction new record to database
func GetInsertParamsTransaction(blocks []model.Block) (query string, values []interface{}) {
	numFields := len(blockParams)
	query = `INSERT INTO transactions (` + strings.Join(blockParams, ",") + `) VALUES `
	values = []interface{}{}
	for i, block := range blocks {
		values = append(values,
			block.Difficulty,
			block.ExtraData,
			block.GasLimit,
			block.GasUsed,
			block.Hash,
			block.LogsBloom,
			block.Miner,
			block.MixHash,
			block.Nonce,
			block.Number,
			block.ParentHash,
			block.ReceiptsRoot,
			block.Sha3Uncles,
			block.Size,
			block.StateRoot,
			block.Timestamp,
			block.TotalDifficulty,
			block.TransactionsRoot,
			block.TransactionCount,
		)
		n := i * numFields
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1] // remove the trailing comma
	return
}
