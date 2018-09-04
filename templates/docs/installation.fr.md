<% seoDescription("Installer le framework Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "installation", "framework", "web"]) %>

<%= h1("Installation") %>

Dans ce chapitre, vous allez apprendre à installer Buffalo, depuis un binaire pré-compilé ou depuis les sources.

Buffalo fournit **deux composants majeurs** :
* La commande `buffalo`, une puissante boîte à outils pour vous aider à développer d'une manière rapide et efficace.
* Le *framework* buffalo, un ensemble de pièces pour construire votre application.

<%= title("Prérequis") %>

Avant d'installer Buffalo, assurez-vous d'avoir installé les dépendances suivantes :

* [Guide de configuration d'un environnement de développement Go (EN)](http://gopherguides.com/before-you-come-to-class).
* [Une variable d'environnement `$PATH` correctement configurée, incluant `$GOPATH/bin`](https://golang.org/doc/code.html#GOPATH).
* [Go](https://golang.org) version `<%= goMinVersion %>` ou supérieure.

##### Prérequis pour le frontend

* [node](https://github.com/nodejs/node) version `8` ou supérieure et [yarn](https://yarnpkg.com/en/) ou [npm](https://github.com/npm/npm) pour la [gestion des ressources](/docs/assets) avec [webpack](https://github.com/webpack/webpack).

##### Prérequis spécifiques aux bases de données

* **SQLite 3**: GCC, ou compilateur équivalent pour [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

<%= title("Installation depuis une archive pré-compilée - 64 bits", {title: "Installation depuis une archive pré-compilée"}) %>

<%= note() { %>
Ces archives sont fournies sans support pour SQLite.
<% } %>

Depuis la version `v0.10.3`, des archives contenant une version pré-compilée de Buffalo sont fournies à chaque publication officielle. Si vous n'avez pas besoin de la toute dernière version de développement et que vous souhaitez une version stable, vous préférerez sans-doute cette version. 

### GNU / Linux

```bash
$ wget  https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_linux_amd64.tar.gz
$ tar -xvzf buffalo_<%= version %>_linux_amd64.tar.gz
$ sudo mv buffalo-no-sqlite /usr/local/bin/buffalo
```

### MacOS

```bash
$ wget  https://github.com/gobuffalo/buffalo/releases/download/v<%= version %>/buffalo_<%= version %>_darwin_amd64.tar.gz
$ tar -xvzf buffalo_<%= version %>_darwin_amd64.tar.gz
$ sudo mv buffalo-no-sqlite /usr/local/bin/buffalo
# ou si vous avez configuré votre PATH pour inclure le dossier ~/bin
$ mv buffalo-no-sqlite ~/bin/buffalo
```

### Windows
Buffalo peut être installé en utilisant le gestionnaire de paquets [Scoop](http://scoop.sh/) :

```powershell
PS C:\> scoop install buffalo
```

<%= title("Installation avec Homebrew") %>

Sur macOS, vous pouvez également installer Buffalo avec [Homebrew](https://brew.sh/). Après avoir [installé](https://docs.brew.sh/Installation) Homebrew, vous pouvez installer Buffalo en une simple ligne de commande :

```bash
brew install gobuffalo/tap/buffalo
```

<%= title("Installation avec GoFish") %>

[GoFish](https://gofi.sh/index.html) est un gestionnaire de paquets multi-plateformes ; qui fonctionne donc avec Windows, MacOSX et Linux.

Après avoir [installé](https://gofi.sh/index.html#install) GoFish, vous pouvez installer Buffalo comme suit :

```bash
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo 0.12.6: installed in 3.223672926s
```

<%= title("Installation personnalisée **avec** support pour SQLite3", {title: "Installation avec support pour SQLite3"}) %>

**SQLite 3** nécessite GCC, ou un compilateur C équivalent pour compiler [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3). Vous **devez** avoir installé GCC **avant** d'installer Buffalo.

Pour installer Buffalo, assurez-vous que le `GOPATH` est défini, puis&nbsp;:

```bash
$ go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo
```

**Utilisateurs de Windows**&nbsp;: Suivez le guide d'installation [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3 (EN)](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) pour installer GCC sur Windows 10. GCC peut également être installé via le gestionnaire de paquets [Scoop](http://scoop.sh/) :

```powershell
PS C:\> scoop install gcc
```

<%= note() { %>
Ces instructions peuvent aussi être utilisées pour mettre à jour votre version de Buffalo.
<% } %>

<%= title("Installation personnalisée **sans** support pour SQLite3", {title: "Installation sans support pour SQLite3"}) %>

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
Ces instructions peuvent également être utilisées pour mettre à jour votre version de Buffalo.
<% } %>

<%= title("Vérifier votre installation") %>

Vous pouvez vérifier que votre installation fonctionne, en exécutant la commande `buffalo` dans un terminal (ou console) :

```bash
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
  update      will attempt to upgrade a Buffalo application to version v0.11.1
  version     Print the version number of buffalo

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```

Si vous avez un retour similaire de la commande, votre boîte à outils Buffalo est prête à fonctionner !

<%= title("Prochaines étapes") %>

* [Intégrations](/fr/docs/integrations) - Configurez votre environnement pour mieux travailler avec Buffalo.
* [Générer un nouveau projet](/fr/docs/new-project) - Créez votre premier projet Buffalo !
