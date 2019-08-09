# Middleware

Middleware allows for the interjection of code in the request/response cycle. Common use cases for middleware are things like logging (which Buffalo already does), authentication requests, etc.

A list of "known" middleware packages can be found at [https://toolkit.gobuffalo.io/tools?topic=middleware](https://toolkit.gobuffalo.io/tools?topic=middleware).

<%= partial("en/docs/middleware/interface.md") %>
<%= partial("en/docs/middleware/using.md") %>
<%= partial("en/docs/middleware/one_action.md") %>
<%= partial("en/docs/middleware/group.md") %>
<%= partial("en/docs/middleware/skipping.md") %>
<%= partial("en/docs/middleware/skipping_resource.md") %>
<%= partial("en/docs/middleware/clearing.md") %>
<%= partial("en/docs/middleware/listing.md") %>
