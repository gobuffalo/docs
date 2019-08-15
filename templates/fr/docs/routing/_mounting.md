## Connecter des applications http.Handler

<%= sinceVersion("0.9.4") %>

Parfois, vous souhaiterez réutiliser certains composants d'autres applications. En utilisant la méthode [`Mount`](https://godoc.org/github.com/gobuffalo/buffalo#App.Mount), vous pouvez connecter un [`http.Handler`](https://golang.org/pkg/net/http/#Handler) standard à une route, tout comme vous le feriez avec une action Buffalo.

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

Puisque les applications Buffalo implémentent l'interface `http.Handler`, vous pouvez également connecter une autre application Buffalo et construire des applications modulaires.
