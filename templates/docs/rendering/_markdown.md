<%= title("Markdown", {}) %>

Files passed into the `render.HTML` or `render.Template` functions, that have an extension of `.md`, will be converted from Markdown (using GitHub flavored Markdown) to HTML before being run through the templating engine. This makes for incredibly easy templating for simpler pages.

```md
// beatles.md
# The Beatles

\<%= for (name) in names { %>
* \<%= name %>
\<% } %>
```

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(200, r.HTML("beatles.md"))
}
```

```html
// output
&lt;h1>The Beatles</h1>

&lt;ul>
  &lt;li>John</li>
  &lt;li>Paul</li>
  &lt;li>George</li>
  &lt;li>Ringo</li>
&lt;/ul>
```

