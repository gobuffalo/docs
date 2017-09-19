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
* `form` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version)
* `form_for` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version) to build a form for a model

Plush also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

<%= title("Content Helpers") %>

Plush ships with two complimentary helpers that let you create dynamic HTML snippets and re-use them later in the template.

### The `contentFor` Helper

The `contentFor` helper takes a block of HTML and holds on to using the given name. Take the following example. We want to set up a group of buttons to be displayed at the top and bottom of a form. Using the `contentFor` helper we can store that HTML in the template under the name `buttons`.

```html
\<% contentFor("buttons") { %>
  &lt;button>Save&lt;/button>&lt;button>Cancel&lt;/button>
\<% } %>
```

### The `contentOf` Helper

To retrieve the content stored as `buttons`, we can use the `contentOf` helper:

```html
\<%= contentOf("buttons") %>
```
