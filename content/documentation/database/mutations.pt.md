---
name: Mutations
seoDescription: "Handle mutations with Pop"
weight: 20
aliases:
  - /docs/db/mutations
  - /pt/docs/db/mutations
---

# Mutations

In this chapter, you'll learn how to create, update and delete objects from you database using Pop.

## Create

### Without validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Create a fruit without running validations
err := c.Create(fruit)
```

### With validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Run validations and create if all validations passed
vErrors, err := c.ValidateAndCreate(fruit)
```

## Update

### Without validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Update a fruit without running validations
err := c.Update(fruit)
```

### With validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Run validations and update if all validations passed
vErrors, err := c.ValidateAndUpdate(fruit)
```

## Save

Save checks for the ID in you model: if the ID is the zero value of the type (so for example if it's an `int` and its value is `0`), `Save` calls `Create`.
Otherwise, it calls `Update`.

### Without validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Create a fruit without running validations
err := c.Save(fruit)
```

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Update a fruit without running validations
err := c.Save(fruit)
```

### With validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Run validations and create if all validations passed
vErrors, err := c.ValidateAndSave(fruit)
```

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Run validations and update if all validations passed
vErrors, err := c.ValidateAndSave(fruit)
```

## Delete

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Destroy the fruit
err := c.Destroy(fruit)
```

## Next Steps

* [Querying](/documentation/database/querying) - Fetch the data you inserted in the database.
* [Associations and Relationships](/documentation/database/relations) - Handle relations between models.
