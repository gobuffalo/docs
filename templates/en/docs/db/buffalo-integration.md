<% seoDescription("How to use Pop with Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "pop", "buffalo", "integration"]) %>

# Buffalo Integration

## Generate a new app

When you generate a new Buffalo application you can choose the target database with the `--db-type` flag. For instance, to generate a new app with MySQL database support, you can write the following:

```bash
$ buffalo new coke --db-type mysql
```

**By default, Buffalo will generate an app with PostgreSQL as the backing database.**

### Skip database support

If you want to handle the database without using Pop, or if you're building an app without database, it's also possible to skip generation of all database components with the `--skip-pop` flag.

```bash
$ buffalo new coke --skip-pop
```

## The Pop Transaction Middleware

Buffalo provides a Pop middleware to ease database usage within Buffalo: https://github.com/gobuffalo/buffalo-pop

### Setup

This middleware is configured for you by default, if you choose to use Pop when creating a new project.

```bash
$ go get github.com/gobuffalo/buffalo-pop
```

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
