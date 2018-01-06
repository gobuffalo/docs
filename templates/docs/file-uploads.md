# File Uploads
<%= sinceVersion("0.10.3") %>

Buffalo allows for the easily handling of files uploaded from a form. Storing those files, such as to disk or S3, is up to you the end developer, Buffalo just gives you easy access to the file from the request.

<%= title("Configuring the Form") %>

The `f.FileTag` form helper can be used to quickly add a file element to the form. When using this the `enctype` of the form is *automatically* switched to be `multipart/form-data`.

```erb
&lt;%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %&gt;
  &lt;%= f.InputTag("Name") %&gt;
  &lt;%= f.FileTag("MyFile") %&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="&lt;%= widgetsPath() %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

<%= partial("docs/uploads/file.md") %>
<%= partial("docs/uploads/model.md") %>
<%= partial("docs/uploads/test.md") %>
