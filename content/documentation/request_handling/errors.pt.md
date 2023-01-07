---
name: Error Handling
seoDescription: "Errors Handling"
seoKeywords: ["buffalo", "go", "golang", "errors"]
weight: 7
aliases:
  - /pt/docs/errors
  - /pt/docs/errors
---

# Error Handling

An `error` is Go way to tell something went wrong. In this chapter, you'll learn how to return errors from a route handler and how Buffalo will catch any non-handled error.

## Returning Errors From a Handler

The easiest way to produce an error response is to return a standard Go error:
```go
func MyHandler(c buffalo.Context) error {
  // Return any go error, this will result in a 500 status code.
  return errors.New("boom!")
}
```

A `nil` error will produce a raw HTTP 200 response:
```go
func MyHandler(c buffalo.Context) error {
  // HTTP 200
  return nil
}
```

If you need to customize the error message or the HTTP code, use the {{< doclink href="github.com/gobuffalo/buffalo#DefaultContext.Error" message="Error" >}} method:
```go
func MyHandler(c buffalo.Context) error {
  // Use the Error function on the context.
  // This will result in a status code of 401.
  return c.Error(401, errors.New("Unauthorized!"))
}
```

## Default Error Handling (Development)

In "development" mode (`GO_ENV=development`), Buffalo will generate some helpful errors pages for you.

<figure>
  <img src="/assets/images/500_example.png" title="screenshot">
  <figcaption>An example of a `500` error in development mode.</figcaption>
</figure>

If you use a JSON or an XML content type, the error is returned in the proper type:

```json
{
  "error": "could not find test/",
  "trace": "could not find test/\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\nnet/http.StripPrefix.func1\n\t/usr/local/go/src/net/http/server.go:1986\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162\ngithub.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2694\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1830\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_amd64.s:2361",
  "code": 404
}
```

```xml
<response code="404">
  <error>could not find test/</error>
  <trace>could not find test/ github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 net/http.StripPrefix.func1 /usr/local/go/src/net/http/server.go:1986 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162 github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127 net/http.serverHandler.ServeHTTP /usr/local/go/src/net/http/server.go:2694 net/http.(*conn).serve /usr/local/go/src/net/http/server.go:1830 runtime.goexit /usr/local/go/src/runtime/asm_amd64.s:2361</trace>
</response>
```

In "production" mode (`GO_ENV=production`), Buffalo will not generate pages that have developer style information, because this would give precious information to hackers. Instead the pages are simpler.

## Custom Error Handling

While Buffalo will handle errors for you out of the box, it can be useful to handle errors in a custom way. To accomplish this, Buffalo allows for the mapping of HTTP status codes to specific handlers. This means the error can be dealt with in a custom fashion.

```go
app = buffalo.New(buffalo.Options{
  Env: ENV,
})

// We associate the HTTP 422 status to a specific handler.
// All the other status code will still use the default handler provided by Buffalo.
app.ErrorHandlers[422] = func(status int, err error, c buffalo.Context) error {
  res := c.Response()
  res.WriteHeader(422)
  res.Write([]byte(fmt.Sprintf("Oops!! There was an error: %s", err.Error())))
  return nil
}

app.GET("/oops", MyHandler)

func MyHandler(c buffalo.Context) error {
  return c.Error(422, errors.New("Oh no!"))
}
```

```text
GET /oops -> [422] Oops!! There was an error: Oh no!
```

In the above example any error from your application that returns a status of `422` will be caught by the custom handler and will be dealt with accordingly.

