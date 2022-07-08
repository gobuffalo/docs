---
name: Middleware
weight: 6
aliases:
  - /docs/middleware
  - /es/docs/middleware
---
# Middleware

El middleware permite interponer código en el ciclo de petición/respuesta. Los casos de uso comunes para el middleware son cosas como el registro (que Buffalo ya hace), las peticiones de autenticación, etc.

En [https://toolkit.gobuffalo.io/tools?topic=middleware](https://toolkit.gobuffalo.io/tools?topic=middleware) se puede encontrar una lista de paquetes de middleware "conocidos".

## Escribiendo tu propio Middleware

La interfaz [`buffalo.MiddlewareFunc`](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareFunc) es cualquier función que toma un `buffalo.Handler` y devuelve un `buffalo.Handler`.

```go
func MyMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    // do some work before calling the next handler
    err := next(c)
    // do some work after calling the next handler
    return err
  }
}
```

Implementando la interfaz `buffalo.MiddlewareFunc` podrás controlar el flujo de ejecución de tu aplicación. Piensa en un middleware de autorización; envía errores a tu herramienta de monitorización favorita; carga datos en el `buffalo.Context`, y mucho más.

### Ejemplo

```go
// UserIPMiddleware gets the user IP and sets it to the context.
func UserIPMiddleware(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    if xRealIP := c.Request().Header.Get("X-Real-Ip"); len(xRealIP) > 0 {
      c.Set("user_ip", xRealIP)
      return next(c)
    }

    if xForwardedFor := c.Request().Header.Get("X-Forwarded-For"); len(xForwardedFor) > 0 {
      c.Set("user_ip", xForwardedFor)
      return next(c)
    }

    h, _, err := net.SplitHostPort(c.Request().RemoteAddr)
    if err != nil {
      return err
    }
    c.Set("user_ip", h)
    return next(c)
  }
}
```

## Uso de un Middleware

```go
a := buffalo.New(buffalo.Options{})

a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)
// or
a.Use(MyMiddleware, AnotherPieceOfMiddleware)
```

En el ejemplo anterior todas las peticiones pasarán primero por el middleware `MyMiddleware`, y luego por el middleware `AnotherPieceOfMiddleware` antes de llegar a su handler final.

{{<note>}}
**NOTA**: El middleware definido en una aplicación es heredado automáticamente por todas las rutas y grupos de esa aplicación.
{{</note>}}


## Uso de un Middleware con una Acción

A menudo, hay casos en los que se quiere utilizar un middleware en una sola acción, y no en toda la aplicación o recurso.

Dado que la definición de un middleware es que toma un `buffalo.Handler` y devuelve un `buffalo.Handler` puedes envolver cualquier `buffalo.Handler` en un middlware.

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(MyHandler))
```

Esto no afecta al resto de la pila de middleware que ya está en marcha, sino que se añade a la cadena de middleware sólo para esa acción.

Esto puede llevarse un paso más allá, envolviendo un número ilimitado de middleware alrededor de un `buffalo.Handler`.

```go
a := buffalo.New(buffalo.Options{})
a.GET("/foo", MyMiddleware(AnotherPieceOfMiddleware(MyHandler)))
```


## Agrupar Middlewares

```go
a := buffalo.New(buffalo.Options{})
a.Use(MyMiddleware)
a.Use(AnotherPieceOfMiddleware)

g := a.Group("/api")
// authorize the API end-point
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

a.GET("/foo", FooHandler)
```

En el ejemplo anterior los middlewares `MyMiddleware` y `AnotherPieceOfMiddleware` serán llamados en _todos_ los requests, pero el middleware `AuthorizeAPIMiddleware` sólo será llamado en las rutas `/api/*`.

```text
GET /foo       -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> MyMiddleware -> AnotherPieceOfMiddleware -> AuthorizeAPIMiddleware -> UsersHandler
```

## Omitir un Middleware

Hay ocasiones en las que, en una aplicación, se quiere añadir middleware a toda la aplicación, o a un grupo, pero no llamar a ese middleware en unos pocos handlers individuales. Buffalo permite crear este tipo de mapeos.

{{< codetabs >}}
{{< tab "actions/app.go" >}}
```go
// actions/app.go
a := buffalo.New(buffalo.Options{})
a.Use(AuthorizeUser)

// skip the AuthorizeUser middleware for the NewUser and CreateUser handlers.
a.Middleware.Skip(AuthorizeUser, NewUser, CreateUser)

a.GET("/users/new", NewUser)
a.POST("/users", CreateUser)
a.GET("/users", ListUsers)
a.GET("/users/{id}", ShowUser)
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```text
// OUTPUT
GET /users/new  -> NewUser
POST /users     -> CreateUser
GET /users      -> AuthorizeUser -> ListUsers
GET /users/{id} -> AuthorizeUser -> ShowUser
```
{{< /tab >}}
{{< /codetabs>}}

---

{{<note>}}
**IMPORTANTE:** La función del middleware y las funciones de la acción que quieres omitir **DEBEN** ser la misma instancia de Go.
{{</note>}}

### Ejemplos

{{< codetabs >}}
{{< tab "EJEMPLO 1" >}}
```go
// EJEMPLO 1
m1 := MyMiddleware()
m2 := MyMiddleware()

app.Use(m1)

app.Skip(m2, Foo, Bar) // NO FUNCIONA m2 != m1
app.Skip(m1, Foo, Bar) // FUNCIONA
```
{{< /tab >}}
{{< tab "EJEMPLO 2" >}}
```go
// EJEMPLO 2
app.Resource("/widgets", WidgetResource{})
app.Skip(mw, WidgetResource{}.Show) // NO FUNCIONA

wr := WidgetResource{}
app.Resource("/widgets", wr)
app.Skip(mw, wr.Show) // FUNCIONA
```
{{< /tab >}}
{{< /codetabs>}}

</div>

Véase [https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip](https://godoc.org/github.com/gobuffalo/buffalo#MiddlewareStack.Skip) para más detalles sobre la función `Skip`.

## Omitir acciones de recursos

A menudo es necesario querer omitir el middleware para una o más acciones. Por ejemplo, permitir a los usuarios invitados ver las acciones `List` y `Show` en un recurso, pero exigir autorización en el resto de las acciones.

Entendiendo la sección [Omitir un Middleware](#omitir-un-middleware), tenemos que asegurarnos de que estamos usando las mismas funciones cuando registramos el recurso que cuando queremos saltarnos el middleware en esas funciones más adelante.

La línea que fue generada en `actions/app.go` por `buffalo generate resource` tendrá que ser cambiada para acomodar este requisito.

{{< codetabs >}}
{{< tab "Antes" >}}
```go
app.Resource("/widgets", WidgetResource{})
```
{{< /tab >}}
{{< tab "Después" >}}
```go
res := WidgetResource{}
wr := app.Resource("/widgets", res)
wr.Middleware.Skip(Authorize, res.Index, res.Show)
```
{{< /tab >}}
{{< /codetabs>}}


## Reemplazar un Middleware

Puedes utilizar el método [`Middleware.Replace`](https://pkg.go.dev/github.com/gobuffalo/buffalo#MiddlewareStack.Replace) que permite sustituir un middleware por otro manteniendo la misma posición de ejecución.

{{< codetabs >}}
{{< tab "actions/app.go" >}}
```go
// actions/app.go

app := buffalo.New(buffalo.Options{})
app.Use(Middleware1, Middleware2, Middleware3)

app.GET("/foo/", FooHandler)


g := app.Group("/group")
g.Middleware.Replace(Middleware1, Middleware4)

g.GET("/", GroupListHandler)

```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```text
GET /foo    -> Middleware1 -> Middleware2 -> Middleware3 -> FooHandler
GET /group/ -> Middleware4 -> Middleware2 -> Middleware3 -> GroupListHandler
```
{{< /tab >}}
{{< /codetabs>}}


## Limpiar Middleware

Dado que el middleware es [heredado](#uso-de-un-middleware) de su padre, puede que haya ocasiones en las que sea necesario empezar con un conjunto de middleware "en blanco".

{{< codetabs >}}
{{< tab "actions/app.go" >}}
```go
// actions/app.go
app := buffalo.New(buffalo.Options{})
app.Use(MyMiddleware)
app.Use(AnotherPieceOfMiddleware)

app.GET("/foo", FooHandler)

g := app.Group("/api")
// clear out any previously defined middleware
g.Middleware.Clear()
g.Use(AuthorizeAPIMiddleware)
g.GET("/users", UsersHandler)

```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```text
GET /foo       -> MyMiddleware -> AnotherPieceOfMiddleware -> FooHandler
GET /api/users -> AuthorizeAPIMiddleware -> UsersHandler
```
{{< /tab >}}
{{< /codetabs>}}


## Listando middlewares de una aplicación

Para obtener una lista completa de los middleware que utiliza tu aplicación, desglosada por grupos, se debe ejecutar el comando `buffalo task middleware`.

{{<codetabs>}}
{{<tab "actions/app.go">}}
```go
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})

		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(csrf.New)
		app.Use(translations())
		app.Use(Middleware1)
		app.Use(Middleware2)

		app.GET("/", HomeHandler)

		app.ServeFiles("/", http.FS(public.FS()))
	}

	return app
}
```
{{</tab>}}
{{<tab "Lista de Middlewares">}}
```bash
$ buffalo t middleware
-> http://127.0.0.1:3000/
	github.com/gobuffalo/buffalo.*App.defaultErrorMiddleware
	github.com/gobuffalo/buffalo.*App.PanicHandler
	github.com/gobuffalo/buffalo.RequestLoggerFunc
	github.com/gobuffalo/mw-forcessl.Middleware.func1
	github.com/gobuffalo/mw-paramlogger.ParameterLogger
	github.com/gobuffalo/mw-csrf.glob..func1
	github.com/gobuffalo/mw-i18n/v2.*Translator.Middleware.func1
	coke/actions.Middleware1
	coke/actions.Middleware2
```


{{</tab>}}
{{</codetabs>}}
