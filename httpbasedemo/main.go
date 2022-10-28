package main

import (
	"go-demo/httpbase"
	"net/http"
)

func main() {
	engine := httpbase.New()
	engine.Get("/index", func(ctx *httpbase.Context) {
		ctx.HTML(http.StatusOK, "<h1>index page</h1>")
	})
	v1 := engine.Group("/v1")
	{
		//v1.Get("/", func(c *httpbase.Context) {
		//	c.HTML(http.StatusOK, "<h1>Hello httpbase</h1>")
		//})

		v1.Get("/hello", func(c *httpbase.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	//engine.Get("/hello", func(ctx *httpbase.Context) {
	//	ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	//})
	//engine.Get("/hello/:name", func(ctx *httpbase.Context) {
	//	ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	//})
	//engine.Get("/assets/*filepath", func(ctx *httpbase.Context) {
	//	ctx.Json(http.StatusOK, httpbase.H{"filepath": ctx.Param("filepath")})
	//})
	//engine.Post("/login", func(ctx *httpbase.Context) {
	//	ctx.Json(http.StatusOK, httpbase.H{
	//		"userName": ctx.PostForm("userName"),
	//		"password": ctx.PostForm("password"),
	//	})
	//})
	//engine.Post("/userInfo", func(ctx *httpbase.Context) {
	//	b := ctx.PostBody()
	//	ctx.Json(http.StatusOK, b)
	//})
	engine.Run(":8080")

}
