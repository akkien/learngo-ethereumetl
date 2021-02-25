package main

import (
	"fmt"
	"log"
	"time"

	"github.com/akkien/ethereumetl/job"
)

//0x26955f189e5c0eeccc3368d7e24ca3d41ac17c6d8229e548ac1d292db836a2a2
func main() {
	// var startBlock = flag.Int("s", 0, "block would be parsed from")
	// var endBlock = flag.Int("e", 0, "block would be parsed to")
	// flag.Parse()
	//fmt.Println(*startBlock, *endBlock)

	RopstenHTTP := "https://mainnet.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	//RopstenHTTP := "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"
	connStr := "postgres://akkien:lekien9@127.0.0.1:5432/blockchain?sslmode=disable"
	start := time.Now()

	job.ExportAll(11775644, 11775743, 40, 5, RopstenHTTP, connStr, 10)

	elapsed := time.Since(start)
	log.Printf("Parse block took %s", elapsed)

	fmt.Println("OK")

}
