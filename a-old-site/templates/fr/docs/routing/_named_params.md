## Paramètres nommés

Puisque le routeur Buffalo n'est rien d'autre qu'une version « habillée » du routeur [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux), vous pouvez accéder à toutes les fonctionnalités du routeur gorilla/mux. Par exemple, vous pouvez utiliser des pseudo-expressions régulières dans vos routes pour définir des paramètres nommés. Ils seront alors disponibles via le [`buffalo.Context`](/fr/docs/context).

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Dans l'exemple ci-dessus, si l'on appelle la route `GET /users/ringo`, la réponse devrait être `200: ringo`.

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, r.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Vous devriez pouvoir créer des chemins similaires, comme `/users/new` et `/users/{name}` sans problème. Le routeur s'assurera de les mener au bon endroit.

### Expressions régulières

[github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) permet d'utiliser des expressions régulières dans les URLs. Vous pouvez donc pré-filtrer les requêtes :

```go
a.GET("/articles/{id:[0-9]+}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("id")))
})
```
