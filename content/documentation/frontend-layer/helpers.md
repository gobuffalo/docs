---
Name: "Helpers"
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
METHOD | PATH                         | ALIASES | NAME              | HANDLER
------ | ----                         | ------- | ----              | -------
GET    | /                            |         | rootPath          | github.com/gobuffalo/coke/actions.HomeHandler
GET    | /about                       |         | aboutPath         | github.com/gobuffalo/coke/actions.AboutHandler
GET    | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.List
POST   | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.Create
GET    | /drinks/new                  |         | newDrinksPath     | github.com/gobuffalo/coke/actions.DrinksResource.New
GET    | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Show
PUT    | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Update
DELETE | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Destroy
GET    | /drinks/{drink_id}/edit      |         | editDrinkPath     | github.com/gobuffalo/coke/actions.DrinksResource.Edit
GET    | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.List
POST   | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.Create
GET    | /api/v1/users/new            |         | newApiV1UsersPath | github.com/gobuffalo/coke/actions.UsersResource.New
GET    | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Show
PUT    | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Update
DELETE | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Destroy
GET    | /api/v1/users/{user_id}/edit |         | editApiV1UserPath | github.com/gobuffalo/coke/actions.UsersResource.Edit
```

Going down this list we start with the path *NAME*d `rootPath` which represents *PATH* `/` or the root route of the server and as a bonus, with all of these we can even see exactly which *HANDLER* code is being run for this METHOD+PATH combination.

Next we have a standard `app.GET("/about", AboutHandler)` which generates to `aboutPath`.

Then we use a resource `app.Resource("/drinks", DrinksResource{})`, which generates a path for each of our standard actions, and for each of those a helper to be used in templates. Those that take a parameter can be used like this `\<%= drinkPath({drink_id: drink.ID}) %>`. All helpers take a `map[string]interface{}` that is used to fill-in parameters.

Finally, when we use a group we can see that this changes the generated helpers. Here is the routing for those last paths:

```
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

```html
&lt;%= linkTo([user, widget], {class: "btn"}) %&gt;

&lt;a class="btn" href="/users/id/widget/slug"&gt;&lt;/a&gt;
```

#### Example 2:

```html
&lt;%= linkTo("foo", {class: "btn"}) %&gt;

&lt;a class="btn" href="/foo"&gt;&lt;/a&gt;
```

#### Example 3:

```html
&lt;%= linkTo(user, {class: "btn"}) { %&gt;
Click Me!
&lt;% } %&gt;

&lt;a class="btn" href="/users/id"&gt;Click Me!&lt;/a&gt;
```

### RemoteLinkTo

The [`github.com/gobuffalo/helpers/tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo) helper provides the same functionality as
[`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) but adds the `data-remote` element for use with
[https://www.npmjs.com/package/rails-ujs](https://www.npmjs.com/package/rails-ujs) which is included in the default generated Webpack configuration.

#### Example 1:

```html
&lt;%= remoteLinkTo([user, widget], {class: "btn"}) %&gt;

&lt;a class="btn" data-remote="true" href="/users/id/widget/slug"&gt;&lt;/a&gt;
```

#### Example 2:

```html
&lt;%= remoteLinkTo("foo", {class: "btn"}) %&gt;

&lt;a class="btn" data-remote="true" href="/foo"&gt;&lt;/a&gt;
```

#### Example 3:

```html
&lt;%= remoteLinkTo(user, {class: "btn"}) { %&gt;
Click Me!
&lt;% } %&gt;

&lt;a class="btn" data-remote="true" href="/users/id"&gt;Click Me!&lt;/a&gt;
```

## Content Helpers

Plush ships with two complementary helpers that let you create dynamic HTML snippets and re-use them later in the template.

### The `contentFor` and `contentOf` Helpers

The `contentFor` helper takes a block of HTML and holds on to it using the given name. This block can then be used elsewhere in a template file, even when the content defined in a `contentFor` block is in a yielded-to template and is expanded into a `contentOf` block in a `yield`-calling template. The default `templates/application.html` calls `yield` like this.

Take the following example: suppose we have a `templates/application.html` that fully specifies everything in `&lt;head>` and the outermost contents of `&lt;body>`. This template yields to other subtemplates, like `templates/users/show.html`, to fill `&lt;body>`. However, if we want to add or override something in the `&lt;head>` from a subtemplate, we'll need to use `contentFor`. In this example, we'll add a way for subtemplates to add an extra chunk of CSS to the `&lt;head>` of `application.html`:

```html
&lt;!DOCTYPE html>
&lt;html>
  &lt;head>
    &lt;meta charset="utf-8">
    &lt;title>My Site&lt;/title>
    \<%= stylesheetTag("application.css") %>
    \<%= contentOf("extraStyle") %>
  &lt;/head>
  &lt;body>
    &lt;div class="container">
      \<%= partial("flash.html") %>
      \<%= yield %>
    &lt;/div>
  &lt;/body>
&lt;/html>
```

As it turns out, our `users/index.html` template could use a little page-wide styling instead of adding a bunch of `style` attributes to different elements, so it defines a block of CSS that doesn't show up anywhere inside the template:

```html
&lt;div class="page-header">
  &lt;h1>Users&lt;/h1>
&lt;/div>
&lt;table class="table table-striped">
  &lt;thead>
    &lt;th>Username&lt;/th> &lt;th>Password&lt;/th> &lt;th>Email&lt;/th> &lt;th>Admin?&lt;/th> &lt;th>&nbsp;&lt;/th>
  &lt;/thead>
  &lt;tbody>
    \<%= for (user) in users { %>
      &lt;!-- … -->
    \<% } %>
  &lt;/tbody>
&lt;/table>

\<% contentFor("extraStyle") { %>
  &lt;style>
    .online {
      color: limegreen;
      background: black;
    }

    .offline {
      color: lightgray;
      background: darkgray;
    }
  &lt;/style>
\<% } %>
```

The styling for the `online` and `offline` classes then appears at the end of `&lt;head>` in `/users`. In other pages, nothing is added.

Of course, if you'd rather do extensive processing on what goes into a chunk that goes on a webpage, you may want to do your processing in Go code instead of in templates. In that case, call, say, `c.Set("moonPhase", mp)` where `c` is a `buffalo.Context` in a function in an action like in `actions/users.go`, and `mp` is some string or object. Then, in your templates, refer to `\<%= moonPhase %>` to display your expertly-calculated phase of the moon.

