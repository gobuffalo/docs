## Writing your own middleware

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

By implementing the `buffalo.MiddlewareFunc` interface you are able to control the flow of execution in your application. Think an authorization middleware; send errors off to your favorite monitoring tool; load data on to the `buffalo.Context`, and more.

### Example

```go
// UserIPMiddleware gets the user IP and sets it to the context.
func UserIPMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if xRealIP := c.Request().Header.Get("X-Real-Ip"); len(xRealIP) > 0 {
      c.Set("user_ip", xRealIP)
      return next(c)
    }

    if xForwardedFor := c.Request().Header.Get("X-Forwarded-For"); len(xForwardedFor) > 0 {
      c.Set("user_ip", xForwardedFor)
      return next(c)
    }

    h, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
      return err
    }
    c.Set("user_ip", h)
    return next(c)
  }
}
```