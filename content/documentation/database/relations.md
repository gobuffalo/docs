---
seoDescription: "How to manage associations and relationships with Pop?"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "associations", "relations", "entity"]
name: Associations and Relationships
---

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

* `has_many`: This tag is used to describe [one-to-many](https://en.wikipedia.org/wiki/One-to-many_(data_model)) relationships in the database. In the example, `User` type defines a one-to-many relation with `Books` slice type through the use of `has_many` tag, meaning a `User` can own many `Books`. When querying to the database, Pop will load all records from the `books` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` value.

* `belongs_to`: This tag is used to describe the owner in the relationship. An owner represents a highly coupled dependency between the model and the target association field where `belongs_to` tag was defined. This tag is mostly used to indicate that model owns its "existence" to the association field with `belongs_to`. In the example above, `Book` type use `belongs_to` to indicate that it is owned by a `User` type. When querying to the database, Pop will load a record from the `users` table with `id` that matches with `Book.UserID` value.

* `has_one`: This tag is used to describe [one-to-one](https://en.wikipedia.org/wiki/One-to-one_(data_model)) relationships in the database. In the example above, there is only one `FavoriteSong` within all songs records that `User` type like the most. When querying to the database, Pop will load a record from the `songs` table that have a column named `user_id`, or the column specified with `fk_id` that matches the `User.ID` field value.

* `many_to_many`: This tag is used to describe [many-to-many](https://en.wikipedia.org/wiki/Many-to-many_(data_model)) relationships in the database. In the example above, the relationship between `User` type and `Addresses` slice type exists to indicate an `User` can own many `Houses` and a `House` can be owned by many `Users`. It is important to notice that value for `many_to_many` tag is the associative table that connects both sides in the relationship; in the example above this value is defined as `users_addresses`. When querying to the database, Pop will load all records from the `addresses` table through the associative table `users_addresses`. Table `users_addresses` **MUST** define `address_id` and `user_id` columns to match `User.ID` and `Address.ID` field values. You can also define a `fk_id` tag that will be used in the target association i.e. `addresses` table.

* `fk_id`: This tag can be used to define the column name in the target association that matches model ID. In the example above, `Song` has a column named `u_id` that references the id of the `users` table. When loading `FavoriteSong`, `u_id` column will be used instead of `user_id`.

* `order_by`: This tag can be used in combination with `has_many` and `many_to_many` tags to indicate the order for the association when loading. The format to use is `order_by:"&lt;column_name> &lt;asc | desc>"`

## Loading Associations
Pop currently provides two modes for loading associations; each mode will affect the way pop loads associations and queries to the database.

[Eager](#eager-mode). Default mode. By enabling this mode, pop will perform "n" queries for every association defined in the model. This means more hits to the database in order to not affect memory use.

[EagerPreload](#eagerpreload-mode). Optional mode. By enabling this mode, pop will perform one query for every association defined in the model. This mode will hit the database with a reduced frequency by sacrifing more memory space.

* `pop.SetEagerMode`: Pop allows enabling any of these modes globally which will affect **ALL** queries handle performance. Use `EagerDefault` or `EagerPreload` as parameter to activate any of these modes.


* `tx.EagerPreload | q.EagerPreload`: Pop allows developers to take control in which situations they want Pop to perform any of these modes when necessary. This method will activate `EagerPreload` mode only for the query in action.

* `tx.Eager | q.Eager`: Pop allows developers to take control in which situations they want Pop to perform any of these modes when necessary. This method will activate `Eager` mode only for the query in action.


## Eager Mode

The [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) method tells Pop to load the associations for a model once that model is loaded from the database. This mode will perform "n" queries for every association defined in the model.

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
```

```go
u := Users{}
err := tx.Eager().All(&u)  // loads all associations for every user registered, i.e Books, Houses and FavoriteSong
```

`Eager` mode will:
 1. Load all users.  
```
 SELECT * FROM users;
```
2. Iterate on every user and load its associations:  
```
 SELECT * FROM books WHERE user_id=1)
```  
```
 SELECT * FROM books WHERE user_id=2)  
```  
```
 SELECT * FROM books WHERE user_id=3)
``` 

## EagerPreload Mode
The [`pop.Connection.EagerPreload()`](https://github.com/gobuffalo/pop/pull/146/files#diff-f49e947ec94f65964b0845af2b62845aR180) method tells Pop to load the associations for a model once that model is loaded from the database. This mode will hit the database with a reduced frequency by sacrifing more memory space.

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
```

```go
u := Users{}
err := tx.EagerPreload().All(&u)  // loads all associations for every user registered, i.e Books, Houses and FavoriteSong
```

`EagerPreload` mode will:
 1. Load all users.  
 ```
  SELECT * FROM users;
 ```
2. Load associations for all users in one single query.  
```
  SELECT * FROM books WHERE user_id IN (1,2,3))
``` 

## Load Specific Associations
By default `Eager` and `EagerPreload` will load all the assigned associations for the model. To specify which associations should be loaded you can pass in the names of those fields to the `Eager` or `EagerPreload` methods and only those associations will be loaded.

```go
err  = tx.Eager("Books").Where("name = 'Mark'").All(&u) // load only Books association for user with name 'Mark'.
// OR
err  = tx.EagerPreload("Books").Where("name = 'Mark'").All(&u) // load only Books association for user with name 'Mark'.
```

Pop also allows you to eager load nested associations by using the `.` character to concatenate them. Take a look at the example below.

```go
// will load all Books for u and for every Book will load the user which will be the same as u.
tx.Eager("Books.User").First(&u)
// OR
tx.EagerPreload("Books.User").First(&u)
```

```go
// will load all Books for u and for every Book will load all Writers and for every writer will load the Book association.
tx.Eager("Books.Writers.Book").First(&u)
// OR
tx.EagerPreload("Books.Writers.Book").First(&u)
```

```go
// will load all Books for u and for every Book will load all Writers. And Also it will load the favorite song for user.
tx.Eager("Books.Writers").Eager("FavoriteSong").First(&u)
// OR
tx.EagerPreload("Books.Writers").EagerPreload("FavoriteSong").First(&u)
```

## Loading Associations for an Existing Model

The [`pop.Connection.Load()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Load) method takes a model struct, that has already been populated from the database, and an optional list of associations to load.

```go
tx.Load(&u) // load all associations for user, i.e Books, Houses and FavoriteSong
tx.Load(&u, "Books") // load only the Books associations for user
```

The `Load` method will not retrieve the `User` from the database, only its associations.

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

2. `FavoriteSong` is a `has_one` association and it uses same logic described in `has_many` association. Since `User` data was previously saved before updating all affected books, it already knows that `User` has got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then updated in the database.

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables in this case: `users` and `addresses`. Because `User` was already stored, it already has its `ID`.  It will then use the `ID`s passed with the `Addresses` to create the coresponding entries in the join table.

For a `belongs_to` association like shown in the example below, it fills its `UserID` field before being saved in the database.

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

1. It will notice `Books` is a `has_many` association and it will realize that to actually store every book it will need to get the `User ID` first. So, it proceeds to first store/create the `User` data so it can retrieve an **ID** and then use that ID to fill the `UserID` field in every `Book` in `Books`. Later it stores all books in the database.

2. `FavoriteSong` is a `has_one` association and it uses same logic described in the `has_many` association. Since `User` data was previously saved before creating all books, it already knows that `User` has got an `ID` so it fills its `UserID` field with that value and `FavoriteSong` is then stored in the database.

3. `Houses` in this example is a `many_to_many` relationship and it will have to deal with two tables, in this case: `users` and `addresses`. It will need to store all addresses first in the `addresses` table before saving them in the many to many(join) table. Because `User` was already stored, it already has an `ID`. * This is a special case to deal with, since this behavior is different from all other associations, it is solved by implementing the `AssociationCreatableStatement` interface, all other associations by default implement the `AssociationCreatable` interface.

For a `belongs_to` association like shown in the example below, it will need to first create the `User` to retrieve its **ID** value and then fill its `UserID` field before being saved in the database.

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

In the case where you feed the eager create with associated models that already exist, it will, instead of creating duplicates of them or updating the contents of them, simply create/update the associations with them.

## Next steps

* [One to one relations](/en/docs/db/relations/onetoone/)
* [One to many relations](/en/docs/db/relations/onetomany/)
