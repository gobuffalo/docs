## Stop Listening for Events

When registering a new <%= doclink("github.com/gobuffalo/events#Listener") %> a <%= doclink("github.com/gobuffalo/events#DeleteFn") %> is returned. This function should be held on to and used when you want to remove the added listener.

```go
deleteFn, err := events.Listen(func(e events.Event) {
  // do work
})
if err != nil {
  return err
}
defer deleteFn()
```
