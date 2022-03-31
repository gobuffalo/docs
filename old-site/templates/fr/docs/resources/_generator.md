## Générer des ressources

La commande `buffalo generate resource` génère les modèles nécessaires, les migrations, le code Go et les templates HTML pour gérer la ressource.

En utilisant le générateur avec une application de type API, Buffalo génère le code nécessaire pour implémenter l'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource).

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

Avec une application Web classique, le même générateur va générer le code nécessaire pour implémenter l'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource), mais aussi pour les méthodes facultatives `New` et `Edit`.

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  New(Context) error
  Create(Context) error
  Edit(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

<%= partial("fr/docs/resources/example.md") %>

