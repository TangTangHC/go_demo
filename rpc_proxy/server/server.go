package main

import (
	"go-demo/rpc_proxy/server_proxy"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"go-demo/rpc_proxy/handler"
)

func main() {
	listen, _ := net.Listen("tcp", ":1234")
	server_proxy.RegisterHelloService(handler.NewHelloService{})
	//rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})

	for {
		conn, _ := listen.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}
