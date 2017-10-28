# Skipping Model Fields

Sometimes you need to let `Pop` know that certain field should not be stored in the database table because it's just a field you use in-memory or other logical reason related with the application you're building

The way you let `Pop` know about this is by usind the `db` struct tag on your model and setting it to be `-` like the following example:

```go
type User struct {
    ID                  uuid.UUID   `db:"id"`
    Email               string      `db:"email"`
    Password            string      `db:"-"`
    EncryptedPassword   string      `db:"encrypted_password"`
}
```

As you may see the `Password` field is marked as `db:"-"` what means `Pop` will not store this field in the database.