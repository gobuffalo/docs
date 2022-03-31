## Méthodes de ressources optionnelles

<%= sinceVersion("0.14.1") %>

Avec la version `v0.14.1`, l'interface [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) a été simplifiée. Les méthodes suivantes sont désormais facultatives :

```go
New(Context) error
Edit(Context) error
```

Si ces méthodes sont implémentées, elles apparaissent dans la table de routage sans configuration supplémentaire :

```bash
METHOD | PATH                   | ALIASES | NAME         | HANDLER
------ | ----                   | ------- | ----         | -------
GET    | /users/new/            |         | newUsersPath | coke/actions.UsersResource.New
GET    | /users/{user_id}/edit/ |         | editUserPath | coke/actions.UsersResource.Edit
```
