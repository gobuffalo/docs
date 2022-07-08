---
Name: Sesiones
seoDescription: "Sesiones en Buffalo"
seoKeywords: ["buffalo", "go", "golang", "http", "session"]
weight: 8
aliases:
  - /docs/sessions
  - /es/docs/sessions
---

# Sesiones

Una sesión HTTP es un almacenamiento de datos no persistente, que se destruye al cerrar el navegador (en la configuración por defecto del navegador). Puede utilizarse para almacenar mensajes flash, o cualquier dato temporal específico del usuario. Utilice [cookies](/es/documentación/request_handling/cookies) en su lugar si necesita un almacenamiento más persistente del lado del cliente.

La sesión está disponible directamente desde el `buffalo.Context` dentro de un handler.

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

## El tipo Session

El tipo `buffalo.Session` tiene todo lo necesario para trabajar con una sesión durante una petición. Bajo las cubiertas Buffalo utiliza el paquete [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) para gestionar la sesión.

```go
type Session
  // Clear a session of all values
  func (s *Session) Clear()
  // Delete a specific value from the session
  func (s *Session) Delete(name interface{})
  // Get a value from the session
  func (s *Session) Get(name interface{}) interface{}
  // GetOnce gets a value from the current session and then deletes it.
  func (s *Session) GetOnce(name interface{}) interface{}
  // Save a session
  func (s *Session) Save() error
  // Set a value on the session
  func (s *Session) Set(name, value interface{})
```



## Almacenamiento de la sesión

Por defecto, Buffalo configurará un almacenamiento de sesiones utilizando [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

Esto se puede cambiar al configurar una nueva aplicación de Buffalo utilizando la opción `SessionStore`:

```go
app = buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

La variable ENV `SESSION_SECRET` debe ser establecida antes de ejecutar la aplicación. Si no se establece, verá una advertencia en sus registros de que su sesión no está asegurada.

Para más información sobre esto, consulte la documentación de [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).





## Almacenamiento de tipos complejos

Generalmente se considera que **NO** es una buena práctica almacenar tipos complejos en una sesión. Hay muchas razones para ello, pero se recomienda almacenar el ID de un tipo, en lugar del valor "completo".

Si necesitas almacenar un tipo complejo, como un `struct`, primero tendrás que registrar el tipo con el paquete [`encoding/gob`](https://golang.org/pkg/encoding/gob/).

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

## Guardar una sesión

Buffalo guarda automáticamente la sesión por ti, para que no tengas que hacerlo tú. Si se produce un error al guardar la sesión, Buffalo devolverá un error a través del proceso normal de [Manejo de errores](/es/documentation/request_handling/errors).


## Null Sessions para APIs

Cuando se construyen servidores API, el almacenamiento de sesiones de cookies por defecto es indeseable. El tipo `sessions.Null` es el reemplazo recomendado para el almacenamiento de sesiones por defecto.

```go
app = buffalo.New(buffalo.Options{
  Env:          ENV,
  SessionStore: sessions.Null{},
  SessionName: "_coke_session",
})
```

Cuando se ejecuta `buffalo new` con el flag `--api` la sesión por defecto se establece en `sessions.Null`.

