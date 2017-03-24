<%= title("Generating Migrations (SQL)", {name: "generating-sql"}) %>

If you don't want to use Fizz, or you have a complicated query you want to execute, you can use SQL.

```text
$ buffalo db g sql --help

Generates Up/Down migrations for your database using SQL.

Usage:
  buffalo db generate sql [name] [flags]

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```
