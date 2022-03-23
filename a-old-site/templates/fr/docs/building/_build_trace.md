```bash
$ buffalo build
```

```bash
Buffalo version <%= version %>

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/coke -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```