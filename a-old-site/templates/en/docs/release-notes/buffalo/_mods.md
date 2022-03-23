## Module Changes

### Buffalo Library/CLI Changes

These are changes to the Buffalo library, as well as the `buffalo` binary.

<%= codeTabs() { %>

```go
// gobuffalo/buffalo/go.mod@<%= to %>

<%= partial(rn.Path("buffalo", to, "buffalo", "go.mod") ) %>
```

```go
// gobuffalo/buffalo/go.mod@<%= from %>

<%= partial(rn.Path("buffalo", from, "buffalo", "go.mod") ) %>
```

```diff
// DIFF
<%= rn.Diff("buffalo", from, to, "buffalo", "_go.mod") %>
```

<% } %>

---

### Application Changes

These are changes for a brand new application `@<%= to %>` with the following command:

```bash
$ buffalo new coke
```

<%= codeTabs() { %>

```go
// <your-app>@<%= to %>/go.mod

<%= partial(rn.Path("buffalo", to, "coke", "go.mod")) %>
```

```go
// <your-app>@<%= from %>/go.mod

<%= partial(rn.Path("buffalo", from, "coke", "go.mod")) %>
```

```diff
// DIFF
<%= rn.Diff("buffalo", from, to, "coke", "_go.mod") %>
```

<% } %>

