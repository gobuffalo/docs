---
Name: "Enrutamiento"
seoDescription: "How to handle routes in Buffalo?"
seoKeywords: ["buffalo", "go", "golang", "http", "route", "gorilla", "mux", "router"]
weight: 1
aliases:
  - /docs/routing
  - /es/docs/routing
---

# Enrutamiento

Buffalo utiliza el paquete [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) para manejar el enrutamiento dentro de las aplicaciones de Buffalo. Dicho esto, Buffalo envuelve la API `mux` con la suya propia. Esta guía le ayudará a conocer todo lo que necesita saber sobre cómo Buffalo maneja el enrutamiento.

## Necesitamos tener la configuración de buffalo.App creada

La configuración de la aplicación se encuentra en el archivo `actions/app.go`.

```go
// actions/app.go
app = buffalo.New(buffalo.Options{
    Env:         ENV,
    SessionName: "_coke_session",
  })
```

Por defecto, Buffalo sólo requiere 2 parámetros para la configuración de su aplicación:

- `Env`: El entorno donde se ejecutará la aplicación. Valor por defecto: `development`.
- `SessionName`: Es la cookie de sesión que se establece. Valor por defecto: `_buffalo_session`.

Puedes personalizarlo para adaptarlo a tu caso de uso.

Puede consultar la lista de opciones disponibles aquí: [https://godoc.org/github.com/gobuffalo/buffalo#Options](https://godoc.org/github.com/gobuffalo/buffalo#Options)


## Buffalo.Handler

Si ya conoces el patrón **MVC**, las funciones de `buffalo.Handler` gestionan la parte del Controlador. Su estructura es la siguiente:

```go
func (c buffalo.Context) error {
  // do some work
}
```

Aquí es donde va toda la lógica de la aplicación. El handler toma un parámetro `buffalo.Context`, que contiene todo lo que necesitas sobre la petición actual.


{{<note>}}
Consulte [Contexto](/es/documentation/request_handling/context) para entender la interfaz `buffalo.Context`.
{{</note>}}




## Mapeando Handlers

Para mapear un `buffalo.Handler`, tendrás que asociarlo a una ruta específica con un método HTTP.

##### Métodos HTTP soportados

Buffalo soporta los siguientes métodos HTTP:

{{< codetabs >}}
{{< tab "GET" >}}
```go
app.GET("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "POST" >}}
```go
app.POST("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "PUT" >}}
```go
app.PUT("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "PATCH" >}}
```go
app.PATCH("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "DELETE" >}}
```go
app.DELETE("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "OPTIONS" >}}
```go
app.OPTIONS("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< tab "HEAD" >}}
```go
app.HEAD("/your/path", buffalo.Handler)
```
{{< /tab >}}
{{< /codetabs >}}

También puede hacer coincidir todos los métodos HTTP utilizando `ANY`.

Por defecto, Buffalo establece una ruta raíz dentro de la configuración de bufalo.App:

```go
// actions/app.go
func App() *buffalo.App {
  // ...
  app.GET("/", HomeHandler)
  // ...
}
```

La asociación de múltiples `buffalo.Handlers` a los métodos HTTP se ve de la siguiente manera:

```go
// actions/app.go
app.GET("/", HomeHandler)
app.GET("/some/path", SomeHandler)
app.POST("/another/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

Como puedes ver, puedes declarar buffalo.Handlers en la misma línea si quieres.
Sin embargo, para una mayor legibilidad, a menudo es mejor separar los handlers en varios archivos. Por ejemplo, si tienes muchos handlers gestionando cosas de usuarios, puedes agruparlos en un archivo `users.go` en la carpeta [`actions`](/es/documentation/getting_started/directory-structure#actions), por ejemplo.

## Denominanción de Rutas

Por defecto, Buffalo nombrará sus rutas por usted en la forma de `<nombre de la ruta>Path`.

Por ejemplo: `a.GET("/coke", CokeHandler)` dará como resultado una ruta llamada `cokePath`.

```go
a.GET("/coke", CokeHandler) // cokePath()
```

Estos nombres se convierten en el nombre de los helpers de ruta en tus plantillas.

```html
<a href="<%= cokePath() %>">Coke</a>
```

## Denominación de Rutas Personalizadas

Buffalo también le proporciona una forma de establecer un nombre personalizado para su ruta, El método [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) le permite establecer un nombre personalizado para los helpers de la ruta.
Para personalizar el nombre de su ruta, sólo tiene que utilizar el método Name después de asignar el método HTTP.

```go
app.GET("/coke", CokeHandler).Name("customCoke") // customCokePath()
```

Esta ruta se llama ahora `customCokePath` y puedes referenciarla como tal en tus plantillas.


```html
<a href="<%= customCokePath() %>">Coke</a>
```

## Lista de Rutas

Puedes revisar todas tus rutas ejecutando `buffalo routes` desde la línea de comandos.

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                       |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/             |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/          |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                  |         | customCokePath       | coke/actions.CokeHandler
```


{{<note>}}
**IMPORTANTE:** Dado que los nombres de los helpers de las rutas se calculan usando el **`path`** pe. **`/widgets/new -> newWidgetsPath`**; si la ruta cambia, entonces el nombre del helper de la ruta **también** cambia.
{{</note>}}

Ejemplo:

Mapeando `WidgetResource` en la ruta `/widgets`:

```go
app.Resource("/widgets", WidgetsResource{})
```

Obtendrás los siguientes nombres de rutas:

```bash
$ buffalo routes

METHOD | HOST                  | PATH                       | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                       | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                          |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/                |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/             |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                     |         | customCokePath       | coke/actions.CokeHandler

GET    | http://127.0.0.1:3000 | /widgets/                  |         | widgetsPath    | coke/actions.WidgetResource.List
POST   | http://127.0.0.1:3000 | /widgets/                  |         | widgetsPath    | coke/actions.WidgetResource.Create
GET    | http://127.0.0.1:3000 | /widgets/new/              |         | newWidgetsPath | coke/actions.WidgetResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/      |         | widgetPath     | coke/actions.WidgetResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/edit/ |         | editWidgetPath | coke/actions.WidgetResource.Edit
```

Pero, si cambias el nombre de la ruta a `/fooz`:

```go
app.Resource("/fooz", WidgetsResource{})
```

Los nombres de las rutas pasarán a llamarse:

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /                       |         | rootPath             | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /some/path/             |         | somePath             | coke/actions.SomeHandler
POST   | http://127.0.0.1:3000 | /another/path/          |         | anotherPath          | coke/actions.App.func1
GET    | http://127.0.0.1:3000 | /coke/                  |         | customCokePath       | coke/actions.CokeHandler

GET    | http://127.0.0.1:3000 | /fooz/                  |         | foozPath             | coke/actions.WidgetResource.List
POST   | http://127.0.0.1:3000 | /fooz/                  |         | foozPath             | coke/actions.WidgetResource.Create
GET    | http://127.0.0.1:3000 | /fooz/new/              |         | newFoozPath          | coke/actions.WidgetResource.New
GET    | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Show
PUT    | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Update
DELETE | http://127.0.0.1:3000 | /fooz/{widget_id}/      |         | foozWidgetIDPath     | coke/actions.WidgetResource.Destroy
GET    | http://127.0.0.1:3000 | /fooz/{widget_id}/edit/ |         | editFoozWidgetIDPath | coke/actions.WidgetResource.Edit
```

Vea [`Denominación de Rutas Personalizadas`](#denominación-de-rutas-personalizadas) para más detalles sobre cómo cambiar el nombre generado.


## Uso de los Helpers de ruta en las plantillas

Los helpers de ruta se pueden utilizar directamente en las plantillas utilizando el nombre del helper:

```erb
<%= widgetsPath() %> // /widgets
```

Las rutas que requieren parámetros con nombre, deben ser alimentadas con un mapa de esos parámetros.

```erb
<%= editWidgetPath({widget_id: 1}) %> --> /widgets/1/edit
```


## El Helper `pathFor`

El helper `pathFor` recibe una `interface{}` o un `slice` de estas, y trata de convertirlo en una ruta URL de estilo `/foos/{id}`

Reglas:
* si es `string` se devuelve tal cual
* si es `Pathable`, se devuelve el método `ToPath`.
* si se trata de un `slice` o de un `array`, cada elemento se pasa por el helper y luego se unen
* si es `struct`, el nombre de la estructura se utiliza el plural para el nombre
* si es `Paramable` el método `ToParam` se utiliza para rellenar el campo `{id}`
* si es `struct.Slug` el slug se utiliza para rellenar el campo `{id}` de la URL
* si es `struct.ID` el ID se utiliza para rellenar el campo `{id}` de la URL

```go
// Car{1} => "/cars/1"
// Car{} => "/cars"
// &Car{} => "/cars"
type Car struct {
  ID int
}

// Boat{"titanic"} => "/boats/titanic"
type Boat struct {
  Slug string
}

// Plane{} => "/planes/aeroPlane"
type Plane struct{}

func (Plane) ToParam() string {
  return "aeroPlane"
}

// Truck{} => "/a/Truck"
// {[]interface{}{Truck{}, Plane{}} => "/a/Truck/planes/aeroPlane"
type Truck struct{}

func (Truck) ToPath() string {
  return "/a/Truck"
}
```


## Uso de los Helpers de ruta en las acciones

### Redireccionamiento con helpers de ruta

También se pueden utilizar nombres de ruta cuando se redirige a otra url.

```go
func MyHandler(c buffalo.Context) error {
  return c.Redirect(http.StatusSeeOther, "widgetsPath()")
  // Or with parameters
  return c.Redirect(http.StatusSeeOther, "widgetPath()", render.Data{"widget_id": "1"})
}
```

---

### Encontrar/llamar a un Helper de ruta

{{< since "0.13.0-beta.1" >}}

La función [`buffalo.RouteList#Lookup`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList.Lookup) permite buscar una ruta por su nombre desde la aplicación. Con el valor `RouteInfo` de la ruta dada se puede generar la ruta para la misma.

```go
func MyHandler(c buffalo.Context) error {
  ri, err := App().Routes().Lookup("widgetPath")
  if err != nil {
    return errors.WithStack(err)
  }
  h := ri.BuildPathHelper()
  u, err := h(render.Data{"widget_id": 1})
  if err != nil {
    return errors.WithStack(err)
  }
  return c.Redirect(307, string(u))
}
```


## Parámetros

Query string y otros parámetros están disponibles en el [`buffalo.Context`](/documentation/request_handling/context) que se pasa al `buffalo.Handler`.

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users?name=ringo`, the response should be `200: ringo`.

## Parámetros con nombre

Dado que Buffalo es el [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) bajo las cubiertas, significa que podemos acceder a algunas de las bondades que proporciona. En este caso, la capacidad de crear patrones de expresiones pseudo-regulares en la ruta mapeada que se convertirán en parámetros a los que se puede acceder desde un [`buffalo.Context`](/documentation/request_handling/context).

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Dado el ejemplo de código anterior, si hacemos una petición con `GET /users?name=ringo`, la respuesta debería ser `200: ringo`.

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, r.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Puedes asignar rutas aparentemente similares, como `/users/new` y `/users/{name}` sin ningún problema. El router se asegurará de que lleguen al mismo lugar.

### Expresiones regulares

[github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) proporciona una forma de utilizar expresiones regulares, para que puedas pre-filtrar las queries:

```go
a.GET("/articles/{id:[0-9]+}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("id")))
})
```


## Grupos

Las aplicaciones de Buffalo permiten agrupar los end-points. Esto permite reunir funcionalidades comunes, como los [middleware](/es/documentation/request_handling/middleware). Un gran ejemplo de esto sería un end-point de la API.

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // responds to GET /api/v1/users
})
```

Por defecto, un grupo heredará cualquier middleware de su aplicación padre.

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

En el ejemplo anterior el grupo `/api/v1` utilizará tanto `SomeMiddleware` como `APIAuthorizer`. Ver [middleware](/documentation/request_handling/middleware) para más información sobre el uso, la omisión y la eliminación de middleware.

## Hosts Virtuales

{{< since "0.18.2" >}}

Las aplicaciones de Buffalo también admiten la agrupación de end-points por host. La función `VirtualHost` crea un nuevo grupo que coincide con el dominio introducido. Esto es útil para crear grupos de end-points para diferentes dominios o subdominios.

```go
app := buffalo.New(buffalo.Options{
    Env:         envy.Get("GO_ENV", "development"),
    SessionName: "_coke_session",
})

subApp := app.VirtualHost("docs.domain.com")
subApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("docs.domain.com Homepage"))
})

domainApp := app.VirtualHost("example.com")
domainApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("example.com Homepage"))
})

app.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("Main App Homepage"))
})
```

También se admiten las variables asignadas a los parámetros:

```go
app.VirtualHost("{subdomain}.example.com")
app.VirtualHost("{subdomain:[a-z]+}.example.com")
```

## Montaje de aplicaciones http.Handler

{{< since "0.9.4" >}}

A veces, querrás reutilizar algunos componentes de otras aplicaciones. Usando el método [`Mount`](https://godoc.org/github.com/gobuffalo/buffalo#App.Mount), puedes vincular un [`http.Handler`](https://golang.org/pkg/net/http/#Handler) estándar a una ruta, igual que harías con un handler de ruta normal.

```go
func muxer() http.Handler {
  f := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "%s - %s", req.Method, req.URL.String())
  }
  mux := mux.NewRouter()
  mux.HandleFunc("/foo/", f).Methods("GET")
  mux.HandleFunc("/bar/", f).Methods("POST")
  mux.HandleFunc("/baz/baz/", f).Methods("DELETE")
  return mux
}

a.Mount("/admin", muxer())
```

Dado que Buffalo `App` implementa la interfaz `http.Handler`, también puede montar otra aplicación de Buffalo y construir aplicaciones modulares.
