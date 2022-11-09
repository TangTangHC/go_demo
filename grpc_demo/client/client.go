package main

import (
	"context"
	"fmt"
	"go-demo/grpc_demo/proto"
	"google.golang.org/grpc"
)

func main() {
	// 推荐
	//dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {

	}
	defer dial.Close()

	client := proto.NewGreeterClient(dial)
	hello, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "一个人"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(hello.Message)
}
