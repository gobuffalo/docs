<%= title("Group Middleware", {}) %>

```go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// authorize the API end-point
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```

In the above example the `MyMiddleware` and `AnotherPieceOfMiddleware` middlewares will be called on _all_ requests, but the `AuthorizeAPIMiddleware` middleware will only be called on the `/api/*` routes.

```text
GET /foo -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> MyMiddleware -> AnotherPieceOfMiddleware -> AuthorizeAPIMiddleware -> UsersHandler
```
