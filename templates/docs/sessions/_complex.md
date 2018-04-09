<%= title("Storing Complex Types") %>

It is generally considered **not** good practice to store complex types in a session. There are lots of reasons for this, but it is recommended to store the ID of a type, instead of the "whole" value.

Should you need to store a complex type, like a `struct` you will first need to register the type with the [`encoding/gob`](https://golang.org/pkg/encoding/gob/) package.

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

