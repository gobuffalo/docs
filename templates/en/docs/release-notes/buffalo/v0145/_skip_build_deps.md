## Skip Build-time Dependencies

When running the `buffalo build` command Buffalo may need packages that are not part of the applications directly.

For example, in order have access to runtime information in your application, the [`github.com/gobuffalo/buffalo/runtime`](https://godoc.org/github.com/gobuffalo/buffalo/runtime) is used to provide versioning information.


There are times when you may not want/need this, for a variety of reasons.

The new `--skip-build-deps` flag allows you to disable the `buffalo` binary from trying to satisfy those dependencies automatically.

```bash
$ buffalo build --skip-builds-deps
```
