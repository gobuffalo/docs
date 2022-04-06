---
Name: "Routing"
seoDescription: "How to handle routes in Buffalo?"
seoKeywords: ["buffalo", "go", "golang", "http", "route", "gorilla", "mux", "router"]
weight: 1
aliases:
  - /docs/routing
  - /en/docs/routing
---

# Routing

Buffalo uses the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers, to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with its own. This guide walks you through all you'll need to know about how Buffalo handles routing.

## Creating a new Buffalo App (and router)

The app configuration is located in the `app.go` file.

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
})
```

The default setup should handle most of your needs, but you are free to customize it to fit your use case.

You can check the available options list here: [https://godoc.org/github.com/gobuffalo/buffalo#Options](https://godoc.org/github.com/gobuffalo/buffalo#Options)

## Mapping Handlers

All routing in Buffalo results in the calling of a `buffalo.Handler` function. The signature for a `buffalo.Handler` looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

If you already know about **MVC pattern**, `buffalo.Handler` functions manages the Controller part: this is the place where all the app logic goes. The handler function takes a `buffalo.Context` struct, which contains everything you need about the current request.

See the [Context](/documentation/request_handling/context) to understand the `buffalo.Context` interface.

##### Supported HTTP Methods

Buffalo supports the following HTTP methods out of the box:

* GET
* POST
* PUT
* PATCH
* DELETE
* OPTIONS
* HEAD

You can also match all HTTP methods using `ANY`.

Mapping a `buffalo.Handler` to an HTTP method takes the form of:

```go
a.GET("/some/path", SomeHandler)
a.POST("/some/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

As you can see, you can use inline handlers if you want. For more readability though, it's often better to separate your handlers into multiple files. If you have many handlers managing users stuff, you can group them into a `users.go` file in the [`actions`](/documentation/getting_started/directory-structure#actions) folder, for instance.


## Named Routes

By default, Buffalo will name your routes for you in the form of `pathnamePath`. For example `a.GET("/coke", CokeHandler)` will result in a route named `cokePath`.

```go
a.GET("/coke", CokeHandler) // cokePath()
```

These names become the name of the route helpers in your templates.

```html
<a href="<%= cokePath() %>">Coke</a>
```

You can inspect all of your paths by running `buffalo routes` from the command line.

```plain
$ buffalo routes

METHOD | PATH                       | ALIASES | NAME           | HANDLER
------ | ----                       | ------- | ----           | -------
GET    | /                          |         | rootPath       | github.com/markbates/coke/actions.HomeHandler
GET    | /widgets/                  |         | widgetsPath    | github.com/markbates/coke/actions.WidgetsResource.List
POST   | /widgets/                  |         | widgetsPath    | github.com/markbates/coke/actions.WidgetsResource.Create
GET    | /widgets/new/              |         | newWidgetsPath | github.com/markbates/coke/actions.WidgetsResource.New
GET    | /widgets/{widget_id}/      |         | widgetPath     | github.com/markbates/coke/actions.WidgetsResource.Show
PUT    | /widgets/{widget_id}/      |         | widgetPath     | github.com/markbates/coke/actions.WidgetsResource.Update
DELETE | /widgets/{widget_id}/      |         | widgetPath     | github.com/markbates/coke/actions.WidgetsResource.Destroy
GET    | /widgets/{widget_id}/edit/ |         | editWidgetPath | github.com/markbates/coke/actions.WidgetsResource.Edit
```

---

**IMPORTANT:** Because route helper names are calculated using the `path`, (`/widgets/new` -> `newWidgetsPath`), if the path changes, then the route helper name **also** changes.

```go
app.Resource("/fooz", WidgetsResource{})
```

```bash
$ buffalo routes

METHOD | PATH                    | ALIASES | NAME         | HANDLER
------ | ----                    | ------- | ----         | -------
GET    | /                       |         | rootPath     | github.com/markbates/coke/actions.HomeHandler
GET    | /fooz/                  |         | foozPath     | github.com/markbates/coke/actions.WidgetsResource.List
POST   | /fooz/                  |         | foozPath     | github.com/markbates/coke/actions.WidgetsResource.Create
GET    | /fooz/new/              |         | newFoozPath  | github.com/markbates/coke/actions.WidgetsResource.New
GET    | /fooz/{widget_id}/      |         | foozPath     | github.com/markbates/coke/actions.WidgetsResource.Show
PUT    | /fooz/{widget_id}/      |         | foozPath     | github.com/markbates/coke/actions.WidgetsResource.Update
DELETE | /fooz/{widget_id}/      |         | foozPath     | github.com/markbates/coke/actions.WidgetsResource.Destroy
GET    | /fooz/{widget_id}/edit/ |         | editFoozPath | github.com/markbates/coke/actions.WidgetsResource.Edit
```

See [`Custom Named Routes`](#custom-named-routes) for details on how to change the generated name.

## Using Route Helpers in Templates

Route helpers can be used directly in templates using the name of the helper:

```erb
<%= widgetsPath() %> // /widgets
```

Routes that require named parameters, must be fed a map of those parameters.

```erb
<%= editWidgetPath({widget_id: 1}) %> // /widgets/1/edit
```


## The `pathFor` Helper

The `pathFor` helper takes an `interface{}`, or a `slice` of them,
and tries to convert it to a `/foos/{id}` style URL path.

Rules:
* if `string` it is returned as is
* if `Pathable` the `ToPath` method is returned
* if `slice` or an `array` each element is run through the helper then joined
* if `struct` the name of the struct, pluralized is used for the name
* if `Paramable` the `ToParam` method is used to fill the `{id}` slot
* if `struct.Slug` the slug is used to fill the `{id}` slot of the URL
* if `struct.ID` the ID is used to fill the `{id}` slot of the URL

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


## Using Route Helpers in Actions

### Redirecting with Route Helpers

You can also use route names when redirecting to another url.

```go
func MyHandler(c buffalo.Context) error {
  return c.Redirect(307, "widgetsPath()")
  // Or with parameters
  return c.Redirect(307, "widgetPath()", render.Data{"widget_id": "1"})
}
```

---

### Finding/Calling a Route Helper

{{< since "0.13.0-beta.1" >}}

The [`buffalo.RouteList#Lookup`](https://godoc.org/github.com/gobuffalo/buffalo#RouteList.Lookup) allows you to look up a route by its name from the application. With the `RouteInfo` value for the given route you can generate the path for the route.

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


## Custom Named Routes

The [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) function allows you to set a custom name for route helpers.

```go
a.GET("/coke", CokeHandler).Name("customPath")
```

This route is now called `customPath` and you can reference it as such in your templates.

```erb
<a href="<%= customPath() %>">Coke</a>
```

## Parameters

Query string and other parameters are available from the [`buffalo.Context`](/documentation/request_handling/context) that is passed into the `buffalo.Handler`.

```go
a.GET("/users", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users?name=ringo`, the response should be `200: ringo`.

## Named Parameters

Since Buffalo is the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) under the covers, it means we can get access to some of the goodness it provides. In this case, the ability to create pseudo-regular expression patterns in the mapped path that will get converted into parameters that can be accessed from a [`buffalo.Context`](/documentation/request_handling/context).

```go
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

Given the above code sample, if we make a request with `GET /users/ringo`, the response should be `200: ringo`.

```go
a.GET("/users/new", func (c buffalo.Context) error {
  return c.Render(200, r.String("new"))
})
a.GET("/users/{name}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("name")))
})
```

You may map seemingly similar paths, like `/users/new` and `/users/{name}` without any issues. The router will make sure they get to the same place.

### Regular expressions

[github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) provides a way to use regular expressions, so you can pre-filter queries:

```go
a.GET("/articles/{id:[0-9]+}", func (c buffalo.Context) error {
  return c.Render(200, r.String(c.Param("id")))
})
```


## Groups

Buffalo apps allow for the grouping of end-points. This allows for common functionality, such as [middleware](/documentation/request_handling/middleware) to be collected together. A great example of this would be an API end-point.

```go
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
g.GET("/users", func (c buffalo.Context) error {
  // responds to GET /api/v1/users
})
```

By default a group will inherit any middleware from its parent app.

```go
a.Use(SomeMiddleware)
g := a.Group("/api/v1")
g.Use(APIAuthorizer)
```

In the above example the `/api/v1` group will use both `SomeMiddleware` and `APIAuthorizer`. See [middleware](/documentation/request_handling/middleware) for more information about using, skipping, and clearing middleware.


## Hosts

{{< since "0.18.2" >}}

Buffalo apps also support grouping of end-points by host. `Host` creates a new group that matches the domain passed. This is useful for creating groups of end-points for different domains or subdomains.

```go
app := buffalo.New(buffalo.Options{
    Env:         envy.Get("GO_ENV", "development"),
    SessionName: "_coke_session",
})

subApp := app.Host("docs.domain.com")
subApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("docs.domain.com Homepage"))
})

domainApp := app.Host("example.com")
domainApp.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("example.com Homepage"))
})

app.GET("/", func (c buffalo.Context) error {
  return c.Render(http.StatusOK, r.String("Main App Homepage"))
})
```

Variables mapped to parameters are also supported:

```go
app.Host("{subdomain}.example.com")
app.Host("{subdomain:[a-z]+}.example.com")
```

## Mounting http.Handler Apps

{{< since "0.9.4" >}}

Sometimes, you'll want to reuse some components from other apps. Using the [`Mount`](https://godoc.org/github.com/gobuffalo/buffalo#App.Mount) method, you can bind a standard [`http.Handler`](https://golang.org/pkg/net/http/#Handler) to a route, just like you'll do with a normal route handler.

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

Since Buffalo `App` implements the `http.Handler` interface, you can also mount another Buffalo app and build modular apps.

