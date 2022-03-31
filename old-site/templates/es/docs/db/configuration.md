<% seoDescription("¿Cómo configurar mi base de datos con Pop?") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "ORM", "pop", "configuración"]) %>

# Configuración de la base de datos

La configuración Pop se gestiona mediante un archivo `database.yml`, ubicado en la raíz de su proyecto. Este archivo se genera para usted si utiliza Buffalo &#8211; si elige usar Pop &#8211; y contiene una configuración básica para la base de datos que seleccionó en la generación con el _flag_ `--db-type`. PostgreSQL se usa como base de datos predeterminada.

Aquí hay una configuración de muestra generada para una nueva aplicación basada en PostgreSQL:

```yaml
development:
  dialect: postgres
  database: myapp_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}
```

Puede ver tres conexiones definidas:
* `development` es el utilizado cuando su aplicación se ejecuta en modo dev.
* `test` sirve para ejecutar las pruebas de integración.
* `production` es la configuración que usará en la aplicación final, en el servidor.

Por supuesto, puede configurar cualquier conexión nueva que desee, pero Buffalo no los seleccionará por defecto.

## Generador

<%= partial("en/docs/db/soda_buffalo_note.md") %>

Puede generar un archivo de configuración predeterminado utilizando el comando `init`:

```bash
$ soda g config
```

El valor predeterminado generará un archivo `database.yml` en el directorio actual para una base de datos PostgreSQL. Puede anular el tipo de base de datos usando el indicador `-t` y pasar cualquiera de los tipos de bases de datos compatibles:` postgres`, `cockroach`,` mysql`, o `sqlite3`.

## Personalizar la ubicación de configuración

El archivo de configuración Pop &#8211; `database.yml` &#8211; se puede encontrar:
* En su raíz del proyecto (predeterminado).
* En el directorio `config /`, en la raíz de su proyecto.

Si desea poner su archivo de configuración en otra ubicación, puede usar [`AddLookupPaths`](https://godoc.org/github.com/gobuffalo/pop#AddLookupPaths).

También puedes personalizar el nombre del archivo:

```go
pop.ConfigName = "my_pop_config.yml"
```

## Env vs configuración detallada

<%= note() { %>
Tenga en cuenta que el archivo `database.yml` también es una plantilla Go, que permite el uso de la sintaxis de la plantilla Go. Hay dos funciones especiales que se incluyen, `env` y` envOr`.
<% } %>

Como puede ver, tiene dos formas de configurar una nueva conexión:
* El que usa la conexión `development` es el más detallado. Le permite establecer cada parámetro disponible, uno por uno.
* El usado por las conexiones `test` y` production` es un poco diferente: usa una variable (ver las marcas `{{}}`?) Para establecer el valor, y el _helper_ `envOr`.

El _helper_ `envOr` intenta obtener un valor de una variable de entorno y, de forma predeterminada, el segundo valor. Por ejemplo:

```yaml
envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"
```

Intenta obtener el valor `TEST_DATABASE_URL` del entorno y su valor predeterminado es `postgres://postgres:postgres@127.0.0.1:5432/myapp_test`.

De esta forma, puede proporcionar un valor predeterminado para fines de desarrollo y permitir reconfigurar la configuración de la base de datos desde una variable de entorno.

<%= warning() { %>
El parámetro `url` para una conexión anulará cualquier otro parámetro de conexión. Asegúrese de establecer todas las configuraciones que desee desde el URL.
<% } %>

Para obtener detalles adicionales, consulte la documentación de [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

** Asegúrese de haber configurado este archivo correctamente antes de trabajar con Pop! **

## Opciones disponibles

### database

El nombre de la base de datos a usar.

### dialect

El dialecto de la base de datos a usar con la conexión. Los valores aceptados son:
* MySQL driver: "mysql"
* PostgreSQL driver: "postgres", "postgresql" or "pg"
* Cockroach driver: "cockroach", "cockroachdb" or "crdb"
* SQLite driver: "sqlite" or "sqlite3"

### encoding

<%= sinceVersion("4.6.0") %>

Esta opción actualmente solo es compatible con **dialecto de mysql**. Esta codificación se usará para crear la base de datos (si la creas usando `soda`), y como el parámetro` collation` para la cadena de conexión. Si se omite esta opción, el valor predeterminado es `utf8mb4_general_ci`.

```yaml
development:
  dialect: mysql
  options:
    encoding: "utf8_general_ci"
```

### host

La dirección de host de la base de datos para conectarse.

### password

La contraseña para el usuario que usa para conectarse a la base de datos.

### port

El puerto de host de la base de datos para la base de datos.

**Valores predeterminados**:

| Driver    | Port  |
|-----------|-------|
| PostgreSQL| 5432  |
| MySQL     | 3306  |
| Cockroach | 26257 |

### user

El usuario a usar para conectarse a la base de datos.