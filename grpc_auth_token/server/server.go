package main

import (
	"context"
	"fmt"
	"go_demo/grpc_auth_token/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type Server struct {
	proto.UnimplementedGretterServer
}

func (s Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Data: "hello, " + req.Name,
	}, nil
}

func (s Server) Ping(ctx context.Context, empty *emptypb.Empty) (*proto.Pong, error) {
	return &proto.Pong{Data: "1"}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, b := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !b {
			return nil, status.Error(codes.Unauthenticated, "metadata is null")
		}
		sli, ok := md["token"]
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "not found token from metadata")
		}
		if sli[0] != "thc" {
			return nil, status.Error(codes.Unavailable, "token is unavailable")
		}
		i, err := handler(ctx, req)
		return i, err
	}
	serverOption := grpc.UnaryInterceptor(interceptor)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic("listen error " + err.Error())
	}
	server := grpc.NewServer(serverOption)
	proto.RegisterGretterServer(server, &Server{})
	err = server.Serve(lis)
	if err != nil {
		panic("serve error " + err.Error())
	}

}
