---
name: Extending Buffalo
seoDescription: ""
seoKeywords: ["buffalo", "go", "golang", "cli", "plugins"]
weight: 3
aliases:
  - /docs/cli/extending
  - /en/docs/cli/extending
---

# Extending Buffalo

Extending the CLI implies overriding the default plugins that the CLI ships with. To do it the CLI package provides two starting points:

- The `cli.NewApp` function
- The `cli.NewWithDefaults` function

Both of these return an instance of the Buffalo CLI, which can be modified with the help of the `Clear`, `Add` and `Remove` methods. The two main use cases for modifying the CLI are:

- To add/replace based on project specifics
- To add/replace based on user preferences

The following sections would describe how to do each of these.

## Extending for the project
One of the common needs is to add functionality to the CLI that is related to the project or team that uses the codebase. This can be done by adding a `main.go` file in the `cmd/buffalo` folder within the codebase. When the buffalo command runs it will run that file instead of the default one.

A typical `cmd/buffalo/main.go` file would look like this:

```go
package main

func main() {
  a := cli.NewWithDefaults()
  a.Add(&my.Plugin{})
  a.Add(&my.SeccondPlugin{})
  
  a.Run()
}
```

## Extending for the user
Another use case for extending the CLI is using custom plugins to initialize Buffalo applications, to do this one could override the CLI at the user level by adding a `.buffalo/cmd/` folder in the user `$HOME` root. When the Buffalo CLI finds this file it will attempt to run it instead of the default plugins.

Like with the project specific plugins, the user can override the CLI with custom plugins. One important thing to note is that the project specific CLI will take precedence over the user specific CLI.



