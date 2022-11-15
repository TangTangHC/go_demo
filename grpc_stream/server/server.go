package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"

	"go_demo/grpc_stream/proto"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) GetStream(data *proto.StreamReqData, streamServer proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++

		streamServer.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})

		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(streamServer proto.Greeter_PutStreamServer) error {
	for {
		if recv, err := streamServer.Recv(); err != nil {
			fmt.Println(err.Error())
			break
		} else {
			fmt.Println(recv.Data)
		}
	}
	return nil
}

func (s *server) AllStream(streamServer proto.Greeter_AllStreamServer) error {
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		defer group.Done()
		for {
			err := streamServer.Send(&proto.StreamResData{
				Data: "这是服务端",
			})
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer group.Done()
		for {
			recv, err := streamServer.Recv()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println("服务端收到：", recv.Data)
		}
	}()
	group.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err.Error())
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	s.Serve(listen)
}
