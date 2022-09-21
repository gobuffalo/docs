---
Name: "Parciales"
weight: 4
aliases:
  - /docs/partials
  - /es/docs/partials
---

# Parciales

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

## Uso

Puedes llamar tus parciales usando el helper de plush `partial`:

{{<codetabs>}}
{{<tab "templates/users/new.plush.html">}}
```html
<h1>Create New User</h1>

<%= partial("users/form.html") %>
```
{{</tab>}}
{{<tab "templates/users/form.plush.html">}}
```html
<form action="/users/" method="POST">
<!-- form content here  -->
<form>
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

## Contexto

Todo el [contexto de renderizado](/documentation/frontend-layer/rendering) de la plantilla padre se pasará automáticamente a través del parcial, y a cualquier parcial que puedas llamar. (Ve también [Context](/documentation/request_handling/context))

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

<form action="/users/077acb8d-6cf7-4e7c-ba6c-c60e58ea5fcb/">
<!-- form content here  -->
</form>
```
{{</tab>}}
{{</codetabs>}}



## Contexto local

Además de tener el [context](/documentation/request_handling/context) global, puedes establecer variables adicionales sólo para los parciales como variables "locales".

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

Las parciales no son muy diferentes de las [plantillas](/documentación/capa frontal/plantillas) estándar en Buffalo. También incluyen todos los mismos [helpers](/documentation/frontend-layer/helpers) también.
