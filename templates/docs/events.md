<% seoDescription("Listening for events in a Buffalo application") %>
<% seoKeywords(["buffalo", "go", "golang", "events", "plugins"]) %>

# Events

<%= sinceVersion("0.13.0-beta.2") %>

The <%= doclink("github.com/gobuffalo/events") %> package allows for Go applications, including Buffalo applications, to listen, and emit, global event messages.

<%= partial("docs/events/listening.md") %>
<%= partial("docs/events/emitting.md") %>
<%= partial("docs/events/filtering.md") %>
<%= partial("docs/events/stop_listening.md") %>
<%= partial("docs/events/plugins.md") %>
<%= partial("docs/events/message_queue.md") %>
<%= partial("docs/events/known.md") %>
