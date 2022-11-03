package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	// 1、实例化server
	listen, _ := net.Listen("tcp", ":1234")
	// 2、注册handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3、服务启动
	conn, _ := listen.Accept()
	rpc.ServeConn(conn)
}
