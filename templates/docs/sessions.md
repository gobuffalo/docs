# Sessions

The session is available directly from the `buffalo.Context` inside of a handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

<%= partial("docs/sessions/type.md") %>
<%= partial("docs/sessions/store.md") %>
<%= partial("docs/sessions/complex.md") %>
<%= partial("docs/sessions/save.md") %>
<%= partial("docs/sessions/null.md") %>

