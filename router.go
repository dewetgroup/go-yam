package yam

import (
	"fmt"
	"net/http"
	"regexp"
)

type Multiplexer interface {
	Handle(pattern string, handler http.Handler)
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
}

type Router struct {
	handler     http.Handler
	middleware  []Middleware
	multiplexer Multiplexer
}

func NewRouter() *Router {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", adapt(DefaultHandler))
	multiplexer := Multiplexer(serveMux)
	return &Router{nil, make([]Middleware, 0), multiplexer}
}

// Applies a middleware to this router.
func (router *Router) Use(middleware Middleware) {
	router.middleware = append(router.middleware, middleware)
}

// Route registers a handler for the given method and path.
// The path must start with a "/" and end with a "/".
// The path must not contain spaces.
// The path must not contain consecutive slashes.
func (router *Router) Route(method Method, path string, handler Handler) *Router {
	validate(path)
	pattern := fmt.Sprintf("%s %s", method, path)
	router.multiplexer.Handle(pattern, adapt(handler))
	return router
}

// Link links the otherRouter router to the path.
// The path must not already be handled by the router.
// When path == "/", this is equivalent to merging the routers.
func (router *Router) Link(path string, otherRouter *Router) *Router {
	validate(path)
	prefix := path[:len(path)-1]
	router.multiplexer.Handle(path, http.StripPrefix(prefix, otherRouter))
	return router
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	router.build()
	router.handler.ServeHTTP(writer, request)
}

func (router *Router) build() {
	router.handler = router.multiplexer
	// Iterate in reverse to preserve order of middleware.
	for i := len(router.middleware) - 1; i >= 0; i-- {
		middleware := router.middleware[i]
		router.handler = middleware(router.handler)
	}
}

func validate(path string) {
	// TODO: Revisit this validation function...
	regex := `^\/(?:[^\/\s{}]+|{[^\/\s{}]+})*(?:\/(?:[^\/\s{}]+|{[^\/\s{}]+}))*\/?$`
	if !regexp.MustCompile(regex).MatchString(path) {
		panic("path is invalid")
	}
}
