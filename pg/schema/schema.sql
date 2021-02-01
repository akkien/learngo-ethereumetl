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