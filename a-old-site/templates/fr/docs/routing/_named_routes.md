## Routes nommées

Par défaut, Buffalo donne un nom à chacune de vos routes, sous la forme `pathnamePath`. Par exemple `a.GET("/coke", CokeHandler)` donne le nom de route `cokePath`.

```go
a.GET("/coke", CokeHandler) // cokePath()
```

Ces noms deviennent ceux des *helpers* utilisables dans vos templates.

```html
<a href="\<%= cokePath() %>">Coke</a>
```

Vous pouvez obtenir la liste de toutes vos routes en exécutant la commande `buffalo routes`.

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

**IMPORTANT** : Puisque les noms des *helpers* de routes sont générés en utilisant le chemin de l'URL, (`/widgets/new` -> `newWidgetsPath`), si ce chemin change, le nom du *helper* de route **change aussi**.

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

Consultez [`Routes nommées personnalisées`](#routes-nommées-personnalisées) pour plus de détails sur comment définir les noms de routes vous-même.
