```bash
$ buffalo build -o bin/cookies
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/cookies -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```