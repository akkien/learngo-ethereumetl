package model

import (
	"github.com/akkien/ethereumetl/util"
)

// ReceiptRPC : transaction receipt
type ReceiptRPC struct {
	BlockHash         string      `json:"blockHash"`
	BlockNumber       string      `json:"blockNumber"`
	ContractAddress   interface{} `json:"contractAddress"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           string      `json:"gasUsed"`
	Logs              []LogRPC    `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	Status            string      `json:"status"`
	To                string      `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  string      `json:"transactionIndex"`
}

// Receipt : transaction receipt
type Receipt struct {
	BlockHash         string      `json:"blockHash"`
	BlockNumber       int         `json:"blockNumber"`
	ContractAddress   interface{} `json:"contractAddress"`
	CumulativeGasUsed int         `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           int         `json:"gasUsed"`
	LogsCount         int         `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	Status            int         `json:"status"`
	To                string      `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  int         `json:"transactionIndex"`
}

// mapTransaction map rpc result to block
func mapReceipt(in ReceiptRPC) Receipt {
	out := Receipt{}

	out.BlockHash = in.BlockHash
	out.BlockNumber = util.HexToDec(in.BlockNumber)
	out.ContractAddress = in.ContractAddress
	out.CumulativeGasUsed = util.HexToDec(in.CumulativeGasUsed)
	out.From = in.From
	out.GasUsed = util.HexToDec(in.GasUsed)
	out.LogsCount = len(in.Logs)
	out.LogsBloom = in.LogsBloom
	out.Status = util.HexToDec(in.Status)
	out.To = in.To
	out.TransactionHash = in.TransactionHash
	out.TransactionIndex = util.HexToDec(in.TransactionIndex)

	return out
}
