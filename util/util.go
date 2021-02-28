package util

import (
	"strconv"
)

// HexToDec convert hexa string to integer number
func HexToDec(hexNum string) (int64, error) {
	numberStr := hexNum[2:] // Remove 0x
	return strconv.ParseInt(numberStr, 16, 64)
}

// MakeRange create range of item
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// GeneratePatitions get export block partitions
func GeneratePatitions(startBlock, endBlock, partitionSize int) [][]int {
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

// GenerateBatchs get export txs, address partitions
func GenerateBatchs(input []string, batchSize int) [][]string {
	length := len(input)

	numBatch := 0
	if length%batchSize == 0 {
		numBatch = length / batchSize
	} else {
		numBatch = length/batchSize + 1
	}

	ret := make([][]string, numBatch)
	for i, j := 0, 0; i < len(input); i, j = i+batchSize, j+1 {
		if i+batchSize < length {
			ret[j] = input[i : i+batchSize]
		} else {
			ret[j] = input[i:length]
		}
	}
	return ret
}
