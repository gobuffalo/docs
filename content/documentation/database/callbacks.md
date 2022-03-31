---
name: Callbacks
seoDescription: "Attach callbacks to database operations"
seoKeywords: ["buffalo", "go", "golang", "callback", "hook", "database", "pop"]
---

# Callbacks

Pop provides a means to execute code before and after database operations. This is done by defining specific methods for your models.

For example, to hash a user password you may want to define the following method:

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

In the above example, when the connection's `Save` method is called with a `User`, the `BeforeCreate` method
will be called before writing to the database.

The available callbacks include:

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

## Related Content

* [Models](/en/docs/db/models) - Define a database model.