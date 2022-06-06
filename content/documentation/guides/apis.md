---
name: API Applications
weight: 10
aliases:
  - /docs/apis
  - /en/docs/apis
---

# API Applications

Applications that only serve API end-points, typically JSON, are very different from those that serve HTML, JavaScript, and CSS. In this guide, you'll learn how to build an API-only app, using Buffalo.

## Creating a New API Application

When creating a new Buffalo application using the `buffalo new` command, the optional `--api` flag will generate an application that is better suited to serving APIs than a stock Buffalo application.

```bash
$ buffalo new --api coke
```

### Slimmed Project Layout

Applications generated with the `--api` flag don't contain any front systems. This means there is no templating, stylesheets, etc...

#### <code>buffalo new coke --api</code>

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

#### <code>buffalo new coke</code>

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

```go
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})
		app.Use(forceSSL())
		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}
		app.Use(csrf.New)
		app.Use(middleware.PopTransaction(models.DB))
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
		HTMLLayout:   "application.html",
		TemplatesFS:  templates.FS(),
		AssetsFS:     public.FS(),
		Helpers:      render.Helpers{},
	})
}
```
