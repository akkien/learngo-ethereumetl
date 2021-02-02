package rpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// RopstenHTTP Fullnode HTTP Url for Ropsten Network
// const RopstenHTTP = "https://ropsten.infura.io/v3/2ee8969fa00742efb10051fc923552e1"

// Call calling RPC method from fullnode
func Call(providerURI string, jsonRPC []byte) []byte {
	fmt.Println(string(jsonRPC))
	response, err := http.Post(providerURI, "application/json", bytes.NewBuffer(jsonRPC))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}
