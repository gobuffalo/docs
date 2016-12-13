package actions

import (
	"net/http"
	"os"

	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/markbates/gobuffalo/models"
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
	a.Use(func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			// TODO: get from the db
			c.Set("version", "0.4.5")
			return next(c)
		}
	})
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)
	a.GET("/docs/{name}", Docs)

	return a
}
