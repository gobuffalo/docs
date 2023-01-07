---
Name: "Helpers"
weight: 6
aliases:
  - /docs/helpers
  - /es/docs/helpers
---

# Helpers

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

## Helpers incorporados

Una lista completa de todas las funciones de helpers para [`github.com/gobuffalo/plush`](https://godoc.org/github.com/gobuffalo/plush) se puede encontrar en [`github.com/gobuffalo/helpers`](https://godoc.org/github.com/gobuffalo/helpers).

## Helpers de rutas

Buffalo generará los helpers de rutas para todas las rutas que agregas a la aplicación. La forma mas fácil de ver cuáles son todos los helpers de ruta generados y a qué apuntan, es ejecutar `buffalo routes`. Esto imprimirá una lista como esta:

```text
$ buffalo routes
METHOD | HOST                   | PATH                         | ALIASES | NAME              | HANDLER
------ | ----                   | ----                         | ------- | ----              | -------
GET    | http://127.0.0.1:3000  | /                            |         | rootPath          | github.com/gobuffalo/coke/actions.HomeHandler
GET    | http://127.0.0.1:3000  | /about                       |         | aboutPath         | github.com/gobuffalo/coke/actions.AboutHandler
GET    | http://127.0.0.1:3000  | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.List
POST   | http://127.0.0.1:3000  | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.Create
GET    | http://127.0.0.1:3000  | /drinks/new                  |         | newDrinksPath     | github.com/gobuffalo/coke/actions.DrinksResource.New
GET    | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Show
PUT    | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Update
DELETE | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Destroy
GET    | http://127.0.0.1:3000  | /drinks/{drink_id}/edit      |         | editDrinkPath     | github.com/gobuffalo/coke/actions.DrinksResource.Edit
GET    | http://127.0.0.1:3000  | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.List
POST   | http://127.0.0.1:3000  | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.Create
GET    | http://127.0.0.1:3000  | /api/v1/users/new            |         | newApiV1UsersPath | github.com/gobuffalo/coke/actions.UsersResource.New
GET    | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Show
PUT    | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Update
DELETE | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Destroy
GET    | http://127.0.0.1:3000  | /api/v1/users/{user_id}/edit |         | editApiV1UserPath | github.com/gobuffalo/coke/actions.UsersResource.Edit
```

Bajando por esta lista podemos ver la ruta en la columna *NAME* `rootPath`, la cual se representa en la columna *PATH* `/` o la ruta raíz del servidor y como bonus, con todos estos podemos incliso ver exactamente que controlador en la columna *HANDLER* se está ejecutando para esta combinación *METHOD* + *PATH*.

A continuación tenemos un estándar `app.GET("/about", AboutHandler)` que genera `aboutPath`.

Entonces usamos un recurso `app.Resource("/drinks", DrinksResource{})`, el cual genera una ruta para cada una de nuestras acciones estándar, y para cada una de ellas, un helper para usarlo en las plantillas. Los que toman un parametro se pueden usar asi `<%= drinkPath({drink_id: drink.ID}) %>`. Todos los helpers toman un `map[string]interface{}` que se usa para rellenar los parámetros.

Finalmente, cuando usamos un grupo podemos ver que esto cambia los helpers generados. Aquí está el enrutamiento para esas últimas rutas:

```go
api := app.Group("/api/v1")
api.Resource("/users", UsersResource{})
```

**Nota** que los helpers se generan para que coincidan con las rutas generadas. Es posible sobreescribir los nombres de las rutas en `App.Routes`, pero se conseja encarecidamente que encuentres un camino diferente a tu objetivo. Slack siempre está abierto a estas conversaciones.

### Helper PathFor

El helper [`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) toma una `interface{}`, or a `slice` de ellas, e intenta convertirlo en una ruta de URL al estilo `/foos/{id}`.

Reglas:

* Si es `cadena` se devuelve tal cual
* Si es de tipo [`github.com/gobuffalo/helpers/paths#Pathable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Pathable), se retornará el método `ToPath`.
* Si es un `slice` o de un `array`, cada elemento se pasa por el helper y luego se une.
* Si des de tipo [`github.com/gobuffalo/helpers/paths#Paramable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Paramable) se usa el método `ToParam` para rellenar el espacio `{id}`.
* Si es `<T>.Slug`, el slug se utiliza para rellenar el espacio `{id}` de la URL.
* Si es `<T>.ID` el ID se utiliza para rellenar el espacio `{id}` de la URL.

### LinkTo Helpers

### LinkTo

El helper [`tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) crea una etiqueta HTML <a> usando [`tags`](https://godoc.org/github.com/gobuffalo/tags) para crear la etiqueta con las [`tags#Options`](https://godoc.org/github.com/gobuffalo/tags#Options) dadas, y usando [`paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) para establecer el atributo `href`.

Si se le da un bloque, se interrumpirá y se añadirá dentro de la etiqueta `<a>`.

#### Ejemplo 1:

```html
<%= linkTo([user, widget], {class: "btn"}) %>

<a class="btn" href="/users/id/widget/slug"></a>
```

#### Ejemplo 2:

```html
<%= linkTo("foo", {class: "btn"}) %>

<a class="btn" href="/foo"></a>
```

#### Ejemplo 3:

```html
<%= linkTo(user, {class: "btn"}) { %>
  Click Me!
<% } %>

<a class="btn" href="/users/id">Click Me!</a>
```

### RemoteLinkTo


El helper [`tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo) proporciona la misma funcionalidad que [`tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) pero agrega el atributo `data-remote` para ser usado con [https://www.npmjs.com/package/rails-ujs](https://www.npmjs.com/package/rails-ujs) el cual se incluye en la configuración por defecto de Webpack.

#### Ejemplo 1:

```html
<%= remoteLinkTo([user, widget], {class: "btn"}) %>

<a class="btn" data-remote="true" href="/users/id/widget/slug"></a>
```

#### Ejemplo 2:

```html
<%= remoteLinkTo("foo", {class: "btn"}) %>

<a class="btn" data-remote="true" href="/foo"></a>
```

#### Ejemplo 3:

```html
<%= remoteLinkTo(user, {class: "btn"}) { %>
  Click Me!
<% } %>

<a class="btn" data-remote="true" href="/users/id">Click Me!</a>
```

## Helpers de contenido

Plush viene con dos helpers complementarios que te permiten crear fragmentos dinámicos de HTML y reutilizarlos despues en la plantilla.

### Helpers `contentFor` y `contentOf`

El helper `contentFor` toma un bloque de HTML y lo mantiene usando el nombre dado. Este bloque se puede usar en cualquier parte del archivo de plantilla, aun cuando el contenido definido en un bloque `contentFor` es una plantilla cedida y se expande dentro de un bloque `contentOf` una plantilla con llamada `yield`. La plantilla por defecto `templates/application.html` llama `yield` así:

Tomemos el siguiente ejemplo: supongamos que tenemos una plantilla `templates/application.html` que especifica completamente  todo en `<head>` y el contenido mas externo dentro de `<body>`. Esta plantilla cede a otras subplantillas, como `templates/users/show.html`, para llenar `<body>`. Sin embargo, si queremos agregar o sobrescribir  algo en el `<head>` desde una subplantilla, necesitaremos usar `contentFor`. En este ejemplo, agregaremos una manera de que las subplantillas agreguen un trozo extra de CSS al `<head>` de `application.plush.html`.

```html
<!-- aplication.plush.html -->

<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My Site</title>
    <%= stylesheetTag("application.css") %>
    <%= contentOf("extraStyle") %>
  </head>
  <body>
    <div class="container">
      <%= partial("flash.html") %>
      <%= yield %>
    </div>
  </body>
</html>
```

Resulta que a nuestra plantilla `users/index.html` le vendría bien usar un poco de estilo para toda la página en vez de agregar un monton de atributos `style` a diferentes elementos, así que se define un bloque de CSS que no aparece dentro de la plantilla:

```html
<!-- users/index.html -->

<div class="page-header">
  <h1>Users</h1>
</div>
<table class="table table-striped">
  <thead>
    <th>Username</th> <th>Password</th> <th>Email</th> <th>Admin?</th> <th>&nbsp;</th>
  </thead>
  <tbody>
    <%= for (user) in users { %>
      <!-- … -->
    <% } %>
  </tbody>
</table>

<% contentFor("extraStyle") { %>
  <style>
    .online {
      color: limegreen;
      background: black;
    }

    .offline {
      color: lightgray;
      background: darkgray;
    }
  </style>
<% } %>
```

El estilo para las clases `.outline` y `.offline` aparecen entonces al final de la etiqueta `<head>` en `/users`. En otras paginas no se agrega nada.

Por supuesto, si prefieres hacer un procesamiento extenso de lo que va en un trozo que va en una página web, es posible que desees hacer tu procesamiento en el código de Go en vez de en las plantillas. En ese caso, llama, digamos, `c.Set("moonPhase", mp)`, donde `c` es `buffalo.Context` en una función de una accion como en `actions/users.go`, y `mp` es alguna cadena u objeto. Luego en tus plantillas, haz referencia `<%= moonPhase %>` para mostrar tu fase de la luna calculada por expertos.
