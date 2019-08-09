<% seoDescription("Cómo crear y usar migraciones de bases de datos con Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "ORM", "pop", "migración"]) %>

<%= h1("Migraciones") %>

El mantenimiento del software es una tarea difícil, y es probable que necesite parchear su base de datos para agregar, modificar o eliminar algunos campos. La forma de manejar eso con Pop es usar **migraciones**.

Puede crear nuevas migraciones usando `fizz`, un lenguaje personalizado que describe los cambios de la base de datos de la manera más independiente de las bases de datos; o use declaraciones SQL si lo prefiere.

## Escribir migraciones
<%= partial("en/docs/db/soda_buffalo_note.md") %>

<%= partial("en/docs/db/fizz.md") %>
<%= partial("en/docs/db/sql.md") %>

## Ejecutar migraciones
<%= partial("en/docs/db/soda_buffalo_note.md") %>

### Aplicar migraciones
Una vez que se han creado las migraciones, se pueden ejecutar con cualquiera de los siguientes comandos:

```bash
$ soda migrate
$ soda migrate up
```

Ambos comandos son idénticos, ¡uno es más corto de escribir! Las migraciones se ejecutarán en orden secuencial.

### Revertir una migración
Si desea deshacer la última migración aplicada, use el siguiente comando:

```bash
$ soda migrate down
```

---

Se puede encontrar más información sobre el comando de migración ejecutando:

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
  reset       The equivalent of running `migrate down` and then `migrate up`
  status      Displays the status of all migrations.
  up          Apply all of the 'up' migrations.

Flags:
  -h, --help   help for migrate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "soda migrate [command] --help" for more information about a command.
```

## Orientación de una base de datos

Desde Pop [v4.4.0](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), las migraciones pueden dirigirse a una base de datos específica, utilizando un sufijo. Esto permite usar comandos específicos de un dialecto, solo para una base de datos dada.

Por ejemplo, si desea admitir PostgreSQL y MySQL, puede crear dos migraciones:

* `my-migration.mysql.up.sql` y `my-migration.mysql.down.sql` se usarán al migrar una base de datos MySQL.
* `my-migration.postgres.up.sql` y `my-migration.postgres.down.sql` se usarán al migrar una base de datos PostgreSQL.

Si no se puede encontrar ninguna versión para el dialecto, Pop recurrirá a la versión no sufijada, si existe.

## Tabla de migraciones personalizadas

De forma predeterminada, las migraciones aplicadas se rastrean en la tabla `schema_migration`. Esta tabla es creada por Pop si no existe.

Sin embargo, en algunos casos, es posible que desee utilizar un nombre diferente para esta tabla. Desde Pop v4.5.0, se puede personalizar el nombre de esta tabla usando la opción `migration_table_name`. El siguiente ejemplo usará `migrations` como el nombre de la tabla:

```yaml
development:
  dialect: "postgres"
  url: "your_db_development"
  options:
    migration_table_name: migrations
```

<%= partial("en/docs/db/deployed_app.md") %>