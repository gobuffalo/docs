# Context

At the heart of every Buffalo request handler is the `Context`. This context gives handlers a simple, and clean, function definition, while being immensely powerful.

{{ partial "topics.html" }}

{{#panel title="The Context Interface" name="interface"}}

Since `buffalo.Context` is an interface it is possible to create an application specific implementation that is tailor suited to the needs of the application being built.

```go
type Context interface {
	Response() http.ResponseWriter
	Request() *http.Request
	Session() *Session
	Params() ParamValues
	Param(string) string
	ParamInt(string) (int, error)
	Set(string, interface{})
	Get(string) interface{}
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
{{/panel}}

{{#panel title="Context and Rendering"}}

As part of the context interface is a `Render` function that takes a type of `render.Renderer`, see [rendering](/docs/rendering) for more information.

Any values that are "set" on the context will automatically be available to the `render.Renderer` that is passed into the `Render` function.

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, render.String("Hi {{name}}"))
}
// Hi Mark
```

{{/panel}}
