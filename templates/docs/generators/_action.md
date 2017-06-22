<%= title("Actions") %>

<%= code("text") { %>
$ buffalo g action --help

Generates new action(s)

Usage:
  buffalo generate action [name] [actionName...] [flags]

Aliases:
  action, a, actions
<% } %>

<%= code("text") { %>
$ buffalo g a users show index create

--> templates/users/show.html
--> templates/users/index.html
--> templates/users/create.html
--> actions/users.go
--> actions/users_test.go
--> goimports -w .

<% } %>

In some cases you will neeed to generate an action with an HTTP method different than GET, for that case you can use the --method flag, like in the following example:

<%= code("text") { %>
$ buffalo g actions users message --method POST
<% } %>

In some other scenarios you will need to generate an action without generating an HTML template, to skip the generation of the HTML template for your newly created action you can pass the --skip-template flag to the generator, p.e:

<%= code("text") { %>
$ buffalo g actions users update --method POST
<% } %>

