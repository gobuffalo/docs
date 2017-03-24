<h1>Templating</h1>

<%= partial("docs/disclaimer.html") %>

Buffalo defaults to using [plush](https://github.com/gobuffalo/plush) as its template engine.

<div class="video">
<iframe src="https://player.vimeo.com/video/207200621?portrait=0" width="640" height="360" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>

<%= title("General Usage", {name: "general"}) %>

<div class="code-tabs">
<%= code("html", {file: "templates/index.html"}) { %>
<h1>\<%= name %></h1>
<ul>
  \<%= for (name) in names { %>
    <li>\<%= name %></li>
  \<% } %>
</ul>
<% } %>
<%= code("go", {file: "actions/index.go"}) { %>
func IndexHandler(c buffalo.Context) error {
  c.Set("names", []string{"John", "Paul", "George", "Ringo"})
  return c.Render(200, r.HTML("index.html"))
}
<% } %>

<%= code("html", {file: "output"}) { %>
<h1>Mark</h1>
<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
<% } %>
</div>

<%= title("If Statements", {name: "if"}) %>

<%= code("html") { %>
\\<%= if (true) { %>
  \<!-- render this -->
\\<% } %>
<% } %>



<%= title("Else Statements", {name: "else"}) %>

<%= code("html") { %>
\\<%= if (false) { %>
  \<!-- won't render this -->
\\<% } else { %>
  \<!-- render this -->
\\<% } %>
<% } %>
