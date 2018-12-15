<%= title("Using Route Helpers in Templates") %>

Route helpers can be used directly in templates using the name of the helper:

```html
\<%= widgetsPath() %> // /widgets
```

Routes that require named parameters, must be feed a map of those parameters.

```html
\<%= editWidgetPath({widget_id: 1}) %> // /widgets/1/edit
```
