<%= h1("Resources") %>

Often web applications need to build very similar "CRUD" end-points. To help reduce the amount of thought and complexity involved in this, Buffalo supports the concept of a "Resource".

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

```go
type UserResource struct{
  buffalo.Resource
}

a.Resource("/users", UserResource{})
```

The above code example would be the equivalent of the following:

```go
ur := UserResource{}
a.GET("/users", ur.List)
a.GET("/users/new", ur.New)
a.GET("/users/{user_id}", ur.Show)
a.GET("/users/{user_id}/edit", ur.Edit)
a.POST("/users", ur.Create)
a.PUT("/users/{user_id}", ur.Update)
a.DELETE("/users/{user_id}", ur.Destroy)
```
