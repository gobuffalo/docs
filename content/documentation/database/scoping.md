---
name: Scoping
weight: 31
aliases:
  - /docs/db/scoping
  - /en/docs/db/scoping
---
# Scoping

Scoping is a way to structure your DB calls, when it needs the same "base" query. Let's say you want to create a book store: the store provides books for everyone, but some special editions are reserved to customers with a registered account. It means that for the whole store, you'll need to filter the books, so the "guest" customers can only see the restricted list of books.

## The usual way

A "naive" way can be writing each full query.

```go
type Book struct {
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

## The scoped way

The scope factorizes the common part of the query:

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// restrictedScope defines a base query which shares the common constraint.
func restrictedScope(registeredAccount bool) pop.ScopeFunc {
  return func(q *pop.Query) *pop.Query {
    if !registeredAccount {
      return q
    }
    return q.Where("is_restricted = false")
  }
}
```

```go
// Get available books list
books := Books{}

if err := tx.Scope(restrictedScope(registeredAccount)).All(&books); err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Get a specific book
book := Book{}
bookID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
tx := c.Value("tx").(*pop.Connection)

var err error

if err := tx.Scope(restrictedScope(registeredAccount)).Find(&book, bookID) != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", book)
}
```

See how we factorized the common restriction for each query, using the `restrictedScope` function?
