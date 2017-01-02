package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/gobuffalo/models"
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

		app.Use(middleware.NewRelic(os.Getenv("NEW_RELIC_LICENSE_KEY"), "gobuffalo.io"))
		app.Use(middleware.PopTransaction(models.DB))
		app.Use(SetVersion)
		app.GET("/", HomeHandler)
		app.GET("/docs/{name}", Docs)

		// ensure this is a JSON request
		g := app.Group("/version")
		g.Use(middleware.SetContentType("application/json"))
		g.GET("/", VersionList)
		g.GET("/current", VersionCurrent)
		g.POST("/", VersionUpdate)

		app.ServeFiles("/assets", assetsPath())
	}
	return app
}
