# Go Modules

**NOTE**: Please read [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules) to understand more about Go Modules **before** using them.

## Enabling Go Module Support

<%= sinceVersion("v0.13.0") %>

The support for Go Modules in "Buffalo" packages is experimental, as are Go Modules (as of `v1.11.x`). To "opt-in" to using Go Modules you need to turn them using the `GO111MODULE` environment variable and setting it to `on`.

This is **REQUIRED** to use Go Modules with "Buffalo" packages. The `auto` setting for `GO111MODULE` is **NOT** supported.

```bash
$ export GO111MODULE=on
```

## Working Outside of the `GOPATH`

In addition to repeatable builds, Go Modules, allows you to easily work outside of the `GOPATH`.

With `GO111MODULE=on` the `buffalo` command should work as it previously did _inside_ the `GOPATH`.

```bash
$ export GO111MODULE=on
$ buffalo new -h
```

## Working Inside the `GOPATH`

Because Go Modules are still experimental, and not complete, it is recommended to continue to work **INSIDE** the `GOPATH`. This will allow you to easily move between using, and not using modules.

When working inside the `GOPATH` you should continue to use `GOPATH` style module names.

#### Recommended

This style of module name works both inside, and outside, of the `GOPATH` easily. It also makes your projects work with `go get`.

```go
module github.com/markbates/coke
```

#### Not-Recommended

This style of module, can work inside of the `GOPATH`, but it is less flexible, although shorter, than the longer format module name.

```go
module coke
```

Regardless of which module name style you pick, you **MUST** be consistent within your application.

For example, if your module name is `coke` your actions package is `coke/actions`. If you module name is `github.com/markbates/coke` your actions package is `github.com/markbates/coke/actions`.

## FAQs

### I Get `invalid import`

When I run `buffalo build` I get strange errors like this when I run **outside** of my `GOPATH`:

```text
invalid import path: "D:/projects/testBuffalo/src/my-project/actions"
```

Make sure you have `GO111MODULE=on`. If you don't, Buffalo, tries to use your `GOPATH` to determine your package locations. Enable Go Modules support and try again.

### How Do I Migrate From Dep?

The `go mod init` tool can read your `Gopkg.toml` files and create a new `go.mod` for you. [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)

### How Do I Use The `development` Branch?

If you want to live on the "edge" and use the latest, bleeding edge, version of Buffalo you can tell Go Modules to get that version:

```bash
$ go get -u github.com/gobuffalo/buffalo@development
$ go mod tidy
```
