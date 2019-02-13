<% seoDescription("Getting started with Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "migrations"]) %>

<%= h1("Getting started with Pop") %>

The [pop](https://godoc.org/github.com/gobuffalo/pop) package is included with Buffalo by default, but you can use it outside of Buffalo. It wraps the absolutely amazing https://github.com/jmoiron/sqlx library, cleans up some of the common patterns and work flows usually associated with dealing with databases in Go.

**Pop makes it easy to do CRUD operations with basic ORM functionality, run migrations, and build/execute queries.**

Pop, by default, follows conventions that were influenced by the ActiveRecord Ruby gem. What does this mean?

* Tables must have an "id" column and a corresponding "ID" field on the struct being used.
* If there is a timestamp column named `created_at`, and a `CreatedAt time.Time` attribute on the struct, it will be set with the current time when the record is created.
* If there is a timestamp column named `updated_at`, and a `UpdatedAt time.Time` attribute on the struct, it will be set with the current time when the record is updated.
* Default database table names are lowercase, plural, and underscored versions of the struct name. Examples: `User{}` is "users", `FooBar{}` is "foo_bars", etc...

Buffalo has a deep integration with Pop, and it'll help you to generate all the stuff you need to get started. You can still use another package if you want, but you'll be by yourself. :)

## Supported Databases

Pop supports the following databases:
* [PostgreSQL](https://www.postgresql.org/) (>= 9.3)
* [CockroachDB](https://www.cockroachlabs.com/) (>= 1.1.1)
* [MySQL](https://www.mysql.com/) (>= 5.7)
* [SQLite3](https://sqlite.org/) (>= 3.x)

## Buffalo Integration

When you generate a new Buffalo application you can choose the target database with the `--db-type` flag. For instance, to generate a new app with MySQL database support, you can write the following:

```bash
$ buffalo new coke --db-type mysql
```

**By default, Buffalo will generate an app with PostgreSQL as the backing database.**

### Skip database support

If you want to handle the database without using Pop, or if you're building an app without database, it's also possible to skip generation of all database components with the `--skip-pop` flag.

```bash
$ buffalo new coke --skip-pop
```

## Next Steps

* [Configuration](/en/docs/db/configuration) - Configure your database connections.