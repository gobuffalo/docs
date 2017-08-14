<%= title("Has Many Through") %>

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
}
```

```fizz
// migrations/create_pets.fizz
create_table("pets", func(t) {
  t.Column("id", "uuid", {"primary": true})
  t.Column("name", "string", {})
})
```

```go
// models/pet_owner.go
type PetOwner struct {
  ID        uuid.UUID `json:"id" db:"id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
  PersonID  uuid.UUID `json:"person_id" db:"person_id"`
  PetID     uuid.UUID `json:"pet_id" db:"pet_id"`
}
```

```fizz
// migrations/create_pet_owners.fizz
create_table("pet_owners", func(t) {
  t.Column("id", "uuid", {"primary": true})
  t.Column("person_id", "uuid", {})
  t.Column("pet_id", "uuid", {})
})
```

```go
// some seed code to insert a few records
models.DB.Transaction(func(tx *pop.Connection) error {
  person := &models.Person{Name: "Mark"}
  err := tx.Create(person)
  if err != nil {
    return errors.WithStack(err)
  }
  pet := &models.Pet{Name: "Ringo"}
  err = tx.Create(pet)
  if err != nil {
    return errors.WithStack(err)
  }
  owner := &models.PetOwner{PersonID: person.ID, PetID: pet.ID}
  return tx.Create(owner)
})
```

```go
// some code to find pets for a person through pet_owners
models.DB.Transaction(func(tx *pop.Connection) error {
  person := &models.Person{}
  err := tx.First(person)
  if err != nil {
    return errors.WithStack(err)
  }
  pets := &models.Pets{}
  err = tx.BelongsToThrough(person, "pet_owners").All(pets)
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
SELECT pets.created_at, pets.id, pets.name, pets.updated_at FROM pets AS pets, pet_owners AS pet_owners WHERE pet_owners.person_id = $1 AND pets.id = pet_owners.pet_id
```
