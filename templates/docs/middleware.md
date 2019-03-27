<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Middleware"}) %>
<% } %>

Middleware allows for the interjection of code in the request/response cycle. Common use cases for middleware are things like logging (which Buffalo already does), authentication requests, etc.

A list of "known" middleware packages can be found at [https://toolkit.gobuffalo.io/tools?topic=middleware](https://toolkit.gobuffalo.io/tools?topic=middleware).

<%= partial("docs/middleware/interface.md") %>
<%= partial("docs/middleware/using.md") %>
<%= partial("docs/middleware/one_action.md") %>
<%= partial("docs/middleware/group.md") %>
<%= partial("docs/middleware/skipping.md") %>
<%= partial("docs/middleware/skipping_resource.md") %>
<%= partial("docs/middleware/clearing.md") %>
<%= partial("docs/middleware/listing.md") %>
