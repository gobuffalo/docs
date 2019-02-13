## Using Middleware

```go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)
```

In the above example all requests will first go through the `MyMiddleware` middleware, and then through the `AnotherPieceOfMiddleware` middleware before first getting to their final handler.

_NOTE: Middleware defined on an application is automatically inherited by all routes and groups in that application._

