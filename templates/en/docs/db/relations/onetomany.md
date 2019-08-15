<% seoDescription("Pop one to many associations") %>
<% seoKeywords(["orm", "buffalo", "go", "golang", "pop", "one to many", "associations"]) %>

<%= h1("One to many associations") %>

In this chapter, you'll learn how to write a one to many [association](/en/docs/db/relations/) in Pop.

## Tags

One to one associations work using a pair of tags:
* `belongs_to` for the model with the foreign key.
* `has_many` for the model without the foreign key (the one with the slice).

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
// Eager creation:
// Create an apple tree with 2 fruits.
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
// Eager fetch all the trees with their fruits.
trees := &models.Trees{}

if err := c.Eager().All(trees); err != nil {
    log.Printf("err: %v", err)
    return
}

log.Printf("eager fetch: %v", trees)
```

## Custom association order

Since `has_many` is mapped to a slice, you'll probably want to customize the order of this slice. `has_many` tag allows you to indicate the order for the association when loading it:

```go
type Tree struct {
    ID     int     `json:"id" db:"id"`
    Name   string  `json:"name" db:"name"`
    Fruits []Fruit `json:"fruits,omitempty" has_many:"fruits" order_by:"id desc"`
}
```

The format to use is `order_by:"<column_name> <asc | desc>"`.

## Related Content

* [Associations with Pop: 1 to n](https://blog.gobuffalo.io/associations-with-pop-1-to-n-2fb3e1c3833f) - An article about 1 to n associations in Pop.