## Binding to a Struct

The [`c.Bind`](https://godoc.org/github.com/gobuffalo/buffalo#Context) allows form elements to be bound to a struct, but it can also attach uploaded files to the struct. To do this, the type of the struct attribute **must** be a `binding.File` type.

In the example below you can see a model, which is configured to have a `MyFile` attribute that is of type `binding.File`. There is an `AfterCreate` callback on this example model that saves the file to disk after the model has been successfully saved to the database.

```go
// models/widget.go
type Widget struct {
  ID        uuid.UUID    `json:"id" db:"id"`
  CreatedAt time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
  Name      string       `json:"name" db:"name"`
  MyFile    binding.File `db:"-" form:"someFile"`
}

func (w *Widget) AfterCreate(tx *pop.Connection) error {
  if !w.MyFile.Valid() {
    return nil
  }
  dir := filepath.Join(".", "uploads")
  if err := os.MkdirAll(dir, 0755); err != nil {
    return errors.WithStack(err)
  }
  f, err := os.Create(filepath.Join(dir, w.MyFile.Filename))
  if err != nil {
    return errors.WithStack(err)
  }
  defer f.Close()
  _, err = io.Copy(f, w.MyFile)
  return err
}
```

**Note:** The `MyFile` attribute is not being saved to the database because of the `db:"-"` struct tag.
