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
