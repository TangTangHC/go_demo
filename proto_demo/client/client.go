package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"go_demo/proto_demo/proto"
)

func main() {
	dial, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("连接失败" + err.Error())
	}
	defer dial.Close()

	s := proto.HelloReply_Result{Name: "hh", Url: "www"}
	fmt.Println(s)

	h := proto.HelloReply{Sex: proto.Gender_FEMALE}
	fmt.Println(h.String())

	p := proto.HelloRequest{Map: map[string]string{
		"k1": "v1",
		"k2": "v2",
	}, AddTime: timestamppb.New(time.Now())}
	fmt.Println(p.String())

	client := proto.NewGreeterClient(dial)
	pong, _ := client.Ping(context.Background(), &empty.Empty{})
	//pong, _ = client.Ping(context.Background(), &emptypb.Empty{})
	fmt.Println(pong)
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "一个人",
	})
	if err != nil {
		fmt.Println("调用失败", err.Error())
	}
	fmt.Println(reply.Msg)

}
