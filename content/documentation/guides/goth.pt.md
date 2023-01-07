---
name: Third Party Authentication
seoDescription: Third Party Authentication
seoKeywords: ["buffalo", "go", "golang", "users", "oauth2", "authentication", "third-party", "google", "facebook"]
weight: 22
aliases:
  - /docs/goth
  - /pt/docs/goth
---

# Third Party Authentication

In many use-cases, you'll need to implement user authentication in your apps. [Goth](https://github.com/markbates/goth) provides a simple, clean, and idiomatic way to write authentication packages for Go web applications.

If you're looking for authentication through Facebook, Google and others, that's probably the solution you're looking for.

{{< note >}}
Buffalo had a native support for Goth until version `v0.9.4`. Since then, it was moved into it's own plugin, [https://github.com/gobuffalo/buffalo-goth](https://github.com/gobuffalo/buffalo-goth).
{{< /note >}}

## Installation

To install the `buffalo-goth` plugin, run the following command:

```bash
$ buffalo plugins install github.com/gobuffalo/buffalo-goth
```

## Generator
```bash
$ buffalo g goth twitter facebook linkedin github

--> actions/auth.go
--> go get github.com/markbates/goth/...
--> goimports -w .
```

## Example Usage
```go
// actions/app.go
package actions

import (
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/buffalo/middleware"
  "github.com/gobuffalo/buffalo/middleware/csrf"
  "github.com/gobuffalo/buffalo/middleware/i18n"

  "github.com/markbates/coke/models"

  "github.com/gobuffalo/envy"
  "github.com/gobuffalo/packr"

  "github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
  if app == nil {
    app = buffalo.New(buffalo.Options{
      Env:         ENV,
      SessionName: "_coke_session",
    })

    if ENV == "development" {
      app.Use(middleware.ParameterLogger)
    }
    if ENV != "test" {
      // Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
      // Remove to disable this.
      app.Use(csrf.Middleware)
    }

    // Wraps each request in a transaction.
    //  c.Value("tx").(*pop.PopTransaction)
    // Remove to disable this.
    app.Use(middleware.PopTransaction(models.DB))

    // Setup and use translations:
    var err error
    if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
      app.Stop(err)
    }
    app.Use(T.Middleware())

    app.GET("/", HomeHandler)

    app.ServeFiles("/assets", packr.NewBox("../public/assets"))
    auth := app.Group("/auth")
    auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
    auth.GET("/{provider}/callback", AuthCallback)
  }

  return app
}
```

```go
// actions/auth.go
package actions

import (
  "fmt"
  "os"

  "github.com/gobuffalo/buffalo"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/facebook"
  "github.com/markbates/goth/providers/github"
  "github.com/markbates/goth/providers/linkedin"
  "github.com/markbates/goth/providers/twitter"
)

func init() {
  gothic.Store = App().SessionStore

  goth.UseProviders(
    twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/twitter/callback")),
    facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/facebook/callback")),
    linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/linkedin/callback")),
    github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
  )
}

func AuthCallback(c buffalo.Context) error {
  user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
  if err != nil {
    return c.Error(401, err)
  }
  // Do something with the user, maybe register them/sign them in
  return c.Render(200, r.JSON(user))
}
```

## Video Tutorial

{{< vimeo 223666374>}}

## See Also

* [Local Authentication](/documentation/guides/auth) - Manage internal users auth.
