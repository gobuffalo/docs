# Context

{{ partial "topics.html" }}

{{#panel title="The Context Interface" name="interface"}}
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
