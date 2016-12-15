# Partials

{{ partial "docs/disclaimer.html" }}

{{ partial "topics.html" }}

{{#panel title="Naming"}}
All partial file names must start with an `_`. For example: `_form.html`. This helps to different partials from other view templates in your application.

**templates/users/new.html**
```html
<h1>Create New User</h1>

\{{ partial "users/form.html" }}
```

**templates/users/_form.html**
```html
<form action="/users">
<!-- form stuff here  -->
</form>
```

**output**
```html
<h1>Create New User</h1>

<form action="/users">
<!-- form stuff here  -->
</form>
```
{{/panel}}

{{#panel title="Context"}}

All [rendering context](/docs/rendering) from parent template will automatically pass through to the partial, and any partials that partial may call. (see also [context](/docs/context))

**actions/users.go**
```html
func UsersEdit(c buffalo.Context) error {
  // do some work to find the user
  c.Set("user", user)
  return c.Render(200, render.HTML("users/edit.html"))
}
```

**templates/users/edit.html**
```html
<h1>Edit \{{user.Name}} (\{{user.ID}})</h1>

\{{ partial "users/form.html" }}
```

**templates/users/_form.html**
```html
<form action="/users/\{{user.ID}}">
<!-- form stuff here  -->
</form>
```

**output**
```html
<h1>Edit Mark Bates (1)</h1>

<form action="/users/1">
<!-- form stuff here  -->
</form>
```
{{/panel}}

{{#panel title="\"Local\" Context" name="local-context"}}

In addition to have the [context](/docs/context) of the parent template, partials can also be sent additional information as "local" variables.

**templates/users/index.html**
```html
<h1>All Users</h1>

<ul>
  \{{#each users as |user|}}
    \{{ partial "users/user.html" user=user }}
  \{{/each}}
</ul>
```

**templates/users/_user.html**
```html
<li>\{{user.Name}}</li>
```

**outputs**
```html
<h1>All Users</h1>

<ul>
  <li>John Lennon</li>
  <li>Paul McCartney</li>
  <li>George Harrison</li>
  <li>Ringo Starr</li>
</ul>
```

{{/panel}}

{{#panel title="Helpers"}}

Partials are not much different from standard [templates](/docs/templating) in Buffalo. They include all of the same [helpers](/docs/helpers) as well.

{{/panel}}
