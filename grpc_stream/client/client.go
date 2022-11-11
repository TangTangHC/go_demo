package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"

	"go-demo/grpc_stream/proto"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	stream, _ := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "一个信息",
	})

	for {
		a, err := stream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(a)
	}

	i := 0
	putStream, _ := c.PutStream(context.Background())
	for {
		i++
		putStream.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("一个人发送第%d条消息", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	allStream, _ := c.AllStream(context.Background())
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		defer group.Done()
		for {
			recv, err := allStream.Recv()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println("收到服务端消息：", recv.Data)
		}
	}()
	go func() {
		defer group.Done()
		for {
			err = allStream.Send(&proto.StreamReqData{Data: "一个人QQ"})
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			time.Sleep(time.Second)
		}
	}()
	group.Wait()
}
