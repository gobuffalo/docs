# Callbacks

Pop provides a means to execute code before and after database operations. This is done by defining specific methods for your models. For example, to hash a user password you may want to define the following method:

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}

func (u *User) BeforeSave(tx *pop.Connection) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return errors.WithStack(err)
  }

  u.Password = string(hash)

  return nil
}
```

In the above example, when the connection's `Save` method is called with a `User`, the `BeforeSave` method
will be called before writing to the database. The available callbacks include:

* BeforeSave
* BeforeCreate
* BeforeUpdate
* BeforeDestroy
* AfterSave
* AfterCreate
* AfterUpdate
* AfterDestroy
* AfterFind
