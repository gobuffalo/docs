<% seoDescription("How to install the Buffalo framework") %>
<% seoKeywords(["buffalo", "go", "golang", "installation", "framework", "web", "mac", "windows", "linux"]) %>

<%= h1("Install Buffalo") %>

In this chapter, you'll learn how to install Buffalo, either from pre-built binaries or from source.

Buffalo provides **two major components**:
* The `buffalo` tool, a powerful toolbox to help you develop in a fast and efficient way.
* The buffalo framework, a collection of pieces to construct your app.

Buffalo is currently available and tested on the following platforms:
* GNU/Linux
* Mac OSX
* Windows

## Requirements

Before installing make sure you have the required dependencies installed:

* [A working Go environment](http://gopherguides.com/before-you-come-to-class)
* [A configured `$PATH` environment variable that includes `$GOPATH/bin`.](https://golang.org/doc/code.html#GOPATH)
* [Go](https://golang.org) version `<%= goMinVersion %>`.

##### Frontend Requirements

The following requirements are optional. You don't need them if you want to build an API or if you prefer to build your app in an old-fashioned way.

* [node](https://github.com/nodejs/node) version `8` or greater
* either [yarn](https://yarnpkg.com/en/) or [npm](https://github.com/npm/npm) for the [asset pipeline](/en/docs/assets) built upon [webpack](https://github.com/webpack/webpack).

##### Database Specific Requirements

Again, if you don't need a database, you won't need these.

* **SQLite 3**: GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

## Installation from a release archive - 64 bits

<%= note() { %>
The release packages contain Buffalo without SQLite support.
<% } %>

Since `v0.10.3`, pre-compiled archives are provided with each release. If you don't need the latest cutting-edge version, you'll probably prefer to install this version.

### GNU / Linux

```bash
$ wget https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_Linux_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Linux_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### MacOS

```bash
$ curl -OL https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
# or if you have ~/bin folder setup in the environment PATH variable
$ mv buffalo ~/bin/buffalo
```

## Scoop (Windows)
Buffalo can be installed using the [Scoop](http://scoop.sh/) package manager:

```powershell
PS C:\> scoop install buffalo
```

## Chocolatey (Windows)
Buffalo can be installed using the [Chocolatey](https://chocolatey.org/packages/buffalo) package manager. Versions on Chocolatey are published with a potential delay and must go through moderation before they are available:

```powershell
PS C:\> choco install buffalo
```

## Homebrew (macOS)

On macOS, you can also install Buffalo with [Homebrew](https://brew.sh/). After you have Homebrew [installed](https://docs.brew.sh/Installation), you can easily install Buffalo:

```bash
brew install gobuffalo/tap/buffalo
```

## GoFish (Cross-Platforms)

[GoFish](https://gofi.sh/index.html) is a cross-platform systems package manager, that works across Windows, MacOSX and Linux.

After you have GoFish [installed](https://gofi.sh/index.html#install), you can very simply install Buffalo:

```bash
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo <%= version %>: installed in 3.223672926s
```

## Custom installation **with** SQLite3 Support

**SQLite 3** requires a GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) to compile. You **must** have a GCC installed **first** before installing Buffalo.

To install Buffalo, ensure your `GOPATH` is defined, then:

```bash
$ go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo
```

**Windows Users**: Follow the installation guide at [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) to install a GCC for Windows 10. Alternatively, GCC can be installed with the [Scoop](http://scoop.sh/) package manager:

```powershell
PS C:\> scoop install gcc
```

<%= note() { %>
These instructions can also be used for upgrading to a newer version of Buffalo.
<% } %>

## Custom installation **without** SQLite3 Support

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
These instructions can also be used for upgrading to a newer version of Buffalo.
<% } %>

## Verify your installation

You can check if your installation is working, by executing the `buffalo` command in a terminal/command prompt:

```bash
$ buffalo help
Build Buffalo applications with ease

Usage:
  buffalo [command]

Available Commands:
  build       Build the application binary, including bundling of assets (packr & webpack)
  destroy     Destroy generated components
  dev         Run the Buffalo app in 'development' mode
  fix         Attempt to fix a Buffalo application's API to match version v0.16.17
  generate    Generate application components
  help        Help about any command
  info        Print diagnostic information (useful for debugging)
  new         Creates a new Buffalo application
  plugins     tools for working with buffalo plugins
  routes      Print all defined routes
  setup       Setup a newly created, or recently checked out application.
  task        Run grift tasks
  test        Run the tests for the Buffalo app. Use --force-migrations to skip schema load.
  version     Print the version information

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

If you have a similar output, your Buffalo toolbox is ready to work!

## Next Steps

* [Tooling Integration](/en/docs/getting-started/integrations) - Work with Buffalo, using existing tools.
* [Generate a New Project](/en/docs/getting-started/new-project) - Create your first Buffalo project!
