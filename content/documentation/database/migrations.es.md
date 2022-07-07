---
name: Migraciones
seoDescription: "Como crear y usar las migraciones de base de datos con Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "migration"]
weight: 12
aliases:
  - /docs/db/migrations
  - /es/docs/db/migrations
---
# Migraciones

El mantenimiento de Software es una tarea difícil, y probablemente necesites parchear tu base de datos para agregar, modificar o remover algunos campos. La forma para manejar esto con Pop es usar **migraciones**.

Puedes crear nuevas migraciones usando `fixx`, un lenguaje personalizado que describe los cambios de la manera mas independiente a la base de datos; o usa sentencias SQL y lo prefieres.

## Escribiendo migraciones

{{<note>}}
**Nota para usuarios de Buffalo**: Los comandos de `Soda` estan adheridos de los comandos de `Buffalo`, detrás del espacio de nombres de `pop`. Asi que cada vez que desees usar un comando de `soda`, solo ejecuta `buffalo pop` en su lugar. No necesitas instalar la CLI de `soda`.
{{</note>}}

### Migraciones Fizz

El comando de `soda` que generará los archivos migracion SQL para ti (tanto la migración arriba como abajo).

```bash
$ soda generate fizz name_of_migration
```

Ejecutando este comando generará los siguientes archivos **vacíos**:

```bash
./migrations/20220706213354_name_of_migration.up.fizz
./migrations/20220706213354_name_of_migration.down.fizz
```

Los archivos generados son `fizz`. Pop usa[Fizz](https://github.com/gobuffalo/fizz/blob/master/README.md) para generar las migraciones con las que son fáciles para trabajar y funcionan en múltiples tipos de bases de datos.

Puedes encontrar más información sobre este comando usando el flag `--help`:

```bash
$ soda g migration --help
Generates Up/Down migrations for your database using fizz.

Usage:
  buffalo-pop pop generate fizz [name] [flags]

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

{{<warning>}}
Por defecto, la migracion creará un UUID `id` que sirve como la llave primaria, tan bien como las columnas de fecha y hota `created_at` y `updated_at`, por lo que no es necesario crear una propia. Estos valores son predeterminados pero tu puedes cambiarlos si deseas.
{{</warning>}}

### Migraciones SQL

SI no deseas usar `fizz`, o tienes una consulta muy complicada que deseas ejecutar, puedes usar SQL.
If you don't want to use Fizz, or you have a complicated query you want to execute, you can use SQL.


Para generar una nueva migración **vacía**, usa el siguiente comando:

```bash
$ soda generate sql name_of_migration
```

Ejecutando este comando, se generarán los siguientes archivos:
Running this command will generate the following files:

```bash
./migrations/20220706213354_name_of_migration.up.sql
./migrations/20220706213354_name_of_migration.down.sql
```

Puedes encontrar más información sobre este comando usando el flag `--help`:

```bash
$ soda g sql --help
Generates Up/Down migrations for your database using SQL.

Usage:
  buffalo-pop pop generate sql [name] [flags]

Flags:
  -h, --help   help for sql

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```


## Ejecutando migraciones

{{<note>}}
**Nota para usuarios de Buffalo**: Los comandos de `Soda` estan adheridos de los comandos de `Buffalo`, detrás del espacio de nombres de `pop`. Asi que cada vez que desees usar un comando de `soda`, solo ejecuta `buffalo pop` en su lugar. No necesitas instalar la CLI de `soda`.
{{</note>}}


### Aplicando las migraciones
Una vez que las migraciones hayan sido creadas, ellas pueden ser ejecutadas con cualquiera de los siguientes comandos:

```bash
$ soda migrate
$ soda migrate up
```
Ambos comandos son identicos, uno es más corto de escribir! Las migraciones se ejecutarán en orden secuencial.

### Revertir una migración
Si deseas revertir la ultima migración aplicada, usa el siguiente comando:

```bash
$ soda migrate down
```

---

Puedes encontrar más información sobre el comando de migración ejecutando:

```bash
$ soda migrate --help
Runs migrations against your database.

Usage:
  soda migrate [flags]
  soda migrate [command]

Aliases:
  migrate, m

Available Commands:
  down        Apply one or more of the 'down' migrations.
  status      Displays the status of all migrations.
  up          Apply one or more of the 'up' migrations.

Flags:
  -h, --help   help for migrate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "soda migrate [command] --help" for more information about a command.
```

## Dirigirse a una base de datos

Desde Pop [`v4.4.0`](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), las migraciones pueden dirigirse a una base de datos específica, usando un sufijo. Esto permite usar comandos específicos a un dialecto, solo para unsa base de datos determinada.

Por ejemplo, si quieres soporte para PostgreSQL y MySQL, puedes crear dos migraciones:

* `my-migration.mysql.up.sql` y `my-migration.mysql.down.sql` se usarán al migrar a una base de datos MySQL.
* `my-migration.postgres.up.sql` y `my-migration.postgres.down.sql` se usarán al migrar a una base de datos PostgreSQL.

Si no se puede encontrar una versión para el dialecto, Pop recurrirá a la versión sin sufijo, si existe.

## Custom Migrations Table

Por defecto, las migraciones estan registradas en la tabla `schema_migration`. esta tabla es creada por Pop si no existe.

Sin embargo en algunos casos, esposible que desees usar un nombre diferente para esta tabla. Desde Pop `v4.5.0`, puedes personalizar el nombre de esta tabla usando la opción `migration_table_name`. El siguiente ejemplo usaremos `migrations` como el nombre de la tabla:

```yaml
development:
  dialect: "postgres"
  url: "your_db_development"
  options:
    migration_table_name: migrations
```

## Migraciones una vez desplegadas

{{< note "Esta sección es solo para usuarios de Buffalo.">}}

Cuando cionstruyes tu aplicaión, las migraciones se almacenan dentro de tu binario. Tu binario tiene un comando `migrate` escondido que realiza las migraciones, tal como lo hace cuando usa `buffalo pop migrate`:

```bash
$ ./myapp migrate
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213171622
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172104
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172249
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213173148
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219070903
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219071524

0.0010 seconds
```
