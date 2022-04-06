---
Name: "Forms"
weight: 9
aliases:
  - /docs/forms
  - /en/docs/forms
---

# Forms

Buffalo uses the [`github.com/gobuffalo/tags`](https://github.com/gobuffalo/tags) package to make form building easier.

Plush includes two helpers from this package that produce Bootstrap v3 style forms. These helpers are `form` and `form_for`.

Both types of form helpers have the following features in common:

* Automatically setting the CSRF authenticity token
* Support for all HTTP methods (PUT, POST, DELETE, etc...)
* [Error Handling](#error-handling)
* Multipart form support
* Customizable input types
* Pass through HTML tag attributes

## Basic Forms

The `form` helper can be used to generate HTML forms. Since this type of form isn't attached to any particular "model" all information must be passed as options to the form and it's methods.

{{< codetabs >}}
{{< tab "templates/talks/edit.html" >}}
```erb
// templates/talks/edit.html

<%= form({action: talkPath({id: 3}), method: "PUT"}) { %>
  <div class="row">
    <div class="col-md-12">
      <%= f.InputTag({name:"Title", value: talk.Title }) %>
    </div>

    <div class="col-md-6">
      <%= f.TextArea({value: talk.Abstract, hide_label: true }) %>
    </div>

    <div class="col-md-6">
      <%= f.SelectTag({name: "TalkFormatID", value: talk.TalkFormatID, options: talk_formats}) %>
      <%= f.SelectTag({name: "AudienceLevel", value: talk.AudienceLevel, options: audience_levels }) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea({name: "Description", value: talk.Description, rows: 10}) %>
    </div>
    <div class="col-md-12">
      <%= f.TextArea({notes:"Notes", value: talk.Notes, rows: 10 }) %>
    </div>

  </div>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```html
// OUTPUT
<form action="/talks/3" method="POST">
  <input name="authenticity_token" type="hidden" value="e0c536b7a1a7d752066727b771f1e5d02220ceff5143f6c77b">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <input class=" form-control" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class=" form-control">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
        <select class=" form-control" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <select class=" form-control" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" notes="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
{{< /tab >}}
{{<  /codetabs  >}}



## Model Forms

The `form_for` helper can be used to generate HTML forms for a specified model. This makes the code easier to write, and maintains a level of "consistency" across your application.

The `form_for` helper behaves in a similar matter to the `form` helper, with several key differences.

The first difference is that the `form_for` takes a "model" as a first argument. This "model" only needs to be a `struct` it does not have to be database backed.

The second difference is in the tag calls the models directly. These tags, such as `InputTag`, take the name of the attribute on the model you want to build a field for, then they take an optional set of options as the second argument.

{{< codetabs >}}
{{< tab "models/talk.go" >}}
```go
// models/talk.go
type Talk struct {
  ID            int          `json:"id" db:"id"`
  CreatedAt     time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
  UserID        int          `json:"user_id" db:"user_id"`
  Title         string       `json:"title" db:"title"`
  Description   nulls.String `json:"description" db:"description"`
  Notes         nulls.String `json:"notes" db:"notes"`
  ParentID      nulls.Int    `json:"parent_id" db:"parent_id"`
  Abstract      string       `json:"abstract" db:"abstract"`
  AudienceLevel string       `json:"audience_level" db:"audience_level"`
  IsPublic      nulls.Bool   `json:"is_public" db:"is_public"`
  TalkFormatID  int          `json:"talk_format_id" db:"talk_format_id"`
}
```

{{< /tab>}}
{{< tab "templates/talks/edit.html" >}}
```erb
<%= form_for( talk, {action: talkPath({id: 3}), method: "PUT"}) { %>
  <div class="row">
    <div class="col-md-12">
      <%= f.InputTag("Title") %>
    </div>
    <div class="col-md-6">
      <%= f.TextArea("Abstract", {hide_label: true}) %>
    </div>


    <div class="col-md-6">
      <%= f.SelectTag("TalkFormatID", {options: talk_formats}) %>
      <%= f.SelectTag("AudienceLevel", , {options: audience_levels}) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea("Description", {rows: 10}) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea("Notes", {rows: 10}) %>
    </div>
  </div>
<% } %>
```
{{< /tab>}}
{{< tab "OUTPUT">}}
```html
// OUTPUT
<form action="/talks/3" id="talk-form" method="POST">
  <input name="authenticity_token" type="hidden" value="cd998be98a99b452481c43fd3e4715e4e85333a45b982ac999">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <label>Title</label>
        <input class="form-control" id="talk-Title" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class="form-control" id="talk-Abstract" name="Abstract">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
      <label>TalkFormatID</label>
        <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <label>AudienceLevel</label>
        <select class=" form-control" id="talk-AudienceLevel" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Description</label>
        <textarea class=" form-control" id="talk-Description" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Notes</label>
        <textarea class=" form-control" id="talk-Notes" name="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## Select Tags

To build your `<select>` tags inside forms Tags provide 3 convenient ways to add your `<select>` options: `form.SelectOptions`, `map[string]interface{}` or `[]string`, all of them by passing an `options` field into the `form.SelectTag` options like:


```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats}) %>
```

or

```erb
<%= f.SelectTag("TalkFormatID", {options: ["one", "two"]}) %>
```

Which will use the same value for the `value` attribute and the body of the option, or:

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}}) %>
```

Which allows us to define the options map inside the view.

### Selectable Interface

Another alternative for the select options is to pass a list of structs that meet the `form.Selectable` interface.

Which consist of two functions:

```go
//Selectable allows any struct to become an option in the select tag.
type Selectable interface {
  SelectValue() interface{}
  SelectLabel() string
}
```

By implementing this interface tags will call `SelectValue` and `SelectLabel` to get the option Value and Label from implementer.

### Selected

Tags will add the `selected` attribute to the option that has the same value than the one it receives on the `value` option of the `form.SelectTag`, so you don't have to look for the option that has equal value than the selected one manually, p.e:

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}, value: 2}) %>
```

Produces:

```html
<div class="form-group">
  <label>TalkFormatID</label>
  <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
    <option value="1">one</option>
    <option value="2" selected>two</option>
  </select>
</div>
```

And similarly with the `form.SelectOptions` slice:

```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats, value: 2}) %>
```

## Checkbox Tags

Tags provide a convenient way to build an HTML `<input>` element with `type="checkbox"`:

```erb
<%= f.CheckboxTag("IsPublic") %>
```

That produces:

```html
<div class="form-group">
  <label>
    <input class="" id="talk-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
    IsPublic
  </label>
</div>
```

You can easily change the label content with

```erb
<%= f.CheckboxTag("IsPublic", {label: "Is the talk public?"}) %>
```

That produces:

```html
<div class="form-group">
  <label>
    <input class="" id="post-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
     Is the Talk public?
  </label>
</div>
```

### Non-Checked Checkbox Values

By default when a checkbox is not "checked" no value will be sent to the server. Often, it is useful to send a value indicating a non-checked checkbox. This can be set by passing in a `unchecked` value.

```erb
<%= f.CheckboxTag("IsPublic", {unchecked: false}) %>
```

```html
<div class="form-group">
  <label>
    <input id="widget-IsPublic" name="IsPublic" type="checkbox" value="true">
    <input name="IsPublic" type="hidden" value="false"> IsPublic
  </label>
</div>
```

When the form is submitted the `hidden` tag will be posted and the server will see the `false` value.

## Error Handling

Both `form` and `form_for` helpers have support for handling errors from the [`github.com/gobuffalo/validate`](https://github.com/gobuffalo/validate) package.

In an action simply set a value of type `*validate.Errors` on the context as `errors` and the form helpers will pick it up and add error messages to the appropriate form tags.

{{< codetabs >}}
{{< tab "actions/widgets.go" >}}
```go
// actions/widgets.go
func (v WidgetsResource) Create(c buffalo.Context) error {
  tx := c.Value("tx").(*pop.Connection)
  widget := &models.Widget{}
  if err := c.Bind(widget); err != nil {
    return err
  }
  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(widget)
  if err != nil {
    return errors.WithStack(err)
  }
  if verrs.HasAny() {
    c.Set("widget", widget)
    // Make the errors available inside the html template
    c.Set("errors", verrs)
    return c.Render(422, r.HTML("widgets/new.html"))
  }
  c.Flash().Add("success", "Widget was created successfully")
  return c.Redirect(302, "/widgets/%s", widget.ID)
}
```
{{< /tab >}}
{{< tab "templates/widgets/new.html" >}}
```erb
// templates/widgets/new.html
<%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= f.InputTag("Name") %>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="<%= widgetsPath() %>" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}

```html
// OUTPUT
<form action="/widgets" id="widget-form" method="POST">
  <input name="authenticity_token" type="hidden" value="AI0pb5YFBw2xU/EfcS6FaEOwTLWaGv58Y+w0ArfJoknfqu7l/j6tRLWybbcm+YZqXbBmi7f80l3Sf0WfnR7COA==">
  <div class="form-group has-error">
    <label>Widget</label>
    <input class=" form-control" id="widget-Widget" name="Widget" type="text" value="">
    <span class="help-block">Widget can not be blank.</span>
  </div>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## Using Non-Bootstrap Form Helpers

The default form helpers, `form` and `form_for`, generate forms that are compatible with Bootstrap 3. If this is not for you, you can easily use the non-Bootstrap versions of these helpers.

<em><small>Requires Plush version `v3.6.8` or greater</small></em>

{{< codetabs >}}
{{< tab "actions/render.go" >}}
```go
// actions/render.go
func init() {
  r = render.New(render.Options{
    // ...
    // Add template helpers here:
    Helpers: render.Helpers{
      "form":     plush.FormHelper,
      "form_for": plush.FormForHelper,
    },
    // ...
  })
}
```
{{< /tab >}}
{{< tab "templates/widgets/new.html" >}}
```erb
// templates/widgets/new.html
<%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= f.InputTag("Name") %>
  <%= f.InputTag("Body") %>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="<%= widgetsPath() %>" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```html
// OUTPUT
<form action="/widgets" id="widget-form" method="POST">
  <input name="authenticity_token" type="hidden" value="jN3nYOhCTqxZvmYnO9v1maso2VMs8fslj3rmKg1TS281W6JKpMd6Uezqp1dd3VBu2su41nKRBkd5AWDyCM4BzQ==">
  <input id="widget-Name" name="Name" type="text" value="">
  <input id="widget-Body" name="Body" type="text" value="">
  <button class="btn btn-success" role="submit">Save</button>
  <a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## FAQs

### How Do I Map a Form to a Model/Struct?

See the [Request Binding](/documentation/request_handling/bind) page for more information on request binding.

### Can I Change the Name of the `f` Variable in My Template?

By default the form value inside the block is given the name `f`, however this can be changed when creating the form and passing the `var` option.

```erb
<%= form({var: "xyz"}) { %>
  <%= xyz.InputTag({name: "Foo"}) %>
<% } %>
```

### How Do I Create a Multipart Form?

```erb
<%= form({multipart: true}) { %>
<% } %>
```

```html
<form enctype="multipart/form-data" method="POST">
</form>
```

### Can I Just Use My Own Form (Without the Use of the Form Helper)?

Yes! You most definitely can create and use your own form! The forms provided from having Buffalo generate your resources are simply a placeholder to get you up and running quickly! It is important to note, however, that asking Buffalo to generate your resources, using the supplied generators, will also generate the resource's CRUD related routes.  This is important to note since the route associated with the UPDATE action makes use of the PUT method and is not a valid value for an HTML form method according to the [HTML Standard](https://www.w3.org/TR/html5/forms.html#association-of-controls-and-forms). That being said, you need to ensure that you structure your form (for editing a resource) to use the POST method to tunnel the HTTP method, while using a hidden input to indicate your intention to make use of the PUT method server side.  An example of this would look like the follow:

```html
<form method="POST" ...>
  <input type="hidden" name="_method" value="PUT" />
...
```

#### How Do I Handle CSRF Tokens If I Use My Own Form?

If you do decide to use your own forms you are going to need a way to provide the form with the authenticity token. There are two ways to solve this issue.

The first way is to use the `authenticity_token` directly in form, since it is already in the context.

```html
<form method="POST" ...>
  <input name="authenticity_token" type="hidden" value="<%= authenticity_token %>">
</form>
```

Another way is to write a helper to generate that line of code for you.

```go
"csrf": func(ctx plush.HelperContext) (template.HTML, error) {
  tok, ok := ctx.Value("authenticity_token").(string)
  if !ok {
    return "", fmt.Errorf("expected CSRF token got %T", ctx.Value("authenticity_token"))
  }
  t := tags.New("input", tags.Options{
    "value": tok,
    "type":  "hidden",
    "name":  "authenticity_token",
  })
  return t.HTML(), nil
},
```

Now that you have defined a helper to use in your templates you can use your helper inside your form with `<%= csrf() %>`. So your custom form should end up looking like this:

```erb
<form method="POST" ...>
  <%= csrf() %>
</form>
```

