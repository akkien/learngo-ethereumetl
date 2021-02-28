package rpc

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Call calling RPC method from fullnode
func Call(providerURI string, jsonRPC []byte) ([]byte, error) {
	response, err := http.Post(providerURI, "application/json", bytes.NewBuffer(jsonRPC))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}
