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

<%= title("Configuration") %>

Pop configuration is managed by a `database.yml`, located at the root of your project. This file is generated for you by Buffalo, if you choose to use Pop, and contains a basic configuration for the database you selected on generation with the `--db-type` flag. PostgreSQL is considered as the default.

Here is a sample configuration generated for a new app based on PostgreSQL:

```yaml
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

You can see three connections defined:
* `development` is the one used when your app runs on dev mode.
* `test` serves to run the integration tests.
* `production` is the config you'll use on the final app, on the server.

Of course, you can configure any new connection you want, but Buffalo won't pick them by default.

### Env vs detailed configuration

<%= note() { %>
Note that the `database.yml` file is also a Go template, so you can use Go template syntax. There are two special functions that are included, `env` and `envOr`.
<% } %>

As you can see, you have two ways to configure a new connection:
* The one used by the `development` connection is the most detailed. It allows you to set each available parameter, one by one.
* The one used by the `test` and `production` connections is a bit different: it uses a variable (see the `{{ }}` marks?) to set the value, and the `envOr` helper.

The `envOr` helper tries to get a value from an environment variable, and default to the second value. For instance:

```yaml
envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"
```

Tries to get the `TEST_DATABASE_URL` value from environment, and defaults to `postgres://postgres:postgres@127.0.0.1:5432/myapp_test`.

This way, you can provide a default value for development purposes, and allow to reconfigure the database settings from an environment variable!

<%= note() { %>
The `url` param for a connection will override any other connection param. Make sure you set all the settings you want from the URL string.
<% } %>

For additional details, check the documentation for [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Make sure you have configured this file properly before working with Pop!**