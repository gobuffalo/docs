# Routing

Buffalo uses the wonderful, [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with it's own. This guide walks you through all you'll need to know about how Buffalo handles routing.

#### Automatic

```go
func App() http.Handler {
  a := buffalo.Automatic(buffalo.Options{})

  return a
}
```

#### Standard

```go
func App() http.Handler {
  a := buffalo.New(buffalo.Options{})

  return a
}
```

### `buffalo.Handler`

All routing in Buffalo results in the calling of a `buffalo.Handler` function. The signature for a `buffalo.Handler` looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

### Groups
