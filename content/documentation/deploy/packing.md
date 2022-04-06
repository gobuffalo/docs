---
name: Packing
seoDescription: "Packing a Buffalo application"
seoKeywords: ["buffalo", "go", "golang", "pack", "binary", "build"]
weight: 1
aliases:
  - /documentation/deploy/building
  - /docs/building
  - /en/docs/building
---

# Packing

Now, your project is ready to be deployed. In this section, you will learn how to package a version of your app to deploy it on a server.

## The build Command

Buffalo features a command, `build`, that will build a **full binary** of your application including, but not limited to; assets, migrations, templates, etc. If you buy into the “Buffalo Way”, things just work. It's a wonderful experience. :)

```bash
$ buffalo build
```

```bash
Buffalo version {{< latestclirelease >}}

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

When the build finishes, you have a fresh baked binary in the `bin` folder. It will also have the **compilation time** and the **git commit SHA** burnt in, thus making the binaries “versioned”.

## Customize the Build

To get the list of available options, use the help command:

```bash
$ buffalo help build
```

```bash
Buffalo version {{< latestclirelease >}}

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

### Binary name / location

By default, your application will be built in the `bin` directory of your project, and the name of the executable will be the name you used to create the project with the `new` command.

You can change this default name by using the `-o` or `-output` flag:

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

In fact, you can change the target directory too:

```bash
$ # Put the app in my home directory, as "coke"
$ buffalo build -o ~/coke
```

```bash
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o ~/coke -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

### Extract Assets in a Zip File

By default, your whole app is packed into a single executable, assets included. In production setups, you may want to serve these assets with a proxy server (like Apache or NGINX), to lower the app load. You may even use a *CDN* to handle your assets.

Buffalo provides a way to extract compiled app assets into a single archive, using the `-e` or `-extract-assets` flag:

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

By default, the assets archive is put in the *bin* directory, but if you change the executable output directory with the `-o` flag, the assets will be put in the same directory.

```bash
$ ls -la bin
```

```bash
total 36280
drwxr-xr--@  4 markbates  staff   136B Apr  3 10:10 ./
drwxr-xr-x@ 20 markbates  staff   680B Apr  3 10:10 ../
-rwxr-xr-x@  1 markbates  staff    17M Apr  3 10:10 coke*
-rw-r--r--@  1 markbates  staff   691K Apr  3 10:10 coke-assets.zip
```

## Advanced Options

### Building “Static”/CGO Binaries

Building statically linked binaries that contain CGO, think SQLite3, can be tricky. By using the `--static` flag with `buffalo build`, the flags `--ldflags '-linkmode external -extldflags "-static"'` will be added to the `go build` command.

### Build Tags

When building a Buffalo binary using the `buffalo build` command, you can pass `--tags` and `--ldflags` to the built binary; just as you normally would when using the `go build` tools.

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

## Binary Commands

### Modes
Binaries, by default, run in `development` mode, which means all of the sub-commands will run in that mode as well. To change the mode, you must use the `GO_ENV` environment variable.

```bash
$ GO_ENV=production ./coke
```

### Available commands

Once a binary has been built, there are several sub-commands that can be run on that binary:

#### Default

The default command, if you just run the binary, will start the application.

#### migrate

The `migrate` sub-command will run the migrations for the application.

#### version

The `version` sub-command will output the version information for the binary, including the name, the git commit SHA used to build the binary, and the time the binary was built.

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task

The `task` sub-command runs tasks.

```bash
$ ./coke task greet

Hello World!
```

## Next Steps

* [Using a Proxy](/documentation/deploy/proxy) - Integrate your app with a server like NGINX.
* [Systemd Service](/documentation/deploy/systemd) - Run your app as a systemd service.
* [Cloud Providers](/documentation/deploy/cloud-providers) - Deploy your app on a cloud provider.
