package actions

import (
	"fmt"
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

		// Automatically redirect to SSL
		app.Use(forceSSL())

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Setup and use translations:
		app.Use(translations())

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", "0.12.0")
				c.Set("goMinVersion", "1.8.1")
				c.Set("year", time.Now().Year())
				c.Set("trainingURL", "http://www.gopherguides.com")

				c.Set("lang", "en")
				langs := c.Value("languages").([]string)
				for _, l := range langs {
					if l == "fr" || l == "en" {
						c.Set("lang", l)
						break
					}
				}
				return next(c)
			}
		})

		app.Redirect(302, "/docs/overview", "/")
		app.Redirect(302, "/docs/repl", "/")
		app.Redirect(302, "/docs/test-suites", "/docs/testing")
		app.Redirect(302, "/docs/env-vars", "/docs/config-vars")
		app.GET("/search", func(c buffalo.Context) error {
			return c.Redirect(302, fmt.Sprintf("/%s/search", c.Value("lang").(string)))
		})
		app.GET("/sponsors", func(c buffalo.Context) error {
			return c.Redirect(302, fmt.Sprintf("/%s/sponsors", c.Value("lang").(string)))
		})
		app.GET("/docs/{name:.+}", func(c buffalo.Context) error {
			return c.Redirect(302, fmt.Sprintf("/%s/docs/%s", c.Value("lang").(string), c.Param("name")))
		})
		app.GET("/", func(c buffalo.Context) error {
			return c.Redirect(302, fmt.Sprintf("/%s", c.Value("lang").(string)))
		})

		app.GET("/{lang:fr|en}/search", Search)
		app.GET("/{lang:fr|en}/docs/{name:.+}", Docs)

		app.POST("/lang", ChangeLanguage)
		app.GET("/{lang:fr|en}/sponsors", Sponsors)
		app.GET("/{lang:fr|en}/", HomeHandler)

		app.ServeFiles("/", assetBox)

		indexDocs(app)
		go func() {
			indexBlog(app)
			t := time.NewTicker(60 * time.Minute)
			defer t.Stop()
			for {
				select {
				case <-app.Context.Done():
					return
				case <-t.C:
					indexBlog(app)
				}
			}
		}()
	}
	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en"); err != nil {
		app.Stop(err)
	}

	T.LanguageExtractors = []i18n.LanguageExtractor{
		i18n.URLPrefixLanguageExtractor,
		i18n.HeaderLanguageExtractor,
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
