package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(res string, reply *string) error {
	*reply = "hello, " + res
	return nil
}

func main() {
	listen, _ := net.Listen("tcp", ":1234")
	_ = rpc.RegisterName("HelloService", &HelloService{})
	for {
		conn, _ := listen.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
