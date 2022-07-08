---
name: Configuración de base de datos
seoDescription: "Como configurar my base de datos con Pop?"
seoKeywords: ["buffalo", "go", "golang", "database", "ORM", "pop", "configuration"]
weight: 3
aliases:
  - /docs/db/configuration
  - /es/docs/db/configuration
---

# Configuración de la base de datos

La configuración de `pop` se gestiona mediante un archivo `database.yml`, ubicado en la raíz de tu proyecto. Este archivo se genera para ti si usas Buffalo, si eliges usar `pop`, y contiene una configuración básica para la base de datos que seleccionaste en la generación con el flag `--db-type`. PostgreSQL se considera como predeterminado.

Aquí hay un ejemplo configuración generada para una nueva aplicación basada en PostgreSQL:

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

Puedes ver 3 conexiones definidas:

* `development` Es usada cuando tu app se ejecura en modo `dev`.
* `test` sirve para ejecutar test de intergración.
* `production` es la configuración que usarás en la version final de la pp, en el servidor.

Por supuesto, puedes configurar cualquier nueva conexión que desees, pero Buffalo no los eligirá por defecto.

## Generador

{{<note>}}
**Nota para usuarios de Buffalo**: Los comandos de `Soda` estan adheridos de los comandos de `Buffalo`, detrás del espacio de nombres de `pop`. Asi que cada vez que desees usar un comando de `soda`, solo ejecuta `buffalo pop` en su lugar. No necesitas instalar la CLI de `soda`.
{{</note>}}

Puedes generar el archivo de configuración pr defecto con el comando init:

```bash
$ soda g config
```

Por defecto, se generará un archivo `database.yml` en el directorio actual para una base de datos `PostgreSQL`. Puedes sobreescribir el tipo de base de datos usando el flag `-t` y pasando cualquiera de los tipos de base de datos admitidos: `postgres`, `cockroach`, `mysql` o `sqlite3`.

## Ubicación de archivo de configuración

El archivo de configuración de Pop `database.yml` puede ser encontrado ya sea:
* En tu raíz del proyecto (defecto).
* En el directorio `config/`, en la raíz de tu proyecto

Si deseas colocar tu archivo de configuración en alguna otra ubicación, puedes usar [`AddLookupPaths`](https://godoc.org/github.com/gobuffalo/pop#AddLookupPaths).

Tambien puedes personalizar el nombre del archivo:

```go
pop.ConfigName = "my_pop_config.yml"
```

## Env vs Configuración detallada

{{<note>}}
Nota que el archivo `database.yml` es tambien un template de Go, para que puedas usar la sintaxis del template de Go. Hay dos funciones especiales que estan incluidas, `env` y `envOr`
{{</note>}}

Como puedes ver, tienes dos maneras de configurar una nueva conexión:
* El que usa la conexión `development` es el más detallado. Te permite configurar cada parámetro disponible, uno por uno.
* El que usa las conexiones de `prueba` y `producción` es un poco diferente: usa una variable (ver las marcas `{{ }}`) para establecer el valor, y el helper `envOr`.

El helper `envOr` trata de obtener un valor de una variable de entorno, y por defecto utiliza un segundo valor, por ejemplo:

```yaml
envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"
```

Intenta obtener el valor de `TEST_DATABASE_URL` del entorno y el valor por defecto es `postgres://postgres:postgres@127.0.0.1:5432/myapp_test`.

Esta forma, puedes proporcionar un valor por defecto para fines de desarrollo, y permitir reconfigurar la configuración de la base de datos desde una variable de entorno!

{{<warning>}}
El parámetro `url` para una conexión va a sobreescribir cualquier otro parámetro de conexión. Asegúrate de establecer todas la configuraciones que desses para la cadena URL.
{{</warning>}}

Para detalles adicionales, revisa la documentación para [github.com/gobuffalo/pop](https://github.com/gobuffalo/pop).

**Asegúrate de haber configurado correctamente este archivo antes de trabajar con Pop!**

## Opciones disponibles

### database

El nombre de la base de datos a usar.

### dialect


El dialecto de la base de datos que se usará con la conexión. Los valores aceptados son:
* Controlador MySQL: "mysql"
* Controlador PostgreSQL: "postgres", "postgresql" o "pg"
* Controlador Cockroach: "cockroach", "cockroachdb" o "crdb"
* Controlador SQLite: "sqlite" o "sqlite3"

### driver

{{< since "4.11.2" >}}

Usa esta opción para personalizar el controlador de la base de datos y sobreescribir el por defecto usado por Pop.

Aquí esta la lista de controladores SQL por defecto:
* MySQL: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* PostgreSQL: [github.com/lib/pq](https://github.com/lib/pq)
* Cockroach DB: [github.com/cockroachdb/cockroach-go/crdb](https://github.com/cockroachdb/cockroach-go/tree/master/crdb)
* SQLite: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

### encoding

{{< since "4.6.0" >}}

Esta opción es actualmente soportada por el **dialecto mysql**. Este codificador sera usado para crear la base de dator (si la creas usando `soda`, y como la ``)

Actualmente, esta opción solo es compatible con el **dialecto mysql**. Esta codificación se usará para crear la base de datos (si la creas usando `soda`), y como el parámetro `collation` para la cadena de conexión. Si se omite esta opción, el valor predeterminado es `utf8mb4_general_ci`.


```yaml
development:
  dialect: mysql
  options:
    encoding: "utf8_general_ci"
```

### host

La dirección del host de la base de datos a la cual conectarse.

### port

Es el puerto del host para la base de datos.

**Predeterminado**:

| Driver     | Port    |
|:-----------|:--------|
| PostgreSQL | 5432    |
| MySQL      | 3306    |
| Cockroach  | 26257   |

### user

El usuario para conectar a la base de datos.

### password

Es la contraseña para el usuario que usas para conectarte a la base de datos.
