# Context

At the heart of every Buffalo request handler is the `Context`. This context gives handlers a simple, and clean, function definition, while being immensely powerful.

<%= title("The Context Interface", {name:"interface"}) %>

The `buffalo.Context` interface supports `context.Context` so it can be passed around and used as a "standard" Go Context.

Since `buffalo.Context` is an interface it is possible to create an application specific implementation that is tailor suited to the needs of the application being built.

```go
type Context interface {
  context.Context
  Response() http.ResponseWriter
  Request() *http.Request
  Session() *Session
  Params() ParamValues
  Param(string) string
  Set(string, interface{})
  LogField(string, interface{})
  LogFields(map[string]interface{})
  Logger() Logger
  Bind(interface{}) error
  Render(int, render.Renderer) error
  Error(int, error) error
  Websocket() (*websocket.Conn, error)
  Redirect(int, string, ...interface{}) error
  Data() map[string]interface{}
}
```

<%= title("Context and Rendering") %>

As part of the context interface, there is a `Render` function that takes a type of `render.Renderer`. See [rendering](/docs/rendering) for more information.

Any values that are "set" on the context will automatically be available to the `render.Renderer` that is passed into the `Render` function.

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, render.String("Hi \<%= name %>"))
}
```

<%= title("Implementing the Interface", {name:"implementing"}) %>

The `buffalo.Context` is never meant to be "fully" implemented. Instead it is recommended that you use composition and implement only the functions that you want to provide custom implementations of.

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

<%= title("Ranging Over Parameters") %>

The `buffalo.Context#Params` method returns [`buffalo.ParamValues`](https://godoc.org/github.com/gobuffalo/buffalo#ParamValues) which is an interface around [`url.Values`](https://golang.org/pkg/net/url/#Values). You can cast to this type in a handler to range over the parameter values.

```go
func HomeHandler(c buffalo.Context) error {
  if m, ok := c.Params().(url.Values); ok {
    for k, v := range m {
      fmt.Println(k, v)
    }
  }
  return c.Render(200, r.HTML("index.html"))
}
```
