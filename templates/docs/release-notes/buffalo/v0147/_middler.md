## Middler Interface

The [`gobuffalo/buffalo#Middler`](https://godoc.org/github.com/gobuffalo/buffalo#Middler) interface
allows [`gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) implementations to provide additional custom
[`gobuffalo/buffalo#MiddlewareFunc`](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareFunc)
to the those already configured for the application

```go
// Middler allows Resource implementations to provide
// additional custom Middleware to the those already
// configured.
type Middler interface {
  Use() []MiddlewareFunc
}
```

```go
 func (w WidgetResource) AuthMW(next buffalo.Handler) buffalo.Handler {
   return func(c buffalo.Context) error {
     // do work
     return next(c)
  }
}

func (w WidgetResource) Use() []MiddlewareFunc {
  mw := []buffalo.MiddlewareFunc{
    w.AuthMW,
  }

  mw = append(mw, func(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
      // do work
      return next(c)
    }
  })

  return mw
}
```

