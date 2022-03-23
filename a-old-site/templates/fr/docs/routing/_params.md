## Paramètres

Les paramètres d'URL et autres paramètres de requête sont disponibles depuis le [`buffalo.Context`](/fr/docs/context) qui est passé au `buffalo.Handler`.

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Si l'on prend l'exemple ci-dessus : en appelant la route `GET /users?name=ringo`, la réponse devrait être `200: ringo`.
