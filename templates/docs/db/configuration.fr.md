<% seoDescription("Comment configurer ma base de données avec Pop ?") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "bdd", "ORM", "pop", "configuration"]) %>

<%= h1("Configuration de base de données") %>

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

### Générateur

<%= partial("docs/db/soda_buffalo_note.md") %>

Vous pouvez générer un fichier de configuration par défaut en utilisant la commande d'initialisation:

```bash
$ soda g config
```

Cette commande génère un fichier `database.yml` dans le répertoire courant, pour une base de données PostgreSQL. Vous pouvez choisir le type de base de données en utilisant l'option `-t` et en passant l'un des types supportés : `postgres`, `cockroach`, `mysql` ou `sqlite3`.

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

<%= warning() { %>
Le paramètre `url` de la connexion écrasera tout autre configuration détaillée. Assurez-vous de passer tout paramètre nécessaire dans l'URL de connexion.
<% } %>

Pour plus d'informations, consultez la documentation de [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Assurez-vous d'avoir configuré ce fichier correctement avant de travailler avec Pop !**

<%= title("Options disponibles") %>

### encoding

<%= sinceVersion("4.6.0") %>

Cette option n'est pour le moment supportée que par le **dialecte mysql**. Ce paramètre est utilisé pour créer la base de données (si vous la créez avec `soda`), et en tant que paramètre `collation` de l'URL de connexion à la base de données. Si cette option n'est pas renseignée, la valeur par défaut est  `utf8_general_ci`.

```yaml
development:
  dialect: mysql
  database: myapp_development
  user: root
  password: root
  host: 127.0.0.1
  encoding: "utf8mb4_general_ci"
```