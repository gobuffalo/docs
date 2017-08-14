# Sessions

The session is available directly from the `buffalo.Context` inside of a handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

<%= title("The Session Type") %>

The `buffalo.Session` type has everything needed to work with a session during a request. Under the covers Buffalo uses the [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) package for managing the session.

```go
type Session
  // Clear a session of all values
  func (s *Session) Clear()
  // Delete a specific value from the session
  func (s *Session) Delete(name interface{})
  // Get a value from the session
  func (s *Session) Get(name interface{}) interface{}
  // GetOnce gets a value from the current session and then deletes it.
  func (s *Session) GetOnce(name interface{}) interface{}
  // Save a session
  func (s *Session) Save() error
  // Set a value on the session
  func (s *Session) Set(name, value interface{})
```

<%= title("Session Store") %>

By default Buffalo will setup a session store using [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

This can be changed when setting up a new Buffalo application.

```go
app = buffalo.Automatic(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

The ENV variable `SESSION_SECRET` should be set before running the application. If this is not set, you will see a warning in your logs that your session is not secured.

For more information on this see the docs for [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).

<%= title("Storing Complex Types") %>

It is generally considered **not** good practice to store complex types in a session. There are lots of reasons for this, but it is recommended to store the ID of a type, instead of the "whole" value.

Should you need to store a complex type, like a `struct` you will first need to register the type with the [`encoding/gob`](https://golang.org/pkg/encoding/gob/) package.

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

<%= title("Saving a Session") %>

Saving a session requires simply calling the `Session#Save` method. However, it is important to **always** check the error returned. Often when people run into problems saving a session, they are not checking the error. The error will tell you what the problem is. The most common problem is [storing complex types](#storing-complex-types).

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
  // do some work
  err := s.Save()
  if err != nil {
    return err
  }
  // ...
}
```

Alternatively you can use the `middleware.SessionSaver` middleware to automatically save sessions at the end of every request.

```go
app.Use(middleware.SessionSaver)
```
