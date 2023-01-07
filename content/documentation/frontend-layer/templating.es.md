---
Name: "Plantillas"
weight: 2
aliases:
  - /docs/plantillas
  - /es/docs/plantillas
---

# Templating

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para mas detalles sobre el paquete de plantillas.
{{</note>}}

Buffalo utiliza por defecto [plush](https://github.com/gobuffalo/plush) as its template engine.

## Introducción a Plush
{{< vimeo 207200621>}}

## Plush - Consejos, trucos y pruebas

{{< vimeo 267643588>}}

## Uso general

Plush te permite capturar las variables de `contexto` para usarlas en cualquier lugar en tus plantillas.

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

## Plush Ejemplos

#### Condicionales

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

### Iteraciones

#### A través de Slices

Cuando recorremos a través de `slices`, el bloque en el que se realiza el bucle tendrá acceso al contexto "global".

La sentencia `for` toma 1 o 2 argumentos. Cuando se usa la versión de 2 argumentos, el primer argumento es el "indice" del bucle y el segundo argumento es el valor del elemento del slice.

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

#### A través de Mapas

El buble a través de `mapas` usando el helper `each` tambien está soportado, y sigue directrices similares al bucle a través de `slices`.

Cuando se usa la version de 2 argumentos, el primer argumento es la llave del `mapa` y el segundo argumento es el valor del elemento en el `mapa`.

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
Puedes ver más ejemplos en el [repositorio plush](https://github.com/gobuffalo/plush).
{{</note>}}
