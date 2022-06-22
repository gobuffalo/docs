---
Name: "Templating"
weight: 2
aliases:
  - /docs/templating
  - /en/docs/templating
---

# Templating

{{<note>}}
This document only applies when using [https://github.com/gobuffalo/buffalo/tree/main/render](https://github.com/gobuffalo/buffalo/tree/main/render).
Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.
{{</note>}}


Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

## Introduction to Plush
{{< vimeo 207200621>}}

## Plush - Tips, Tricks and Testing

{{< vimeo 267643588>}}

## General Usage

Plush allows you to capture the `context` variables to use anywhere in your templates.

{{<codetabs>}}
{{<tab "actions/index.go">}}
```go
func myHandler(c buffalo.Context) error {
  c.Set("name", "John Smith")
  return c.Render(http.StatusOK, r.HTML("index.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/index.plush.html">}}
```erb
<h1><%= name %></h1>
```
{{</tab>}}
{{<tab "Output">}}
```html
<h1>John Smith</h1>
```
{{</tab>}}
{{</codetabs>}}

## Plush Examples

#### Conditional Statements

{{<codetabs>}}
{{<tab "IF">}}
```erb
<%= if (true) { %>
  <!-- some template content -->
<% } %>
```
{{</tab>}}
{{<tab "ELSE">}}
```erb
<%= if (true) { %>
  <!-- content when statement is true -->
<% } else { %>
  <!-- content when statement is false -->
<% } %>
```
{{</tab>}}
{{<tab "ELSE IF">}}
```erb
<%= if (value == 0) { %>
  <!-- content when value is 0 -->
<% } else if (value == 1) { %>
  <!-- content when value is 1 -->
<% } else { %>
  <!-- content when value is different to 0 and 1 -->
<% } %>
```
{{</tab>}}
{{<tab "Multiple Conditions">}}
```erb
<%= if ((value > 0) && (value < 10)) { %>
  <!-- some template content -->
<% } else { %>
  <!-- some template content -->
<% } %>
```
{{</tab>}}
{{</codetabs>}}

{{<codetabs>}}
{{<tab "actions/main.go">}}
```go
func MyHandler(c buffalo.Context) error {
	// ...
	c.Set("userName", "John Smith")
	return c.Render(http.StatusOK, r.HTML("templates/index.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/index.plush.html">}}
```erb
<%= if (userName != "") { %>
  <span>Welcome <strong><%= userName %>!</strong></span>
<% } else { %>
  <span>Welcome Visitor</span>
<% } %>
```
{{</tab>}}
{{<tab "Output">}}
```html
<span>Welcome <strong>John Smith!</strong></span>
```
{{</tab>}}
{{</codetabs>}}

### Iterating

#### Through Slices

When looping through `slices`, the block being looped through will have access to the "global" context.

The `for` statement takes 1 or 2 arguments. When using the two arguments version, the first argument is the "index" of the loop and the second argument is the value from the slice.

{{<codetabs>}}
{{<tab "actions/main.go">}}
```go
func MyHandler(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(http.StatusOK, r.HTML("index.plush.html"))
}
```
{{</tab>}}
{{<tab "Loop using 2 Arguments">}}
```erb
<!-- templates/index.plush.html -->
<ul>
  <%= for (index, name) in names { %>
    <li><%= index %> - <%= name %></li>
  <% } %>
</ul>
```

```html
<!-- Output -->
<ul>
  <li>0 - John</li>
  <li>1 - Paul</li>
  <li>2 - George</li>
  <li>3 - Ringo</li>
</ul>
```
{{</tab>}}
{{<tab "Loop using 1 Arguments">}}
```erb
<!-- templates/index.plush.html -->
<ul>
  <%= for (name) in names { %>
    <li><%= name %></li>
  <% } %>
</ul>
```

```html
<!-- Output -->
<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
```
{{</tab>}}
{{</codetabs>}}

#### Through Maps

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

When using the two argument version, the first argument is the key of the map and the second argument is the value from the map:

{{<codetabs>}}
{{<tab "actions/main.go">}}
```go
func ColorsHandler(c buffalo.Context) error {
	colors := map[string]interface{}{
		"White":  "#FFFFFF",
		"Maroon": "#800000",
		"Red":    "#FF0000",
		"Purple": "#800080",
	}

	c.Set("colors", colors)
	return c.Render(http.StatusOK, r.HTML("home/colors.plush.html"))
}
```
{{</tab>}}
{{<tab "Loop using 2 Arguments">}}
```erb
<!-- templates/index.plush.html -->
<div>
  <%= for (name, code) in colors { %>
    <span><%= name %>: <%= code %></span>
  <% } %>
</div>
```
```html
<!-- Output -->
<div>
    <span>White:  #FFFFFF</span>
    <span>Maroon: #800000</span>
    <span>Red:    #FF0000</span>
    <span>Purple: #800080</span>
</div>
```
{{</tab>}}
{{<tab "Loop using 1 Arguments">}}
```erb
<!-- templates/index.plush.html -->
<div>
  Color codes:
  <%= for (code) in colors { %>
    <span><%= code %></span>
  <% } %>
</div>
```
```html
<!-- Output -->
<div>
    Color codes:
    <span>#FFFFFF</span>
    <span>#800000</span>
    <span>#FF0000</span>
    <span>#800080</span>
</div>
```
{{</tab>}}
{{</codetabs>}}

{{<note>}}
You can see more examples in [plush repository](https://github.com/gobuffalo/plush).
{{</note>}}