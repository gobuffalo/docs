<% seoDescription("Migrations") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "migration"]) %>

<%= h1("Migrations") %>

Software maintenance is a hard task, and you'll probably need to patch your database to add, modify or remove some fields. The way to handle that with Pop is to use **migrations**.

You can create new migrations using `fizz`, a custom language describing the database changes in the most database-agnostic way; or use SQL statements if you prefer.

<%= partial("docs/db/fizz.md") %>
<%= partial("docs/db/sql.md") %>

<%= title("Running Migrations") %>

<%= partial("docs/db/soda_buffalo_note.md") %>

Once migrations have been created they can be run with either of the following commands:

```bash
$ soda migrate
$ soda migrate up
```

Both commands are identical, one is shorter to type!

If you want to rollback the last applied migration, use the following command:

```bash
$ soda migrate down
```

More information about the migration command be found by running:

```bash
$ soda migrate --help

Runs migrations against your database.

Usage:
  soda migrate [flags]
  soda migrate [command]

Aliases:
  migrate, m

Available Commands:
  down        Apply one or more of the 'down' migrations.
  reset       The equivalent of running `migrate down` and then `migrate up`
  status      Displays the status of all migrations.
  up          Apply all of the 'up' migrations.

Flags:
  -h, --help   help for migrate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "soda migrate [command] --help" for more information about a command.
```

<%= title("Targeting a Database") %>

Since Pop [v4.4.0](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), migrations can target a specific database, using a suffix. This allows to use commands specific to a dialect, only for a given database.

For instance, if you want to support both PostgreSQL and MySQL, you can create two migrations:

* `my-migration.mysql.up.sql` and `my-migration.mysql.down.sql` will be used when migrating a MySQL database.
* `my-migration.postgres.up.sql` and `my-migration.postgres.down.sql` will be used when migrating a PostgreSQL database.

If no version for the dialect can be found, Pop will fallback to the non-suffixed version, if it exists.

<%= partial("docs/db/deployed_app.md") %>