---
name: Manejo de Errores
seoDescription: "Manejo de Errores"
seoKeywords: ["buffalo", "go", "golang", "errors"]
weight: 7
aliases:
  - /docs/errors
  - /es/docs/errors
---

# Manejo de Errores

Un `error` es la forma de Go de decir que algo ha ido mal. En este capítulo, aprenderás cómo devolver errores desde un handler de ruta y cómo Buffalo atrapará cualquier error no manejado.

## Retornar errores desde un handler

La forma más sencilla de producir una respuesta de error es devolver un error estándar de Go:
```go
func MyHandler(c buffalo.Context) error {
  // Return any go error, this will result in a 500 status code.
  return errors.New("boom!")
}
```

Un error `nil` producirá una respuesta HTTP 200 básica:
```go
func MyHandler(c buffalo.Context) error {
  // HTTP 200
  return nil
}
```

Si necesitas personalizar el mensaje de error o el código HTTP, utiliza el método {{< doclink href="github.com/gobuffalo/buffalo#DefaultContext.Error" message="Error" >}}:
```go
func MyHandler(c buffalo.Context) error {
  // Use the Error function on the context.
  // This will result in a status code of 401.
  return c.Error(401, errors.New("Unauthorized!"))
}
```

## Manejo de errores por defecto (Development)

En modo "development " (`GO_ENV=development`), Buffalo generará algunas páginas de errores útiles para ti.

<figure>
  <img src="/assets/images/500_example.png" title="screenshot">
  <figcaption>An example of a `500` error in development mode.</figcaption>
</figure>

Si utilizas un tipo de contenido JSON o XML, el error se devuelve en el tipo apropiado:

```json
{
  "error": "could not find test/",
  "trace": "could not find test/\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\nnet/http.StripPrefix.func1\n\t/usr/local/go/src/net/http/server.go:1986\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162\ngithub.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2694\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1830\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_amd64.s:2361",
  "code": 404
}
```

```xml
<response code="404">
  <error>could not find test/</error>
  <trace>could not find test/ github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 net/http.StripPrefix.func1 /usr/local/go/src/net/http/server.go:1986 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162 github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127 net/http.serverHandler.ServeHTTP /usr/local/go/src/net/http/server.go:2694 net/http.(*conn).serve /usr/local/go/src/net/http/server.go:1830 runtime.goexit /usr/local/go/src/runtime/asm_amd64.s:2361</trace>
</response>
```

En el modo " production" (`GO_ENV=production`), Buffalo no generará páginas que tengan información de estilo para desarrolladores, porque esto daría información valiosa a los hackers. En su lugar, las páginas son más sencillas.

## Manejo de personalizado errores

Aunque Buffalo gestiona los errores por ti de forma inmediata, puede ser útil gestionar los errores de forma personalizada. Para ello, Buffalo permite asignar códigos de estado HTTP a handlers específicos. Esto significa que el error puede ser tratado de forma personalizada.

```go
app = buffalo.New(buffalo.Options{
  Env: ENV,
})

// We associate the HTTP 422 status to a specific handler.
// All the other status code will still use the default handler provided by Buffalo.
app.ErrorHandlers[422] = func(status int, err error, c buffalo.Context) error {
  res := c.Response()
  res.WriteHeader(422)
  res.Write([]byte(fmt.Sprintf("Oops!! There was an error: %s", err.Error())))
  return nil
}

app.GET("/oops", MyHandler)

func MyHandler(c buffalo.Context) error {
  return c.Error(422, errors.New("Oh no!"))
}
```

```text
GET /oops -> [422] Oops!! There was an error: Oh no!
```

En el ejemplo anterior, cualquier error de tu aplicación que devuelva un estado de `422` será capturado por el handler personalizado y será tratado en base a ello.

