<% seoDescription("Pop Models") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "models"]) %>

<%= h1("Pop Models") %>

Pop, as an ORM, allows you to translate database tables into Go structs. This way, you can manipulate Go structs instead of writing SQL statements.

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

We'll start by creating a new file in the `models` directory, called `soda.go` (the convention used here is to take the singular for of the word). In this file, we'll create the structure for the `sodas` table (the structure is singular too, since it will contain a single line of the table):

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

That's it! You don't need anything else to work with Pop! Note, for each table field, we defined a `db` tag matching the field name.

<%= title("Using the generator") %>

Writing the files by hand is not the most efficient way to work. Buffalo provides a generator to help you:

<%= partial("docs/db/model.md") %>