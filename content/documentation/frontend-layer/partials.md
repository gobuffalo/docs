---
Name: "Partials"
weight: 4
aliases:
  - /docs/partials
  - /en/docs/partials
---

# Partials


{{<note>}}
This document only applies when using [https://github.com/gobuffalo/buffalo/tree/main/render](https://github.com/gobuffalo/buffalo/tree/main/render).
Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.
{{</note>}}


## Usage

You can call your partials using `partial` plush helper:

{{<codetabs>}}
{{<tab "templates/users/form.plush.html">}}
```html
<form action="/users/" method="POST">
<!-- form content here  -->
<form>
```
{{</tab>}}
{{<tab "templates/users/new.plush.html">}}
```html
<h1>Create New User</h1>

<%= partial("users/form.html") %>
```
{{</tab>}}
{{<tab "Output">}}
```html
<h1>Create New User</h1>

<form action="/users/" method="POST">
<!-- form content here  -->
<form>
```
{{</tab>}}
{{</codetabs>}}

## Context

All [rendering context](/documentation/frontend-layer/rendering) from the parent template will automatically pass through to the partial, and any partials that partial may call. (see also [Context](/documentation/request_handling/context))



{{<codetabs>}}
{{<tab "actions/users.go">}}
```go
func UsersEdit(c buffalo.Context) error {
	user := User{
		Name: "John Smith",
	}
	// ...
	c.Set("user", user)
	return c.Render(http.StatusOK, render.HTML("users/edit.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/users/edit.plush.html">}}
```html
<h1>User to edit: <strong><%= user.Name %></strong></h1>

<%= partial("users/form.plush.html") %>
```
{{</tab>}}
{{<tab "templates/users/form.plush.html">}}
```html
<form action="/users/<%= user.ID %>/">
<!-- form content here  -->
</form>
```
{{</tab>}}
{{<tab "Output">}}
```html
<h1>User to edit: <strong>John Smith</strong></h1>

<form action="/users/<%= user.ID %>/">
<!-- form content here  -->
</form>
```
{{</tab>}}
{{</codetabs>}}



## Local Context

In addition to have the global [context](/documentation/request_handling/context), you can set additional variable only for partials as "local" variables.

{{<codetabs>}}
{{<tab "actions/colors.go">}}
```go
func ColorsHandler(c buffalo.Context) error {
  colors := map[string]interface{}{
		"White":  "#FFFFFF",
		"Maroon": "#800000",
		"Red":    "#FF0000",
		"Purple": "#800080",
	}

	c.Set("colors", colors)
	return c.Render(http.StatusOK, r.HTML("colors/index.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/colors/index.plush.html">}}
```html
<div class="list">
  <%= for (name, code) in colors { %>
      <%= partial("colors/details.plush.html", {colorName: name, hexCode: code}) %>
  <% } %>
</div>
```
{{</tab>}}
{{<tab "templates/colors/details.plush.html">}}
```erb
<div>
  <span>Color: <%= colorName %></span>
  <span>Hex Code: <strong><%= hexCode %></strong></span>
</div>
```
{{</tab>}}
{{<tab "Output">}}
```html
<div class="list">
  <div>
    <span>Color: White</span>
    <span>Hex Code: <strong>#FFFFFF</strong></span>
  </div>
  <div>
    <span>Color: Maroon</span>
    <span>Hex Code: <strong>#800000</strong></span>
  </div>
  <div>
    <span>Color: Red</span>
    <span>Hex Code: <strong>#FF0000</strong></span>
  </div>
  <div>
    <span>Color: Purple</span>
    <span>Hex Code: <strong>#800080</strong></span>
  </div>
</div>
```
{{</tab>}}
{{</codetabs>}}

## Helpers

Partials are not much different from standard [templates](/documentation/frontend-layer/templating) in Buffalo. They include all of the same [helpers](/documentation/frontend-layer/helpers) as well.
