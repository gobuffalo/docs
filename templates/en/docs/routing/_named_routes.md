## Named Routes

By default, Buffalo will name your routes for you in the form of `pathnamePath`. For example `a.GET("/coke", CokeHandler)` will result in a route named `cokePath`.

```go
a.GET("/coke", CokeHandler) // cokePath()
```

These names become the name of the route helpers in your templates.

```html
&lt;a href="\<%= cokePath() %>">Coke&lt;/a>
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
