<%= title("Tasks") %>

Easily generate new [https://github.com/markbates/grift](https://github.com/markbates/grift) tasks.

<div class="code-tabs">

<%= code("bash", {file: "command"}) { %>
$ buffalo g task foo:bar

--> grifts/bar.go
<% } %>

<%= code("go", {file: "grifts/bar.go"}) { %>
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
<% } %>

</div>
