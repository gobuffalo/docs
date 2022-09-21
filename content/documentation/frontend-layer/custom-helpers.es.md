---
Name: "Helpers Personalizados"
weight: 7
aliases:
  - /docs/custom-helpers
  - /es/docs/custom-helpers
---

# Helpers personalizados

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

Ningún paquete de plantillas estaría completo sim permitirte construir tus propias funciónes helper personalizadas.

{{< vimeo 229572343>}}

## Registrar Helpers

Las funciónes helpers se pueden registrar en dos diferentes lugares, dependiendo de como se vayan a usar.

### Helper globales

*La mayoría* de helpers seran helpers globales, esto quiere decir que se podrán incluir en cualquier plantilla. Los tipos de helpers pueden ser establecidos en `actions/render.go`:

```go
func init() {
  r = render.New(render.Options{
    // ...
    Helpers: render.Helpers{
      "myHelper": func() string {
        return "hello"
      },
    },
    // ...
  })
}
```

### Helpers por petición

Otros helpers, que son específicos a una determinada petición se pueden agregar al `buffalo.Context` de esa petición.

```go
func HomeHandler(c buffalo.Context) error {
  // ...
  c.Set("myHelper", func() string {
    return "hello"
  })
  // ...
}
```

## Valores retornados

Plush te permite retornar cualquier valor que desees de una función helper. Esta guía se enfocará en los helpers que estan diseñados para generar "salidas".

Cuando retornamos multiples valores de una función, el primer valor será el que se use para renderizar en la plantilla.
Si el ultimo valor retornado es un `error`, Plush manejará ese error.

---

#### `string`

Retornar sólo un `string`. El `string` será escapado de HTML y considerada "no" seguro.

```go
func() string {
  return ""
}
```

---

#### `template.HTML`

[https://golang.org/pkg/html/template/#HTML](https://golang.org/pkg/html/template/#HTML)

Retornar sólo una cadena de `template.HTML`. El `template.HTML` **no** será escapado de HTML y será considerado seguro.

```go
func() template.HTML {
  return template.HTML("")
}
```

## Valores de entrada

Las funciónes helper personalizadas pueden tomar cualquier tipo, y cualquier numero de argumentos. Incluso puede utilizar funciónes variádicas. Hay un ultimo argumento opcional: [`plush.HelperContext`](https://godoc.org/github.com/gobuffalo/plush#HelperContext), que se puede recibir. Es bastante útil, y te recomendaría tomarlo, ya que te proporciona acceso a cosas como el contexto de la llamada, el bloque asociado al helper, etc...

## Helpers simples

```go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "greet": func(name string) string {
      return fmt.Sprintf("Hi %s!", name)
    },
  },
  // ...
})
```
La función `greet` ahora está disponible para todas las plantillas que usen ese motor `render.Engine`.

```go
// actions/greet.go
func Greeter(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(http.StatusOK, r.String("<h1><%= greet(name) %></h1>"))
}
```

```go
// output
<h1>Hi Mark!</h1>
```

## Helpers en bloque

Como las sentencias `if` o `for`, los helpers en bloques toman un "bloque" de texto que se puede evaluar y potencialmente renderizar, manipular o lo que quisieras.
Para escribir un helper en bloque, tienes que tomar el `plush.HelperContext` como el último argumento de tu función helper. Esto te dará acceso al bloque asociado con esa llamada.

{{<codetabs>}}
{{<tab "actions/render.go" >}}
```go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "upblock": upblock,
  },
  // ...
})
```
{{</tab>}}
{{<tab "helper" >}}
```go
func upblock(help plush.HelperContext) (template.HTML, error) {
  s, err := help.Block()
  if err != nil {
    return "", err
  }
  return strings.ToUpper(s), nil
}
```
{{</tab>}}
{{<tab "actions/upper.go" >}}
```go
func Upper(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("up.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/up.plush.html" >}}
```html
<%= upblock() { %>
  hello world
<% } %>
```
{{</tab>}}

{{<tab "Output" >}}
```text
HELLO WORLD
```
{{</tab>}}
{{</codetabs>}}


## Obteniendo valores del contexto

{{<codetabs>}}
{{<tab "actions/render.go" >}}
```go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "is_logged_in": isLoggedIn,
  },
  // ...
})
```
{{</tab>}}
{{<tab "helper" >}}
```go
func isLoggedIn(help plush.HelperContext) bool {
  return help.Value("current_user") != nil
}
```
{{</tab>}}
{{<tab "actions/users.go" >}}
```go
func Show(c buffalo.Context) error {
  c.Set("current_user", models.User{Name: "Ringo"})
  return c.Render(http.StatusOK, r.HTML("users/show.plush.html"))
}
```
{{</tab>}}
{{<tab "templates/users/show.plush.html" >}}
```html
<%= if (is_logged_in()) { %>
  Hello <%= current_user.Name %>
<% } %>
```
{{</tab>}}
{{<tab "Output" >}}
```text
// output
Hello Ringo
```
{{</tab>}}
{{</codetabs>}}
