# Iterating

<%= partial("en/docs/disclaimer.html") %>

## Iterating Through Arrays

When looping through `arrays` or `slices`, the block being looped through will have access to the "global" context.

The `for` statement takes 1 - 2 arguments. When using the two argument version, the first argument is the "index" of the loop and the second argument is the value from the array or slice.

```html
&lt;ul>
  \<%= for (index, name) in names { %>
    &lt;li>\<%= index %> - \<%= name %></li>
  \<% } %>
&lt;/ul>
```

When using the one argument version the index is omitted and just the value is returned:

```html
&lt;ul>
  \<%= for (name) in names { %>
    &lt;li>\<%= name %></li>
  \<% } %>
&lt;/ul>
```

## Iterating Through Maps

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

When using the two argument version, the first argument is the key of the map and the second argument is the value from the map:

```html
&lt;ul>
  \<%= for (key, value) in users { %>
    &lt;li>\<%= key %> - \<%= value %></li>
  \<% } %>
&lt;/ul>
```

When using the one argument version the key is omitted and just the value is returned:

```html
&lt;ul>
  \<%= for (user) in users { %>
    &lt;li>\<%= user %></li>
  \<% } %>
&lt;/ul>
```
