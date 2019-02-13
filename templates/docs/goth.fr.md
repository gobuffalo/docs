<% seoDescription("Authentification tierce") %>
<% seoKeywords(["buffalo", "go", "golang", "users", "oauth2", "authentification", "tierce", "google", "facebook"]) %>

<%= h1("Authentification tierce") %>

Dans bien des cas, vous allez devoir implémenter un méchanisme d'authentification dans vos applications. [Goth](https://github.com/markbates/goth) fournit un moyen simple, propre et idiomatique d'écrire des paquets d'authentification pour les applications Web en Go.

Si vous cherchez à intégrer la connexion via Facebook, Google ou autres, c'est sans doute la solution que vous recherchez.

<%= note() { %>
Buffalo fournissait un support native en standard pour Goth jusqu'à la version `v0.9.4`. Depuis, ce support à été déplacé dans un plugin à part, [https://github.com/gobuffalo/buffalo-goth](https://github.com/gobuffalo/buffalo-goth).
<% } %>

## Installation

Pour installer le plugin `buffalo-goth`, lancez le commande suivante :

```bash
$ go get -u github.com/gobuffalo/buffalo-goth
```

## Générateur
```bash
$ buffalo g goth twitter facebook linkedin github

--> actions/auth.go
--> go get github.com/markbates/goth/...
--> goimports -w .
```

## Exemple d'utilisation
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
  // Faire quelque chose avec l'utilisateur, par exemple créer un compte associé ou le connecter.
  return c.Render(200, r.JSON(user))
}
```

## Tutoriel vidéo (EN)

<%= vimeo("223666374") %>

## Voir aussi

* [Authentification locale](/fr/docs/auth) - Gérer l'authentification d'utilisateurs internes.