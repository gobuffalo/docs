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

<%= title("database.yml", {name: "configuring"}) %>

When you first generate a Buffalo application, a `database.yml` file will be generated for you, based on the type of database that was selected with the `--db-type` flag. PostgreSQL is considered as the default.

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

**CONFIGURE THIS FILE!**

Make sure to set up the appropriate usernames, passwords, hosts, etc... that are appropriate for the environment that will be running the application. Buffalo **does not** install these databases, or start up any services for you. If the database is running on a different port, you can add it to the default configuration file as `port: [port]`.

For example, if the development database is running on Docker using a random port `32768`, the configuration would be as follows:

```yaml
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  port: 32768
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

Note that the `database.yml` file is also a Go template, so you can use Go template syntax. There are two special functions that are included, `env` and `envOr`.

The generated `database.yml` file contains a template helper, `envOr`, used to define the URL for the test and production databases. It will attempt to find the corresponding ENV var, for example `DATABASE_URL` for production, if that ENV var does not exist, it will load the "default" string.

For additional details, check the documentation for [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).


<%= title("Creating Databases") %>

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Buffalo can create all of the databases in the `database.yml` file with a simple command:

```bash
$ buffalo db create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ buffalo db create -e test
```

<%= title("Dropping Databases") %>

Buffalo can drop all of your databases, should you want to, with one command:

```bash
$ buffalo db drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ buffalo db drop -e test
```

