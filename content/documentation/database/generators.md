---
name: Generators
weight: 11
aliases:
  - /docs/db/generators
  - /en/docs/db/generators
---
# Generators

```bash
$ buffalo pop g --help
Generates config, model, and migrations files.

Usage:
  buffalo-pop pop generate [command]

Aliases:
  generate, g

Available Commands:
  config      Generates a database.yml file for your project.
  fizz        Generates Up/Down migrations for your database using fizz.
  model       Generates a model for your database
  sql         Generates Up/Down migrations for your database using SQL.

Flags:
  -h, --help   help for generate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "buffalo-pop pop generate [command] --help" for more information about a command.
```


## Migrations

For information on generating migrations see [Migrations](/documentation/database/migrations).
