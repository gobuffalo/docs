---
name: Eventos
seoDescription: Escuchando por eventos en una aplicación de Buffalo
seoKeywords: ["buffalo", "go", "golang", "events", "plugins"]
aliases:
  - /docs/events
  - /es/docs/events
---

# Eventos

{{< since "0.13.0-beta.2" >}}

La librería de {{< doclink href="github.com/gobuffalo/events" message="events" >}} permite a las aplicaciones en Go, incluyendo aplicaciones de Buffalo, escuchen y emitan mensajes de eventos globales.

## Escuchar Eventos

Para iniciar a escuchar eventos, primero se debe registrar un {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}} con la librería {{< doclink href="github.com/gobuffalo/events" message="events"  >}}.

```go
func init() {
  _, err := events.Listen(func(e events.Event) {
    // do work
  })
}
```

Una vez registrada esta nueva función `listener`, se enviarán todos los eventos emitidos a través de la librería {{< doclink message="events" href="github.com/gobuffalo/events" >}}.

## Emitiendo Eventos

Cuando emitimos eventos, el atributo `Kind` debe ser una cadena única, pero constante. Es este atributo que los usuarios usarán para determinar cómo responder a los eventos que reciben.

Se recomienda asignar un espacio de nombres a este atributo como tal, con los eventos de error sean con el sufijo: `:err`.

```plain
"<package-name>:<additional-names>:<optional-error>"

"myapp:foo:start"
"myapp:foo:stop"
"mypkg:workers:bar:err"
```

Este patrón de nomenclatura facilita a los usuarios filtrar solo aquellos eventos que les interesan. Ve a [Filtrando Eventos](#filtering-events) para mas detalles.

---

Hay varias maneras de emitir un {{< doclink message="events#Event" href="github.com/gobuffalo/events#Event" >}} en tu código de Go. LAs funciones {{< doclink message="events#EmitError" href="github.com/gobuffalo/events#EmitError" >}} y {{< doclink message="events#EmitPayload" href="github.com/gobuffalo/events#EmitPayload" >}} aceptan un argumento de tipo `payload interface{}`. Se recomienda usar {{< doclink message="events#Payload" href="github.com/gobuffalo/events#Payload" >}} para datos de entrada; cualquier otro tipo pasado será convertido a {{< doclink message="events#Payload" href="github.com/gobuffalo/events#Payload" >}} con el argumento establecido en los datos de entrada con la llave `data`.

* {{< doclink message="events#Emit" href="github.com/gobuffalo/events#Emit" >}}

```go
func MyHandler(c buffalo.Context) error {
e := events.Event{
    Kind:    "coke:myhandler:hello",
    Message: "hi!",
    Payload: events.Payload{"context": c},
}
if err := events.Emit(e); err != nil {
    return err
}
return c.Render(200, r.HTML("index.html"))
}
```

* {{< doclink message="events#EmitError" href="github.com/gobuffalo/events#EmitError" >}}

```go
func MyHandler(c buffalo.Context) error {
  if err := events.EmitError("coke:myhandler:hello:err", errors.New("boom"), c); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

* {{< doclink message="events#EmitPayload" href="github.com/gobuffalo/events#EmitPayload" >}}

```go
func MyHandler(c buffalo.Context) error {
  p := events.Payload{
    "message": "hi!",
  }
  if err := events.EmitPayload("coke:myhandler:hello", p); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

## Filtrando Eventos

En la sección de [Emitiendo Eventos](#emitiendo-eventos) se describe la convención de nomenclaturas para {{< doclink message="events#Event.Kind" href="github.com/gobuffalo/events#Event.Kind" >}}. Al comprobar el valor de {{< doclink message="events#Event.Kind" href="github.com/gobuffalo/events#Event.Kind" >}}.


{{< codetabs >}}
{{< tab "Coincidencia directa" >}}
```go
events.Listen(func(e events.Event) {
  if e.Kind != buffalo.EvtRouteStarted {
    // do nothing
    return
  }
  // do work on the route event
})
```
{{< /tab >}}

{{< tab "Coincidencias con sentencia switch" >}}
```go
events.Listen(func(e events.Event) {
  switch e.Kind {
  case buffalo.EvtAppStart, buffalo.EvtAppStop:
    // do work
  case "buffalo:dev:build:finished":
    // do work
  default:
    // do nothing
  }
})
```
{{< /tab >}}

{{< tab "Coincidencias por eventos de error" >}}
```go
func init() {
  events.Listen(func(e events.Event) {
    if !e.IsError() {
      // do nothing
      return
    }
    // do something with e.Error
  })
}
```
{{< /tab >}}

{{< tab "Coincidencias en el prefijo" >}}
```go
events.Listen(func(e events.Event) {
  if !strings.HasPrefix(e.Kind, "buffalo:") {
    // do nothing
    return
  }
  // do work only on events emitted by Buffalo
})
```
{{< /tab >}}
{{< /codetabs >}}

## Dejar de escuchar Eventos

Cuando registramos un nuevo {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}}, se retorna una{{< doclink message="events#DeleteFn" href="github.com/gobuffalo/events#DeleteFn" >}}. Esta función debe conservarse y usarse cuando desees eliminar el listener agregado.

```go
deleteFn, err := events.Listen(func(e events.Event) {
  // do work
})
if err != nil {
  return err
}
defer deleteFn()
```
## Escuchando con Complementos

Para permitir que un complemento reciba una version JSON de los eventos emitidos, el complemento puede establecer el valor de {{< doclink message="events#Command.BuffaloCommand" href="github.com/gobuffalo/buffalo-plugins/plugins#Command.BuffaloCommand" >}} al evento al listar los comandos `disponibles` para el complemento.

{{< codetabs >}}
{{< tab "availableCmd" >}}
```go
// availableCmd
var availableCmd = &cobra.Command{
  Use:   "available",
  Short: "a list of available buffalo plugins",
  RunE: func(cmd *cobra.Command, args []string) error {
    plugs := plugins.Commands{
      {Name: "listen", UseCommand: "listen", BuffaloCommand: "events", Description: listenCmd.Short, Aliases: listenCmd.Aliases},
    }
    return json.NewEncoder(os.Stdout).Encode(plugs)
  },
}

```
{{< /tab >}}
{{< tab "listenCmd" >}}
```go
// listenCmd
var listenCmd = &cobra.Command{
  Use:   "listen",
  Short: "listens to github.com/gobuffalo/events",
  RunE: func(cmd *cobra.Command, args []string) error {
    if len(args) == 0 {
      return errors.New("must pass a payload")
    }

    e := events.Event{}
    err := json.Unmarshal([]byte(args[0]), &e)
    if err != nil {
      return errors.WithStack(err)
    }

    // do work with event
    return nil
  },
}
```
{{< /tab >}}
{{< /codetabs >}}

## Integración de una cola de mensajería

A menudo se desea tomar los eventos emitidos y enviarlos a una cola de mensajes, como Kafka o Redis, para que se procesen externamente. La librería {{< doclink message="events" href="github.com/gobuffalo/events" >}} no tiene una forma directa para esta funcionalidad, la forma mas directa para habilitar este comportamiento es registrar un {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}} que luego puede transferir el evento a la cola de mensajes correspondientes.

```go
events.Listen(func(e events.Event) {
  myMessageQ.DoWork(e)
})
```
## Eventos Conocidos

### Eventos de la Aplicación

Se sabe que Buffalo emite los siguientes eventos durante el ciclo de vida de la aplicación.

| constante                   | Cadena                       | Se emite cuando                                                                                                             | Datos de entrada                                                                                                                                                                                                                                                                                  |
| ---                         | ---                          | ---                                                                                                                         | ---                                                                                                                                                                                                                                                                                               |
| `buffalo.EvtAppStart`       | `"buffalo:app:start"`        | Se llama {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}}                          | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtAppStartErr`    | `"buffalo:app:start:err"`    | Ocurre un error llamando {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}}          | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtAppStop`        | `"buffalo:app:stop"`         | Se llama {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}}                            | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtAppStopErr`     | `"buffalo:app:stop:err"`     | Ocurre un error llamando {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}}            | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtRouteStarted`   | `"buffalo:route:started"`    | Se está procesando una ruta solicitada                                                                                      | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtRouteFinished`  | `"buffalo:route:finished"`   | Se completa una ruta solicitada                                                                                             | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtRouteErr`       | `"buffalo:route:err"`        | Hay un problema al manejar el procesamiento de una ruta                                                                     | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtWorkerStart`    | `"buffalo:worker:start"`     | Se llama {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}} y se inician los Workers | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtWorkerStartErr` | `"buffalo:worker:start:err"` | Se produce un error al iniciar los Workers                                                                                  | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtWorkerStop`     | `"buffalo:worker:stop"`      | Se llama {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}} y se detienen los Workers  | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtWorkerStopErr`  | `"buffalo:worker:stop:err"`  | Se produce un error al detener los Workers                                                                                  | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                                                                                                             |
| `buffalo.EvtFailureErr`     | `"buffalo:failure:err"`      | Algo no se puede procesar en absoluto. Es algo malo                                                                         | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}}                                                                                                        |


### Eventos de Buffalo Dev

Se sabe que los siguientes eventos son emitidos por `buffalo dev` durante el ciclo de vida del desarrollo.

| Cadena                         | Se emite cuando                    | Datos de entrada                                                                                                                                                                     |
| ---                            | ---                                | ---                                                                                                                                                                                  |
| `"buffalo:dev:raw"`            | Se modifica un archivo aplicable   | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}}                                                                                          |
| `"buffalo:dev:build:started"`  | Comienza una una construcción      | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}} <br> `cmd`: string of the `go build` command (example: `"go build foo"`)                 |
| `"buffalo:dev:build:finished"` | Se completa una construcción       | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}} <br> `pid`: PID of the newly running binary <br> `build_time`: the duration of the build |
| `"buffalo:dev:build:err"`      | Ocurre un error en la construcción | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}} <br> `cmd`: string of the `go build` command (example: `"go build foo"`)                 |
