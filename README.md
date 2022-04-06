# Golang Ethereum ETL

## Introduction

This project is in the Learning Golang series. The purpose is to pull ethereum data from fullnode to database

The project is inspired by Python Ethereum ETL: <https://github.com/blockchain-etl/ethereum-etl>

## Setup

1. Start postgres database

Create database tables by running sql script at: pg/schema/schema.sql

2. Build application: run `make build`

## Start application

```sh
# ./ethereumetl -mode=<mode> -start=<start_block> -end=<end_block>

./ethereumetl -mode=pasttime -start=10 -end=20 # parse block 10 to block 20

./ethereumetl -mode=realtime # parse on going block
```
