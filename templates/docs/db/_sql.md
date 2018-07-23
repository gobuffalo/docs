### SQL Migrations

<%= partial("docs/db/soda_buffalo_note.md") %>

If you don't want to use Fizz, or you have a complicated query you want to execute, you can use SQL.

To generate a new **empty** migration, use the following command:

```bash
$ soda generate sql name_of_migration
```

Running this command will generate the following files:

```text
./migrations/20160815134952_name_of_migration.up.sql
./migrations/20160815134952_name_of_migration.down.sql
```

Further info about this command can be found by using the `--help` flag:

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
