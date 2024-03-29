---
Name: "Layouts"
weight: 3
aliases:
  - /docs/layouts
  - /en/docs/layouts
---

# Layouts

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

## Using a Standard Layout

It is quite common to want to use the same layout across most, if not all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

{{<codetabs>}}
{{<tab "actions/render.go">}}
```go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html",
    // ...
  })
}
```
{{</tab>}}
{{<tab "templates/application.plush.html">}}
```html
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
{{</tab>}}
{{<tab "templates/hello.plush.html">}}
```html
<h1>Hello!!</h1>
```
{{</tab>}}
{{<tab "actions/hello.go">}}
```go
func Hello(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("hello.html"))
}
```
{{</tab>}}
{{<tab "Output">}}
```html
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
{{</tab>}}
{{</codetabs>}}

## Using a Custom Layout

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`.

{{<note>}}
Custom layouts do **NOT** work with **`render.Auto`**.
{{</note>}}

{{<codetabs>}}
{{<tab "actions/render.go">}}
```go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout: "application.plush.html", // You can realize that render continues using the application.plush.html
    // ...
  })
}
```
{{</tab>}}
{{<tab "templates/custom.plush.html">}}
```html
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
{{</tab>}}
{{<tab "templates/hello.plush.html">}}
```html
<h1>Hello!!</h1>
```
{{</tab>}}
{{<tab "actions/hello.go">}}
```go
func Hello(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("hello.plush.html", "custom.plush.html"))
}
```
{{</tab>}}
{{<tab "Output">}}
```html
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
{{</tab>}}
{{</codetabs>}}
