## 🚀 Features & Advantages

### ✨ Supported Functionality

- **HTTP methods**: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- **Exact-path routing** (no wildcard or parameter support yet)
- **Context object (`*Context`)** wrapping `http.Request` and `http.ResponseWriter`
- **Response rendering** in multiple formats:
  - `JSON`
  - `XML`
  - `YAML`
  - raw `string` / `[]byte`
- **Request binding**:
  - `BindJSON`
  - `BindXML`
  - `BindYAML`
- **Helper type**: `vega.H` (`map[string]any`) for quick JSON responses
- **Minimal internal abstractions**, fully compatible with Go's `net/http`

### 💡 Why Choose Vega 0.1.0?

- **Lightweight and minimal** – only core functionality, ideal as a foundation
- **Zero external dependencies** (except for optional `gopkg.in/yaml.v3`)
- **Focused on performance** – simple routing with minimal overhead
- **Transparent and easy to extend or debug**
- **Compatible** with standard Go tooling and HTTP abstractions

---

**Note:** This is an early release. Middleware (`Use`, `Next`, `Abort`), route grouping, parameterized routes, error recovery, and static file serving are planned for future versions.
