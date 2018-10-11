<%= title("Custom Named Routes") %>

The [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) function allows you to set a custom name for route helpers.

```go
a.GET("/coke", CokeHandler).Name("customPath")
```

This route is now called `customPath` and you can reference it as such in your templates.

```html
&lt;a href="\<%= customPath() %>">Coke&lt;/a>
```

