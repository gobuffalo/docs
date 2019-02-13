## Error Handling

Both `form` and `form_for` helpers have support for handling errors from the [`github.com/gobuffalo/validate`](https://github.com/gobuffalo/validate) package.

In an action simply set a value of type `*validate.Errors` on the context as `errors` and the form helpers will pick it up and add error messages to the appropriate form tags.

<%= codeTabs() { %>
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

```html
// templates/widgets/new.html
&lt;%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %&gt;
  &lt;%= f.InputTag("Name") %&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="&lt;%= widgetsPath() %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

```html
// OUTPUT
&lt;form action="/widgets" id="widget-form" method="POST"&gt;
  &lt;input name="authenticity_token" type="hidden" value="AI0pb5YFBw2xU/EfcS6FaEOwTLWaGv58Y+w0ArfJoknfqu7l/j6tRLWybbcm+YZqXbBmi7f80l3Sf0WfnR7COA=="&gt;
  &lt;div class="form-group has-error"&gt;
    &lt;label&gt;Widget&lt;/label&gt;
    &lt;input class=" form-control" id="widget-Widget" name="Widget" type="text" value=""&gt;
    &lt;span class="help-block"&gt;Widget can not be blank.&lt;/span&gt;
  &lt;/div&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;/form&gt;
```
<% } %>

