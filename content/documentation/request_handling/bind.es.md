---
name: Vinculación de Peticiones
seoDescription: "Bind request parameters to Buffalo context"
seoKeywords: ["buffalo", "go", "golang", "bind", "parameter", "context"]
weight: 5
aliases:
  - /docs/bind
  - /es/docs/bind
---

# Vinculación de Peticiones

La interfaz `buffalo.Context` tiene un método llamado `Bind`. Este método permite la vinculación de una petición, como un formulario o un cuerpo JSON, para ser mapeado a un struct para facilitar el manejo del formulario.

Para más información sobre la vinculación de peticiones en Buffalo, consulta el godoc de [github.com/gobuffalo/buffalo/binding](https://godoc.org/github.com/gobuffalo/buffalo/binding)

## Cómo funciona?

La función `Bind` funciona mirando a los headers `Content-Type` o al `Accept` de una petición y buscando una implementación mapeada de la interfaz [`binding.Binder`](https://godoc.org/github.com/gobuffalo/buffalo/binding#Binder).

#### Tipos de contenido Mapeados (HTML)

* `application/html`
* `text/html`
* `application/x-www-form-urlencoded`
* `multipart/form-data`
* `html`

#### Tipos de contenido Mapeados (JSON)

* `application/json`
* `text/json`
* `json`

#### Tipos de contenido Mapeados (XML)

* `application/xml`
* `text/xml`
* `xml`

## Vinculación HTML/Form

La vinculación de formularios HTML, por defecto, utiliza el paquete [github.com/monoculum/formam](https://github.com/monoculum/formam) para vincular los formularios HTML a un struct.

Toma el siguiente struct `User` y el formulario HTML que sigue.

```go
type User struct {
  Name     string `form:"name"`
  Email    string
  Password string `form:"-"`
}
```

```html
<form>
  <input type="text" value="ringo" name="name"/>
  <input type="text" value="ringo@beatles.com" name="Email"/>
  <input type="text" value="starr" name="Password"/>
  <input type="submit"/>
</form>
```

En una acción podemos vincular este formulario HTML al struct `User` de la siguiente manera:

```go
func MyAction(c buffalo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
  u.Name // "Ringo"
  u.Email // "ringo@beatles.com"
  u.Password // ""
  // do more work
}
```

El mapeo por defecto entre el formulario y el struct es el nombre del atributo en el struct, y debe coincidir con el atributo "name" del campo del formulario. Observe que el campo `Email` en los ejemplos coincide tanto con el atributo del struct como con el atributo `name` del campo del formulario.

Usando las etiquetas `form` en los structs, podemos asignar los campos del formulario HTML al struct `User`, incluyendo el ignorar la `Contraseña` usando un `-`.

Consulta [github.com/monoculum/formam](https://github.com/monoculum/formam) para obtener más información sobre la etiqueta `form` en los structs.

## Vinculación JSON y XML


La vinculación de peticiones JSON y XML, por defecto, utiliza los paquetes `encoding/json` y `encoding/xml` de la biblioteca estándar.

Toma el siguiente struct `User` y el JSON de abajo. (XML funciona exactamente igual, pero en lugar de las etiquetas struct `json`, sustituye `xml` en su lugar).

```go
type User struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"-"`
}
```

```json
{
  "name": "Ringo",
  "email": "ringo@beatles.com",
  "password": "starr"
}
```

En una acción podemos vincular este JSON al struct `User` de la siguiente manera:

```go
func MyAction(c buffalo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
  u.Name // "Ringo"
  u.Email // "ringo@beatles.com"
  u.Password // ""
  // do more work
}
```

Usando las etiquetas de struct `json` o `xml` podemos mapear los campos de un JSON al struct `User`, incluyendo el ignorar la `Contraseña` usando un `-`.

Consulta la documentación de la biblioteca estándar para obtener más información sobre las etiquetas de struct `json` y `xml`.

## Registrando un Binder personalizado

Tal vez no quieras usar los binders por defecto, o tienes un tipo de contenido diferente que quieres poder vincular, por ejemplo, digamos que quieres vincular una petición YAML.

```go
package actions

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

func init() {
  binding.Register("text/yaml", func(req *http.Request, model interface{}) error {
    b, err := ioutil.ReadAll(req.Body)
    if err != nil {
      return err
    }
    return yaml.Unmarshal(b, model)
  })
}
```
