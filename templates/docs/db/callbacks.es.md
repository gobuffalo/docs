<% seoDescription("Adjuntar _callbacks_ a las operaciones de la base de datos") %>
<% seoKeywords(["buffalo", "go", "golang", "callback", "hook", "database", "pop"]) %>

# Callbacks

Pop proporciona un medio para ejecutar código antes y después de las operaciones de la base de datos. Esto se hace definiendo métodos específicos para sus modelos. Por ejemplo, para hash una contraseña de usuario, es posible que desee definir el siguiente método:

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return errors.WithStack(err)
  }

  u.Password = string(hash)

  return nil
}
```

En el ejemplo anterior, cuando se llama al método `save` de la conexión con` User`, el método `BeforeCreate`
se llamará antes de escribir en la base de datos.

Los _callbacks_ disponibles incluyen:

* [BeforeSave](https://godoc.org/github.com/gobuffalo/pop#BeforeSaveable)
* [BeforeCreate](https://godoc.org/github.com/gobuffalo/pop#BeforeCreateable)
* [BeforeUpdate](https://godoc.org/github.com/gobuffalo/pop#BeforeUpdateable)
* [BeforeDestroy](https://godoc.org/github.com/gobuffalo/pop#BeforeDestroyable)
* [AfterSave](https://godoc.org/github.com/gobuffalo/pop#AfterSaveable)
* [AfterCreate](https://godoc.org/github.com/gobuffalo/pop#AfterCreateable)
* [AfterUpdate](https://godoc.org/github.com/gobuffalo/pop#AfterUpdateable)
* [AfterDestroy](https://godoc.org/github.com/gobuffalo/pop#AfterDestroyable)
* [AfterFind](https://godoc.org/github.com/gobuffalo/pop#AfterFindable)

<%= title("Contenido relacionado") %>

* [Models](/en/docs/db/models) - Defina un modelo de base de datos.