### Removing Stack Traces from Errors

Since the beginning Buffalo has used the [`github.com/pkg/errors`](https://godoc.org/github.com/pkg/errors) package to provide "stack traces" to developers to help them gain more context to the error.

Unfortunately, Go, doesn't yet provide a good, clean way of getting that information.

In earlier versions of Buffalo if you had a typo in your `index.html` you would get an error that looked something like this:

```text
index.html: line 5: "RootPath": unknown identifier
github.com/gobuffalo/buffalo.sessionSaver.func1
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/session.go:72
github.com/gobuffalo/buffalo.RequestLoggerFunc.func1
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/request_logger.go:54
github.com/gobuffalo/buffalo.(*App).PanicHandler.func1
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/errors.go:95
github.com/gobuffalo/buffalo.(*App).defaultErrorMiddleware.func1
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/errors.go:101
github.com/gobuffalo/buffalo.RouteInfo.ServeHTTP
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/route_info.go:109
github.com/gorilla/mux.(*Router).ServeHTTP
	$GOPATH/pkg/mod/github.com/gorilla/mux@v1.7.3/mux.go:212
github.com/markbates/refresh/refresh/web.ErrorChecker.func1
	$GOPATH/pkg/mod/github.com/markbates/refresh@v1.8.0/refresh/web/web.go:23
net/http.HandlerFunc.ServeHTTP
	$GOROOT/src/net/http/server.go:1995
github.com/gobuffalo/buffalo.(*App).ServeHTTP
	$GOPATH/pkg/mod/github.com/gobuffalo/buffalo@v0.14.0/server.go:134
net/http.serverHandler.ServeHTTP
	$GOROOT/src/net/http/server.go:2774
net/http.(*conn).serve
	$GOROOT/src/net/http/server.go:1878
runtime.goexit
	$GOROOT/src/runtime/asm_amd64.s:1337
```

While the first line is useful, the rest offers no contextual help to those trying to debug the issue. At no point in that error message is a reference to my application. Showing users this sort of "internal plumbing" does not help, and often, confuses the issue even more.

We have been systematically removing the [`github.com/pkg/errors`](https://godoc.org/github.com/pkg/errors) package from all of the Buffalo related projects.

In `<%= to %>` you will now get an error that looks more like this:

```text
index.html: line 5: "RootPath": unknown identifier
```
