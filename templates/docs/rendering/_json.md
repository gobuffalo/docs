<%= title("JSON and XML", {name: "json-xml"}) %>

When rendering JSON, or XML, using the [`r.JSON`](https://godoc.org/github.com/gobuffalo/buffalo/render#JSON) or [`r.XML`](https://godoc.org/github.com/gobuffalo/buffalo/render#XML), you pass the value that you would like to be marshaled and the appropriate marshaler will encode the value you passed and write it to the response with the correct content/type.

**NOTE**: If you already have a string that contains JSON or XML do **NOT** use these methods as they will attempt to marshal the string into JSON or XML causing strange responses.

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(200, r.JSON(User{}))
}
```

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(200, r.XML(User{}))
}
```

