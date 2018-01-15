# Migrations

<%= partial("docs/db/fizz.md") %>
<%= partial("docs/db/sql.md") %>

<%= title("Running Migrations") %>

Once migrations have been created they can be run with either of the following commands:

```bash
$ buffalo db migrate
$ buffalo db migrate up
```

Both commands are identical, one is shorter to type!

More information about the migration command be found by running:

```bash
$ buffalo db migrate --help

Runs migrations against your database.

Usage:
  buffalo db migrate [flags]
  buffalo db migrate [command]

Aliases:
  migrate, m


Available Commands:
  down        Apply all of the 'down' migrations.
  reset       The equivalent of running `migrate down` and then `migrate up`
  up          Apply all of the 'up' migrations.

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "buffalo db migrate [command] --help" for more information about a command.
```

<%= partial("docs/db/deployed_app.md") %>
