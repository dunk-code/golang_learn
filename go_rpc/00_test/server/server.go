package main

import (
	"golang_learn/go_rpc/00_test/model"
	"net"
	"net/rpc"
)

type RpcServer struct {
}

func (p *RpcServer) CalcArea(req model.CalcRequest, rsp *model.CalcResponse) error {
	rsp.Area = req.Long * req.Width
	return nil
}

func main() {
	rpc.RegisterName("RpcServer", new(RpcServer))

	listen, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		panic(err)
	}

	for true {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}
