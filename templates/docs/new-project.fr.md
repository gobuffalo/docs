<% seoDescription("Générer un nouveau projet Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "nouveau projet", "générateur", "framework", "web"]) %>

<%= h1("Générer un nouveau projet") %>

Vous disposez maintenant d'une installation de Buffalo fonctionnelle. Dans cette section, vous allez apprendre comment créer **une toute nouvelle application web**, en utilisant la commande `buffalo`.

<%= title("Créer un nouveau projet") %>

Le but de Buffalo est de pouvoir construire de nouvelles applications Web en Go, aussi **vite et simplement** que possible. Peut-on faire plus simple qu'un *générateur d'application* ?

Commencez par vous placer dans votre `$GOPATH` et créez votre nouvelle application !

```bash
$ cd $GOPATH/src/github.com/$USER/
```

Assurez-vous que `$GOPATH/bin` est dans votre `$PATH`, puis&nbsp;:

```bash
$ buffalo new coke
```

Cette commande va vous générer une nouvelle application Buffalo dénommée **coke**, avec le minimum syndical&nbsp;:
* la **structure du framework Buffalo** et la configuration par défaut ([pop/soda](https://github.com/gobuffalo/pop) avec support pour PostgreSQL),
* toutes les **dépendances Go** nécessaires pour exécuter votre application,
* les **dépendances frontend** et une configuration fonctionnelle pour [webpack](https://webpack.js.org/)
* et un **dépôt Git** initial.

<%= partial("docs/new-project/new.md") %>

<%= title("Créer une application personnalisée") %>

Les options par défaut sont suffisantes, mais ne peuvent pas convenir à tout le monde. Buffalo vous permet donc de changer cette configuration via des *flags* à passer à la commande `new`.

Vous pouvez lister les *flags* disponibles en utilisant la commande `help` :

<%= partial("docs/new-project/help.md") %>

Vous pouvez choisir de générer une application de type API, en excluant les fichiers propres à l'interface utilisateur. Vous voulez utiliser un système d'intégration continue ? Ou même utiliser votre propre dépendance pour gérer les interactions avec la base de données ? Utilisez les *flags* !

<%= title("Changer la config par défaut") %>

Par défaut, la commande `buffalo new` recherche le fichier de configuration `$HOME/.buffalo.yml`, et tente de le charger s'il existe. Vous pouvez écraser les options de ce fichier en passant directement les flags d'options nécessaires, ou en utilisant l'option `--config` pour prendre un compte un autre fichier YAML à la place. Utilisez l'option `--skip-config` si vous ne souhaitez pas que la commande `buffalo new` charge un fichier de configuration par défaut.

Voici un exemple de fichier de configuration `.buffalo.yml` :

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

<%= partial("docs/dev.md") %>

<%= title("Prochaines étapes") %>

* [Structure d'un projet](/fr/docs/directory-structure) - En savoir plus sur la structure d'un projet Buffalo.
* [Configuration](/fr/docs/config-vars) - Comment gérer la configuration de mon application ?