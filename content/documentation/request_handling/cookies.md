---
name: Cookies
seoDescription: "Cookies"
seoKeywords: ["buffalo", "go", "golang", "http", "cookie"]
aliases:
  - /docs/cookies
  - /en/docs/cookies
---

# Cookies

An HTTP cookie is a small piece of data that a server sends to the user's web browser. The browser can store this data and send it back to the same server, even after the browser restart (unlike a [browser session](/documentation/request_handling/sessions)).

(HTTP) cookies are commonly used to save users state (like whether the user logged-in). See [https://golang.org/pkg/net/http/#Cookie](https://golang.org/pkg/net/http/#Cookie) for more information on cookies in Go.

## Setting a Cookie

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().Set("user_id", user.ID, 30 * 24 * time.Hour)
  // ...
}
```

## Setting a Cookie with Expiration

```go
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire in 1 year
  c.Cookies().SetWithExpirationTime("user_id", user.ID, exp)
  // ...
}
```

## Setting a Cookie with Path

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().SetWithPath("user_id", user.ID, "/user")
  // ...
}
```

## Advanced setting a Cookie way

```go
import "net/http"
```

```go
func MyHandler(c buffalo.Context) error {
  // ...
  ck := http.Cookie{
    Name:    "token",
    Value:   token,
    Path:    "/",
    Expires: time.Now().Add(30 * 24 * time.Hour), // expire in 1 month
  }

  http.SetCookie(c.Response(), &ck)
  // ...
}
```

See [Cookie struct](https://golang.org/src/net/http/cookie.go) for other parameters.

## Getting a Cookie

```go
func MyHandler(c buffalo.Context) error {
  value, err := c.Cookies().Get("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(value))
}
```

## Deleting a Cookie


```go
func MyHandler(c buffalo.Context) error {
  c.Cookies().Delete("user_id")
  // ...
}
```
