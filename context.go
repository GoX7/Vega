package vega

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net"
	"net/http"

	"gopkg.in/yaml.v3"
)

type (
	Context struct {
		Request *http.Request
		Writer  responseWriter
	}
)

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  newResponseWriter(w),
	}
}

func (ctx *Context) Status(code int) {
	ctx.Writer.writer.WriteHeader(code)
}

func (ctx *Context) ClientIp() string {
	ip, _, _ := net.SplitHostPort(ctx.Request.RemoteAddr)
	if ip == "::1" {
		return "127.0.0.1"
	}
	return ip
}

func (ctx *Context) Redirect(code int, url string) {
	http.Redirect(ctx.Writer.writer, ctx.Request, url, code)
}

func (ctx *Context) Write(code int, obj []byte) {
	ctx.Writer.SetCode(code)
	ctx.Writer.Write(obj)
}

func (ctx *Context) WriteString(code int, obj string) {
	ctx.Writer.SetCode(code)
	ctx.Writer.Write([]byte(obj))
}

func (ctx *Context) Text(code int, obj string) {
	ctx.Writer.SetCode(code)
	ctx.Writer.Write([]byte(obj))
}

func (ctx *Context) Bind() ([]byte, error) {
	return io.ReadAll(ctx.Request.Body)
}

func (ctx *Context) BindString() (string, error) {
	data, err := io.ReadAll(ctx.Request.Body)
	return string(data), err
}

func (ctx *Context) Form(key string) string {
	ctx.Request.ParseForm()
	return ctx.Request.FormValue(key)
}

func (ctx *Context) FormDefault(key string, fallback string) string {
	ctx.Request.ParseForm()
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) QueryDefault(key string, fallback string) string {
	value := ctx.Request.URL.Query().Get(key)
	if value == "" {
		return fallback
	}
	return value
}

func (ctx *Context) JSON(code int, obj any) error {
	ctx.Writer.SetCode(code)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(ctx.Writer.writer).Encode(obj)
}

func (ctx *Context) BindJSON(obj any) error {
	decoder := json.NewDecoder(ctx.Request.Body)
	return decoder.Decode(obj)
}

func (ctx *Context) XML(code int, obj any) error {
	ctx.Writer.SetCode(code)
	ctx.Writer.Header().Set("Content-Type", "application/xml")
	return xml.NewEncoder(ctx.Writer.writer).Encode(obj)
}

func (ctx *Context) BindXML(obj any) error {
	decoder := xml.NewDecoder(ctx.Request.Body)
	return decoder.Decode(obj)
}

func (ctx *Context) YAML(code int, obj any) error {
	ctx.Writer.SetCode(code)
	ctx.Writer.Header().Set("Content-Type", "application/yaml")
	return yaml.NewEncoder(ctx.Writer.writer).Encode(obj)
}

func (ctx *Context) BindYAML(obj any) error {
	decoder := yaml.NewDecoder(ctx.Request.Body)
	return decoder.Decode(obj)
}
