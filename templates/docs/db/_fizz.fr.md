### Migrations Fizz

<%= partial("docs/db/soda_buffalo_note.md") %>

La commande `soda` permet de générer des fichiers de migrations SQL (une pour appliquer, l'autre pour annuler) pour vous.

```bash
$ soda generate fizz nom_de_la_migration
```

Lancer cette commande génère les fichiers **vides** suivants :

```text
./migrations/20160815134952_nom_de_la_migration.up.fizz
./migrations/20160815134952_nom_de_la_migration.down.fizz
```

Ces migrations sont des fichiers `fizz`. Pop utilise [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) pour générer des migrations à la fois faciles à écrire, et qui s'adaptent autant que faire se peut aux différents types de bases de données.

Pour en savoir plus sur cette commande, vous pouvez utiliser l'option `--help` :

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