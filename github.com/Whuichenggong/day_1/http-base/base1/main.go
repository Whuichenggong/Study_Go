package main

import (
	"fmt"
	"log"
	"net/http"
)

//第一天 使用内置的net/http库 启动Web服务

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	//启动Web服务 第二个参数是 基于net/http标准库实现Web框架的
	//Handler是一个接口 需要实现方法 只要传入 只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了
	log.Fatal(http.ListenAndServe(":9999", nil))

}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
