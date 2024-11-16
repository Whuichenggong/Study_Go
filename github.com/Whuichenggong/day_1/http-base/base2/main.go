package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

//我们定义了一个空的结构体Engine，实现了方法ServeHTTP。
//这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；
//第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprint(w, "URL.Path = %q\n", req.URL.Path)

	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)

		}
	default:
		fmt.Fprintf(w, "404 page not found: %s\n", req.URL.Path)

	}
}

// 我们给 ListenAndServe 方法的第二个参数传入了刚才创建的engine实例。
// 即，将所有的HTTP请求转向了我们自己的处理逻辑。
// 在实现Engine之前，我们调用 http.HandleFunc 实现了路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。比如/hello。
// 在实现Engine之后 在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
