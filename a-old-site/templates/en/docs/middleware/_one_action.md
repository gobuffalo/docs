## Using Middleware with One Action

Often there are cases when you want to use a piece of middleware on just one action, and not on the whole application or resource.

Since the definition of a piece of middleware is that it takes in a `buffalo.Handler` and returns a `buffalo.Handler` you can wrap any `buffalo.Handler` in a piece of middlware.

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(MyHandler))
```

This does not affect the rest of the middleware stack that is already in place, instead it appends to the middleware chain for just that one action.

This can be taken a step further, by wrapping unlimited numbers of middleware around a `buffalo.Handler`.

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(AnotherPieceOfMiddleware(MyHandler)))
```
