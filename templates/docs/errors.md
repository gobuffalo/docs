# Error Handling

<%= title("Returning Errors From a Handler", {name: "returning-errors"}) %>

<div class="code-snippet">
  <%= code("go", {"data-line": 3}) { %>
func MyHandler(c buffalo.Context) error {
  // Return any old error, this will result in a 500 status code.
  return errors.New("boom!")
}
<% } %>

<%= code("go", {"data-line": 4}) { %>
func MyHandler(c buffalo.Context) error {
  // Use the Error function on the context.
  // This will result in a status code of 401.
  return c.Error(401, errors.New("Unauthorized!"))
}<% } %>
</div>


<%= title("Default Error Handling (Development)", {name: "dev-error-handling"}) %>

In "development" mode (`GO_ENV=development`), Buffalo will generate some helpful errors pages for you.

<figure>
  <img src="/assets/images/404_example.png" title="screenshot">
  <figcaption>An example of a `404` error in development mode.</figcaption>
</figure>

---

<figure>
  <img src="/assets/images/500_example.png" title="screenshot">
  <figcaption>An example of a `500` error in development mode.</figcaption>
</figure>

In "production" mode (`GO_ENV=production`), Buffalo will not generate pages that have developer style information. Instead the pages are simpler.

<%= title("Custom Error Handling", {}) %>

While Buffalo will handle errors for you out of the box, it can be useful to handle errors in a custom way. To accomplish this, Buffalo allows for the mapping of HTTP status codes to specific handlers. This means the error can be dealt with in a custom fashion.

<div class="code-tabs">
<%= code("go") { %>
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
<% } %>

<%= code("output") { %>
  GET /oops -> [422] On no!
<% } %>
</div>

In the above example any error from your application that returns a status of `422` will be caught by the custom handler and will be dealt with accordingly.

