<%= title("Filtering Events") %>

In the [Emitting Events](#emitting-events) section the naming convention for <%= doclink("github.com/gobuffalo/events#Event.Kind") %> is described. By the checking the value of <%= doclink("github.com/gobuffalo/events#Event.Kind") %>.

<%= codeTabs() { %>
```go
// direct match
events.Listen(func(e events.Event) {
  if e.Kind != buffalo.EvtRouteStarted {
    // do nothing
    return
  }
  // do work on the route event
})
```

```go
// matching with a switch statement
events.Listen(func(e events.Event) {
  switch e.Kind {
  case buffalo.EvtAppStart, buffalo.EvtAppStop:
    // do work
  case "buffalo:dev:build:finished":
    // do work
  default:
    // do nothing
  }
})
```

```go
// matching error events
func init() {
  events.Listen(func(e events.Event) {
    if !e.IsError() {
      // do nothing
      return
    }
    // do something with e.Error
  })
}
```

```go
// matching on prefix
events.Listen(func(e events.Event) {
  if !strings.HasPrefix(e.Kind, "buffalo:") {
    // do nothing
    return
  }
  // do work only on events emitted by Buffalo
})
```
<% } %>
