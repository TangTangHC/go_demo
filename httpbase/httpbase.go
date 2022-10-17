package httpbase

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method, url string, handlerFunc HandlerFunc) {
	key := method + "-" + url
	engine.router[key] = handlerFunc
}

func (engine *Engine) Get(url string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", url, handlerFunc)
}

func (engine *Engine) Post(url string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", url, handlerFunc)
}

func (engine *Engine) Run(port string) error {
	return http.ListenAndServe(port, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	key := res.Method + "-" + res.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, res)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND, %q\n", res.URL.Path)
	}
}
