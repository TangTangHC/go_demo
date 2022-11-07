package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	conn, _ := net.Dial("tcp", ":1234")
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	client.Call("HelloService.Hello", "一个人", &reply)
	fmt.Println(reply)

}
