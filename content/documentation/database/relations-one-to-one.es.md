---
name: Asociaciones Uno a Uno
seoDescription: "Asociaciones uno a uno en Pop"
seoKeywords: ["orm", "buffalo", "go", "golang", "pop", "one to one", "associations"]
weight: 52
aliases:
  - /documentation/database/onetoone
  - /docs/db/relations/onetoone
  - /es/docs/db/relations/onetoone
---

# Asociaciones Uno a Uno

En este capítulo, aprenderás a como escribir una [asociación](/es/documentation/database/relations/) uno a uno en Pop.

## Tags

Las asociaciones Uno a Uno usan un par de tags:

* `belongs_to` para el modelo con una llave foránea.
* `has_one` para un modelo sin llave foránea.

## Ejemplo

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
// Creación con Eager:
// Crear un `Body` con su `Head`.
b := &models.Body{
    Head: models.Head{},
}

if err := tx.Eager().Create(b); err != nil {
    return err
}
```

```go
// Buscar todos los `Bodies` con sus `Head` usando Eager.
bodies = &models.Bodies{}

if err := c.Eager().All(bodies); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", bodies)
```

## Contenido Relacionado

* [Associations with Pop: 1 to 1](https://blog.gobuffalo.io/associations-with-pop-1-to-1-592f02e2bdd8) - Un artículo sobre asociaciones 1 a 1 en Pop.
