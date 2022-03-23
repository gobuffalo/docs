```bash
$ buffalo help build
```

```bash
Buffalo version <%= version %>

Builds a Buffalo binary, including bundling of assets (packr & webpack)

Usage:
  buffalo build [flags]

Aliases:
  build, b, bill

Flags:
  -c, --compress         compress static files in the binary (default true)
  -e, --extract-assets   extract the assets and put them in a distinct archive
  -h, --help             help for build
      --ldflags string   set any ldflags to be passed to the go build
  -o, --output string    set the name of the binary (default "bin/coke")
  -s, --static           build a static binary using  --ldflags '-linkmode external -extldflags "-static"' (USE FOR CGO)
  -t, --tags string      compile with specific build tags
```