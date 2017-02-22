package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/going/defaults"
)

var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env: ENV,
		})

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", "0.7.3")
				return next(c)
			}
		})
		app.GET("/", HomeHandler)
		app.GET("/docs/{name}", Docs)

		app.ServeFiles("/assets", assetsPath())
	}
	return app
}
