### Migrations SQL

<%= partial("docs/db/soda_buffalo_note.md") %>

Si vous ne voulez pas utiliser Fizz, ou si vous souhaitez exécuter une requête complexe, vous pouvez utiliser des migrations SQL.

Pour générer une nouvelle migration, utilisez la commande suivante :

```bash
$ soda generate sql nom_de_la_migration
```

Lancer cette commande génère les fichiers **vides** suivants :

```text
./migrations/20160815134952_nom_de_la_migration.up.sql
./migrations/20160815134952_nom_de_la_migration.down.sql
```

Pour en savoir plus sur cette commande, vous pouvez utiliser l'option `--help` :

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
