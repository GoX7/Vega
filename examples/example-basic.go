package main

import "github.com/gox7/vega"

func main() {
	router := vega.NewRouter()

	router.Get("/ping", func(ctx *vega.Context) {
		ctx.WriteString(200, "Hello, world!")
	})
	router.Get("/json", func(ctx *vega.Context) {
		ctx.JSON(200, vega.H{"status": "ok"})
	})
	router.Get("/yaml", func(ctx *vega.Context) {
		ctx.YAML(200, vega.H{"status": "ok"})
	})
	router.Get("/xml", func(ctx *vega.Context) {
		ctx.XML(200, "status - ok")
	})

	router.Run(":8080")
}
