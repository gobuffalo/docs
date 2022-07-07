---
name: Asociaciones Uno a Muchos
seoDescription: "Asociaciones uno a muchos en Pop"
seoKeywords: ["orm", "buffalo", "go", "golang", "pop", "one to many", "associations"]
weight: 53
aliases:
  - /documentation/database/onetomany
  - /docs/db/relations/onetomany
  - /en/docs/db/relations/onetomany
---

# Asociaciones Uno a Muchos

En este capítulo, aprenderás a como escribir una [asociación](/es/documentation/database/relations/) uno a muchos en Pop.

## Tags

Las asociaciones de uno a muchos funcionan usando un par de tags:

* `belongs_to` para el modelo con una llave foránea.
* `has_many` para un modelo sin llave foránea (el de la lista).

## Example

```go
// Models

type Fruit struct {
    ID     int   `json:"id,omitempty" db:"id"`
    TreeID int   `json:"-" db:"tree_id"`
    Tree   *Tree `json:"tree,omitempty" belongs_to:"tree"`
}

type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits"`
}
```

```go
// Creación usando Eager:
// Crear un árbol de manzana con 2 frutas.
t := &models.Tree{
    Name: "Apple tree",
    Fruits: []models.Fruit{
        {},
        {},
    },
}

if err := tx.Eager().Create(t); err != nil {
    return err
}
```

```go
// Buscar todos los `Tress` con sus `Fruits` usando Eager.
trees := &models.Trees{}

if err := c.Eager().All(trees); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", trees)
```

## Orden de asociación personalizado

Dado que `has_many` está asignado a una lista, probablemente desearás personalizar el orden de esa lista. La etiqueta `order_by` te permite indicar el orden de la asociación al cargarla:

```go
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits" order_by:"id desc"`
}
```

La forma de usarla es `order_by:"<column_name> <asc | desc>"`.

## Personalizar búsqueda de Llaves foráneas.

Por defecto, `has_many` buscará registros relacionados usando una convención para la columna de llave foránea. En nuestro ejemplo anterior, la tabla `fruits` (asignada a la estructura `Fruit`) contiene una columna de llave foránea `tree_id` que hace referencia al ID del `tree` al que está asociada `fruit`.

Puedes usar el tag `fk_id` para personalizar esta columna de llave foránea:

```go
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits" fk_id:"custom_tree_id"`
}
```

Aquí, la relación se buscará utilizando la columna `custom_tree_id` en la tabla `fruits`, en lugar de la predeterminada `tree_id`.

Esto puede ser realmente útil cuando tienes estructuras con varios campos que apuntan al mismo modelo:

```go
type Player struct {
    ID            int     `json:"id" db:"id"`
    Name          string  `json:"name" db:"name"`
    CurrentBandID int     `json:"current_band_id" db:"current_band_id"`
    FormerBandID  int     `json:"former_band_id" db:"former_band_id"`
}

type Band struct {
    ID             int      `json:"id" db:"id"`
    Name           string   `json:"name" db:"name"`
    CurrentPlayers []Player `json:"current_players,omitempty" has_many:"players" fk_id:"current_band_id"`
    FormerPlayers  []Player `json:"former_players,omitempty" has_many:"players" fk_id:"former_band_id"`
}
```

## Contenido Relacionado

* [Associations with Pop: 1 to n](https://blog.gobuffalo.io/associations-with-pop-1-to-n-2fb3e1c3833f) - Un artículo sobre asociaciones 1 a n en Pop.
