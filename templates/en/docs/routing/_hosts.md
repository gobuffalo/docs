## Hosts

<%= sinceVersion("0.18.2") %>

Buffalo apps also support grouping of end-points by host. Host creates a new group that matches the domain passed. This is useful for creating groups of end-points for different domains or subdomains.

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

Variables mapped to parameters are also supported:

```go
app.Host("{subdomain}.example.com")
app.Host("{subdomain:[a-z]+}.example.com")
```
