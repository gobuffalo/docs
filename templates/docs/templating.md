# Templating

<%= partial("docs/disclaimer.html") %>

<%= partial("topics.html") %>

<%= panel("General Usage", {name: "general"}) { %>

Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

Let's assume you have a template (a string of some kind):

```erb
<!-- templates/index.html -->
<h1>\<%= name %></h1>
<ul>
  \<%= for (name) in names { %>
    <li>\<%= name %></li>
  \<% } %>
</ul>

\<%# "this is a comment and won't get printed" %>
```

Given that string, you can render the template like such:

```go
func IndexHandler(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(200, r.HTML("index.html"))
}
```

Which would result in the following output:

```html
<h1>Mark</h1>
<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
```

<% } %>

<%= panel("If Statements", {name: "if"}) { %>

What to do? Should you render the content, or not? Using plush's built in `if`, `else`, and `unless` helpers let you figure it out for yourself.

```erb
\<%= if (true) { %>
  <!-- render this -->
\<% } %>
```

<% } %>

<%= panel("Else Statements", {name: "else"}) { %>

```erb
\<%= if (false) { %>
  <!-- won't render this -->
\<% } else { %>
  <!-- render this -->
\<% } %>
```

<% } %>
