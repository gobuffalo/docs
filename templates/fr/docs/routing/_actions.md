## Utiliser les *helpers* de routes dans les actions

### Redirections avec un *helper* de route

Vous pouvez également utiliser les noms de routes pour vos redirections (à la place des chemins en dur).

```go
func MyHandler(c buffalo.Context) error {
  return c.Redirect(307, "widgetsPath()")
  // Ou avec des paramètres
  return c.Redirect(307, "widgetPath()", render.Data{"widget_id": "1"})
}
```

---

### Trouver/appeler un *helper* de route

<%= sinceVersion("0.13.0-beta.1") %>

La méthode [`buffalo.RouteList#Lookup`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList.Lookup) vous permet de chercher une route à partir de son nom, depuis l'application. Vous obtenez alors une struct `RouteInfo` pour la route demandée, et cette struct vous permet de générer le chemin pour cette route.

```go
func MyHandler(c buffalo.Context) error {
  ri, err := App().Routes().Lookup("widgetPath")
  if err != nil {
    return errors.WithStack(err)
  }
  h := ri.BuildPathHelper()
  u, err := h(render.Data{"widget_id": 1})
  if err != nil {
    return errors.WithStack(err)
  }
  return c.Redirect(307, string(u))
}
```

