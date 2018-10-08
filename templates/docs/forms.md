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
<%= partial("docs/forms/non_bootstrap.md") %>

<%= title("FAQs") %>

### How Do I Map a Form to a Model/Struct?

See the [Request Binding](/docs/bind) page for more information on request binding.

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
&lt;form enctype="multipart/form-data" method="POST"&gt;
&lt;/form&gt;
```

### Can I Just Use My Own Form (Without the Use of the Form Helper)?

Yes! You most definitely can create and use your own form! The forms provided from having Buffalo generate your resources are simply a placeholder to get you up and runnning quickly! It is important to note, however, that asking Buffalo to generate your resources, using the supplied generators, will also generate the resource's CRUD related routes.  This is important to note since the route associated with the UPDATE action makes use of the PUT method and is not a valid value for an HTML form method according to the [HTML Standard](https://www.w3.org/TR/html5/forms.html#association-of-controls-and-forms). That being said, you need to ensure that you structure your form (for editing a resource) to use the POST method to tunnel the HTTP method, while using a hidden input to indicate your intention to make use of the PUT method server side.  An example of this would look like the follow:

```html
&lt;form method="POST" ...&gt;
  &lt;input type="hidden" name="_method" value="PUT" /&gt;
...
```

#### How Do I Handle CSRF Tokens If I Use My Own Form?

If you do decide to use your own forms you are going to need a way to provide the form with the authenticity token.  To solve this problem you can create a helper inside `render.go`.  That code would look something like:

```go
"csrf": func() template.HTML {
	return template.HTML("&lt;input name=\"authenticity_token\" value=\"&lt;%= authenticity_token %&gt;\" type=\"hidden\"&gt;")
},
```

Now that you have defined a helper to use in your templates you can use your helper inside your form with `&lt;%= csrf() %&gt;`. So your custom form should end up looking like this:

```html
&lt;form method="POST" ...&gt;
  &lt;input type="hidden" name="_method" value="PUT" /&gt;
  &lt;%= csrf() %&gt;
...
```

