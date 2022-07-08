---
name: Callbacks
seoDescription: "Attach callbacks to database operations"
seoKeywords: ["buffalo", "go", "golang", "callback", "hook", "database", "pop"]
weight: 30
aliases:
  - /docs/db/callbacks
  - /es/docs/db/callbacks
---

# Callbacks

Pop proporciona una forma para ejecutar código antes y despues de una operación de base de datos. Esto se hace definiendo métodos específicos para tus modelos.

Por ejemplo, para codificar una contraseña de usuario, es posible que desees definir el siguiente método:

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}

func (u User) BeforeCreate(tx *pop.Connection) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return errors.WithStack(err)
  }

  u.Password = string(hash)

  return nil
}
```

En el ejemplo anterior, cuando se llama al método `Save` de la conexión con un `User`, se llamará al método `BeforeCreate` antes de escribir en la base de datos.

La lista de callbacks disponibles incluye:

* [BeforeSave](https://godoc.org/github.com/gobuffalo/pop#BeforeSaveable)
* [BeforeCreate](https://godoc.org/github.com/gobuffalo/pop#BeforeCreateable)
* [BeforeUpdate](https://godoc.org/github.com/gobuffalo/pop#BeforeUpdateable)
* [BeforeDestroy](https://godoc.org/github.com/gobuffalo/pop#BeforeDestroyable)
* [BeforeValidate](https://godoc.org/github.com/gobuffalo/pop#BeforeValidateable)
* [AfterSave](https://godoc.org/github.com/gobuffalo/pop#AfterSaveable)
* [AfterCreate](https://godoc.org/github.com/gobuffalo/pop#AfterCreateable)
* [AfterUpdate](https://godoc.org/github.com/gobuffalo/pop#AfterUpdateable)
* [AfterDestroy](https://godoc.org/github.com/gobuffalo/pop#AfterDestroyable)
* [AfterFind](https://godoc.org/github.com/gobuffalo/pop#AfterFindable)

## Contenido Relacionado

* [Modelos](/es/documentation/database/models) - Define un modelo de base de datos.
