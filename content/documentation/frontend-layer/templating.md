---
Name: "Templating"
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

## If Statements

```html
\<%= if (true) { %>
  &lt;!-- render this -->
\<% } %>
```



## Else Statements

```html
\<%= if (false) { %>
  &lt;!-- won't render this -->
\<% } else { %>
  &lt;!-- render this -->
\<% } %>
```
