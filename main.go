package main

import (
	"fmt"
	"math/rand"
)

//0x26955f189e5c0eeccc3368d7e24ca3d41ac17c6d8229e548ac1d292db836a2a2
func main() {
	// var startBlock = flag.Int("s", 0, "block would be parsed from")
	// var endBlock = flag.Int("e", 0, "block would be parsed to")
	// flag.Parse()
	//fmt.Println(*startBlock, *endBlock)

	// start := time.Now()

	// job.ExportAll("postgres://akkien:trungkien@127.0.0.1:5432/ropsten?sslmode=disable", 5)

	// elapsed := time.Since(start)
	// log.Printf("Parse block took %s", elapsed)

	a := rand.Perm(10)
	fmt.Println(a)

}
