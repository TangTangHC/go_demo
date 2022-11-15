package main

import (
	"context"
	"fmt"
	"go_demo/grpc_auth_token/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		t := time.Now()
		md := metadata.New(map[string]string{
			"token": "thc1 ",
		})
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("Invoker time is ", time.Since(t))
		return err
	}
	dialOption := grpc.WithUnaryInterceptor(interceptor)
	var options []grpc.DialOption
	options = append(append(options, dialOption), grpc.WithInsecure())
	dial, err := grpc.Dial(":9000", options...)
	if err != nil {
		panic("dial error " + err.Error())
	}
	client := proto.NewGretterClient(dial)

	hello, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "一个人"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(hello.Data)
}
