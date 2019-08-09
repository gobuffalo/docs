<% seoDescription("Local authentication") %>
<% seoKeywords(["buffalo", "go", "golang", "users", "password", "authentication"]) %>

<%= h1("Local Authentication") %>

In many use-cases, you'll need to implement user authentication in your apps.

Buffalo had a native support for Auth until version `v0.9.4`. Since then, it was moved into it's own plugin, [https://github.com/gobuffalo/buffalo-auth](https://github.com/gobuffalo/buffalo-auth).

## Installation

To install the `buffalo-auth` plugin, run the following command:

```bash
$ go get -u github.com/gobuffalo/buffalo-auth
$ buffalo plugins install github.com/gobuffalo/buffalo-auth
```

## Generator
```bash
$ buffalo g auth

create  models/user.go
create  models/user_test.go
    run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go models/user.go models/user_test.go
create  migrations/20180910062057_create_users.up.fizz
create  migrations/20180910062057_create_users.down.fizz
create  actions/auth.go
create  actions/auth_test.go
create  actions/users.go
create  actions/users_test.go
create  models/user_test.go
create  actions/home_test.go
create  templates/auth/new.html
create  templates/index.html
create  templates/users/new.html
```

## Example Usage

### Actions
<%= codeTabs() { %>
```go
// actions/app.go
package actions

import (
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/buffalo/middleware"
  "github.com/gobuffalo/buffalo/middleware/ssl"
  "github.com/gobuffalo/envy"
  "github.com/unrolled/secure"

  "coke/models"

  "github.com/gobuffalo/buffalo/middleware/csrf"
  "github.com/gobuffalo/buffalo/middleware/i18n"
  "github.com/gobuffalo/packr"
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

    app.Use(SetCurrentUser)
    app.Use(Authorize)
    app.GET("/users/new", UsersNew)
    app.POST("/users", UsersCreate)
    app.GET("/signin", AuthNew)
    app.POST("/signin", AuthCreate)
    app.DELETE("/signout", AuthDestroy)
    app.Middleware.Skip(Authorize, HomeHandler, UsersNew, UsersCreate, AuthNew, AuthCreate)
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
// actions/auth.go
package actions

import (
  "database/sql"
  "strings"

    "coke/models"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/gobuffalo/validate"
  "github.com/pkg/errors"
  "golang.org/x/crypto/bcrypt"
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
  c.Set("user", models.User{})
  return c.Render(200, r.HTML("auth/new.html"))
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
  u := &models.User{}
  if err := c.Bind(u); err != nil {
    return errors.WithStack(err)
  }

  tx := c.Value("tx").(*pop.Connection)

  // find a user with the email
  err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)

  // helper function to handle bad attempts
  bad := func() error {
    c.Set("user", u)
    verrs := validate.NewErrors()
    verrs.Add("email", "invalid email/password")
    c.Set("errors", verrs)
    return c.Render(422, r.HTML("auth/new.html"))
  }

  if err != nil {
    if errors.Cause(err) == sql.ErrNoRows {
      // couldn't find an user with the supplied email address.
      return bad()
    }
    return errors.WithStack(err)
  }

  // confirm that the given password matches the hashed password from the db
  err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
  if err != nil {
    return bad()
  }
  c.Session().Set("current_user_id", u.ID)
  c.Flash().Add("success", "Welcome Back to Buffalo!")

  return c.Redirect(302, "/")
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
  c.Session().Clear()
  c.Flash().Add("success", "You have been logged out!")
  return c.Redirect(302, "/")
}
```

```go
// actions/auth_test.go
package actions

import (
  "coke/models"
)

func (as *ActionSuite) Test_Auth_New() {
  res := as.HTML("/signin").Get()
  as.Equal(200, res.Code)
  as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }
  verrs, err := u.Create(as.DB)
  as.NoError(err)
  as.False(verrs.HasAny())

  res := as.HTML("/signin").Post(u)
  as.Equal(302, res.Code)
  as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
  u := &models.User{
    Email:    "mark@example.com",
    Password: "password",
  }
  res := as.HTML("/signin").Post(u)
  as.Equal(422, res.Code)
  as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }
  verrs, err := u.Create(as.DB)
  as.NoError(err)
  as.False(verrs.HasAny())

  u.Password = "bad"
  res := as.HTML("/signin").Post(u)
  as.Equal(422, res.Code)
  as.Contains(res.Body.String(), "invalid email/password")
}
```

```go
// actions/home.go
package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
  return c.Render(200, r.HTML("index.html"))
}
```

```go
// actions/home_test.go
package actions

import "coke/models"

func (as *ActionSuite) Test_HomeHandler() {
  res := as.HTML("/").Get()
  as.Equal(200, res.Code)
  as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }
  verrs, err := u.Create(as.DB)
  as.NoError(err)
  as.False(verrs.HasAny())
  as.Session.Set("current_user_id", u.ID)

  res := as.HTML("/").Get()
  as.Equal(200, res.Code)
  as.Contains(res.Body.String(), "Sign Out")

  as.Session.Clear()
  res = as.HTML("/").Get()
  as.Equal(200, res.Code)
  as.Contains(res.Body.String(), "Sign In")
}
```

```go
// actions/users.go
package actions

import (
  "coke/models"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/pkg/errors"
)

func UsersNew(c buffalo.Context) error {
  u := models.User{}
  c.Set("user", u)
  return c.Render(200, r.HTML("users/new.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
  u := &models.User{}
  if err := c.Bind(u); err != nil {
    return errors.WithStack(err)
  }

  tx := c.Value("tx").(*pop.Connection)
  verrs, err := u.Create(tx)
  if err != nil {
    return errors.WithStack(err)
  }

  if verrs.HasAny() {
    c.Set("user", u)
    c.Set("errors", verrs)
    return c.Render(200, r.HTML("users/new.html"))
  }

  c.Session().Set("current_user_id", u.ID)
  c.Flash().Add("success", "Welcome to Buffalo!")

  return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if uid := c.Session().Get("current_user_id"); uid != nil {
      u := &models.User{}
      tx := c.Value("tx").(*pop.Connection)
      err := tx.Find(u, uid)
      if err != nil {
        return errors.WithStack(err)
      }
      c.Set("current_user", u)
    }
    return next(c)
  }
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if uid := c.Session().Get("current_user_id"); uid == nil {
      c.Flash().Add("danger", "You must be authorized to see that page")
      return c.Redirect(302, "/")
    }
    return next(c)
  }
}
```

```go
// actions/users_test.go
package actions

import (
  "coke/models"
)

func (as *ActionSuite) Test_Users_New() {
  res := as.HTML("/users/new").Get()
  as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
  count, err := as.DB.Count("users")
  as.NoError(err)
  as.Equal(0, count)

  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }

  res := as.HTML("/users").Post(u)
  as.Equal(302, res.Code)

  count, err = as.DB.Count("users")
  as.NoError(err)
  as.Equal(1, count)
}
```
<% } %>

### Models
<%= codeTabs() { %>
```go
// models/user.go
package models

import (
  "encoding/json"
  "time"

  "github.com/gobuffalo/pop"
  "github.com/gobuffalo/uuid"
  "github.com/gobuffalo/validate"
  "github.com/gobuffalo/validate/validators"
"strings"
"github.com/pkg/errors"
"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID           uuid.UUID `json:"id" db:"id"`
  CreatedAt    time.Time `json:"created_at" db:"created_at"`
  UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
  Email        string    `json:"email" db:"email"`
  PasswordHash string    `json:"password_hash" db:"password_hash"`
Password string `json:"-" db:"-"`
PasswordConfirmation string `json:"-" db:"-"`
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
    var err error
    return validate.Validate(
        &validators.StringIsPresent{Field: u.Email, Name: "Email"},
        &validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
        // check to see if the email address is already taken:
        &validators.FuncValidator{
            Field:   u.Email,
            Name:    "Email",
            Message: "%s is already taken",
            Fn: func() bool {
                var b bool
                q := tx.Where("email = ?", u.Email)
                if u.ID != uuid.Nil {
                    q = q.Where("id != ?", u.ID)
                }
                b, err = q.Exists(u)
                if err != nil {
                    return false
                }
                return !b
            },
        },
    ), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
    var err error
    return validate.Validate(
        &validators.StringIsPresent{Field: u.Password, Name: "Password"},
        &validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
    ), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
  return validate.NewErrors(), nil
}


// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
    u.Email = strings.ToLower(strings.TrimSpace(u.Email))
    ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return validate.NewErrors(), errors.WithStack(err)
    }
    u.PasswordHash = string(ph)
    return tx.ValidateAndCreate(u)
}
```

```go
// models/user_test.go
package models_test

import (
  "coke/models"
)

func (ms *ModelSuite) Test_User_Create() {
  count, err := ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(0, count)

  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }
  ms.Zero(u.PasswordHash)

  verrs, err := u.Create(ms.DB)
  ms.NoError(err)
  ms.False(verrs.HasAny())
  ms.NotZero(u.PasswordHash)

  count, err = ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
  count, err := ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(0, count)

  u := &models.User{
    Password: "password",
  }
  ms.Zero(u.PasswordHash)

  verrs, err := u.Create(ms.DB)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
  count, err := ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(0, count)

  u := &models.User{
    Email:                "mark@example.com",
    Password:             "password",
    PasswordConfirmation: "password",
  }
  ms.Zero(u.PasswordHash)

  verrs, err := u.Create(ms.DB)
  ms.NoError(err)
  ms.False(verrs.HasAny())
  ms.NotZero(u.PasswordHash)

  count, err = ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(1, count)

  u = &models.User{
    Email:    "mark@example.com",
    Password: "password",
  }
  verrs, err = u.Create(ms.DB)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("users")
  ms.NoError(err)
  ms.Equal(1, count)
}
```
<% } %>

### Migrations
<%= codeTabs() { %>
```go
// migrations/20180910062057_create_users.down.fizz
drop_table("users")
```

```go
// migrations/20180910062057_create_users.up.fizz
create_table("users") {
  t.Column("id", "uuid", {"primary": true})
  t.Column("email", "string", {})
  t.Column("password_hash", "string", {})
}
```
<% } %>

### Templates

```html
// templates/auth/new.html
&lt;style&gt;
  .auth-wrapper{
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .auth-wrapper .sign-form{
    max-width: 350px;
    width: 100%;
    padding: 0 20px;
  }

  .auth-wrapper h1{margin-bottom: 20px;}
&lt;/style&gt;

&lt;div class="auth-wrapper"&gt;
  &lt;div class="sign-form"&gt;
    &lt;h1&gt;Sign In&lt;/h1&gt;

    &lt;%= form_for(user, {action: signinPath()}) { %&gt;
      &lt;%= f.InputTag("Email") %&gt;
      &lt;%= f.InputTag("Password", {type: "password"}) %&gt;
      &lt;button class="btn btn-success"&gt;Sign In!&lt;/button&gt;
    &lt;% } %&gt;
  &lt;/div&gt;
&lt;/div&gt;
```


```html
// templates/new/new.html
&lt;style&gt;
  .auth-wrapper{
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .auth-wrapper .sign-form{
    max-width: 350px;
    width: 100%;
    padding: 0 20px;
  }

  .auth-wrapper h1{margin-bottom: 20px;}
&lt;/style&gt;

&lt;div class="auth-wrapper"&gt;
  &lt;div class="sign-form"&gt;
    &lt;h1&gt;register&lt;/h1&gt;

    &lt;%= form_for(user, {action: userspath()}) { %&gt;
      &lt;%= f.inputtag("email") %&gt;
      &lt;%= f.inputtag("password", {type: "password"}) %&gt;
      &lt;%= f.inputtag("passwordconfirmation", {type: "password"}) %&gt;

      &lt;button class="btn btn-success"&gt;register!&lt;/button&gt;
    &lt;% } %&gt;
  &lt;/div&gt;
&lt;/div&gt;
```

```html
// templates/index.html
&lt;style&gt;
  .auth-center{
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
  }
  .sign-in-btn{
    margin-right: 10px;
  }
&lt;/style&gt;

&lt;div class="auth-center"&gt;
  &lt;%= if (current_user) { %&gt;
    &lt;h1&gt;&lt;%= current_user.email %&gt;&lt;/h1&gt;
    &lt;a href="/signout" data-method="delete"&gt;sign out&lt;/a&gt;
  &lt;% } else { %&gt;
    &lt;a href="/signin" class="btn btn-primary"&gt;sign in&lt;/a&gt;
    &lt;a href="/users/new" class="btn btn-success"&gt;register&lt;/a&gt;
  &lt;% } %&gt;
&lt;/div&gt;
```
