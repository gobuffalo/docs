## Utiliser les *helpers* de routes dans les templates

Les *helpers* de toutes peuvent être directement utilisés dans les templates, en utilisant le nom du *helper* :

```html
\<%= widgetsPath() %> // /widgets
```

Les routes qui nécessitent des paramètres nommés doivent recevoir une *map* avec ces paramètres.

```html
\<%= editWidgetPath({widget_id: 1}) %> // /widgets/1/edit
```
