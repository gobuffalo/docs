## Skipping Middleware

There are times when, in an application, you want to add middleware to the entire application, or a group, but not call that middleware on a few individual handlers. Buffalo allows you to create these sorts of mappings.

<%= codeTabs() { %>
```go
// actions/app.go
a := buffalo.New(buffalo.Options{})
a.Use(AuthorizeUser)

// skip the AuthorizeUser middleware for the NewUser and CreateUser handlers.
a.Middleware.Skip(AuthorizeUser, NewUser, CreateUser)

a.GET("/users/new", NewUser)
a.POST("/users", CreateUser)
a.GET("/users", ListUsers)
a.GET("/users/{id}", ShowUser)
```

```text
// OUTPUT
GET /users/new -> NewUser
POST /users -> CreateUser
GET /users -> AuthorizeUser -> ListUsers
GET /users/{id} -> AuthorizeUser -> ShowUser
```
<% } %>

---

<div class="alert alert-warning" role="alert">
<b>IMPORTANT:</b> The middleware function and the action functions you want to skip <b>MUST</b> be the same Go instance.

<%= codeTabs() { %>
```go
// EXAMPLE 1
m1 := MyMiddleware()
m2 := MyMiddleware()

app.Use(m1)

app.Middleware.Skip(m2, Foo, Bar) // WON'T WORK m2 != m1
app.Middleware.Skip(m1, Foo, Bar) // WORKS
```

```go
// EXAMPLE 2
app.Resource("/widgets", WidgetResource{})
app.Middleware.Skip(mw, WidgetResource{}.Show) // WON'T WORK

wr := WidgetResource{}
app.Resource("/widgets", wr)
app.Middleware.Skip(mw, wr.Show) // WORKS
```

<% } %>

</div>

See [https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip) for more details on the `Skip` function.
