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

<%= title("Configuration") %>

La configuration de Pop est gérée par le fichier `database.yml`, qui se trouve à la racine de votre projet. Ce fichier est généré par Buffalo si vous avez choisi de travailler avec Pop : il contient une configuration basique pour la base de données que vous avez sélectionnée lors de la génération de votre application, avec l'option `--db-type`. PostgreSQL est considéré comme le choix par défaut, comme indiqué précédement.

Voici un exemple de configuration générée pour une nouvelle application basée sur PostgreSQL :

```yaml
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

Comme vous pouvez le voir, trois connexions sont définies :
* `development` est la connexion utilisée par votre application en mode `development`.
* `test` sert à exécuter les tests d'intégration.
* `production` est la configuration que vous utiliserez avec l'application finale, sur le serveur.

Bien entendu, vous pouvez configurer autant de connexions que vous le désirez, mais Buffalo ne les utilisera pas par défaut.

### Configuration par environnement vs détaillée

<%= note() { %>
Notez que le fichier `database.yml` est aussi un template Go, vous pouvez donc utiliser la syntaxe de ces templates. Deux fonctions spéciales sont disponibles, `env` et `envOr`.
<% } %>

Comme vous pouvez le voir, il est possible de configurer une nouvelle connexion de deux manières :
* Celle utilisée par la connexion `development` est la plus détaillée. Elle permet de configurer chaque paramètre un par un.
* Celle utilisée par les connexions `test` et `production` est un peu différente : elle utilise des variables (comme l'indique les symboles `{{ }}`) pour définir la valeur, et la fonction `envOr`.

Cette fonction `envOr` essaie d'obtenir une valeur nommée depuis les variables d'environnement, et utilise la deuxième valeur si elle est introuvable. Par exemple :

```yaml
envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"
```

`envOr` essaie de récupérer la valeur du paramètre `TEST_DATABASE_URL` depuis l'environnement, et utilise  `postgres://postgres:postgres@127.0.0.1:5432/myapp_test` si le paramètre n'est pas défini.

De cette manière, vous pouvez utiliser une valeur par défaut pour la phase de développement, et permettre de modifier cette configuration via une variable d'environnement !

<%= note() { %>
Le paramètre `url` de la connexion écrasera tout autre configuration détaillée. Assurez-vous de passer tout paramètre nécessaire dans l'URL de connexion.
<% } %>

Pour plus d'informations, consultez la documentation de [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Assurez-vous d'avoir configuré ce fichier correctement avant de travailler avec Pop !**