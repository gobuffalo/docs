---
Name: "Diseños"
weight: 3
aliases:
  - /docs/layouts
  - /es/docs/layouts
---

# Diseño

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

## Usando un diseño estándar

Es bastante común querer usar el mismo diseño en la mayoria, si no en toda la aplicación. Al crear un nuevo `render.Engine` the se puede establecer la propiedad `HTMLLayout` a un archivo que será usado automáticamente por la función `render.HTML`.

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

## Usando un diseño personalizado

A veces, en ciertas peticiones, se necesita un diseño diferente. Este diseó alternativo se puede pasar como segundo parámetro al `render.HTML`.

{{<note>}}
Los diseños personalizados **NO** funcionan con **`render.Auto`**.
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
