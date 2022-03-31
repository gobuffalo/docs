## New Link Helpers

The new [`github.com/gobuffalo/helpers/tags`](https://godoc.org/github.com/gobuffalo/helpers/tags) package brings some new helper functions.

### LinkTo

The [`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo)
helpers creates HTML for a `<a>` tag using [`github.com/gobuffalo/tags`](https://godoc.org/github.com/gobuffalo/tags)
to create tag with the given [`github.com/gobuffalo/tags#Options`](https://godoc.org/github.com/gobuffalo/tags#Options) and using
[`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) to set the `href` element.

If given a block it will be interrupted and appended inside of the `<a>` tag.

#### Example 1:

```html
\<%= linkTo([user, widget], {class: "btn"}) %>

<a class="btn" href="/users/id/widget/slug"></a>
```

#### Example 2:

```html
\<%= linkTo("foo", {class: "btn"}) %>

<a class="btn" href="/foo"></a>
```

#### Example 3:

```html
\<%= linkTo(user, {class: "btn"}) { %>
Click Me!
\<% } %>

<a class="btn" href="/users/id">Click Me!</a>
```

### RemoteLinkTo

The [`github.com/gobuffalo/helpers/tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo) helper provides the same functionality as
[`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) but adds the `data-remote` element for use with
[https://www.npmjs.com/package/rails-ujs](https://www.npmjs.com/package/rails-ujs) which is included in the default generated Webpack configuration.

#### Example 1:

```html
\<%= remoteLinkTo([user, widget], {class: "btn"}) %>

<a class="btn" data-remote="true" href="/users/id/widget/slug"></a>
```

#### Example 2:

```html
\<%= remoteLinkTo("foo", {class: "btn"}) %>

<a class="btn" data-remote="true" href="/foo"></a>
```

#### Example 3:

```html
\<%= remoteLinkTo(user, {class: "btn"}) { %>
Click Me!
\<% } %>

<a class="btn" data-remote="true" href="/users/id">Click Me!</a>
```
