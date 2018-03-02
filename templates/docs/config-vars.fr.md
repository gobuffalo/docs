<%= h1("Configuration") %>

Dans ce chapitre, vous allez apprendre comment gérer la configuration de votre application avec Buffalo.

Les variables d'environnement sont une bonne méthode pour séparer les valeurs sensibles (ou non) liées à un environnement spécifique, du reste de votre application ([comme le décrit The Twelve Factor app (EN)](https://12factor.net/config)). Cela peut aider à définir un comportement basé sur le contexte de l'application (forcer le SSL en production, par exemple), mais aussi isoler les clefs secrètes d'API (comme une API d'accès à une banque). De cette manière, les développeurs peuvent utiliser des valeurs de type « bac-à-sable », et ne pas mettre en danger la version de production.

<%= title("Variables d'environnement disponibles", {name: "variables-disponibles"}) %>

Les variables suivantes sont utilisées par Buffalo :

| Variable              | Valeur par défaut        | Rôle                                                       |
| ---                   | ---                      | ---                                                        |
| `GO_ENV`              | `development`            | L'« environment » d'exécution de Buffalo                   |
| `GO_BIN`              | `go`                     | Le compilateur Go à utiliser                               |
| `BUFFALO_PLUGIN_PATH` | `$PATH`                  | L'endroit où Buffalo va chercher ses plugins               |
| `ADDR`                | `127.0.0.1` or `0.0.0.0` | L'adresse d'écoute du serveur HTTP                         |
| `PORT`                | `3000`                   | Le port d'écoute du serveur HTTP                           |
| `HOST`                | `http://127.0.0.1:$PORT` | L'« URL » de l'application (c-à-d son adresse publique)    |
| `SESSION_SECRET`      | `""`                     | Un sel utilisé pour sécuriser les sessions                 |

<%= title("Configuration personnalisée") %>

Vous pouvez toujours fournir vos propres variables, et en récupérer les valeurs depuis votre application. Le paquet [envy](https://github.com/gobuffalo/envy) rend cela très facile&nbsp;!

```go
import "github.com/gobuffalo/envy"

// [...]

// Get MYSECRET env variable, default to empty string if it's not set
var MYSECRET = envy.Get("MYSECRET", "")

// Get REQUIREDSECRET env variable, return an error if it's not set
REQUIREDSECRET, err := envy.MustGet("REQUIREDSECRET")
```

<%= title("Support des fichiers .env") %>

<%= sinceVersion("0.10.3") %>

Buffalo est fourni avec le support des fichiers `.env` (**depuis buffalo >= 0.10.3**), ce qui signifie que Buffalo chargera le contenu de ces fichiers comme variables d'environnement, dès le démarrage de l'application. Pour ce faire, Buffalo utilise [`envy.Load`](https://github.com/gobuffalo/envy/blob/e613c80275b86293880eddeb27417c9a7c670ff3/envy.go#L53) qui cherchera un fichier `.env` à la racine de votre application.

Si vous n'êtes pas familier avec la structure d'un fichier `.env`, en voici un exemple :

```text
SENDGRID_API_KEY=ABCCOQ7GFRVCW0ODHPFQ3FTP5SLL1Q
SENDGRID_EMAIL=email@myapp.com

APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://myapp.com
```

Les applications générées (**avec buffalo >= 0.10.3**) contiennent déjà un fichier `.env` par défaut à la racine du projet. Ce fichier est surveillé par Buffalo pour relever les changements éventuels, mais il sera ignoré par git (il est ajouté au fichier `.gitignore`). C'est donc une bonne méthode pour éviter que les dévelopeurs poussent des données sensibles ou propres à leur environnement par erreur.
