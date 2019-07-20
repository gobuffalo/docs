<% seoDescription("Comment utiliser Pop avec Buffalo ?") %>
<% seoKeywords(["buffalo", "go", "golang", "pop", "buffalo", "intégration"]) %>

# Intégration avec Buffalo

## Génération d'une nouvelle application

Lorsque vous générez une nouvelle application Buffalo, vous pouvez choisir la base de données à utiliser à l'aide de l'option `--db-type`. Par exemple, pour générer une nouvelle application avec le support de MySQL, vous pouvez utiliser la commande suivante :

```bash
$ buffalo new coke --db-type mysql
```

**Si vous ne choisissez pas le type de base de données, Buffalo utilisera le support pour PostgreSQL par défaut.**

### Ignorer la configuration de base de données

Si vous souhaitez gérer la base de données par vous-même (sans utiliser Pop), ou si vous ne souhaitez pas utiliser de base de données, il est possible d'ignorer la génération du code associé aux bases de données : il suffit d'utiliser l'option `--skip-pop` pour cela.

```bash
$ buffalo new coke --skip-pop
```

## Le middleware Pop pour les transactions

Buffalo fournit un middleware pour Pop, afin de rendre plus facile l'utilisation des bases de données au sein de Buffalo : https://github.com/gobuffalo/buffalo-pop

### Installation

Ce middleware est configuré pour vous par défaut, si vous choisissez d'utiliser Pop à la création d'un nouveau projet.

**actions/app.go**

```go
func App() *buffalo.App {
	if app == nil {
        // [...]

        app.Use(poptx.PopTransaction(models.DB))

        // [...]

        app.GET("/", HomeHandler)
    }

    return app
}
```

`poptx.PopTransaction(models.DB)` utilise la connexion à la base de données configurée, pour créer un nouveau middleware `PopTransaction`. Ce middleware se charge des tâches suivantes :

* Historiser le temps mis par toutes les requêtes en base de données à s'exécuter, pour une requête HTTP.
* Contenir **chaque requête HTTP** dans une transaction de base de données.
* Valider les modifications en base de données (commit) **si aucune erreur n'est survenue** lors de l'exécution des middlewares et de l'action ; **et que le statut de réponse HTTP est de type 2xx ou 3xx**.
* Invalider (rollback) les modifications de base de données le cas échéant.

### Gérer une transaction à la main

Si vous avez besoin de gérer une transaction à la main, vous pouvez neutraliser le middleware pour une route donnée :

```go
func App() *buffalo.App {
	if app == nil {
        // [...]
        txm := poptx.PopTransaction(models.DB)
        app.Use(txm)
        a.Middleware.Skip(txm, HomeHandler)

        // [...]

        app.POST("/form", FormHandler)
        app.GET("/", HomeHandler)
    }

    return app
}
```