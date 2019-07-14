# Factorisation des requêtes

La factorisation est une façon de structurer vos requêtes lorsqu'elles partagent une composante commune. Prenons par exemple une librairie : celle-ci fournit des livres pour tout le monde, mais certaines éditions ne sont accessibles qu'aux clients disposant d'un compte de fidélité. Cela signifie que pour l'ensemble de la librairie, vous allez devoir filtrer les livres de manière à appliquer la règle « seuls les clients disposant d'un compte fidélité peuvent avoir accès à l'ensemble des livres ».

## La manière normale

Sans trop réfléchir, on peut écrire chaque requête en entier.

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// On cherche à récupérer la liste des livres
books := Books{}
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false").All(&books)
} else {
    // On récupère l'ensemble des livres sans filtre
    err = tx.All(&books)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// On récupère un livre en particulier
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

## La manière factorisée

En utilisant un *scope*, on peut réécrire nos requêtes de la manière suivante :

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// restrictedScope définit une requête de base qui applique notre contrainte partagée.
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
// On cherche à récupérer la liste des livres
books := Books{}

if err := tx.Scope(restrictedScope(registeredAccount)).All(&books); err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// On récupère un livre en particulier
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

Et voilà comment factoriser une (ou plusieurs) conditions communes entre plusieurs requêtes !
