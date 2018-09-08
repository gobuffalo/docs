<% seoDescription("Fizz") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "fizz", "pop", "DSL"]) %>

<%= h1("Fizz") %>

Fizz est langage dédié (DSL) commun pour migrer des bases de données. Il essaie d'être aussi agnostique aux bases de données que possible.
C'est le langage de définition des [migrations](/fr/docs/db/migrations) utilisé par défaut par Pop.

## Créer une table

``` javascript
create_table("users") {
  t.Column("email", "string", {})
  t.Column("twitter_handle", "string", {"size": 50})
  t.Column("age", "integer", {"default": 0})
  t.Column("admin", "bool", {"default": false})
  t.Column("company_id", "uuid", {"default_raw": "uuid_generate_v1()"})
  t.Column("bio", "text", {"null": true})
  t.Column("joined_at", "timestamp", {})
}

create_table("todos") {
  t.Column("user_id", "integer", {})
  t.Column("title", "string", {"size": 100})
  t.Column("details", "text", {"null": true})
  t.ForeignKey("user_id", {"users": ["id"]}, {"on_delete": "cascade"})
}
```

La fonction `create_table` génère une colonne `id` de type `integer` en mode auto-incrément. Il est possible de changer ce type pour utiliser le type [`UUID`](https://github.com/gobuffalo/uuid) :

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

Deux colonnes de type `timestamp` seront aussi générées par défaut ; `created_at` et `updated_at`.

La méthode `t.Column` accepte les arguments suivants : nom de la colonne, le type du champ, et finalement les options que vous souhaitez attribuer à cette colonne.

#### <a name="column-info"></a> Types « courants » :

* `string`
* `text`
* `timestamp`, `time`, `datetime`
* `integer`
* `bool`
* `uuid`

Tout autre type sera passé à la base de données sans aucune transformation.

Par exemple, pour PostgreSQL, vous pouvez passer le type `jsonb` et il sera supporté. Si vous faites la même chose avec SQLite par contre, il vous criera dessus !

#### Options supportées :

* `size` - La taille de la colonne. Par exemple, si vous voulez un `varchar(50)` avec PostgreSQL, vous procéderiez de la façon suivante : `t.Column("column_name", "string", {"size": 50})`
* `null` - Par défaut, les colonnes ne peuvent pas être `null`.
* `default` - La valeur par défaut de la colonne. La valeur par défaut est `null`.
* `default_raw` - La valeur par défaut, si vous souhaitez utiliser une fonction de base de données (comme `NOW()`).
* `after` - (MySQL seul) Ajouter la colonne après une autre colonne existante. `example: {"after":"created_at"}`
* `first` - (MySQL seul) Ajouter la colonne en première position dans la table. `example: {"first": true}`

#### Désactiver l'horodatage automatique

```javascript
create_table("users") {
  t.Column("id", "uuid", {primary: true})
  // ...
  // Désactiver la création automatique des colonnes created_at et updated_at
  t.DisableTimestamps()
}
```

or

```javascript
create_table("users", {timestamps: false}) {
  t.Column("id", "uuid", {primary: true})
  // ...
}
```

## Supprimer une table

``` javascript
drop_table("table_name")
```

## Renommer une table

``` javascript
rename_table("old_table_name", "new_table_name")
```

## Ajouter une colonne

``` javascript
add_column("table_name", "column_name", "string", {})
```

Voir [ci-dessus](#column-info) pour plus d'informations sur les types de colonnes et les options disponibles.

## Modifier une colonne

``` javascript
change_column("table_name", "column_name", "string", {})
```

## Renommer une colonne

``` javascript
rename_column("table_name", "old_column_name", "new_column_name")
```

## Supprimer une colonne

``` javascript
drop_column("table_name", "column_name")
```

## Add an Index

#### Options supportées

* `name` - Défaut : `table_name_column_name_idx`
* `unique`

### Index simple

``` javascript
add_index("table_name", "column_name", {})
```

### Indexes multi-colonnes

``` javascript
add_index("table_name", ["column_1", "column_2"], {})
```

### Indexes uniques

``` javascript
add_index("table_name", "column_name", {"unique": true})
```

### Nom des indexes

``` javascript
add_index("table_name", "column_name", {}) # name => table_name_column_name_idx
add_index("table_name", "column_name", {"name": "custom_index_name"})
```

## Renommer un index

``` javascript
rename_index("table_name", "old_index_name", "new_index_name")
```

## Supprimer un index

``` javascript
drop_index("table_name", "index_name")
```

## Ajouter une clef étrangère

```javascript
add_foreign_key("table_name", "field", {"ref_table_name": ["ref_column"]}, {
    "name": "optional_fk_name",
    "on_delete": "action",
    "on_update": "action",
})

```

#### Options supportées

* `name` - Par défaut : `table_name_ref_table_name_ref_column_name_fk`
* `on_delete` - `CASCADE`, `SET NULL`, ...
* `on_update`

**Note** : `on_update` et `on_delete` ne sont pas encore supportés par CockroachDB.

## Supprimer une clef étrangère

```javascript
drop_foreign_key("table_name", "fk_name", {"if_exists": true})
```

#### Options supportées

* `if_exists` - Ajoute une condition `IF EXISTS`


## SQL pur

``` javascript
sql("select * from users;")
```

## Exécuter une commande externe

Parfois, pendant une migration, vous souhaiterez exécuter une commande shell :

```javascript
exec("echo hello")
```
