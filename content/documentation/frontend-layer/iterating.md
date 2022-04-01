---
Name: "Iterating"
---

# Iterating

<p>
<em>This document only applies when using <a href="https://github.com/gobuffalo/buffalo/tree/master/render" rel="nofollow">https://github.com/gobuffalo/buffalo/render</a>.</em><br>
<em>Please see <a href="https://github.com/gobuffalo/plush" target="_blank">github.com/gobuffalo/plush</a> for more details on the underlying templating package.</em>
</p

## Iterating Through Arrays

When looping through `arrays` or `slices`, the block being looped through will have access to the "global" context.

The `for` statement takes 1 - 2 arguments. When using the two argument version, the first argument is the "index" of the loop and the second argument is the value from the array or slice.

```html
<ul>
  <%= for (index, name) in names { %>
    <li><%= index %> - <%= name %></li>
  <% } %>
</ul>
```

When using the one argument version the index is omitted and just the value is returned:

```html
<ul>
  <%= for (name) in names { %>
    <li><%= name %></li>
  <% } %>
</ul>
```

## Iterating Through Maps

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

When using the two argument version, the first argument is the key of the map and the second argument is the value from the map:

```html
<ul>
  <%= for (key, value) in users { %>
    <li><%= key %> - <%= value %></li>
  <% } %>
</ul>
```

When using the one argument version the key is omitted and just the value is returned:

```html
<ul>
  <%= for (user) in users { %>
    <li><%= user %></li>
  <% } %>
</ul>
```
