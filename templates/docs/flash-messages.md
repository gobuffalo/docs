# Flash Messages [![GoDoc](https://godoc.org/github.com/gobuffalo/buffalo?status.svg)](https://godoc.org/github.com/gobuffalo/buffalo#Flash)

<%= partial("docs/disclaimer.html") %>

<%= partial("topics.html") %>

<%= panel("What are Flash Messages?", {name:"what-is-flash"}) { %>

Flash messages are a means of communicating messages to the end user from inside of an application. These messages might be errors, warnings, or success types of messages.

Some examples of flash messages are:

* "You have been successfully logged out."
* "Your widget could not be updated."
* "There was a problem accessing your account."

Being able to set these messages in a Buffalo handler and then pass them down to views is incredibly helpful.

<% } %>

<%= panel("Setting Flash Messages", {}) { %>

Creating flash messages can easily be done by using the `c.Flash()` function provided on the [`buffalo.Context`](/docs/context).

```go
func WidgetsCreate(c buffalo.Context) error {
  // do some work
	c.Flash().Add("success", "Widget was successfully created!")
  // do more work and return
}
```

The names of the "keys", in this example, "success", are left up to your application to use as is appropriate. There are "special" or "pre-defined" keys.

<% } %>

<%= panel("Accessing Flash Messages in Templates", {name: "accessing-in-templates"}) { %>

```handlebars
<!-- loop through all of the flash messages  -->
{{#each flash as |k messages|}}
  <!-- messages is a slice, loop through each  -->
  {{#each messages as |msg|}}
    <div class="alert alert-{{k}}" role="alert">
      <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
      {{msg}}
    </div>
  {{/each}}
{{/each}}
```

```handlebars
<!-- loop through a specific key  -->
{{#each flash.success as |msg|}}
  <div class="alert alert-{{k}}" role="alert">
    <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    {{msg}}
  </div>
{{/each}}
```

<% } %>
