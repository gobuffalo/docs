## Plugins Moved Into Core

The [`gobuffalo/buffalo-plugins`](https://godoc.org/github.com/gobuffalo/buffalo-plugins) package, and binary, have been moved into the core [`gobuffalo/buffalo`](https://godoc.org/github.com/gobuffalo/buffalo) repository.

This does **not** mean that plugins like [`gobuffalo/buffalo-pop`](https://godoc.org/github.com/gobuffalo/buffalo-pop) or [`gobuffalo/buffalo-heroku`](https://godoc.org/github.com/gobuffalo/buffalo-heroku) are going to be brought into core.

We are pulling in the tool chain that third-party plugins use to develop with.

<%= codeTabs() { %>

```go
// buffalo-coke/cmd/available.go@<%= to %>

<%= partial(rn.Path("buffalo", to, "coke", "available.go.txt") ) %>
```

```go
// buffalo-coke/cmd/available.go@<%= from %>

<%= partial(rn.Path("buffalo", from, "coke", "available.go.txt") ) %>
```

```diff
// DIFF
<%= rn.Diff("buffalo", from, to, "coke", "_available.go.txt") %>
```

<% } %>
