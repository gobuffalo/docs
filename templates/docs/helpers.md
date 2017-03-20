# Helpers

<%= partial("docs/disclaimer.html") %>

<%= partial("topics.html") %>


<%= panel("Other Builtin Helpers", {name: "other"}) { %>

* `json` - returns a JSON marshaled string of the value passed to it.
* `js_escape` - safely escapes a string to be used in a JavaScript bit of code.
* `html_escape` - safely escapes a string to be used in an HTML bit of code.
* `upcase` - upper cases the entire string passed to it.
* `downcase` - lower cases the entire string passed to it.
* `markdown` - converts markdown to HTML.
* `eq` - works like the `if` helper, but compares to values
* `neq` - the opposite of `eq` (not equal)
* `len` - returns the length of a value

Velvet also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

<% } %>

