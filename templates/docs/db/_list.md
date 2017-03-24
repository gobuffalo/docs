<%= title("Database Generators", {name: "generators"}) %>

```text
$ buffalo db g --help

Usage:
  buffalo db generate [command]

Aliases:
  generate, g


Available Commands:
  config      Generates a database.yml file for your project.
  fizz        Generates Up/Down migrations for your database using fizz.
  model       Generates a model for your database
  sql         Generates Up/Down migrations for your database using SQL.

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "buffalo db generate [command] --help" for more information about a command.
```
