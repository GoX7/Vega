package vega

import (
	"fmt"
	"net/http"
	"time"
)

type (
	Engine struct {
		router map[string]func(*Context)
	}
	H map[string]any
)

func NewRouter() *Engine {
	return &Engine{
		router: make(map[string]func(*Context)),
	}
}

func (eng *Engine) Get(pattern string, handler func(ctx *Context)) {
	eng.router["GET||"+pattern] = handler
}

func (eng *Engine) Post(pattern string, handler func(ctx *Context)) {
	eng.router["POST||"+pattern] = handler
}

func (eng *Engine) Put(pattern string, handler func(ctx *Context)) {
	eng.router["PUT||"+pattern] = handler
}

func (eng *Engine) Patch(pattern string, handler func(ctx *Context)) {
	eng.router["PATCH||"+pattern] = handler
}

func (eng *Engine) Delete(pattern string, handler func(ctx *Context)) {
	eng.router["DELETE||"+pattern] = handler
}

func (eng *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := eng.router[r.Method+"||"+r.URL.Path]; ok {
		context := newContext(w, r)
		handler(context)
	} else {
		http.NotFound(w, r)
	}
}

func (eng *Engine) Run(addr string) error {
	server := http.Server{
		Handler:      eng,
		Addr:         addr,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	}

	fmt.Println("Vega: start server on " + addr)
	return server.ListenAndServe()
}
