# Custom Helpers

<%= partial("docs/disclaimer.html") %>

No templating package would be complete without allowing for you to build your own, custom, helper functions.

<%= title("Registering Helpers") %>

Helper functions can be registered in two different places, depending on how they are to be used.

### Global Helpers

_Most_ helpers will be global helpers, meaning that they should be included in every template. The types of
helpers can can be set in `actions/render.go`:

<%= code("go", {file: "actions/render.go", "data-line": "4-8"}) { %>
func init() {
  r = render.New(render.Options{
    // ...
    Helpers: render.Helpers{
      "myHelper": func() string {
        return "hello"
      },
    },
    // ...
  })
}
<% } %>

### Per Request Helpers

Other helpers, that are specific to a certain request can be added to the `buffalo.Context` for that request.

<%= code("go", {file: "actions/home.go", "data-line": "2-4"}) { %>
func HomeHandler(c buffalo.Context) error {
  c.Set("myHelper", func() string {
    return "hello"
  })
  // ...
}
<% } %>

<%= title("Return Values") %>

Plush allows you to return any values you would like from a helper function. This guide will focus on helpers that
are designed to generate "output".

When returning multiple values from a function, the first value will be the one used for rendering in the template.
If the last return value is an `error`, Plush will handle that error.

---

#### `string`

Return just a `string`. The `string` will be HTML escaped, and deemed "not"-safe.

<%= code("go") { %>
func() string {
  return ""
}
<% } %>

#### `string, error`

Return a `string` and an error. The `string` will be HTML escaped, and deemed "not"-safe.

<%= code("go") { %>
func() (string, error) {
  return "", nil
}
<% } %>

---

#### `template.HTML`

[https://golang.org/pkg/html/template/#HTML](https://golang.org/pkg/html/template/#HTMLlate/#HTML)

Return a `template.HTML` string. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

<%= code("go") { %>
func() template.HTML {
  return template.HTML("")
}
<% } %>

#### `template.HTML, error`

Return a `template.HTML` string and an error. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

<%= code("go") { %>
func() (template.HTML, error) {
  return template.HTML(""), error
}
<% } %>

<%= title("Input Values") %>

Custom helper functions can take any type, and any number of arguments. There is an optional last argument, [`plush.HelperContext`](https://godoc.org/github.com/gobuffalo/plush#HelperContext), that can be received. It's quite useful, and I would recommend taking it, as it provides you access to things like the context of the call, the block associated with the helper, etc...

<%= title("Simple Helpers") %>

<%= code("go", {"data-line": "4-6"}) { %>
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "greet": func(name string) string {
      return fmt.Sprintf("Hi %s!", name)
    },
  },
  // ...
})
<% } %>

The `greet` function is now available to all templates that use that `render.Engine`.

<div class="code-tabs">
<%= code("go", {file: "actions/greet.go"}) { %>
func Greeter(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, r.String(`<h1>\<%= greet(name) %></h1>`))
}
<% } %>

<%= code("html", {file: "output"}) { %>
&lt;h1>Hi Mark!&lt;/h1>
<% } %>
</div>

<%= title("Block Helpers") %>

Like the `if` or `for` statements, block helpers take a "block" of text that can be evaluated and potentially rendered, manipulated, or whatever you would like. To write a block helper, you have to take the `plush.HelperContext` as the last argument to your helper function. This will give you access to the block associated with that call.

<div class="code-tabs">

<%= code("go", {file: "actions/render.go"}) { %>
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "upblock": upblock,
  },
  // ...
})
<% } %>

<%= code("go", {file: "helper"}) { %>
func upblock(help plush.HelperContext) (template.HTML, error) {
  s, err := help.Block()
  if err != nil {
    return "", err
  }
  return strings.ToUpper(s), nil
}
<% } %>

<%= code("go", {file: "actions/upper.go"}) { %>
func Upper(c buffalo.Context) error {
  return c.Render(200, r.HTML("up.html"))
}
<% } %>

<%= code("html", {file: "templates/up.html"}) { %>
\<%= upblock() { %>
  hello world
\<% } %>
<% } %>

<%= code("html", {file: "output"}) { %>
HELLO WORLD
<% } %>
</div>
