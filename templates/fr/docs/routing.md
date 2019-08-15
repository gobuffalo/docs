<% seoDescription("Comment gérer les routes dans Buffalo ?") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "route", "gorilla", "mux", "routeur"]) %>

<%= h1("Routage") %>

Buffalo utilise le paquet [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) pour gérer le routage au sein des applications Buffalo. L'API de `mux` est néanmoins embarquée dans celle de Buffalo. Dans ce chapitre, vous allez apprendre tout ce qu'il y a à savoir sur les routes et Buffalo.

<%= partial("fr/docs/routing/new.md") %>
<%= partial("fr/docs/routing/mapping.md") %>
<%= partial("fr/docs/routing/named_routes.md") %>
<%= partial("fr/docs/routing/templates.md") %>
<%= partial("en/docs/routing/path_for.md") %>
<%= partial("fr/docs/routing/actions.md") %>
<%= partial("fr/docs/routing/custom_named.md") %>
<%= partial("fr/docs/routing/params.md") %>
<%= partial("fr/docs/routing/named_params.md") %>
<%= partial("fr/docs/routing/groups.md") %>
<%= partial("fr/docs/routing/mounting.md") %>
