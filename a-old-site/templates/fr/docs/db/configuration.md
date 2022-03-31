<% seoDescription("Comment configurer ma base de données avec Pop ?") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "bdd", "ORM", "pop", "configuration"]) %>

# Configuration de base de données

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

## Générateur

<%= note() { %>
**Note pour les utilisateurs de Buffalo** : les commandes de `soda` sont intégrées à la commande `buffalo`, sous la commande `pop`. À chaque fois que vous souhaitez utiliser une commande de `soda`, utilisez `buffalo pop` à la place.
<% } %>

Vous pouvez générer un fichier de configuration par défaut en utilisant la commande d'initialisation:

```bash
$ soda g config
```

Cette commande génère un fichier `database.yml` dans le répertoire courant, pour une base de données PostgreSQL. Vous pouvez choisir le type de base de données en utilisant l'option `-t` et en passant l'un des types supportés : `postgres`, `cockroach`, `mysql` ou `sqlite3`.

## Emplacement du fichier de configuration

Le fichier de configuration de Pop, `database.yml` peut se situer :
* À la racine de votre projet (défaut).
* Dans un dossier `config/`, à la racine de votre projet.

Si vous souhaitez placer votre fichier de configuration à un autre endroit, vous pouvez utiliser [`AddLookupPaths`](https://godoc.org/github.com/gobuffalo/pop#AddLookupPaths).

Il est également possible de personnaliser le nom de ce fichier :

```go
pop.ConfigName = "ma_config_pop.yml"
```

## Configuration par environnement vs détaillée

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

<%= warning() { %>
Le paramètre `url` de la connexion écrasera tout autre configuration détaillée. Assurez-vous de passer tout paramètre nécessaire dans l'URL de connexion.
<% } %>

Pour plus d'informations, consultez la documentation de [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Assurez-vous d'avoir configuré ce fichier correctement avant de travailler avec Pop !**

## Options disponibles

### database

Le nom de la base de données à utiliser.

### dialect

Le dialecte de base de données à utiliser avec la connection. Les valeurs acceptées sont :
* Pilote MySQL : « mysql »
* Pilote PostgreSQL : « postgres », « postgresql » ou « pg »
* Pilote Cockroach DB : « cockroach », « cockroachdb » ou « crdb »
* Pilote SQLite : « sqlite » ou « sqlite3 »

### driver

<%= sinceVersion("4.11.2") %>

Vous pouvez utiliser cette option pour personnaliser le pilote de base de données utilisé.

Voici les pilotes utilisés par défaut :
* MySQL : [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* PostgreSQL : [github.com/lib/pq](https://github.com/lib/pq)
* Cockroach DB : [github.com/cockroachdb/cockroach-go/crdb](https://github.com/cockroachdb/cockroach-go/tree/master/crdb)
* SQLite : [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

### encoding

<%= sinceVersion("4.6.0") %>

Cette option n'est pour le moment supportée que par le **dialecte mysql**. Ce paramètre est utilisé pour créer la base de données (si vous la créez avec `soda`), et en tant que paramètre `collation` de l'URL de connexion à la base de données. Si cette option n'est pas renseignée, la valeur par défaut est  `utf8mb4_general_ci`.

```yaml
development:
  dialect: mysql
  options:
    encoding: "utf8_general_ci"
```

### host

L'adresse de l'hôte de la base de données.

### password

Le mot de passe de l'utilisateur de la base de données.

### port

Le port de la base de données sur l'hôte.

**Valeurs par défaut** :

| Pilote    | Port  |
|-----------|-------|
| PostgreSQL| 5432  |
| MySQL     | 3306  |
| Cockroach | 26257 |

### user

L'utilisateur de la base de données.