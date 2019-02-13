## Renderer Interface

In order for a renderer to be able to be used with [`Context#Render`](/docs/context) it must implement the following interface:

```go
// Renderer interface that must be satisified to be used with
// buffalo.Context.Render
type Renderer interface {
  ContentType() string
  Render(io.Writer, Data) error
}

// Data type to be provided to the Render function on the
// Renderer interface.

type Data map[string]interface{}
```

The [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/master/render) [[godoc]](https://godoc.org/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!

