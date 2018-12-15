<%= title("Using Context Variables") %>

<%= sinceVersion("0.13.0-rc1") %>

To use context variables such as [RouteHelpers](/docs/routing#using-route-helpers-in-templates) or those set with
`c.Set(...)`, `mail.New` accepts a `buffalo.Context`.

```go
func SendMail(c buffalo.Context) error {
  m := mail.New(c)
  ...

  m.AddBody(r.HTML("mail.html"))
  return SMTP.Send(m)
}
```

```html
&lt;a href="\<%= awesomePath() %>">Click here&lt;/a>
```
