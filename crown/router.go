package crown

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		handlers:make(map[string]HandlerFunc),
	}
}

func (router *Router)addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handler
}