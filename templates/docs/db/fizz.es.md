<% seoDescription("Fizz") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "fizz", "pop", "DSL"]) %>

<%= h1("Fizz") %>

Fizz es una DSL común para migrar bases de datos. Intenta ser lo más independiente de la base de datos posible. Este es el lenguaje utilizado por defaut por Pop para definir las [migraciones de base de datos](/es/docs/db/migrations).

## Crear una tabla

``` javascript
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

La función `create_table` generará una columna `id` de tipo `integer` que se autoincrementará. Esto se puede cambiar para usar el tipo [`UUID`](https://github.com/gobuffalo/uuid) type:

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

También generará dos columnas `timestamp`; `created_at` y  `updated_at`.

El método `t.Column` toma los siguientes argumentos: nombre de la columna, tipo de campo y finalmente el último argumento es cualquier opción que quiera establecer en esa columna.

#### <a name="column-info"></a> "Common" Types:

* `string`
* `text`
* `timestamp`, `time`, `datetime`
* `integer`
* `bool`
* `uuid`

Cualquier otro tipo pasado pasará directamente a la base de datos subyacente.

Por ejemplo, para PostgreSQL puede pasar `jsonb` y será compatible, sin embargo, SQLite le gritará muy fuerte si hace lo mismo.

#### Opciones compatibles:

* `size` - El tamaño de la columna. Por ejemplo, si quisiera un `varchar (50)` en Postgres, haría: `t.Column (" column_name "," string ", {" size ": 50})`
* `null` - Por defecto, las columnas no pueden ser` null`.
* `default` - El valor predeterminado que desea para esta columna. Por defecto, esto es `nulo`.
* `default_raw` -  El valor predeterminado definido como una función de base de datos.
* `after` - (MySQL solamente) Agregue una columna después de otra columna en la tabla. `ejemplo: {"after":"created_at"}`
* `first` - (MySQL solamente) Agregue una columna a la primera posición en la tabla. `ejemplo: {"first": true}`

#### Desactivar marcas de tiempo automáticas

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
  // Disable auto-creation of created_at and updated_at columns
  t.DisableTimestamps()
}
```

o

```javascript
create_table("users", {timestamps: false}) {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

## Borrar una tabla

``` javascript
drop_table("table_name")
```

## Renombrar una table

``` javascript
rename_table("old_table_name", "new_table_name")
```

## Agregar una columna

``` javascript
add_column("table_name", "column_name", "string", {})
```

Ver [arriba](#column-info) para más detalles sobre tipos de columnas y opciones.

## Alterar una columna

``` javascript
change_column("table_name", "column_name", "string", {})
```

## Renombrar una columna

``` javascript
rename_column("table_name", "old_column_name", "new_column_name")
```

## Borrar una columna

``` javascript
drop_column("table_name", "column_name")
```

## Agregar un indice

#### Opciones compatibles:

* `name` - Esto es `table_name_column_name_idx`
* `unique`

### Índice simple:

``` javascript
add_index("table_name", "column_name", {})
```

### Índice de múltiples columnas:

``` javascript
add_index("table_name", ["column_1", "column_2"], {})
```

### Índice único:

``` javascript
add_index("table_name", "column_name", {"unique": true})
```

### Nombres de índice:

``` javascript
add_index("table_name", "column_name", {}) # name => table_name_column_name_idx
add_index("table_name", "column_name", {"name": "custom_index_name"})
```

## Renombrar un índice

``` javascript
rename_index("table_name", "old_index_name", "new_index_name")
```

## Borrar un índice

``` javascript
drop_index("table_name", "index_name")
```

## Agregar una clave externa

```javascript
add_foreign_key("table_name", "field", {"ref_table_name": ["ref_column"]}, {
    "name": "optional_fk_name",
    "on_delete": "action",
    "on_update": "action",
})

```

#### Opciones compatibles:

* `name` - Por defecto es `table_name_ref_table_name_ref_column_name_fk`
* `on_delete` - `CASCADE`, `SET NULL`, ...
* `on_update`

**Nota:** `on_update` y` on_delete` aún no son compatibles con CockroachDB.

## Borrar una clave externa

```javascript
drop_foreign_key("table_name", "fk_name", {"if_exists": true})
```

#### Opciones compatibles:

* `if_exists` - Agrega la condición` IF EXISTS`


## SQL sin formato

``` javascript
sql("select * from users;")
```

## Ejecute un comando externo

Algunas veces, durante una migración, se necesita ejecutar un comando externo.

```javascript
exec("echo hello")
```
