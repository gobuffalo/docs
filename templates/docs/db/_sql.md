<%= panel("Generating Migrations (SQL)", {name: "generating-sql"}) { %>

```text
$ buffalo db g sql --help

Generates Up/Down migrations for your database using SQL.

Usage:
  buffalo soda generate sql [name] [flags]

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

<% } %>
