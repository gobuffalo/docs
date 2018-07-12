### Migrations SQL

<%= partial("docs/db/soda_buffalo_note.md") %>

Si vous ne voulez pas utiliser Fizz, ou si vous souhaitez exécuter une requête complexe, vous pouvez utiliser des migrations SQL.

Pour générer une nouvelle migration **vide**, utilisez la commande suivante :

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
