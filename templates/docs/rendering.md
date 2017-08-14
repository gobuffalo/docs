# Rendering

<%= partial("docs/disclaimer.html") %>

<%= title("Renderer Interface", {name: "interface"}) %>

In order for a renderer to be able to be used with [`Context#Render`](/docs/context) it must implement the following interface:

```go
// Renderer interface that must be satisified to be used with
// buffalo.Context.Render
type Renderer interface {
  ContentType() string
  Render(io.Writer, Data) error
}

// Data type to be provided to the Render function on the
// Renderer interface.

type Data map[string]interface{}
```

Thankfully the [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/master/render) [[godoc]](https://godoc.org/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!

<%= title("Creating a Render Engine", {}) %>

A render engine is used to store information about configuration needed for rendering. Examples include: [helpers](/docs/helpers), [layouts](/docs/layouts), etc. Multiple engines can be initialized. For example one engine for the "main" site, and another for the "admin" portion.

By default an initial render engine is created for the application during application generation:

```go
var r *render.Engine

func init() {
  r = render.New(render.Options{
    HTMLLayout: "application.html",
    TemplatesBox: packr.NewBox("../templates"),
    Helpers: render.Helpers{},
  })

}
```

<%= title("Markdown", {}) %>

Files passed into the `render.HTML` or `render.Template` functions, that have an extension of `.md`, will be converted from Markdown (using GitHub flavored Markdown) to HTML before being run through the templating engine. This makes for incredibly easy templating for simpler pages.

```md
// beatles.md
# The Beatles

\<%= for (name) in names { %>
* \<%= name %>
\<% } %>
```

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(200, r.HTML("beatles.md"))
}
```

```html
// output
&lt;h1>The Beatles</h1>

&lt;ul>
  &lt;li>John</li>
  &lt;li>Paul</li>
  &lt;li>George</li>
  &lt;li>Ringo</li>
&lt;/ul>
```
