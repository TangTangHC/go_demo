package main

import (
	"context"
	"fmt"
	"go_demo/grpc_demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("消耗时间：%s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	//var opts []grpc.DialOption
	opts := make([]grpc.DialOption, 2)
	opts[0] = grpc.WithTransportCredentials(insecure.NewCredentials())
	opts[1] = opt
	//opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, opt)
	// 推荐
	dial, err := grpc.Dial(":8080", opts...)
	//dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	//dial, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {

	}
	defer dial.Close()

	client := proto.NewGreeterClient(dial)
	md := metadata.New(map[string]string{
		"name":     "one person",
		"password": "any pass",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	hello, err := client.SayHello(ctx, &proto.HelloRequest{Name: "一个人"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(hello.Message)
}
