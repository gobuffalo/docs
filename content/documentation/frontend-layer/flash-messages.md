---
Name: "Flash Messages"
---

# Flash Messages

## What are Flash Messages?

Flash messages are a means of communicating messages to the end user from inside of an application. These messages might be errors, warnings, or success types of messages.

Some examples of flash messages are:

* "You have been successfully logged out."
* "Your widget could not be updated."
* "There was a problem accessing your account."

Being able to set these messages in a Buffalo handler and then pass them down to views is incredibly helpful.

## Setting Flash Messages

Creating flash messages can easily be done by using the `c.Flash()` function provided on the [`buffalo.Context`](/en/context).

```go
func WidgetsCreate(c buffalo.Context) error {
  // do some work
  c.Flash().Add("success", "Widget was successfully created!")
  // do more work and return
}
```

The names of the "keys", in this example, "success", are left up to your application to use as is appropriate. There are no "special" or "pre-defined" keys.

## Accessing Flash Messages in Templates

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

### Looping Over all Flash Messages

```html
&lt;div class="row">
  &lt;div class="col-md-12">
    \<%= for (k, messages) in flash { %>
      \<%= for (msg) in messages { %>
        &lt;div class="alert alert-\<%= k %>" role="alert">
          &lt;button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          \<%= msg %>
        &lt;/div>
      \<% } %>
    \<% } %>
  &lt;/div>
&lt;/div>
```

### Looping Over a Specific Flash Message Key

```html
&lt;div class="row">
  &lt;div class="col-md-12">
    \<%= for (message) in flash["success"] { %>
      &lt;div class="alert alert-success" role="alert">
        &lt;button type="button" class="close" data-dismiss="alert" aria-label="Close">&lt;span aria-hidden="true">&times;</span></button>
        \<%= message %>
      &lt;/div>
    \<% } %>
  &lt;/div>
&lt;/div>
```
