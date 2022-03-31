<% seoDescription("Tasks") %>
<% seoKeywords(["buffalo", "go", "golang", "tâches", "scripts", "grift"]) %>

<%= h1("Tâches") %>

Les tâches sont des petits scripts qui sont souvent nécessaires lorsque l'on écrit une application.
Ces tâches peuvent par exemple remplir la base de données avec des données de test, parser un fichier de log, ou encore empaqueter une nouvelle version. Buffalo utilise la bibliothèque [grift](https://github.com/markbates/grift) pour faciliter l'écriture de ces tâches.

<%= vimeo("213096302") %>

## Écrire des tâches

Les tâches doivent toutes se trouver dans le package `grifts`. Une tâche simple devrait ressembler à ceci :

```go
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
```

<%= partial("fr/docs/generators/tasks.md") %>

## Lister les tâches disponibles

```bash
$ buffalo task list

Available grifts
================
buffalo task middleware    # Prints out your middleware stack
buffalo task routes        # Print out all defined routes
buffalo task secret        # Generate a cryptographically secure secret key
```

## Exécuter des tâches

### En mode développement

Les tâches peuvent être exécutées en mode développement en utilisant la commande `buffalo task`.

```bash
$ buffalo task hello
```

### Depuis le binaire de votre app

Après la [compilation de votre binaire](/fr/docs/deploy/building), les tâches peuvent être lancées en utilisant la sous-commande `task` :

```bash
$ monapp task hello
```
