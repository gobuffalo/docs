<% seoDescription("How to manage associations and relationships with Pop?") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "associations", "relations", "entity"]) %>

# Associations and Relationships

Associations are the Pop way to define **a relation between two objects in the database**. In this chapter, you'll learn how to define associations using struct tags; and how to manipulate them with the `Eager()` modifier.

<%= vimeo("253683926") %>

## Example

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

type Books []Book
type Addresses []Address
```

## Available Struct Tags

Using the above [example](#example) code below is a list of available struct tags and how to use them.

* `has_many`: Will load all records from the `books` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.
* `belongs_to`: Will load a record from `users` table that have a column named `id` that matches with `Book.UserID` value.
* `has_one`: Will load a record from the `songs` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.
* `many_to_many`: Will load all records from the `addresses` table through the table `users_addresses`. Table `users_addresses` **MUST** define `address_id` and `user_id` columns to match `User.ID` and `Address.ID` values. You can also define a `fk_id` tag that will be used in the target association i.e. `addresses` table.
* `fk_id`: Defines the column name in the target association that matches model ID. In the example above `Song` has a column named `u_id` that represents id of `users` table. When loading `FavoriteSong`, `u_id` will be used instead of `user_id`.
* `order_by`: Used in `has_many` and `many_to_many` to indicate the order for the association when loading. The format to use is `order_by:"&lt;column_name> &lt;asc | desc>"`

## Eager Loading Associations

The [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) method tells Pop to load the associations for a model when that model is loaded from the database.

```go
u := Users{}
err := tx.Eager().Where("name = 'Mark'").All(&u)  // preload all associations for user with name 'Mark', i.e Books, Houses and FavoriteSong
```

By default `Eager` will load all the assigned associations for the model. To specify which associations should be loaded you can pass in the names of those fields to the `Eager` method and only those associations will be loaded.

```go
err  = tx.Eager("Books").Where("name = 'Mark'").All(&u) // preload only Books association for user with name 'Mark'.
```

## Loading Associations for an Existing Model

The [`pop.Connection.Load()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Load) method takes a model struct, that has already been populated from the database, and an optional list of associations to load.

```go
tx.Load(&u) // load all associations for user, i.e Books, Houses and FavoriteSong
tx.Load(&u, "Books") // load only the Books associations for user
```

The `Load` method will not retreive the `User` from the database only its associations.

## Flat Nested Creation

Pop allows you to create the models and their associations with other models in one step by default. You no longer need to create every association separately anymore. Pop will even create join table records for `many_to_many` associations.

Assuming the following pieces of pseudo-code:

```go
book := Book{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}
tx.Create(&book)
song := Song{Title: "Don't know the title"}
tx.Create(&song)
addr := Address{HouseNumber: 1, Street: "Golang"}
tx.Create(&addr)

user := User{
  Name: "Mark Bates",
  Books: Books{Book{ID: book.ID}},
  FavoriteSong: song,
  Houses: Addresses{
    addr,
  },
}
```

```go
err := tx.Create(&user)
```

1. It will notice `Books` is a `has_many` association and it will realize that to actually update each book it will need to get the `User ID` first. So, it proceeds to store first `User` data so it can retrieve an **ID** and then use that ID to fill `UserID` field in every `Book` in `Books`. It updates all affected books in the database using their `ID`s to target them.

2. `FavoriteSong` is a `has_one` association and it uses same logic described in `has_many` association. Since `User` data was previously saved before updating all affected books, it already knows that `User` got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then updated in the database.

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables in this case: `users` and `addresses`. Because `User` was already stored, it already has its `ID`.  It will then use the `ID`s passed with the `Addresses` to create the coresponding entries in the join table.

For a `belongs_to` association like shown in the example below, it fill its `UserID` field before be saved in database.

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: user,
}
```

```go
tx.Create(&book)
```

## Eager Creation

Pop also allows you to create models and embed the creation of their associations in one step as well.

Assuming the following pieces of pseudo-code:

```go
user := User{
  Name: "Mark Bates",
  Books: Books{{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}},
  FavoriteSong: Song{Title: "Don't know the title"},
  Houses: Addresses{
    Address{HouseNumber: 1, Street: "Golang"},
  },
}
```

```go
err := tx.Eager().Create(&user)
```

1. It will notice `Books` is a `has_many` association and it will realize that to actually store every book it will need to get the `User ID` first. So, it proceeds to store/create first the `User` data so it can retrieve an **ID** and then use that ID to fill `UserID` field in every `Book` in `Books`. Later it stores all books in the database.

2. `FavoriteSong` is a `has_one` association and it uses same logic described in `has_many` association. Since `User` data was previously saved before creating all books, it already knows that `User` got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then stored in the database.

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables in this case: `users` and `addresses`. It will need to store all addresses first in `addresses` table before saving them in the many to many(join) table. Because `User` was already stored, it already have an `ID`. * This is a special case to deal with, since this behavior is different from all other associations, it is solved by implementing the `AssociationCreatableStatement` interface, all other associations by default implement the `AssociationCreatable` interface.

For a `belongs_to` association like shown in the example below, it will need first to create `User` to retrieve **ID** value and then fill its `UserID` field before be saved in database.

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: User{
        Name: nulls.NewString("Larry"),
   },
}
```

```go
tx.Eager().Create(&book)
```

In the case that you feed the eager create with associate models that already exist it will, instead of creating duplicates of them or updating the contents of them, simply create/update the associations with them.

## Related Content

* [Associations with Pop: 1 to 1](https://blog.gobuffalo.io/associations-with-pop-1-to-1-592f02e2bdd8) - An article about 1 to 1 associations in Pop.
* [Associations with Pop: 1 to n](https://blog.gobuffalo.io/associations-with-pop-1-to-n-2fb3e1c3833f) - An article about 1 to n associations in Pop.
