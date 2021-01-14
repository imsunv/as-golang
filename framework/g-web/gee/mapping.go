package gee

import (
	"log"
	"net/http"
)

type mapping struct {
	handlers map[string]HandlerFunc
}

func newMapping() *mapping {
	return &mapping{handlers: make(map[string]HandlerFunc)}
}

func (r *mapping) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *mapping) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
