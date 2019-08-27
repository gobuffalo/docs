package actions

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/docs/search/vimeo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/x/sessions"
	"github.com/pkg/errors"
	"github.com/unrolled/secure"
)

var app *buffalo.App
var supportedLanguages = map[string]string{
	"English":  "en",
	"Español":  "es",
	"Français": "fr",
	"Italiano": "it",
}

// T is used to provide translations
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		env := envy.Get("GO_ENV", "development")
		app = buffalo.New(buffalo.Options{
			SessionName:  "_gobuffalo_session",
			Env:          env,
			SessionStore: sessions.Null{},
		})
		defer func() {
			go StartSearch(app)
		}()

		// Automatically redirect to SSL
		app.Use(forceSSL(env))

		if env == "development" {
			app.Use(paramlogger.ParameterLogger)
		}

		// Setup and use translations:
		app.Use(translations())

		langCache := map[string]struct{}{}
		for _, v := range supportedLanguages {
			langCache[v] = struct{}{}
		}

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", buffaloVersion)
				c.Set("goMinVersion", "1.10.8")
				c.Set("year", time.Now().Year())
				c.Set("trainingURL", "http://www.gopherguides.com")
				c.Set("videoList", vimeo.Videos())

				c.Set("current_path", strings.TrimRight(c.Value("current_path").(string), "/"))
				langs, ok := c.Value("languages").([]string)
				if !ok {
					return errors.New("could not get user languages")
				}
				c.Set("lang", "en")
				for _, l := range langs {
					if _, ok := langCache[l]; ok {
						c.Set("lang", l)
						break
					}
				}

				c.Set("supported_languages", supportedLanguages)

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

		enableProfiling(app)
		bindRedirects(app)

		app.POST("/lang", ChangeLanguage)

		languagesKey := createLanguagesKey()

		app.GET(fmt.Sprintf("/{lang:%s}/search", languagesKey), Search)
		app.GET(fmt.Sprintf("/{lang:%s}/docs/{name:.+}", languagesKey), Docs)
		app.GET(fmt.Sprintf("/{lang:%s}/sponsors", languagesKey), Sponsors)
		app.GET(fmt.Sprintf("/{lang:%s}", languagesKey), HomeHandler)

		app.ServeFiles("/", assetBox)
	}
	return app
}

func createLanguagesKey() string {
	allKeys := ""
	for _, code := range supportedLanguages {
		allKeys = fmt.Sprintf("%s|%s", allKeys, code)
	}
	return allKeys
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("../locales", "../locales"), "en"); err != nil {
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
func forceSSL(env string) buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     env == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func bindRedirects(app *buffalo.App) {
	// Deleted pages
	app.GET("/en/docs/generators", func(c buffalo.Context) error {
		return c.Render(http.StatusGone, nil)
	})

	// Remapped pages
	redirectsTable := []struct {
		From string
		To   string
	}{
		{
			From: "/docs/db",
			To:   "/docs/db/getting-started",
		},
		{
			From: "/docs/systemd",
			To:   "/docs/deploy/systemd",
		},
		{
			From: "/docs/proxy",
			To:   "/docs/deploy/proxy",
		},
		{
			From: "/docs/building",
			To:   "/docs/deploy/building",
		},
		{
			From: "/docs/installation",
			To:   "/docs/getting-started/installation",
		},
		{
			From: "/docs/integrations",
			To:   "/docs/getting-started/integrations",
		},
		{
			From: "/docs/new-project",
			To:   "/docs/getting-started/new-project",
		},
		{
			From: "/docs/directory-structure",
			To:   "/docs/getting-started/directory-structure",
		},
		{
			From: "/docs/config-vars",
			To:   "/docs/getting-started/config-vars",
		},
	}

	for _, route := range redirectsTable {
		func(from string, to string) {
			app.GET(from, func(c buffalo.Context) error {
				return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s%s", c.Value("lang").(string), to))
			})
			app.GET(fmt.Sprintf("/{lang:fr|en}%s", from), func(c buffalo.Context) error {
				return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s%s", c.Value("lang").(string), to))
			})
		}(route.From, route.To)
	}

	// Specific cases
	app.GET("/search", func(c buffalo.Context) error {
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/search", c.Value("lang").(string)))
	})
	app.GET("/sponsors", func(c buffalo.Context) error {
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/sponsors", c.Value("lang").(string)))
	})
	app.GET("/docs/{name:.+}", func(c buffalo.Context) error {
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/docs/%s", c.Value("lang").(string), c.Param("name")))
	})
	app.GET("/", func(c buffalo.Context) error {
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/", c.Value("lang").(string)))
	})
}

func enableProfiling(app *buffalo.App) {
	app.GET("/debug/pprof/", buffalo.WrapHandlerFunc(pprof.Index))
	app.GET("/debug/pprof/block", buffalo.WrapHandler(pprof.Handler("block")))
	app.GET("/debug/pprof/goroutine", buffalo.WrapHandler(pprof.Handler("goroutine")))
	app.GET("/debug/pprof/heap", buffalo.WrapHandler(pprof.Handler("heap")))
	app.GET("/debug/pprof/mutex", buffalo.WrapHandler(pprof.Handler("mutex")))
	app.GET("/debug/pprof/threadcreate", buffalo.WrapHandler(pprof.Handler("threadcreate")))
	app.GET("/debug/pprof/profile", buffalo.WrapHandler(pprof.Handler("profile")))
	app.GET("/debug/pprof/symbol", buffalo.WrapHandler(pprof.Handler("symbol")))
	app.GET("/debug/pprof/trace", buffalo.WrapHandler(pprof.Handler("trace")))
}
