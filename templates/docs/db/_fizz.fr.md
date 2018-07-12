### Migrations Fizz

<%= partial("docs/db/soda_buffalo_note.md") %>

Pop utilise [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) pour générer des migrations à la fois faciles à écrire, et qui s'adaptent autant que faire se peut aux différents types de bases de données.

Pour générer une nouvelle migration **vide**, utilisez la commande ci-dessous :

```bash
$ soda g migration --help

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

<%= warning() { %>
Par défaut, la migration créera un `id` de type UUID en tant que clé primaire, ainsi que les colonnes `created_at` et `updated_at` de type datetime. Il n'est donc pas nécessaire de les mentionner lors de la création de la migration, à moins de vouloir les définir autrement.
<% } %>