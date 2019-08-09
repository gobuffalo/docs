<% seoDescription("Listening for events in a Buffalo application") %>
<% seoKeywords(["buffalo", "go", "golang", "events", "plugins"]) %>

# Events

<%= sinceVersion("0.13.0-beta.2") %>

The <%= doclink("github.com/gobuffalo/events") %> package allows for Go applications, including Buffalo applications, to listen, and emit, global event messages.

<%= partial("en/docs/events/listening.md") %>
<%= partial("en/docs/events/emitting.md") %>
<%= partial("en/docs/events/filtering.md") %>
<%= partial("en/docs/events/stop_listening.md") %>
<%= partial("en/docs/events/plugins.md") %>
<%= partial("en/docs/events/message_queue.md") %>
<%= partial("en/docs/events/known.md") %>
