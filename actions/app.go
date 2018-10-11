package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/gobuffalo/search/vimeo"
	"github.com/gobuffalo/mw-forcessl"
	"github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/x/sessions"
	"github.com/unrolled/secure"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var supportedLanguages = map[string]string{
	"en": "English",
	"fr": "Français",
}

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
			app.Use(paramlogger.ParameterLogger)
		}

		// Setup and use translations:
		app.Use(translations())

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", buffaloVersion)
				c.Set("goMinVersion", "1.9.7")
				c.Set("year", time.Now().Year())
				c.Set("trainingURL", "http://www.gopherguides.com")
				c.Set("videoList", vimeo.Videos)

				c.Set("lang", "en")
				c.Set("current_path", strings.TrimRight(c.Value("current_path").(string), "/"))
				langs := c.Value("languages").([]string)
				for _, l := range langs {
					if _, ok := supportedLanguages[l]; ok {
						c.Set("lang", l)
						break
					}
				}

				c.Set("localized_current_path", func(lang string) string {
					cp := c.Value("current_path").(string)
					l := c.Param("lang")
					tp := strings.TrimPrefix(cp, fmt.Sprintf("/%s", l))
					if lang == "" {
						return tp
					}
					return fmt.Sprintf("/%s%s", lang, tp)
				})

				return next(c)
			}
		})

		app.GET("/docs/db", func(c buffalo.Context) error {
			return c.Redirect(301, fmt.Sprintf("/%s/docs/db/getting-started", c.Value("lang").(string)))
		})
		app.GET("/{lang:fr|en}/docs/db", func(c buffalo.Context) error {
			return c.Redirect(301, fmt.Sprintf("/%s/docs/db/getting-started", c.Value("lang").(string)))
		})

		// Patch old URLs
		oldURLs := []string{"systemd", "proxy", "building"}
		for _, url := range oldURLs {
			app.GET(fmt.Sprintf("/docs/%s", url), func(c buffalo.Context) error {
				return c.Redirect(301, fmt.Sprintf("/%s/docs/deploy/%s", c.Value("lang").(string), url))
			})
			app.GET(fmt.Sprintf("/{lang:fr|en}/docs/%s", url), func(c buffalo.Context) error {
				return c.Redirect(301, fmt.Sprintf("/%s/docs/deploy/%s", c.Value("lang").(string), url))
			})
		}

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
		app.GET("/{lang:fr|en}", HomeHandler)

		app.ServeFiles("/", assetBox)
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
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
