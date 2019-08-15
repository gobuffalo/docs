## Associer des actions

Toutes les routes dans Buffalo sont branchées à des fonctions de type `buffalo.Handler`. La signature d'une telle fonction ressemble à ceci :

```go
func (c buffalo.Context) error {
  // do some work
}
```

Si vous connaissez déjà le **pattern MVC**, les fonctions `buffalo.Handler` gèrent la partie contrôleur : c'est l'endroit où l'on gère tout l'aspect logique de l'application. Les fonctions `buffalo.Handler` prennent en paramètre une struct `buffalo.Context`, qui contient tout le nécessaire sur la requête courante.

Consultez la page [Contexte](/fr/docs/context) pour mieux comprendre l'interface `buffalo.Context`.

##### Méthodes HTTP supportées

Buffalo supporte de base les méthodes HTTP suivantes :

* GET
* POST
* PUT
* PATCH
* DELETE
* OPTIONS
* HEAD

Vous pouvez également gérer toutes les méthodes HTTP d'un coup en utilisant `ANY`.

Pour associer un `buffalo.Handler` à une méthode HTTP, on procède de la sorte :

```go
a.GET("/some/path", SomeHandler)
a.POST("/some/path", func (c buffalo.Context) error {
  // do some work
})
// etc...
```

Comme vous pouvez le voir, il est possible d'utiliser des fonctions en ligne directement dans la déclaration de la route. Pour rendre les choses plus lisibles, il est toutefois préférable de séparer vos `buffalo.Handler` des routes et de les placer dans plusieurs fichiers. Par exemple, si vous avez plusieurs `buffalo.Handler` qui gèrent les utilisateurs, vous pouvez les grouper dans un fichier `users.go` dans le dossier [`actions`](/fr/docs/getting-started/directory-structure).
