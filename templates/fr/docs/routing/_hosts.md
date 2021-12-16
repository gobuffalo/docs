## Hôtes

<%= sinceVersion("0.18.2") %>

Buffalo prennent également en charge le regroupement des points de terminaison par hôte. L'hôte crée un nouveau groupe qui correspond au domaine transmis. Ceci est utile pour créer des groupes de points de terminaison pour différents domaines ou sous-domaines.

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

Les variables mappées aux paramètres sont également prises en charge:

```go
app.Host("{subdomain}.example.com")
app.Host("{subdomain:[a-z]+}.example.com")
```
