# Helpers

<%= partial("docs/disclaimer.html") %>

<%= title("Builtin Helpers") %>

Listed below are a few of the helpers that ship with Plush. Please see the Plush [documentation](https://github.com/gobuffalo/plush) for more information on its helpers.

* `json` - converts the interface to a JSON object
* `jsEscape` - escapes the interface to be JavaScript safe
* `htmlEscape` - escapes the interface to be HTML safe
* `upcase` - converts the string to upper case
* `downcase` - converts the string to lower case
* `contentFor` - stores a block of HTML to be used later
* `contentOf` - retrieves a block of HTML previously stored with `contentFor`
* `markdown` - converts the string from Markdown into HTML
* `len` - returns the length of the interface
* `debug` - returns the `%+v` of the interface wrapped in `%lt;pre%gt;` tags.
* `inspect` - returns the `%+v` of the interface
* `range` - interate between, and including two numbers
* `between` - iterate between, but not including, two numbers
* `until` - iterate until a number is reached
* `groupBy` - splits a slice or array into `n` groups
* `env` - returns the ENV variable for the specified key
* `truncate` - truncates a string to a specified length
* `raw` - converts a string to `template.HTML`
* `form` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version)
* `form_for` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version) to build a form for a model

Plush also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

<%= title("Content Helpers") %>

Plush ships with two complimentary helpers that let you create dynamic HTML snippets and re-use them later in the template.

### The `contentFor` and `contentOf` Helpers

The `contentFor` helper takes a block of HTML and holds on to it using the given name. This block can then be used elsewhere in a template file, even when the content defined in a `contentFor` block is in a yielded-to template and is expanded into a `contentOf` block in a `yield`-calling template. The default `templates/application.html` calls `yield` like this.

Take the following example. Suppose we have a `templates/application.html` that fully specifies everything in `<head>` and the outermost contents of `<body>`. This template yields to other subtemplates, like `templates/users/show.html`, to fill `<body>`. However, if we want to add or override something in the `<head>` from a subtemplate, we'll need to use `contentFor`. In this example, we'll add a way for subtemplates to add an extra chunk of CSS to the `<head>` of `application.html`:

```html
&lt;!DOCTYPE html>
&lt;html>
  &lt;head>
    &lt;meta charset="utf-8">
    &lt;title>My Site&lt;/title>
    \<%= stylesheetTag("application.css") %>
    \<%= contentOf("extraStyle") %>
  &lt;/head>
  &lt;body>
    &lt;div class="container">
      \<%= partial("flash.html") %>
      \<%= yield %>
    &lt;/div>
  &lt;/body>
&lt;/html>
```

As it turns out, our `users/index.html` template could use a little page-wide styling instead of adding a bunch of `style` attributes to different elements, so it defines a block of CSS that doesn't show up anywhere inside the template:

```html
&lt;div class="page-header">
  &lt;h1>Users&lt;/h1>
&lt;/div>
&lt;table class="table table-striped">
  &lt;thead>
    &lt;th>Username&lt;/th> &lt;th>Password&lt;/th> &lt;th>Email&lt;/th> &lt;th>Admin?&lt;/th> &lt;th>&nbsp;&lt;/th>
  &lt;/thead>
  &lt;tbody>
    \<%= for (user) in users { %>
      &lt;!-- … -->
    \<% } %>
  &lt;/tbody>
&lt;/table>

\<% contentFor("extraStyle") %>
&lt;style>
    .online {
        color: limegreen;
        background: black;
    }
    
    .offline {
        color: lightgray;
        background: darkgray;
    }
&lt;/style>
\<% } %>
```

The styling for the `online` and `offline` classes then appears at the end of `<head>` in `/users`. In other pages, nothing is added.

Of course, if you'd rather do extensive processing on what goes into a chunk that goes on a webpage, you may want to do your processing in Go code instead of in templates. In that case, call, say, `c.Set("moonPhase", mp)` where `c` is a `buffalo.Context` in a function in an action like in `actions/users.go`, and `mp` is some string or object. Then, in your templates, refer to `<%= moonPhase %>` to display your expertly-calculated phase of the moon.
