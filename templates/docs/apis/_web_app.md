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
