<%= title("Tasks") %>

Easily generate new [https://github.com/markbates/grift](https://github.com/markbates/grift) tasks.

```bash
$ buffalo g task foo:bar

--> grifts/bar.go
```

```go
// grifts/bar.go
package grifts

import (
  . "github.com/markbates/grift/grift"
)

var _ = Namespace("foo", func() {

  Desc("bar", "TODO")
  Add("bar", func(c *Context) error {
    return nil
  })

})
```
