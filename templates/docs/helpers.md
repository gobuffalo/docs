# Helpers

<%= partial("docs/disclaimer.html") %>

<%= title("Builtin Helpers") %>

Listed below are a few of the helpers that ship with Plush. Please see the Plush [documentation](https://github.com/gobuffalo/plush) for more information on its helpers.

* `json` - converts the interface to a JSON object
* `jsEscape` - escapes the interface to be JavaScript safe
* `htmlEscape` - escapes the interface to be HTML safe
* `upcase` - converts the string to upper case
* `downcase` - converts the string to lower case
* `contentFor` - stores a block of HTML to be used later
* `contentOf` - retrieves a block of HTML previously stored with `contentFor`
* `markdown` - converts the string from Markdown into HTML
* `len` - returns the length of the interface
* `debug` - returns the `%+v` of the interface wrapped in `%lt;pre%gt;` tags.
* `inspect` - returns the `%+v` of the interface
* `range` - interate between, and including two numbers
* `between` - iterate between, but not including, two numbers
* `until` - iterate until a number is reached
* `groupBy` - splits a slice or array into `n` groups
* `env` - returns the ENV variable for the specified key
* `truncate` - truncates a string to a specified length
* `raw` - converts a string to `template.HTML`
* `form` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version)
* `form_for` - support for the [github.com/gobuffalo/tags/form](https://github.com/gobuffalo/tags/tree/master/form) package (Bootstrap version) to build a form for a model

Plush also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

<%= title("Path Helpers") %>

Buffalo will generate path helpers for all of the routes you add to the App. The easiest way to see what all of the generated path helpers are and what they point to is to run `buffalo routes`. This will print out a list that looks something like this:

```
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

Going down this list we start with the path *NAME*d `rootPath` which represents *PATH* `/` or the root route of the server and as a bonus with all of these we can even see exactly which *HANDLER* code is being run for this METHOD+PATH combination.

Next we have a standard `app.GET("/about", AboutHandler)` which generates to `aboutPath`. 

Then we use a resource `app.Resource("/drinks", DrinksResource{})` which generates a path for each of our standard actions, and for each of those a helper to be used in templates. Those that take a parameter can be used like this `<%= drinkPath({drink_id: drink.ID}) %>`. All helpers take a `map[string]interface{}` that is used to fill-in parameters.

Finally, when we use a group we can see that this changes the generated helpers. Here is the routing for those last paths:

```
api := app.Group("/api/v1")
api.Resource("/users", UsersResource{})
```

**Note** that the helpers are generated to match the generated paths. It is possible to override the path names in the `App.Routes`, but it is highly adviced that you find a different way to your goal than this. Slack is always open to these conversations.

<%= title("Content Helpers") %>

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

