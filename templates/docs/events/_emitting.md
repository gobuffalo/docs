<%= title("Emitting Events") %>

When emitting events the `Kind` attribute should be a unique, but constant, string. It is this attribute that users will use to determine how to respond to events they receive.

It is recommended to namespace this attribute like such, with error events being suffixed with `:err`.

```plain
"&lt;package-name&gt;:&lt;additional-names&gt;:&lt;optional-error&gt;"
"myapp:foo:start"
"myapp:foo:stop"
"mypkg:workers:bar:err"
```

This naming pattern makes it easier for users to filter events to only those that they care about. See [Filtering Events](#filtering-events) for more details.

---

There are multiple ways to emit an <%= doclink("github.com/gobuffalo/events#Event") %> in your Go code. The <%= doclink("github.com/gobuffalo/events#EmitError") %> and <%= doclink("github.com/gobuffalo/events#EmitPayload") %> functions both accept a `payload interface{}` argument. It is recommended to use <%= doclink("github.com/gobuffalo/events#Payload") %> for payloads; any other type passed in will get converted into a <%= doclink("github.com/gobuffalo/events#Payload") %> with the argument set in the payload with the key, `data`.

* <%= doclink("github.com/gobuffalo/events#Emit") %>

```go
func MyHandler(c buffalo.Context) error {
  e := events.Event{
    Kind:    "coke:myhandler:hello",
    Message: "hi!",
    Payload: events.Payload{"context": c},
  }
  if err := events.Emit(e); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

* <%= doclink("github.com/gobuffalo/events#EmitError") %>

```go
func MyHandler(c buffalo.Context) error {
  if err := events.EmitError("coke:myhandler:hello:err", errors.New("boom"), c); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

* <%= doclink("github.com/gobuffalo/events#EmitPayload") %>

```go
func MyHandler(c buffalo.Context) error {
  p := events.Payload{
    "message": "hi!",
  }
  if err := events.EmitPayload("coke:myhandler:hello", p); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

