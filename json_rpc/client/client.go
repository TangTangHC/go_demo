package main

import "net/rpc"

func main() {

	client, _ := rpc.Dial("tcp", ":1234")
	var reply string
	client.Call("")

}
