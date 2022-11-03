package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic("连接失败")
	}
	var reply string
	err = client.Call("HelloService.Hello", "一个人", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

	// rpc调用需要解决的问题：1、call id 2、序列化和反序列化

}
