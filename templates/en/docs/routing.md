<% seoDescription("How to handle routes in Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "http", "route", "gorilla", "router"]) %>

# Routing

Buffalo uses the [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package under the covers, to handle routing within Buffalo applications. With that said, Buffalo wraps the `mux` API with its own. This guide walks you through all you'll need to know about how Buffalo handles routing.

<%= partial("en/docs/routing/new.md") %>
<%= partial("en/docs/routing/mapping.md") %>
<%= partial("en/docs/routing/named_routes.md") %>
<%= partial("en/docs/routing/templates.md") %>
<%= partial("en/docs/routing/actions.md") %>
<%= partial("en/docs/routing/custom_named.md") %>
<%= partial("en/docs/routing/params.md") %>
<%= partial("en/docs/routing/named_params.md") %>
<%= partial("en/docs/routing/groups.md") %>
<%= partial("en/docs/routing/mounting.md") %>
