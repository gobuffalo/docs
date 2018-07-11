# Routing

Buffalo uses the wonderful [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers, to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with its own. This guide walks you through all you'll need to know about how Buffalo handles routing.

<%= title("Creating a new Buffalo App (and router)", {name: "new-app"}) %>

The app configuration is located in the `app.go` file.

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
})
```

The default setup should handle most of your needs, but you are free to customize it to fit your use case.

You can check the available options list here: https://godoc.org/github.com/gobuffalo/buffalo#Options

<%= title("Mapping Handlers", {}) %>

All routing in Buffalo results in the calling of a `buffalo.Handler` function. The signature for a `buffalo.Handler` looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

If you already know about **MVC pattern**, `buffalo.Handler` functions manages the Controller part: this is the place where all the app logic goes. The handler function takes a `buffalo.Context` struct, which contains everything you need about the current request.

See the [context](/docs/context) to understand the `buffalo.Context` interface.

##### Supported HTTP Methods

Buffalo supports the following HTTP methods out of the box:

* GET
* POST
* PUT
* PATCH
* DELETE
* OPTIONS
* HEAD

You can also match all HTTP methods using `ANY`.

Mapping a `buffalo.Handler` to an HTTP method takes the form of:

```go
a.GET("/some/path", SomeHandler)
a.POST("/some/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

As you can see, you can use inline handlers if you want. For more readability though, it's often better to separate your handlers into multiple files. If you have many handlers managing users stuff, you can group them into a `users.go` file in the [`actions`](/en/docs/directory-structure) folder, for instance.

<%= title("Parameters", {})  %>

Query string and other parameters are available from the [`buffalo.Context`](/docs/context) that is passed into the `buffalo.Handler`.

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users?name=ringo`, the response should should be `200: ringo`.

<%= title("Named Parameters", {}) %>

Since Buffalo is the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) under the covers, it means we can get access to some of the goodness it provides. In this case, the ability to create pseudo-regular expression patterns in the mapped path that will get converted into parameters that can be accessed from a [`buffalo.Context`](/docs/context).

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users/ringo`, the response should should be `200: ringo`.

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, r.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

You may map seemingly similar paths, like `/users/new` and `/users/{name}` without any issues. The router will make sure they get to the same place.

### Regular expressions

[github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) provides a way to use regular expressions, so you can pre-filter queries:

```go
a.GET("/articles/{id:[0-9]+}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("id")))
})
```

<%= title("Groups", {}) %>

Buffalo apps allow for the grouping of end-points. This allows for common functionality, such as [middleware](/docs/middleware) to be collected together. A great example of this would be an API end-point.

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // responds to GET /api/v1/users
})
```

By default a group will inherit any middleware from its parent app.

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

In the above example the `/api/v1` group will use both `SomeMiddleware` and `APIAuthorizer`. See [middleware](/docs/middleware) for more information about using, skipping, and clearing middleware.

<%= title("Loose Slash", {}) %>

<%= sinceVersion("0.10.3") %>

By default, the configured routes for your app **match strictly** the pattern you defined: if the pattern ends with a slash, the URL won't be accessible without a slash. Conversely, a pattern without an ending slash won't match an URL with an ending slash.

To allow your routes to ignore the ending slash, you can use the `LooseSlash` option:

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  LooseSlash:  true,
})
```
