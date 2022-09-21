---
Name: "Renderizado"
weight: 1
aliases:
  - /docs/renderizado
  - /es/docs/renderizado
---

# Renderizado

El paquete [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render) [[godoc]](https://pkg.go.dev/github.com/gobuffalo/buffalo/render) implementa esta interfaz y tuiene una útil colleccion de tipos de renderización ya definidos. Se recomienda que uses este paquete pero sientete libre y escribe tus propios renderizadores!

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

## Renderizado automático

{{< since "0.11.0" >}}

En varios casos, tendrás que proporcionar el mismo contenido en diferentes formatos: JSON, XML, HTML... Buffalo ofrece una forma sencilla de hacerlo usando una única sentencia.

```go
func Beatles(c buffalo.Context) error {
  members := models.Members{}
  // ...
  return c.Render(http.StatusOK, r.Auto(c, members))
}
```

{{< vimeo 257736901>}}

## JSON y XML

Cuando se renderiza JSON o XML, usando [`render.JSON`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#JSON) o [`render.XML`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#XML), pasas el valor que deseas que se formatee y el formateador adecuado codificará el valor que pasaste y lo escribirá en la respuesta con su correcto `content-type`.

{{<note>}}
**NOTA:** Si ya tienes una cadena que contiene JSON or XML, **NO** uses estos metodos ya que intentará, formatear la cadena a JSON o XML causando respuestas extrañas.
Lo que podrias hacer en su lugar es escribir una funcion de **renderizado personalizado** como se explica en la sección de [Renderizado personalizado](#renderizado-personalizado):
{{</note>}}

```go
// models/user.go

type User struct {
	FirstName string
	LastName  string
	Gender    string
}
```
{{<codetabs>}}
{{<tab "JSON">}}
```go
func MyHandler(c buffalo.Context) error {
  user := models.User{
		FirstName: "John",
		LastName:  "Smith",
		Gender:    "Male",
	}

  return c.Render(http.StatusOK, r.JSON(user))
}
```

```json
// output
{
  "FirstName": "John",
  "LastName": "Smith",
  "Gender": "Male"
}
```
{{</tab>}}
{{<tab "XML">}}
```go
func MyHandler(c buffalo.Context) error {
  user := models.User{
		FirstName: "John",
		LastName:  "Smith",
		Gender:    "Male",
	}

  return c.Render(http.StatusOK, r.XML(user))
}
```

```xml
<!-- output -->
<User>
  <FirstName>John</FirstName>
  <LastName>Smith</LastName>
  <Gender>Male</Gender>
</User>
```
{{</tab>}}
{{</codetabs>}}

## Markdown

Los archivos pasados a los métodos [`render.HTML`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.HTML) or [`render.Template`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.Template), que tienen la extension `plush.md`, serán convertidos de Markdown (usando GitHub flavored Markdown) a HTML antes de ser ejecutados a través del motor de plantillas. Esto facilita increíblemente la creación de plantillas para páginas más sencillas.

```md
<!-- beatles.plush.md -->

# The Beatles

<%= for (name) in names { %>
* <%= name %>
<% } %>
```

```go
// actions/beatles.go

func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})

  return c.Render(http.StatusOK, r.HTML("beatles.plush.md"))
}
```

```html
<!-- output -->
<h1>The Beatles</h1>

<ul>
  <li><p>John</p></li>
  <li><p>Paul</p></li>
  <li><p>George</p></li>
  <li><p>Ringo</p></li>
</ul>
```

## JavaScript
{{< since "0.10.0" >}}

El paquete [`render`](https://godoc.org/github.com/gobuffalo/buffalo/render) tiene una nueva implementacion de [`render.Renderer`](https://godoc.org/github.com/gobuffalo/buffalo/render#Renderer), [`render.JavaScript`](https://godoc.org/github.com/gobuffalo/buffalo/render#JavaScript).

Esto significa que dentro de una acción puedes hacer lo siguiente:

```go
func HomeHandler(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.JavaScript("index.js"))
}
```

El tipo [`render.Options`](https://godoc.org/github.com/gobuffalo/buffalo/render#Options) tiene ahora un nuevo atributo, `JavaScriptLayout`. Esta nueva opción es similar a la opción `HTMLLayout` en el sentido en que envolverá archivos `*.js` dentro de otro `*.js`.

El nuevo renderizador de JavaScript también tiene su propia implementación de la función `partial`. Esta nueva implementación funciona de manera similar a la implementación original, pero es lo suficientemente inteligente como para saber si estás renderizando un archivo `*.html` dentro de un archivo `*.js` que necesitará ser escapado apropiadamente, y así lo hace por ti.

```javascript
$("#new-goal-form").replaceWith("<%= partial("goals/new.html") %>");
```

## Extensiones automáticas

{{< since "0.10.2" >}}

Puedes usar los renderizadores `HTML`, `JavaScript` y `Markdown` sin especificar la extensión del archivo:

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  // Render beatles.html
  return c.Render(http.StatusOK, r.HTML("beatles"))
}
```

{{<note>}}
Tambien funciona con [partials](/documentation/frontend-layer/partials).
{{</note>}}

## Descargar archivos

El método [`r.Download`](https://pkg.go.dev/github.com/gobuffalo/buffalo/render#Engine.Download) te permite descargar archivos en tu aplicacion fácilmente.

```go
func DownloadHandler(c buffalo.Context) error {
	// ...
	f, err := os.Open("your/path/file_name.extension")
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.Download(c, "file_name.extension", f))
}
```

## Renderizado personalizado

Para otro tipo de renderizado, el método [`r.Func`](https://godoc.org/github.com/gobuffalo/buffalo/render#Func) te permitirá pasar un tipo de contenido y una funcion para renderizar tus datos al `io.Writer` proporcionado, que es comúnmente, la respuesta HTTP, en particular, un [`*buffalo.Response`](https://godoc.org/github.com/gobuffalo/buffalo#Response).

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.Func("application/csv", csvWriter))
}

func csvWriter(w io.Writer, d render.Data) error {
  cw := csv.NewWriter(w)
  if err := cw.Write([]string{"a", "b", "c"}); err != nil {
    return errors.WithStack(err)
  }
  cw.Flush()
  return nil
}
```

Para situaciones más pequeñas, o únicas, usar la función anónima cuede ser incluso mas fácil.
En este ejemplo puedes ver como usar una función anónima para renderizar una cadena que ya contiene JSON.

```go
var myJSONString string
return c.Render(http.StatusOK, r.Func("application/json", func(w io.Writer, d render.Data) error {
  _, err := w.Write([]byte(myJSONString))
  return err
}))
```

## La interfaz Renderer

Para que un renderizador pueda ser usado con [`Context#Render`](/documentation/request_handling/context#context-and-rendering), este debe implementar la siguiente interfaz:

```go
// Renderer interface that must be satisified to be used with
// buffalo.Context.Render
type Renderer interface {
  ContentType() string
  Render(io.Writer, Data) error
}

// Data type to be provided to the Render function on the
// Renderer interface.

type Data map[string]interface{}
```

El paquete [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render) [[godoc]](https://pkg.go.dev/github.com/gobuffalo/buffalo/render) implementa esta interfaz y tuiene una útil colleccion de tipos de renderización ya definidos. Se recomienda que uses este paquete pero sientete libre y escribe tus propios renderizadores!
