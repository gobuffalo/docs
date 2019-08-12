<% seoDescription("Modelos") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "ORM", "pop", "modelos"]) %>

<%= h1("Modelos") %>

Pop, como ORM, le permite traducir tablas de base de datos en estructuras Go. De esta forma, puede manipular las estructuras de Go en lugar de escribir frases de SQL. El código Go que gestiona esta parte se denomina "modelos", como referencia a la arquitectura MVC.

En este capítulo, aprenderá a trabajar con modelos a mano; y cómo mejorar su flujo de trabajo utilizando los generadores brindados.

## El directorio de modelos

Los archivos de modelos Pop se almacenan en el directorio `models`, en la raíz de su proyecto (consulte el capítulo [Estructura de directorio](/es/docs/getting-started/directory-structure) para obtener más información sobre la forma en que Buffalo organiza sus archivos).

Este directorio contiene:

* Un archivo `models.go`, que define las partes comunes para cada modelo definido. También contiene un _pointer_ a la conexión configurada. Recuerda que el código es suyo, así que puede colocar lo que quiera aquí.
* Archivos de definición de modelo, uno para cada modelo (por lo que uno por cada tabla de base de datos a la que desee acceder de esta manera).

## Definir un modelo simple

Un archivo de modelo define una asignación para la tabla de la base de datos, los métodos de validación y las devoluciones de llamada Pop si desea agregar más lógica relacionada con el modelo.

Tomemos la siguiente definición de tabla SQL y escriba una estructura que coincida:

<%= partial("en/docs/db/models_sodas_sql.md") %>

Comenzaremos por crear un nuevo archivo en el directorio `models`, llamado `soda.go` (la convención utilizada aquí es tomar la forma singular de la palabra). En este archivo, crearemos la estructura para la tabla `sodas` (la estructura también es singular, ya que contendrá una sola línea de la tabla):

<%= partial("en/docs/db/models_sodas_go.md") %>

¡Eso es! ¡No necesita nada más para trabajar con Pop! Tenga en cuenta que, para cada campo de tabla, definimos una etiqueta `pop` que coincide con el nombre del campo, pero no es obligatorio. Si no proporciona un nombre, Pop usará el nombre del campo _struct_ para generar uno.

## Uso del generador

<%= note() { %>
**Nota para los usuarios de Buffalo**: los comandos `soda` están incrustados en el comando` buffalo`, detrás del espacio de nombres `pop`. Así que cada vez que quiera usar un comando de `soda`, simplemente ejecute `buffalo pop` en su lugar.<% } %>

Escribir los archivos a mano no es la manera más eficiente de trabajar. Soda (y Buffalo, si siguió el capítulo sobre Soda) proporciona un generador para ayudarlo a:

<%= partial("en/docs/db/model.md") %>

Puede eliminar el modelo generado ejecutando:

```bash
$ soda destroy model [name]
```

O en forma corta:

```bash
$ soda d m [name]
```

## Personalizar modelos

### Asignación de campos de modelo

De forma predeterminada, al tratar de asignar una estructura a una tabla de base de datos, Pop usará el nombre del campo en la estructura como el nombre de la columna en la base de datos.

```go
type User struct {
  ID       uuid.UUID
  Email    string
  Password string
}
```

Con la estructura anterior, se supone que los nombres de las columnas en la base de datos son `ID`, `Email` y `Password`.

Estos nombres de columna se pueden cambiar utilizando la etiqueta struct `db`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

Ahora se espera que los nombres de las columnas sean `id`, `email` y `password`.

Esto es muy similar a cómo funciona [el enlace de formulario](/es/docs/bind).

Se puede usar cualquier tipo que se adhiera a las interfaces [Scanner](https://golang.org/pkg/database/sql/#Scanner) y [Valuer](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces, sin embargo, para que no tenga que escribir estos usted mismo, se recomienda que se adhiera a los siguientes tipos:

| Base type             | Nullable        | Slice/Array |
|-----------------------|:---------------:|------------:|
|int                    |nulls.Int        |slices.Int   |
|int32                  |nulls.Int32      | ------      |
|int64                  |nulls.Int64      | ------      |
|uint32                 |nulls.UInt32     | ------      |
|float32                |nulls.Float32    | ------      |
|float, float64         |nulls.Float64    |slices.Float |
|bool                   |nulls.Bool       | ------      |
|[]byte                 |nulls.ByteSlice  | ------      |
|string                 |nulls.String     |slices.String|
|uuid.UUID              |nulls.UUID       |slices.UUID  |
|time.Time              |nulls.Time       | ------      |
|map[string]interface{} | ---------       |slices.Map   |

### Campos de solo lectura

A menudo es necesario leer un campo de una base de datos, pero no desea escribir ese campo en la base de datos. Esto se puede hacer usando la etiqueta struct `rw`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"r"`
}
```

En este ejemplo, todos los campos se leerán **de** la base de datos y todos los campos, **excepto** para `Password` podrán escribir en la base de datos.

### Campos de solo escritura

Los campos solo de escritura son el reverso de los campos de solo lectura. Estos son campos que desea escribir en la base de datos, pero nunca recuperar. De nuevo, esto hace uso de la etiqueta struct `rw`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" rw:"w"`
}
```

### Saltar los campos del modelo

En ocasiones, debe informar al Pop que cierto campo no debe almacenarse en la tabla de la base de datos. Tal vez sea solo un campo que utilice en memoria u otro motivo lógico relacionado con la aplicación que está creando.

La forma en que le informa a Pop sobre esto es usando la etiqueta struct `db` en su modelo y estableciendo que sea `-` como en el siguiente ejemplo:

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"-"`
}
```

Como puede ver, el campo `Password` está marcado como `db:"-"` lo que significa que Pop **no almacenará ** ni **recuperar** este campo de la base de datos.

### Cambiar la cláusula Select para una columna

El valor predeterminado, al intentar construir la consulta `select` para una estructura es usar todos los nombres de campo para construir una consulta.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password"`
}
```

La declaración `select` resultante se vería así:

```sql
select id, email, password from users
```

Podemos cambiar la declaración de una columna usando la etiqueta `select`.

```go
type User struct {
  ID       uuid.UUID `db:"id"`
  Email    string    `db:"email"`
  Password string    `db:"password" select:"password as p"`
}
```

La declaración `select` resultante se vería así:

```sql
select id, email, password as p from users
```

### Uso de un nombre de tabla personalizado

A veces, tendrá que trabajar con un esquema existente, con los nombres de tabla que no coinciden con las convenciones Pop. Puede anular este comportamiento y proporcionar un nombre de tabla personalizado implementando la interfaz [`TableNameAble`](https://godoc.org/github.com/gobuffalo/pop#TableNameAble):

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

## Modelos de vistas

Una [vista](https://en.wikipedia.org/wiki/View_(SQL)) es un objeto de colección de base de datos que almacena el resultado de una consulta. Como este objeto actúa como una tabla de solo lectura, puede asignarlo a los modelos Pop como una tabla.

Si desea utilizar un modelo con más de una tabla, definir una vista es probablemente la mejor solución para usted.

### Ejemplo

El siguiente ejemplo usa la sintaxis de PostgreSQL. Comenzaremos por crear dos tablas:

```sql
-- Crear una tabla de sodas
CREATE TABLE sodas (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    provider_id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE sodas ADD CONSTRAINT sodas_pkey PRIMARY KEY (id);

-- Crear una tabla de providers
CREATE TABLE providers (
    id uuid NOT NULL,
    label character varying(255) NOT NULL
);

ALTER TABLE providers ADD CONSTRAINT providers_pkey PRIMARY KEY (id);

-- Crear unn Private Key entre las dos tablas
ALTER TABLE sodas ADD FOREIGN KEY (provider_id) REFERENCES providers(id);
```

Luego crea una vista desde las dos tablas:

```sql
CREATE VIEW sodas_with_providers AS
SELECT s.id, s.created_at, s.updated_at, p.label AS provider_label, s.label
FROM sodas s
LEFT JOIN providers p ON p.id = s.provider_id;
```

Como Pop considera que la vista es una tabla, terminemos por declarar un nuevo modelo:

```sql
type Soda struct {
	ID                   uuid.UUID    `db:"id" rw:"r"`
	CreatedAt            time.Time    `db:"created_at" rw:"r"`
	UpdatedAt            time.Time    `db:"updated_at" rw:"r"`
	Label                string       `db:"label" rw:"r"`
	ProviderLabel        string       `db:"provider_label" rw:"r"`
}
```

Como aprendimos en este capítulo, cada atributo de la estructura tiene una etiqueta de solo lectura `rw:"r"`. Como una vista es un objeto de solo lectura, impide cualquier operación de escritura antes de llegar a la base de datos.

## Contenido relacionado

* [Migraciones](/es/docs/db/migrations) - Escribe migraciones de bases de datos.
* [Consulta](/es/docs/db/querying) - Consulta dato de su base de datos.
