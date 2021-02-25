package rpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Call calling RPC method from fullnode
func Call(providerURI string, jsonRPC []byte) []byte {
	response, err := http.Post(providerURI, "application/json", bytes.NewBuffer(jsonRPC))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}
