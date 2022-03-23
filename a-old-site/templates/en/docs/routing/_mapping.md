## Mapping Handlers

All routing in Buffalo results in the calling of a `buffalo.Handler` function. The signature for a `buffalo.Handler` looks like this:

```go
func (c buffalo.Context) error {
  // do some work
}
```

If you already know about **MVC pattern**, `buffalo.Handler` functions manages the Controller part: this is the place where all the app logic goes. The handler function takes a `buffalo.Context` struct, which contains everything you need about the current request.

See the [context](/en/docs/context) to understand the `buffalo.Context` interface.

##### Supported HTTP Methods

Buffalo supports the following HTTP methods out of the box:

* GET
* POST
* PUT
* PATCH
* DELETE
* OPTIONS
* HEAD

You can also match all HTTP methods using `ANY`.

Mapping a `buffalo.Handler` to an HTTP method takes the form of:

```go
a.GET("/some/path", SomeHandler)
a.POST("/some/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

As you can see, you can use inline handlers if you want. For more readability though, it's often better to separate your handlers into multiple files. If you have many handlers managing users stuff, you can group them into a `users.go` file in the [`actions`](/en/docs/getting-started/directory-structure) folder, for instance.
