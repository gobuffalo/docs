<%= h1("Installation") %>

<%= title("Prérequis") %>

Avant d'installer Buffalo, assurez-vous d'avoir installé les dépendances suivantes :

* [Guide de configuration d'un environnement de développement Go (EN)](http://gopherguides.com/before-you-come-to-class).
* [Go](https://golang.org) version `<%= goMinVersion %>` ou supérieure.

##### Prérequis pour la partie interface utilisateur

* [node](https://github.com/nodejs/node) et [npm](https://github.com/npm/npm) pour la [gestion des ressources](/docs/assets) avec [webpack](https://github.com/webpack/webpack).

##### Prérequis spécifiques aux bases de données

* **SQLite 3**: GCC, ou compilateur équivalent pour [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

<%= title("Installation basique") %>

Buffalo fournit **deux composants majeurs** :
* La commande `buffalo`, une puissante boîte à outils pour vous aider à développer d'une manière rapide et efficace.
* Le *framework* buffalo, un ensemble de pièces pour construire votre application.

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
Ces instructions peuvent également être utilisées pour mettre à jour votre version de Buffalo.
<% } %>
