<% seoDescription("Sessions") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "session"]) %>

<%= h1("Sessions") %>

An HTTP session is a non-persistent data storage, which is destroyed on browser shutdown (in the default browser configuration). It can be used to store flash messages, or any temporary user-specific data. Use [cookies](/en/docs/cookies) instead if you need a more persistent client side storage.

The session is available directly from the `buffalo.Context` inside of a handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

<%= partial("en/docs/sessions/type.md") %>
<%= partial("en/docs/sessions/store.md") %>
<%= partial("en/docs/sessions/complex.md") %>
<%= partial("en/docs/sessions/save.md") %>
<%= partial("en/docs/sessions/null.md") %>

