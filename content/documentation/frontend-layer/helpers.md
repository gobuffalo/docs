---
Name: "Helpers"
weight: 6
aliases:
  - /docs/helpers
  - /en/docs/helpers
---

# Helpers

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

## Builtin Helpers

A full list of all helper functions for [`github.com/gobuffalo/plush`](https://godoc.org/github.com/gobuffalo/plush) can be found at [`github.com/gobuffalo/helpers`](https://godoc.org/github.com/gobuffalo/helpers).

## Path Helpers

Buffalo will generate path helpers for all of the routes you add to the App. The easiest way to see what all of the generated path helpers are and what they point to is to run `buffalo routes`. This will print out a list that looks something like this:

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

Going down this list we start with the path *NAME*d `rootPath` which represents *PATH* `/` or the root route of the server and as a bonus, with all of these we can even see exactly which *HANDLER* code is being run for this METHOD+PATH combination.

Next we have a standard `app.GET("/about", AboutHandler)` which generates to `aboutPath`.

Then we use a resource `app.Resource("/drinks", DrinksResource{})`, which generates a path for each of our standard actions, and for each of those a helper to be used in templates. Those that take a parameter can be used like this `<%= drinkPath({drink_id: drink.ID}) %>`. All helpers take a `map[string]interface{}` that is used to fill-in parameters.

Finally, when we use a group we can see that this changes the generated helpers. Here is the routing for those last paths:

```go
api := app.Group("/api/v1")
api.Resource("/users", UsersResource{})
```

**Note** that the helpers are generated to match the generated paths. It is possible to override the path names in the `App.Routes`, but it is highly advised that you find a different way to your goal than this. Slack is always open to these conversations.

### PathFor Helper

The [`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) helper takes an `interface{}`, or a `slice` of them, and tries to convert it to a `/foos/{id}` style URL path.

Rules:

* if `string` it is returned as is
* if [`github.com/gobuffalo/helpers/paths#Pathable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Pathable) the `ToPath` method is returned
* if `slice` or an `array` each element is run through the helper then joined
* if [`github.com/gobuffalo/helpers/paths#Paramable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Paramable) the `ToParam` method is used to fill the `{id}` slot
* if `<T>.Slug` the slug is used to fill the `{id}` slot of the URL
* if `<T>.ID` the ID is used to fill the `{id}` slot of the URL

### LinkTo Helpers

### LinkTo

The [`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo)
helpers creates HTML for a `<a>` tag using [`github.com/gobuffalo/tags`](https://godoc.org/github.com/gobuffalo/tags)
to create tag with the given [`github.com/gobuffalo/tags#Options`](https://godoc.org/github.com/gobuffalo/tags#Options) and using
[`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) to set the `href` element.

If given a block it will be interrupted and appended inside of the `<a>` tag.

#### Example 1:

```erb
<%= linkTo([user, widget], {class: "btn"}) %>

<a class="btn" href="/users/id/widget/slug"></a>
```

#### Example 2:

```erb
<%= linkTo("foo", {class: "btn"}) %>

<a class="btn" href="/foo"></a>
```

#### Example 3:

```erb
<%= linkTo(user, {class: "btn"}) { %>
Click Me!
<% } %>

<a class="btn" href="/users/id">Click Me!</a>
```

### RemoteLinkTo

The [`github.com/gobuffalo/helpers/tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo) helper provides the same functionality as
[`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) but adds the `data-remote` element for use with
[https://www.npmjs.com/package/rails-ujs](https://www.npmjs.com/package/rails-ujs) which is included in the default generated Webpack configuration.

#### Example 1:

```erb
<%= remoteLinkTo([user, widget], {class: "btn"}) %>

<a class="btn" data-remote="true" href="/users/id/widget/slug"></a>
```

#### Example 2:

```erb
<%= remoteLinkTo("foo", {class: "btn"}) %>

<a class="btn" data-remote="true" href="/foo"></a>
```

#### Example 3:

```erb
<%= remoteLinkTo(user, {class: "btn"}) { %>
Click Me!
<% } %>

<a class="btn" data-remote="true" href="/users/id">Click Me!</a>
```

## Content Helpers

Plush ships with two complementary helpers that let you create dynamic HTML snippets and re-use them later in the template.

### The `contentFor` and `contentOf` Helpers

The `contentFor` helper takes a block of HTML and holds on to it using the given name. This block can then be used elsewhere in a template file, even when the content defined in a `contentFor` block is in a yielded-to template and is expanded into a `contentOf` block in a `yield`-calling template. The default `templates/application.html` calls `yield` like this.

Take the following example: suppose we have a `templates/application.html` that fully specifies everything in `<head>` and the outermost contents of `<body>`. This template yields to other subtemplates, like `templates/users/show.html`, to fill `<body>`. However, if we want to add or override something in the `<head>` from a subtemplate, we'll need to use `contentFor`. In this example, we'll add a way for subtemplates to add an extra chunk of CSS to the `<head>` of `application.html`:

```erb
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

As it turns out, our `users/index.html` template could use a little page-wide styling instead of adding a bunch of `style` attributes to different elements, so it defines a block of CSS that doesn't show up anywhere inside the template:

```erb
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

The styling for the `online` and `offline` classes then appears at the end of `<head>` in `/users`. In other pages, nothing is added.

Of course, if you'd rather do extensive processing on what goes into a chunk that goes on a webpage, you may want to do your processing in Go code instead of in templates. In that case, call, say, `c.Set("moonPhase", mp)` where `c` is a `buffalo.Context` in a function in an action like in `actions/users.go`, and `mp` is some string or object. Then, in your templates, refer to `<%= moonPhase %>` to display your expertly-calculated phase of the moon.

