package httpbase

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{prefix: group.prefix + prefix, parent: group, engine: engine}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method, url string, handlerFunc HandlerFunc) {
	pattern := group.prefix + url
	log.Printf("Route %s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handlerFunc)
}

func (group *RouterGroup) Get(url string, handlerFunc HandlerFunc) {
	group.addRouter("GET", url, handlerFunc)
}

func (group *RouterGroup) Post(url string, handlerFunc HandlerFunc) {
	group.addRouter("POST", url, handlerFunc)
}

func (engine *Engine) Run(port string) error {
	return http.ListenAndServe(port, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	c := newContext(w, res)
	engine.router.handle(c)
}
