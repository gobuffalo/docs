package actions

import (
	"log"
	"net/http"
	"os"

	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/markbates/gobuffalo/models"
	"github.com/markbates/going/defaults"
	newrelic "github.com/newrelic/go-agent"
)

var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})

	nrkey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if nrkey != "" {
		config := newrelic.NewConfig("gobuffalo.io", nrkey)
		app, err := newrelic.NewApplication(config)
		if err != nil {
			log.Fatal(err)
		}
		a.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				tx := app.StartTransaction(c.Request().URL.String(), c.Response(), c.Request())
				defer tx.End()
				return next(c)
			}
		})
	}

	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)

	return a
}
