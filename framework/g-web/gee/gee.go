package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type HandlerFuncOrigin func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	mapping *mapping
}

func (engine *Engine) GetMapping() *mapping {
	return engine.mapping
}

func New() *Engine {
	return &Engine{mapping: newMapping()}
}

func (engine *Engine) addMapping(method string, pattern string, handler HandlerFunc) {
	engine.GetMapping().addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addMapping("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addMapping("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	log.Printf("http server start on %s \n", addr)
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.GetMapping().handle(newContext(w, req))
}
