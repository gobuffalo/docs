# Trouble Shooting

<%= title("buffalo new <project> fails looking for golang.org/x/tools/go/gcimporter") %>

This is caused by an outdated copy of the `github.com/motemen/gore` package. To fix simply update `gore`:

```text
$ go get -u github.com/motemen/gore
```

For information see [https://github.com/gobuffalo/buffalo/issues/108](https://github.com/gobuffalo/buffalo/issues/108) and [https://github.com/motemen/gore/issues/63](https://github.com/motemen/gore/issues/63).

<%= title("buffalo dev fails to start", {name: "dev-fails"}) %>

When starting `$ buffalo dev`, and you encounter this error:

`There was a problem starting the dev server: Unknown, Please review the troubleshooting docs`

This may be due to your system missing NodeJS/NPM, Ensure that Node/NPM is installed and is in your `$PATH`. If  Node/NPM are indeed in your `$PATH`, try renaming webpack.config.js.

If you are still having issues after attempting the steps above, please reach out to the community in the #buffalo channel on Gophers Slack.

<%= title("package context: unrecognized import path \"context\" (import path does not begin with hostname)", {name: "unrecognized-context"}) %>

When trying to install Buffalo `go get` returns this error:

`package context: unrecognized import path "context" (import path does not begin with hostname)`

This is due to an outdated version of Go. Buffalo requires Go 1.7 or higher. Please check your installation of Go and ensure you running the latest version.

<%= title("error \"unexpected directory layout:\" during \"go get\"") %>

Occasionally when running `go get` on Buffalo you will get the following error:

```
unexpected directory layout:
    import path: github.com/mattn/go-colorable
    dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
    expand dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
    separator: /
```

This issue has been reported previously the Go team, [https://github.com/golang/go/issues/17597](https://github.com/golang/go/issues/17597).

The best way to solve this problem is to run `go get` again, and it seems to fix itself.
