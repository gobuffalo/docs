<%= title("Fizz Migrations", {name:"generating-fizz"}) %>

<%= partial("docs/db/soda_buffalo_note.md") %>

Pop uses [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) to generate migrations that are both easy to work with and work across multiple types of databases.

To generate a new **empty** migration, use the following command:

```bash
$ soda g migration --help

Generates Up/Down migrations for your database using fizz.

Usage:
  soda generate fizz [name] [flags]

Aliases:
  fizz, migration

Flags:
  -h, --help   help for fizz

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

