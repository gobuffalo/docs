## Utiliser des ressources

Après avoir implémenté les méthodes de l'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource), la ressource peut être branchée à l'application en utilisant la méthode [`github.com/gobuffalo/buffalo#App.Resource`](https://godoc.org/github.com/gobuffalo/buffalo#App.Resource).

```go
type UsersResource struct{ }

func (u UsersResource) List(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Show(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Create(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Update(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Destroy(c buffalo.Context) error {
  // do work
}

a.Resource("/users", UsersResource{})
```

Le code ci-dessus aurait pour équivalent celui ci-dessous :

```go
ur := UsersResource{}
a.GET("/users", ur.List)
a.GET("/users/{user_id}", ur.Show)
a.POST("/users", ur.Create)
a.PUT("/users/{user_id}", ur.Update)
a.DELETE("/users/{user_id}", ur.Destroy)
```

La table de routage (`$ buffalo routes`) ressemble alors à ceci :

```bash
METHOD | PATH                   | NAME         | HANDLER
------ | ----                   | ----         | -------
GET    | /users/                | usersPath    | coke/actions.UsersResource.List
POST   | /users/                | usersPath    | coke/actions.UsersResource.Create
GET    | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Show
PUT    | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Update
DELETE | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Destroy
```
