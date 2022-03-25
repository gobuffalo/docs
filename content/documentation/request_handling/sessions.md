---
Name: Sessions
seoDescription: "Sessions"
seoKeywords: ["buffalo", "go", "golang", "http", "session"]
---

# Sessions

An HTTP session is a non-persistent data storage, which is destroyed on browser shutdown (in the default browser configuration). It can be used to store flash messages, or any temporary user-specific data. Use [cookies](/en/docs/cookies) instead if you need a more persistent client side storage.

The session is available directly from the `buffalo.Context` inside of a handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

## The Session Type

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



## Session Store

By default Buffalo will setup a session store using [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

This can be changed when setting up a new Buffalo application using the `SessionStore` option:

```go
app = buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

The ENV variable `SESSION_SECRET` should be set before running the application. If this is not set, you will see a warning in your logs that your session is not secured.

For more information on this see the docs for [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).





## Storing Complex Types

It is generally considered **not** good practice to store complex types in a session. There are lots of reasons for this, but it is recommended to store the ID of a type, instead of the "whole" value.

Should you need to store a complex type, like a `struct` you will first need to register the type with the [`encoding/gob`](https://golang.org/pkg/encoding/gob/) package.

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

## Saving a Session

Buffalo automatically saves your session for you, so you don't have to. If there is an error when saving the session, Buffalo will return an error through the normal [ error handling ](/en/docs/errors) process.


## Null Sessions for APIs

When building API servers the default cookie session store is undesirable. The [`sessions.Null`](`sessions.Null`) type is the recommended replacement for the default session store.

```go
app = buffalo.New(buffalo.Options{
  Env:          ENV,
  SessionStore: sessions.Null{},
  SessionName: "_coke_session",
})
```

When running `buffalo new` with the `--api` flag the default session will be set to `sessions.Null`.


