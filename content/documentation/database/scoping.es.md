---
name: Scoping
weight: 31
aliases:
  - /docs/db/scoping
  - /en/docs/db/scoping
---
# Scoping

El Scoping es una forma de estructurar las llamadas a la base de datos, cuando se nedesita la misma consulta "base". Supongamos que necesitas crear una librería: la tienda ofrece libros para todos, pero algunas ediciones especiales están reservadas para clientes con una cuenta registrada. Esto significa que para toda la tienda, necesitaras filtrar los libros, para que los clientes "invitados" solo puedan ver la lista de libros.

## La forma habitual

Una forma "inocente" puede ser escribir cada consulta completa.

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// Obtener la lista de libros disponibles
books := Books{}
tx := c.Value("tx").(*pop.Connection)

var err error

if !registeredAccount {
    err = tx.Where("is_restricted = false").All(&books)
} else {
    // Crear un query vacío
    err = tx.All(&books)
}

if err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Obtener un libro específico
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

## La forma Scope

El scope factoriza la parte común de la consulta:

```go
type Book struct {
    ID         uuid.UUID `json:"id" db:"id"`
    Label      string    `json:"label" db:"label"`
    Restricted bool      `json:"is_restricted" db:"is_restricted"`
}

type Books []Book
```

```go
// restrictedScope define una consulta base la cual comparte la restricción común.
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
// Obtener la lista de libros disponibles
books := Books{}

if err := tx.Scope(restrictedScope(registeredAccount)).All(&books); err != nil {
    fmt.Printf("ERROR: %v\n", err)
} else {
    fmt.Printf("%v\n", books)
}

// Obtener un libro específico
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

Ves cómo factorizamos la restricción común para cada consulta, usando la función `restrictedScope`?
