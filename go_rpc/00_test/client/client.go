package main

import (
	"fmt"
	"golang_learn/go_rpc/00_test/model"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	response := &model.CalcResponse{}
	err = client.Call("RpcServer.CalcArea", model.CalcRequest{Long: 10, Width: 10}, response)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Area)
}
