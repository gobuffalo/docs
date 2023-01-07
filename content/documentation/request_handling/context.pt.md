---
Name: Context
weight: 4
aliases:
  - /pt/docs/context
  - /pt/docs/context
---
# Context

At the heart of every Buffalo request handler is the `Context`. This context gives handlers a simple, and clean, function definition, while being immensely powerful.

## The Context Interface

The `buffalo.Context` interface supports `context.Context` so it can be passed around and used as a "standard" Go Context.

Since `buffalo.Context` is an interface it is possible to create an application specific implementation that is tailored to the needs of the application being built.

{{< since "0.12.0" >}}

```go
type Context interface {
	context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	Session() *Session
	Cookies() *Cookies
	Params() ParamValues
	Param(string) string
	Set(string, interface{})
	LogField(string, interface{})
	LogFields(map[string]interface{})
	Logger() Logger
	Bind(interface{}) error
	Render(int, render.Renderer) error
	Error(int, error) error
	Redirect(int, string, ...interface{}) error
	Data() map[string]interface{}
	Flash() *Flash
	File(string) (binding.File, error)
}
```

The `Websocket() (*websocket.Conn, error)` function was removed from `buffalo.Context` in version `v0.12.0`. Use the [http://www.gorillatoolkit.org/pkg/websocket](http://www.gorillatoolkit.org/pkg/websocket) package directly instead

## Context and Rendering

As part of the context interface, there is a `Render` function that takes a type of `render.Renderer`. See [rendering](/documentation/frontend-layer/rendering) for more information.

Any values that are "set" on the context will automatically be available to the `render.Renderer` that is passed into the `Render` function.

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")

  return c.Render(http.StatusOK, render.String("Hi <%= name %>"))
}
```

## Implementing the Interface

The `buffalo.Context` is never meant to be "fully" implemented. Instead it is recommended that you use [composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) and implement only the functions that you want to provide custom implementations of.

Below is an example of changing the `Error` function to log the error and kill application:

```go
// actions/context.go
type MyContext struct {
  buffalo.Context
}

func (my MyContext) Error(status int, err error) error {
  my.Logger().Fatal(err)
  return err
}
```

```go
// actions/app.go
// ...
func App() *buffalo.App {
  if app != nil {
    // ...
    app.Use(func (next buffalo.Handler) buffalo.Handler {
      return func(c buffalo.Context) error {
      // change the context to MyContext
      return next(MyContext{c})
      }
    })
    // ...
  }
  return app
}
// ...
```

## Ranging Over Parameters

The `buffalo.Context#Params` method returns [`buffalo.ParamValues`](https://godoc.org/github.com/gobuffalo/buffalo#ParamValues) which is an interface around [`url.Values`](https://golang.org/pkg/net/url/#Values). You can cast to this type in a handler to range over the parameter values.

```go
import "net/url"

func HomeHandler(c buffalo.Context) error {
  if m, ok := c.Params().(url.Values); ok {
    for k, v := range m {
      fmt.Println(k, v)
    }
  }

  return c.Render(http.StatusOK, r.HTML("index.html"))
}
```

## What's in the Context

Buffalo stuffs the context of each request with a lot of information that could be useful in your application, such as the `current_route` or the `session`. Below is a list of what Buffalo adds to the context on each request that you can access from in your actions or templates.

| Key             | Type                                                                                 | Usage                                                                                                                                   |
| ---             | ---                                                                                  | ---                                                                                                                                     |
| `app`           | [`*buffalo.App`](https://godoc.org/github.com/gobuffalo/buffalo#App)                | The current Buffalo application that's running.                                                                                         |
| `env`           | `string`                                                                             | The current environment the app is running in. Example: `test`, `development`, `production`                                                   |
| `routes`        | [`buffalo.RouteList`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList)     | A list of all of the routes mapped on the application.                                                                                  |
| `current_route` | [`buffalo.RouteInfo`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo)     | The current route that is being accessed.                                                                                               |
| `current_path`  | `string`                                                                             | The current path being requested. Example: `/users/1/edit`                                                                              |
| `*Path`         | [`RouteHelperFunc`](https://godoc.org/github.com/gobuffalo/buffalo#RouteHelperFunc) | Helpers to create paths based off of mapped routes. Example: `editUserPath`. Run `buffalo task routes` to see a full list for your app. |
| `params`        | `map[string]string`                                                                  | Query parameters for the requested page.                                                                                                |
| `flash`         | `map[string][]string`                                                                | A map of messages set using `buffalo.Context#Flash`.                                                                                    |
| `session`       | [`*buffalo.Session`](https://godoc.org/github.com/gobuffalo/buffalo#Session)        | The current user's session.                                                                                                             |
| `request`       | [`*http.Request`](https://godoc.org/net/http#Request)                               | The current request.                                                                                                                    |
| `tx`            | [`*pop.Connection`](https://godoc.org/github.com/gobuffalo/pop#Connection)          | Only set if using the `github.com/gobuffalo/buffalo/middleware.PopTransaction` middleware (on by default).                              |

See [Helpers](/documentation/frontend-layer/helpers#builtin-helpers) for a list of built-in helper functions available inside of templates.
