<% seoDescription("Querying a Database with Pop") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "querying", "pop", "finders"]) %>

<%= h1("Querying") %>

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

#### Find Last

```go
// Last() orders by created_at
user := models.User{}
err := models.DB.Last(&user)
```

### Find Where

```go
users := []models.User{}
query := models.DB.Where("id = 1").Where("name = 'Mark'")
err := query.All(&users)

err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
```

#### Using `in` Queries

```go
err = models.DB.Where("id in (?)", 1, 2, 3).All(&users)
err = models.DB.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

Unfortunately, for a variety of reasons you can't use an `and` query in the same `Where` call as an `in` query.

```go
// does not work:
err = models.DB.Where("id in (?) and foo = ?", 1, 2, 3, "bar").All(&users)
// works:
err = models.DB.Where("id in (?)", 1, 2, 3).Where("foo = ?", "bar").All(&users)
```

### Select specific columns
`Select` allows you to load specific columns from a table. Useful when you don't want all columns from a table to be loaded in a query.
```go
err = models.DB.Select("name").All(&users)
// SELECT name FROM users

err = models.DB.Select("max(age)").All(&users)
// SELECT max(age) FROM users

err = models.DB.Select("age", "name").All(&users)
// SELECT age, name FROM users
```

### Join Query

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
