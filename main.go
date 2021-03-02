package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/akkien/ethereumetl/job"
	"github.com/akkien/ethereumetl/model"
	"github.com/akkien/ethereumetl/rpc"
	"github.com/akkien/ethereumetl/util"
)

func main() {
	var mode = flag.String("m", "realtime", "mode to execute the module: pasttime / realtime")
	var startBlock = flag.Int("s", -1, "start block: block would be parsed from")
	var endBlock = flag.Int("e", -1, "end block: block would be parsed to")
	flag.Parse()

	//RopstenHTTP := "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	RopstenHTTP := "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	connStr := "postgres://akkien:lekien9@127.0.0.1:5432/blockchain?sslmode=disable"

	if *mode == "pasttime" {
		if *startBlock <= 0 || *endBlock <= 0 || *startBlock < *endBlock {
			fmt.Println("Please provide valid block range")
		}
		start := time.Now()
		//9759324
		job.ExportAll(*startBlock, *endBlock, 1000, 5, RopstenHTTP, connStr, 5)

		elapsed := time.Since(start)
		log.Printf("Parse block took %s", elapsed)
		fmt.Println("Done")
	} else if *mode == "realtime" {
		lastBlock := -1
		var parse = func() {
			// Get lastest block number
			blockReq, err := rpc.GetBlockNumberRequest()
			if err != nil {
				fmt.Println("Error generate block request", err)
			}
			response, err := rpc.Call(RopstenHTTP, blockReq)
			if err != nil {
				fmt.Println("Error call block rpc", err)
			}

			var blockRes []model.BlockNumberRPCResponse
			err = json.Unmarshal(response, &blockRes)
			if err != nil {
				fmt.Println("Error parse blocks result")
			}
			newBlock, err := util.HexToDec(blockRes[0].Result)
			if err != nil {
				fmt.Println("Error converse block number")
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
