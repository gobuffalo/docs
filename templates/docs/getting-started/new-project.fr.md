<% seoDescription("Générer un nouveau projet Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "nouveau projet", "générateur", "framework", "web"]) %>

<%= h1("Générer un nouveau projet") %>

Vous disposez maintenant d'une installation de Buffalo fonctionnelle. Dans cette section, vous allez apprendre comment créer **une toute nouvelle application web**, en utilisant la commande `buffalo`.

## Créer un nouveau projet

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

## Créer une application personnalisée

Les options par défaut sont suffisantes, mais ne peuvent pas convenir à tout le monde. Buffalo vous permet donc de changer cette configuration via des *flags* à passer à la commande `new`.

Vous pouvez lister les *flags* disponibles en utilisant la commande `help` :

<%= partial("docs/new-project/help.md") %>

Vous pouvez choisir de générer une application de type API, en excluant les fichiers propres à l'interface utilisateur. Vous voulez utiliser un système d'intégration continue ? Ou même utiliser votre propre dépendance pour gérer les interactions avec la base de données ? Utilisez les *flags* !

## Changer la config par défaut

Par défaut, la commande `buffalo new` recherche le fichier de configuration `$HOME/.buffalo.yml`, et tente de le charger s'il existe. Vous pouvez écraser les options de ce fichier en passant directement les flags d'options nécessaires, ou en utilisant l'option `--config` pour prendre un compte un autre fichier YAML à la place. Utilisez l'option `--skip-config` si vous ne souhaitez pas que la commande `buffalo new` charge un fichier de configuration par défaut.

Voici un exemple de fichier de configuration `.buffalo.yml` :

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

## Lancer votre application en mode développement

<%= note() { %>
Avant de démarrer Buffalo pour la première fois, veuillez consulter la documentation sur les [bases de données](/docs/db) pour paramétrer votre application correctement.
<% } %>

L'un des inconvénients du développement en Go est l'absence de «&nbsp;rechargement&nbsp;» de code. Cela signifie qu'à chaque fois que vous modifiez votre code, **vous devez stopper manuellement** votre application, la recompiler et enfin la redémarrer. Buffalo trouve cela ennuyeux, et veut vous faciliter la vie.

```bash
$ buffalo dev
```

La commande `dev` surveille les fichiers `.go`, `.html` et le dossier des [ressources graphiques](/docs/assets) par défaut. Elle se charge de **recompiler et redémarrer votre application** automatiquement, pour que vous n'ayez pas à vous en soucier.

Lancez la commande `buffalo dev` et affichez la page [localhost:3000/](http://localhost:3000/) pour voir tous vos changements en direct&nbsp;!

<figure>
  <img src="/assets/images/new-coke.png" title="Capture d'écran">
  <figcaption>La toute nouvelle application Coke.</figcaption>
</figure>

#### Lancer le serveur de développement sur un port personnalisé

Il arrive souvent que vous ayez déjà une application utilisant le port 3000. Vous pouvez configurer le port utilisé par le serveur de développement en utilisant la variable d'environnement `PORT`&nbsp;:

```bash
$ PORT=3001 buffalo dev
```

Vous pouvez également consulter le chapitre sur les [variables d'environnement](/fr/docs/getting-started/config-vars) pour plus d'informations sur la configuration de Buffalo.

## Prochaines étapes

* [Structure d'un projet](/fr/docs/getting-started/directory-structure) - En savoir plus sur la structure d'un projet Buffalo.
* [Configuration](/fr/docs/getting-started/config-vars) - Comment gérer la configuration de mon application ?