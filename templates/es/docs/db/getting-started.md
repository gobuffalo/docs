<% seoDescription("Comenzando con Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "ORM", "pop", "migraciones"]) %>

<%= h1("Comenzando con Pop") %>

El paquete [pop](https://godoc.org/github.com/gobuffalo/pop) se incluye con Buffalo de manera predeterminada, pero se puede utilizar fuera de Buffalo. Adjunta la biblioteca absolutamente increíble https://github.com/jmoiron/sqlx y limpia algunos de los patrones comunes y flujos de trabajo generalmente asociados con el manejo de bases de datos en Go.

**Pop hace que sea fácil hacer operaciones CRUD con funcionalidad ORM básica, ejecutar migraciones y generar/ejecutar consultas.**

Pop, por defecto, sigue las convenciones que fueron influenciadas por la gema ActiveRecord de Ruby. ¿Qué significa esto?

* Las tablas tienen que tener una columna "id" y un campo "ID" correspondiente en la estructura que se está utilizando.
* Si hay una columna de marca de tiempo denominada `created_at`, y un atributo` CreatedAt time.Time` en la estructura, se establecerá con la hora actual cuando se crea el registro.
* Si hay una columna de marca de tiempo denominada `updated_at`, y un atributo` UpdatedAt time.Time` en la estructura, se establecerá con la hora actual cuando se actualice el registro.
* Los nombres predeterminados de la tabla de la base de datos son versiones en minúsculas, en plural y subrayadas del nombre de la estructura. Ejemplos: `Usuario {}` es "usuarios", `FooBar {}` es "foo_bars", etc...

Buffalo tiene una profunda integración con Pop, y te ayudará a generar todo lo que usted necesita para empezar. Todavía puede usar otro paquete si lo desea, pero estará solo. :)

## Bases de datos compatibles

Pop es compatible con las siguientes bases de datos:
* [PostgreSQL](https://www.postgresql.org/) (>= 9.3)
* [CockroachDB](https://www.cockroachlabs.com/) (>= 1.1.1)
* [MySQL](https://www.mysql.com/) (>= 5.7)
* [SQLite3](https://sqlite.org/) (>= 3.x)

## Siguientes pasos

* [Configuración](/es/docs/db/configuration) - Configure sus conexiones a la base de datos..