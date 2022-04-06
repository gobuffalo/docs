---
Name: "Partials"
weight: 4
aliases:
  - /docs/partials
  - /en/docs/partials
---

# Partials

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

## Naming

All partial file names must start with an `_`. For example: `_form.html`. This helps to differentiate partials from other view templates in your application.

```erb
// templates/users/new.html
<h1>Create New User</h1>

<%= partial("users/form.html") %>
```

```html
// templates/users/_form.html
<form action="/users">
<!-- form stuff here  -->
<form>
```

```html
// output
<h1>Create New User</h1>

<form action="/users">
<!-- form stuff here  -->
<form>
```


## Context

All [rendering context](/documentation/frontend-layer/rendering) from the parent template will automatically pass through to the partial, and any partials that partial may call. (see also [Context](/documentation/request_handling/context))


```go
// actions/users.go
func UsersEdit(c buffalo.Context) error {
  // do some work to find the user
  c.Set("user", user)
  return c.Render(200, render.HTML("users/edit.html"))
}
```

```erb
// templates/users/edit.hml
<h1>Edit <%= user.Name %> (<%= user.ID %>)</h1>

<%= partial("users/form.html") %>
```

```html
// templates/users/_form.html
<form action="/users/<%= user.ID %>">
<!-- form stuff here  -->
</form>
```

```html
// output
<h1>Edit Mark Bates (1)</h1>

<form action="/users/1">
<!-- form stuff here  -->
</form>
```


## Local Context

In addition to have the [context](/documentation/request_handling/context) of the parent template, partials can also be sent additional information as "local" variables.

```go
// actions/users.go
func UsersIndex(c buffalo.Context) error {
  c.Set("users", []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"})
  return c.Render(200, r.HTML("users/index.html"))
}
```

```erb
// templates/users/index.html
<h1>All Users</h1>

<ul>
  <%= for (u) in users { %>
    <%= partial("users/user.html", {user: u}) %>
  <% } %>
</ul>
```

```erb
// templates/users/_user.html
<li><%= user.Name %></li>
```

```html
// output
<h1>All Users</h1>

<ul>
  <li>John Lennon</li>
  <li>Paul McCartney</li>
  <li>George Harrison</li>
  <li>Ringo Starr</li>
</ul>
```

## Helpers

Partials are not much different from standard [templates](/documentation/frontend-layer/templating) in Buffalo. They include all of the same [helpers](/documentation/frontend-layer/helpers) as well.
