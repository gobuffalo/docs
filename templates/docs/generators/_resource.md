<%= title("Resources") %>

```text
$ buffalo g resource --help

Generates a new actions/resource file

Usage:
  buffalo generate resource [name] [flags]

Aliases:
  resource, r
```

<div class="code-tabs">
<%= code("text", {file: "command"}) { %>
$ buffalo g resource users name email bio:nulls.Text

--> actions/users.go
--> actions/users_test.go
--> buffalo db g model user name email bio:nulls.Text
--> models/user.go
--> models/user_test.go
--> goimports -w .
> migrations/20170324195738_create_users.up.fizz
> migrations/20170324195738_create_users.down.fizz
--> goimports -w .
<% } %>

<%= code("go", {file: "actions/app.go", "data-line": "33-35"}) { %>
package actions

import (
	"github.com/gobuffalo/buffalo"

	"github.com/gobuffalo/buffalo/middleware"
	"github.com/markbates/cure/models"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_cure_session",
		})

		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
		var userResource buffalo.Resource
		userResource = UserResource{&buffalo.BaseResource{}}
		app.Resource("/user", userResource)
	}

	return app
}
<% } %>

<%= code("go", {file: "actions/users.go"}) { %>
package actions

import "github.com/gobuffalo/buffalo"

type UsersResource struct {
	buffalo.Resource
}

// List default implementation.
func (v UsersResource) List(c buffalo.Context) error {
	return c.Render(200, r.String("Users#List"))
}

// Show default implementation.
func (v UsersResource) Show(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Show"))
}

// New default implementation.
func (v UsersResource) New(c buffalo.Context) error {
	return c.Render(200, r.String("Users#New"))
}

// Create default implementation.
func (v UsersResource) Create(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Create"))
}

// Edit default implementation.
func (v UsersResource) Edit(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Edit"))
}

// Update default implementation.
func (v UsersResource) Update(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Update"))
}

// Destroy default implementation.
func (v UsersResource) Destroy(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Destroy"))
}
<% } %>

<%= code("go", {file: "models/users.go"}) { %>
package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
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

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (u *User) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
<% } %>

<%= code("fizz", {file: "migration"}) { %>
create_table("users", func(t) {
	t.Column("id", "uuid", {"primary": true})
	t.Column("name", "string", {})
	t.Column("email", "string", {})
	t.Column("bio", "text", {"null": true})
})
<% } %>
</div>
