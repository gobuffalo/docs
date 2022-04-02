---
name: API Applications
---

# API Applications

Applications that only serve API end-points, typically JSON, are very different from those that serve HTML, JavaScript, and CSS. In this guide, you'll learn how to build an API-only app, using Buffalo.

## Creating a new API Application

When creating a new Buffalo application using the `buffalo new` command, the optional `--api` flag will generate an application that is better suited to serving APIs than a stock Buffalo application.

```bash
$ buffalo new --api coke
```

### Slimmed Project Layout

Applications generated with the `--api` flag don't contain any front systems. This means there is no templating, stylesheets, etc...

#### <code>buffalo new coke --api</code>
```erb
├── Dockerfile
├── README.md
├── actions
│   ├── actions_test.go
│   ├── app.go
│   ├── home.go
│   ├── home_test.go
│   └── render.go
├── database.yml
├── fixtures
│   └── sample.toml
├── grifts
│   ├── db.go
│   └── init.go
├── inflections.json
├── main.go
└── models
    ├── models.go
    └── models_test.go
```

#### <code>buffalo new coke</code>
```erb
├── Dockerfile
├── README.md
├── actions
│   ├── actions_test.go
│   ├── app.go
│   ├── home.go
│   ├── home_test.go
│   └── render.go
├── assets
│   ├── css
│   │   └── application.scss
│   ├── images
│   │   ├── favicon.ico
│   │   └── logo.svg
│   └── js
│       └── application.js
├── database.yml
├── fixtures
│   └── sample.toml
├── grifts
│   ├── db.go
│   └── init.go
├── inflections.json
├── locales
│   └── all.en-us.yaml
├── main.go
├── models
│   ├── models.go
│   └── models_test.go
├── node_modles
├── package.json
├── public
│   ├── assets
│   │   └── .keep
│   └── robots.txt
├── templates
│   ├── _flash.html
│   ├── application.html
│   └── index.html
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
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}
	return app
}
```

```go
func init() {
	r = render.New(render.Options{
		HTMLLayout:   "application.html",
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,
		Helpers:      render.Helpers{},
	})
}
```
