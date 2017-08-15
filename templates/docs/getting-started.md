<%= h1("Getting Started") %>

<%= title("Installation") %>

Before installing make sure you have the required dependencies installed:

* [Guide to Setting a Go Developer Environment](http://gopherguides.com/before-you-come-to-class)
* [Go](https://golang.org) version `1.8.1` or greater.
* GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).
* [OPTIONAL] [node](https://github.com/nodejs/node) and [npm](https://github.com/npm/npm) for the [asset pipeline](/docs/assets) built upon [webpack](https://github.com/webpack/webpack).

### Basic Installation:

These instructions can also be used for upgrading to a newer version of Buffalo.

```bash
$ go get -u -v github.com/gobuffalo/buffalo/...
```

<%= title("Generating a New Project") %>

Buffalo aims to make building new web applications in Go as quick and simple as possible, and what could be more simple than a *new application* generator? Start by going to your `$GOPATH` and create your new application!

```bash
$ cd $GOPATH/src/github.com/$USER/
$ # Make sure $GOPATH/bin is in your $PATH, then:
$ buffalo new <name>
```

That will generate a whole new Buffalo application that is ready to go. It'll even run `go get` for you to make sure you have all of the necessary dependencies needed to run your application.

<%= partial("docs/getting-started/new.md") %>

To see a list of available flags for the `new` command, just check out its help.

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci] (default "none")
      --db-type string       specify the type of database you want to use [postgres, mysql, sqlite3] (default "postgres")
      --docker string        specify the type of Docker file to generate [none, multi, standard] (default "multi")
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            skip to use npm as the asset package manager
  -v, --verbose              verbosely print out the go get/install commands
      --with-dep             adds github.com/golang/dep to your app
```

Note: by default, Buffalo generates a database.yml targeted for PostgreSQL. If you wish to change this behavior, you can pass in a `--db-type` flag into the `new` command.

```bash
$ buffalo new coke --db-type sqlite3
```

If your app doesn't need a database, or if you want to handle it by yourself, you can use the `--skip-pop` flag.

```bash
$ buffalo new coke --skip-pop
```

<%= partial("docs/dev.md") %>

<%= title("Building Your Application", {name: "building"}) %>

Buffalo features a command, `build`, that will build a full binary of your application including, but not limited to; assets, migrations, templates, etc... If you buy into the "Buffalo Way" things just work. It's a wonderful experience. :)

```bash
$ buffalo build
```

```bash
Buffalo version <%= version %>

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/gobuffalo -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

See [Building](/docs/building) for more options on the `build` command.
