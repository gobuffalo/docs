<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Plugins"}) %>
<% } %>

<%= sinceVersion("0.9.1") %>

Plugins allow for 3rd party code to extend the `buffalo` command as well as its sub-commands.


<%= partial("docs/plugins/installation.md") %>
<%= partial("docs/plugins/finding.md") %>
<%= partial("docs/plugins/search_paths.md") %>
<%= partial("docs/plugins/installing.md") %>
<%= partial("docs/plugins/removing.md") %>
<%= partial("docs/plugins/writing.md") %>

<%= partial("docs/plugins/no_go.md") %>
