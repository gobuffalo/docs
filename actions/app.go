package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			SessionName: "_gobuffalo_session",
			Env:         ENV,
		})

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", "0.9.2")
				c.Set("year", time.Now().Year())
				c.Set("trainingURL", "http://www.gopherguides.com")
				return next(c)
			}
		})

		app.GET("/docs/{name:.+}", Docs)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
		app.Redirect(302, "/", "/docs/getting-started")
	}
	return app
}
