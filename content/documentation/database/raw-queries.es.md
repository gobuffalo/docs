---
name: Consultas SQL nativo
seoDescription: "Escribir consultas SQL nativocon Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "raw", "query", "custom"]
weight: 22
aliases:
  - /docs/db/raw-queries
  - /es/docs/db/raw-queries
---

# Consultas en SQL nativo

A veces necesitarás escribir una consulta personalizada en lugar de dejar que Pop la genere por ti. En este capítulo, aprenderás a escribir consultas SQL nativo utilizando Pop.

## Escribiendo SQL nativo

### Select

```go
player := Player{}

q := tx.RawQuery("SELECT * FROM players WHERE id = ?", 1)
err := q.First(&player)
```

### Update

```go
err := tx.RawQuery("UPDATE players SET instrument = ? WHERE id = ?", "guitar", 1).Exec()
```

### Delete

```go
err := tx.RawQuery("DELETE FROM players WHERE id = ?", 1).Exec()
```

## Sintaxis de tokens

Con `RawQuery`, puedes seguir usando el token `?` para asegurar sus valores de entrada. No necesitas usar la sintaxis de token para su base de datos subyacente.
