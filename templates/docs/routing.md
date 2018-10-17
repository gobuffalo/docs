<% seoDescription("How to handle routes in Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "route", "gorilla", "router"]) %>

# Routing

Buffalo uses the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers, to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with its own. This guide walks you through all you'll need to know about how Buffalo handles routing.

<%= partial("docs/routing/new.md") %>
<%= partial("docs/routing/mapping.md") %>
<%= partial("docs/routing/named_routes.md") %>
<%= partial("docs/routing/templates.md") %>
<%= partial("docs/routing/actions.md") %>
<%= partial("docs/routing/custom_named.md") %>
<%= partial("docs/routing/params.md") %>
<%= partial("docs/routing/named_params.md") %>
<%= partial("docs/routing/groups.md") %>
<%= partial("docs/routing/mounting.md") %>
<%= partial("docs/routing/loose_slash.md") %>
