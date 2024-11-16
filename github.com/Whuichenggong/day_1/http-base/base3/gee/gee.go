package gee

import (
	"fmt"
	"net/http"
)

//首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法

type HandlerFunc func(http.ResponseWriter, *http.Request)

// 在Engine中，添加了一张路由映射表router
// key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello
// 如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法

type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + pattern
	engine.router[key] = handler
}

// 用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，(*Engine).Run()方法，是 ListenAndServe 的包装。
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 Not Found: %s\n", req.URL)
	}

}
