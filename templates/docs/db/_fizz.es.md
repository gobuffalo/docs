### Migraciones Fizz

El comando `soda` generará migraciones de SQL (tanto arriba como abajo) para usted.

```bash
$ soda generate fizz name_of_migration
```

Al ejecutar este comando, se generarán los siguientes archivos **vacios**:

```text
./migrations/20160815134952_name_of_migration.up.fizz
./migrations/20160815134952_name_of_migration.down.fizz
```

Los archivos generados son archivos `fizz`. Pop usa [Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) para generar migraciones que son fáciles de utlizar y funcionan en varios tipos de bases de datos.

Se puede encontrar más información sobre este comando usando el _flag_ `--help`:

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
De forma predeterminada, la migración creará un UUID `id` que sirve como la clave principal, así como` created_at` y `updated_at` columnas de fecha y hora, por lo que no es necesario crearlas. Estos son los predeterminados, pero puede anularlos si lo desea.
<% } %>