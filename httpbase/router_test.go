package httpbase

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hello/d/c", nil)
	r.addRouter("GET", "/hello/e/c", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	n, hash := r.getRouter("GET", "/hello/thc")
	fmt.Println(r.roots["GET"])
	if n == nil {
		t.Fatal("should match /hello/:name")
	}
	fmt.Println(hash)
}

func TestTripInsert(t *testing.T) {
	node := new(node)
	url := "/"
	node.insert("/", strings.Split(url, "/"), len(strings.Split(url, "/")))
	url = "/hello/b/c"
	node.insert(url, strings.Split(url, "/"), 3)
	url = "/hello/b/d"
	node.insert(url, strings.Split(url, "/"), 3)
	url = "/hello/:name"
	node.insert(url, strings.Split(url, "/"), 2)
	url = "/hello/:name/:log"
	node.insert(url, strings.Split(url, "/"), 3)
	fmt.Println(node)

	url = ":name"
	if url[0] != ':' {
		t.Fatal("not equals ':'")
	}

	fmt.Println(strings.Split("/", "/"))
}
