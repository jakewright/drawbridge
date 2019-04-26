package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jakewright/drawbridge/middleware"
	"github.com/urfave/negroni"
)

type RouterFactory struct {
	negroni *negroni.Negroni // Use negroni to handle middleware
	mux     *mux.Router      // Use gorilla mux internally to do the legwork
}

// Return a new factory with some defaults
func NewRouterFactory() *RouterFactory {
	negroni := negroni.New()
	mux := mux.NewRouter().StrictSlash(true)
	return &RouterFactory{negroni, mux}
}

// Return an http.Handler that can be used as the argument to http.ListenAndServe
func (f *RouterFactory) Build() http.Handler {
	// The mux router needs to be the last item of middleware added to the negroni instance.
	f.negroni.UseHandler(f.mux)
	return f.negroni
}

// Add middleware that will be applied to every request. Middleware handlers are executed in the order they are defined.
func (f *RouterFactory) AddMiddleware(middlewares ...middleware.Middleware) {
	for _, middleware := range middlewares {
		f.negroni.UseFunc(middleware)
	}
}

func (f *RouterFactory) HandleFunc(
	method string,
	path string,
	handlerFunc http.HandlerFunc,
	middlewares ...middleware.Middleware) {

	// Use the adapter to transform the http.handlerFunc into a http.Handler
	handler := http.HandlerFunc(handlerFunc)

	// Defer to the Handle method
	f.Handle(method, path, handler, middlewares...)
}

// Register a route for the given path and method. Optionally add middleware.
// see https://github.com/gorilla/mux for options available for the path, including variables.
func (f *RouterFactory) Handle(
	method string,
	path string,
	handler http.Handler,
	middlewares ...middleware.Middleware) {

	// A slice to hold all of the middleware once it's converted (including the handler itself)
	var stack []negroni.Handler

	// The middleware functions have type Middleware but they need to conform to the negroni.Handler interface.
	// By using the negroni.HandlerFunc adapter, they will be given the method required by the interface.
	for _, middleware := range middlewares {
		stack = append(stack, negroni.HandlerFunc(middleware))
	}

	// The handler needs to be treated like middleware.
	// The negroni.Wrap function can convert an http.Handler into a negroni.Handler.
	stack = append(stack, negroni.Wrap(handler))

	// Create the new route using mux
	route := f.mux.NewRoute()

	// If the last character of the path is an asterisk, create a path prefix
	if path[len(path)-1] == '*' {
		// Be sure to strip the asterisk off again
		route.PathPrefix(path[:len(path)-1])
		// Otherwise just add the path normally
	} else {
		route.Path(path)
	}

	// Use a new instance of negroni with our handler stack as the route handler
	route.Handler(negroni.New(stack...))

	// If a method is defined, restrict to that method only.
	if len(method) > 0 {
		route.Methods(method)
	}
}

// A helper function to add a GET route
func (f *RouterFactory) Get(path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	f.Handle("GET", path, handler, middlewares...)
}

// A helper function to add a POST route
func (f *RouterFactory) Post(path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	f.Handle("POST", path, handler, middlewares...)
}

// A helper function to add a PUT route
func (f *RouterFactory) Put(path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	f.Handle("PUT", path, handler, middlewares...)
}

// A helper function to add a PATCH route
func (f *RouterFactory) Patch(path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	f.Handle("PATCH", path, handler, middlewares...)
}

// A helper function to add a DELETE route
func (f *RouterFactory) Delete(path string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {
	f.Handle("DELETE", path, handler, middlewares...)
}
