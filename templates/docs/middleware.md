# Middleware

Middleware allows for the interjection of code in the request/response cycle. Common use cases for middleware are things like logging (which Buffalo already does), authentication requests, etc. Buffalo ships with some common middleware, so please checkout out [https://godoc.org/github.com/gobuffalo/buffalo/middleware](https://godoc.org/github.com/gobuffalo/buffalo/middleware) for details on those.

<%= partial("topics.html") %>

<%= panel("The Middleware Interface", {name: "interface"}) { %>
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
<% } %>

<%= panel("Using Middleware", {}) { %>

```go
a := buffalo.Automatic(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)
```

In the above example all requests will first go through the `MyMiddleware` middleware, and then through the `AnotherPieceOfMiddleware` middleware before first getting to their final handler.

_NOTE: Middleware defined on an application is automatically inherited by all routes and groups in that application._

<% } %>

<%= panel("Group Middleware", {}) { %>

```go
a := buffalo.Automatic(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// authorize the API end-point
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```

In the above example the `MyMiddleware` and `AnotherPieceOfMiddleware` middlewares will be called on _all_ requests, but the `AuthorizeAPIMiddleware` middleware will only be called on the `/api/*` routes.

```
GET /foo -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> MyMiddleware -> AnotherPieceOfMiddleware -> AuthorizeAPIMiddleware -> UsersHandler
```

<% } %>

<%= panel("Skipping Middleware", {}) { %>

There are times when, in an application, you want to add middleware to the entire application, or a group, but not call that middleware on a few individual handlers. Buffalo allows you to create these sorts of mappings.

```go
a := buffalo.Automatic(buffalo.Options{})
a.Use(AuthorizeUser)

// skip the AuthorizeUser middleware for the NewUser and CreateUser handlers.
a.Middleware.Skip(AuthorizeUser, NewUser, CreateUser)

a.GET("/users/new", NewUser)
a.POST("/users", CreateUser)
a.GET("/users", ListUsers)
a.GET("/users/{id}", ShowUser)
```

```
GET /users/new -> NewUser
POST /users -> CreateUser
GET /users -> AuthorizeUser -> ListUsers
GET /users/{id} -> AuthorizeUser -> ShowUser
```

See [https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip) for more details on the `Skip` function.
<% } %>

<%= panel("Clearing Middleware", {}) { %>

Since middleware is [inherited](#using-middleware) from its parent, there maybe times when it is necessary to start with a "blank" set of middleware.

```go
a := buffalo.Automatic(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// clear out any previously defined middleware
g.Middleware.Clear()
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```
```
GET /foo -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> AuthorizeAPIMiddleware -> UsersHandler
```
<% } %>
