---
name: API Applications
weight: 10
aliases:
  - /pt/docs/apis
  - /pt/docs/apis
---

# API Applications

Applications that only serve API end-points, typically JSON, are very different from those that serve HTML, JavaScript, and CSS. In this guide, you'll learn how to build an API-only app, using Buffalo.

## Creating a New API Application

When creating a new Buffalo application using the `buffalo new` command, the optional `--api` flag will generate an application that is better suited to serving APIs than a stock Buffalo application.

```bash
$ buffalo new coke --api
```

### Slimmed Project Layout

Applications generated with the `--api` flag don't contain any front systems. This means there is no templating, stylesheets, etc...

{{<codetabs>}}
{{<tab "API">}}
```bash
$ buffalo new coke --api
```

```erb
├── actions/
│	├── app.go
│	└── render.go
├── cmd/
│	└── app/
│		└── main.go
├── config/
├── fixtures/
├── grifts/
├── locales/
├── models/
├── .buffalo.dev.yml
├── .codeclimate.yml
├── .docketignore
├── .env
├── .gitignore
├── database.yml
├── Dockerfile
├── go.mod
├── go.sum
├── inflections.json
├── README.md
```
{{</tab>}}
{{<tab "Default">}}
```bash
$ buffalo new coke
```

```erb
├── .yarn/
├── actions/
│	├── app.go
│	└── render.go
├── assets/
├── cmd/
│	└── app/
│		└── main.go
├── config/
├── fixtures/
├── grifts/
├── locales/
├── models/
├── public/
├── templates/
├── .babelrc
├── .buffalo.dev.yml
├── .codeclimate.yml
├── .docketignore
├── .env
├── .gitignore
├── .pnp.cjs
├── .pnp.loader.mjs
├── .yarnrc.yml
├── database.yml
├── Dockerfile
├── go.mod
├── go.sum
├── inflections.json
├── package.json
├── postcss.config.js
├── README.md
├── webpack.config.js
└── yarn.lock
```
{{</tab>}}
{{</codetabs>}}

### Tuned `actions/app.go actions/render.go` Files

API applications have `actions/app.go` and `actions/render.go` files that are a good starting point for API applications.

{{<codetabs>}}
{{<tab "API">}}
```bash
$ buffalo new coke --api
```

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
		app.Use(paramlogger.ParameterLogger)
		app.Use(contenttype.Set("application/json"))

		app.Use(popmw.Transaction(models.DB))
		app.GET("/", HomeHandler)
	}

	return app
}
```

```go
func init() {
	r = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}
```
{{</tab>}}
{{<tab "Default">}}
```bash
$ buffalo new coke
```

```go
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})
		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(csrf.New)
		app.Use(popmw.Transaction(models.DB))
		app.Use(translations())
		app.GET("/", HomeHandler)
		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}
```

```go
func init() {
	r = render.New(render.Options{
		HTMLLayout: "application.plush.html",
		TemplatesFS: templates.FS(),
		AssetsFS: public.FS(),
		Helpers: render.Helpers{},
	})
}
```
{{</tab>}}
{{</codetabs>}}
