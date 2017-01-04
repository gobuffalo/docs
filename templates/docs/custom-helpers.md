# Custom Helpers

{{ partial "docs/disclaimer.html" }}

No templating package would be complete without allowing for you to build your own, custom, helper functions.

{{ partial "topics.html" }}

{{#panel title="Return Values"}}

The first thing to understand about building custom helper functions is their are a few "valid" return values:

#### `string`

Return just a `string`. The `string` will be HTML escaped, and deemed "not"-safe.

```go
func() string {
  return ""
}
```

#### `string, error`

Return a `string` and an error. The `string` will be HTML escaped, and deemed "not"-safe.

```go
func() (string, error) {
  return "", nil
}
```

#### `template.HTML`

[https://golang.org/pkg/html/template/#HTML](https://golang.org/pkg/html/https://golang.org/pkg/html/template/#HTMLlate/#HTML)

Return a `template.HTML` string. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

```go
func() template.HTML {
  return template.HTML("")
}
```


#### `template.HTML, error`

Return a `template.HTML` string and an error. The `template.HTML` will **not** be HTML escaped, and will be deemed safe.

```go
func() ( template.HTML, error ) {
  return template.HTML(""), error
}
```

{{/panel}}

{{#panel title="Input Values"}}

Custom helper functions can take any type, and any number of arguments. There is an option last argument, [`velvet.HelperContext`](https://godoc.org/github.com/gobuffalo/velvet#HelperContext), that can be received. It's quite useful, and I would recommend taking it, as it provides you access to things like the context of the call, the block associated with the helper, etc...

{{/panel}}

{{#panel title="Simple Helpers"}}

```go
r := render.New(render.Options{
  // ...
  Helpers: map[string]interface{}{
    "greet": greet,
  },
  // ...
})

func greet(name string) string {
  return fmt.Sprintf("Hi %s!", name)
}
```

The `greet` function is now available to all templates that use that `render.Engine`.

```go
func Greeter(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, r.String(`<h1>\{{greet "mark"}}</h1>`))
}
// <h1>Hi Mark!</h1>
```

{{/panel}}

{{#panel title="Block Helpers"}}

Like the [`if`](/docs/templating#if) and [`each`](/docs/helpers#each-array) helpers, block helpers take a "block" of text that can be evaluated and potentially rendered, manipulated, or whatever you would like. To write a block helper, you have to take the `velvet.HelperContext` as the last argument to your helper function. This will give you access to the block associated with that call.

```go
r := render.New(render.Options{
  // ...
  Helpers: map[string]interface{}{
    "upblock": upblock,
  },
  // ...
})

func upblock(help velvet.HelperContext) (template.HTML, error) {
  s, err := help.Block()
  if err != nil {
    return "", err
  }
  return strings.ToUpper(s), nil
}

func Upper(c buffalo.Context) error {
  return c.Render(200, r.String(`\{{#upblock}}hi\{{/upblock}}`))
}
// HI
```

{{/panel}}
