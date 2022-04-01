---
Name: "Custom Helpers"
---

# Custom Helpers

<p>
  <em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
  <em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

No templating package would be complete without allowing for you to build your own, custom, helper functions.

{{< vimeo 229572343>}}

## Registering Helpers

Helper functions can be registered in two different places, depending on how they are to be used.

### Global Helpers

_Most_ helpers will be global helpers, meaning that they should be included in every template. The types of
helpers can can be set in `actions/render.go`:

```go
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
```

### Per Request Helpers

Other helpers, that are specific to a certain request can be added to the `buffalo.Context` for that request.

```go
func HomeHandler(c buffalo.Context) error {
  // ...
  c.Set("myHelper", func() string {
    return "hello"
  })
  // ...
}
```

## Return Values

Plush allows you to return any values you would like from a helper function. This guide will focus on helpers that
are designed to generate "output".

When returning multiple values from a function, the first value will be the one used for rendering in the template.
If the last return value is an `error`, Plush will handle that error.

---

#### `string`

Return just a `string`. The `string` will be HTML escaped, and deemed "not"-safe.

```go
func() string {
  return ""
}
```

---

#### `template.HTML`

[https://golang.org/pkg/html/template/#HTML](https://golang.org/pkg/html/template/#HTMLlate/#HTML)

Return a `template.HTML` string. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

```go
func() template.HTML {
  return template.HTML("")
}
```

## Input Values

Custom helper functions can take any type, and any number of arguments. You can even use variadic functions. There is an optional last argument, [`plush.HelperContext`](https://godoc.org/github.com/gobuffalo/plush#HelperContext), that can be received. It's quite useful, and I would recommend taking it, as it provides you access to things like the context of the call, the block associated with the helper, etc...

## Simple Helpers

```go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "greet": func(name string) string {
      return fmt.Sprintf("Hi %s!", name)
    },
  },
  // ...
})
```

The `greet` function is now available to all templates that use that `render.Engine`.

```go
// actions/greet.go
func Greeter(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, r.String("&lt;h1>\<%= greet(name) %></h1>"))
}
```

```go
// output
&lt;h1>Hi Mark!&lt;/h1>
```

## Block Helpers

Like the `if` or `for` statements, block helpers take a "block" of text that can be evaluated and potentially rendered, manipulated, or whatever you would like. To write a block helper, you have to take the `plush.HelperContext` as the last argument to your helper function. This will give you access to the block associated with that call.

{{< codetabs >}}
{{< tab "actions/render.go" >}}
```go
// actions/render.go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "upblock": upblock,
  },
  // ...
})
```
{{< /tab >}}
{{< tab "helper" >}}
```go
// helper
func upblock(help plush.HelperContext) (template.HTML, error) {
  s, err := help.Block()
  if err != nil {
    return "", err
  }
  return strings.ToUpper(s), nil
}
```
{{< /tab >}}
{{< tab "actions/upper.go" >}}
```go
// actions/upper.go
func Upper(c buffalo.Context) error {
  return c.Render(200, r.HTML("up.html"))
}
```
{{< /tab >}}
{{< tab "templates/up.html" >}}
```html
// templates/up.html
<%= upblock() { %>
  hello world
<% } %>
```
{{< /tab >}}

{{< tab "Output" >}}
```html
// output
HELLO WORLD
```
{{< /tab >}}
{{< /codetabs >}}


## Getting Values From the Context

{{< codetabs >}}
{{< tab "actions/render.go" >}}
```go
// actions/render.go
r := render.New(render.Options{
  // ...
  Helpers: render.Helpers{
    "is_logged_in": isLoggedIn,
  },
  // ...
})
```
{{< /tab >}}

{{< tab "helper" >}}
```go
// helper
func isLoggedIn(help plush.HelperContext) bool {
  return help.Value("current_user") != nil
}
```
{{< /tab >}}
{{< tab "action" >}}
```go
// actions/users.go
func Show(c buffalo.Context) error {
  c.Set("current_user", models.User{Name: "Ringo"})
  return c.Render(200, r.HTML("users/show.html"))
}
```
{{< /tab >}}
{{< tab "template" >}}
```html
// templates/users/show.html
\<%= if (is_logged_in()) { %>
  Hello \<%= current_user.Name %>
\<% } %>
```
{{< /tab >}}
{{< tab "Output" >}}
```html
// output
Hello Ringo
```
{{< /tab >}}
{{< /codetabs >}}
