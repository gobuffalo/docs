<% seoDescription("Requêter une base de données avec Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "base de données", "requêtage", "pop"]) %>

<%= h1("Requêtage") %>

Dans ce chapitre, vous allez apprendre comment récupérer des données depuis votre base de données, en utilisant Pop.

### Rechercher par ID

```go
user := User{}
err := db.Find(&user, id)
```

### Rechercher toutes les occurrences

```go
users := []User{}
err := db.All(&users)
err = db.Where("id in (?)", 1, 2, 3).All(&users)
```

### Rechercher toutes les occurrences et ordonner les résultats

```go
// Vous pouvez utiliser la méthode Order pour ordonner les résultats.
users := []User{}
err := db.Order("id desc").All(&users)
```

#### Rechercher le dernier résultat

```go
// Last() trie les résultats par date (created_at)
user := models.User{}
err := tx.Last(&user)
```

### Recherche conditionnelle

```go
users := []models.User{}
query := db.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = tx.Where("id in (?)", 1, 2, 3).All(&users)
```

#### Rechercher avec `IN`

```go
err = db.Where("id in (?)", 1, 2, 3).All(&users)
err = db.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

Malheureusement, pour certaines raisons, vous ne pouvez pas utiliser une requête `AND` en même temps que `IN`.

```go
// ne fonctionne pas :
err = db.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)
// fonctionne  :
err = db.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Sélectionner des colonnes spécifiques
`Select` vous permet de sélectionner des colonnes spécifiques d'une table. C'est pratique lorsque vous n'avez pas besoin de toutes les colonnes d'une table, et cela évite de charger des données pour rien.
```go
err = db.Select("name").All(&users)
// SELECT name FROM users

err = db.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = db.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Requête JOIN

```go
// page: numéro de page
// perpage: limit
roles := []models.UserRole{}
query := models.DB.LeftJoin("roles", "roles.id=user_roles.role_id").
  LeftJoin("users u", "u.id=user_roles.user_id").
  Where(`roles.name like ?`, name).Paginate(page, perpage)

count, _ := query.Count(models.UserRole{})
count, _ := query.CountByField(models.UserRole{}, "*")
sql, args := query.ToSQL(&pop.Model{Value: models.UserRole{}}, "user_roles.*",
  "roles.name as role_name", "u.first_name", "u.last_name")
err := models.DB.RawQuery(sql, args...).All(&roles)
```
