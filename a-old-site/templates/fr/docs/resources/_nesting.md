## Combiner des ressources

Pour simplifier la création de ressources hiérarchisées, Buffalo supporte la combinaison de ressources.

```go
type UsersResource struct {
  buffalo.Resource
}

type ImagesResource struct {
  buffalo.Resource
}

u := a.Resource("/users", UsersResource{})
u.Resource("/images", ImagesResource{})
```

Cela donne les routes suivantes :

```bash
$ buffalo routes
METHOD | PATH                                    | ALIASES | NAME              | HANDLER
------ | ----                                    | ------- | ----              | -------
GET    | /                                       |         | rootPath          | github.com/gobuffalo/coke/actions.HomeHandler
GET    | /users                                  |         | usersPath         | github.com/gobuffalo/coke/actions.UsersResource.List
POST   | /users                                  |         | usersPath         | github.com/gobuffalo/coke/actions.UsersResource.Create
GET    | /users/new                              |         | newUsersPath      | github.com/gobuffalo/coke/actions.UsersResource.New
GET    | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Show
PUT    | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Update
DELETE | /users/{user_id}                        |         | userPath          | github.com/gobuffalo/coke/actions.UsersResource.Destroy
GET    | /users/{user_id}/edit                   |         | editUserPath      | github.com/gobuffalo/coke/actions.UsersResource.Edit
GET    | /users/{user_id}/images                 |         | userImagesPath    | github.com/gobuffalo/coke/actions.ImagesResource.List
POST   | /users/{user_id}/images                 |         | userImagesPath    | github.com/gobuffalo/coke/actions.ImagesResource.Create
GET    | /users/{user_id}/images/new             |         | newUserImagesPath | github.com/gobuffalo/coke/actions.ImagesResource.New
GET    | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Show
PUT    | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Update
DELETE | /users/{user_id}/images/{image_id}      |         | userImagePath     | github.com/gobuffalo/coke/actions.ImagesResource.Destroy
GET    | /users/{user_id}/images/{image_id}/edit |         | editUserImagePath | github.com/gobuffalo/coke/actions.ImagesResource.Edit
```