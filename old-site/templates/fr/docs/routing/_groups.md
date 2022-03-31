## Groupes

Buffalo permet de grouper des routes ensemble. Cela permet de partager des fonctionnalités communes, telles que l'utilisation de [middlewares](/fr/docs/middleware). Un bon exemple serait une racine d'API.

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // répond à GET /api/v1/users
})
```

Par défaut, un groupe de routes hérite de tous les middlewares de son application parente.

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

Dans l'exemple ci-dessus, le groupe `/api/v1` utilisera les middlewares `SomeMiddleware` et `APIAuthorizer`. Consultez la page [Middleware](/fr/docs/middleware) pour plus d'informations sur l'utilisation des middlewares.
