package util

import (
	"fmt"
	"strconv"
)

// HexToDec convert hexa string to integer number
func HexToDec(hexNum string) int {
	numberStr := hexNum[2:] // Remove 0x
	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(output)
}

// MakeRange create range of item
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// GeneratePatition get export block partitions
func GeneratePatition(startBlock, endBlock, partitionSize int) [][]int {
	numPartitions := (endBlock-startBlock)/partitionSize + 1
	ret := make([][]int, numPartitions)

	for i, start := 0, startBlock; start <= endBlock; i, start = i+1, start+partitionSize {
		end := start + partitionSize - 1
		if end > endBlock {
			end = endBlock
		}
		ret[i] = []int{start, end}
	}
	return ret
}
