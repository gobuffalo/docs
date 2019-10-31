## Support `-timeout` in Buffalo Test

When running tests with `buffalo test` you can now pass the `-timeout` flag, as with `go test`, to limit how long tests can run.

```go
$ buffalo test -timeout 3s
```

* [https://github.com/gobuffalo/buffalo/pull/1809](https://github.com/gobuffalo/buffalo/pull/1809)
