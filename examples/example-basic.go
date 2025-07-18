package examples

import "github.com/gox7/vega"

type (
	xml struct {
		Status string `xml:"status"`
	}
)

func basic() {
	router := vega.NewRouter()

	router.Get("/ping", func(ctx *vega.Context) {
		ctx.WriteString(200, "pong")
	})
	router.Get("/json", func(ctx *vega.Context) {
		ctx.JSON(200, vega.H{"status": "ok"})
	})
	router.Get("/yaml", func(ctx *vega.Context) {
		ctx.YAML(200, vega.H{"status": "ok"})
	})
	router.Get("/xml", func(ctx *vega.Context) {
		ctx.XML(200, xml{Status: "ok"})
	})

	router.Run(":8080")
}
