<%= title("Has Many") %>

<%= codeTabs() { %>
```go
// models/person.go
type Person struct {
  ID        uuid.UUID `json:"id" db:"id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
  Name      string    `json:"name" db:"name"`
}
```

```fizz
// migrations/create_people.fizz
create_table("people", func(t) {
  t.Column("id", "uuid", {"primary": true})
  t.Column("name", "string", {})
})
```

```go
// models/pet.go
type Pet struct {
  ID        uuid.UUID `json:"id" db:"id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
  Name      string    `json:"name" db:"name"`
  PersonID  uuid.UUID `json:"person_id" db:"person_id"`
}
```

```fizz
// migrations/create_pets.fizz
create_table("pets", func(t) {
  t.Column("id", "uuid", {"primary": true})
  t.Column("name", "string", {})
  t.Column("person_id", "uuid", {})
})
```

```go
// seed a few records
models.DB.Transaction(func(tx *pop.Connection) error {
  person := &models.Person{Name: "Mark"}
  err := tx.Create(person)
  if err != nil {
    return errors.WithStack(err)
  }
  pet := &models.Pet{Name: "Ringo", PersonID: person.ID}
  return tx.Create(pet)
})
```

```go
// find pets for a person
models.DB.Transaction(func(tx *pop.Connection) error {
  person := &models.Person{}
  err := tx.First(person)
  if err != nil {
    return errors.WithStack(err)
  }
  pets := &models.Pets{}
  err = tx.BelongsTo(person).All(pets)
  if err != nil {
    return errors.WithStack(err)
  }
  fmt.Println(pets)
  return nil
})
```

```sql
// resulting SQL queries
SELECT people.created_at, people.id, people.name, people.updated_at FROM people AS people LIMIT 1
SELECT pets.created_at, pets.id, pets.name, pets.person_id, pets.updated_at FROM pets AS pets WHERE person_id = $1
```
<% } %>
