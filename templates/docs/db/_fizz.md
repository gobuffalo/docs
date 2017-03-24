<%= title("Generating Migrations (fizz)", {name:"generating-fizz"}) %>

Buffalo uses [Fizz](https://github.com/markbates/pop/blob/master/fizz%2FREADME.md) to generate migrations that are both easy to work with and work across multiple types of databases.

```text
$ buffalo db g migration --help

Generates Up/Down migrations for your database using fizz.

Usage:
  buffalo db generate fizz [name] [flags]

Aliases:
  fizz, migration


Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```
