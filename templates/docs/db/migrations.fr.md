<% seoDescription("Migrations") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "ORM", "pop", "migration"]) %>

<%= h1("Migrations") %>

Maintenir un logiciel n'est pas une tâche simple, et vous devrez sans doûte patcher votre base de données pour ajouter, modifier ou supprimer des champs. Pop résout ce problème à l'aide des **migrations**.

Vous pouvez écrire de nouvelles migrations en utilisant `fizz`, un langage de description des changements de votre base de données agnostique ; ou utiliser du SQL si vous préférez.

<%= title("Écrire des migrations") %>
<%= partial("docs/db/fizz.md") %>
<%= partial("docs/db/sql.md") %>

<%= title("Exécuter des migrations") %>

<%= partial("docs/db/soda_buffalo_note.md") %>

### Appliquer les migrations
Après avoir créé vos migrations, vous pouvez les exécuter à l'aide de l'une des commandes suivantes :

```bash
$ soda migrate
$ soda migrate up
```

Ces deux commandes sont identiques d'un point de vue fonctionnel, mais la première est plus courte à taper !

### Annuler une migration
Si vous souhaitez annuler l'effet de la dernière migration, vous pouvez utiliser la commande suivante :

```bash
$ soda migrate down
```

---

Pour plus d'informations sur la commande `migrate`, vous pouvez accéder à l'aide en tapant :

```bash
$ soda migrate --help

Runs migrations against your database.

Usage:
  soda migrate [flags]
  soda migrate [command]

Aliases:
  migrate, m

Available Commands:
  down        Apply one or more of the 'down' migrations.
  reset       The equivalent of running `migrate down` and then `migrate up`
  status      Displays the status of all migrations.
  up          Apply all of the 'up' migrations.

Flags:
  -h, --help   help for migrate

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")

Use "soda migrate [command] --help" for more information about a command.
```

<%= title("Cibler un type de base de données") %>

Depuis Pop [v4.4.0](https://github.com/gobuffalo/pop/releases/tag/v4.4.0), les migrations peuvent cibler un type de base de données en utilisant un suffixe. Cela permet d'utiliser des commandes SQL propres à un dialecte de base de données, tout en ignorant les autres.

Par exemple, si vous voulez supporter à la fois PostgreSQL et MySQL, vous pouvez créer deux migrations :

* `my-migration.mysql.up.sql` et `my-migration.mysql.down.sql` seront utilisés lors de la migration d'une base de données MySQL.
* `my-migration.postgres.up.sql` et `my-migration.postgres.down.sql` seront appliqués lors de la migration d'une base de données PostgreSQL.

Si aucune version spécifique au dialecte ne peut être trouvée, Pop tentera d'utiliser la version classique non-suffixée à la place, si elle existe.

<%= title("Table de migrations personnalisée") %>

Par défaut, les migrations appliquées sont suivies dans la table `schema_migration`. Cette dernière est créée par Pop si elle n'existe pas.

Dans certains cas, cependant, vous souhaiterez peut-être utiliser un nom différent pour cette table. C'est possible depuis Pop v4.5.0, en utilisant l'option `migration_table_name`. Dans l'exemple ci-dessous, les migrations seront suivies dans la table `migrations` :

```yaml
development:
  dialect: "postgres"
  url: "your_db_development"
  options:
    migration_table_name: migrations
```

<%= partial("docs/db/deployed_app.md") %>
