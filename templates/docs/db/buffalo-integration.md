<% seoDescription("How to use Pop with Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "pop", "buffalo", "integration"]) %>

<%= h1("Buffalo Integration") %>

Buffalo provides a Pop middleware to ease database usage within Buffalo: https://github.com/gobuffalo/buffalo-pop

<%= title("The Pop Transaction Middleware") %>

### Setup

This middleware is configured for you by default, if you choose to use Pop when creating a new project.

**actions/app.go**

```go
func App() *buffalo.App {
	if app == nil {
        // [...]

        app.Use(poptx.PopTransaction(models.DB))

        // [...]

        app.GET("/", HomeHandler)
    }

    return app
}
```

`poptx.PopTransaction(models.DB)` uses the connection to the configured database to create a new `PopTransaction` middleware. This middleware does the following:

* Log the total duration spent during the request making database calls.
* Wrap **each HTTP request** in a database transaction.
* Commit **if there was no error** executing the middlewares and action; **and the response status is a 2xx or 3xx**.
* Rollback otherwise.

### Handle Transaction By Hand

If you need to handle a transaction by hand, you can skip the middleware for a given route:

```go
func App() *buffalo.App {
	if app == nil {
        // [...]
        txm := poptx.PopTransaction(models.DB)
        app.Use(txm)
        a.Middleware.Skip(txm, HomeHandler)

        // [...]

        app.POST("/form", FormHandler)
        app.GET("/", HomeHandler)
    }

    return app
}
```