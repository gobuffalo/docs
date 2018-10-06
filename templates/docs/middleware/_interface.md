<%= title("The Middleware Interface") %>

The [`buffalo.MiddlewareFunc`](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareFunc) interface is any function that takes a `buffalo.Handler` and returns a `buffalo.Handler`.

```go
func MyMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    // do some work before calling the next handler
    err := next(c)
    // do some work after calling the next handler
    return err
  }
}
```

By implementing the `buffalo.MiddlewareFunc` interface you are able to control the flow of execution in your application, think an authorization middleware; send errors off to your favorite monitoring tool; load data on to the `buffalo.Context`, and more.
