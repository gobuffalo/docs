<% seoDescription("Structure d'un projet Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "structure", "dossiers", "framework", "web"]) %>

<%= h1("Structure d'un projet") %>

Buffalo utilise **une structure minimale de dossiers**  pour travailler sur votre projet. Cette structure permet de **ranger le projet proprement**, et est intimement liée aux [générateurs](/docs/generators). N'essayez pas de réinventer la roue, et **laissez Buffalo vous faire gagner du temps** pour travailler sur la partie la plus intéressante&nbsp;! :)

Maintenant que vous disposez d'un projet minimal, voyons sa composition.

## Le répertoire racine

Voici la structure d'un projet Buffalo&nbsp;:

* `go/` &mdash; racine GOPATH.
	* `src/` &mdash; répertoire des sources Go
		* `github.com/username/myapp/` &mdash; racine de votre application
			* `actions/`
			* `assets/`
			* `grifts/`
			* `locales/`
			* `models/`
			* `public/`
			* `templates/`
			* `tmp/`
			* `database.yml`
			* `main.go`

### actions

Ce dossier contient la partie **Contrôleur** du design-pattern MVC. Il comporte les traitements de vos URLs, ainsi que&nbsp;:

* le fichier `app.go` pour configurer votre application et ses routes,
* le fichier `render.go` pour paramétrer le(s) moteur(s) de templates.

### assets

<%= note() { %>
Ce répertoire est facultatif. Si vous n'avez pas besoin d'interface utilisateur (pour une API, par exemple), vous pouvez le supprimer.
<% } %>

Ce dossier contient les composants graphiques **non-optimisés** qui seront compressés & placés dans le dossier [`public`](#public).

### grifts

<%= note() { %>
Ce répertoire est facultatif. Si vous n'avez pas besoin d'utiliser les [tâches](/docs/tasks), vous pouvez le supprimer.
<% } %>

Ce dossier contient les [tâches](/docs/tasks) basées sur [grift](https://github.com/markbates/grift).

### locales

<%= note() { %>
Ce répertoire est facultatif. Si vous n'utilisez qu'une seule langue, vous pouvez le supprimer&nbsp;; ainsi que le module i18n pré-configuré dans le fichier `app.go` du dossier `actions`.
<% } %>

Ce dossier est utilisé par le système d'<abbr title="internationalization">i18n</abbr>. Les traductions seront récupérées depuis cet endroit.

### models

<%= note() { %>
Si vous utilisez pop/soda avec le générateur intégré, les fichiers de modèles de données seront placés ici.
<% } %>

<%= note() { %>
Ce répertoire est facultatif. Si vous n'avez pas besoin d'utiliser une base de données, vous pouvez le supprimer.
<% } %>

Ce dossier se charge de la partie **modèle** du design-pattern MVC. Il contient le fichier `models.go` pour initialiser la connexion à la base de données, ainsi que les fichiers de modèles de données pour refléter les tables de la base de données.

### public

<%= note() { %>
Le contenu de ce répertoire est auto-généré.
<% } %>

Ce dossier contient les ressources graphiques publiques (compilées/compressées). Si vous utilisez webpack, les ressources produites seront placées ici.

### templates

<%= note() { %>
Ce répertoire est facultatif. Si vous n'avez pas besoin d'interface utilisateur (pour une API, par exemple), vous pouvez le supprimer.
<% } %>

Ce dossier se charge de la partie **vue** du design-pattern MVC. Il contient les gabarits utilisés pour générer les pages.

### tmp

<%= note() { %>
Le contenu de ce répertoire est auto-généré.
<% } %>

Ce dossier est utilisé par la commande `buffalo dev` pour reconstruire votre projet à chaque changement. Les fichiers temporaires de Buffalo sont placés ici.

### database.yml

<%= note() { %>
Ce répertoire est facultatif. Si vous n'avez pas besoin d'utiliser une base de données, vous pouvez le supprimer.
<% } %>

Ce fichier contient la configuration des bases de données utilisées par [pop/soda](https://github.com/gobuffalo/pop).

### main.go

Ce fichier se charge d'amorcer votre application et de la démarrer.

## Prochaines étapes

* [Configuration](/fr/docs/config-vars) - Comment gérer la configuration de mon application ?