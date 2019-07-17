### Clara

The Go/Buffalo tool for checking a system environment for Go and/or Buffalo development, [https://github.com/gobuffalo/clara](https://github.com/gobuffalo/clara), has been added to the `buffalo info` command to help users understand/diagnosis common environmental issues.

In addition, the cleaner, and more expressive output of `buffalo info` will be very helpful to those helping to assist on issues.

```plain
-> Go: Checking installation
✓ The `go` executable was found on your system at: $GOROOT/bin/go

-> Go: Checking minimum version requirements
✓ Your version of Go, 1.12.6, meets the minimum requirements.

-> Go: Checking GOPATH
✓ You are using Go Modules, so no need to worry about the GOPATH.

-> Go: Checking Package Management
✓ You are using Go Modules (`go`) for package management.

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
WithModules true
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
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c // indirect
	github.com/codegangsta/negroni v1.0.0 // indirect
	github.com/gobuffalo/buffalo v0.14.7-beta.2
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
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/markbates/grift v1.1.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
	github.com/unrolled/secure v1.0.0
)
```
