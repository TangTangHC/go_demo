package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type method struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int      `json:"id"`
}

func main() {
	dial, err := net.Dial("tcp", ":1234")
	if err != nil {
		return
	}

	m := method{
		Method: "HelloService.Hello",
		Params: []string{"一个人"},
		Id:     0,
	}

	marshal, err := json.Marshal(m)
	dial.Write(marshal)

	b := make([]byte, 1024)
	dial.Read(b)
	fmt.Println(string(b))

}
