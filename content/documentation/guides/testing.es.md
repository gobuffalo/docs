---
name: Pruebas
seoDescription: Cómo probar una aplicación de Buffalo
seoKeywords: ["buffalo", "go", "golang", "testing"]
aliases:
  - /docs/testing
  - /en/docs/testing
---

# Pruebas

## Ejecutando Pruebas

El ejecutador de pruebas de Buffalo se asegurará que tu entorno de prueba sea correcto y las ejecutará.

Por ejemplo, si se usa Pop (database), Primero intentará de configurar tu base de datos de prueba usando el esquema de la base de datos de desarrollo. Si no existe (por ejemplo CI), ejecutará las migraciones en la base de datos de prueba.

El ejecutador de pruebas tambien se asegurará de ignorar el temido directorio `vendor`.

```bash
$ buffalo test

dropped database authrecipe_test
created database authrecipe_test
dumped schema for authrecipe_development
loaded schema for authrecipe_test
go test -p 1 github.com/gobuffalo/authrecipe github.com/gobuffalo/authrecipe/actions github.com/gobuffalo/authrecipe/grifts github.com/gobuffalo/authrecipe/models
?   	github.com/gobuffalo/authrecipe	[no test files]
ok  	github.com/gobuffalo/authrecipe/actions	0.640s
?   	github.com/gobuffalo/authrecipe/grifts	[no test files]
ok  	github.com/gobuffalo/authrecipe/models	0.327s
```

### Ejecutando una sola prueba

{{< since "0.10.2" >}}

La depuración de una prueba específica es una tarea difícil, si debes ejecutar todas las pruebas existentes. Puedes usar el flag `-m` para ejecutar un solo metodo de prueba:
```bash
$ buffalo test -m "FooMethod"
```

Esto iterará a través de todos los paquetes y ejecutará cualquier prueba que coincida con `FooMethod` en cualquier paquete.

{{< since "0.14.10" >}}

Para limitar a un solo paquete, especifica el nombre del paquete:

```bash
$ buffalo test models -m "FooMethod"
```

## Suites de prueba

Buffalo usa la librería [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) para crear suites de prueba.

Cuando ejecutamos un test que es parte de una suite de prueba, lo siguiente está disponible para la prueba:

* La aplicación, `as.App`.
* La base de datos, `as.DB` (si se está usando Pop).
* La sesión, `as.Session`.
* Las afirmaciones de pruebas [`github.com/stretchr/testify/require`](https://github.com/stretchr/testify).
* La librería de pruebas de HTTP [`github.com/gobuffalo/httptest`](https://github.com/gobuffalo/httptest).

## Ejemplo de Prueba

```go
func (as *ActionSuite) Test_WidgetsResource_Create() {
  // setup a Widget model
  w := &models.Widget{Name: "My Widget"} // make a POST /widgets request
  res := as.HTML("/widgets").Post(w)
  // assert that the response status code was 302 as.Equal(302, res.Code)
  // retrieve the first Widget from the database
  err := as.DB.First(w)
  as.NoError(err)
  as.NotZero(w.ID)
  // assert the Widget title was saved correctly
  as.Equal("My Widget", w.Name)
  // assert the redirect was sent to the place expected
  as.Equal(fmt.Sprintf("/widgets/%s", w.ID), res.Location())
}
```

## Fixtures

{{< since "0.12.0" >}}

A menudo, es útil cargar datos de muestra en la base de datos al comienzo de una prueba. Por ejemplo, si necesitas tener un usuario en la base de datos para registrar a una persona en la aplicación, o necesitas algunos datos en la base de datos para probar que una ruta genera las respuestas correctamente. Los fixtures nos ayudan a resolver fácilmente estos problemas.

los fistures son archivos `.toml` que se ubican en el directorio `fixtures`. Puedes tener tantos fixtures como desees y se pueden nombrar como quieras.

Por ejemplo, si tenemos un simple modelo `Widget` que se muestra en una página que se ve así:

```go
type Widget struct {
  ID          uuid.UUID
  CreatedAt   time.Time
  UpdatedAt   time.Time
  Name        string
  Description string
}
```

Podemos crear un fixture como este:

```toml
[[scenario]]
name = "lots of widgets"

  [[scenario.table]]
    name = "widgets"

    [[scenario.table.row]]
      id = "\<%= uuidNamed("widget") %>"
      name = "This is widget #1"
      description = "some widget body #1"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"

    [[scenario.table.row]]
      id = "\<%= uuid() %>"
      name = "This is widget #2"
      description = "some widget body #2"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"
```

Cuando ejecutamos nuestra suite, estos dos registros se crearán en nuestra base de datos de prueba y luego podremos probar usando estos registros.
Todo lo que necesitas hacer para cargar el fixture es ubicarlo por su nombre con `ActionSuite.LoadFixture`.

```go
func (as *ActionSuite) Test_WidgetsResource_List() {
  as.LoadFixture("lots of widgets")
  res := as.HTML("/widgets").Get()

  as.Equal(200, res.Code)
  body := res.Body.String()
  as.Contains(body, "This is widget #1")
  as.Contains(body, "This is widget #2")
}
```

Puedes encontrar más información detallada de los fixtures en el [README del repo de gobuffalo/suite](https://github.com/gobuffalo/suite#fixtures-test-data).

## Acceso a la sesión

Ser capaz de manipular la sesión para realizar pruebas es muy importante. Las suites de prueba en Buffalo te dan acceso a una sesión de prueba que puedes usar.

consulta [https://github.com/gobuffalo/authrecipe](https://github.com/gobuffalo/authrecipe) para un ejemplo más detallado.

```go
func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
  // get a user from the DB

  // set the user ID onto the session
  as.Session.Set("current_user_id", user.ID)

  res := as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged in"
  as.Contains(res.Body.String(), "Sign Out")

  // clear the session
  as.Session.Clear()
  res = as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged out"
  as.Contains(res.Body.String(), "Sign In")
}
```

## Reportes de cobertura

{{< note >}}
La siguiente caracteristica requiere de **Go 1.10** o superior.
Go cover no es compatible con el operador `./...` en versiones anteriores e intentar usarlo generará un error.
{{< /note >}}

Es posible generar un reporte de cobertura de pruebas con Buffalo especificando el flag `-coverprofile` de la siguiente manera:

```bash
$ buffalo test -coverprofile=c.out ./...

created database authrecipe_test
loaded schema for authrecipe_test
INFO[0010] go test -p 1 -coverprofile=c.out ./...
?       github.com/gobuffalo/authrecipe [no test files]
ok      github.com/gobuffalo/authrecipe/actions 2.770s  coverage: 76.9% of statements
?       github.com/gobuffalo/authrecipe/grifts  [no test files]
ok      github.com/gobuffalo/authrecipe/models  2.609s  coverage: 71.4% of statements
```
