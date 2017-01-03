# Rendering

{{ partial "docs/disclaimer.html" }}

{{ partial "topics.html" }}

{{#panel title="Renderer Interface" name="interface"}}

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

{{/panel}}

{{#panel title="Creating a Render Engine"}}

A render engine is used to store information about configuration needed for rendering. Examples include: [helpers](/docs/helpers), [layouts](/docs/layouts), etc... Multiple engines can be initialized. For example one engine for the "main" site, and another for the "admin" portion.

By default an initial render engine is created for the application during application generation:

```go
var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "application.html",
		CacheTemplates: ENV == "production",
		FileResolverFunc: func() resolvers.FileResolver {
			return &resolvers.RiceBox{
				Box: rice.MustFindBox("../templates"),
			}
		},
	})
}
```

{{/panel}}
