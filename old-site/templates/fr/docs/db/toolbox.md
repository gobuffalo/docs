<% seoDescription("CLI Soda") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "outils", "pop", "CLI", "soda"]) %>

<%= h1("CLI Soda") %>

Pop est une bibliothèque permettant de communiquer avec des bases de données, mais elle fournit également `soda`, une petite boîte à outils en ligne de commande pour gérer vos bases de données. Elle peut vous aider à créer une nouvelle base de données, supprimer des bases existantes, et bien plus.

<%= note() { %>
**Note pour les utilisateurs de Buffalo**: les commandes de `soda` sont intégrées à la commande `buffalo`, sous la commande `pop`. À chaque fois que vous voulez utiliser une commande de `soda`, il vous suffit d'utiliser `buffalo pop` à la place. Vous n'avez pas besoin d'installer la CLI `soda`.
<% } %>

## Installer la CLI

### Depuis une archive pré-compilée

Les archives pré-compilées contiennent Soda **avec support pour SQLite**.

Téléchargez la version appropriée pour votre plate-forme depuis les [versions de Pop](https://github.com/gobuffalo/pop/releases).

Placez-la quelque part dans votre `PATH`, et assurez-vous que le binaire `soda` est exécutable.

### Homebrew (macOS)

```bash
$ brew install gobuffalo/tap/pop
```

### Depuis les sources

Pour go version 1.16 et versions ultérieures,

**Sans** support pour sqlite 3 :

```bash
$ go install github.com/gobuffalo/pop/v6/soda@latest
```

**Avec** support pour sqlite 3 (nécessite un compilateur C, GCC ou équivalent) :

```bash
$ go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

Si vous ne compilez pas votre code avec `buffalo build` (donc vous n'utilisez probablement pas Buffalo), vous devrez aussi passer l'option `-tags sqlite` à `go build` lors de la compilation.

## Créer des bases de données

Une fois que le fichier `database.yml` a été correctement configuré, et que le serveur de base de données fonctionne, Soda peut créer toutes les bases déclarées dans le fichier `database.yml`, à l'aide d'une simple commande :

```bash
$ soda create -a
```

Vous pouvez également choisir d'en créer qu'une seule, en utilisant l'option `-e` pour choisir le nom de la base de données à créer :

```bash
$ soda create -e test
```

## Supprimer une base de données

Soda peut supprimer toutes les bases de données configurées dans le fichier `database.yml`, à l'aide d'une simple commande :

```bash
$ soda drop -a
```

De la même manière que pour créer une base de données, il est possible de choisir seulement l'une des configurations du fichier, en utilisant l'option `-e` :

```bash
$ soda drop -e test
```
