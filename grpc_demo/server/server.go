package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"

	"go_demo/grpc_demo/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	if name, ok := md["name"]; ok {
		for i, v := range name {
			fmt.Println(i, v)
		}
	}
	return &proto.HelloReply{
		Message: "Hello, " + request.Name,
	}, nil
}

func main() {

	inteceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")
		i, err := handler(ctx, req)
		fmt.Println("返回数据：", i)
		return i, err
	}
	opt := grpc.UnaryInterceptor(inteceptor)
	g := grpc.NewServer(opt)
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
