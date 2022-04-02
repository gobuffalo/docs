---
Name: "Rendering"
---

# Rendering


The [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/master/render) [[godoc]](https://godoc.org/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

## Render Auto

{{< since "0.11.0" >}}

In many cases, you'll have to provide the same contents in different formats: JSON, XML, HTML... Buffalo provides an easy way to do that using a single statement.

```go
func Beatles(c buffalo.Context) error {
  members := models.Members{}
  // ...
  return c.Render(200, r.Auto(c, members))
}
```

{{< vimeo 257736901>}}

## JSON and XML

When rendering JSON, or XML, using the [`r.JSON`](https://godoc.org/github.com/gobuffalo/buffalo/render#JSON) or [`r.XML`](https://godoc.org/github.com/gobuffalo/buffalo/render#XML), you pass the value that you would like to be marshaled and the appropriate marshaler will encode the value you passed and write it to the response with the correct content/type.

**NOTE**: If you already have a string that contains JSON or XML do **NOT** use these methods as they will attempt to marshal the string into JSON or XML causing strange responses.
What you could do instead is write a custom render function as explained in the [Custom Rendering](rendering#custom-rendering) section.
```go
func MyHandler(c buffalo.Context) error {
  return c.Render(200, r.JSON(User{}))
}
```

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(200, r.XML(User{}))
}
```


## Markdown

Files passed into the `render.HTML` or `render.Template` functions, that have an extension of `.md`, will be converted from Markdown (using GitHub flavored Markdown) to HTML before being run through the templating engine. This makes for incredibly easy templating for simpler pages.

```md
// beatles.md
# The Beatles

<%= for (name) in names { %>
* <%= name %>
<% } %>
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
<h1>The Beatles</h1>

<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
```


## JavaScript
{{< since "0.10.0" >}}

The [`render`](https://godoc.org/github.com/gobuffalo/buffalo/render) package has a new implementation of [`render.Renderer`](https://godoc.org/github.com/gobuffalo/buffalo/render#Renderer), [`render.JavaScript`](https://godoc.org/github.com/gobuffalo/buffalo/render#JavaScript).

This means inside of an action you can do the following:

```go
func HomeHandler(c buffalo.Context) error {
  return c.Render(200, r.JavaScript("index.js"))
}
```

The [`render.Options`](https://godoc.org/github.com/gobuffalo/buffalo/render#Options) type now has a new attribute, `JavaScriptLayout`. This new option is similar to the `HTMLLayout` option in that it will wrap `*.js` files inside of another `*.js`.

The new JavaScript renderer also has it’s own implementation of the `partial` function. This new implementation behaves almost the same as the original implementation, but is smart enough to know that if you are rendering an `*.html` file inside of a `*.js` file that it will need to be escaped properly, and so it does it for you.

```javascript
$("#new-goal-form").replaceWith("<%= partial("goals/new.html") %>");
```


## Automatic extensions

{{< since "0.10.2" >}}

You can use HTML, Javascript and Markdown renderers without specifying the file extension:

```go
// actions/beatles.go
func Beatles(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  // Render beatles.html
  return c.Render(200, r.HTML("beatles"))
}
```

This works with [partials](/en/docs/partials) too.


## Custom Rendering

The [`r.Func`](https://godoc.org/github.com/gobuffalo/buffalo/render#Func) method allows you to pass in a content type and a function to render your data to the provided `io.Writer`, which is commonly, the HTTP response, in particular, a [`*buffalo.Response`](https://godoc.org/github.com/gobuffalo/buffalo#Response).

```go
func MyHandler(c buffalo.Context) error {
  return c.Render(200, r.Func("application/csv", csvWriter))
}

func csvWriter(w io.Writer, d render.Data) error {
  cw := csv.NewWriter(w)
  if err := cw.Write([]string{"a", "b", "c"}); err != nil {
    return errors.WithStack(err)
  }
  cw.Flush()
  return nil
}
```

For smaller, or one off situations, using an anonymous function can be even easier.
In this example you can see how to use an anonymous function to render a string that already contains JSON.
```go
var myJSONString string
return c.Render(200, r.Func("application/json", func(w io.Writer, d render.Data) error {
  _, err := w.Write([]byte(myJSONString))
  return err
}))
```

## Renderer Interface

In order for a renderer to be able to be used with [`Context#Render`](/en/docs/context) it must implement the following interface:

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

The [https://github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/master/render) [[godoc]](https://godoc.org/github.com/gobuffalo/buffalo/render) package implements that interface, and has a collection of useful render types already defined. It is recommended that you use this package, but feel free and write your own renderers!


