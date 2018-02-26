# Database

Buffalo uses [github.com/markbates/pop](https://github.com/markbates/pop) as its default database package for migrations, transactions, basic ORM functionality, and more.

<%= title("Getting Started") %>

Buffalo supports [PostgreSQL](https://www.postgresql.org/) (default), [CockroachDB](https://www.cockroachlabs.com/), [MySQL](https://www.mysql.com/), and [SQLite3](https://sqlite.org/). When you generate a new Buffalo application you can change this with the `--db-type` flag. It is also possible to skip generation of all database components with the `--skip-pop` flag.

<%= title("database.yml", {name: "configuring"}) %>

When you first generate a Buffalo application a `database.yml` file will be generated for you, based on the type of database that was selected with the `--db-type` flag, with PostgreSQL being the default.

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

For additional details, check the documentation for [github.com/markbates/pop](https://github.com/markbates/pop).


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

