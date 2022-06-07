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

## We need to have the buffalo.App configuration created

The app configuration is located in the `actions/app.go` file.

```go
// actions/app.go
app = buffalo.New(buffalo.Options{
    Env:         ENV,
    SessionName: "_coke_session",
  })
```

By default, buffalo requires only 2 options for its app setup:

- `Env`: The enviroment where the application will run. Default value: `development`.
- `SessionName`: Is the session cookie that is set. Default value: `_buffalo_session`.

You are free to customize it to fit your use case.

You can check the available options list here: [https://godoc.org/github.com/gobuffalo/buffalo#Options](https://godoc.org/github.com/gobuffalo/buffalo#Options)


## Buffalo.Handler

If you already know about **MVC pattern**, `buffalo.Handler` functions manages the Controller part. Its signature looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

This is the place where all the app logic goes. The handler function takes a `buffalo.Context` param, which contains everything you need about the current request.


{{<note>}}
See the [Context](/documentation/request_handling/context) to understand the `buffalo.Context` interface.
{{</note>}}




## Mapping Handlers

To map a `buffalo.Handler`, you'll need to associate it with an specific path with an HTTP method.

##### Supported HTTP Methods

Buffalo supports the following HTTP methods:

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

You can also match all HTTP methods using `ANY`.

As default, Buffalo sets a root path inside bufalo.App setup:

```go
// actions/app.go
func App() *buffalo.App {
  // ...
  app.GET("/", HomeHandler)
  // ...
}
```

Mapping multiple `buffalo.Handlers` to HTTP methods take the form of:

```go
// actions/app.go
app.GET("/", HomeHandler)
app.GET("/some/path", SomeHandler)
app.POST("/another/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

As you can see, you can use inline buffalo.Handlers if you want.
For more readability though, it's often better to separate your handlers into multiple files. For example, if you have many handlers managing users stuff, you can group them into a `users.go` file in the [`actions`](/documentation/getting_started/directory-structure#actions) folder, for instance.

## Named Routes

By default, Buffalo will name your routes for you in the form of `<pathName>Path`.

For example: `a.GET("/coke", CokeHandler)` will result in a route named `cokePath`.

```go
a.GET("/coke", CokeHandler) // cokePath()
```

These names become the name of the route helpers in your templates.

```html
<a href="<%= cokePath() %>">Coke</a>
```

## Custom Named Routes

Buffalo also provides you a way to set a custom name for your route, The [`buffalo.RouteInfo#Name`](https://godoc.org/github.com/gobuffalo/buffalo#RouteInfo.Name) method allows you to set a custom name for route helpers.
To customize your route name, just use the Name method after mapping the HTTP Method.

```go
app.GET("/coke", CokeHandler).Name("customCoke") // customCokePath()
```

This route is now called `customCokePath` and you can reference it as such in your templates.


```html
<a href="<%= customCokePath() %>">Coke</a>
```

## Route list

You can inspect all of your paths by running `buffalo routes` from the command line.

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
**IMPORTANT:** Because route helper names are calculated using the **`path`** pe. **`/widgets/new -> newWidgetsPath`**; if path changes, then the route helper name **also** changes.
{{</note>}}

Example:

Mapping `WidgetResource` in `/widgets` path:

```go
app.Resource("/widgets", WidgetsResource{})
```

You will get the following route path names:

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

But, if you rename the route path to `/fooz`:

```go
app.Resource("/fooz", WidgetsResource{})
```

The route names will be renamed to:

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

See [`Custom Named Routes`](#custom-named-routes) for details on how to change the generated name.


## Using Route Helpers in Templates

Route helpers can be used directly in templates using the name of the helper:

```erb
<%= widgetsPath() %> // /widgets
```

Routes that require named parameters, must be fed a map of those parameters.

```erb
<%= editWidgetPath({widget_id: 1}) %> --> /widgets/1/edit
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
  return c.Redirect(http.StatusSeeOther, "widgetsPath()")
  // Or with parameters
  return c.Redirect(http.StatusSeeOther, "widgetPath()", render.Data{"widget_id": "1"})
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


## Virtual Hosts

{{< since "0.18.2" >}}

Buffalo apps also support grouping of end-points by host. `VirtualHost` creates a new group that matches the domain passed. This is useful for creating groups of end-points for different domains or subdomains.

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

Variables mapped to parameters are also supported:

```go
app.VirtualHost("{subdomain}.example.com")
app.VirtualHost("{subdomain:[a-z]+}.example.com")
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
