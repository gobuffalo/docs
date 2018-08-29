# Plugins

Plugins, introduced in `v0.9.1`, allow for 3rd party code to extend the `buffalo` command as well as its sub-commands.

<%= title("Plugins list") %>

Here is a list of plugins made by the Buffalo community:

| Plugin  | Description | Author | Docs / Examples |
|---------|-------------|--------|-----------------|
| [buffalo-goth](https://github.com/gobuffalo/buffalo-goth) | A generator to use with the [Goth](https://github.com/markbates/goth) third-party auth package. | [@markbates](https://github.com/markbates) |   |
| [buffalo-bootstrap](https://github.com/markbates/buffalo-bootstrap) | A plugin for generating Bootstrap stuff. | [@markbates](https://github.com/markbates) |   |
| [buffalo-azure](https://github.com/Azure/buffalo-azure) | A plugin for deploying to and generating for Microsoft Azure. | [@Microsoft](https://open.microsoft.com) | |
| [buffalo-heroku](https://github.com/markbates/buffalo-heroku) | A plugin for deploying docker to heroku.  | [@markbates](https://github.com/markbates) |   |
| [buffalo-ocean](https://github.com/wolves/buffalo-ocean) | A plugin for deploying docker to Digital Ocean.  | [@wolves](https://github.com/wolves) | [Demo](https://blog.wolvesdesign.com/post/buffalo-ocean-walkthrough/) |

A plugin is missing from this list? Open [a PR](https://github.com/gobuffalo/gobuffalo/pulls) to add it!

<%= title("Writing a Plugin") %>

There are only a few details to know when implementing a plugin for Buffalo.

* Plugins must be named in the format of `buffalo-&lt;plugin-name>`. For example, `buffalo-myplug`.
* Plugins must be executable and must be available in one of two places:
  * in the `$PATH`
  * in the `plugins/` folder of your Buffalo application
* Plugins must implement an `available` command that prints a JSON response listing the available commands.

<%= title("Plugin: Case Study") %>

Using the [https://github.com/markbates/buffalo-heroku](https://github.com/markbates/buffalo-heroku) plugin let's examine how it's fits into the `buffalo` command.

Since this plugin was written in Go, it can be installed using the Go tool.

```bash
$ go get -v github.com/markbates/buffalo-heroku
```

After installation it will be in the `$GOPATH/bin` folder. Since most Go installations have `$GOPATH/bin` added to the `$PATH`, this plugin will be found by the `buffalo` command.

---

If we were to run `buffalo-heroku available` we would get a JSON response printed out that lists the available commands for this plugin.

```json
[
  {
  "buffalo_command" : "root",
  "description" : "Tools for deploying Buffalo to Heroku",
  "aliases" : [ "h" ],
  "name" : "heroku"
  }
]
```

* `buffalo_command`: This is the Buffalo command you want to nest your plugin commands underneath. In this case that command is `root`, meaning it will show up when you run `buffalo --help`. Other examples are `generate`, `task`, `dev`, etc...
* `description`: This should be a small (one sentence) description of what your plugin does.
* `aliases`: This is an **optional** list of aliases for your command.
* `name`: This is the name the plugin will be run as. For example, `buffalo heroku`.

With all of this in place when we run `buffalo --help` we should see the plugin listed with the rest of the commands.

```bash
$ buffalo --help

Helps you build your Buffalo applications that much easier!

Usage:
  buffalo [command]

Available Commands:
  // ...
  help        Help about any command
  heroku      [PLUGIN] Tools for deploying Buffalo to Heroku
  info        Prints off diagnostic information useful for debugging.
  // ...

Flags:
  -h, --help   help for buffalo

  Use "buffalo [command] --help" for more information about a command.
```

<%= title("Writing Non-Go Plugins") %>

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

---

```bash
$ buffalo --help

Helps you build your Buffalo applications that much easier!

Usage:
  buffalo [command]

Available Commands:
  // ...
  hello       [PLUGIN] says hello to you
  help        Help about any command
  heroku      [PLUGIN] Tools for deploying Buffalo to Heroku
  // ...

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

```bash
$ buffalo hello

Hi there!
```
