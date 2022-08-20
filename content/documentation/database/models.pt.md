---
name: Models
seoDescription: "Models"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "models"]
weight: 10
aliases:
  - /docs/db/models
  - /pt/docs/db/models
---

# Models

Pop, as an ORM, allows you to translate database tables into Go structs. This way, you can manipulate Go structs instead of writing SQL statements. The Go code managing this part is named "models", as a reference to the MVC architecture.

In this chapter, you'll learn how to work with models by hand; and how to improve your workflow using the provided generators.

## The Models Directory

Pop model files are stored in the `models` directory, at your project root (see [the Directory Structure](/documentation/getting_started/directory-structure) chapter for more info about the Buffalo way to organize your files).

This directory contains:

* A `models.go` file, which defines the common parts for every defined model. It also contains a pointer to the configured connection. Remember the code is your own, so you can place whatever you like here.
* Model definition files, one for each model (so one per database table you want to access this way).

## Define a Simple Model

A model file defines a mapping for the database table, validation methods and Pop callbacks if you want to add more model-related logic.

Let's take the following SQL table definition, and write a matching structure:

```sql
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    label character varying(255)
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);
```

We'll start by creating a new file in the `models` directory, called `soda.go` (the convention used here is to take the singular form of the word). In this file, we'll create the structure for the `sodas` table (the structure is singular too, since it will contain a single line of the table):

```go
package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
)

type Soda struct {
	ID                   uuid.UUID    `db:"id"`
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            time.Time    `db:"updated_at"`
	Label                nulls.String `db:"label"`
}
```

That's it! You don't need anything else to work with Pop! Note, for each table field, we defined a `pop` tag matching the field name, but it's not required. If you don't provide a name, Pop will use the name of the struct field to generate one.

## Using the Generator

{{< note >}}
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So everytime you want to use a command from `soda`, just execute `buffalo pop` instead.
{{< /note >}}

Writing the files by hand is not the most efficient way to work. Soda (and Buffalo, if you followed the chapter about Soda) provides a generator to help you:

```bash
$ soda g model --help
Generates a model for your database

Usage:
  soda generate model [name] [flags]

Aliases:
  model, m


Flags:
 -h, --help                    help for model
      --migration-type string   sets the type of migration files for model (sql or fizz) (default "fizz")
      --models-path string      the path the model will be created in (default "models")
  -s, --skip-migration   Skip creating a new fizz migration for this model.
      --struct-tag string       sets the struct tags for model (xml/json/jsonapi) (default "json")

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

You can remove generated model by running:

```bash
$ soda destroy model [name]
```

Or in short form:

```bash
$ soda d m [name]
```

## Nulls Handling

If you need to store `NULL` values in your table, you'll have to use special types: for instance, you can't store a `NULL` value if your type is `int`.

The [Go standard library](https://golang.org/pkg/database/sql) provides special types for that use case, like [`sql.NullBool`](https://golang.org/pkg/database/sql/#NullBool) or [`sql.NullInt64`](https://golang.org/pkg/database/sql/#NullInt64).

If you need more than what the standard library offers, you can use the [gobuffalo/nulls](https://github.com/gobuffalo/nulls) package which provides more nulls types and a better handling for JSON serialization and unserialization.

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password nulls.String
}
```

## Customize Models

### Mapping Model Fields

By default when trying to map a struct to a database table, Pop, will use the name of the field in the struct as the name of the column in the database.

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}
```

With the above struct it is assumed the column names in the database are `ID`, `Email`, and `Password`.

These column names can be changed by using the `db` struct tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Now the columns names are expected to be `id`, `email`, and `password`.

This is very similar to how [form binding](/docs/bind) works.

Any types can be used that adhere to the [Scanner](https://golang.org/pkg/database/sql/#Scanner) and [Valuer](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces, however, so that you don't have to write these yourself it is recommended you stick with the following types:

| Base type             | Nullable        | Slice/Array |
|:-----------------------|:---------------|:------------|
|int                    |nulls.Int        |slices.Int   |
|int32                  |nulls.Int32      | ------      |
|int64                  |nulls.Int64      | ------      |
|uint32                 |nulls.UInt32     | ------      |
|float32                |nulls.Float32    | ------      |
|float, float64         |nulls.Float64    |slices.Float |
|bool                   |nulls.Bool       | ------      |
|[]byte                 |nulls.ByteSlice  | ------      |
|string                 |nulls.String     |slices.String|
|uuid.UUID              |nulls.UUID       |slices.UUID  |
|time.Time              |nulls.Time       | ------      |
|map[string]interface{} | ---------       |slices.Map   |

{{< note >}}
**Note**: Any `slices.Map` typed fields will need to be initialized before `Bind`ing or accessing.
```go
widget := &models.Widget{Data: slices.Map{}}
```
{{< /note >}}

### Read Only Fields

It is often necessary to read a field from a database, but not want to write that field to the database. This can be done using the `rw` struct  tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

In this example all fields will be read **from** the database and all fields, **except** for `Password` will be able to write to the database.

### Write Only Fields

Write only fields are the reverse of read only fields. These are fields that you want to write to the database, but never retrieve. Again, this makes use of the `rw` struct tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Skipping Model Fields

Sometimes you need to let Pop know that certain field should not be stored in the database table. Perhaps it's just a field you use in-memory or other logical reason related with the application you're building.

The way you let Pop know about this is by using the `db` struct tag on your model and setting it to be `-` like the following example:

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```

As you may see the `Password` field is marked as `db:"-"` that means Pop will neither **_store_** nor **_retrieve_** this field from the database.

### Changing the Select Clause for a Column

The default, when trying to build the `select` query for a struct is to use all of the field names to build a query.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

The resulting `select` statement would look like this:

```sql
select id, email, password from users
```

We can change the statement for a column using the `select` tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" select:"password as p"`
}
```

The resulting `select` statement would look like this:

```sql
select id, email, password as p from users
```

### Using a Custom Table Name

Sometimes, you'll have to work with an existing schema, with the table names non-matching the Pop conventions. You can override this behavior, and provide a custom table name by implementing the [`TableNameAble`](https://godoc.org/github.com/gobuffalo/pop#TableNameAble) interface:

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}

// TableName overrides the table name used by Pop.
func (u User) TableName() string {
  return "my_users"
}
```

It is recommended to use a value receiver over a pointer receiver if the struct is used as a value anywhere in the code.

```go
// recommended:
func (u User) TableName() string {

// can cause issues:
func (u *User) TableName() string {
```

### UNIX Timestamps

{{< since "v4.7.0" >}}

If you define the `CreatedAt` and `UpdatedAt` fields in your model struct (and they are created by default when you use the model generator), Pop will manage them for you. It means when you create a new entity in the database, the `CreatedAt` field will be set to the current datetime, and `UpdatedAt` will be set each time you update an existing entity.

These fields are defined as time.Time, but now you can define them as `int` and handle them as UNIX timestamps.

```go
type User struct {
  ID        int    `db:"id"`
  CreatedAt int    `db:"created_at"`
  UpdatedAt int    `db:"updated_at"`
  FirstName string `db:"first_name"`
  LastName  string `db:"last_name"`
}
```

If you use fizz migrations, make sure to define these fields by yourself, and disable the default datetime timestamps:

```go
create_table("users") {
  t.Column("id", "int", {primary: true})
  t.Column("created_at", "int")
  t.Column("updated_at", "int")
  t.Column("first_name", "string")
  t.Column("last_name", "string")
  t.DisableTimestamps()
}
```

## Views Models

A [view](https://en.wikipedia.org/wiki/View_(SQL)) is a database collection object which stores the result of a query. Since this object acts as a read-only table, you can map it with Pop models just like a table.

If you want to use a model with more than one table, defining a view is probably the best solution for you.

### Example

The following example uses the PostgreSQL syntax. We'll start by creating two tables:

```sql
-- Create a sodas table
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    provider_id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);

-- Create a providers table
CREATE TABLE providers (
    id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE providers ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

-- Create a foreign key between the two tables
ALTER TABLE sodas ADD FOREIGN KEY (provider_id) REFERENCES providers(id);
```

Then create a view from the two tables:

```sql
CREATE VIEW sodas_with_providers AS
SELECT s.id, s.created_at, s.updated_at, p.label AS provider_label, s.label
FROM sodas s
LEFT JOIN providers p ON p.id = s.provider_id;
```

Since the view is considered as a table by Pop, let's finish by declaring a new model:

```go
type SodasWithProvider struct {
	ID            uuid.UUID `db:"id" rw:"r"`
	CreatedAt     time.Time `db:"created_at" rw:"r"`
	UpdatedAt     time.Time `db:"updated_at" rw:"r"`
	Label         string    `db:"label" rw:"r"`
	ProviderLabel string    `db:"provider_label" rw:"r"`
}
```

As we learned in this chapter, each attribute on the structure has a read-only tag `rw:"r"`. Since a view is a read-only object, it prevents any writing operation before hitting the database.

## Related Content

* [Migrations](/documentation/database/migrations) - Write database migrations.
* [Querying](/documentation/database/querying) - Query data from your database.
