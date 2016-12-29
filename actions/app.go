package actions

import (
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/gobuffalo/models"
	"github.com/markbates/going/defaults"
)

var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})

	a.Use(middleware.NewRelic(os.Getenv("NEW_RELIC_LICENSE_KEY"), "gobuffalo.io"))
	a.Use(middleware.PopTransaction(models.DB))
	a.Use(SetVersion)
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)
	a.GET("/docs/{name}", Docs)

	// ensure this is a JSON request
	g := a.Group("/version")
	g.Use(middleware.SetContentType("application/json"))
	g.GET("/", VersionList)
	g.GET("/current", VersionCurrent)
	g.POST("/", VersionUpdate)

	return a
}
