<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Videos"}) %>
<% } %>

<%= for (video) in videoList { %>
# <%= video.Title %>
  <%= raw(video.Description) %>
  <%= vimeoFromVideo(video) %>
<% } %>
