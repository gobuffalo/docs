---
name: Aplicaciones API
weight: 10
aliases:
  - /docs/apis
  - /es/docs/apis
---

# Aplicaciones API

Las aplicaciones que solo sirven end-points de API, normalmente JSON, son muy diferentes de las que sirven HTML, JavaScript y CSS. En esta guía, aprenderás a crear una aplicación solo de API con Buffalo.

## Creando una nueva aplicacion API

Cuando creamos una nueva aplicacion de Buffalo usando el comando `buffalo new`, el flag opcional `--api` generará una aplicación que se adapta mejor para API que para una aplicación por defecto de Buffalo.

```bash
$ buffalo new coke --api
```

### Estructura simplificada del proyecto

Las aplicaciones creadas con el flag `--api` no contiene ningun sistema frontal. Es decir, que no hay plantillas, hojas de estilo, etc...

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

### Archivos `actions/app.go actions/render.go` ajustados

Las aplicaciones API tienen los archivos `actions/app.go` and `actions/render.go`, que son un buen punto de inicio para estas.

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