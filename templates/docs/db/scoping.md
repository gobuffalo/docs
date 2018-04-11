# Scoping

Scoping is a way to structure your DB calls, when it needs the same "base" query. Let's say you want to create a book store: the store provides books for everyone, but some special editions are reserved to customers with a registered account. It means that for the whole store, you'll need to filter the books, so the "guest" customers can only see the restricted list of books.

<%= title("The usual way") %>

A "naive" way can be writing each full query.

```go
type struct Book {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// Get available books list
books := Books{}
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false").All(&books)
} else {
    // Create an empty query
    err = tx.All(&books)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Get a specific book
book := Book{}
bookID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false AND id = ?", bookID).First(&book)
} else {
    err = tx.Find(&book, bookID)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", book)
}
```

<%= title("The scoped way") %>

The scope factorizes the common part of the query:

```go
type struct Book {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
func restrictedScope(c buffalo.Context, registeredAccount bool) *pop.Query {
    tx := c.Value("tx").(*pop.Connection)

    if registeredAccount {
        return tx.Where("is_restricted = false")
    }
    // Create an empty query
    return tx.Q()
}
```

```go
// Get available books list
books := Books{}

if err := restrictedScope(c).All(&books); err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Get a specific book
book := Book{}
bookID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
tx := c.Value("tx").(*pop.Connection)

var err error

if err := restrictedScope(c).Find(&book, bookID) != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", book)
}
```

See how we factorized the common restriction for each query, using the `restrictedScope` function?