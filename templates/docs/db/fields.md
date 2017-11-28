# Managing Model Fields

<%= title("Mapping Model Fields") %>

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

<%= title("Read Only Fields") %>

It is often necessary to read a field from a database, but not want to write that field to the database. This can be done using the `rw` struct  tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

In this example all fields will be read **from** the database and all fields, **except** for `Password` will be able to write to the database.

<%= title("Write Only Fields") %>

Write only fields are the reverse of read only fields. These are fields that you want to write to the database, but never retreive. Again, this makes use of the `rw` struct tag.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

<%= title("Skipping Model Fields") %>

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

<%= title("Changing the Select Clause for a Column") %>

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

