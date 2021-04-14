package crown

import (
	"net/http"
)

// 重新定义框架请求的方法
type HandlerFunc func(*Context)

// 核心结构体
type Engine struct {
	router *Router  // 使用map记录路由信息
}


// 实例化结构体
func New() *Engine  {
	engine := &Engine{
		router: newRouter(),
	}
	return engine
}

// 添加路由到结构体
func (engine *Engine)addRouter(method, pattern string, handler HandlerFunc)  {
	engine.router.addRouter(method, pattern, handler)
}

// GET请求
func (engine *Engine)GET(pattern string, handler HandlerFunc)  {
	engine.addRouter("GET", pattern, handler)
}

// POST请求
func (engine *Engine) POST(pattern string,handler HandlerFunc){
	engine.addRouter("POST",pattern,handler)
}


// 启动服务
func (engine *Engine)Run(addr string) (err error)  {
	return http.ListenAndServe(addr, engine)
}

//engine 实现ServeHTTP接口（所有的请求都会走到这）
//查找是否路由映射表存在，如果存在则调用，否则返回404
func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	c := newContext(w, req)
	engine.handleHTTPRequest(c)
}

//v2 新增
func (engine *Engine) handleHTTPRequest(c *Context){
	key := c.Method + "-" + c.Path
	if handler, ok := engine.router.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}


