<%= title("Sessions nulles pour APIs") %>

Si vous construisez une API, vous allez probablement devoir neutraliser les sessions (vu que les API fonctionnent très souvent sans notion de persistance de session, on parle de « stateless »). Le type [`sessions.Null`](`sessions.Null`) est la solution recommandée pour neutraliser le système de sessions.

```go
app = buffalo.New(buffalo.Options{
  Env:          ENV,
  SessionStore: sessions.Null{},
  SessionName: "_coke_session",
})
```

Lorsque vous utilisez la commande `buffalo new` avec l'option `--api`, la session utilisera par défaut une session nulle (`sessions.Null`).
