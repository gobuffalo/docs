package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/x/sessions"
	"github.com/unrolled/secure"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// T is used to provide translations
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			SessionName:  "_gobuffalo_session",
			Env:          ENV,
			SessionStore: sessions.Null{},
		})
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", "0.10.3")
				c.Set("goMinVersion", "1.8.1")
				c.Set("year", time.Now().Year())
				c.Set("trainingURL", "http://www.gopherguides.com")

				c.Set("lang", "en")
				if lang, err := c.Cookies().Get("lang"); err == nil {
					c.Set("lang", lang)
				} else {
					// Ensure the correct language is set, even without the cookie
					langs := c.Value("languages").([]string)
					for _, l := range langs {
						if l == "fr" || l == "en" {
							c.Set("lang", l)
							break
						}
					}
				}
				return next(c)
			}
		})

		app.GET("/search", Search)
		app.Redirect(302, "/docs/overview", "/")
		app.Redirect(302, "/docs/repl", "/")
		app.Redirect(302, "/docs/test-suites", "/docs/testing")
		app.Redirect(302, "/docs/env-vars", "/docs/config-vars")
		app.GET("/docs/{name:.+}", Docs)

		app.ServeFiles("/assets", assetBox)
		app.POST("/lang", ChangeLanguage)
		app.GET("/sponsors", func(c buffalo.Context) error {
			return c.Render(200, r.HTML("sponsors.html"))
		})
		app.GET("/", HomeHandler)
		indexDocs(app)
	}
	return app
}
