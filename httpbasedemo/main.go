package main

import (
	"fmt"
	"go-demo/httpbase"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUNT: %s\n", req.URL)
	}
}

func main() {
	//engine := new(Engine)
	//log.Fatalln(http.ListenAndServe(":8080", engine))

	engine := httpbase.New()
	engine.Get("/index", indexHandler)
	engine.Post("/hello", helloHandler)
	engine.Run(":8080")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
