package main

import (
	"fmt"
	"log"
	"time"

	"github.com/akkien/ethereumetl/job"
)

func main() {
	// var startBlock = flag.Int("s", 0, "block would be parsed from")
	// var endBlock = flag.Int("e", 0, "block would be parsed to")
	// flag.Parse()
	//fmt.Println(*startBlock, *endBlock)

	RopstenHTTP := "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	//RopstenHTTP := "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	connStr := "postgres://akkien:lekien9@127.0.0.1:5432/blockchain?sslmode=disable"

	start := time.Now()

	job.ExportAll(11946583, 11946583, 1000, 5, RopstenHTTP, connStr, 5)

	elapsed := time.Since(start)
	log.Printf("Parse block took %s", elapsed)

	fmt.Println("OK")
}
