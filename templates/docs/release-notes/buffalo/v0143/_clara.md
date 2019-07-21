## Improved Buffalo Info Output

The `buffalo info` command can be useful for helper developers gather information about their application and environment to help, or help others, diagnosis issues that may arise.

In previous versions of Buffalo the output would look something like the follow:

```plain
### Buffalo Version
v0.14.0

### App Information
Pwd=$GOPATH/src/github.com/markbates/coke
Root=$GOPATH/src/github.com/markbates/coke
GoPath=$GOPATH
PackagePkg=github.com/markbates/coke
ActionsPkg=github.com/markbates/coke/actions
ModelsPkg=github.com/markbates/coke/models
GriftsPkg=github.com/markbates/coke/grifts
WithModules=true
Name=coke
Bin=bin/coke
VCS=git
WithPop=true
WithSQLite=false
WithDep=false
WithWebpack=true
WithNodeJs=true
WithYarn=true
WithDocker=true
WithGrifts=true
AsWeb=true
AsAPI=false
PackageJSON={map[]}

### Go Version
go version go1.12.3 darwin/amd64

### Go Env
GOARCH="amd64"
GOBIN=""
GOCACHE="$HOME/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="$GOPATH"
GOPROXY=""
GORACE=""
GOROOT="$GOROOT"
GOTMPDIR=""
GOTOOLDIR="$GOROOT/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="$GOPATH/src/github.com/markbates/coke/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/zj/ktv0trrj4l79dfq0dkm1b6d40000gn/T/go-build329153793=/tmp/go-build -gno-record-gcc-switches -fno-common"

### Node Version
v10.11.0

### NPM Version
6.4.1

### Yarn Version
1.10.1

### PostgreSQL Version
pg_ctl (PostgreSQL) 10.5

### MySQL Version
mysql  Ver 8.0.12 for osx10.13 on x86_64 (Homebrew)

### SQLite Version
3.24.0 2018-06-04 14:10:15 95fbac39baaab1c3a84fdfc82ccb7f42398b2e92f18a2a57bce1d4a713cbaapl

### Dep Version
could not find a Gopkg.toml file

### Dep Status
could not find a Gopkg.toml file

### go.mod
module github.com/markbates/coke

go 1.12

require (
	cloud.google.com/go v0.36.0 // indirect
	github.com/codegangsta/negroni v1.0.0 // indirect
	github.com/gobuffalo/buffalo v0.14.7-beta.2
	github.com/gobuffalo/buffalo-docker v1.0.7 // indirect
	github.com/gobuffalo/buffalo-pop v1.16.0
	github.com/gobuffalo/envy v1.7.0
	github.com/gobuffalo/mw-csrf v0.0.0-20190129204204-25460a055517
	github.com/gobuffalo/mw-forcessl v0.0.0-20190224202501-6d1ef7ffb276
	github.com/gobuffalo/mw-i18n v0.0.0-20190224203426-337de00e4c33
	github.com/gobuffalo/mw-paramlogger v0.0.0-20190224201358-0d45762ab655
	github.com/gobuffalo/packr v1.30.1
	github.com/gobuffalo/packr/v2 v2.5.2
	github.com/gobuffalo/pop v4.11.2+incompatible
	github.com/gobuffalo/suite v2.8.1+incompatible
	github.com/gobuffalo/x v0.0.0-20190224155809-6bb134105960 // indirect
	github.com/markbates/grift v1.1.0
	github.com/unrolled/secure v1.0.0
)
```

While all of that information is helpful to members of the core team, few others understand it and offers little of value to them.

The [Clara](https://github.com/gobuffalo/clara) tool was built as a way of helping developers check their environment for common setup issues for both Go and Buffalo. It then provides simple troubleshooting tips and links.

The `buffalo info` command now uses Clara, as well as custom checks, to, hopefully, help developers with issues.

```plain
-> Go: Checking installation
✓ The `go` executable was found on your system at: $GOROOT/bin/go

-> Go: Checking minimum version requirements
✓ Your version of Go, 1.12.6, meets the minimum requirements.

-> Go: Checking GOPATH
✓ You appear to be operating inside of your GOPATH.

-> Go: Checking Package Management
⚠ You do not appear to be using a package management system.

It is strongly suggested that you use one of the following package management systems:

* Go Modules (Recommended) - https://gobuffalo.io/en/docs/gomods
* Dep - https://github.com/golang/dep

For help setting up your Go environment please follow the instructions for you platform at:

https://www.gopherguides.com/courses/preparing-your-environment-for-go-development

-> Go: Checking PATH
✓ Your PATH contains $GOPATH/bin.

-> Node: Checking installation
✓ The `node` executable was found on your system at: /usr/local/bin/node

-> Node: Checking minimum version requirements
✓ Your version of Node, v10.11.0, meets the minimum requirements.

-> NPM: Checking installation
✓ The `npm` executable was found on your system at: /usr/local/bin/npm

-> NPM: Checking minimum version requirements
✓ Your version of NPM, 6.4.1, meets the minimum requirements.

-> Yarn: Checking installation
✓ The `yarnpkg` executable was found on your system at: /usr/local/bin/yarnpkg

-> Yarn: Checking minimum version requirements
✘ Your version of Yarn, 1.10.1, does not meet the minimum requirements.

Minimum versions of Yarn are:

* >=1.12

For help setting up your Yarn environment please follow the instructions for you platform at:

https://yarnpkg.com/en/docs/install

-> PostgreSQL: Checking installation
✓ The `postgres` executable was found on your system at: /Applications/Postgres.app/Contents/Versions/latest/bin/postgres

-> PostgreSQL: Checking minimum version requirements
✓ Your version of PostgreSQL, 10.5, meets the minimum requirements.

-> MySQL: Checking installation
✓ The `mysql` executable was found on your system at: /usr/local/bin/mysql

-> MySQL: Checking minimum version requirements
✓ Your version of MySQL, 8.0.12, meets the minimum requirements.

-> SQLite3: Checking installation
✓ The `sqlite3` executable was found on your system at: /usr/bin/sqlite3

-> SQLite3: Checking minimum version requirements
✓ Your version of SQLite3, 3.24.0, meets the minimum requirements.

-> Cockroach: Checking installation
✓ The `cockroach` executable was found on your system at: /usr/local/bin/cockroach

-> Cockroach: Checking minimum version requirements
✓ Your version of Cockroach, 2.0.5, meets the minimum requirements.

-> Buffalo: Checking installation
✓ The `buffalo` executable was found on your system at: $GOPATH/bin/buffalo

-> Buffalo: Checking minimum version requirements
✓ Your version of Buffalo, v0.14.7, meets the minimum requirements.

-> Buffalo: Application Details
Pwd         $GOPATH/src/github.com/markbates/coke
Root        $GOPATH/src/github.com/markbates/coke
GoPath      $GOPATH
PackagePkg  github.com/markbates/coke
ActionsPkg  github.com/markbates/coke/actions
ModelsPkg   github.com/markbates/coke/models
GriftsPkg   github.com/markbates/coke/grifts
WithModules false
Name        coke
Bin         bin/coke
VCS         git
WithPop     true
WithSQLite  false
WithDep     false
WithWebpack true
WithNodeJs  true
WithYarn    true
WithDocker  true
WithGrifts  true
AsWeb       true
AsAPI       false
InApp       true
PackageJSON {map[]}

-> Buffalo: config/buffalo-app.toml
name = "coke"
bin = "bin/coke"
vcs = "git"
with_pop = true
with_sqlite = false
with_dep = false
with_webpack = true
with_nodejs = true
with_yarn = true
with_docker = true
with_grifts = true
as_web = true
as_api = false

-> Buffalo: config/buffalo-plugins.toml
[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"

-> Buffalo: go.mod
module github.com/markbates/coke

go 1.12

require (
	cloud.google.com/go v0.36.0 // indirect
	github.com/codegangsta/negroni v1.0.0 // indirect
	github.com/gobuffalo/buffalo v0.14.7-beta.2
	github.com/gobuffalo/buffalo-docker v1.0.7 // indirect
	github.com/gobuffalo/buffalo-pop v1.16.0
	github.com/gobuffalo/envy v1.7.0
	github.com/gobuffalo/mw-csrf v0.0.0-20190129204204-25460a055517
	github.com/gobuffalo/mw-forcessl v0.0.0-20190224202501-6d1ef7ffb276
	github.com/gobuffalo/mw-i18n v0.0.0-20190224203426-337de00e4c33
	github.com/gobuffalo/mw-paramlogger v0.0.0-20190224201358-0d45762ab655
	github.com/gobuffalo/packr v1.30.1
	github.com/gobuffalo/packr/v2 v2.5.2
	github.com/gobuffalo/pop v4.11.2+incompatible
	github.com/gobuffalo/suite v2.8.1+incompatible
	github.com/gobuffalo/x v0.0.0-20190224155809-6bb134105960 // indirect
	github.com/markbates/grift v1.1.0
	github.com/unrolled/secure v1.0.0
)
```