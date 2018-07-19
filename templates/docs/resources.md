<% seoDescription("Resources") %>
<% seoKeywords(["buffalo", "go", "golang", "resources", "routing", "generator"]) %>

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

a.Resource("/users", &UserResource{&buffalo.BaseResource{}})
```

The above code example would be the equivalent of the following:

```go
ur := &UserResource{}
a.GET("/users", ur.List)
a.GET("/users/new", ur.New)
a.GET("/users/{user_id}", ur.Show)
a.GET("/users/{user_id}/edit", ur.Edit)
a.POST("/users", ur.Create)
a.PUT("/users/{user_id}", ur.Update)
a.DELETE("/users/{user_id}", ur.Destroy)
```

<%= title("Generator", {}) %>

### Video Presentation

<%= vimeo("212302823") %>

### Example

```bash
$ buffalo g resource --help

Generates a new actions/resource file
Usage:
  buffalo generate resource [name] [flags]

Aliases:
  resource, r
```

```bash
$ buffalo g resource users name email bio:nulls.Text

      create  actions/users.go
      create  actions/users_test.go
      create  locales/users.en-us.yaml
      create  templates/users/_form.html
      create  templates/users/edit.html
      create  templates/users/index.html
      create  templates/users/new.html
      create  templates/users/show.html
         run  buffalo db g model user name email bio:nulls.Text
<%= version %>

      create  models/user.go
      create  models/user_test.go
         run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go actions/users.go actions/users_test.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go models/user.go models/user_test.go
      create  migrations/20180719054002_create_users.up.fizz
      create  migrations/20180719054002_create_users.down.fizz
         run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go actions/users.go actions/users_test.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go models/user.go models/user_test.go
```

```go
// actions/app.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/markbates/coke/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})
		// Automatically redirect to SSL
		app.Use(forceSSL())

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		app.Resource("/users", UsersResource{})
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
```

```go
// actions/users.go
package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/markbates/coke/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (User)
// DB Table: Plural (users)
// Resource: Plural (Users)
// Path: Plural (/users)
// View Template Folder: Plural (/templates/users/)

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

// List gets all Users. This function is mapped to the path
// GET /users
func (v UsersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	users := &models.Users{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Users from the DB
	if err := q.All(users); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, users))
}

// Show gets the data for one User. This function is mapped to
// the path GET /users/{user_id}
func (v UsersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, user))
}

// New renders the form for creating a new User.
// This function is mapped to the path GET /users/new
func (v UsersResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.User{}))
}

// Create adds a User to the DB. This function is mapped to the
// path POST /users
func (v UsersResource) Create(c buffalo.Context) error {
	// Allocate an empty User
	user := &models.User{}

	// Bind user to the html form elements
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, user))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User was created successfully")

	// and redirect to the users index page
	return c.Render(201, r.Auto(c, user))
}

// Edit renders a edit form for a User. This function is
// mapped to the path GET /users/{user_id}/edit
func (v UsersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, user))
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (v UsersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind User to the html form elements
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, user))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User was updated successfully")

	// and redirect to the users index page
	return c.Render(200, r.Auto(c, user))
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "User was destroyed successfully")

	// Redirect to the users index page
	return c.Render(200, r.Auto(c, user))
}
```

```go
// models/users.go
package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type User struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      string       `json:"name" db:"name"`
	Email     string       `json:"email" db:"email"`
	Bio       nulls.String `json:"bio" db:"bio"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
```

```fizz
// migration
create_table("users") {
	t.Column("id", "uuid", {"primary": true})
	t.Column("name", "string", {})
	t.Column("email", "string", {})
	t.Column("bio", "text", {"null": true})
}
```

```html
// templates/users/index.html
&lt;div class="page-header"&gt;
  &lt;h1&gt;Users&lt;/h1&gt;
&lt;/div&gt;
&lt;ul class="list-unstyled list-inline"&gt;
  &lt;li&gt;&lt;a href="&lt;%= newUsersPath() %&gt;" class="btn btn-primary"&gt;Create New User&lt;/a&gt;&lt;/li&gt;
&lt;/ul&gt;

&lt;table class="table table-striped"&gt;
  &lt;thead&gt;
  &lt;th&gt;Name&lt;/th&gt;
    &lt;th&gt;Email&lt;/th&gt;
    &lt;th&gt;&nbsp;&lt;/th&gt;
  &lt;/thead&gt;
  &lt;tbody&gt;
    &lt;%= for (user) in users { %&gt;
      &lt;tr&gt;
      &lt;td&gt;&lt;%= user.Name %&gt;&lt;/td&gt;
        &lt;td&gt;&lt;%= user.Email %&gt;&lt;/td&gt;
        &lt;td&gt;
          &lt;div class="pull-right"&gt;
            &lt;a href="&lt;%= userPath({ user_id: user.ID }) %&gt;" class="btn btn-info"&gt;View&lt;/a&gt;
            &lt;a href="&lt;%= editUserPath({ user_id: user.ID }) %&gt;" class="btn btn-warning"&gt;Edit&lt;/a&gt;
            &lt;a href="&lt;%= userPath({ user_id: user.ID }) %&gt;" data-method="DELETE" data-confirm="Are you sure?" class="btn btn-danger"&gt;Destroy&lt;/a&gt;
          &lt;/div&gt;
        &lt;/td&gt;
      &lt;/tr&gt;
    &lt;% } %&gt;
  &lt;/tbody&gt;
&lt;/table&gt;

&lt;div class="text-center"&gt;
  &lt;%= paginator(pagination) %&gt;
&lt;/div&gt;
```

```html
// templates/users/show.html
&lt;div class="page-header"&gt;
  &lt;h1&gt;User#Show&lt;/h1&gt;
&lt;/div&gt;

&lt;ul class="list-unstyled list-inline"&gt;
  &lt;li class="list-inline-item"&gt;&lt;a href="&lt;%= usersPath() %&gt;" class="btn btn-info"&gt;Back to all Users&lt;/a&gt;&lt;/li&gt;
  &lt;li class="list-inline-item"&gt;&lt;a href="&lt;%= editUserPath({ user_id: user.ID })%&gt;" class="btn btn-warning"&gt;Edit&lt;/a&gt;&lt;/li&gt;
  &lt;li class="list-inline-item"&gt;&lt;a href="&lt;%= userPath({ user_id: user.ID })%&gt;" data-method="DELETE" data-confirm="Are you sure?" class="btn btn-danger"&gt;Destroy&lt;/a&gt;
&lt;/ul&gt;

&lt;p&gt;
  &lt;strong&gt;Name&lt;/strong&gt;: &lt;%= user.Name %&gt;
&lt;/p&gt;
&lt;p&gt;
  &lt;strong&gt;Email&lt;/strong&gt;: &lt;%= user.Email %&gt;
&lt;/p&gt;
&lt;p&gt;
  &lt;strong&gt;Bio&lt;/strong&gt;: &lt;%= user.Bio %&gt;
&lt;/p&gt;
```

```html
// templates/users/new.html
&lt;div class="page-header"&gt;
  &lt;h1&gt;New User&lt;/h1&gt;
&lt;/div&gt;

&lt;%= form_for(user, {action: usersPath(), method: "POST"}) { %&gt;
  &lt;%= partial("users/form.html") %&gt;
  &lt;a href="&lt;%= usersPath() %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

```html
// templates/users/_form.html
&lt;%= f.InputTag("Name") %&gt;
&lt;%= f.InputTag("Email") %&gt;
&lt;%= f.TextAreaTag("Bio", {rows: 10}) %&gt;
&lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
```

```html
// templates/users/edit.html
&lt;div class="page-header"&gt;
  &lt;h1&gt;Edit User&lt;/h1&gt;
&lt;/div&gt;

&lt;%= form_for(user, {action: userPath({ user_id: user.ID }), method: "PUT"}) { %&gt;
  &lt;%= partial("users/form.html") %&gt;
  &lt;a href="&lt;%= userPath({ user_id: user.ID }) %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

### Destroy Resources

You can remove files generated by this generator by running:

```bash
$ buffalo destroy resource users
```

This command will ask you which files you want to remove, you can either answer each of the questions with y/n or you can pass the -y flag to the command like:

```bash
$ buffalo destroy resource users -y
```

Or in short form:

```bash
$ buffalo d r users -y
```
