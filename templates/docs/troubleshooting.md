# Trouble Shooting

{{#panel title="buffalo new <project> fails looking for golang.org/x/tools/go/gcimporter"}}

This is caused by an outdated copy of the `github.com/motemen/gore` package. To fix simply update `gore`:

```text
$ go go get -u github.com/motemen/gore
```

For information see [https://github.com/gobuffalo/buffalo/issues/108](https://github.com/gobuffalo/buffalo/issues/108) and [https://github.com/motemen/gore/issues/63](https://github.com/motemen/gore/issues/63).
{{/panel}}
