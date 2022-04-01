---
name: File Uploads
---

# File Uploads
{{< since "0.10.3" >}}

Buffalo allows for the easily handling of files uploaded from a form. Storing those files, such as to disk or S3, is up to you the end developer: Buffalo just gives you easy access to the file from the request.

## Configuring the Form

The `f.FileTag` form helper can be used to quickly add a file element to the form. When using this the `enctype` of the form is *automatically* switched to be `multipart/form-data`.

```erb
&lt;%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %&gt;
  &lt;%= f.InputTag("Name") %&gt;
  &lt;%= f.FileTag("MyFile") %&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="&lt;%= widgetsPath() %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

## Accessing a Form File

In the [`buffalo.Context`](https://godoc.org/github.com/gobuffalo/buffalo#Context) the `c.File` takes a string, the name of the form file parameter and will return a [`binding.File`](https://godoc.org/github.com/gobuffalo/buffalo/binding#File) that can be used to easily retrieve a file from the from.

```go
func SomeHandler(c buffalo.Context) error {
  // ...
  f, err := c.File("someFile")
  if err != nil {
    return errors.WithStack(err)
  }
  // ...
}
```

## Binding to a Struct

The [`c.Bind`](https://godoc.org/github.com/gobuffalo/buffalo#Context) allows form elements to be bound to a struct, but it can also attach uploaded files to the struct. To do this, the type of the struct attribute **must** be a `binding.File` type.

In the example below you can see a model, which is configured to have a `MyFile` attribute that is of type `binding.File`. There is an `AfterCreate` callback on this example model that saves the file to disk after the model has been successfully saved to the database.

```go
// models/widget.go
type Widget struct {
  ID        uuid.UUID    `json:"id" db:"id"`
  CreatedAt time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
  Name      string       `json:"name" db:"name"`
  MyFile    binding.File `db:"-" form:"someFile"`
}

func (w *Widget) AfterCreate(tx *pop.Connection) error {
  if !w.MyFile.Valid() {
    return nil
  }
  dir := filepath.Join(".", "uploads")
  if err := os.MkdirAll(dir, 0755); err != nil {
    return errors.WithStack(err)
  }
  f, err := os.Create(filepath.Join(dir, w.MyFile.Filename))
  if err != nil {
    return errors.WithStack(err)
  }
  defer f.Close()
  _, err = io.Copy(f, w.MyFile)
  return err
}
```

**Note:** The `MyFile` attribute is not being saved to the database because of the `db:"-"` struct tag.


## Testing File Uploads

The HTTP testing library, [`github.com/gobuffalo/httptest`](https://github.com/gobuffalo/httptest) (which is included in the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package that Buffalo uses for testing) has been updated to include two new functions: [`MultiPartPost`](https://godoc.org/github.com/gobuffalo/httptest#Request.MultiPartPost) and [`MultiPartPut`](https://godoc.org/github.com/gobuffalo/httptest#Request.MultiPartPut).

These methods work just like the `Post` and `Put` methods, but instead they submit a multipart form, and can accept files for upload.

Like `Post` and `Put`, `MultiPartPost` and `MultiPartPut`, take a struct, or map, as the first argument: this is the equivalent of the HTML form you would post. The methods take a variadic second argument, [`httptest.File`](https://godoc.org/github.com/gobuffalo/httptest#File).

A `httptest.File` requires the name of the form parameter, `ParamName`; the name of the file, `FileName`; and an `io.Reader`, presumably the file you want to upload.


{{< codetabs >}}
{{< tab "actions/widgets_test.go" >}}
```go
// actions/widgets_test.go

func (as *ActionSuite) Test_WidgetsResource_Create() {
  // clear out the uploads directory
  os.RemoveAll("./uploads")

  // setup a new Widget
  w := &models.Widget{Name: "Foo"}

  // find the file we want to upload
  r, err := os.Open("./logo.svg")
  as.NoError(err)
  // setup a new httptest.File to hold the file information
  f := httptest.File{
    // ParamName is the name of the form parameter
    ParamName: "someFile",
    // FileName is the name of the file being uploaded
    FileName: r.Name(),
    // Reader is the file that is to be uploaded, any io.Reader works
    Reader: r,
  }

  // Post the Widget and the File(s) to /widgets
  res, err := as.HTML("/widgets").MultiPartPost(w, f)
  as.NoError(err)
  as.Equal(302, res.Code)

  // assert the file exists on disk
  _, err = os.Stat("./uploads/logo.svg")
  as.NoError(err)

  // assert the Widget was saved to the DB correctly
  as.NoError(as.DB.First(w))
  as.Equal("Foo", w.Name)
  as.NotZero(w.ID)
}
```
{{< /tab >}}

{{< tab "actions/widgets.go" >}}
```go
// actions/widgets.go

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func (v WidgetsResource) Create(c buffalo.Context) error {
  // Allocate an empty Widget
  widget := &models.Widget{}

  // Bind widget to the html form elements
  if err := c.Bind(widget); err != nil {
    return errors.WithStack(err)
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.WithStack(errors.New("no transaction found"))
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(widget)
  if err != nil {
    return errors.WithStack(err)
  }

  if verrs.HasAny() {
    // Make widget available inside the html template
    c.Set("widget", widget)

    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.HTML("widgets/new.html"))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", "Widget was created successfully")

  // and redirect to the widgets index page
  return c.Redirect(302, "/widgets/%s", widget.ID)
}
```
{{< /tab >}}
{{< tab "models/widgets.go" >}}
```go
// models/widgets.go

package models

import (
  "encoding/json"
  "io"
  "os"
  "path/filepath"
  "time"

  "github.com/gobuffalo/buffalo/binding"
  "github.com/gobuffalo/pop"
  "github.com/markbates/validate"
  "github.com/markbates/validate/validators"
  "github.com/pkg/errors"
  "github.com/satori/go.uuid"
)

type Widget struct {
  ID        uuid.UUID    `json:"id" db:"id"`
  CreatedAt time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
  Name      string       `json:"name" db:"name"`
  MyFile    binding.File `db:"-" form:"someFile"`
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

func (w *Widget) AfterCreate(tx *pop.Connection) error {
  if !w.MyFile.Valid() {
    return nil
  }
  dir := filepath.Join(".", "uploads")
  if err := os.MkdirAll(dir, 0755); err != nil {
    return errors.WithStack(err)
  }
  f, err := os.Create(filepath.Join(dir, w.MyFile.Filename))
  if err != nil {
    return errors.WithStack(err)
  }
  defer f.Close()
  _, err = io.Copy(f, w.MyFile)
  return err
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Widget) Validate(tx *pop.Connection) (*validate.Errors, error) {
  return validate.Validate(
    &validators.StringIsPresent{Field: w.Name, Name: "Name"},
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
{{< /codetabs >}}