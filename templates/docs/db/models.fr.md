<% seoDescription("Modèles") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "modèles"]) %>

<%= h1("Modèles") %>

Pop, en tant qu'ORM, vous permet de traduire les tables de votre base de données en des structures en Go. De cette manière, vous pouvez manipuler écrire en Go ce qui nécessiterait d'écrire du SQL. Le code en Go qui permet de faire ce travail est nommé « modèles », en référence à l'architecture MVC. 

Dans ce chapitre, vous allez apprendre comment travailler avec les modèles à la main ; puis comment le faire avec les générateurs fournis, pour améliorer votre productivité.

<%= title("Le dossier models") %>

Les fichiers de modèle de Pop sont placés dans le dossier `models`, à la racine de votre projet (voir le chapitre sur [la structure d'un projet](/fr/docs/directory-structure) pour plus d'informations sur la manière dont Buffalo organise ses fichiers).

Ce répertoire contient :

* Un fichier `models.go`, qui définit le code commun à tous les modèles. Il contient également un pointeur sur la connexion à la base de données courante. N'oubliez pas que le code généré vous appartient, donc vous pouvez placer ce que vous voulez ici.
* Les fichiers de définition des modèles, un par modèle (donc un par table de la base de données cible).

<%= title("Définir un modèle simple") %>

Le fichier de modèle définit une structure pour accueillir une ligne de la table cible, des méthodes de validation et des fonctions callback optionnelles, qui permettent de définir des traitements liés aux modèles.

Prenons par exemple cette définition de table SQL, et créons la structure associée :

<%= partial("docs/db/models_sodas_sql.md") %>

Nous allons commencer en créant un nouveau fichier dans le dossier `models`, que l'on nommera `soda.go` (la convention est d'utiliser la forme au singulier du nom du modèle). Dans ce fichier, nous allons créer une structure pour la table `sodas` (la structure est aussi au singulier, puisqu'elle ne contient d'une seule ligne de la table) :

<%= partial("docs/db/models_sodas_go.md") %>

C'est tout ! Vous n'avez besoin de rien de plus pour travailler avec Pop ! Notez que pour chaque champ, nous avons défini un tag `db` qui correspond au nom du champ de la table, mais cela n'est pas obligatoire. Si vous ne fournissez pas de nom, il sera déterminé à partir de celui du champ de la structure.

<%= title("En utilisant le générateur") %>

<%= note() { %>
**Note pour les utilisateurs de Buffalo**: les commandes de `soda` sont intégrées à la commande `buffalo`, sous la commande `db`. À chaque fois que vous voulez utiliser une commande de `soda`, il vous suffit d'utiliser `buffalo db` à la place.
<% } %>

Écrire les fichiers de modèles à la main n'est pas de tout repos. Soda (et donc Buffalo, si vous avez bien suivi le chapitre sur Soda) fournit un générateur pour vous aider :

<%= partial("docs/db/model.md") %>

Vous pouvez supprimer le code généré pour le modèle à l'aide de la commande suivante :

```bash
$ soda destroy model [name]
```

Ou dans sa forme courte :

```bash
$ soda d m [name]
```

<%= title("Personnaliser les modèles") %>

### Définir les noms des champs

Par défaut, Pop utilisera le nom du champ de la structure pour trouver celui qui correspond dans la base de données.

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}
```

Pour la structure `User` définie ici, les noms de colonnes utilisés seront donc `ID`, `Email` et `Password`.

Vous pouvez changer ces noms en définissant explicitement celui à utiliser, à l'aide du tag `db`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Avec ce changement, les noms recherchés en base sont `id`, `email` et `password`.

À titre de comparaison, c'est très similaire à la manière dont fonctionne l'[association d'un formulaire à une structure](/docs/bind).

### Champs en lecture seule

Il est souvent nécessaire de lire un champ de la base de données, mais de ne pas vouloir écrire ce champ dans la base. C'est possible, grâce au tag `rw`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

Dans cet exemple, tous les champs seront récupérés **depuis la base**, et tous les champs **sauf** `Password` pourront être écrits dans la base.

### Champs en écriture seule

Les champs en écriture seule sont le pendant opposé des champs en lecture seule. Utilisez cette option pour les champs que vous voulez écrire, mais jamais récupérer de la base de données. Là encore, on utilise le tag `rw`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Ignorer des champs

Dans certains cas, vous voudrez faire en sorte que Pop ignore complètement un champ de la structure. Pensez à un champ temporaire en mémoire, ou qui sert dans une certaine logique de votre application.

Vous pouvez signaler à Pop qu'il faut ignorer ce champ, en utilisant le tag `db` avec la valeur `-`, comme dans l'exemple ci-après :

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```

Comme vous pouvez le voir, le champ `Password` possède le tag `db:"-"`, ce qui signifie que Pop **_n'enregistrera pas_** ce champ dans la base, et n'ira pas non-plus **_récupérer_** sa valeur.

### Modifier la clause de sélection d'un champ

Par défaut, Pop essaie de construire ses requêtes de selection pour une structure en utilisant tous les noms des champs.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Pour cette structure, la requête de sélection ressemblerait à ceci :

```sql
select id, email, password from users
```

Il est possible de changer le nom utilisé pour un champ en utilisant le tag `select`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" select:"password as p"`
}
```

La requête de sélection aurait alors cette tête :

```sql
select id, email, password as p from users
```

### Utiliser un nom de table personnalisé

Parfois, vous devrez travailler avec un schéma existant, ne respectant pas les conventions de nommage de Pop. Vous pouvez personnaliser le nom de la table associée à un modèle, en implémentant l'interface [`TableNameAble`](https://godoc.org/github.com/gobuffalo/pop#TableNameAble) :

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}

// TableName personnalise le nom de la table du modèle, utilisé par Pop.
func (u User) TableName() string {
  return "my_users"
}
```

<%= title("Modèles de vues") %>

Une [vue](https://fr.wikipedia.org/wiki/Vue_(base_de_donn%C3%A9es)) est un objet de base de données qui stocke le résultat d'une requête. Puisque cet objet agit comme une table en lecture seule, il est possible de le lier avec un modèle de Pop tout comme vous le feriez avec une table.

Si vous voulez utiliser un modèle avec plus d'une table, définir une vue est probablement la meilleure solution pour vous.

### Exemple

L'exemple suivant utilise la syntaxe de PostgreSQL. Nous allons commencer par créer deux tables :

```sql
-- Créer la table sodas
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    provider_id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);

-- Créer la table providers
CREATE TABLE providers (
    id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE providers ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

-- Créer une clé étrangère pour lier les deux tables
ALTER TABLE sodas ADD FOREIGN KEY (provider_id) REFERENCES providers(id);
```

Puis continuons en créant la vue :

```sql
CREATE VIEW sodas_with_providers AS
SELECT s.id, s.created_at, s.updated_at, p.label AS provider_label, s.label
FROM sodas s
LEFT JOIN providers p ON p.id = s.provider_id;
```

Comme la vue est considérée comme une table par Pop, terminons en déclarant un nouveau modèle :

```sql
type Soda struct {
	ID                   uuid.UUID    `db:"id" rw:"r"`
	CreatedAt            time.Time    `db:"created_at" rw:"r"`
	UpdatedAt            time.Time    `db:"updated_at" rw:"r"`
	Label                string       `db:"label" rw:"r"`
	ProviderLabel        string       `db:"provider_label" rw:"r"`
}
```

Comme nous l'avons appris dans ce chapitre, chaque attribut de la structure possède un tag lecture seule `rw:"r"`. Puisqu'une vue est un objet en lecture seule, cela évite de laisser passer une opération d'écriture; avant même d'atteindre la base de données.

<%= title("Contenu lié") %>

* [Migrations](/fr/docs/db/migrations) - Écrire des migrations de base de données.
* [Requêtage](/fr/docs/db/querying) - Lire des données depuis votre base de données.