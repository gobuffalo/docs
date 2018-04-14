<%= title("Render Auto", {}) %>

<%= sinceVersion("0.11.0") %>

In many cases, you'll have to provide the same contents in different formats: JSON, XML, HTML... Buffalo provides an easy way to do that using a single statement.

```go
func Beatles(c buffalo.Context) error {
  members := models.Members{}
  // ...
  return c.Render(200, r.Auto(c, members))
}
```

<%= vimeo("257736901") %>
