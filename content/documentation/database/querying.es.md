---
name: Consultas
seoDescription: "Consulta en una base de datos con Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "querying", "pop", "finders"]
weight: 21
aliases:
  - /docs/db/querying
  - /es/docs/db/querying
---
# Consultas

en este capítulo, aprenderás cómo recuperar datos de tu base de datos usando Pop.

### Encontrar por ID

```go
user := User{}
err := tx.Find(&user, id)
```

### Encontrar todos

```go
users := []User{}
err := tx.All(&users)
```

```go
users := []User{}
err := tx.Where("id in (?)", 1, 2, 3).All(&users)
```

### Encontrar todos con orden

```go
users := []User{}
// Para recuperar registros de la base de datos en un orden especifico, puedes usar el metodo `Order`
err := tx.Order("id desc").All(&users)
```

#### Encontrar el último

```go
user := models.User{}
// El metodo `Last` ordena por la columna `created_at`
err := tx.Last(&user)
```

### Encontrar donde

```go
users := []models.User{}
query := tx.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)
```

```go
users := []models.User{}
err = tx.Where("id in (?)", 1, 2, 3).All(&users)
```

### Cláusula `in`

```go
users := []models.User{}
err := tx.Where("id in (?)", 1, 2, 3).All(&users)
```

```go
users := []models.User{}
err := tx.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

{{<warning>}}
Desafortunadamente, por una variedad de razones, no puedes usar el condicional `and` en la misma llamada `Where` que contenga una consulta `in`.
{{</warning>}}

```go
// no funciona:
err := tx.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)

// funciona:
err := tx.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Seleccionar columnas específicas

`Select` te permite cargar columnas específicas de una tabla. Útil cuando no deseas que todas las columnas de una tabla se carguen en la consulta.

```go
err = tx.Select("name").All(&users)
// SELECT name FROM users

err = tx.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = tx.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Cláusula `Join`

```go
// page: page number
// perPage: limit

roles := []models.UserRole{}

q := tx.Q()
q.LeftJoin("roles", "roles.id = user_roles.role_id")
q.LeftJoin("users u", "u.id = user_roles.user_id")
q.Where(`roles.name like ?`, name)
q.Paginate(page, perPage)

err := q.All(&roles)
```

### Contar registros

```go
count, err := query.Count(&models.User{})
```

```go
count, err := query.Where("name = ?", "John").Count(&models.User{})
```

```go
count, err := query.CountByField(&models.User{}, "first_name")
// Es igual a
count, err := query.Count(&models.User{},, "first_name")
```

### Pop a SQL

```go
q := tx.Q()
q.LeftJoin("roles", "roles.id = user_roles.role_id")
q.LeftJoin("users u", "u.id = user_roles.user_id")
q.Where(`roles.name like ?`, "john")
q.Paginate(1, 20)

popModel := &pop.Model{Value: models.UserRole{}}
cols := []string{"user_roles.*", "roles.name as role_name", "u.first_name", "u.last_name"}

sql, args := tx.Q().ToSQL(popModel, cols...)
```

{{<codetabs>}}
{{<tab "sql">}}
```sql
-- La consulta original está en una linea
SELECT
    user_roles.*,
    roles.name as role_name,
    u.first_name,
    u.last_name
FROM
    user_roles AS user_roles
LEFT JOIN
    roles ON roles.id = user_roles.role_id
LEFT JOIN
    users u ON u.id = user_roles.user_id
WHERE
    roles.name like $1
LIMIT
    20
OFFSET
    0
```
{{</tab>}}
{{<tab "args">}}
```text
[john]
```
{{</tab>}}
{{</codetabs>}}
