Les applications Web exposent souvent des URLs très similaires de type [CRUD](https://fr.wikipedia.org/wiki/CRUD). Pour vous aider à réduire le temps à copier coller la même chose, Buffalo supporte le concept de « Ressource ».

L'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) permet à Buffalo de brancher un groupe de routes communes et de réduire la quantité de code au strict nécessaire.

<%= sinceVersion("0.14.1") %>

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

Depuis la version `v0.14.1`, l'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) a été réduite.

À titre de comparaison, voici l'ancienne version de cette interface :

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
