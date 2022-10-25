package httpbase

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method, url string, handlerFunc HandlerFunc) {
	engine.router.addRouter(method, url, handlerFunc)
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
	c := newContext(w, res)
	engine.router.handle(c)
}
