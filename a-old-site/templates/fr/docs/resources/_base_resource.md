## buffalo.BaseResource

Lorsque Buffalo génère une ressource, le type [`buffalo.BaseResource`](https://godoc.org/github.com/gobuffalo/buffalo#BaseResource) est embarqué par défaut.

```go
type Widget struct {
  buffalo.BaseResource
}
```

`buffalo.BaseResource` a une implémentation basique pour toutes les méthodes requises par l'interface `buffalo.Resource`. Toutes les implémentations par défaut retournent une 404.

```go
// Edit default implementation. Returns a 404
func (v BaseResource) Edit(c Context) error {
  return c.Error(404, errors.New("resource not implemented"))
}
```