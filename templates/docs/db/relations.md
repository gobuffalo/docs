# Assocations and Relationships

Pop allows you to perform an eager loading for associations defined in a model. By using [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) method plus some struct field tags predefined in your model you can extract associated data from a model.

<%= vimeo("253683926") %>

<%= title("Example") %>

```go
type User struct {
  ID           uuid.UUID
  Email        string
  Password     string
  Books        Books     `has_many:"books" order_by:"title asc"`
  FavoriteSong Song      `has_one:"song" fk_id:"u_id"`
  Houses       Addresses `many_to_many:"users_addresses"`
}

type Book struct {
  ID      uuid.UUID
  Title   string
  Isbn    string
  User    User        `belongs_to:"user"`
  UserID  uuid.UUID
}

type Song struct {
  ID      uuid.UUID
  Title   string
  UserID  uuid.UUID   `db:"u_id"`
}

type Address struct {
  ID           uuid.UUID
  Street       string
  HouseNumber  int
}

type Addresses []Address
```

<%= title("Available Struct Tags") %>

Using the above [example](#example) code below is a list of available struct tags and how to use them.

* `has_many`: Will load all records from the `books` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.
* `belongs_to`: Will load a record from `users` table that have a column named `id` that matches with `Book.UserID` value.
* `has_one`: Will load a record from the `songs` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.
* `many_to_many`: Will load all records from the `addresses` table through the table `users_addresses`. Table `users_addresses` **MUST** define `address_id` and `user_id` columns to match `User.ID` and `Address.ID` values. You can also define a `fk_id` tag that will be used in the target association i.e. `addresses` table.
* `fk_id`: Defines the column name in the target association that matches model ID. In the example above `Song` has a column named `u_id` that represents id of `users` table. When loading `FavoriteSong`, `u_id` will be used instead of `user_id`.
* `order_by`: Used in `has_many` and `many_to_many` to indicate the order for the association when loading. The format to use is `order_by:"&lt;column_name> &lt;asc | desc>"`

<%= title("Eager Loading Associations") %>

The [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) method tells Pop to load the associations for a model when that model is loaded from the database.

```go
u := Users{}
err := tx.Eager().Where("name = 'Mark'").All(&u)  // preload all associations for user with name 'Mark', i.e Books, Houses and FavoriteSong
```

By default `Eager` will load all the assigned assocations for the model. To specify which associations should be loaded you can pass in the names of those fields to the `Eager` method and only those associations will be loaded.

```go
err  = tx.Eager("Books").Where("name = 'Mark'").All(&u) // preload only Books association for user with name 'Mark'.
```

<%= title("Loading Associations for an Existing Model") %>

The [`pop.Connection.Load()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Load) method takes a model struct, that has already been populated from the database, and an optional list of associations to load.

```go
tx.Load(u) // load all associations for user, i.e Books, Houses and FavoriteSong
tx.Load(u, "Books") // load only the Books associations for user
```

The `Load` method will not retreive the `User` from the database only its associations.
