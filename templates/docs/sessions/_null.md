<%= title("Null Sessions for APIs") %>

When building API servers the default cookie session store is undesirable. The [`sessions.Null`](`sessions.Null`) type is the recommended replacement for the default session store.

```go
app = buffalo.New(buffalo.Options{
  Env:          ENV,
  SessionStore: sessions.Null{},
  SessionName: "_coke_session",
})
```

When running `buffalo new` with the `--api` flag the default session will be set to `sessions.Null`.
