package model

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
