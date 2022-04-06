package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akkien/ethereumetl/job"
	"github.com/akkien/ethereumetl/model"
	"github.com/akkien/ethereumetl/rpc"
	"github.com/akkien/ethereumetl/util"
)

func main() {
	var mode = flag.String("mode", "", "mode to execute the module: pasttime / realtime")
	var startBlock = flag.Int("start", -1, "start block: block would be parsed from")
	var endBlock = flag.Int("end", -1, "end block: block would be parsed to")
	flag.Parse()

	// Setup Log
	LOG_FILE := "./parser.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Hello")

	// RopstenHTTP := "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	RopstenHTTP := "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	connStr := "postgres://postgres:mysecret@127.0.0.1:5432/explorer?sslmode=disable"

	if *mode == "pasttime" {
		if *startBlock <= 0 || *endBlock <= 0 || *startBlock > *endBlock {
			fmt.Println("Please provide valid block range")
		}
		start := time.Now()

		job.ExportAll(*startBlock, *endBlock, 1000, 50, RopstenHTTP, connStr, 5)

		elapsed := time.Since(start)

		logger.Printf("Parse block took %s", elapsed)
		logger.Println("Done")
	} else if *mode == "realtime" {
		lastBlock := -1
		var parse = func() {
			// Get lastest block number
			blockReq, err := rpc.GetBlockNumberRequest()
			if err != nil {
				fmt.Println("Error generate block request", err)
				return
			}

			response, err := rpc.Call(RopstenHTTP, blockReq)
			if err != nil {
				fmt.Println("Error call block rpc", err)
				return
			}

			var blockRes []model.BlockNumberRPCResponse
			err = json.Unmarshal(response, &blockRes)
			if err != nil {
				fmt.Println("Error parse blocks result")
				return
			}

			newBlock, err := util.HexToDec(blockRes[0].Result)
			if err != nil {
				fmt.Println("Error converse block number")
				return
			}

			// Export
			newBlockInt := int(newBlock)
			if newBlockInt != lastBlock {
				lastBlock = newBlockInt
				job.ExportAll(newBlockInt, newBlockInt, 1000, 5, RopstenHTTP, connStr, 5)
			}
		}

		for {
			<-time.After(2 * time.Second)
			fmt.Println("Current time: ", time.Now())
			parse()
		}
	} else {
		fmt.Println("Please provide valid mode: pasttime / realtime. Your input:", *mode)
	}
}
