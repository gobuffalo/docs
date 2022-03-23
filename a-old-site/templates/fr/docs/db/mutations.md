<% seoDescription("Modifier ses données avec Pop") %>

# Mutations

Dans ce chapitre, vous allez apprendre à créer, modifier et supprimer des objets de votre base de données en utilisant Pop.

## Créer

### Sans validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Créer un fruit sans valider le modèle
err := c.Create(fruit)
```

### Avec validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Valider le modèle, puis créer le fruit
vErrors, err := c.ValidateAndCreate(fruit)
```

## Modifier

### Sans validation
```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Modifier un fruit sans valider le modèle
err := c.Update(fruit)
```

### Avec validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{}
// Valider le modèle, puis modifier le fruit en base.
vErrors, err := c.ValidateAndUpdate(fruit)
```

## Sauvegarder

`Save` vérifie l'ID de votre modèle : si cet ID est à la valeur 0 du type (par exemple, si c'est un `int` et que sa valeur est `0`), `Save` appelle la méthode `Create`.
Sinon, elle appelle `Update`.

### Sans validation
```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Créer un fruit sans validation
err := c.Save(fruit)
```

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Modifier un fruit sans validation
err := c.Save(fruit)
```

### Avec validation

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 0}
// Valider le fruit, puis le créer en base
vErrors, err := c.ValidateAndSave(fruit)
```

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Valider le fruit, puis le modifier en base
vErrors, err := c.ValidateAndSave(fruit)
```

## Supprimer

```go
c, err := pop.Connect("development")
// [...]
fruit := &models.Fruit{ID: 1}
// Supprimer le fruit
err := c.Destroy(fruit)
```

## Prochaines étapes

* [Requêtage](/fr/docs/db/querying) - Récupérer la donnée insérée en base.
* [Associations et relations](/fr/docs/db/relations) - Gérer les relations entre les modèles.