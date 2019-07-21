## Module Changes

### Buffalo Library/CLI Changes

These are changes to the Buffalo library, as well as the `buffalo` binary.

#### `github.com/gobuffalo/buffalo@<%= to %>/go.mod`

```go
<%= partial(rn.Path("buffalo", to, "buffalo", "go.mod")) %>
```

#### Diff `<%= from %>...<%= to %>`

```diff
<%= rn.Diff("buffalo", from, to, "buffalo", "_go.mod") %>
```

---

### Application Changes

These are changes for a brand new application `@<%= to %>` with the following command:

```bash
$ buffalo new coke
```

#### `<your-app>@<%= to %>/go.mod`

```go
<%= partial(rn.Path("buffalo", to, "coke", "go.mod")) %>
```

#### Diff `<%= from %>...<%= to %>`

```diff
<%= rn.Diff("buffalo", from, to, "coke", "_go.mod") %>
```
