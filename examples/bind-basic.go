package examples

import (
	"fmt"

	"github.com/gox7/vega"
)

type (
	json struct {
		Status string `json:"status"`
	}
	yaml struct {
		Status string `yaml:"status"`
	}
	xml struct {
		Status string `xml:"status"`
	}
)

func bind() {
	router := vega.NewRouter()

	router.Post("/ping", func(ctx *vega.Context) {
		pong, _ := ctx.BindString()
		fmt.Println(pong)
	})
	router.Post("/json", func(ctx *vega.Context) {
		var json json
		ctx.BindJSON(&json)
	})
	router.Post("/yaml", func(ctx *vega.Context) {
		var yaml yaml
		ctx.BindYAML(&yaml)
	})
	router.Post("/xml", func(ctx *vega.Context) {
		var xml xml
		ctx.BindXML(&xml)
	})

	router.Run(":8080")
}
