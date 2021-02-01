package model

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

/*
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
*/

// Transaction for PostgreSQL
type Transaction struct {
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
