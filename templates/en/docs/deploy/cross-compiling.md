<% seoDescription("Cross-compiling a Buffalo application") %>
<% seoKeywords(["buffalo", "go", "golang", "cross-compile", "linux", "arm", "windows", "mac"]) %>

<%= h1("Cross-compiling a Buffalo application") %>

Just like another Go application, you can cross-compile a Buffalo application. This means it's possible to develop your app on a Mac, and compile it for a Linux target on your Mac.

## GOOS and GOARCH

The Go toolchain supports cross-compilation out of the box. You just need to provide the `GOOS` and `GOARCH` env variables.
* `GOOS` sets the target OS (e.g. linux, windows, etc.)
* `GOARCH` sets the target CPU architecture (e.g. amd64, 386, etc.)

You can find the list of supported targets here: https://golang.org/doc/install/source#environment

## Examples

### Build for AMD64 Linux

```go
$ GOOS=linux GOARCH=amd64 buffalo build
```

### Build for ARM64 Linux

```go
$ GOOS=linux GOARCH=arm64 buffalo build
```

### Build for i386 Windows

```go
$ GOOS=windows GOARCH=386 buffalo build
```