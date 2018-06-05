<%= title("SQL Migrations", {name: "generating-sql"}) %>

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `db` namespace. So everytime you want to use a command from `soda`, just execute `buffalo db` instead.
<% } %>

If you don't want to use Fizz, or you have a complicated query you want to execute, you can use SQL.

To generate a new **empty** migration, use the following command:

```bash
$ soda g sql --help

Generates Up/Down migrations for your database using SQL.

Usage:
  soda generate sql [name] [flags]

Flags:
  -h, --help   help for sql

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```
