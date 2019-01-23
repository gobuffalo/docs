## Mounting http.Handler Apps

<%= sinceVersion("0.9.4") %>

Sometimes, you'll want to reuse some components from other apps. Using the [`Mount`](https://godoc.org/github.com/gobuffalo/buffalo#App.Mount) method, you can bind a standard [`http.Handler`](https://golang.org/pkg/net/http/#Handler) to a route, just like you'll do with a normal route handler.

```go
func muxer() http.Handler {
  f := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "%s - %s", req.Method, req.URL.String())
  }
  mux := mux.NewRouter()
  mux.HandleFunc("/foo/", f).Methods("GET")
  mux.HandleFunc("/bar/", f).Methods("POST")
  mux.HandleFunc("/baz/baz/", f).Methods("DELETE")
  return mux
}

a.Mount("/admin", muxer())
```

Since Buffalo `App` implements the `http.Handler` interface, you can also mount another Buffalo app and build modular apps.
