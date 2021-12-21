## Hôtes

<%= sinceVersion("0.18.2") %>

Buffalo supporte également le regroupement d'APIs par domaine. La méthode `Host` crée un nouveau groupe pour le nom de domaine passé en paramètre, ce qui peut être une méthode pratique pour regrouper des APIs appartenant au même domaine (voire gérer plusieurs sous-domaines sur une même application).

```go
app := buffalo.New(buffalo.Options{
    Env:         envy.Get("GO_ENV", "development"),
    SessionName: "_coke_session",
})

subApp := app.Host("docs.domain.com")
subApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("docs.domain.com Homepage"))
})

domainApp := app.Host("example.com")
domainApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("example.com Homepage"))
})

app.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("Main App Homepage"))
})
```

Les variables associéees aux paramètres sont également prises en charge :

```go
app.Host("{subdomain}.example.com")
app.Host("{subdomain:[a-z]+}.example.com")
```
