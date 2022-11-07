package server_proxy

import (
	"net/rpc"

	"go-demo/rpc_proxy/handler"
)

type HelloServicer interface {
	Hello(res string, reply *string) error
}

func RegisterHelloService(svr HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, svr)
}
