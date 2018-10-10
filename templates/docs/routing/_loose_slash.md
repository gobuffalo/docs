<%= title("Loose Slash", {}) %>

<%= sinceVersion("0.13.0-beta.1") %>

The `LooseSlash` option was removed in `v0.13.0-beta.1` in this [PR](https://github.com/gobuffalo/buffalo/pull/1168).

What this PR does is normalizes Buffalo to using a trailing `/`. However, it does not force this behavior on the end user or the end developer.

The Gorilla Mux [`StrictSlash`](http://www.gorillatoolkit.org/pkg/mux#Router.StrictSlash) has been removed and is now handled on the Buffalo side.

```go
table := []struct {
  mapped   string
  browser  string
  expected string
}{
  {"/foo", "/foo", "/foo/"},
  {"/foo", "/foo/", "/foo/"},
  {"/foo/", "/foo", "/foo/"},
  {"/foo/", "/foo/", "/foo/"},
  {"/index.html", "/index.html", "/index.html"},
  {"/foo.gif", "/foo.gif", "/foo.gif"},
}
```

Regardless of whether the developer maps `/foo` or `/foo/` and regardless of which variant the user goes to, the router will respond the same with a `200`.

**NOTE:** There is a slight breaking change in that incoming path `c.Request().URL.Path` will **ALWAYS** have the trailing slash.

---

<%= sinceVersion("0.10.3") %>

By default, the configured routes for your app **match strictly** the pattern you defined: if the pattern ends with a slash, the URL won't be accessible without a slash. Conversely, a pattern without an ending slash won't match an URL with an ending slash.

To allow your routes to ignore the ending slash, you can use the `LooseSlash` option:

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  LooseSlash:  true,
})
```
