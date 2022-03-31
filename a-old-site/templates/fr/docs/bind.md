<% seoDescription("Associer les données d'une requête à une struct") %>
<% seoKeywords(["buffalo", "go", "golang", "bind", "paramètres", "contexte"]) %>

# Associer les données d'une requête à une struct

L'interface `buffalo.Context` a une méthode `Bind`. Cette méthode permet d'associer les paramètres d'une requête (comme un formulaire ou du contenu JSON) à une structure, pour les traiter plus facilement.

Pour plus d'informations sur cette fonctionnalité, vous pouvez consulter la Godoc pour [github.com/gobuffalo/buffalo/binding](https://godoc.org/github.com/gobuffalo/buffalo/binding)

## Comment ça marche ?

La méthode `Bind` utilise l'en-tête HTTP `Content-Type` ou `Accept` d'une requête, et elle cherche une association configurée entre ce `Content-Type` et une implémentation de l'interface [`binding.Binder`](https://godoc.org/github.com/gobuffalo/buffalo/binding#Binder).

##### Content-type pré-associé (HTML)

* `application/html`
* `text/html`
* `application/x-www-form-urlencoded`
* `multipart/form-data`
* `html`

##### Content-type pré-associé (JSON)

* `application/json`
* `text/json`
* `json`

##### Content-type pré-associé (XML)

* `application/xml`
* `text/xml`
* `xml`

## Formulaires HTML

Le traitement des formulaires HTML utilise par défaut le paquet [github.com/monoculum/formam](https://github.com/monoculum/formam) pour associer le formulaire à une `struct`.

Prenez la struct `User` et le formulaire HTML ci-dessous.

```go
type User struct {
  Name     string `form:"name"`
  Email    string
  Password string `form:"-"`
}
```

```html
&lt;form&gt;
  &lt;input type="text" value="ringo" name="name"/&gt;
  &lt;input type="text" value="ringo@beatles.com" name="Email"/&gt;
  &lt;input type="text" value="starr" name="Password"/&gt;
  &lt;input type="submit"/&gt;
&lt;/form&gt;
```

Dans une action, on peut associer ce formulaire HTML avec la struct `User` de la manière suivante :

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

L'association par défaut entre le formulaire et la struct se base sur le nom de l'attribut `name` du champ du formulaire. Si l'on prend le champ `Email` du formulaire, il sera associé au champ `Email` de la struct.

Si vous souhaitez ignorer des champs de la struct lors de l'association, vous pouvez le signaler en utilisant la valeur `-` (comme c'est le cas pour `Password` dans notre exemple).

Pour plus d'informations sur le tag `form`, consultez la documentation du paquet [github.com/monoculum/formam](https://github.com/monoculum/formam).

## JSON et XML


Le traitement des requêtes JSON et XML utilise par défaut les paquets `encoding/json` et `encoding/xml` de la bibliothèque standard de Go.

Prenons par exemple la struct `User` et l'objet JSON ci-dessous. (pour le XML, même fonctionnement, mais en remplaçant les tags `json` par des tags `xml`.)

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

Dans l'action, on peut associer cet objet JSON avec la struct `User` comme suit :

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

En utilisant les tags `json` ou `xml`, il est possible d'associer des champs de l'objet JSON à la struct `User`, et on peut ignorer certains champs en utilisant la valeur `-`.

Veuillez consulter la documentation de la bibliothèque standard de Go pour plus d'informations sur les tags `json` et `xml`.

## Customiser les associations 

Peut-être souhaitez-vous personnaliser le traitement de vos requêtes (au lieu d'utiliser ceux par défaut), ou bien ajouter un traitement pour un type non-supporté ? Voici comment faire, par exemple pour ajouter le support du type YAML :

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
