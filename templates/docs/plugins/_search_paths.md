## How Does Buffalo Find Plugins?

Buffalo plugins have a set of rules that must be followed for them to be consider, by Buffalo, as a plugin.

* Plugins must be named in the format of `buffalo-&lt;plugin-name>`. For example, `buffalo-myplug`.
* Plugins must be executable and must be available in one of the following places:
  * in the `$BUFFALO_PLUGIN_PATH`
  * if not set, `$GOPATH/bin`, is tried
  * in the `./plugins` folder of your Buffalo application
* Plugins must implement an `available` command that prints a JSON response listing the available commands.

The `buffalo plugins list` command will print a table of plugins that Buffalo sees as "available" to you.

## With Configuration

<%= sinceVersion("1.1.0", {pkg: "github.com/gobuffalo/buffalo-plugins"}) %>

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
