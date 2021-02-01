package util

import (
	"fmt"
	"strconv"
)

// HexToDec convert hexa string to integer number
func HexToDec(hexaString string) int {
	numberStr := hexaString[2:]
	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(output)
}
