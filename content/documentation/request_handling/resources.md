---
name: Resources
seoDescription: "How to use Buffalo's resources?"
seoKeywords: ["buffalo", "go", "golang", "resources", "routing", "generator"]
---

# Resources

Often web applications need to build very similar ["CRUD"](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) end-points. To help reduce the amount of thought and complexity involved in this, Buffalo supports the concept of a "Resource".

The [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface allows Buffalo to map common routes and respond to common requests.

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

The [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface was made smaller in release `v0.14.1`. The `New` and `Edit` methods, which serve the HTML forms to edit the resource, are now optional.

Here's what the interface looked like before:

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

## Generating Resources

The `buffalo generate resource` command will generate the necessary models, migrations, Go code, and HTML to CRUD the resource.

When running the generator in an API application Buffalo will generate code to meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface.

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

When running the generator in a Web application Buffalo will generate code to the meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface, as well as the optional `New` and `Edit` methods.

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

## Example Resource Generation

In this example Buffalo will generate the code needed to CRUD a resource named `widget` (Go: `Widget`) that has the following attributes:

|                | Model Attribute | Go Type                                                                   | DB type                  | Form Type                |
|----------------|-----------------|---------------------------------------------------------------------------|--------------------------|--------------------------|
| `title`        | `Title`         | `string`                                                                  | `varchar`                | `text`                   |
| `description`  | `Description`   | [`nulls.String`](https://godoc.org/github.com/gobuffalo/pop/nulls#String) | `varchar (nullable)`     | `textarea`               |

```bash
$ buffalo generate resource widget title description:nulls.Text
```

<%= exampleDir("en/docs/resources/_example/standard") %>

## Destroying Resources

You can remove files generated by this generator by running:

```bash
$ buffalo destroy resource users
```

This command will ask you which files you want to remove, you can either answer each of the questions with y/n or you can pass the `-y` flag to the command like:

```bash
$ buffalo destroy resource users -y
```

Or in short form:

```bash
$ buffalo d r users -y
```


## Nesting resources

To simplify creating resource hierarchies, Buffalo supports nesting resources.

```go
type UsersResource struct {
  buffalo.Resource
}

type ImagesResource struct {
  buffalo.Resource
}

u := a.Resource("/users", UsersResource{})
u.Resource("/images", ImagesResource{})
```

This results in the following routes:

```bash
$ buffalo routes
METHOD | PATH                                    | ALIASES | NAME              | HANDLER
------ | ----                                    | ------- | ----              | -------
GET    | /                                       |         | rootPath          | github.com/gobuffalo/coke/actions.HomeHandler
GET    | /users                                  |         | usersPath         | github.com/gobuffalo/coke/actions.UsersResource.List
POST   | /users                                  |         | usersPath         | github.com/gobuffalo/coke/actions.UsersResource.Create
GET    | /users/new                              |         | newUsersPath      | github.com/gobuffalo/coke/actions.UsersResource.New
GET    | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Show
PUT    | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Update
DELETE | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Destroy
GET    | /users/{user_id}/edit                   |         | editUserPath      | github.com/gobuffalo/coke/actions.UsersResource.Edit
GET    | /users/{user_id}/images                 |         | userImagesPath    | github.com/gobuffalo/coke/actions.ImagesResource.List
POST   | /users/{user_id}/images                 |         | userImagesPath    | github.com/gobuffalo/coke/actions.ImagesResource.Create
GET    | /users/{user_id}/images/new             |         | newUserImagesPath | github.com/gobuffalo/coke/actions.ImagesResource.New
GET    | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Show
PUT    | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Update
DELETE | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Destroy
GET    | /users/{user_id}/images/{image_id}/edit |         | editUserImagePath | github.com/gobuffalo/coke/actions.ImagesResource.Edit
```

## buffalo.BaseResource

When a resource is generated it has [`buffalo.BaseResource`](https://godoc.org/github.com/gobuffalo/buffalo#BaseResource) embedded into it.

```go
type Widget struct {
  buffalo.BaseResource
}
```

The `buffalo.BaseResource` has basic implementations for all of the methods required by `buffalo.Resource`. These methods all `404`.

```go
// Edit default implementation. Returns a 404
func (v BaseResource) Edit(c Context) error {
  return c.Error(404, errors.New("resource not implemented"))
}
```

## Video Presentation

<%= vimeo("212302823") %>

## Related Content

* [Actions](/en/docs/actions) - Learn more about Buffalo actions.