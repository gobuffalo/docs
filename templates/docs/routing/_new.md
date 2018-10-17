<%= title("Creating a new Buffalo App (and router)") %>

The app configuration is located in the `app.go` file.

```go
a := buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
})
```

The default setup should handle most of your needs, but you are free to customize it to fit your use case.

You can check the available options list here: [https://godoc.org/github.com/gobuffalo/buffalo#Options](https://godoc.org/github.com/gobuffalo/buffalo#Options)
