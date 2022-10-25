package main

import (
	"go-demo/httpbase"
	"net/http"
)

func main() {
	engine := httpbase.New()
	engine.Get("/", func(ctx *httpbase.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.Get("/hello", func(ctx *httpbase.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	engine.Post("/login", func(ctx *httpbase.Context) {
		ctx.Json(http.StatusOK, httpbase.H{
			"userName": ctx.PostForm("userName"),
			"password": ctx.PostForm("password"),
		})
	})
	engine.Post("/userInfo", func(ctx *httpbase.Context) {
		b := ctx.PostBody()
		ctx.Json(http.StatusOK, b)
	})
	engine.Run(":8080")

}
