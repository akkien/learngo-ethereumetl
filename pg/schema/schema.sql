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
);

CREATE TABLE transactions (
    block_hash CHAR(66),
    block_number BIGINT REFERENCES blocks(number),
    from_address CHAR(42),
    gas BIGINT,
    gas_price BIGINT,
    hash CHAR(66) PRIMARY KEY,
    input TEXT,
    nonce BIGINT,
    r CHAR(66),
    s CHAR(66),
    to_address CHAR(42),
    transaction_index SMALLINT,
    v VARCHAR(66),
    value DECIMAL(38,0),
    created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE receipts (
    transaction_hash CHAR(66) PRIMARY KEY REFERENCES transactions(hash),
    transaction_index SMALLINT,
    block_hash CHAR(66),
    block_number BIGINT,
    cumulative_gas_used BIGINT,
    gas_used BIGINT,
    contract_address CHAR(42),
    logs_count INT,
    logs_bloom TEXT,
    root CHAR(66),
    status BOOLEAN,
    created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE logs (
    log_index SMALLINT,
    transaction_hash CHAR(66) REFERENCES transactions(hash),
    transaction_index SMALLINT,
    block_number BIGINT,
    address CHAR(42),
    data TEXT,
    topics TEXT,
    removed BOOLEAN,
    created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT pk_logs PRIMARY KEY (transaction_hash, log_index)
);

CREATE TABLE traces ( 
    transaction_hash CHAR(66) PRIMARY KEY REFERENCES transactions(hash),
    block_number BIGINT,
	type TEXT,
	from_address CHAR(42),
	to_address CHAR(42),
	value DECIMAL(38,0),
	gas BIGINT,
	gas_used BIGINT,
	input TEXT,
	output TEXT,
	error TEXT,
	time TEXT,
	calls JSON,
    created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE internal_transactions ( 
    transaction_hash CHAR(66) REFERENCES transactions(hash),
	block_number BIGINT,
	type TEXT,
	from_address CHAR(42),
	to_address CHAR(42),
	value DECIMAL(38,0),
	gas BIGINT,
	gas_used BIGINT,
	input TEXT,
	output TEXT,
	error TEXT,
	time TEXT,
	call_depth TEXT,
    created_timestamp TIMESTAMPTZ NOT NULL
		DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX transactions_block_number_index ON transactions(block_number);
CREATE INDEX transactions_from_to_index ON transactions(from_address, to_address);
CREATE INDEX receipts_block_number_index ON receipts(block_number);
CREATE INDEX logs_transaction_hash_index ON logs(transaction_hash);
