---
name: Raw Queries
seoDescription: "Writing raw queries with Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "raw", "query", "custom"]
weight: 22
aliases:
  - /docs/db/raw-queries
  - /pt/docs/db/raw-queries
---

# Raw Queries

Sometimes you'll need to write a custom query instead of letting Pop generate it for you. In this chapter, you'll learn how to write raw SQL queries using Pop.

## Writing a Raw Query

### Select

```go
player := Player{}
q := db.RawQuery("SELECT * FROM players WHERE id = ?", 1)
err := q.Find(&player, id)
```

### Update

```go
err := db.RawQuery("UPDATE players SET instrument = ? WHERE id = ?", "guitar", 1).Exec()
```

### Delete

```go
err := db.RawQuery("DELETE FROM players WHERE id = ?", 1).Exec()
```

## Tokens Syntax

With `RawQuery`, you can continue to use the `?` tokens to secure your input values. You don't need to use the token syntax for your underlying database.
