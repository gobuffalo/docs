<%= title("JavaScript") %>
<%= sinceVersion("0.10.0") %>

The [`render`](https://godoc.org/github.com/gobuffalo/buffalo/render) package has a new implementation of [`render.Renderer`](https://godoc.org/github.com/gobuffalo/buffalo/render#Renderer), [`render.JavaScript`](https://godoc.org/github.com/gobuffalo/buffalo/render#JavaScript).

This means inside of an action you can do the following:

```go
func HomeHandler(c buffalo.Context) error {
  return c.Render(200, r.JavaScript("index.js"))
}
```

The [`render.Options`](https://godoc.org/github.com/gobuffalo/buffalo/render#Options) type now has a new attribute, `JavaScriptLayout`. This new option is similar to the `HTMLLayout` option in that it will wrap `*.js` files inside of another `*.js`.

The new JavaScript renderer also has it’s own implementation of the `partial` function. This new implementation behaves almost the same as the original implementation, but is smart enough to know that if you are rendering an `*.html` file inside of a `*.js` file that it will need to be escaped properly, and so it does it for you.

```javascript
$("#new-goal-form").replaceWith("&lt;%= partial("goals/new.html") %&gt;");
```

