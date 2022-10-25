package httpbase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Write        http.ResponseWriter
	Req          *http.Request
	Path, Method string
	Params       map[string]string
	StatusCode   int
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func newContext(write http.ResponseWriter, res *http.Request) *Context {
	return &Context{
		Write:  write,
		Req:    res,
		Path:   res.URL.Path,
		Method: res.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostBody() (h H) {
	json.NewDecoder(c.Req.Body).Decode(&h)
	//b, err := io.ReadAll(c.Req.Body)
	//if !json.Valid(b) {
	//	http.Error(c.Write, err.Error(), 500)
	//}
	return h
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Write.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Write.Header().Set(key, value)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Write.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Write)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Write, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Write.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Write.Write([]byte(html))
}
