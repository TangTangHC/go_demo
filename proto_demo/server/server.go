package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"

	"go-demo/proto_demo/proto"
)

type server struct {
	*proto.UnimplementedGreeterServer
}

func (s server) SayHello(context context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(req.Id)
	fmt.Println(req.Name)
	fmt.Println(req.IsDelete)
	fmt.Println(req.Score)
	fmt.Println(req.Titles)
	fmt.Println(req.File)
	return &proto.HelloReply{
		Msg: "hello, " + req.Name,
	}, nil
}

func (s server) Ping(context.Context, *empty.Empty) (*proto.Pong, error) {
	return &proto.Pong{
		Id: "success",
	}, nil
}

func main() {
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, server{})
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic("listen error " + err.Error())
	}
	err = s.Serve(listen)
	if err != nil {
		panic("start error " + err.Error())
	}
}
