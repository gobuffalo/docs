<%= title("Clearing Middleware", {}) %>

Since middleware is [inherited](#using-middleware) from its parent, there maybe times when it is necessary to start with a "blank" set of middleware.

<%= codeTabs() { %>
```go
// actions/app.go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// clear out any previously defined middleware
g.Middleware.Clear()
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```

```text
// OUTPUT
GET /foo -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> AuthorizeAPIMiddleware -> UsersHandler
```
<% } %>
