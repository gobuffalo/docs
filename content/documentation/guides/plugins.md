---
name: Plugins
weight: 20
aliases:
  - /docs/plugins
  - /en/docs/plugins
---

# Plugins

{{< since "0.9.1" >}}

Plugins allow for 3rd party code to extend the `buffalo` command as well as its sub-commands.

## Installing the buffalo-plugins Plugin

```bash
$ go get -u -v github.com/gobuffalo/buffalo-plugins
```

## Finding Available Plugins

A full list of Plugins can be found at [https://toolkit.gobuffalo.io/tools?topic=plugin](https://toolkit.gobuffalo.io/tools?topic=plugin).

To get your project listed on the Buffalo Toolkit you must tag your project on GitHub with `gobuffalo`.

There are a few more tags that you can use that will help the Buffalo Toolkit better categorize your project. You can add as many of this tags to your project as is suitable. Please try to refrain from using more than just a few tags.

* `plugin` - Plugins
* `generator` - Generators
* `middleware` - Middleware
* `pop` - Pop/Soda
* `templating` - Templating
* `grifts` - Grift Tasks
* `deployment` - Deployment
* `testing` - Testing
* `example` - Example Apps
* `worker` - Workers/Adapters
* `webpack` - Webpack/Front-End

Any other tags will still be indexed and searchable, but the tool may not show in the "known" categories section.


## How Does Buffalo Find Plugins?

Buffalo plugins have a set of rules that must be followed for them to be consider, by Buffalo, as a plugin.

* Plugins must be named in the format of `buffalo-<plugin-name>`. For example, `buffalo-myplugin`.
* Plugins must be executable and must be available in one of the following places:
  * in the `$BUFFALO_PLUGIN_PATH`
  * if not set, `$GOPATH/bin`, is tried
  * in the `./plugins` folder of your Buffalo application
* Plugins must implement an `available` command that prints a JSON response listing the available commands.

The `buffalo plugins list` command will print a table of plugins that Buffalo sees as "available" to you.

## With Configuration

{{< since "1.1.0" >}}

When a `config/buffalo-plugins.toml` file is present, Buffalo will scope the list of plugins that are "available" to those listed in the configuration file.

```bash
$ cat config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"
```

```bash
$ buffalo plugins list

Bin         |Command               |Description
---         |---                   |---
buffalo-pop |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop |buffalo destroy model |Destroys model files.
buffalo-pop |buffalo pop           |A tasty treat for all your database needs
```

## Without Configuration

Without a configuration file, Buffalo will try to aforementioned paths to find any, and all Buffalo plugins installed on the users system.

```bash
$ buffalo plugins list

Bin              |Command                    |Description
---              |---                        |---
buffalo-auth     |buffalo generate auth      |Generates a full auth implementation
buffalo-pop      |buffalo db                 |[DEPRECATED] please use `buffalo pop` instead.
buffalo-goth     |buffalo generate goth-auth |Generates a full auth implementation use Goth
buffalo-goth     |buffalo generate goth      |generates a actions/auth.go file configured to the specified providers.
buffalo-heroku   |buffalo heroku             |helps with heroku setup and deployment for buffalo applications
buffalo-plugins  |buffalo events listen      |listens to github.com/gobuffalo/events
buffalo-pop      |buffalo destroy model      |Destroys model files.
buffalo-plugins  |buffalo generate plugin    |generates a new buffalo plugin
buffalo-plugins  |buffalo plugins            |tools for working with buffalo plugins
buffalo-pop      |buffalo pop                |A tasty treat for all your database needs
buffalo-trash    |buffalo trash              |destroys and recreates a buffalo app
buffalo-upgradex |buffalo upgradex           |updates Buffalo and/or Pop/Soda as well as your app
```

## Installing Plugins

{{< since "1.1.0" >}}

To add support for the plugin manager, one can either manually edit `./config/buffalo-plugins.toml` or let `buffalo plugins install` create it for you.

{{< codetabs >}}
{{< tab "Install command" >}}
```bash
// $ buffalo plugins install

go get github.com/gobuffalo/buffalo-pop
./config/buffalo-plugins.toml
```
{{< /tab >}}
{{< tab "Config file" >}}
``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"
```
{{< /tab >}}

{{< tab "Resulting plugin list" >}}
```bash
// $ buffalo plugins list

Bin         |Command               |Description
---         |---                   |---
buffalo-pop |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop |buffalo destroy model |Destroys model files.
buffalo-pop |buffalo pop           |A tasty treat for all your database needs
```
{{< /tab >}}
{{< /codetabs>}}


The `buffalo-pop` plugin was automatically added because the application in this example is a Buffalo application that uses Pop.

New plugins can be install in bulk with the `install` command

{{< codetabs >}}
{{< tab "Bulk Install command" >}}
```bash
// $ buffalo plugins install
$ buffalo plugins install github.com/markbates/buffalo-trash github.com/gobuffalo/buffalo-heroku

go get github.com/gobuffalo/buffalo-heroku
go get github.com/gobuffalo/buffalo-pop
go get github.com/markbates/buffalo-trash
./config/buffalo-plugins.toml
```
{{< /tab >}}
{{< tab "Config file" >}}
``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-heroku"
  go_get = "github.com/gobuffalo/buffalo-heroku"

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"

[[plugin]]
  binary = "buffalo-trash"
  go_get = "github.com/markbates/buffalo-trash"
```
{{< /tab >}}

{{< tab "Resulting plugin list" >}}
```bash
// $ buffalo plugins list

Bin            |Command               |Description
---            |---                   |---
buffalo-pop    |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-heroku |buffalo heroku        |helps with heroku setup and deployment for buffalo applications
buffalo-pop    |buffalo destroy model |Destroys model files.
buffalo-pop    |buffalo pop           |A tasty treat for all your database needs
buffalo-trash  |buffalo trash         |destroys and recreates a buffalo app
```
{{< /tab >}}
{{< /codetabs>}}


## Removing Plugins

{{< since "1.1.0" >}}

Plugins can be removed with the `remove` command. This only removes them from the config file, not from the users system.

{{< codetabs >}}
{{< tab "Remove command" >}}
```bash
// $ buffalo plugins remove
$ buffalo plugins remove github.com/gobuffalo/buffalo-heroku

./config/buffalo-plugins.toml
```
{{< /tab >}}
{{< tab "Config file" >}}
``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"

[[plugin]]
  binary = "buffalo-trash"
  go_get = "github.com/markbates/buffalo-trash"
```
{{< /tab >}}

{{< tab "Resulting plugin list" >}}
```bash
// $ buffalo plugins list

Bin           |Command               |Description
---           |---                   |---
buffalo-pop   |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop   |buffalo destroy model |Destroys model files.
buffalo-pop   |buffalo pop           |A tasty treat for all your database needs
buffalo-trash |buffalo trash         |destroys and recreates a buffalo app
```
<% } %>
{{< /tab >}}
{{< /codetabs>}}

## Writing a Plugin

First, you must understand [how Buffalo finds plugins](#how-does-buffalo-find-plugins), before you can successfully write one.

The `buffalo-plugins` plugin adds a new generator to `buffalo generate` to help you build a new plugin quickly

```bash
$ buffalo generate plugin -h

buffalo generate plugin github.com/foo/buffalo-bar

Usage:
  buffalo-plugins plugin [flags]

Flags:
  -a, --author string       author's name
  -d, --dry-run             run the generator without creating files or running commands
  -f, --force               will delete the target directory if it exists
  -h, --help                help for plugin
  -l, --license string      choose a license from: [agpl, isc, lgpl-v2.1, mozilla, no-license, artistic, bsd, eclipse, lgpl-v3, mit, apache, bsd-3-clause, unlicense, cc0, gpl-v2, gpl-v3] (default "mit")
  -s, --short-name string   a 'short' name for the package
      --with-gen            creates a generator plugin
```

{{< codetabs>}}
{{< tab "LICENSE" >}}
```text
The MIT License (MIT)

Copyright (c) 2018 Mark Bates

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
{{< /tab >}}
{{< tab "Makefile" >}}
```makefile
TAGS ?= "sqlite"
GO_BIN ?= go

install: deps
	packr
	$(GO_BIN) install -tags ${TAGS} -v ./.

deps:
	$(GO_BIN) get github.com/gobuffalo/release
	$(GO_BIN) get github.com/gobuffalo/packr/packr
	$(GO_BIN) get -tags ${TAGS} -t ./...
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

build:
	packr
	$(GO_BIN) build -v .

test:
	packr
	$(GO_BIN) test -tags ${TAGS} ./...

ci-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

lint:
	gometalinter --vendor ./... --deadline=1m --skip=internal

update:
	$(GO_BIN) get -u -tags ${TAGS}
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
	packr
	make test
	make install
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

release-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

release:
	release -y -f bar/version.go
```
{{< /tab >}}
{{< tab "main.go" >}}
```go
package main

import "github.com/foo/buffalo-bar/cmd"

func main() {
	cmd.Execute()
}
```
{{< /tab >}}
{{< tab "bar/version.go" >}}
```go
package bar

const Version = "v0.0.1"
```
{{< /tab >}}
{{< tab "cmd/available.go" >}}
```go
package cmd

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo-plugins/plugins"
	"github.com/spf13/cobra"
)

// availableCmd represents the available command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "a list of available buffalo plugins",
	RunE: func(cmd *cobra.Command, args []string) error {
		plugs := plugins.Commands{
			// {Name: "bar", UseCommand: "generate", BuffaloCommand: "generate", Description: generateCmd.Short, Aliases: generateCmd.Aliases},
		}
		return json.NewEncoder(os.Stdout).Encode(plugs)
	},
}

func init() {
	rootCmd.AddCommand(availableCmd)
}
```
{{< /tab >}}
{{< tab "cmd/bar.go" >}}
```go
package cmd

import (
	"github.com/spf13/cobra"
)

// barCmd represents the bar command
var barCmd = &cobra.Command{
	Use:   "bar",
	Short: "description about this plugin",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(barCmd)
}
```
{{< /tab >}}
{{< tab "cmd/root.go" >}}
```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "buffalo-bar",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
{{< /tab >}}
{{< tab "cmd/version.go" >}}
```go
package cmd

import (
	"fmt"

	"github.com/foo/buffalo-bar/bar"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of bar",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("bar", bar.Version)
		return nil
	},
}

func init() {
	barCmd.AddCommand(versionCmd)
}
```
{{< /tab >}}
{{< /codetabs>}}



## Writing Non-Go Plugins

Plugins do not need to be written in Go. They can be written in any language you would like, as long as they comply with the rules above.

For example, we can write the following plugin using Ruby:

```ruby
#!/usr/bin/env ruby
# ./plugins/buffalo-hello.rb

require 'json'

command = ARGV[0]

case command
when 'available'
  puts JSON.generate([{ name: 'hello', buffalo_command: 'root', description: 'says hello to you' }])
when 'hello'
  puts 'Hi there!'

end
```

To activate the plugin we need to add the file as `buffalo-hello.rb` to somewhere in the `$PATH` or in a directory called `plugins/` inside of a Buffalo application.

Finally the file needs to be made executable. On a Mac/Linux it can be done with `chmod +x buffalo-hello.rb`.

```bash
$ buffalo plugins list

Bin              |Command                    |Description
---              |---                        |---
buffalo-auth     |buffalo generate auth      |Generates a full auth implementation
buffalo-pop      |buffalo db                 |[DEPRECATED] please use `buffalo pop` instead.
buffalo-goth     |buffalo generate goth-auth |Generates a full auth implementation use Goth
buffalo-goth     |buffalo generate goth      |generates a actions/auth.go file configured to the specified providers.
buffalo-hello.rb |buffalo hello              |says hello to you
buffalo-heroku   |buffalo heroku             |helps with heroku setup and deployment for buffalo applications
buffalo-plugins  |buffalo events listen      |listens to github.com/gobuffalo/events
buffalo-pop      |buffalo destroy model      |Destroys model files.
buffalo-plugins  |buffalo generate plugin    |generates a new buffalo plugin
buffalo-plugins  |buffalo plugins            |tools for working with buffalo plugins
buffalo-pop      |buffalo pop                |A tasty treat for all your database needs
buffalo-trash    |buffalo trash              |destroys and recreates a buffalo app
buffalo-upgradex |buffalo upgradex           |updates Buffalo and/or Pop/Soda as well as your app
```

