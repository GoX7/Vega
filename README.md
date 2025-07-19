<p align="center">
  <a href="https://github.com/gox7/vega">
    <img src="https://github.com/GoX7/Vega/blob/main/assets/image.png" height=120px />
  </a>
  <br />
  <a href="https://pkg.go.dev/github.com/gox7/vega">
    <img src="https://pkg.go.dev/badge/github.com/gox7/vega.svg" alt="Go Reference"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/gox7/vega">
    <img src="https://goreportcard.com/badge/github.com/gox7/vega" alt="Go Report Card"/>
  </a>
</p>

# Vega

**Vega** is a minimal, fast, and lightweight HTTP framework for Go.  
It builds directly on `net/http`, offering clear abstractions for routing, request/response handling, and data binding—nothing more, nothing less.

---

## ✨ Features

- **Routing** with `GET`, `POST`, `PUT`, `PATCH`, `DELETE`  
- **Grouped routes** via `router.Group(prefix)`  
- **Context** (`*vega.Context`) wrapping `http.Request` and `http.ResponseWriter`  
- **Response helpers**  
  - `ctx.Status(code)`  
  - `ctx.Text(code, string)` / `ctx.WriteString`  
  - `ctx.Write(code, []byte)`  
  - `ctx.JSON(code, any)`  
  - `ctx.XML(code, any)`  
  - `ctx.YAML(code, any)`  
- **Request binding**  
  - `ctx.Bind()` → `[]byte`  
  - `ctx.BindString()` → `string`  
  - `ctx.BindJSON(obj any)`  
  - `ctx.BindXML(obj any)`  
  - `ctx.BindYAML(obj any)`  
- **Redirect**: `ctx.Redirect(code, url)`  
- **Utility type**: `vega.H` (alias for `map[string]any`)  
- **No extra dependencies** (only `gopkg.in/yaml.v3`)

---

## 📦 Installation

```bash
go get github.com/gox7/vega
````

---

## 🧪 Quick Start

```go
package main

import "github.com/gox7/vega"

func main() {
  router := vega.NewRouter()

  // Simple GET
  router.Get("/", func(ctx *vega.Context) {
    ctx.JSON(200, vega.H{"message": "Hello, Vega!"})
  })

  // Bind JSON
  type Payload struct {
    Name string `json:"name"`
  }
  router.Post("/echo", func(ctx *vega.Context) {
    var p Payload
    if err := ctx.BindJSON(&p); err != nil {
      ctx.Status(400)
      ctx.JSON(400, vega.H{"error": "invalid JSON"})
      return
    }
    ctx.JSON(200, p)
  })

  // Redirect
  router.Get("/old", func(ctx *vega.Context) {
    ctx.Redirect(302, "/")
  })

  // Raw text / bytes
  router.Get("/raw", func(ctx *vega.Context) {
    ctx.Text(200, "plain text")
  })

  // Start server
  router.Run(":8080")
}
```

---

## 📁 Route Groups

```go
router := vega.NewRouter()
api := router.Group("/api/v1")

api.Get("/status", func(ctx *vega.Context) {
  ctx.JSON(200, vega.H{"status": "ok"})
})
api.Get("/ping", func(ctx *vega.Context) {
  ctx.JSON(200, vega.H{"ping": "pong"})
})

router.Run(":8080")
```

---

## 🚧 Project Status

Vega is in **early development**. Planned improvements:

1. Middleware support (`Use`, `Next`, `Abort`)
2. Path parameters (`/users/:id`)
3. Static file serving
4. Built‑in logging & recovery

Contributions and feedback are welcome!

---

## 📄 License

Vega is distributed under the **MIT License**.
See [LICENSE](LICENSE) for details.
