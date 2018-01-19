<%= title("Using Non-Bootstrap Form Helpers") %>

The default form helpers, `form` and `form_for`, generate forms that are compatiable with Bootstrap 3. If this is not for you, you can easily use the non-Bootstrap versions of these helpers.

<em><small>Requires Plush version `v3.6.8` or greater</small></em>

<%= codeTabs() { %>
```go
// actions/render.go
func init() {
  r = render.New(render.Options{
    // ...
    // Add template helpers here:
    Helpers: render.Helpers{
      "form":     plush.FormHelper,
      "form_for": plush.FormForHelper,
    },
    // ...
  })
}
```

```erb
// templates/widgets/new.html
&lt;%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %&gt;
  &lt;%= f.InputTag("Name") %&gt;
  &lt;%= f.InputTag("Body") %&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="&lt;%= widgetsPath() %&gt;" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;% } %&gt;
```

```html
// OUTPUT
&lt;form action="/widgets" id="widget-form" method="POST"&gt;
  &lt;input name="authenticity_token" type="hidden" value="jN3nYOhCTqxZvmYnO9v1maso2VMs8fslj3rmKg1TS281W6JKpMd6Uezqp1dd3VBu2su41nKRBkd5AWDyCM4BzQ=="&gt;
  &lt;input id="widget-Name" name="Name" type="text" value=""&gt;
  &lt;input id="widget-Body" name="Body" type="text" value=""&gt;
  &lt;button class="btn btn-success" role="submit"&gt;Save&lt;/button&gt;
  &lt;a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?"&gt;Cancel&lt;/a&gt;
&lt;/form&gt;
```
<% } %>
