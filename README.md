# Vega

Vega is a minimal and fast HTTP web framework for Go.  
It provides a lightweight core with clean abstractions and built-in support for request routing, data binding, and response rendering.

---

## 🚀 Features & Advantages

### ✅ Supported Functionality

- HTTP methods: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- Exact path routing (no wildcards or parameters yet)
- Context object (`*vega.Context`) wrapping `http.Request` and `http.ResponseWriter`
- Response rendering:
  - HTML (ctx.HTML)
  - JSON (`ctx.JSON`)
  - XML (`ctx.XML`)
  - YAML (`ctx.YAML`)
  - Raw strings / bytes
- Request body binding:
  - `BindJSON`
  - `BindXML`
  - `BindYAML`
- Utility type `vega.H` (alias for `map[string]any`) for JSON responses
- Built on Go's standard `net/http` without hidden magic

### 💡 Why Choose Vega?

- Minimal and focused – everything you need, nothing you don't
- Zero external dependencies (except YAML encoder)
- Transparent and easy to debug
- Fully compatible with Go's ecosystem and tooling
- A great base to build your own features on top

---

## 🧪 Example

```go
package main

import (
	"vega"
)

func main() {
	app := vega.NewRouter()

	app.Get("/", func(ctx *vega.Context) {
		ctx.JSON(200, vega.H{
			"message": "Hello from Vega",
		})
	})

	app.Post("/echo", func(ctx *vega.Context) {
		var data map[string]any
		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(400, vega.H{"error": "Invalid JSON"})
			return
		}
		ctx.JSON(200, data)
	})

	app.Run(":8080")
}
````

---

## 📦 Installation

```bash
go get github.com/gox7/vega
```

---

## ⚠️ Project Status

This is an early-stage release. The core is stable and works, but features like middleware, route grouping, parameterized paths, and error recovery are under development.
