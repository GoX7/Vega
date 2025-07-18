package vega

import (
	"net/http"
)

type (
	responseWriter struct {
		writer http.ResponseWriter
	}
)

func newResponseWriter(writer http.ResponseWriter) responseWriter {
	return responseWriter{
		writer: writer,
	}
}

func (rw responseWriter) SetCode(code int) {
	rw.writer.WriteHeader(code)
}

func (rw responseWriter) Write(data []byte) {
	rw.writer.Write(data)
}

func (rw responseWriter) Header() http.Header {
	return rw.writer.Header()
}
