---
name: Writing Plugins
seoDescription: ""
seoKeywords: ["buffalo", "go", "golang", "cli", "plugins"]
weight: 4
aliases:
  - /docs/cli/plugins
  - /en/docs/cli/plugins
---

# Writing Custom Plugins

As the [Extending buffalo](/documentation/cli/extending) shows the Buffalo CLI can be extended by adding/removing plugins. On this section you will learn how to write your own plugins and add them to the Buffalo CLI.

## Anatomy of a Plugin

A plugin for the Buffalo CLI is a type that satisfies the [`plugin.Plugin`](TODO) interface.

```go
type Plugin interface {
  Name() string
}
```

Beyond the Plugin interface there are some other interfaces one can implement depending on the functionallity that we want our plugin to provide. For more example check the [base commands](/documentation/cli/commands) documentation, which lists the interfaces that the commands provide.

### Writing a Hello Command

In order to show the capabilities we will write a command that prints a message. To achieve this we will declare our SayHello command in the `cmd/buffalo/cmd/hello.go` file.

```go
// in cmd/buffalo/cmd/hello.go
package hello

import(
  "fmt"
)

type SayHello string

func (c SayHello) Name() string {
  return string(c)
}

func (c HelloCommand) Main(ctx context.Context, pwd string, args []string) error {
  fmt.Println("Hello, world!")

  return nil
}
```

And then add it to our `cmd/buffalo/main.go` CLI instance by using the `Add` API method.

```go
package main

import(
  "github.com/gobuffalo/cli/cmd/cli"
  "me/mymodule/cmd/buffalo/cmd/hello" // referencing the hello package
)

func main() {
  app := cli.NewWithDefaults()
  app.Add(hello.SayHello("my-hello-command"))

  app.Run()
}
```

To check that the command has been wired correctly we can run the plugins command, And we should see it listed in the terminal.

```bash
$ buffalo plugins

Plugins loaded (38):
  help                 [command]   Provides help for a given command, p.e. buffalo help list.
  plugins              [command]   List plugins loaded to the CLI.
  ...
  my-hello-command     [command]   (...) <<< 👀👀
  ...
```

And we can invoke our command with its name, Which invoke our `SayHello` "Main" method.

```bash   
$ buffalo my-hello-command
Hello, world!
```


