package model

import (
	"fmt"

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
	BlockNumber       int64       `json:"blockNumber"`
	ContractAddress   interface{} `json:"contractAddress"`
	CumulativeGasUsed int64       `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           int64       `json:"gasUsed"`
	LogsCount         int64       `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	Status            int64       `json:"status"`
	To                string      `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  int64       `json:"transactionIndex"`
}

// mapTransaction map rpc result to block
func mapReceipt(in ReceiptRPC) Receipt {
	out := Receipt{}
	var err error
	out.BlockHash = in.BlockHash
	out.BlockNumber, err = util.HexToDec(in.BlockNumber)
	out.ContractAddress = in.ContractAddress
	out.CumulativeGasUsed, err = util.HexToDec(in.CumulativeGasUsed)
	out.GasUsed, err = util.HexToDec(in.GasUsed)
	out.LogsCount = int64(len(in.Logs))
	out.LogsBloom = in.LogsBloom
	out.Status, err = util.HexToDec(in.Status)
	out.TransactionHash = in.TransactionHash
	out.TransactionIndex, err = util.HexToDec(in.TransactionIndex)

	if err != nil {
		fmt.Println("Map receipt", err)
	}

	return out
}

// RPCResponseToReceipt map rpc result to block
func RPCResponseToReceipt(response *[]ReceiptRPCResponse) ([]Receipt, []Log) {
	receipts := make([]Receipt, len((*response)))
	logs := []Log{}
	for i, receiptRes := range *response {
		receipts[i] = mapReceipt(receiptRes.Result)
		for _, rpcTx := range receiptRes.Result.Logs {
			logs = append(logs, mapLog(rpcTx))
		}
	}

	return receipts, logs
}
