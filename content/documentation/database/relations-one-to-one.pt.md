---
name: One to one associations
seoDescription: "Pop one to one associations"
seoKeywords: ["orm", "buffalo", "go", "golang", "pop", "one to one", "associations"]
weight: 52
aliases:
  - /documentation/database/onetoone
  - /docs/db/relations/onetoone
  - /pt/docs/db/relations/onetoone
---

# One to One Associations

In this chapter, you'll learn how to write a one to one [association](/documentation/database/relations/) in Pop.

## Tags

One to one associations work using a pair of tags:
* `belongs_to` for the model with the foreign key.
* `has_one` for the model without the foreign key.

## Example

```go
// Models

type Head struct {
  ID           int
  BodyID       int        `db:"body_id"`
  Body         *Body      `belongs_to:"body"`
}

type Body struct {
  ID           int
  Head         Head       `has_one:"head"`
}
```

```go
// Eager creation:
// Create a body with its head.
b := &models.Body{
    Head: models.Head{},
}

if err := tx.Eager().Create(b); err != nil {
    return err
}
```

```go
// Eager fetch all bodies with their head.
bodies = &models.Bodies{}

if err := c.Eager().All(bodies); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", bodies)
```

## Related Content

* [Associations with Pop: 1 to 1](https://blog.gobuffalo.io/associations-with-pop-1-to-1-592f02e2bdd8) - An article about 1 to 1 associations in Pop.
