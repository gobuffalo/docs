<%= title("Using Route Helpers in Actions") %>

### Redirecting with Route Helpers

You can also use route names when redirecting to another url.

```go
func MyHandler(c buffalo.Context) error {
  return c.Redirect(307, "widgetsPath()")
  // Or with parameters
  return c.Redirect(307, "widgetPath()", render.Data{"widget_id": "1"})
}
```

---

### Finding/Calling a Route Helper

<%= sinceVersion("0.13.0-beta.1") %>

The [`buffalo.RouteList#Lookup`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList.Lookup) allows you to look up a route by its name from the application. With the `RouteInfo` value for the given route you can generate the path for the route.

```go
func MyHandler(c buffalo.Context) error {
	ri, err := App().Routes().Lookup("widgetPath")
	if err != nil {
		return errors.WithStack(err)
	}
	h := ri.BuildPathHelper()
	u, err := h(render.Data{"widget_id": 1})
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(307, string(u))
}
```

