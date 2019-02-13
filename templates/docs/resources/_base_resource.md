## buffalo.BaseResource

There are often occasions, such as JSON APIs where a resource doesn't need to implement endpoints such as `/widget/{widget_id}/edit`.

Because of the way Go interfaces works you would be unable to remove the `Edit(Context) error` function, for example, if you did not need since `WidgetResource` would no longer conform to the `buffalo.Resource` interface.

When a resource is generated it has [`buffalo.BaseResource`](https://godoc.org/github.com/gobuffalo/buffalo#BaseResource) embedded into it.

```go
type Widget struct {
  buffalo.BaseResource
}
```

The `buffalo.BaseResource` has basic implementations for all of the methods required by `buffalo.Resource`. These methods all `404`.

```go
// Edit default implementation. Returns a 404
func (v BaseResource) Edit(c Context) error {
  return c.Error(404, errors.New("resource not implemented"))
}
```

## Video Presentation

<%= vimeo("212302823") %>
