---
name: Ejecuciones en segundo plano
weight: 12
aliases:
  - /docs/workers
  - /en/docs/workers
---
# Ejecuciones en segundo plano

Al construir aplicaciones complejas, a menudo es bueno poder ejecutar cosas en "segundo plano". Si bien Go proporciona bonitas funciones de concurrencia, como el famoso Goroutine, a menudo uno quiere ejecutarlas en diferentes máquinas, persistirlas usando Redis o cualquier número de posibles razones por las que un simple Goroutine no es suficiente.

{{<note>}}
Los Workers no se deben confundir con las [tareas](/es/documentation/guides/tasks): las tareas son herramientas síncronas, mientras que los ejecutores están destinados a ejecutarse de manera asíncrona.
{{</note>}}

## La interfaz Worker

Para usar ejecuciones en segundo plano, primero se debe cumplir la interfaz `worker.Worker`.

```go
type Worker interface {
  // Start the worker with the given context
  Start(context.Context) error
  // Stop the worker
  Stop() error
  // Perform a job as soon as possibly
  Perform(Job) error
  // PerformAt performs a job at a particular time
  PerformAt(Job, time.Time) error
  // PerformIn performs a job after waiting for a specified amount of time
  PerformIn(Job, time.Duration) error
  // Register a Handler
  Register(string, Handler) error
}
```

Actualmente hay dos implementaciones oficiales de esta interfaz:

* La primera es `worker.Simple`: Esta usa Goroutines al implementar la interfaz. Esto es excelente para aplicaciones simples, pero dado que las colas no persisten, cualquier trabajo en cola se perderá si el servidor se apaga. **Esta implementación está activada por defecto**.

* La otra implementación es la librería [`github.com/gobuffalo/gocraft-work-adapter`](https://github.com/gobuffalo/gocraft-work-adapter), la cual implementa la librería [`github.com/gocraft/work`](https://github.com/gocraft/work) usando Redis como tienda de respaldo.


### Implementaciones de la comunidad

Las siguientes implementaciones de Worker son proporcionadas por los usuarios de Buffalo (Soporte no oficial):


| Nombre | Autor | Descripción |
|:------|:--------|:-------------|
| [AMQP worker adapter](https://github.com/stanislas-m/amqp-work-adapter) | [@stanislas-m](https://github.com/stanislas-m) | Una implementación de Worker para usar con agentes compatibles con AMQP (como [RabbitMQ](https://www.rabbitmq.com/)).|

## El tipo de Job

Un Job es una unidad de trabajo dado para una implementación de Worker determinada.

```go
// Args are the arguments passed into a job
type Args map[string]interface{}

// Job to be processed by a Worker
type Job struct {
  // Queue the job should be placed into
  Queue string
  // Args that will be passed to the Handler when run
  Args Args
  // Handler that will be run by the worker
  Handler string
}
```

## Cómo usar tareas en segundo plano

Para ser poder usar tareas en segundo plano, deberás configurar un adaptador de ejecuciones, registrar controladores de trabajo y activarlos.

### Configurando un Adaptador de Ejecución

Cuando configuras tu aplicación, **puedes** asignar una implementación Worker en la opción `Worker`.

```go
// actions/app.go

import "github.com/gobuffalo/gocraft-work-adapter"
import "github.com/gomodule/redigo/redis"

// ...

app = buffalo.New(buffalo.Options{
  // ...
  Worker: gwa.New(gwa.Options{
    Pool: &redis.Pool{
      MaxActive: 5,
      MaxIdle:   5,
      Wait:      true,
      Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", ":6379")
      },
    },
    Name:           "myapp",
    MaxConcurrency: 25,
  }),
  // ...
})
```

Por favor, nota que este paso es opcional, si deseas usar un ejecutor basado en goroutines.


### Registro de Handlers de ejecución

Handler es una función que se ejecutará para procesar trabajos para un tipo dado en la cola. Estos handlers deben registrarse primero con el worker que los ejecutará.

Cada handler tiene que implementar la siguiente interfaz:
```go
// Handler function that will be run by the worker and given
// a slice of arguments
type Handler func(worker.Args) error
```

Para adjuntar una función dada a un tipo de trabajo, vincúlalo a su ejecutor usando el método `Register`:

```go
import "github.com/gobuffalo/buffalo/worker"

var w worker.Worker

func init() {
  w = App().Worker // Get a ref to the previously defined Worker
  w.Register("send_email", func(args worker.Args) error {
    // do work to send an email
    return nil
  })
}
```

### Poner en cola un Job

Ahora que los handlers están vinculados, necesitaras enviar los trabajos a la cola. Se recomienda solo usar los tipos básicos cuando se pone en cola un trabajo. Por ejemplo, usar el ID de un modelo y no el modelo completo en sí.

Puedes optar por activar los trabajos ahora, o esperar un tiempo o una duración determinada.

#### `worker.Perform`

El método `Perform` pone en cola un trabajo, el worker intentará ejecutar el trabajo tan pronto como sea posible, en función de la implementación del worker en sí.

```go
func doWork() {
  // Send the send_email job to the queue, and process it as soon as possible.
  w.Perform(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  })
}
```

#### `worker.PerformIn`

El método `PerformIn` pone en cola el trabajo, el worker intentará ejecutar el trabajo después que haya transcurrido una duración, en función de la implementación del worker en sí.

```go
func doWork() {
  // Send the send_email job to the queue, and process it in 5 seconds.
  // Please note if no working unit is free at this time, it will wait for a free slot.
  w.PerformIn(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, 5 * time.Second)
}
```

#### `worker.PerformAt`

El método `PerformAt` pone en cola un trabajo, el worker intentará ejecutar un trabajo en (o cerca de) el tiempo especificado, en función de la implementación del worker en sí.

```go
func doWork() {
  // Send the send_email job to the queue, and process it at now + 5 seconds.
  // Please note if no working unit is free at this time, it will wait for a free slot.
  w.PerformAt(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, time.Now().Add(5 * time.Second))
}
```

### Iniciar and detener Workers

Por defecto, todas las aplicaciones de Buffalo tendrán un archivo `main.go` como este:

```go
// cmd/app/main.go

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
```

El método [`buffalo.App#Serve`](https://godoc.org/github.com/gobuffalo/buffalo#App.Serve), por defecto, invocará el método [`worker.Worker#Start`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) para el worker registrado. Esto también llamará al método [`worker.Worker#Stop`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) cuando la aplicación se detenga. Este es el enfoque **recomendado** para las aplicaciones.

Si no deseas que tus Workers inicien automáticamente, puedes establecer la opción [`buffalo.Options#WorkerOff`](https://godoc.org/github.com/gobuffalo/buffalo#Options) a `true` cuando configuras tu aplicación.
