---
Name: Contexto
weight: 4
aliases:
  - /docs/context
  - /es/docs/context
---
# Contexto

En el corazón de cada gestor de peticiones de Buffalo se encuentra el `contexto`. Este contexto proporciona a los controladores una definición de función simple y limpia, a la vez que es inmensamente potente.

## La interfaz context

La interfaz `buffalo.Context` soporta `context.Context` por lo que se puede pasar y utilizar como un Contexto de Go "estándar".

Dado que `buffalo.Context` es una interfaz, es posible crear una implementación específica para la aplicación que se adapte a las necesidades de la aplicación que se está construyendo.

{{< since "0.12.0" >}}

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

La función `Websocket() (*websocket.Conn, error)` fue eliminada de `buffalo.Context` en la versión `v0.12.0`. Utiliza el paquete [http://www.gorillatoolkit.org/pkg/websocket](http://www.gorillatoolkit.org/pkg/websocket) en su lugar.

## Contexto y renderizado

Como parte de la interfaz del contexto, hay una función `Render` que toma un tipo de `render.Renderer`. Revisa [rendering](/documentation/frontend-layer/rendering) para mayor información.

Cualquier valor que sea definido en el contexto estará automáticamente disponible para el `render.Renderer` que se pasa a la función `Render`.

```go
func Hello(c buffalo.Context) error {
  c.Set("name", "Mark")

  return c.Render(http.StatusOK, render.String("Hi <%= name %>"))
}
```

## Implementando la Interfaz

El `buffalo.Context` no está pensado para ser implementado "totalmente". En su lugar, se recomienda utilizar una [composición](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) y utilizar sólo las funciones de las que se desea proporcionar implementaciones personalizadas.

A continuación se muestra un ejemplo de cambio de la función `Error` para registrar el error y matar la aplicación:

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

## Recorriendo los Parámetros

El método `buffalo.Context#Params` devuelve [`buffalo.ParamValues`](https://godoc.org/github.com/gobuffalo/buffalo#ParamValues) que es una interfaz alrededor de url.Values. Se puede hacer un casting a este tipo en el handler para recorrer los valores de los parámetros.

```go
import "net/url"

func HomeHandler(c buffalo.Context) error {
  if m, ok := c.Params().(url.Values); ok {
    for k, v := range m {
      fmt.Println(k, v)
    }
  }

  return c.Render(http.StatusOK, r.HTML("index.html"))
}
```

## Qué hay en el contexto

Buffalo añade al contexto de cada solicitud un montón de información que podría ser útil en tu aplicación, como la `current_route` o la `session`. A continuación se muestra una lista de lo que Buffalo añade al contexto de cada solicitud al que puedes acceder desde las acciones o plantillas.

| Llave             | Tipo                                                                                 | Uso                                                                                                                                   |
| ---             | ---                                                                                  | ---                                                                                                                                     |
| `app`           | [`*buffalo.App`](https://godoc.org/github.com/gobuffalo/buffalo#App)                | La aplicación actual de Buffalo que se está ejecutando.                                                                                         |
| `env`           | `string`                                                                             | El entorno actual en el que se ejecuta la aplicación. Por ejemplo: `test`, `development`, `production`                                                   |
| `routes`        | [`buffalo.RouteList`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList)     | Una lista de todas las rutas mapeadas en la aplicación.                                                                                  |
| `current_route` | [`buffalo.RouteInfo`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo)     | La ruta actual a la que se accede.                                                                                                 |
| `current_path`  | `string`                                                                            | La ruta actual solicitada. Ejemplo: `/users/1/edit`                                                                                                   |
| `*Path`         | [`RouteHelperFunc`](https://godoc.org/github.com/gobuffalo/buffalo#RouteHelperFunc) | Helpers para crear caminos basados en las rutas mapeadas. Por ejemplo: `editUserPath`. Ejecuta `buffalo task routes` para ver una lista completa para tu aplicación. |
| `params`        | `map[string]string`                                                                  | Parámetros de consulta de la página solicitada.                                                                                              |
| `flash`         | `map[string][]string`                                                                | Un mapa de mensajes establecidos mediante `buffalo.Context#Flash`.                                                                                          |
| `session`       | [`*buffalo.Session`](https://godoc.org/github.com/gobuffalo/buffalo#Session)         | La sesión del usuario actual.                                                                                                  |
| `request`       | [`*http.Request`](https://godoc.org/net/http#Request)                                | La petición actual.                                                                                                               |
| `tx`            | [`*pop.Connection`](https://godoc.org/github.com/gobuffalo/pop#Connection)          | Sólo se establece si se utiliza el middleware `github.com/gobuffalo/buffalo/middleware.PopTransaction` (activado por defecto).                              |

Consulta [Helpers](/documentation/frontend-layer/helpers#builtin-helpers) para obtener una lista de las funciones de ayuda integradas disponibles dentro de las plantillas.
