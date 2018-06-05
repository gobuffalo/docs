# Templating

<%= partial("docs/disclaimer.html") %>

Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

<%= vimeo("207200621") %>

<%= title("General Usage", {name: "general"}) %>

```html
// templates/index.html
&lt;h1>\<%= name %></h1>
&lt;ul>
  \<%= for (name) in names { %>
    &lt;li>\<%= name %></li>
  \<% } %>
&lt;/ul>
```

```go
// actions/index.go
func IndexHandler(c buffalo.Context) error {
  c.Set("name", "Mark")
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(200, r.HTML("index.html"))
}
```

```html
// output
&lt;h1>Mark</h1>
&lt;ul>
  &lt;li>John</li>
  &lt;li>Paul</li>
  &lt;li>George</li>
  &lt;li>Ringo</li>
&lt;/ul>
```

<%= title("If Statements", {name: "if"}) %>

```html
\<%= if (true) { %>
  &lt;!-- render this -->
\<% } %>
```



<%= title("Else Statements", {name: "else"}) %>

```html
\<%= if (false) { %>
  &lt;!-- won't render this -->
\<% } else { %>
  &lt;!-- render this -->
\<% } %>
```
