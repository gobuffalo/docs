# Contexte

Dans toutes les actions Buffalo, on retrouve le paramètre `Context`. Ce contexte permet d'avoir une définition simple et efficace de vos actions, tout en étant puissante.

## L'interface Context

L'interface `buffalo.Context` embarque `context.Context`, elle peut donc être utilisée comme un `Context` Go standard.

`buffalo.Context` étant une interface, il est possible de créer une implémentation spécifique qui répond de manière plus précise aux besoins de votre application.

<%= sinceVersion("0.12.0") %>

```go
type Context interface {
  context.Context
  Response() http.ResponseWriter
  Request() *http.Request
  Session() *Session
  Cookies() *Cookies
  Params() ParamValues
  Param(string) string
  Set(string, interface{})
  LogField(string, interface{})
  LogFields(map[string]interface{})
  Logger() Logger
  Bind(interface{}) error
  Render(int, render.Renderer) error
  Error(int, error) error
  Redirect(int, string, ...interface{}) error
  Data() map[string]interface{}
  Flash() *Flash
  File(string) (binding.File, error)
}
```

La méthode `Websocket() (*websocket.Conn, error)` a été retirée de `buffalo.Context` en version `v0.12.0`. Utilisez le paquet [http://www.gorillatoolkit.org/pkg/websocket](http://www.gorillatoolkit.org/pkg/websocket) directement à la place.

## Contexte et rendu

Dans l'interface `Context`, on peut trouver une méthode `Render` de type `render.Renderer`. Consultez le chapitre sur le [rendu](/docs/fr/rendering) pour plus d'informations.

Toutes les valeurs données au contexte sont automatiquement disponibles pour le `render.Renderer` passé à la méthode `Render`.

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")
  return c.Render(200, render.String("Hi \<%= name %>"))
}
```

## Implémenter l'interface Context

L'interface `buffalo.Context` n'est pas faite pour être « pleinement » implémentée. Il est recommandé à la place d'utiliser la [composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) et d'implémenter uniquement les méthodes que vous souhaitez réécrire.

Voici un exemple illustrant la réécriture de la méthode `Error` pour logger l'erreur et tuer l'application :

```go
// actions/context.go
type MyContext struct {
  buffalo.Context
}

func (my MyContext) Error(status int, err error) error {
  my.Logger().Fatal(err)
  return err
}
```

```go
// actions/app.go
// ...
func App() *buffalo.App {
  if app != nil {
    // ...
    app.Use(func (next buffalo.Handler) buffalo.Handler {
      return func(c buffalo.Context) error {
      // change the context to MyContext
      return next(MyContext{c})
      }
    })
    // ...
  }
  return app
}
// ...
```

## Itérer sur les paramètres d'URL

La méthode `buffalo.Context#Params` retourne [`buffalo.ParamValues`](https://godoc.org/github.com/gobuffalo/buffalo#ParamValues) qui est une interface autour d'[`url.Values`](https://golang.org/pkg/net/url/#Values). Vous pouvez caster vers ce type pour itérer sur les valeurs des paramètres d'URL.

```go
import "net/url"

func HomeHandler(c buffalo.Context) error {
  if m, ok := c.Params().(url.Values); ok {
    for k, v := range m {
      fmt.Println(k, v)
    }
  }
  return c.Render(200, r.HTML("index.html"))
}
```

## Que trouve-t-on dans le Context ?

Buffalo remplit le contexte de chaque requête avec tout plein d'informations utiles pour votre application, comme par exemple l'URL de la page courante (`current_route`) ou la `session`. Vous trouverez ci-dessous une liste de ce que Buffalo ajoute au contexte de chaque requête, et auquel vous avez accès dans vos actions ou templates.

| Clef            | Type                                                                                 | Description                                                                                                                                   |
| ---             | ---                                                                                  | ---                                                                                                                                     |
| `app`           | [`*buffalo.App`](https://godoc.org/github.com/gobuffalo/buffalo#App)                | Une référence sur l'application Buffalo.                                                                                         |
| `env`           | `string`                                                                             | L'environnement courant. Exemple : `test`, `development`, `production`                                                   |
| `routes`        | [`buffalo.RouteList`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList)     | Une liste de toutes les routes définies dans l'application.                                                                                  |
| `current_route` | [`buffalo.RouteInfo`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo)     | La route de la page actuelle.                                                                                               |
| `current_path`  | `string`                                                                             | L'URL de la page actuelle. Exemple : `/users/1/edit`                                                                              |
| `*Path`         | [`RouteHelperFunc`](https://godoc.org/github.com/gobuffalo/buffalo#RouteHelperFunc) | Des helpers pour vous aider à construire vos liens vers votre application. Exemple : `editUserPath`. Lancez `buffalo task routes` dans une console pour voir la liste complète des helpers de routes disponibles. |
| `params`        | `map[string]string`                                                                  | Les paramètres de la requête courante.                                                                                                |
| `flash`         | `map[string][]string`                                                                | Une map de messages fournis via `buffalo.Context#Flash`.                                                                                    |
| `session`       | [`*buffalo.Session`](https://godoc.org/github.com/gobuffalo/buffalo#Session)        | La session de l'utilisateur courant.                                                                                                             |
| `request`       | [`*http.Request`](https://godoc.org/net/http#Request)                               | La requête courante.                                                                                                                    |
| `tx`            | [`*pop.Connection`](https://godoc.org/github.com/gobuffalo/pop#Connection)          | Seulement disponible si vous utilisez le middleware `github.com/gobuffalo/buffalo/middleware.PopTransaction` (c'est le cas par défaut).                              |

Consultez le chapitre [Helpers](/fr/docs/helpers#builtin-helpers/docs/helpers#builtin-helpers) pour une liste de helpers disponibles de base dans vos templates.
