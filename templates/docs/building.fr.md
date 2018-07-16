<% seoDescription("Empaquetage d'une app") %>
<% seoKeywords(["buffalo", "go", "golang", "empaquetage", "binaire", "build", "contruction", "compilation"]) %>

<%= h1("Empaquetage d'une app") %>

Votre projet est désormais prêt à être déployé. Dans cette section, vous allez apprendre comment empaqueter une version de votre application et la déployer sur un serveur.

<%= title("La commande build") %>

Buffalo contient une commande `build`. Celle-ci permet de construire un **exécutable complet** de votre application, comprenant (mais pas seulement) les fichiers statiques (c'est-à-dire les ressources graphiques), migrations, templates, etc. Si vous adhérez à la «&nbsp;façon de faire Buffalo&nbsp;», les choses fonctionnent simplement. C'est une expérience incroyable. :)

<%= partial("docs/building/build_trace.md") %>

Lorsque le build est terminé, vous obtenez un binaire tout frais dans le dossier `bin`. Il contient également **l'heure de compilation** et le **hash SHA du commit git** intégrés&nbsp;: cela permet de «&nbsp;versionner&nbsp;» vos binaires.

<%= title("Personnaliser le binaire") %>

Pour lister les options disponibles, utilisez la commande help&nbsp;:

<%= partial("docs/building/build_options.md") %>

### Nom du binaire / emplacement

Par défaut, votre application compilée sera placée dans le dossier `bin` de votre projet, et le nom de l'exécutable sera basé sur celui que vous avez utilisé pour créer le projet à l'aide de la commande `new`.

Vous pouvez changer ce nom en utilisant le flag `-o` ou `-output`&nbsp;:

<%= partial("docs/building/output_flag.md") %>

En réalité, vous pouvez aussi changer le dossier de destination&nsbp;:

<%= partial("docs/building/output_dir.md") %>

### Extraire les fichiers statiques dans un fichier zip

Par défault, toute l'application est contenue dans un seul exécutable, les fichiers statiques inclus. Dans un système en production&nsbp;; vous voudrez peut-être servir ces fichiers statiques avec un serveur proxy (comme Apache ou NGINX), pour réduire la charge de votre application. Vous voudrez peut-être même utiliser un *CDN* pour gérer vos fichiers statiques.

Buffalo fournit un moyen d'extraire les fichiers statiques compilés dans une unique archive, en utilisant le flag `-e` ou `-extract-assets`&nbsp;:

<%= partial("docs/building/extract_assets.md") %>

Par défault, l'archive des fichiers statiques est place dans le répertoire *bin*, mais si vous changez le dossier de destination de l'exécutable avec le flag `-o`, les fichiers statiques seront livrés dans le même répertoire.

<%= partial("docs/building/extract_assets_layout.md") %>

<%= title("Options avancées") %>

### Construire des binaires « statiques » avec CGO

La création de binaire lié statiquement (ceux qui utilisent CGO, comme avec SQLite3) peut être délicat. En utilisant le flag `--static` avec `buffalo build`, les flags `--ldflags '-linkmode external -extldflags "-static"'` seront ajoutés à la commande `go build`.

### Build Tags

Quand vous construisez un binaire Buffalo en utilisant la commande `buffalo build`, vous pouvez passer les flags `--tags` et `--ldflags` pour construire le binaire&nbsp;; comme vous le feriez normalement avec l'outil `go build`.

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

<%= title("Commandes binaires") %>

### Modes
Le binaire fonctionne par défaut en mode `développement`, ce qui signifie que toutes les sous-commandes seront executées aussi dans ce mode. Pour changer le mode, utilisez la variable d'environment `GO_ENV`&nbsp;:

```bash
$ GO_ENV=production ./coke
```

### Commandes disponibles

Une fois que le binaire a été construit, il y a plusieurs sous-commandes que vous pouvez exécuter sur celui-ci&nbsp;:

#### Défault

La commande par défaut, si vous exécutez simplement le binaire, va démarrer l'application.

#### migrate

La sous-commande `migrate` va exécuter les commandes de migration (de base de données) pour l'application.

#### version

La sous-commande `version` va afficher la version du binaire, en y incluant le nom, le SHA du commit git utilisé pour construire le binaire et la date et heure de construction de celui-ci.

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task

La sous-commande `task` exécute les tâches.

```bash
$ ./coke task greet

Hello World!
```
