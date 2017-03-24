# Partials

<%= partial("docs/disclaimer.html") %>

<%= title("Naming") %>

All partial file names must start with an `_`. For example: `_form.html`. This helps to differentiate partials from other view templates in your application.

<div class="code-tabs">
<%= code("html", {file: "templates/users/new.html"}) { %>
<h1>Create New User</h1>

\<%= partial("users/form.html") %>
<% } %>

<%= code("html", {file: "templates/users/_form.html"}) { %>
<h1>Create New User</h1>

&lt;form action="/users">
&lt;!-- form stuff here  -->
&lt;form>
<% } %>

<%= code("html", {file: "output"}) { %>
<h1>Create New User</h1>

&lt;form action="/users">
&lt;!-- form stuff here  -->
&lt;form>
<% } %>
</div>


<%= title("Context") %>

All [rendering context](/docs/rendering) from the parent template will automatically pass through to the partial, and any partials that partial may call. (see also [context](/docs/context))

<div class="code-tabs">

<%= code("go", {file: "actions/users.go"}) { %>
func UsersEdit(c buffalo.Context) error {
  // do some work to find the user
  c.Set("user", user)
  return c.Render(200, render.HTML("users/edit.html"))
}
<% } %>

<%= code("html", {file: "templates/users/edit.html"}) { %>
<h1>Edit \<%= user.Name %> (\<%= user.ID %>)</h1>

\<%= partial("users/form.html") %>
<% } %>

<%= code("html", {file: "templates/users/_form.html"}) { %>
&lt;form action="/users/\<%= user.ID %>">
&lt;!-- form stuff here  -->
&lt;/form>
<% } %>

<%= code("html", {file: "output"}) { %>
<h1>Edit Mark Bates (1)</h1>

&lt;form action="/users/1">
&lt;!-- form stuff here  -->
&lt;/form>
<% } %>
</div>


<%= title("Local Context", {name: "local-context"}) %>

In addition to have the [context](/docs/context) of the parent template, partials can also be sent additional information as "local" variables.

<div class="code-tabs">

<%= code("go", {file: "actions/users.go"}) { %>
func UsersIndex(c buffalo.Context) error {
  c.Set("users", []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"})
  return c.Render(r.HTML("users/index.html"))
}
<% } %>

<%= code("html", {file: "templates/users/index.html"}) { %>
<h1>All Users</h1>

<ul>
  \<%= for (u) in users { %>
    \<%= partial("users/user.html", {user: u}) %>
  \<% } %>
</ul>
<% } %>

<%= code("html", {file: "templates/users/_user.html"}) { %>
<li>\<%= user.Name %></li>
<% } %>

<%= code("html", {file: "output"}) { %>
<h1>All Users</h1>

<ul>
  <li>John Lennon</li>
  <li>Paul McCartney</li>
  <li>George Harrison</li>
  <li>Ringo Starr</li>
</ul>
<% } %>
</div>

<%= title("Helpers") %>

Partials are not much different from standard [templates](/docs/templating) in Buffalo. They include all of the same [helpers](/docs/helpers) as well.
