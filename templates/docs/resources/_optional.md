## Optional Resource Methods

<%= sinceVersion("0.14.1") %>

In `v0.14.1` the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) was made smaller with the following methods now being optional:

```go
New(Context) error
Edit(Context) error
```

When implemented the `New` and `Edit` methods will add the following to the routing table:

```bash
METHOD | PATH                   | ALIASES | NAME         | HANDLER
------ | ----                   | ------- | ----         | -------
GET    | /users/new/            |         | newUsersPath | coke/actions.UsersResource.New
GET    | /users/{user_id}/edit/ |         | editUserPath | coke/actions.UsersResource.Edit
```
