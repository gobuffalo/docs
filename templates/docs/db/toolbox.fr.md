<% seoDescription("Boîte à outils Soda") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "outils", "pop", "soda"]) %>

<%= h1("Boîte à outils de BDD") %>

Pop est une bibliothèque permettant de communiquer avec des bases de données, mais elle founit également `soda`, une petite boîte à outils en ligne de commande pour gérer vos bases de données. Elle peut vous aider à créer une nouvelle base de données, supprimer des bases existantes, et bien plus.

<%= note() { %>
**Note pour les utilisateurs de Buffalo**: les commandes de `soda` sont intégrées à la commande `buffalo`, sous la commande `db`. À chaque fois que vous voulez utiliser une commande de `soda`, il vous suffit d'utiliser `buffalo db` à la place.
<% } %>

<%= title("Créer des bases de données") %>

Une fois que le fichier `database.yml` a été correctement configuré, et que le serveur de base de données fonctionne, Soda peut créer toutes les bases déclarées dans le fichier `database.yml`, à l'aide d'une simple commande :

```bash
$ soda create -a
```

Vous pouvez également choisir d'en créer qu'une seule, en utilisant l'option `-e` pour choisir le nom de la base de données à créer :

```bash
$ soda create -e test
```

<%= title("Supprimer une base de données") %>

Soda peut supprimer toutes les bases de données configurées dans le fichier `database.yml`, à l'aide d'une simple commande :

```bash
$ soda drop -a
```

De la même manière que pour créer une base de données, il est possible de choisir seulement l'une des configurations du fichier, en utilisant l'option `-e` :

```bash
$ soda drop -e test
```