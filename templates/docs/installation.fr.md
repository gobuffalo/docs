<%= h1("Installation") %>

Buffalo fournit **deux composants majeurs** :
* La commande `buffalo`, une puissante boîte à outils pour vous aider à développer d'une manière rapide et efficace.
* Le *framework* buffalo, un ensemble de pièces pour construire votre application.

<%= title("Prérequis") %>

Avant d'installer Buffalo, assurez-vous d'avoir installé les dépendances suivantes :

* [Guide de configuration d'un environnement de développement Go (EN)](http://gopherguides.com/before-you-come-to-class).
* [Go](https://golang.org) version `<%= goMinVersion %>` ou supérieure.

##### Prérequis pour le frontend

* [node](https://github.com/nodejs/node) et [npm](https://github.com/npm/npm) pour la [gestion des ressources](/docs/assets) avec [webpack](https://github.com/webpack/webpack).

##### Prérequis spécifiques aux bases de données

* **SQLite 3**: GCC, ou compilateur équivalent pour [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

<%= title("Installation **avec** support pour SQLite3", {name: "avec-SQLite3", title: "Installation avec support pour SQLite3"}) %>

**SQLite 3** nécessite GCC, ou un compilateur C équivalent pour compiler [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3). Vous **devez** avoir installé GCC **avant** d'installer Buffalo.

Pour installer Buffalo, assurez-vous que le `GOPATH` est défini, puis&nbsp;:

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

**Utilisateurs de Windows**&nbsp;: Suivez le guide d'installation [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3 (EN)](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) pour installer GCC sur Windows 10.

<%= note() { %>
Ces instructions peuvent également être utilisées pour mettre à jour votre version de Buffalo.
<% } %>

<%= title("Installation **sans** support pour SQLite3", {name: "sans-SQLite3", title: "Installation sans support pour SQLite3"}) %>

```bash
$ go get -u -v -tags nosqlite github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
Ces instructions peuvent également être utilisées pour mettre à jour votre version de Buffalo.
<% } %>
