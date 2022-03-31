```bash
$ buffalo build -e
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> build assets archive
--> disable self assets handling
--> running go build -v -o bin/coke -ldflags -X main.version="2017-04-02T08:45:58+02:00" -X main.buildTime="2017-04-02T08:45:58+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```