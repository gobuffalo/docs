## Checkbox Tags

Tags provide a convenient way to build an HTML `<input>` element with `type="checkbox"`:

```erb
\<%= f.CheckboxTag("IsPublic") %>
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
\<%= f.CheckboxTag("IsPublic", {label: "Is the talk public?"}) %>
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

```go
\<%= f.CheckboxTag("IsPublic", {unchecked: false}) %>
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
