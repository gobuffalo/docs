---
name: "Querying"
seoDescription: "Querying a Database with Pop"
seoKeywords: ["buffalo", "go", "golang", "database", "querying", "pop", "finders"]
weight: 21
aliases:
  - /docs/db/querying
  - /pt/docs/db/querying
---
# Querying

In this chapter, you'll learn how to retrieve data from your database using Pop.

### Find By ID

```go
user := User{}
err := models.DB.Find(&user, id)
```

### Find All

```go
users := []User{}
err := models.DB.All(&users)
err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
```

### Find All with Order

```go
// To retrieve records from the database in a specific order, you can use the Order method
users := []User{}
err := models.DB.Order("id desc").All(&users)
```

### Find Last

```go
// Last() orders by created_at
user := models.User{}
err := tx.Last(&user)
```

### Find Where

```go
users := []models.User{}
query := models.DB.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = tx.Where("id in (?)", 1, 2, 3).All(&users)
```

### Using `in` Queries

```go
err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
err = models.DB.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

{{<warning>}}
Unfortunately, for a variety of reasons you can't use an `and` query in the same `Where` call as an `in` query.
{{</warning>}}

```go
// does not work:
err := tx.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)

// works:
err := tx.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Select specific columns

`Select` allows you to load specific columns from a table. Useful when you don't want all columns from a table to be loaded in a query.

```go
err = tx.Select("name").All(&users)
// SELECT name FROM users

err = tx.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = tx.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Join Query

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

### Count records

```go
query := tx.Q()
count, err := query.Count(&models.User{})
```

```go
query := tx.Q()
count, err := query.Where("name = ?", "John").Count(&models.User{})
```

```go
query := tx.Q()
count, err := query.CountByField(&models.User{}, "first_name")
// Equals to
count, err := query.Count(&models.User{},, "first_name")
```

### Pop to SQL

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
-- The original query is in one line
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
