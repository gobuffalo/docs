# Helpers

<%= partial("docs/disclaimer.html") %>

<%= title("Builtin Helpers") %>

Listed below are a few of the helpers that ship with Plush. Please see the Plush [documentation](https://github.com/gobuffalo/plush) for more information on it's helpers.

* `json` - returns a JSON marshaled string of the value passed to it.
* `jsEscape` - safely escapes a string to be used in a JavaScript bit of code.
* `htmlEscape` - safely escapes a string to be used in an HTML bit of code.
* `upcase` - upper cases the entire string passed to it.
* `downcase` - lower cases the entire string passed to it.
* `markdown` - converts markdown to HTML.
* `len` - returns the length of a value

Plush also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

<%= title("Content Helpers") %>

Plush ships with two complimentary helpers that let you create dynamic HTML snippets and re-use them later in the template.

### The `contentFor` Helper

The `contentFor` helper takes a block of HTML and holds on to using the given name. Take the following example. We want to set up a group of buttons to be displayed at the top and bottom of a form. Using the `contentFor` helper we can store that HTML in the template under the name `buttons`.

<%= code("html") { %>
\\<% contentFor("buttons") { %>
  &lt;button>Save&lt;/button>&lt;button>Cancel&lt;/button>
\\<% } %>
<% } %>

### The `contentOf` Helper

To retrieve the content stored as `buttons`, we can use the `contentOf` helper:

<%= code("html") { %>
\\<%= contentOf("buttons") %>
<% } %>
