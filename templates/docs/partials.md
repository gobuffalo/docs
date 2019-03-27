<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Partials"}) %>
<% } %>

<%= partial("docs/disclaimer.html") %>

## Naming

All partial file names must start with an `_`. For example: `_form.html`. This helps to differentiate partials from other view templates in your application.

```html
// templates/users/new.html
&lt;h1>Create New User&lt;/h1>

\<%= partial("users/form.html") %>
```

```html
// templates/users/_form.html
&lt;form action="/users">
&lt;!-- form stuff here  -->
&lt;form>
```

```html
// output
&lt;h1>Create New User&lt;/h1>

&lt;form action="/users">
&lt;!-- form stuff here  -->
&lt;form>
```


## Context

All [rendering context](/docs/rendering) from the parent template will automatically pass through to the partial, and any partials that partial may call. (see also [context](/docs/context))


```go
// actions/users.go
func UsersEdit(c buffalo.Context) error {
  // do some work to find the user
  c.Set("user", user)
  return c.Render(200, render.HTML("users/edit.html"))
}
```

```html
// templates/users/edit.hml
&lt;h1>Edit \<%= user.Name %> (\<%= user.ID %>)&lt;/h1>

\<%= partial("users/form.html") %>
```

```html
// templates/users/_form.html
&lt;form action="/users/\<%= user.ID %>">
&lt;!-- form stuff here  -->
&lt;/form>
```

```html
// output
&lt;h1>Edit Mark Bates (1)&lt;/h1>

&lt;form action="/users/1">
&lt;!-- form stuff here  -->
&lt;/form>
```


## Local Context

In addition to have the [context](/docs/context) of the parent template, partials can also be sent additional information as "local" variables.

```go
// actions/users.go
func UsersIndex(c buffalo.Context) error {
  c.Set("users", []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"})
  return c.Render(200, r.HTML("users/index.html"))
}
```

```html
// templates/users/index.html
&lt;h1>All Users&lt;/h1>

&lt;ul>
  \<%= for (u) in users { %>
    \<%= partial("users/user.html", {user: u}) %>
  \<% } %>
&lt;/ul>
```

```html
// templates/users/_user.html
&lt;li>\<%= user.Name %>&lt;/li>
```

```html
// output
&lt;h1>All Users&lt;/h1>

&lt;ul>
  &lt;li>John Lennon</li>
  &lt;li>Paul McCartney</li>
  &lt;li>George Harrison</li>
  &lt;li>Ringo Starr</li>
&lt;/ul>
```

## Helpers

Partials are not much different from standard [templates](/docs/templating) in Buffalo. They include all of the same [helpers](/docs/helpers) as well.
