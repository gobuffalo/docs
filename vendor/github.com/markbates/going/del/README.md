# del

A wrapper around the [render](https://github.com/unrolled/render) package to make it work with [echo](https://github.com/labstack/echo).

## Installation

```bash
$ go get https://github.com/markbates/going/del
```

## Usage

```go
e := echo.New()

e.SetRenderer(del.New(del.DefaultOptions))

e.Get("/", func(c *echo.Context) error {
  return c.Render(200, "a/b", "Mark")
})

e.Run(":9000")
```
