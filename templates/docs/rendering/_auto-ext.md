## Automatic extensions

<%= sinceVersion("0.10.2") %>

You can use HTML, Javascript and Markdown renderers without specifying the file extension:

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  // Render beatles.html
  return c.Render(200, r.HTML("beatles"))
}
```

This works with [partials](/docs/partials) too.

