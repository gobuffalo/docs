<% seoDescription("Cookies") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "cookie"]) %>

<%= h1("Cookies") %>

An HTTP cookie is a small piece of data that a server sends to the user's web brower. The browser can store this data and send it back to the same server, even after the browser restart (unlike a [browser session](/en/docs/sessions)).

(HTTP) cookies are commonly used to save users state (like whether the user logged-in). See [https://golang.org/pkg/net/http/#Cookie](https://golang.org/pkg/net/http/#Cookie) for more information on cookies in Go.

<%= title("Setting a Cookie") %>

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().Set("user_id", user.ID, 30 * 24 * time.Hour)
  // ...
}
```

<%= title("Setting a Cookie with Expiration") %>

```go
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire in 1 year
  c.Cookies().SetWithExpirationTime("user_id", user.ID, exp)
  // ...
}
```

<%= title("Getting a Cookie") %>

```go
func MyHandler(c buffalo.Context) error {
  value, err := c.Cookies().Get("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(value))
}
```

<%= title("Deleting a Cookie") %>


```go
func MyHandler(c buffalo.Context) error {
  c.Cookies().Delete("user_id")
  // ...
}
```
