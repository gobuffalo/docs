<%= title("Checkbox Tags") %>

Tags provide a convenient way to build an HTML `&lt;input&gt;` element with `type="checkbox"`:

```erb
&lt;%= f.CheckboxTag("IsPublic") %&gt;
```

That produces:

```html
&lt;div class="form-group"&gt;
  &lt;label&gt;
    &lt;input class="" id="talk-IsPublic" name="IsPublic" type="checkbox" value="true" checked=""&gt;
    IsPublic
  &lt;/label&gt;
&lt;/div&gt;
```

You can easily change the label content with

```erb
&lt;%= f.CheckboxTag("IsPublic", {label: "Is the talk public?"}) %&gt;
```

That produces:

```html
&lt;div class="form-group"&gt;
  &lt;label&gt;
    &lt;input class="" id="post-IsPublic" name="IsPublic" type="checkbox" value="true" checked=""&gt;
     Is the Talk public?
  &lt;/label&gt;
&lt;/div&gt;
```

### Non-Checked Checkbox Values

By default when a checkbox is not "checked" no value will be sent to the server. Often, it is useful to send a value indicating a non-checked checkbox. This can be set by passing in a `unchecked` value.

```go
&lt;%= f.CheckboxTag("IsPublic", {unchecked: false}) %&gt;
```

```html
&lt;div class="form-group"&gt;
  &lt;label&gt;
    &lt;input id="widget-IsPublic" name="IsPublic" type="checkbox" value="true"&gt;
    &lt;input name="IsPublic" type="hidden" value="false"&gt; IsPublic
  &lt;/label&gt;
&lt;/div&gt;
```

When the form is submitted the `hidden` tag will be posted and the server will see the `false` value.
