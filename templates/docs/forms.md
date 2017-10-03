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

<%= partial("docs/forms/form.md") %>
<%= partial("docs/forms/form_for.md") %>
<%= partial("docs/forms/select_tag.md") %>
<%= partial("docs/forms/checkbox.md") %>
<%= partial("docs/forms/errors.md") %>

<%= title("FAQs") %>

### Can I Change the Name of the `f` Variable in My Template?

By default the form value inside the block is given the name `f`, however this can be changed when creating the form and passing the `var` option.

```html
&lt;%= form({var: "xyz"}) { %&gt;
  &lt;%= xyz.InputTag({name: "Foo"}) %&gt;
&lt;% } %&gt;
```

### How Do I Create a Multipart Form?

```html
&lt;%= form({multipart: true}) { %&gt;
&lt;% } %&gt;
```

```html
&lt;form enctype="multipart/form-data" errors="&lt;no value&gt;" method="POST"&gt;
&lt;/form&gt;
```
