# Layouts

{{ partial "docs/disclaimer.html" }}

{{ partial "topics.html" }}

{{#panel title="Using a Standard Layout" name="standard"}}

It is quite common to want to use the same layout across most, if not, all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

```go
var r *render.Engine

func init() {
	r = render.New(render.Options{
		// ...
		HTMLLayout:     "application.html",
		// ...
	})
}
```

```html
<!-- templates/application.html -->
<html>
  <head>
    <title>My App</title>
  </head>
  <body>
    <div id="main">
      \{{ yield }}
    </div>
  </body>
</html>
```

```html
<!-- templates/hello.html -->
<h1>Hello!!</h1>
```

```go
func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html"))
}
```

```html
<!-- result -->
<html>
  <head>
    <title>My App</title>
  </head>
  <body>
    <div id="main">
      <h1>Hello!!</h1>
    </div>
  </body>
</html>
```

{{/panel}}

{{#panel title="Using a Custom Layout" name="custom"}}

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`.

```go
var r *render.Engine

func init() {
	r = render.New(render.Options{
		// ...
		HTMLLayout:     "application.html",
		// ...
	})
}
```

```html
<!-- templates/custom.html -->
<html>
  <head>
    <title>My Custom Layout</title>
  </head>
  <body>
    <div id="main">
      \{{ yield }}
    </div>
  </body>
</html>
```

```html
<!-- templates/hello.html -->
<h1>Hello!!</h1>
```

```go
func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html", "custom.html"))
}
```

```html
<!-- result -->
<html>
  <head>
    <title>My Custom Layout</title>
  </head>
  <body>
    <div id="main">
      <h1>Hello!!</h1>
    </div>
  </body>
</html>
```

{{/panel}}
