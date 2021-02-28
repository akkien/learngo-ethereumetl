package model

import (
	"fmt"
	"strings"

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

//Log : transaction log
type Log struct {
	Address          string `json:"address"`
	BlockNumber      int64  `json:"blockNumber"`
	Data             string `json:"data"`
	LogIndex         int64  `json:"logIndex"`
	Removed          bool   `json:"removed"`
	Topics           string `json:"topics"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex int64  `json:"transactionIndex"`
}

// mapTransaction map rpc result to block
func mapLog(in LogRPC) Log {
	out := Log{}
	var err error

	out.Address = in.Address
	out.BlockNumber, err = util.HexToDec(in.BlockNumber)
	out.Data = in.Data
	out.LogIndex, err = util.HexToDec(in.LogIndex)
	out.Removed = in.Removed
	out.Topics = strings.Join(in.Topics, ",")
	out.TransactionHash = in.TransactionHash
	out.TransactionIndex, err = util.HexToDec(in.TransactionIndex)

	if err != nil {
		fmt.Println("Map log", err)
	}

	return out
}
