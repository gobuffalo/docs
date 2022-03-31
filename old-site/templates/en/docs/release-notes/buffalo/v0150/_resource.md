## Resource Name Available on Route Info

The [`github.com/gobuffalo/buffalo#RouteInfo`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo) type now contains the name of the resource it belongs to, if at all.

```go
app.Resource("/widget", WidgetsResource{})
```

```html
{
  "method": "GET",
  "path": "/widgets/",
  "handler": "github.com/markbates/coke/actions.WidgetsResource.List",
  "resourceName": "WidgetsResource",
  "pathName": "widgetsPath",
  "aliases": []
}
```

* [https://github.com/gobuffalo/buffalo/pull/1798](https://github.com/gobuffalo/buffalo/pull/1798)
