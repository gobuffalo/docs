<% seoDescription("Models") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "models"]) %>

<%= h1("Models") %>

Pop, as an ORM, allows you to translate database tables into Go structs. This way, you can manipulate Go structs instead of writing SQL statements. The Go code managing this part is named "models", as a reference to the MVC architecture.

In this chapter, you'll learn how to work with models by hand; and how to improve your workflow using the provided generators.

<%= title("The models directory") %>

Pop model files are stored in the `models` directory, at your project root (see [the Directory Structure](/en/docs/directory-structure) chapter for more info about the Buffalo way to organize your files).

This directory contains:

* A `models.go` file, which defines the common parts for every defined model. It also contains a pointer to the configured connection. Remember the code is your own, so you can place whatever you like here.
* Model definition files, one for each model (so one per database table you want to access this way).

<%= title("Define a simple model") %>

A model file defines a mapping for the database table, validation methods and Pop callbacks if you want to add more model-related logic.

Let's take the following SQL table definition, and write a matching structure:

```sql
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);
```

We'll start by creating a new file in the `models` directory, called `soda.go` (the convention used here is to take the singular form of the word). In this file, we'll create the structure for the `sodas` table (the structure is singular too, since it will contain a single line of the table):

```go
package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

type Soda struct {
	ID                   uuid.UUID `db:"id"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	Label                string    `db:"label"`
}
```

That's it! You don't need anything else to work with Pop! Note, for each table field, we defined a `db` tag matching the field name, but it's not required. If you don't provide a name, Pop will use the name of the struct field to generate one.

<%= title("Using the generator") %>

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `db` namespace. So everytime you want to use a command from `soda`, just execute `buffalo db` instead.
<% } %>

Writing the files by hand is not the most efficient way to work. Soda (and Buffalo, if you followed the chapter about Soda) provides a generator to help you:

<%= partial("docs/db/model.md") %>

You can remove generated model by running:

```bash
$ soda destroy model [name]
```

Or in short form:

```bash
$ soda d m [name]
```

<%= title("Customize models") %>

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

Write only fields are the reverse of read only fields. These are fields that you want to write to the database, but never retreive. Again, this makes use of the `rw` struct tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Skipping Model Fields

Sometimes you need to let Pop know that certain field should not be stored in the database table. Perhaps it's just a field you use in-memory or other logical reason related with the application you're building.

The way you let Pop know about this is by usind the `db` struct tag on your model and setting it to be `-` like the following example:

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```

As you may see the `Password` field is marked as `db:"-"` that means Pop will neither **_store_** nor **_retreive_** this field from the database.

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