---
name: Migrations
seoDescription: "How to create and use database migrations with Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "migration"]
weight: 12
aliases:
  - /docs/db/migrations
  - /en/docs/db/migrations
---
# Migrations

Software maintenance is a hard task, and you'll probably need to patch your database to add, modify or remove some fields. The way to handle that with Pop is to use **migrations**.

You can create new migrations using `fizz`, a custom language describing the database changes in the most database-agnostic way; or use SQL statements if you prefer.

## Writing Migrations
{{< note >}}
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead.
{{< /note >}}


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

{{< warning >}}
By default, the migration will create an UUID `id` that serves as the primary key, as well as `created_at` and `updated_at` datetime columns, so there is no need to create your own. These are the default, but you can override them if you want.
{{< /warning >}}

### SQL Migrations

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


## Running Migrations
{{< note >}}
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead.
{{< /note >}}


### Apply Migrations
Once migrations have been created they can be run with either of the following commands:

```bash
$ soda migrate
$ soda migrate up
```

Both commands are identical, one is shorter to type! Migrations will be run in sequential order.

### Rollback a Migration
If you want to rollback the last applied migration, use the following command:

```bash
$ soda migrate down
```

---

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

## Targeting a Database

Since Pop [v4.4.0](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), migrations can target a specific database, using a suffix. This allows to use commands specific to a dialect, only for a given database.

For instance, if you want to support both PostgreSQL and MySQL, you can create two migrations:

* `my-migration.mysql.up.sql` and `my-migration.mysql.down.sql` will be used when migrating a MySQL database.
* `my-migration.postgres.up.sql` and `my-migration.postgres.down.sql` will be used when migrating a PostgreSQL database.

If no version for the dialect can be found, Pop will fallback to the non-suffixed version, if it exists.

## Custom Migrations Table

By default, the applied migrations are tracked in the table `schema_migration`. This table is created by pop if it doesn't exist.

In some cases, though, you may want to use a different name for this table. Since pop v4.5.0, you can customize the name of this table using the `migration_table_name` option. The example below will use `migrations` as the table name:

```yaml
development:
  dialect: "postgres"
  url: "your_db_development"
  options:
    migration_table_name: migrations
```

## Migrations Once Deployed

{{< note "This section is only for Buffalo users.">}}


When you build your app, the migrations are stored inside your binary. Your binary has a hidden `migrate` command baked in that performs the migrations, just like it does when you use `buffalo pop migrate`:

```bash
$ ./myapp migrate
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213171622
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172104
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172249
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213173148
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219070903
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219071524

0.0010 seconds
```
