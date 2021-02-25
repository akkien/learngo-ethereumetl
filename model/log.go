package model

import (
	"fmt"

	"github.com/akkien/ethereumetl/util"
)

//LogRPC : transaction log
type LogRPC struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

// CREATE TABLE logs (
//     log_index SMALLINT,
//     transaction_hash CHAR(66) REFERENCES transactions(hash),
//     transaction_index SMALLINT,
//     block_hash CHAR(66),
//     block_number BIGINT,
//     address CHAR(42),
//     data TEXT,
//     topics TEXT,
//     decoded_value JSON,
//     created_timestamp TIMESTAMPTZ NOT NULL
// 		DEFAULT CURRENT_TIMESTAMP,
// CONSTRAINT pk_logs PRIMARY KEY (transaction_hash, log_index)
// );

//Log : transaction log
type Log struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      int64    `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         int64    `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex int64    `json:"transactionIndex"`
}

// mapTransaction map rpc result to block
func mapLog(in LogRPC) Log {
	out := Log{}
	var err error

	out.Address = in.Address
	out.BlockHash = in.BlockHash
	out.BlockNumber, err = util.HexToDec(in.BlockNumber)
	out.Data = in.Data
	out.LogIndex, err = util.HexToDec(in.LogIndex)
	out.Removed = in.Removed
	out.Topics = in.Topics
	out.TransactionHash = in.TransactionHash
	out.TransactionIndex, err = util.HexToDec(in.TransactionIndex)

	if err != nil {
		fmt.Println("Map log", err)
	}

	return out
}