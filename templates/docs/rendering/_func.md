<%= title("Custom Rendering") %>

The [`r.Func`](https://godoc.org/github.com/gobuffalo/buffalo/render#Func) method allows you to pass in a content type and a function to render your data to the provided `io.Writer`, which is commonly, the HTTP response, in particular, a [`*buffalo.Response`](https://godoc.org/github.com/gobuffalo/buffalo#Response).

```go
func MyHandler(c buffalo.Context) error {
	return c.Render(200, r.Func("application/csv", csvWriter))
}

func csvWriter(w io.Writer, d render.Data) error {
	cw := csv.NewWriter(w)
	if err := cw.Write([]string{"a", "b", "c"}); err != nil {
		return errors.WithStack(err)
	}
	cw.Flush()
	return nil
}
```

Or it could be something as simple as
```go
return c.Render(200, r.Func("application/json", func(w io.Writer, d render.Data) error {
	_, err := w.Write([]byte(myJSONString))
	return err
}))
```
