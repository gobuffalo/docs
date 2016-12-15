# Routing

Buffalo uses the wonderful [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with it's own. This guide walks you through all you'll need to know about how Buffalo handles routing.

{{#panel title="Creating a new Buffalo App (and router)"}}

Buffalo applications come in two flavors:

#### Automatic (_recommended_)

```go
a := buffalo.Automatic(buffalo.Options{})
```

#### Standard

```go
a := buffalo.New(buffalo.Options{})
```

It is highly recommend to use `buffalo.Automatic` to create your application. The `Automatic` command will configure your new application with a variety of settings and functionality that we believe we be useful to 90% of all web applications. If you want "complete" control over your application, then just create a `New` Buffalo app.
{{/panel}}

{{#panel title="Mapping Handlers"}}

All routing in Buffalo results in the calling of a `buffalo.Handler` function. The signature for a `buffalo.Handler` looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

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

Mapping a `buffalo.Handler` to an HTTP method takes the form of:

```go
a.GET("/some/path", SomeHandler)
a.POST("/some/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```
{{/panel}}

{{#panel title="Parameters"}}

Query string and other parameters are available from the [`buffalo.Context`](/docs/context) that is passed into the `buffalo.Handler`.

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, render.String(c.Param("name")))
})
// etc...
```

Given the above code sample, if we make a request with `GET /users?name=ringo`, the response should should be `200: ringo`.
{{/panel}}

{{#panel title="Named Parameters"}}

Since Buffalo is the [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux) under the covers, it means we can get access to some of the goodness it provides, in this case, the ability to create sudo-regular expression patterns in the mapped path that will get converted into parameters that can be accessed from a [`buffalo.Context`](/docs/context).

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, render.String(c.Params("name")))
})
// etc...
```

Given the above code sample, if we make a request with `GET /users/ringo`, the response should should be `200: ringo`.

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, render.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, render.String(c.Params("name")))
})
// etc...
```

You may map seemingly similar paths, like `/users/new` and `/users/{name}` without any issues. The router will make sure they get to the same place.
{{/panel}}

{{#panel title="Groups"}}
{{/panel}}

{{#panel title="Resources"}}
{{/panel}}
