# Routing

Buffalo uses the wonderful [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with it's own. This guide walks you through all you'll need to know about how Buffalo handles routing.

{{#panel title="Creating a new Buffalo App (and router)" name="new-app"}}

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

Buffalo apps allow for the grouping of end-points. This allows for common functionality, such as [middleware](/docs/middleware) to be collected together. A great example of this would be an API end-point.

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // responds to GET /api/v1/users
})
```

By default a group will inherit any middleware from it's parent app.

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

In the above example the `/api/v1` group will use both `SomeMiddleware` and `APIAuthorizer`. See [middleware](/docs/middleware) for more information about using, skipping, and clearing middleware.

{{/panel}}

{{#panel title="Resources"}}

Often web applications need to build very similiar "CRUD" end-points. To Help reduce the amount of thought and complexity involved in this, Buffalo supports the concept of a "Resource".

```go
type Resource interface {
	List(Context) error
	Show(Context) error
	New(Context) error
	Create(Context) error
	Edit(Context) error
	Update(Context) error
	Destroy(Context) error
}
```

```go
type UserResource struct{
  buffalo.Resource
}

a.Resource("/users", &UserResource{&buffalo.BaseResource{}})
```

The above code example would be the equivalent of the following:

```go
ur := &UserResource{}
a.GET("/users", ur.List)
a.GET("/users/new", ur.New)
a.GET("/users/{id}", ur.Show)
a.GET("/users/{id}/edit", ur.Edit)
a.POST("/users", ur.Create)
a.PUT("/users/{id}", ur.Update)
a.DELETE("/users/{id}", ur.Destroy)
```

{{/panel}}
