<% seoDescription("Installer le framework Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "installation", "framework", "web", "mac", "windows", "linux"]) %>

<%= h1("Installer Buffalo") %>

Dans ce chapitre, vous allez apprendre à installer Buffalo, depuis un binaire pré-compilé ou depuis les sources.

Buffalo fournit **deux composants majeurs** :
* La commande `buffalo`, une puissante boîte à outils pour vous aider à développer d'une manière rapide et efficace.
* Le *framework* buffalo, un ensemble de pièces pour construire votre application.

Buffalo est actuellement disponible et testé sur les plateformes suivantes :
* GNU/Linux
* Mac OSX
* Windows

## Prérequis

Avant d'installer Buffalo, assurez-vous d'avoir installé les dépendances suivantes :

* [En environnement fonctionnel pour Go (EN)](http://gopherguides.com/before-you-come-to-class).
* [Une variable d'environnement `$PATH` correctement configurée, incluant `$GOPATH/bin`](https://golang.org/doc/code.html#GOPATH).
* [Go](https://golang.org) version `<%= goMinVersion %>`.

##### Prérequis pour le frontend

Les pré-requis suivants sont optionnels. Vous n'en aurez pas besoin si vous souhaitez créer une API, ou si vous préférez construire votre application d'une manière plus traditionnelle.

* [node](https://github.com/nodejs/node) version `8` ou supérieure
* [yarn](https://yarnpkg.com/en/) ou [npm](https://github.com/npm/npm) pour la [gestion des ressources](/fr/docs/assets) avec [webpack](https://github.com/webpack/webpack).

##### Prérequis spécifiques aux bases de données

Là encore, si vous n'avez pas besoin de base de données, vous pouvez vous affranchir de ce paragraphe.

* **SQLite 3**: GCC, ou compilateur équivalent pour [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

## Installation depuis une archive pré-compilée - 64 bits

<%= note() { %>
Ces archives sont fournies sans support pour SQLite.
<% } %>

Depuis la version `v0.10.3`, des archives contenant une version pré-compilée de Buffalo sont fournies à chaque publication officielle. Si vous n'avez pas besoin de la toute dernière version de développement et que vous souhaitez une version stable, vous préférerez sans-doute cette version.

### GNU / Linux

```console
$ wget https://github.com/gobuffalo/cli/releases/download/v<%= version %>/buffalo_<%= version %>_Linux_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Linux_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### MacOS

```console
$ curl -OL https://github.com/gobuffalo/cli/releases/download/v<%= version %>/buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ tar -xvzf buffalo_<%= version %>_Darwin_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
# ou si vous avez configuré votre PATH pour inclure le dossier ~/bin
$ mv buffalo ~/bin/buffalo
```

## Scoop (Windows)
Buffalo peut être installé en utilisant le gestionnaire de paquets [Scoop](http://scoop.sh/) :

```powershell
PS C:\> scoop install buffalo
```

## Chocolatey (Windows)
Buffalo peut être installé en utilisant le gestionnaire de paquets [Chocolatey](https://chocolatey.org/packages/buffalo). Les versions sur Chocolatey peuvent avoir du retard sur les autres, vu que chaque nouvelle version doit passer par une étape de modération :

```powershell
PS C:\> choco install buffalo
```

## Homebrew (macOS)

Sur macOS, vous pouvez également installer Buffalo avec [Homebrew](https://brew.sh/). Après avoir [installé](https://docs.brew.sh/Installation) Homebrew, vous pouvez installer Buffalo en une simple ligne de commande :

```console
brew install gobuffalo/tap/buffalo
```

## GoFish (multi-plateformes)

[GoFish](https://gofi.sh/index.html) est un gestionnaire de paquets multi-plateformes ; qui fonctionne donc avec Windows, MacOSX et Linux.

Après avoir [installé](https://gofi.sh/index.html#install) GoFish, vous pouvez installer Buffalo comme suit :

```console
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo <%= version %>: installed in 3.223672926s
```

## Installation personnalisée **avec** support pour SQLite3

**SQLite 3** nécessite GCC, ou un compilateur C équivalent pour compiler [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3). Vous **devez** avoir installé GCC **avant** d'installer Buffalo.

Pour installer Buffalo, assurez-vous que le `GOPATH` est défini, puis&nbsp;:

```console
$ go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@v<%= version %>
```

**Utilisateurs de Windows**&nbsp;: Suivez le guide d'installation [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3 (EN)](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) pour installer GCC sur Windows 10. GCC peut également être installé via le gestionnaire de paquets [Scoop](http://scoop.sh/) :

```powershell
PS C:\> scoop install gcc
```

<%= note() { %>
Ces instructions peuvent aussi être utilisées pour mettre à jour votre version de Buffalo.
<% } %>

## Installation personnalisée **sans** support pour SQLite3

```console
$ go install github.com/gobuffalo/cli/cmd/buffalo@v<%= version %>
```

<%= note() { %>
Ces instructions peuvent également être utilisées pour mettre à jour votre version de Buffalo.
<% } %>

## Vérifier votre installation

Vous pouvez vérifier que votre installation fonctionne, en exécutant la commande `buffalo` dans un terminal (ou console) :

```console
$ buffalo
Helps you build your Buffalo applications that much easier!

Usage:
  buffalo [command]

Available Commands:
  build       Builds a Buffalo binary, including bundling of assets (packr & webpack)
  db          A tasty treat for all your database needs
  destroy     Allows to destroy generated code.
  dev         Runs your Buffalo app in 'development' mode
  generate    A collection of generators to make life easier
  help        Help about any command
  info        Prints off diagnostic information useful for debugging.
  new         Creates a new Buffalo application
  setup       Setups a newly created, or recently checked out application.
  task        Runs your grift tasks
  test        Runs the tests for your Buffalo app
  update      will attempt to upgrade a Buffalo application to version v<%= version %>
  version     Print the version number of buffalo

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

Si vous avez un retour similaire de la commande, votre boîte à outils Buffalo est prête à fonctionner !

## Prochaines étapes

* [Intégrations](/fr/docs/getting-started/integrations) - Configurez votre environnement pour mieux travailler avec Buffalo.
* [Générer un nouveau projet](/fr/docs/getting-started/new-project) - Créez votre premier projet Buffalo !
