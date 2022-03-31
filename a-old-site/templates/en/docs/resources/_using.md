## Using Resources

After implementing the necessary methods on the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface, the resource can then be mapped to the application using the [`github.com/gobuffalo/buffalo#App.Resource`](https://godoc.org/github.com/gobuffalo/buffalo#App.Resource) method.

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

The above code example would be the equivalent of the following:

```go
ur := UsersResource{}
a.GET("/users", ur.List)
a.GET("/users/{user_id}", ur.Show)
a.POST("/users", ur.Create)
a.PUT("/users/{user_id}", ur.Update)
a.DELETE("/users/{user_id}", ur.Destroy)
```

It will produce a routing table (`$ buffalo routes`) that looks similar to:

```bash
METHOD | PATH                   | NAME         | HANDLER
------ | ----                   | ----         | -------
GET    | /users/                | usersPath    | coke/actions.UsersResource.List
POST   | /users/                | usersPath    | coke/actions.UsersResource.Create
GET    | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Show
PUT    | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Update
DELETE | /users/{user_id}/      | userPath     | coke/actions.UsersResource.Destroy
```
