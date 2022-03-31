## Accessing a Form File

In the [`buffalo.Context`](https://godoc.org/github.com/gobuffalo/buffalo#Context) the `c.File` takes a string, the name of the form file parameter and will return a [`binding.File`](https://godoc.org/github.com/gobuffalo/buffalo/binding#File) that can be used to easily retrieve a file from the from.

```go
func SomeHandler(c buffalo.Context) error {
  // ...
  f, err := c.File("someFile")
  if err != nil {
    return errors.WithStack(err)
  }
  // ...
}
```
