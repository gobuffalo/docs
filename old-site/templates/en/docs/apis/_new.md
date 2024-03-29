## Creating a new API Application

When creating a new Buffalo application using the `buffalo new` command, the optional `--api` flag will generate an application that is better suited to serving APIs than a stock Buffalo application.

```bash
$ buffalo new --api coke
```

### Slimmed Project Layout

Applications generated with the `--api` flag don't contain any front systems. This means there is no templating, stylesheets, etc...

<div class="row">
    <div class="col-md-6 col-sm-12">
        <h5><code>buffalo new coke --api</code></h5>
        <%= partial("en/docs/apis/api_ls.md") %>
    </div>
    <div class="col-md-6 col-sm-12">
        <h5><code>buffalo new coke</code></h5>
        <%= partial("en/docs/apis/web_ls.md") %>
    </div>
</div>

### Tuned `actions/app.go actions/render.go` Files

API applications have `actions/app.go` and `actions/render.go` files that are a good starting point for API applications.

<h5><code>buffalo new coke --api</code></h5>

```go
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_coke_session",
		})
		app.Use(forceSSL())
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		app.Use(middleware.PopTransaction(models.DB))
		app.GET("/", HomeHandler)
	}
	return app
}
```

```go
func init() {
	r = render.New(render.Options{})
}
```

<h5><code>buffalo new coke</code></h5>

<%= partial("en/docs/apis/web_app.md") %> 
<%= partial("en/docs/apis/web_render.md") %>