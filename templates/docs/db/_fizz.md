<%= title("Fizz Migrations", {name:"generating-fizz"}) %>

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `db` namespace. So everytime you want to use a command from `soda`, just execute `buffalo db` instead.
<% } %>

Pop uses [Fizz](https://github.com/gobuffalo/pop/blob/master/fizz%2FREADME.md) to generate migrations that are both easy to work with and work across multiple types of databases.

To generate a new **empty** migration, use the following command:

```bash
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
 