<%= title("Session Store") %>

By default Buffalo will setup a session store using [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

This can be changed when setting up a new Buffalo application.

```go
app = buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

The ENV variable `SESSION_SECRET` should be set before running the application. If this is not set, you will see a warning in your logs that your session is not secured.

For more information on this see the docs for [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).

