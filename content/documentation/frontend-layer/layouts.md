---
Name: "Layouts"
---

# Layouts

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

## Using a Standard Layout

It is quite common to want to use the same layout across most, if not all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

```go
// actions/render.go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout:     "application.html",
    // ...
  })
}
```

```html
// templates/application.html
&lt;html>
  &lt;head>
    &lt;title>My App&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      \<%= yield %>
    &lt;/div>
  &lt;/body>
&lt;/html>
```

```html
// templates/hello.html
&lt;h1>Hello!!&lt;/h1>
```

```go
// actions/hello.go
package actions

func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html"))
}
```

```html
// output
&lt;html>
  &lt;head>
    &lt;title>My App&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      &lt;h1>Hello!!&lt;/h1>
    &lt;/div>
  &lt;/body>
&lt;/html>
```


## Using a Custom Layout

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`. Custom layouts do **NOT** work with `render.Auto`.

```go
// actions/render.go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout:     "application.html",
    // ...
  })
}
```

```html
// templates/custom.html
&lt;html>
  &lt;head>
    &lt;title>My Custom Layout&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      \<%= yield %>
    &lt;/div>
  &lt;/body>
&lt;/html>
```

```html
// templates/hello.html
&lt;h1>Hello!!&lt;/h1>
```

```go
// actions/hello.go
package actions

func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html", "custom.html"))
}
```

```html
// output
&lt;html>
  &lt;head>
    &lt;title>My Custom Layout&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      &lt;h1>Hello!!&lt;/h1>
    &lt;/div>
  &lt;/body>
&lt;/html>
```
