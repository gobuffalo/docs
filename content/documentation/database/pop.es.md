---
name: Iniciando con Pop
seoDescription: Iniciando con Pop
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "migrations"]
weight: 1
aliases:
  - /documentation/database/getting-started
  - /docs/db/getting-started
  - /es/docs/db/getting-started
---

# Iniciando con Pop

La librería de [pop](https://godoc.org/github.com/gobuffalo/pop) esta incluída en Buffalo por defecto, pero puedes usarla fuera de Buffalo. Esta envuelve completamente librería [sqlx](https://github.com/jmoiron/sqlx), limpia algunos de los patrones comunes y flujos usualmente asociados con el manejo de bases de datos en Go.

**Pop facilita hacer las operaciones CRUD con la funcionalidad básica de ORM, ejecutar migraciones y construit/ejecutar consultas**

Pop, por defecto, sigue las convenciones que fueron influenciadas por la gema de Ruby `ActiveRecord`. Que significa esto?

* Las tablas en la base de datos deben tener una columna "id" correspondiente al campo "ID" en la estructura a ser usada.
* Si hay una columna de tipo `timestamp` llamada `created_at` y un campo de tipo `time.Time` llamado `CreatedAt`, se le asignará la hora actual cuando el registro es creado.
* Si hay una columna de tipo `timestamp` llamada `updated_at` y un campo de tipo `time.Time` llamado `UpdatedAt`, se le asignará la hora actual cuando el registro es actualizado.
* Los nombres en las tablas de la base de datos por defecto son en minúscula, plural, y en la version *unserscored* del nombre de la estructura. Ejemplo: `User{}` es "users", `FooBar{}` es "foo_bars", etc...

Buffalo tiene una profunda integracion con Pop, y te ayudará a generar todas las cosas que necesites para empezar. Aun puedes usar otro paquete si lo deseas, pero estarás solo. :)

## Base de datos soportadas

Pop soporta las siguientes bases de datos:

* [PostgreSQL](https://www.postgresql.org/) (>= 9.3)
* [CockroachDB](https://www.cockroachlabs.com/) (>= 2.1.0)
* [MySQL](https://www.mysql.com/) (>= 5.7)
* [SQLite3](https://sqlite.org/) (>= 3.x)


## Instalación

```bash
$ go get github.com/gobuffalo/pop/...
```

## Siguientes pasos

* [CLI Soda](/es/documentation/database/soda) - Instala el CLI de the Soda.
* [Configuración](/es/documentation/database/configuration) - Configura tu conexión de la base de datos.
