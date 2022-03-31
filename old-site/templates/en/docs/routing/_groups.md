## Groups

Buffalo apps allow for the grouping of end-points. This allows for common functionality, such as [middleware](/en/docs/middleware) to be collected together. A great example of this would be an API end-point.

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

In the above example the `/api/v1` group will use both `SomeMiddleware` and `APIAuthorizer`. See [middleware](/en/docs/middleware) for more information about using, skipping, and clearing middleware.
