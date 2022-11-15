package main

import (
	"fmt"
	"go_demo/rpc_proxy/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string
	client.Hello("一个人", &reply)
	fmt.Println(reply)
}
