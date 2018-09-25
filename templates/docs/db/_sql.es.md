### Migraciones de SQL

Si no desea usar Fizz, o si tiene una consulta complicada que desea ejecutar, puede usar SQL.

Para generar una nueva migración ** vacía **, use el siguiente comando:

```bash
$ soda generate sql name_of_migration
```

Al ejecutar este comando, se generarán los siguientes archivos:

```text
./migrations/20160815134952_name_of_migration.up.sql
./migrations/20160815134952_name_of_migration.down.sql
```

Se puede encontrar más información sobre este comando usando el _flag_ `--help`:

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
