# Videos

# Buffalo: Managing Plugins

In this video Mark Bates will walk you through understanding how Buffalo Plugins work, and how to use the Buffalo Plugin Manager.

[https://www.gopherguides.tv/buffalo-resources/videos/buffalo-managing-plugins](https://www.gopherguides.tv/buffalo-resources/videos/buffalo-managing-plugins)

<%= for (video) in videoList { %>
# <%= video.Title %>
  <%= raw(video.Description) %>
  <%= vimeoFromVideo(video) %>
<% } %>
