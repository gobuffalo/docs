---
name: Modelos
seoDescription: Modelos
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "models"]
weight: 10
aliases:
  - /docs/db/models
  - /es/docs/db/models
---

# Modelos

Pop, es un ORM, te permite traducir tablas de bases de datos a estructuras de Go. De esta forma, puedes manipular estructuras de go en vez de escribir sentencias SQL. El código de Go que maneja esta parte es llamado "modelos", como referencia a la arquitectura MVC.

En este capitulo, aprenderás a trabajar con modelos a mano; y como mejorar el flujo de trabajo usando los generadores proporcionados.

## Directorio de Modelos

Los archivos de modelos de Pop estan guardados en la carpeta `models`, en la raíz de tu proyecto (Mira el capítulo de [La estructura de directorio](/es/documentation/getting_started/directory-structure) para más informacion sobre la forma que usa Buffalo organiza tus archivos).

Este directorio contiene:

* Un archivo `models.go`, el cual define las partes comunes para cada modelo definido. También contiene un puntero configurado a la conexión. Recuerda que el código es tuyo, así que puedes colocar lo que quieras aquí.

* Archivos de definición de modelos, uno para cada modelo (es decir, uno por tabla de base de datos a la que deseas acceder de esta forma).

## Define un modelo simple

Un archivo de modelo define una asignación para la tabla de la base de datos, métodos de validación y callbacks de Pop si desea agregar más lógica relacionada con el modelo.

Tomemos la siguiente definicion de una tabla en SQL, y escribamos una estructura

```sql
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    label character varying(255)
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);
```
Iniciemos creando un nuevo archivo dentro del directorio de `models` llamado `soda.go` (la convención usada aqui es tomar la forma singular de la palabra). En este archivo, crearemos la estructura para la tabla `sodas` (tambien el nombre de la estructura es en singular, ya que contendrá un solo registro de la tabla)

```go
package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gofrs/uuid"
)

type Soda struct {
	ID                   uuid.UUID    `db:"id"`
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            time.Time    `db:"updated_at"`
	Label                nulls.String `db:"label"`
}
```

Eso es todo! No necesitas nada mas para trabajar con Pop! Ten en cuenta que para cada campo de la estructura, definimos un tag de pop llamado `db` que coincide con el nombre de la columna de la tabla de la base de datos, pero no es obligatorio. Si no proporcionas un nombre, Pop usará el nombre del campo de la estructura para generar uno.

## Usando el Generador

{{<note>}}
**Nota para usuarios de Buffalo**: Los comandos de `Soda` estan adheridos de los comandos de `Buffalo`, detrás del espacio de nombres de `pop`. Asi que cada vez que desees usar un comando de `soda`, solo ejecuta `buffalo pop` en su lugar. No necesitas instalar la CLI de `soda`.
{{</note>}}

Escribir los archivos a mano no es la forma mas eficiente para trabajar. Soda (y Buffalo, si viste el capitulo sobre Soda) proporciona un generador para ayudarte:

```bash
$ soda g model --help
Generates a model for your database

Usage:
  soda generate model [name] [flags]

Aliases:
  model, m

Flags:
  -h, --help                    help for model
      --migration-type string   sets the type of migration files for model (sql or fizz) (default "fizz")
      --models-path string      the path the model will be created in (default "models")
  -s, --skip-migration          Skip creating a new fizz migration for this model.
      --struct-tag string       sets the struct tags for model (xml/json/jsonapi) (default "json")

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

Puedes remover el el modelo generado ejecutando:

```bash
$ soda destroy model [name]
```

O la version corta:

```bash
$ soda d m [name]
```

## Manejo de Nulos

Si necesitas guardar valores `NULL` en tu tabla, tendrás que usar tipos especiales

{{<warning>}}
No puedes almacenar un valor `null` en un campo de una estructura de tipo `int`
{{</warning>}}

La libreria de [sql estandar de Go](https://golang.org/pkg/database/sql) proporciona los tipos especiales para ese caso de uso, como like [`sql.NullBool`](https://golang.org/pkg/database/sql/#NullBool) o [`sql.NullInt64`](https://golang.org/pkg/database/sql/#NullInt64).

Si necesitas mas que la libreria ofrece, puedes usar el paquete [gobuffalo/nulls](https://github.com/gobuffalo/nulls) el cual proporciona mas tipos de nulos y un mejor manejo de la serialización y deserialización de JSON.

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password nulls.String
}
```

## Personalizar Modelos

### Asignando Campos del modelo

Por defecto cuando intentamos asignar una estructura a una tabla de la base de datos, Pop usará el nombre del campo en la estructura como nombre de la columna en la base de atos

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}
```

Con la estructura de arriba, se asume que los nombres de las columnas en la tabla de `users` en la base de datos son `ID`, `Email` y `Password`.

Estos nombres se se pueden cambiar usando el tag de estructura `db`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Ahora se espera que los nombres de las columnas en la tabla `users` de la base de datos sean `id`, `email`, and `password`.

Esto es muy similar a cómo funciona [bind](/docs/bind).

Se puede usar cualquier tipo que implemente las interfaces [Scanner](https://golang.org/pkg/database/sql/#Scanner) y [Valuer](https://golang.org/pkg/database/sql/driver/#Valuer), sin embargo, para que no tengas que escribirlas tu mismo, es recomendable que uses los siguientes tipos:


| Tipo Base             | Nullable        | Slice/Array  |
|:----------------------|:----------------|:------------ |
|int                    |nulls.Int        |slices.Int    |
|int32                  |nulls.Int32      | ------       |
|int64                  |nulls.Int64      | ------       |
|uint32                 |nulls.UInt32     | ------       |
|float32                |nulls.Float32    | ------       |
|float, float64         |nulls.Float64    | slices.Float |
|bool                   |nulls.Bool       | ------       |
|[]byte                 |nulls.ByteSlice  | ------       |
|string                 |nulls.String     |slices.String |
|uuid.UUID              |nulls.UUID       |slices.UUID   |
|time.Time              |nulls.Time       | ------       |
|map[string]interface{} | ---------       |slices.Map    |

{{<note>}}
**Nota**: Cualquier campo de tipo `slices.Map` deberá inicializarse antes de aplicar `Bind` o acceder a este.
```go
widget := &models.Widget{Data: slices.Map{}}
```
{{</note>}}

### Campos de Solo lectura

A menudo es necesario leer un campo de la base dedatos, pero no desamos escribir ese campo en la base de datos. Esto se puede hacer usando la etiqueta de estructura `rw:"r"`

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

En este ejemplo todos los campos pueden ser leidos **desde** la base de datos, pero todos los campos **exepto** `password` podrán escribirse en la base de datos

### Campos de Solo escritura

Los campos de solo lectura son lo opuesto a los campos solo lectura. Estos son campos que deseas escribir en la base de datos, pero nunca deseas recuperar. Esto lo puedes hacer usando la etiquiera de estructura `rw:"w"`

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Saltando campos de modelo

A veces, necesitas indicarle a `Pop` que cierto campo no se debe almacenar en la tabla de la base de datos. Tal vez sea solo un campo que se usa en memoria u otra razón lógica relacionada con la aplicación que estas creando.

La forma que le indicas a Pop sobre esto es usando la etiqueta de estructura `db` en tu modelo con el valor `-` como en el siguiente ejemplo:

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```
Como puedes ver, el campo `Password` esta marcado como `db:"-"`, esto significa que Pop no *_guardará_* ni *_recuperará_* este campo de la base de datos.

### Cambiando la cláusula Select para una columna

Por defecto, cuando se intenta de general la consulta `select` para usa estructura, se usan todos los nombres de los campos para generarla.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

La sentencia `select` resultante se vería asi:

```sql
select user.id, users.email, users.password from users
```

Podemos cambiar la sentencia para una columna usando el tag `select`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" select:"password as p"`
}
```

La sentencia `select` resultante se vería asi:

```sql
select users.id, users.email, password as p from users
```

### Usando un nombre de tabla personalizado

A veces, necesitaras tener que trabajar con un esquema existente, con nombres de tablas que no coinciden con las convenciones de Pop. Puedes sobreescribir este comportamiento y proporcionar un nombre de tabla personalizado implementanto la interfaz [`TableNameAble`](https://godoc.org/github.com/gobuffalo/pop#TableNameAble).

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}

// TableName overrides the table name used by Pop.
func (u User) TableName() string {
  return "my_users"
}
```

Se recomienda usar un receptor de valor en lugar de un puntero si la estructura se usa como valor en cualquier parte del código.

```go
// recommended:
func (u User) TableName() string {

// can cause issues:
func (u *User) TableName() string {
```

### UNIX Timestamps

{{< since "v4.7.0" >}}

Si declaras los campos `CreatedAt` y `UpdatedAt` en tu estructura de modelo (o se crean por defecto cuando usas el generador de modelo), Pop los manejará por ti. Eso significa que cuando creasd un registro en la base de datos, al campo `CreatedAt` se le establecerá la fecha y hora actual, y al campo `UpdateAt` será establecido cada vez que actualices un registro existente.

Estos campos definidos de tipo `time.Time`, pero tu puedes definirlos como `int` y manejarlos como UNIX timestamps.

```go
type User struct {
  ID        int    `db:"id"`
  CreatedAt int    `db:"created_at"`
  UpdatedAt int    `db:"updated_at"`
  FirstName string `db:"first_name"`
  LastName  string `db:"last_name"`
}
```

Si usas las migraciones fizz, asegurate de definir estos campos por tu cuenta y deshabilita las marcas de tiempo fecha y hota por defecto.

```go
create_table("users") {
  t.Column("id", "int", {primary: true})
  t.Column("created_at", "int")
  t.Column("updated_at", "int")
  t.Column("first_name", "string")
  t.Column("last_name", "string")
  t.DisableTimestamps()
}
```

## Modelos para vistas

Una [vista](https://en.wikipedia.org/wiki/View_(SQL)) es un objeto de colección de base de datos el cual almacena el resultado de una consulta. Dado que este objeto actua como una tabla solo lectura, puedes mapearlos como modelosde Pop como una tabla.

Si deseas usar un modelo con más de una tabla, definir una vista es probablemente la mejor solución para ti.

### Ejemplo

El siguiente ejemplo usa la sintaxis de PostgreSQL. Iniciaremos con crear dos tablas:

```sql
-- Creando la tabla sodas
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    provider_id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);

-- Creando la tabla providers
CREATE TABLE providers (
    id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE providers ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

-- Creando una llave foránea entre las dos tablas
ALTER TABLE sodas ADD FOREIGN KEY (provider_id) REFERENCES providers(id);
```

A continuación, crea una vista a partir de las dos tablas:

```sql
CREATE VIEW sodas_with_providers AS (
  SELECT
    s.id,
    s.created_at,
    s.updated_at,
    p.label AS provider_label, s.label
  FROM sodas s
  LEFT JOIN providers p ON p.id = s.provider_id;
)
```

Dado que Pop considera la vista como una tabla, terminaremos declarando un nuevo modelo:

```go
type SodasWithProvider struct {
	ID            uuid.UUID `db:"id" rw:"r"`
	CreatedAt     time.Time `db:"created_at" rw:"r"`
	UpdatedAt     time.Time `db:"updated_at" rw:"r"`
	Label         string    `db:"label" rw:"r"`
	ProviderLabel string    `db:"provider_label" rw:"r"`
}
```

Como aprendimos en este capítulo, cada atributo de la estructura tiene una etiqueta de solo lectura `rw:"r"`. Dado que una vista es un objeto de solo lectura, evita cualquier operación de escritura antes de acceder a la base de datos.

## Related Content

* [Migraciones](/es/documentation/database/migrations) - Escribe migraciones a la base de datos.
* [Consultas](/es/documentation/database/querying) - Consulta datos de tu base de datos.
