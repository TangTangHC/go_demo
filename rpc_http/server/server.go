package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(res string, aply *string) error {
	*aply = "hello, " + res
	return nil
}

func main() {
	rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: request.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
