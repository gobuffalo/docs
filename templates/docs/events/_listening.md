## Listening for Events

To start listening for events a <%= doclink("github.com/gobuffalo/events#Listener") %> must first be registered with the <%= doclink("github.com/gobuffalo/events") %> package.

```go
func init() {
  _, err := events.Listen(func(e events.Event) {
    // do work
  })
}
```

Once registered this new listener function will be sent all events emitted through the <%= doclink("github.com/gobuffalo/events") %> package.
