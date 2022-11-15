package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	__ "go_demo/rpc_proxy/proto"
)

func main() {
	req := __.HelloRequest{
		Name:    "thc",
		Age:     12,
		Courses: []string{"语文", "数学"},
	}

	res, _ := proto.Marshal(&req)
	fmt.Println(res)

	newReq := __.HelloRequest{}
	proto.Unmarshal(res, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
