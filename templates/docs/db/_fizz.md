### Fizz Migrations

The `soda` command will generate SQL migrations (both the up and down) files for you.

```bash
$ soda generate fizz name_of_migration
```

Running this command will generate the **empty** following files:

```text
./migrations/20160815134952_name_of_migration.up.fizz
./migrations/20160815134952_name_of_migration.down.fizz
```

The generated files are `fizz` files. Pop uses [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) to generate migrations that are both easy to work with and work across multiple types of databases.

Further info about this command can be found by using the `--help` flag:

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

<%= warning() { %>
By default, the migration will create an UUID `id` that serves as the primary key, as well as `created_at` and `updated_at` datetime columns, so there is no need to create your own. These are the default, but you can override them if you want.
<% } %>