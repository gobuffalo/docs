# Cookies

Buffalo does not currently have any wrappers around working with cookies. See [https://golang.org/pkg/net/http/#Cookie](https://golang.org/pkg/net/http/#Cookie) for more information on cookies in Go.

<%= title("Setting a Cookie") %>

<%= code("go") { %>
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire in 1 year
  cookie := http.Cookie{Name: "user_id", Value: user.ID, Expires: exp}
  http.SetCookie(c.Response(), &cookie)
  // ...
}
<%  }%>

<%= title("Getting a Cookie") %>

<%= code("go") { %>
func MyHandler(c buffalo.Context) error {
  cookie, err := c.Request().Cookie("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(cookie.Value))
}
<%  }%>
