---
name: Request Binding
seoDescription: "Bind request parameters to Buffalo context"
seoKeywords: ["buffalo", "go", "golang", "bind", "parameter", "context"]
weight: 5
aliases:
  - /docs/bind
  - /en/docs/bind
---

# Request Binding

The `buffalo.Context` interface has a method named, `Bind`. This method allows for the binding of a request, such as a form or JSON body, to be mapped to a struct for easy form handling.

For more information on request binding in Buffalo, see the godoc for [github.com/gobuffalo/buffalo/binding](https://godoc.org/github.com/gobuffalo/buffalo/binding)

## How Does It Work?

The `Bind` function works by looking at the `Content-Type` or `Accept` header of a request and looking for a mapped implementation of the [`binding.Binder`](https://godoc.org/github.com/gobuffalo/buffalo/binding#Binder) interface.

##### Mapped Content Types (HTML)

* `application/html`
* `text/html`
* `application/x-www-form-urlencoded`
* `multipart/form-data`
* `html`

##### Mapped Content Types (JSON)

* `application/json`
* `text/json`
* `json`

##### Mapped Content Types (XML)

* `application/xml`
* `text/xml`
* `xml`

## HTML/Form Binding

Binding HTML forms, by default, uses the [github.com/monoculum/formam](https://github.com/monoculum/formam) package to bind HTML forms to a struct.

Take the following `User` struct and the HTML form below.

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

In an action we can bind this HTML form to the `User` struct as follows:

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

The default mapping between the form and struct is the name of the attribute on the struct, and should match the `name` attribute of the form field. Notice the `Email` field in the examples matches both the struct attribute and the `name` attribute on the form field.

By using the `form` struct tags we can map the fields in the HTML form to the `User` struct, including ignoring the `Password` by using a `-`.

Please refer to the [github.com/monoculum/formam](https://github.com/monoculum/formam) docs for more information about the `form` struct tag.

## JSON and XML Binding


Binding JSON and XML requests, by default, uses the `encoding/json` and `encoding/xml` packages in the standard library.

Take the following `User` struct and the JSON payload below. (XML works exactly the same, but instead of `json` struct tags, substitute `xml` instead.)

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

In an action we can bind this JSON payload to the `User` struct as follows:

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

By using the `json` or `xml` struct tags we can map the fields in the JSON payload to the `User` struct, including ignoring the `Password` by using a `-`.

Please refer to the standard library docs for more information about the `json` and `xml` struct tags.

## Registering a Custom Binder

Perhaps you don't want to use the default binders, or you have a different content type you want to be able to bind to, for example, let's say you want to bind a YAML request.

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
