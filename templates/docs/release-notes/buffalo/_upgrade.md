## How to Upgrade

### Pre-built Binaries

The easiest solution is to download one of the pre-built binaries:

https://github.com/gobuffalo/buffalo/releases/tag/<%= to %>

### Using Go Get

```bash
$ go get -u github.com/gobuffalo/buffalo/buffalo
```

### From Source

```bash
$ go get github.com/gobuffalo/buffalo
$ cd $GOPATH/src/github.com/gobuffalo/buffalo
$ git checkout tags/<%= to %> -b <%= to %>
$ make install
```

---

Once you have an upgraded binary you can run the following command to attempt to upgrade your application from `<%= from %>` to `<%= to %>`.

```bash
$ buffalo fix
```

Note: While we have done our best to make this update command work well, please understand that it might not get you to a complete upgrade depending on your application and its complexities, but it will get you pretty close.
