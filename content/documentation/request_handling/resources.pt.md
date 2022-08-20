---
name: Resources
seoDescription: "How to use Buffalo's resources?"
seoKeywords: ["buffalo", "go", "golang", "resources", "routing", "generator"]
weight: 3
aliases:
  - /pt/docs/resources
  - /pt/docs/resources
---

# Resources

Often web applications need to build very similar ["CRUD"](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) end-points. To help reduce the amount of thought and complexity involved in this, Buffalo supports the concept of a "Resource".

The [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface allows Buffalo to map common routes and respond to common requests.

{{< since "0.14.1" >}}

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
// action/users.go
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
```

Mapping the Resource in app.go:

```go
// actions/app.go
app.Resource("/users", UsersResource{})
```

The above code example would be the equivalent of the following:

```go
// actions/app.go
ur := UsersResource{}

app.GET("/users", ur.List)
app.GET("/users/{user_id}", ur.Show)
app.POST("/users", ur.Create)
app.PUT("/users/{user_id}", ur.Update)
app.DELETE("/users/{user_id}", ur.Destroy)
```

It will produce a routing table that looks similar to:

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /users/                 |         | usersPath            | coke/actions.UsersResource.List
POST   | http://127.0.0.1:3000 | /users/                 |         | usersPath            | coke/actions.UsersResource.Create
GET    | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Show
PUT    | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Update
DELETE | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Destroy
```

## Optional Resource Methods

{{< since "0.14.1" >}}

In `v0.14.1` the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) was made smaller with the following methods now being optional:

```go
New(Context) error
Edit(Context) error
```

When implemented the `New` and `Edit` methods will add the following to the routing table:

```bash
METHOD | HOST                   | PATH                   | ALIASES | NAME         | HANDLER
------ | ----                   | ----                   | ------- | ----         | -------
GET    | http://127.0.0.1:3000  | /users/new             |         | newUsersPath | coke/actions.UsersResource.New
GET    | http://127.0.0.1:3000  | /users/{user_id}/edit/ |         | editUserPath | coke/actions.UsersResource.Edit
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

{{< codetabs >}}
{{< tab "actions/app.go" >}}
```go
package actions

import (
	"net/http"

	"coke/locales"
	"coke/models"
	"coke/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
	T   *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		app.Resource("/widgets", WidgetsResource{})
		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
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
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

```
{{< /tab >}}
{{< tab "actions/widgets.go" >}}
```go
package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"coke/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Widget)
// DB Table: Plural (widgets)
// Resource: Plural (Widgets)
// Path: Plural (/widgets)
// View Template Folder: Plural (/templates/widgets/)

// WidgetsResource is the resource for the Widget model
type WidgetsResource struct {
	buffalo.Resource
}

// List gets all Widgets. This function is mapped to the path
// GET /widgets
func (v WidgetsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	widgets := &models.Widgets{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Widgets from the DB
	if err := q.All(widgets); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("widgets", widgets)
		return c.Render(http.StatusOK, r.HTML("widgets/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(widgets))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(widgets))
	}).Respond(c)
}

// Show gets the data for one Widget. This function is mapped to
// the path GET /widgets/{widget_id}
func (v WidgetsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("widget", widget)

		return c.Render(http.StatusOK, r.HTML("widgets/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(widget))
	}).Respond(c)
}

// New renders the form for creating a new Widget.
// This function is mapped to the path GET /widgets/new
func (v WidgetsResource) New(c buffalo.Context) error {
	c.Set("widget", &models.Widget{})

	return c.Render(http.StatusOK, r.HTML("widgets/new.plush.html"))
}

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func (v WidgetsResource) Create(c buffalo.Context) error {
	// Allocate an empty Widget
	widget := &models.Widget{}

	// Bind widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(widget)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("widget", widget)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("widgets/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "widget.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/widgets/%v", widget.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(widget))
	}).Respond(c)
}

// Edit renders a edit form for a Widget. This function is
// mapped to the path GET /widgets/{widget_id}/edit
func (v WidgetsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("widget", widget)
	return c.Render(http.StatusOK, r.HTML("widgets/edit.plush.html"))
}

// Update changes a Widget in the DB. This function is mapped to
// the path PUT /widgets/{widget_id}
func (v WidgetsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(widget)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("widget", widget)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("widgets/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "widget.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/widgets/%v", widget.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(widget))
	}).Respond(c)
}

// Destroy deletes a Widget from the DB. This function is mapped
// to the path DELETE /widgets/{widget_id}
func (v WidgetsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(widget); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "widget.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/widgets")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(widget))
	}).Respond(c)
}

```
{{< /tab >}}
{{< tab "actions/widgets_test.go" >}}
```go
package actions

func (as *ActionSuite) Test_WidgetsResource_List() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Show() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Create() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Update() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Destroy() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_New() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_WidgetsResource_Edit() {
	as.Fail("Not Implemented!")
}

```
{{< /tab >}}
{{< /codetabs >}}


{{< codetabs >}}
{{< tab "locales/widgets.en-us.yaml" >}}
```yaml
- id: "widget.created.success"
  translation: "Widget was successfully created."
- id: "widget.updated.success"
  translation: "Widget was successfully updated."
- id: "widget.destroyed.success"
  translation: "Widget was successfully destroyed."

```
{{< /tab >}}
{{< /codetabs >}}

{{< codetabs >}}
{{< tab "migrations/20181005153028_create_widgets.up.fizz" >}}
```erb
create_table("widgets") {
	t.Column("id", "uuid", {primary: true})
	t.Column("title", "string", {})
	t.Column("description", "text", {null: true})
	t.Timestamps()
}

```
{{< /tab >}}
{{< tab "migrations/20181005153028_create_widgets.down.fizz" >}}
```erb
drop_table("widgets")
```
{{< /tab >}}
{{< /codetabs >}}

{{< codetabs >}}
{{< tab "models/widget.go" >}}
```go
package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Widget is used by pop to map your widgets database table to your go code.
type Widget struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Description nulls.String `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (w Widget) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Widgets is not required by pop and may be deleted
type Widgets []Widget

// String is not required by pop and may be deleted
func (w Widgets) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Widget) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

```
{{< /tab >}}
{{< tab "models/widget_test.go" >}}
```go
package models

func (ms *ModelSuite) Test_Widget() {
	ms.Fail("This test needs to be implemented!")
}

```
{{< /tab >}}
{{< /codetabs >}}

{{< codetabs >}}
{{< tab "templates/widgets/_form.plush.html" >}}
```html
<%= f.InputTag("Title") %>
<%= f.TextAreaTag("Description", {rows: 10}) %>
<button class="btn btn-success" role="submit">Save</button>

```
{{< /tab >}}
{{< tab "templates/widgets/edit.plush.html" >}}
```html
<div class="py-4 mb-2">
  <h3 class="d-inline-block">Edit Widget</h3>
</div>

<%= formFor(widget, {action: widgetPath({ widget_id: widget.ID }), method: "PUT"}) { %>
  <%= partial("widgets/form.html") %>
  <%= linkTo(widgetPath({ widget_id: widget.ID }), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>

```
{{< /tab >}}
{{< tab "templates/widgets/index.plush.html" >}}
```html
<div class="py-4 mb-2">
  <h3 class="d-inline-block">Widgets</h3>
  <div class="float-end">
    <%= linkTo(newWidgetsPath(), {class: "btn btn-primary"}) { %>
      Create New Widget
    <% } %>
  </div>
</div>

<table class="table table-hover table-bordered">
  <thead class="thead-light">
    <th>Title</th>
    <th>&nbsp;</th>
  </thead>
  <tbody>
    <%= for (widget) in widgets { %>
      <tr>
        <td class="align-middle"><%= widget.Title %></td>
        <td>
          <div class="float-end">
            <%= linkTo(widgetPath({ widget_id: widget.ID }), {class: "btn btn-info", body: "View"}) %>
            <%= linkTo(editWidgetPath({ widget_id: widget.ID }), {class: "btn btn-warning", body: "Edit"}) %>
            <%= linkTo(widgetPath({ widget_id: widget.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
          </div>
        </td>
      </tr>
    <% } %>
  </tbody>
</table>

<div class="text-center">
  <%= paginator(pagination) %>
</div>

```
{{< /tab >}}
{{< tab "templates/widgets/new.plush.html" >}}
```html
<div class="py-4 mb-2">
  <h3 class="d-inline-block">New Widget</h3>
</div>

<%= formFor(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= partial("widgets/form.html") %>
  <%= linkTo(widgetsPath(), {class: "btn btn-warning", "data-confirm": "Are you sure?", body: "Cancel"}) %>
<% } %>

```
{{< /tab >}}
{{< tab "templates/widgets/show.plush.html" >}}
```html
<div class="py-4 mb-2">
  <h3 class="d-inline-block">Widget Details</h3>

  <div class="float-end">
    <%= linkTo(widgetsPath(), {class: "btn btn-info"}) { %>
      Back to all Widgets
    <% } %>
    <%= linkTo(editWidgetPath({ widget_id: widget.ID }), {class: "btn btn-warning", body: "Edit"}) %>
    <%= linkTo(widgetPath({ widget_id: widget.ID }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>
  </div>
</div>



<ul class="list-group mb-2 ">


  <li class="list-group-item pb-1">
    <label class="small d-block">Title</label>
    <p class="d-inline-block"><%= widget.Title %></p>
  </li>



  <li class="list-group-item pb-1">
    <label class="small d-block">Description</label>
    <p class="d-inline-block"><%= widget.Description %></p>
  </li>


</ul>

```
{{< /tab >}}
{{< /codetabs >}}

## Destroying Resources

You can remove files generated by this generator by running:

```bash
$ buffalo destroy resource users
```

This command will ask you which files you want to remove, you can either answer each of the questions with `y/n` or you can pass the `-y` flag to the command like:

```bash
$ buffalo destroy resource users -y
```

Or in short form:

```bash
$ buffalo d r users -y
```


## Nesting Resources

To simplify creating resource hierarchies, Buffalo supports nesting resources.

```go
type WidgetsResource struct {
	buffalo.Resource
}

type ImagesResource struct {
  buffalo.Resource
}

w := app.Resource("/widgets", WidgetsResource{})
w.Resource("/images", ImagesResource{})
```

This results in the following routes:

```bash
$ buffalo routes

METHOD | HOST                  | PATH                                         | ALIASES | NAME                | HANDLER
------ | ----                  | ----                                         | ------- | ----                | -------
GET    | http://127.0.0.1:3000 | /                                            |         | rootPath            | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /widgets/                                    |         | widgetsPath         | coke/actions.WidgetsResource.List
POST   | http://127.0.0.1:3000 | /widgets/                                    |         | widgetsPath         | coke/actions.WidgetsResource.Create
GET    | http://127.0.0.1:3000 | /widgets/new/                                |         | newWidgetsPath      | coke/actions.WidgetsResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/edit/                   |         | editWidgetPath      | coke/actions.WidgetsResource.Edit
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/                 |         | widgetImagesPath    | coke/actions.ImagesResource.List
POST   | http://127.0.0.1:3000 | /widgets/{widget_id}/images/                 |         | widgetImagesPath    | coke/actions.ImagesResource.Create
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/new/             |         | newWidgetImagesPath | coke/actions.ImagesResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/edit/ |         | editWidgetImagePath | coke/actions.ImagesResource.Edit
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
  return c.Error(http.StatusNotFound, errors.New("resource not implemented"))
}
```

## Video Presentation

{{< vimeo 212302823>}}

## Related Content

* [Actions](/documentation/request_handling/actions) - Learn more about Buffalo actions.

## Next Steps

* [Context](/documentation/request_handling/context/) - Learn more about Buffalo context.