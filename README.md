# Vega

Vega is a minimal HTTP framework for Go, designed to be fast, simple, and easy to extend. Inspired by frameworks like Gin and Echo, Vega provides a lightweight foundation for building web applications.

## Supported

- **HTTP methods**: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- **Simple routing** with exact path matching
- **Request context** (`*Context`) wrapping `http.Request` and `http.ResponseWriter`
- **Response rendering**:
  - `JSON`
  - `XML`
  - `YAML`
  - raw `[]byte` and `string`
- **Body binding**:
  - `BindJSON`
  - `BindXML`
  - `BindYAML`
- **Utility type**: `vega.H` (`map[string]any`) for clean JSON responses

## Installation

```bash
go get github.com/yourusername/vega
````

## Example

```go
package main

import (
	"vega"
)

func main() {
	app := vega.NewRouter()

	app.Get("/", func(ctx *vega.Context) {
		ctx.JSON(200, vega.H{
			"message": "Hello, Vega!",
		})
	})

	app.Post("/echo", func(ctx *vega.Context) {
		var data map[string]any
		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(400, vega.H{"error": "invalid JSON"})
			return
		}
		ctx.JSON(200, data)
	})

	app.Run(":8080")
}
```

## Status

This is an early version intended as a minimal starting point.
Current limitations:

* No middleware support
* No route grouping
* No path parameters
* No built-in logging or recovery

## License
MIT
