# Forms

<%= title("Creating Forms") %>

Forms utilize the `form_for` function. A simple code generator that will create the HTML for your form, would look like following:

```go
<h1>New Post</h1>
<%= form_for(post, {action: postsPath(), method:"POST"}) { %>
  <%= partial("posts/form.html") %><ahref="<%= postsPath() %>" class="btn btn-warning"data-confirm="Are you sure?">Cancel</a><% } %>
```