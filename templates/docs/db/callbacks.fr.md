<% seoDescription("Attacher des callbacks aux opérations de base de données") %>
<% seoKeywords(["buffalo", "go", "golang", "callback", "base de données", "pop"]) %>

# Callbacks

Pop permet d'exécuter du code avant ou après les opérations de base de données. Pour cela, il suffit de définir des méthodes spécifiques sur vos modèles.

Par exemple, pour générer le hash du mot de passe d'un utilisateur, vous pourriez définir la méthode suivante :

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

Dans l'exemple ci-dessus, lorsque la méthode `Save` de la connection est appelée avec un `User`, la méthode `BeforeCreate` sera appelée avant d'écrire dans la base de données.

Les méthodes callback disponibles sont les suivantes :

* [BeforeSave](https://godoc.org/github.com/gobuffalo/pop#BeforeSaveable)
* [BeforeCreate](https://godoc.org/github.com/gobuffalo/pop#BeforeCreateable)
* [BeforeUpdate](https://godoc.org/github.com/gobuffalo/pop#BeforeUpdateable)
* [BeforeDestroy](https://godoc.org/github.com/gobuffalo/pop#BeforeDestroyable)
* [AfterSave](https://godoc.org/github.com/gobuffalo/pop#AfterSaveable)
* [AfterCreate](https://godoc.org/github.com/gobuffalo/pop#AfterCreateable)
* [AfterUpdate](https://godoc.org/github.com/gobuffalo/pop#AfterUpdateable)
* [AfterDestroy](https://godoc.org/github.com/gobuffalo/pop#AfterDestroyable)
* [AfterFind](https://godoc.org/github.com/gobuffalo/pop#AfterFindable)

<%= title("Contenu lié") %>

* [Modèles](/fr/docs/db/models) - Définir un modèle de base de données.