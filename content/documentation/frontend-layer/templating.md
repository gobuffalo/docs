---
Name: "Templating"
weight: 2
aliases:
  - /docs/templating
  - /en/docs/templating
---

# Templating

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

<div class="row">
  <div class="col-md-6">
    {{< vimeo 207200621>}}
  </div>
  <div class="col-md-6">
    {{< vimeo 267643588>}}
  </div>
</div>

## General Usage

```erb
// templates/index.html
<h1><%= name %></h1>
<ul>
  <%= for (name) in names { %>
    <li><%= name %></li>
  <% } %>
</ul>
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
<h1>Mark</h1>
<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
```

## If Statements

```erb
<%= if (true) { %>
  <!-- render this -->
<% } %>
```



## Else Statements

```erb
<%= if (false) { %>
  <!-- won't render this -->
<% } else { %>
  <!-- render this -->
<% } %>
```
