package model

import (
	"fmt"

	"github.com/akkien/ethereumetl/util"
)

// TransactionRPC RPC response
type TransactionRPC struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	V                string `json:"v"`
	Value            string `json:"value"`
}

// Transaction for PostgreSQL
type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      int64  `json:"blockNumber"`
	From             string `json:"from"`
	Gas              int64  `json:"gas"`
	GasPrice         int64  `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            int64  `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex int64  `json:"transactionIndex"`
	V                string `json:"v"`
	Value            string `json:"value"` ///////////////////////// TOTO: change type to big.Int
}

// mapTransaction map rpc result to block
func mapTransaction(in TransactionRPC) Transaction {
	out := Transaction{}
	var err error = nil

	out.BlockHash = in.BlockHash
	out.BlockNumber, err = util.HexToDec(in.BlockNumber)
	out.From = in.From
	out.Gas, err = util.HexToDec(in.Gas)
	out.GasPrice, err = util.HexToDec(in.GasPrice)
	out.Hash = in.Hash
	out.Input = in.Input
	out.Nonce, err = util.HexToDec(in.Nonce)
	out.R = in.R
	out.S = in.S
	out.To = in.To
	out.TransactionIndex, err = util.HexToDec(in.TransactionIndex)
	out.V = in.V
	out.Value = in.Value

	if err != nil {
		fmt.Println("Map transaction", err)
	}

	return out
}
