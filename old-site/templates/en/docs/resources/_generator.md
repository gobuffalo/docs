## Generating Resources

The `buffalo generate resource` command will generate the necessary models, migrations, Go code, and HTML to CRUD the resource.

When running the generator in an API application Buffalo will generate code to meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface.

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

When running the generator in a Web application Buffalo will generate code to the meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface, as well as the optional `New` and `Edit` methods.

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  New(Context) error
  Create(Context) error
  Edit(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

<%= partial("en/docs/resources/example.md") %>

