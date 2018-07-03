<% seoDescription("Démarrer avec Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "bdd", "ORM", "pop", "migrations"]) %>

<%= h1("Base de données") %>

Le paquet [Pop](https://github.com/gobuffalo/pop) est fourni par défaut avec Buffalo. Il permet de gérer les migrations de schémas, les transactions, fournit un ORM, et bien d'autres fonctionnalités. Comme tout ORM, Pop permet de traduire des structures de modèles en Go vers les tables de la base de données.

Buffalo a une intégration forte avec Pop, et la boîte à outils vous aidera à générer tout ce qui est nécessaire pour démarrer. Vous pouvez toujours utiliser une autre biliothèque pour communiquer avec vos bases de données, mais vous devrez vous débrouiller sans notre aide. :)

<%= title("Démarrer avec Pop") %>

Pop supporte les bases de données suivantes :
* [PostgreSQL](https://www.postgresql.org/)
* [CockroachDB](https://www.cockroachlabs.com/)
* [MySQL](https://www.mysql.com/)
* [SQLite3](https://sqlite.org/)

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