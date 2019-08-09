<% seoDescription("Consultar una base de datos con Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "consulta", "pop", "buscadores"]) %>

<%= h1("Consultas (_Querying_)") %>

En este capítulo, aprenderá cómo acceder y recuperar datos de su base de datos usando Pop.

### Buscar por ID

```go
user := User{}
err := db.Find(&user, id)
```

### Encontrar todo

```go
users := []User{}
err := db.All(&users)
err = db.Where("id in (?)", 1, 2, 3).All(&users)
```

#### Encontrar el último

```go
// Last() orders by created_at
user := models.User{}
err := tx.Last(&user)
```

### Encontrar dónde

```go
users := []models.User{}
query := db.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = tx.Where("id in (?)", 1, 2, 3).All(&users)
```

#### Usando consultas `in`

```go
err = db.Where("id in (?)", 1, 2, 3).All(&users)
err = db.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

Desafortunadamente, por una variedad de razones, no se puede usar una consulta `y` en la misma llamada `Where` como una consulta `in`.

```go
// no funciona:
err = db.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)
// funciona:
err = db.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Seleccionar columnas específicas

`Select` le permite cargar columnas específicas de una tabla. Útil cuando no quiere que se carguen todas las columnas de una tabla en una consulta.

```go
err = db.Select("name").All(&users)
// SELECT name FROM users

err = db.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = db.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Unir consultas

```go
// page: page number
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
