---
Name: Fizz
seoDescription: "Fizz"
seoKeywords: ["buffalo", "go", "golang", "database", "fizz", "pop", "DSL"]
weight: 13
aliases:
  - /docs/db/fizz
  - /es/docs/db/fizz
---

# Fizz
Fizz es un DSL común para migrar bases de datos. Intenta ser tan independiente posible de la base de datos. Este es el lenguaje por defecto que usa Pop para definir [migraciones en base de datos](/documentation/database/migrations).

## Crear una tabla

``` go
create_table("users") {
  t.Column("email", "string", {})
  t.Column("twitter_handle", "string", {"size": 50})
  t.Column("age", "integer", {"default": 0})
  t.Column("admin", "bool", {"default": false})
  t.Column("company_id", "uuid", {"default_raw": "uuid_generate_v1()"})
  t.Column("bio", "text", {"null": true})
  t.Column("joined_at", "timestamp", {})
}

create_table("todos") {
  t.Column("user_id", "integer", {})
  t.Column("title", "string", {"size": 100})
  t.Column("details", "text", {"null": true})
  t.ForeignKey("user_id", {"users": ["id"]}, {"on_delete": "cascade"})
}
```

La columna id no tiene que ser un entero. Por ejemplo, en su lugar, puedes usar un tipo [`UUID`](https://github.com/gofrs/uuid):

```go
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

Por defecto, fizz generará dos columnas de tipo `timestamp`: `created_at` and `updated_at`.

El metodo `t.Column` toma los siguientes argumentos: Nombre de la columna, el tipo de campo y finalmente el último argumento es cualquier opción que deseas establecer a tu columna.

#### Tipos "Comunes":

* `string`
* `text`
* `timestamp`, `time`, `datetime`
* `integer`
* `bool`
* `uuid`


Cualquier otro tipo pasado se pasará directamente a la base de datos subyacente.

Por ejemplo para PostgreSQL puedes pasar `jsonb` y podrá ser admitido, SQLite te gritará muy fuerte si haces lo mismo!

#### Opciones admitidas:

| Opción               | Descripción                                                                                                    | Ejemplo                                                     |
|:---------------------|:---------------------------------------------------------------------------------------------------------------|:------------------------------------------------------------|
| `primary`            | Si la columna es la llave primaria. Para llaves primarias compuestas, ve [abajo](#llaves-primarias-compuestas) | `{"primary": true}`                                         |
| `size`               | El tamaño de la columna. El valor predeterminado para una cadena es 255 (o 191 en MariaDB)                     | `{"size": 50}`                                              |
| `scale`, `precision` | La escala y la precisión para una columna de tipo float                                                        | `{"scale": 4, "precision": 2} `                             |
| `null`               | Por defecto las columnas no se permiten ser nulas                                                              | `{"null": true}`                                            |
| `default`            | El valor por defecto que deseas para tu columns. Por defecto es `null`                                         | `{"default": 0}` `{"default": false}` `{"default": "foo"}`  |
| `default_raw`        | El valor predeterminado definido como una función de base de datos.                                            | `{"default_raw": "uuid_generate_v1()"}`                     |
| `after`              | (**Solo para MySQL**) Agrega una columna después de otra en la tabla.                                          | `{"after": "created_at"}`                                   |
| `first`              | (**Solo para MySQL**) Agrega una columna en la primera pocisión en la tabla.                                   | `{"first": true}`                                           |


#### Deshabilitar marcas de tiempo automáticas

```go
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
  // Disable auto-creation of created_at and updated_at columns
  t.DisableTimestamps()
}
```

o

```go
create_table("users", {timestamps: false}) {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

## Borrar una tabla

``` go
drop_table("table_name")
```

## Renombrar tabla

``` go
rename_table("old_table_name", "new_table_name")
```

## Agregar una columna

``` go
add_column("table_name", "column_name", "string", {})
```

Mira [arriba](#tipos-comunes) para mas detalles sobre tipos de columnas y opciones.

## Alterar una columna

``` go
change_column("table_name", "column_name", "string", {})
```

## Renombrar una columna

``` go
rename_column("table_name", "old_column_name", "new_column_name")
```

## Borrar una columna

``` go
drop_column("table_name", "column_name")
```

## Llaves primarias compuestas

``` go
t.PrimaryKey("column_1", "column_2")
```

Ten en cuenta que la declaración `t.PrimaryKey` DEBE estar después de las definiciones de las columnas.

## Agregar un Index

#### Optiones admitidas:

* `name` - Por defecto es: `table_name_column_name_idx`
* `unique`

### Index simple:

``` go
add_index("table_name", "column_name", {})
```

### Index de varias columnas:

``` go
add_index("table_name", ["column_1", "column_2"], {})
```

### Index unico:

``` go
add_index("table_name", "column_name", {"unique": true})
```

### Nombre de Index:

``` go
add_index("table_name", "column_name", {}) // index name => table_name_column_name_idx
add_index("table_name", "column_name", {"name": "custom_index_name"})
```

## Renombrar un Index

``` go
rename_index("table_name", "old_index_name", "new_index_name")
```

## Borrar un Index

``` go
drop_index("table_name", "index_name")
```

## Agregar una Llave foránea

``` go
add_foreign_key("table_name", "field", {"ref_table_name": ["ref_column"]}, {
    "name": "optional_fk_name",
    "on_delete": "action",
    "on_update": "action",
})

```

#### Opciones admitidas

* `name` - Por defecto es: `table_name_ref_table_name_ref_column_name_fk`
* `on_delete` - `CASCADE`, `SET NULL`, ...
* `on_update`

{{<warning>}}
**Nota:** `on_update` y `on_delete` aun no son soportados en CockroachDB.
{{</warning>}}

## Borrar una Llave foránea

``` go
drop_foreign_key("table_name", "fk_name", {"if_exists": true})
```

#### Opciones admitidas

* `if_exists` - Agrega la condición `IF EXISTS`

## Raw SQL

``` go
sql("select * from users;")
```

## Ejecutar un comando externo

A veces, durante una migración, debe pagar a un comando externo.

``` go
exec("echo hello")
```
