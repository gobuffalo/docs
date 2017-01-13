# Error Handling

{{ partial "topics.html" }}

{{#panel title="Returning Errors From a Handler" name="returning-errors"}}

```go
func MyHandler(c buffalo.Context) error {
  // Return any old error, this will result in a 500 status code.
  return errors.New("boom!")
}
```

```go
func MyHandler(c buffalo.Context) error {
  // Use the Error function on the context.
  // This will result in a status code of 401.
  return c.Error(401, errors.New("Unauthorized!"))
}
```

{{/panel}}

{{#panel title="Default Error Handling (Development)" name="dev-error-handling"}}

In "development" mode (`GO_ENV=development`), Buffalo, will generate some helpful errors pages for you.

<table width="100%">
<tr>
  <td valign="top">
    <img src="/assets/images/404_example.png" width="100%">
  </td>
  <td>
    <img src="/assets/images/500_example.png" width="100%">
  </td>
</tr>
</table>

In "production" mode (`GO_ENV=production`), Buffalo, will not generate pages that have developer style information, instead the pages are simpler.
{{/panel}}

{{#panel title="Custom Error Handling"}}

While Buffalo, out of the box, will handle errors for you, it can be useful to handle errors in a custom way. To accomplish this, Buffalo, allows for the mapping of HTTP status codes to specific handlers, so the error can be dealt with in a custom fashion.

```go
app = buffalo.Automatic(buffalo.Options{
  Env: ENV,
})

app.ErrorHandlers[422] = func(status int, err error, c buffalo.Context) error {
  res := c.Response()
  res.WriteHeader(422)
  res.Write([]byte(fmt.Sprintf("Oops!! There was an error %s", err.Error())))
  return nil
}

app.GET("/oops", MyHandler)

func MyHandler(c buffalo.Context) error {
  return c.Error(422, errors.New("Oh no!"))
}
```

In the above example any error from you application that returns a status of `422` will be caught by the custom handler and will be dealt with accordingly.

{{/panel}}
