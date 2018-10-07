<%= title("Skipping Resource Actions") %>

Often it is necessary to want to skip middleware for one or more actions. For example, allowing guest users to view the `List` and `Show` actions on a resource, but requiring authorization on the rest of the actions.

Understanding from the [Skipping Middleware](#skipping-middleware) section we need to make sure that we are using the same functions when we register the resource as we do when we want to skip the middleware on those functions later.

The line that was generated in `actions/app.go` by `buffalo generate resource` will need to be changed to accommodate this requirement.

<%= codeTabs() { %>
```go
// BEFORE
app.Resource("/widgets", WidgetResource{})
```

```go
// AFTER
res := WidgetResource{}
wr := app.Resource("/widgets", res)
wr.Middleware.Skip(Authorize, res.Index, res.Show)
```
<% } %>
