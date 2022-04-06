---
name: Middleware
weight: 6
aliases:
  - /docs/middleware
  - /en/docs/middleware
---
# Middleware

Middleware allows for the interjection of code in the request/response cycle. Common use cases for middleware are things like logging (which Buffalo already does), authentication requests, etc.

A list of "known" middleware packages can be found at [https://toolkit.gobuffalo.io/tools?topic=middleware](https://toolkit.gobuffalo.io/tools?topic=middleware).

## Writing Your Own Middleware

The [`buffalo.MiddlewareFunc`](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareFunc) interface is any function that takes a `buffalo.Handler` and returns a `buffalo.Handler`.

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

By implementing the `buffalo.MiddlewareFunc` interface you are able to control the flow of execution in your application. Think an authorization middleware; send errors off to your favorite monitoring tool; load data on to the `buffalo.Context`, and more.

### Example

```go
// UserIPMiddleware gets the user IP and sets it to the context.
func UserIPMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if xRealIP := c.Request().Header.Get("X-Real-Ip"); len(xRealIP) > 0 {
      c.Set("user_ip", xRealIP)
      return next(c)
    }

    if xForwardedFor := c.Request().Header.Get("X-Forwarded-For"); len(xForwardedFor) > 0 {
      c.Set("user_ip", xForwardedFor)
      return next(c)
    }

    h, _, err := net.SplitHostPort(c.Request().RemoteAddr)
    if err != nil {
      return err
    }
    c.Set("user_ip", h)
    return next(c)
  }
}
```

## Using Middleware

```go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)
```

In the above example all requests will first go through the `MyMiddleware` middleware, and then through the `AnotherPieceOfMiddleware` middleware before first getting to their final handler.

_NOTE: Middleware defined on an application is automatically inherited by all routes and groups in that application._


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


## Group Middleware

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

## Skipping Middleware

There are times when, in an application, you want to add middleware to the entire application, or a group, but not call that middleware on a few individual handlers. Buffalo allows you to create these sorts of mappings.


{{< codetabs >}}
{{< tab "actions/app.go" >}}
```go
// actions/app.go
a := buffalo.New(buffalo.Options{})
a.Use(AuthorizeUser)

// skip the AuthorizeUser middleware for the NewUser and CreateUser handlers.
a.Middleware.Skip(AuthorizeUser, NewUser, CreateUser)

a.GET("/users/new", NewUser)
a.POST("/users", CreateUser)
a.GET("/users", ListUsers)
a.GET("/users/{id}", ShowUser)
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```text
// OUTPUT
GET /users/new -> NewUser
POST /users -> CreateUser
GET /users -> AuthorizeUser -> ListUsers
GET /users/{id} -> AuthorizeUser -> ShowUser
```
{{< /tab >}}
{{< /codetabs>}}

---

<div class="alert alert-warning" role="alert">
<b>IMPORTANT:</b> The middleware function and the action functions you want to skip <b>MUST</b> be the same Go instance.

{{< codetabs >}}
{{< tab "EXAMPLE 1" >}}
```go
// EXAMPLE 1
m1 := MyMiddleware()
m2 := MyMiddleware()

app.Use(m1)

app.Skip(m2, Foo, Bar) // WON'T WORK m2 != m1
app.Skip(m1, Foo, Bar) // WORKS
```
{{< /tab >}}
{{< tab "EXAMPLE 2" >}}
```go
// EXAMPLE 2
app.Resource("/widgets", WidgetResource{})
app.Skip(mw, WidgetResource{}.Show) // WON'T WORK

wr := WidgetResource{}
app.Resource("/widgets", wr)
app.Skip(mw, wr.Show) // WORKS
```
{{< /tab >}}
{{< /codetabs>}}

</div>

See [https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip) for more details on the `Skip` function.


## Skipping Resource Actions

Often it is necessary to want to skip middleware for one or more actions. For example, allowing guest users to view the `List` and `Show` actions on a resource, but requiring authorization on the rest of the actions.

Understanding from the [Skipping Middleware](#skipping-middleware) section we need to make sure that we are using the same functions when we register the resource as we do when we want to skip the middleware on those functions later.

The line that was generated in `actions/app.go` by `buffalo generate resource` will need to be changed to accommodate this requirement.

{{< codetabs >}}
{{< tab "Before" >}}
```go
// BEFORE
app.Resource("/widgets", WidgetResource{})
```
{{< /tab >}}
{{< tab "After" >}}
```go
// AFTER
res := WidgetResource{}
wr := app.Resource("/widgets", res)
wr.Middleware.Skip(Authorize, res.Index, res.Show)
```
{{< /tab >}}
{{< /codetabs>}}


## Clearing Middleware

Since middleware is [inherited](#using-middleware) from its parent, there maybe times when it is necessary to start with a "blank" set of middleware.

{{< codetabs >}}
{{< tab "actions/app.go" >}}
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
{{< /tab >}}
{{< tab "OUTPUT" >}}
```text
// OUTPUT
GET /foo -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> AuthorizeAPIMiddleware -> UsersHandler
```
{{< /tab >}}
{{< /codetabs>}}


## Listing an Application's Middleware

To get a complete list of the middleware your application is using, broken down by grouping, can be found by running the `buffalo task middleware` command.

```bash
$ buffalo t middleware

-> /
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
-> /courses/{course_slug}
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
github.com/markbates/coke/actions.FindCourse
-> /admin
github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
github.com/gobuffalo/buffalo.*App.PanicHandler
github.com/gobuffalo/buffalo.RequestLoggerFunc
github.com/gobuffalo/buffalo.sessionSaver
github.com/gobuffalo/mw-forcessl.Middleware.func1
github.com/markbates/coke/actions.App.func1
github.com/markbates/coke/actions.trackLastURL
github.com/markbates/coke/actions.TrackingCookie
github.com/markbates/coke/actions.App.func3
github.com/gobuffalo/mw-paramlogger.ParameterLogger
github.com/gobuffalo/mw-csrf.glob..func1
github.com/gobuffalo/buffalo-pop/pop/popmw.Transaction.func2
github.com/markbates/coke/actions.SetCurrentUser
github.com/markbates/coke/actions.SetPageTitle
github.com/markbates/coke/actions.Authorize
github.com/markbates/coke/actions.AuthorizeAdmin
```

