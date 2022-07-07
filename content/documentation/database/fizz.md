---
Name: Fizz
seoDescription: "Fizz"
seoKeywords: ["buffalo", "go", "golang", "database", "fizz", "pop", "DSL"]
weight: 13
aliases:
  - /docs/db/fizz
  - /en/docs/db/fizz
---

# Fizz
Fizz is a common DSL for migrating databases. It tries to be as database-agnostic as possible. This is the default language used by Pop to define [database migrations](/documentation/database/migrations).

## Create a Table

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

The `id` column doesn't have to be an integer. For instance, you can use an [`UUID`](https://github.com/gofrs/uuid) type instead:

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

By default, fizz will generate two `timestamp` columns: `created_at` and `updated_at`.

The `t.Column` method takes the following arguments: name of the column, the type of the field, and finally the last argument is any options you want to set on that column.

#### "Common" Types:

* `string`
* `text`
* `timestamp`, `time`, `datetime`
* `integer`
* `bool`
* `uuid`

Any other type passed will be passed straight through to the underlying database.

For example for PostgreSQL you could pass `jsonb` and it will be supported, however, SQLite will yell very loudly at you if you do the same thing!

#### Supported Options:

| Option               | Description                                                                                          | Example                                                    |
|:---------------------|:-----------------------------------------------------------------------------------------------------|:-----------------------------------------------------------|
| `primary`            | Whether the column is the primary key. To have a composite key look [below](#composite-primary-keys) | `{"primary": true}`                                        |
| `size`               | The size of the column. The default value for a string column is 255 (or 191 for MariaDB)            | `{"size": 50}`                                             |
| `scale`, `precision` | The scale and the precision for a float column                                                       | `{"scale": 4, "precision": 2} `                            |
| `null`               | By default columns are not allowed to be `null`                                                      | `{"null": true}`                                           |
| `default`            | The default value you want for this column. By default this is `null`                                | `{"default": 0}` `{"default": false}` `{"default": "foo"}` |
| `default_raw`        | The default value defined as a database function                                                     | `{"default_raw": "uuid_generate_v1()"}`                    |
| `after`              | (**MySQL Only**) Add a column after another column in the table                                      | `{"after": "created_at"}`                                  |
| `first`              | (**MySQL Only**) Add a column to the first position in the table                                     | `{"first": true}`                                          |

#### Disable Auto Timestamps

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
  // Disable auto-creation of created_at and updated_at columns
  t.DisableTimestamps()
}
```

or

```javascript
create_table("users", {timestamps: false}) {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

## Drop a Table

``` javascript
drop_table("table_name")
```

## Rename a Table

``` javascript
rename_table("old_table_name", "new_table_name")
```

## Add a Column

``` javascript
add_column("table_name", "column_name", "string", {})
```

See [above](#common-types) for more details on column types and options.

## Alter a column

``` javascript
change_column("table_name", "column_name", "string", {})
```

## Rename a Column

``` javascript
rename_column("table_name", "old_column_name", "new_column_name")
```

## Drop a Column

``` javascript
drop_column("table_name", "column_name")
```

## Composite Primary Keys

``` javascript
t.PrimaryKey("column_1", "column_2")
```

Please note that the `t.PrimaryKey` statement MUST be after the columns definitions.

## Add an Index

#### Supported Options:

* `name` - This defaults to `table_name_column_name_idx`
* `unique`

### Simple Index:

``` javascript
add_index("table_name", "column_name", {})
```

### Multi-Column Index:

``` javascript
add_index("table_name", ["column_1", "column_2"], {})
```

### Unique Index:

``` javascript
add_index("table_name", "column_name", {"unique": true})
```

### Index Names:

``` javascript
add_index("table_name", "column_name", {}) # name => table_name_column_name_idx
add_index("table_name", "column_name", {"name": "custom_index_name"})
```

## Rename an Index

``` javascript
rename_index("table_name", "old_index_name", "new_index_name")
```

## Drop an Index

``` javascript
drop_index("table_name", "index_name")
```

## Add a Foreign Key

```javascript
add_foreign_key("table_name", "field", {"ref_table_name": ["ref_column"]}, {
    "name": "optional_fk_name",
    "on_delete": "action",
    "on_update": "action",
})

```

#### Supported Options

* `name` - This defaults to `table_name_ref_table_name_ref_column_name_fk`
* `on_delete` - `CASCADE`, `SET NULL`, ...
* `on_update`

{{<warning>}}
**Note:** `on_update` and `on_delete` are not supported on CockroachDB yet.
{{</warning>}}

## Drop a Foreign Key

```javascript
drop_foreign_key("table_name", "fk_name", {"if_exists": true})
```

#### Supported Options

* `if_exists` - Adds `IF EXISTS` condition


## Raw SQL

``` javascript
sql("select * from users;")
```

## Execute an External Command

Sometimes during a migration you need to shell out to an external command.

```javascript
exec("echo hello")
```
