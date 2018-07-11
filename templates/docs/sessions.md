<% seoDescription("Sessions") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "session"]) %>

<%= h1("Sessions") %>
An HTTP session is a non-persistant data storage, which is destroyed on browser shutdown (in the default browser configuration). It can be used to store flash messages, or any temporary user-specific data. Use [cookies](/en/docs/cookies) instead if you need a more persistant client side storage.

The session is available directly from the `buffalo.Context` inside of a handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

<%= partial("docs/sessions/type.md") %>
<%= partial("docs/sessions/store.md") %>
<%= partial("docs/sessions/complex.md") %>
<%= partial("docs/sessions/save.md") %>
<%= partial("docs/sessions/null.md") %>

