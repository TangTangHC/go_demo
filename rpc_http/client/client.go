package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type method struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int      `json:"id"`
}

func main() {
	client := new(http.Client)
	method := method{
		Method: "HelloService.Hello",
		Params: []string{"一个人"},
		Id:     0,
	}

	buf := new(bytes.Buffer)
	b, _ := json.Marshal(method)
	buf.Write(b)

	resp, _ := client.Post("http://localhost:1234/jsonrpc", "application/json", buf)
	all, _ := io.ReadAll(resp.Body)
	fmt.Println(string(all))
}
