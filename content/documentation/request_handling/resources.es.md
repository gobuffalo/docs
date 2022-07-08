---
name: Recursos
seoDescription: "Cómo usar los 'Recursos' de Buffalo?"
seoKeywords: ["buffalo", "go", "golang", "resources", "routing", "generator"]
weight: 3
aliases:
  - /docs/resources
  - /es/docs/resources
---

# Recursos

A menudo las aplicaciones Web necesitan desarrollar end-points muy parecidos al modelo [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete). Para reducir la complejidad que esto añade, Buffalo soporta el concepto de "Recurso".

La interfaz [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) permite a Buffalo asignar rutas comunes y responder a solicitudes comunes.

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

La interfaz [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) se ha reducido en la versión `v0.14.1`. Los métodos `New` y `Edit`, que sirven a los formularios HTML para editar el recurso, son ahora opcionales.

Así es como se veía la interfaz antes:

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

Luego de implementar los métodos necesarios de la interfaz [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource), el recurso puede ser mapeado en la aplicación utilizando el método [`github.com/gobuffalo/buffalo#App.Resource`](https://godoc.org/github.com/gobuffalo/buffalo#App.Resource).

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

Mapeando el Recurso en app.go:

```go
// actions/app.go
app.Resource("/users", UsersResource{})
```

El código de ejemplo anterior sería el equivalente al siguente:

```go
// actions/app.go
ur := UsersResource{}

app.GET("/users", ur.List)
app.GET("/users/{user_id}", ur.Show)
app.POST("/users", ur.Create)
app.PUT("/users/{user_id}", ur.Update)
app.DELETE("/users/{user_id}", ur.Destroy)
```

Produciría una tabla de rutas que luce similar a la siguiente:

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

## Métodos de Recursos Opcionales

{{< since "0.14.1" >}}

En la version `v0.14.1` el [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) se redujo con los siguientes métodos que son ahora opcionales:

```go
New(Context) error
Edit(Context) error
```

Cuando se implementan los métodos `New` y `Edit` se añade lo siguiente a la tabla de enrutamiento:

```bash
METHOD | HOST                   | PATH                   | ALIASES | NAME         | HANDLER
------ | ----                   | ----                   | ------- | ----         | -------
GET    | http://127.0.0.1:3000  | /users/new             |         | newUsersPath | coke/actions.UsersResource.New
GET    | http://127.0.0.1:3000  | /users/{user_id}/edit/ |         | editUserPath | coke/actions.UsersResource.Edit
```

## Generando Recursos

El comando `buffalo generate resource` generará los modelos, las migraciones, el código Go y el HTML necesarios para el CRUD del recurso.

Cuando se ejecuta el generador en una aplicación API Buffalo generará código para cumplir con la interfaz [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource).

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

Cuando se ejecuta el generador en una aplicación web, Buffalo generará código para cumplir con la interfaz  [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource), así como los métodos opcionales New y Edit

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

## Ejemplo de Generación de Recursos
En este ejemplo Buffalo generará el código necesario para hacer un CRUD de un recurso llamado `widget` (Go: `Widget`) que tiene los siguientes atributos:

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

## Eliminando recursos

Puedes eliminar los archivos generados por este generador ejecutando:

```bash
$ buffalo destroy resource users
```

Este comando te preguntará qué archivos quiere eliminar, puede responder a cada una de las preguntas con `y/n` o puede pasar el flag `-y` al comando de la siguiente manera:

```bash
$ buffalo destroy resource users -y
```

O en forma abreviada:

```bash
$ buffalo d r users -y
```


## Anidación de Recursos

Para simplificar la creación de jerarquías de recursos, Buffalo admite la anidación de recursos.

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

Esto da como resultado las siguientes rutas:

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

Cuando se genera un recurso, éste tiene incorporado el [`buffalo.BaseResource`](https://godoc.org/github.com/gobuffalo/buffalo#BaseResource).

```go
type Widget struct {
  buffalo.BaseResource
}
```

El `buffalo.BaseResource` tiene implementaciones básicas para todos los métodos requeridos por `buffalo.Resource`. Todos estos métodos retornan un `404` por defecto.

```go
// Edit default implementation. Returns a 404
func (v BaseResource) Edit(c Context) error {
  return c.Error(http.StatusNotFound, errors.New("resource not implemented"))
}
```

## Video de presentación

{{< vimeo 212302823>}}

## Contenido relacionado

* [Acciones](/es/documentation/request_handling/actions) - Más información sobre las acciones de Buffalo.

## Siguientes pasos

* [Contexto](/es/documentation/request_handling/context/) - Más información sobre el contexto de Buffalo.