# Layouts

<%= partial("docs/disclaimer.html") %>

<%= title("Using a Standard Layout", {name: "standard"})  %>

It is quite common to want to use the same layout across most, if not all of an application. When creating a new `render.Engine` the `HTMLLayout` property can be set to a file that will automatically be used by the `render.HTML` function.

<div class="code-tabs">
<%= code("go", {file: "actions/render.go"}) { %>
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout:     "application.html",
    // ...
  })
}
<% } %>

<%= code("html", {file: "templates/application.html"}) { %>
&lt;html>
  &lt;head>
    &lt;title>My App</title>
  &lt;/head>
  &lt;body>
    <div id="main">
      \<%= yield %>
    </div>
  &lt;/body>
&lt;/html>
<% } %>

<%= code("html", {file: "templates/hello.html"}) { %>
<h1>Hello!!</h1>
<% } %>

<%= code("go", {file: "actions/hello.go"}) { %>
package actions

func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html"))
}
<% } %>

<%= code("html", {file: "output"}) { %>
&lt;html>
  &lt;head>
    &lt;title>My App</title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      &lt;h1>Hello!!</h1>
    &lt;/div>
  &lt;/body>
&lt;/html>
<% } %>
</div>


<%= title("Using a Custom Layout", {name: "custom"})  %>

Sometimes, on certain requests, a different layout is needed. This alternate layout can be passed in as the second parameter to `render.HTML`.

<div class="code-tabs">

<%= code("go", {file: "actions/render.go"}) { %>
var r *render.Engine

func init() {
  r = render.New(render.Options{
    // ...
    HTMLLayout:     "application.html",
    // ...
  })
}
<% } %>

<%= code("html", {file: "templates/custom.html"}) { %>
&lt;html>
  &lt;head>
    &lt;title>My Custom Layout&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      \<%= yield %>
    &lt;/div>
  &lt;/body>
&lt;/html>
<% } %>

<%= code("html", {file: "templates/hello.html"}) { %>
<h1>Hello!!</h1>
<% } %>

<%= code("go", {file: "actions/hello.go"}) { %>
package actions

func Hello(c buffalo.Context) error {
  return c.Render(200, r.HTML("hello.html", "custom.html"))
}
<% } %>

<%= code("html", {file: "output"}) { %>
&lt;html>
  &lt;head>
    &lt;title>My Custom Layout&lt;/title>
  &lt;/head>
  &lt;body>
    &lt;div id="main">
      &lt;h1>Hello!!</h1>
    &lt;/div>
  &lt;/body>
&lt;/html>
<% } %>
</div>
