# Helpers

{{ partial "docs/disclaimer.html" }}

{{ partial "topics.html" }}


{{#panel title="Each Statements (Array)" name="each-array"}}

When looping through `arrays` or `slices`, the block being looped through will have access to the "global" context, as well as have four new variables available within that block:

* `@first` [`bool`] - is this the first pass through the iteration?
* `@last` [`bool`] - is this the last pass through the iteration?
* `@index` [`int`] - the counter of where in the loop you are, starting with `0`.
* `@value` - the current element in the array or slice that is being iterated over.

```handlebars
<ul>
  \{{#each names}}
    <li>\{{ @index }} - \{{ @value }}</li>
  \{{/each}}
</ul>
```

By using "block parameters" you can change the "key" of the element being accessed from `@value` to a key of your choosing.

```handlebars
<ul>
  \{{#each names as |name|}}
    <li>\{{ name }}</li>
  \{{/each}}
</ul>
```

To change both the key and the index name you can pass two "block parameters"; the first being the new name for the index and the second being the name for the element.

```handlebars
<ul>
  \{{#each names as |index, name|}}
    <li>\{{ index }} - \{{ name }}</li>
  \{{/each}}
</ul>
```

{{/panel}}

{{#panel title="Each Statements (Maps)" name="each-maps"}}

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

* `@first` [`bool`] - is this the first pass through the iteration?
* `@last` [`bool`] - is this the last pass through the iteration?
* `@key` - the key of the pair being accessed.
* `@value` - the value of the pair being accessed.

```handlebars
<ul>
  \{{#each users}}
    <li>\{{ @key }} - \{{ @value }}</li>
  \{{/each}}
</ul>
```

By using "block parameters" you can change the "key" of the element being accessed from `@value` to a key of your choosing.

```handlebars
<ul>
  \{{#each users as |user|}}
    <li>\{{ @key }} - \{{ user }}</li>
  \{{/each}}
</ul>
```

To change both the key and the value name you can pass two "block parameters"; the first being the new name for the key and the second being the name for the value.

```handlebars
<ul>
  \{{#each users as |key, user|}}
    <li>\{{ key }} - \{{ user }}</li>
  \{{/each}}
</ul>
```

{{/panel}}

{{#panel title="Other Builtin Helpers" name="other"}}

* `json` - returns a JSON marshaled string of the value passed to it.
* `js_escape` - safely escapes a string to be used in a JavaScript bit of code.
* `html_escape` - safely escapes a string to be used in an HTML bit of code.
* `upcase` - upper cases the entire string passed to it.
* `downcase` - lower cases the entire string passed to it.
* `markdown` - converts markdown to HTML.
* `eq` - works like the `if` helper, but compares to values

Velvet also imports all of the helpers found [https://github.com/markbates/inflect/blob/master/helpers.go](https://github.com/markbates/inflect/blob/master/helpers.go)

{{/panel}}

