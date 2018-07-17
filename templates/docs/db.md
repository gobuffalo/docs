<% seoDescription("Getting started with databases") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "migrations"]) %>

<%= h1("Database") %>

The [Pop](https://github.com/gobuffalo/pop) package is included with Buffalo by default. It allows managing migrations, transactions, basic ORM functionality, and more. Just like any ORM-like package, it provides a way to map Model structures to database tables.

Buffalo has a deep integration with Pop, and it'll help you to generate all the stuff you need to get started. You can still use another package if you want, but you'll be by yourself. :)

<%= title("Getting Started") %>

Pop supports the following databases:
* [PostgreSQL](https://www.postgresql.org/)
* [CockroachDB](https://www.cockroachlabs.com/)
* [MySQL](https://www.mysql.com/)
* [SQLite3](https://sqlite.org/)

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

<%= title("Next Steps") %>

* [Configuration](/en/docs/db/configuration) - Configure your database connections.