---
name: Plugins
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

<%= exampleDir("docs/plugins/_example/standard") %>


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

