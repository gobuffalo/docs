# Routing

Buffalo uses the wonderful, [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with it's own. This guide walks you through all you'll need to know about how Buffalo handles routing.

```go
func App() http.Handler {
  a := buffalo.Automatic()

  return a
}
```
