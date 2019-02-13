<% seoDescription("How to configure my database with Pop?") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "configuration"]) %>

<%= h1("Database Configuration") %>

Pop configuration is managed by a `database.yml` file, located at the root of your project. This file is generated for you if you use Buffalo &#8211; if you choose to use Pop &#8211; and contains a basic configuration for the database you selected on generation with the `--db-type` flag. PostgreSQL is considered as the default.

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

### Generator

<%= partial("docs/db/soda_buffalo_note.md") %>

You can generate a default configuration file using the init command:

```bash
$ soda g config
```

The default will generate a `database.yml` file in the current directory for a PostgreSQL database. You can override the type of database using the `-t` flag and passing in any of the supported database types: `postgres`, `cockroach`, `mysql`, or `sqlite3`.

### database.yml

The Pop configuration file &#8211; `database.yml` &#8211; can be found either:
* At your project root (default).
* In the `config/` directory, at your project root.

If you want to put your config file in another location, you can use the [`AddLookupPaths`](https://godoc.org/github.com/gobuffalo/pop#AddLookupPaths).

You can also customize the file name:

```go
pop.ConfigName = "my_pop_config.yml"
```

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

<%= warning() { %>
The `url` param for a connection will override any other connection param. Make sure you set all the settings you want from the URL string.
<% } %>

For additional details, check the documentation for [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Make sure you have configured this file properly before working with Pop!**

## Available Options

### database

The name of the database to use.

### dialect

The database dialect to use with the connection. Accepted values are:
* MySQL driver: "mysql"
* PostgreSQL driver: "postgres", "postgresql" or "pg"
* Cockroach driver: "cockroach", "cockroachdb" or "crdb"
* SQLite driver: "sqlite" or "sqlite3"

### encoding

<%= sinceVersion("4.6.0") %>

This option is currently only supported by the **mysql dialect**. This encoding will be used to create the database (if you create it using `soda`), and as the `collation` parameter for the connection string. If this option is omitted, the default value is `utf8mb4_general_ci`.

```yaml
development:
  dialect: mysql
  options:
    encoding: "utf8_general_ci"
```

### host

The database host address to connect to.

### password

The password for the user you use to connect to the database.

### port

The database host port for the database.

**Defaults**:

| Driver    | Port  |
|-----------|-------|
| PostgreSQL| 5432  |
| MySQL     | 3306  |
| Cockroach | 26257 |

### user

The user to use to connect to the database.