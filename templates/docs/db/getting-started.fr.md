<% seoDescription("Démarrer avec Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "bdd", "ORM", "pop", "migrations"]) %>

# Démarrer avec Pop

Le paquet [pop](https://godoc.org/github.com/gobuffalo/pop) est fourni par défaut avec Buffalo, mais vous pouvez l'utiliser en dehors de l'écosystème Buffalo. Pop se sert du fantastique paquet https://github.com/jmoiron/sqlx, normalise quelques patterns courants et workflows généralement associés à la manipulation de bases de données en Go.

**Pop rend les opérations de type CRUD simples à l'aide de fonctionnalités basiques d'ORM, permet de déployer des migrations et de construire et d'exécuter des requêtes.**

Pop, par défaut, suit des conventions influencées par la gem Ruby ActiveRecord. Qu'est-ce que cela signifie ?

* Les tables doivent avoir une colonne « id » et un champ « ID » correspondant dans la structure Go utilisée.
* Si une colonne `created_at` de type timestamp est définie, et que la structure Go correspondante possède un attribut `CreatedAt time.Time` ; la valeur de ce champ sera automatiquement modifiée avec le temps courant lors de la création de l'entrée dans la base.
* Si une colonne `updated_at` de type timestamp est définie, et que la structure Go correspondante possède un attribut `UpdatedAt time.Time` ; la valeur de ce champ sera automatiquement modifiée avec le temps courant à chaque fois que l'entrée sera modifiée dans la base.
* Les noms des tables dans la base de données sont par défaut en minuscules, au pluriel et en notation [snake_case](https://fr.wikipedia.org/wiki/Snake_case). Par exemple : `User{}` donne « users », `FooBar{}` donne « foo_bars », etc.

Buffalo a une intégration forte avec Pop, et la boîte à outils vous aidera à générer tout ce qui est nécessaire pour démarrer. Vous pouvez toujours utiliser une autre biliothèque pour communiquer avec vos bases de données, mais vous devrez vous débrouiller sans notre aide. :)

### Bases de données supportées

Pop supporte les bases de données suivantes :
* [PostgreSQL](https://www.postgresql.org/) (>= 9.3)
* [CockroachDB](https://www.cockroachlabs.com/) (>= 1.1.1)
* [MySQL](https://www.mysql.com/) (>= 5.7)
* [SQLite3](https://sqlite.org/) (>= 3.x)

### Intégration avec Buffalo

Lorsque vous générez une nouvelle application Buffalo, vous pouvez choisir la base de données à utiliser à l'aide de l'option `--db-type`. Par exemple, pour générer une nouvelle application avec le support de MySQL, vous pouvez utiliser la commande suivante :

```bash
$ buffalo new coke --db-type mysql
```

**Si vous ne choisissez pas le type de base de données, Buffalo utilisera le support pour PostgreSQL par défaut.**

#### Ignorer la configuration de base de données

Si vous souhaitez gérer la base de données par vous-même (sans utiliser Pop), ou si vous ne souhaitez pas utiliser de base de données, il est possible d'ignorer la génération du code associé aux bases de données : il suffit d'utiliser l'option `--skip-pop` pour cela.

```bash
$ buffalo new coke --skip-pop
```

## Installation

```bash
$ go get github.com/gobuffalo/pop/...
```

## Prochaines étapes

* [Configuration](/fr/docs/db/configuration) - Configurer vos connexions de base de données.