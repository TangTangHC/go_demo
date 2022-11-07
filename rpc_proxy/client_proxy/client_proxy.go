package client_proxy

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"go-demo/rpc_proxy/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		panic("连接失败")
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return HelloServiceStub{client}
}

func (s *HelloServiceStub) Hello(res string, reply *string) error {
	err := s.Call(handler.HelloServiceName+".Hello", res, reply)
	if err != nil {
		return err
	}
	return nil
}
