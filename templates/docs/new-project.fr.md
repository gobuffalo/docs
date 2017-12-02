<%= h1("Générer un nouveau projet") %>

Vous disposez maintenant d'une installation de Buffalo fonctionnelle. Dans cette section, vous allez apprendre comment créer **une toute nouvelle application web**, en utilisant la commande `buffalo`.

<%= title("Créer un nouveau projet") %>

Le but de Buffalo est de pouvoir construire de nouvelles applications Web en Go, aussi **vite et simplement** que possible. Peut-on faire plus simple qu'un générateur d'application ?

Commencez par vous placer dans votre `$GOPATH` et créez votre nouvelle application !

```bash
$ cd $GOPATH/src/github.com/$USER/
```

Assurez-vous que `$GOPATH/bin` est dans votre `$PATH`, puis&nbsp;:

```bash
$ buffalo new coke
```

Cette commande va vous générer une nouvelle application Buffalo dénommée **coke**, avec le minimum syndical&nbsp;:
* la **structure du framework Buffalo** et la configuration par défaut ([pop/soda](https://github.com/markbates/pop) avec support pour PostgreSQL),
* toutes les **dépendances Go** nécessaires pour exécuter votre application,
* les **dépendances frontend** et une configuration fonctionnelle pour [webpack](https://webpack.js.org/)
* et un **dépôt Git** initial.

<%= partial("docs/new-project/new.md") %>

<%= title("Créer une application personnalisée") %>

Les options par défaut sont suffisantes, mais ne peuvent pas convenir à tout le monde. Buffalo vous permet donc de changer cette configuration via des *flags* à passer à la commande `new`.

Vous pouvez lister les *flags* disponibles en utilisant la commande `help` : 

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci] (default "none")
      --db-type string       specify the type of database you want to use [postgres, mysql, sqlite3] (default "postgres")
      --docker string        specify the type of Docker file to generate [none, multi, standard] (default "multi")
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
      --with-dep             adds github.com/golang/dep to your app
```

Vous pouvez choisir de générer une application de type API, en excluant les fichiers propres à l'interface utilisateur. Vous voulez utiliser un système d'intégration continue ? Ou même utiliser votre propre dépendance pour gérer les interactions avec la base de données ? Utilisez les *flags* !

<%= partial("docs/dev.md") %>
