package vega

import (
	"fmt"
	"net/http"
	"time"
)

type (
	Engine struct {
		handlers map[string]func(*Context)
		group    Group
	}
	H map[string]any
)

func NewRouter() *Engine {
	return &Engine{
		handlers: make(map[string]func(*Context)),
	}
}

func (eng *Engine) Get(pattern string, handler func(ctx *Context)) {
	eng.handlers["GET||"+pattern] = handler
}

func (eng *Engine) Post(pattern string, handler func(ctx *Context)) {
	eng.handlers["POST||"+pattern] = handler
}

func (eng *Engine) Put(pattern string, handler func(ctx *Context)) {
	eng.handlers["PUT||"+pattern] = handler
}

func (eng *Engine) Patch(pattern string, handler func(ctx *Context)) {
	eng.handlers["PATCH||"+pattern] = handler
}

func (eng *Engine) Delete(pattern string, handler func(ctx *Context)) {
	eng.handlers["DELETE||"+pattern] = handler
}

func (eng *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := eng.handlers[r.Method+"||"+r.URL.Path]; ok {
		context := newContext(w, r)
		handler(context)
	} else if handler, ok := eng.group.handlers[r.Method+"||"+r.URL.Path]; ok {
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
