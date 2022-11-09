package main

import (
	"context"
	"google.golang.org/grpc"
	"net"

	"go-demo/grpc_demo/proto"
)

type Server struct{}

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello, " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("fail to listen:" + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("fail to start:" + err.Error())
	}
}
