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

// 这段代码的作用将HTTP请求的路由和对应的处理函数注册到路由表中的核心方法
// pattern路由路径
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//将HTTp方法和路径拼接成唯一一个键 作为路由表的router的键
	//
	key := method + "-" + pattern

	//将处理函数 handler 存入路由表中，关联到对应的路由键。
	engine.router[key] = handler
}

// 用户调用 addRoute("GET", "/home", someHandlerFunc) 在 engine.router 映射表中，会存储一个键值对：
// 调用engine.GET("/home", someHandlerFunc)： 实际是 等价 engine.addRoute("GET", "/home", someHandlerFunc)
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 这段代码隐藏了调用ServeHTTP
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
// 不需要显式调用 ServeHTTP
// 在 Go 的 HTTP 框架中，ServeHTTP 是 http.Handler 接口的约定方法。当你把 Engine 作为服务器的处理器传递时，它会被 ListenAndServe 自动调用。

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {

		w.WriteHeader(http.StatusNotFound) // 设置 404 状态码
		fmt.Fprintf(w, "404 Not Found: %s\n", req.URL.Path)
	}

}
