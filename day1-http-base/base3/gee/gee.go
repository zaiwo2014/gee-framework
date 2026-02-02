package gee

import (
	"fmt"
	"net/http"
)

// 定义一个handlerFunc函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 定义路由结构体
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// 构造函数
func New() *Engine {
	//构造一个新的 map 实例, 本质是 路由地址 + 处理器的映射
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 往 map 中添加对应路由以及请求处理器的映射
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	/**
	method: 请求方法
	pattern: 请求路径
	handler: 请求处理器
	*/
	key := method + "-" + pattern
	engine.router[key] = handler
}

// Get方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
