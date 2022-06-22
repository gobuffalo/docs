---
Name: "Layouts"
weight: 3
aliases:
  - /docs/layouts
  - /en/docs/layouts
---

# Layouts

{{<note>}}
This document only applies when using [https://github.com/gobuffalo/buffalo/tree/main/render](https://github.com/gobuffalo/buffalo/tree/main/render).
Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.
{{</note>}}

## Using a Standard Layout

It is quite common to want to use the same layout across most, if not all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

```go
// actions/render.go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html",
    // ...
  })
}
```

```html
<!-- templates/application.plush.html -->
<html>
  <head>
    <title>My App</title>
  </head>
  <body>
    <div id="main">
      <%= yield %>
    </div>
  </body>
</html>
```

```html
<!-- templates/hello.plush.html -->

<h1>Hello!!</h1>
```

```go
// actions/hello.go

func Hello(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("hello.html"))
}
```

```html
<!-- Output -->
<html>
  <head>
    <title>My App</title>
  </head>
  <body>
    <div id="main">
      <h1>Hello!!</h1>
    </div>
  </body>
</html>
```

## Using a Custom Layout

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`.

{{<note>}}
Custom layouts do **NOT** work with **`render.Auto`**.
{{</note>}}

```go
// actions/render.go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html", // You can realize that render continues using the application.plush.html
    // ...
  })
}
```

```html
<!-- templates/custom.plush.html -->
<html>
  <head>
    <title>My Custom Layout</title>
  </head>
  <body>
    <div id="main">
      <%= yield %>
    </div>
  </body>
</html>
```

```html
<!-- templates/hello.plush.html -->
<h1>Hello!!</h1>
```

```go
// actions/hello.go
func Hello(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("hello.plush.html", "custom.plush.html"))
}
```

```html
<!-- Output -->
<html>
  <head>
    <title>My Custom Layout</title>
  </head>
  <body>
    <div id="main">
      <h1>Hello!!</h1>
    </div>
  </body>
</html>
```
