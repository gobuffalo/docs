---
name: Asociaciones y Relaciones
seoDescription: "Como manejar asociaciones y relaciones con Pop?"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "associations", "relations", "entity"]
weight: 51
aliases:
  - /docs/db/relations
  - /es/docs/db/relations
---

# Asociaciones y Relaciones

Las asociaciones son la forma de Pop de definir **una relación entre dos objetos en la base de datos**. En este capítulo, aprenderás como definir asociaciones usando tags de estructura; y cómo manipularlos con el modificador `Eager`.

{{< vimeo 253683926>}}

## Ejempl

```go
type User struct {
  ID           uuid.UUID
  Email        string
  Password     string
  Books        Books     `has_many:"books" order_by:"title asc"`
  FavoriteSong Song      `has_one:"song" fk_id:"u_id"`
  Houses       Addresses `many_to_many:"users_addresses"`
}

type Book struct {
  ID      uuid.UUID
  Title   string
  Isbn    string
  User    User        `belongs_to:"user"`
  UserID  uuid.UUID
}

type Song struct {
  ID      uuid.UUID
  Title   string
  UserID  uuid.UUID   `db:"u_id"`
}

type Address struct {
  ID           uuid.UUID
  Street       string
  HouseNumber  int
}

type Books []Book
type Addresses []Address
```

## Tags de estructura disponibles

Usando el [ejemplo](#ejemplo) anterior, el código a continuación, es una lista de los tags de estructura disponiblesy como usarlos.

* `has_many`: Este tag es usado para discribir relaciones [uno a muchos](https://en.wikipedia.org/wiki/One-to-many_(data_model)) en la base de datos. En el ejemplo, El tipo `User` define una relación con una lista de tipo `Books` mediante el uso del tag `has_many`, lo que significa que un `User` puede tener varios `Books`. Cuando consultamos ne la base de datos, Pop cargará todos los registros de la tabla `books` que tengan una columna llamada `users_id`, o la columna especificada con `fk_id` que coincida con el valor de `User.ID`.

* `belongs_to`: Este tag se usa para describir el propietario en la relación. Un propietario representa una dependencia altamente acoplada entre el modelo y el campo de asociación de destino donde el tag `belongs_to` fue definido. Este tag se usa principalmente para indicar que el modelo posee su "existencia" en el campo de asociación `belongs_to`. En el ejemplo de arriba, el tipo `Book` usa `belongs_to` para indicar que su propietario es el tipo `User`. cuando consultamos en la base de datos, Pop cargará un registro de la tabla `users` con un `id` que coincida con el valor del campo `Book.UserID`.

* `has_one`: Este tag se usa para describir las relaciones [uno a uno](https://en.wikipedia.org/wiki/One-to-one_(data_model)) en la base de datos. En el ejemplo de arriba, hay una `FavoriteSong` dentro de todas los registros de canciones que el tipo `User` le gusta más. Cuando consultamos en la base de datos, Pop cargará un registro de la tabla `songs` que tiene una columna llamada `user_id`, una columna especificada con `fk_id` que coincide con el valor del campo `User.ID`.

* `many_to_many`: Este tag se usa para describir relaciones [muchos a muchos](https://en.wikipedia.org/wiki/Many-to-many_(data_model)) en la base de datos. En el ejemplo de arriba, la relacion entre `User` y la lista de tipo `Addresses` existe para indicar que un `USer` puede ser propietario de muchas `Houses` y una `House` ser propiedad de muchos `Users`. Es importante notar que el valor para el tag `many_to_many` es la tabla que conecta ambos lados en la relación; en el ejemplo de arriba, este valor se define como `users_addresses`. Cuando consultamos en la base de datos, Pop cargará todos los registros de la tabla `addresses` mediante la tabla asociativa `users_addresses`. La tabla `users_addresses` **DEBE** tener definidas las columnas `address_id` y `user_id` para coincidir con los valores de los campos `Address.ID` y `User.ID`. Tambien puedes definir un tag `fk_id` que se usará en la asociacion de destino, es decir, la tabla `addresses`.

* `fk_id`: Este tag se puede usar para definir el nombre de la columna en la asociación de destino que coincida con el `ID` del modelo. En el ejemplo de arriba, `Song` tiene una columna llamada `u_id` que hace referencia al id de la tabla `users`. Cuando cargamos `FavoriteSong`, `u_id` será usado en lugar de `user_id`.

* `order_by`: Este tag se puede usar en combinacion con los tags `has_many` y `many_to_many` para indicar el orden de la asociación cuando se cargan. La forma de uso es: `order_by:"<column_name> <asc | desc>"`

## Cargando asociaciones

Pop actualmente proporciona dos modos para cargar asociaciones; cada modo afectará la forma en que Pop carga las asociaciones y consultas a la base de datos.

[Eager](#modo-eager). El modo por defecto. Al habilitar este modo, pop realizará "n" consultas para cada asociación definida en el modelo. Esto significa más llamadas a la base de datos para no afectar el uso de la memoria.

[EagerPreload](#modo-eagerpreload). Modo opcional. Al habilitar este modo, Pop ejecutará una consulta para cada asociación definida en el modelo. Este modo llamará a la base de datos con una frecuencia reducida sacrificando más espacio en memoria.

* `pop.SetEagerMode`: Pop permite habilitar alguno de estos modos globalmente, lo que afectará el rendimiento del manejo de **TODAS** las consultas. Usa `EagerDefault` o `EagerPreload` como parametro para activar alguno de estos modos.

* `tx.EagerPreload | q.EagerPreload`: Pop permite a los desarrolladores controlar en cuales situaciones quienen que Pop ejecute cualquiera de estos modos cuando sea necesario. Este método activará el modo `EagerPreload` solo para la consulta en acción.

* `tx.Eager | q.Eager`: Pop permite a los desarrolladores controlar en cuales situaciones quienen que Pop ejecute cualquiera de estos modos cuando sea necesario. Este metodo activará el modo `Eager` solo para la consulta en acción.

## Modo Eager

El método [`pop.Connection.Eager()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Eager) le indica a Pop que cargue las asociaciones para un modelo una vez que ese modelo se carge desde la base de datos.
Este modo ejecutará "n" consultas para cada asociacion definida en el modelo.

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
```

```go
u := Users{}

// Carga todas las asociaciones para cada usuario registrado. Por ejemplo `Books`, `Houses` y `FavoriteSong`
err := tx.Eager().All(&u)
```

El modo `Eager` va a:

1. Cargar todos los users.

```sql
SELECT * FROM users;
```

2. Recorrer cada usuario y cargar sus asociaciones

```sql
SELECT * FROM books WHERE user_id = 1
```

```sql
 SELECT * FROM books WHERE user_id = 2
```

```sql
SELECT * FROM books WHERE user_id = 3
```

## Modo EagerPreload
El método [`pop.Connection.EagerPreload()`](https://github.com/gobuffalo/pop/pull/146/files#diff-f49e947ec94f65964b0845af2b62845aR180) le indica a pop que cargue las asociaciones para un modelo una vez que el modelo es cargado desde la base de datos. Este modo llamará a la base de datos con una frecuencia reducida sacrificando más espacio en memoria.

```go
for i := 0; i < 3; i++ {
  user := User{ID: i + 1}
  tx.Create(&user)
}

for i := 0; i < 3; i++ {
  book := Book{UserID: i +1}
  tx.Create(&book)
}
```

```go
u := Users{}

// Carga todas las asociaciones para cada usuario registrado. Por ejemplo `Books`, `Houses` y `FavoriteSong`
err := tx.EagerPreload().All(&u)
```

El modo `EagerPreload` va a:

1. Cargará todos los usuarios.

```sql
SELECT * FROM users;
```

2. Carga las asociaciones para todos los usuarios en una sola consulta.

```sql
SELECT * FROM books WHERE user_id IN (1, 2, 3)
```

## Cargar asociaciones específicas

Por defecto, `Eager` y `EagerPreload` cargarán todas las asociaciones asignadas para el modelo. PAra especificar cuales asociaciones deben ser cargadas, puedes pasar los nombres de esos campos a los metodos `Eager` o `EagerPreload` y solo se cargarán esas asociaciones.

```go
// Cargar solo las asociaciones `Books` para los usuarios con nombre `Mark`

u := []User{}

// Usando modo `Eager`
err := tx.Eager("Books").Where("name = 'Mark'").All(&u)

// Usando modo `EagerPreload`
err := tx.EagerPreload("Books").Where("name = 'Mark'").All(&u)
```
Pop también te permite cargar asociaciones anidadas usando el caracter `.` para concatenarlos. Echale un vistazo al siguiente ejemplo.

```go
// Cargará todos los `Books` para `u` y para cada `Book`, cargará el usuario que será el mismo que `u`

u := User{}

// Usando modo `Eager`
tx.Eager("Books.User").First(&u)

// Usando modo `EagerPreload`
tx.EagerPreload("Books.User").First(&u)
```

```go
// Cargará todos los `Books` para `u` y para cada `Book` cargará todos los `Writers` y para cada `Writer` cargará la asociación `Book`.

u := User{}

// Usando modo `Eager`
tx.Eager("Books.Writers.Book").First(&u)

// Usando modo `EagerPreload`
tx.EagerPreload("Books.Writers.Book").First(&u)
```

```go
// Cargará todos los `Books` para `u` y para cada `Book` cargará todos los `Writers`. Y también cargará `FavoriteSong` de `u`.

u := User{}

// Usando modo `Eager`
tx.Eager("Books.Writers").Eager("FavoriteSong").First(&u)

// Usando modo `EagerPreload`
tx.EagerPreload("Books.Writers").EagerPreload("FavoriteSong").First(&u)
```

## Cargar asociaciones para un modelo existente

El método [`pop.Connection.Load()`](https://godoc.org/github.com/gobuffalo/pop#Connection.Load) toma una estructura de modelo, que ha sido llenado desde la base de datos, y una lista opcional de asociaciones para cargar.


```go
u := User{}

// Carga todas las asociaciones para `u`. Por ejemplo `Books`, `Houses` y `FavoriteSong`
tx.Load(&u)


// Carga solo las asociaciones `Books` para `u`
tx.Load(&u, "Books")
```

El método `Load` no recuperará el `User` de la base de datos, solo sus asociaciones.

## Creación anidada plana

Pop te permite crear los modelos y sus asociaciones con otros modelos en un solo paso de manera predeterminada. Ya no necesitas crear cada asociación por separado. Pop incluso creará registros de las tablas asociativas para las asociaciones `many_to_many`.


Suponiendo las siguientes piezas de pseudocódigo:

```go
book := Book{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}
tx.Create(&book)

song := Song{Title: "Don't know the title"}
tx.Create(&song)

addr := Address{HouseNumber: 1, Street: "Golang"}
tx.Create(&addr)

user := User{
  Name: "Mark Bates",
  Books: Books{Book{ID: book.ID}},
  FavoriteSong: song,
  Houses: Addresses{
    addr,
  },
}

err := tx.Create(&user)
```

1. Notarás que `Books` es una asociacion `has_many` y se dará cuenta que para actualizar realmente cada `Book`, primero necesitará tener el `User ID`. Por lo tanto, procede a guardar los datos de `User` para que pueda recuperar el **ID** y lueo usar ese ID para llenar el campo `UserID` en todos los `Books`. Actualiza todos los `Books` en la base de datos usando sus `ID`s para orientarlos.

2. `FavoriteSong` es una asociacion `has_one` y usa la misma logica descrita en la asociación `has_many`. Dado que los datos de `User` se guardaron previamente antes de actuializar todos los `Books` afectados, Se sabe que el `User` tiene un `ID`, por lo que llena su campo `UserID` con ese valor y `FavoriteSong` se actualiza en la base de datos.

3. `Houses` en este ejemplo es una relacion `many_to_many` y se tendrá que tratar con dos tablas en este caso: `users` y `addresses`. Debido a que `User` ya estaba almacenado, ya se tiene su `ID`. A continación, se usarán los `ID`'s pasados con los `Addresses` para crear las entradas correcpondientes en la tabla asociativa.

Para una asociación `belongs_to` como se muestra en el siguiente ejemplo, llena su campo `UserID` antes de guardarse en la base de datos.

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: user,
}


tx.Create(&book)
```

## Creacion con Eager

Pop tambien te permite crear modelos e integrar la creación de sus asociaciones en un solo paso

Suponiendo las siguientes piezas de pseudocódigo:

```go
user := User{
  Name: "Mark Bates",
  Books: Books{{Title: "Pop Book", Description: "Pop Book", Isbn: "PB1"}},
  FavoriteSong: Song{Title: "Don't know the title"},
  Houses: Addresses{
    Address{HouseNumber: 1, Street: "Golang"},
  },
}

err := tx.Eager().Create(&user)
```

1. Notará que `Books` es una asociación `has_many` y se dará cuenta de que para almacenar cada `Book`, primero necesitará obtener el `UserID`. Por lo tanto, procede primero a guardar/crear los datos de `User` para que pueda recuperar un **ID** y luego usar esa ID para rellenar el campo `UserID` en cada `Book` en `Books`. Después guarda todos los `Books` en la base de datos.

2. `FavoriteSong` es una asociación `has_one` y usa la misma lógica descrita en la asociación `has_many`. Dado que los datos del `User` se guardaron previamente antes de crear todos los `Books`, se sabe que el `User` tiene un `ID`, por lo que llena su campo `UserID` con ese valor y luego `FavoriteSong` se guarda en la base de datos.

3. `Houses` en este ejemplo es una relación `many_to_many` y tendrá que tratar con dos tablas, en este caso: `users` y `addresses`. Primero deberá guardar todas las `Addresses` en la tabla `addresses` antes de guardarlas en la tabla asociativa. Debido a que `User` ya está guardado, ya tiene un `ID`.
    * Este es un caso especial a tratar, ya que este comportamiento es diferente de todas las demás asociaciones, se soluciona implementando la interfaz `AssociationCreatableStatement`, todas las demás asociaciones implementan por defecto la interfaz `AssociationCreatable`.



Para una asociación `belongs_to` como la que se muestra en el siguiente ejemplo, primero deberá crear `User` para recuperar su valor **ID** y luego rellenar su campo `UserID` antes de guardarlo en la base de datos.

```go
book := Book{
   Title:      "Pop Book",
   Description: "Pop Book",
   Isbn:        "PB1",
   User: User{
        Name: nulls.NewString("Larry"),
   },
}

tx.Eager().Create(&book)
```

En caso de que alimentes la creación con `Eager` con modelos asociados que ya existen, en lugar de crear duplicados o actualizar su contenido, simplemente creará/actualizará las asociaciones con ellos.

## Siguientes pasos

* [Relaciones uno a uno](/es/documentation/database/relations-one-to-one)
* [Relaciones uno a muchos](/es/documentation/database/relations-one-to-many)
