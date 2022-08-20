---
name: Background Job Workers
weight: 12
aliases:
  - /docs/workers
  - /pt/docs/workers
---
# Background Job Workers

When building complex applications it is often nice to be able to run things in the “background”. While Go provides beautiful concurrency features, like the famed Goroutine, often one wants to run these on different machines, persist them using Redis, or any number of potential reasons why a simple Goroutine isn’t sufficient.

{{<note>}}
Workers shouldn't be confused with [tasks](/documentation/guides/tasks): tasks are synchronous tools, whereas workers are intended to run asynchronously.
{{</note>}}

## The Worker Interface

In order to use background jobs, the `worker.Worker` interface must first be satisfied.

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

Currently there are two official implementations of this interface:
* The first is `worker.Simple`; It uses Goroutines to implement the interface. This is great for simple applications, but since the queues are not persisted, any jobs enqueued will be lost if the server was to be shut down. **This implementation is turned on by default**.
* The other implementation is the [`github.com/gobuffalo/gocraft-work-adapter`](https://github.com/gobuffalo/gocraft-work-adapter) package, which implements the [`github.com/gocraft/work`](https://github.com/gocraft/work) package using Redis as the backing store.

### Community implementations

The following Worker implementations are provided by Buffalo users (no official support):

| Name | Author | Description |
|:------|:--------|:-------------|
| [AMQP worker adapter](https://github.com/stanislas-m/amqp-work-adapter) | [@stanislas-m](https://github.com/stanislas-m) | A Worker implementation to use with AMQP-compatible brokers (such as [RabbitMQ](https://www.rabbitmq.com/)). |

## The Job type

A Job is a unit of work for a given Worker implementation.

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

## How to Use Background Tasks

To be able to use background tasks, you'll need to setup a worker adapter, register job handlers and trigger jobs.

### Setting Up a Worker Adapter

When setting up your application you *can* assign a worker implementation to the `Worker` option.

**app.go**

```go
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

Please note this step is optional, if you want to use the goroutines-based runner.

### Registering a Worker Handler

Handler is a function that will be run to process jobs for a given type in the queue. These handlers have to be first registered with the worker that will be running them.

Each handler has to implement the following interface:
```go
// Handler function that will be run by the worker and given
// a slice of arguments
type Handler func(worker.Args) error
```

To attach a given function to a job type, bind it to your runner using the `Register` function:
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

### Enqueueing a Job

Now that the worker handlers are bound, you'll need to send jobs to the queue. It is recommended to only use basic types when enqueueing a job. For example, use the ID of a model, and not the whole model itself.

You can choose to trigger jobs right now, or wait for a given time or duration.

#### `worker.Perform`

The `Perform` method enqueues the job, so the worker should try and run the job as soon as possible, based on the implementation of the worker itself.

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

The `PerformIn` method enqueues the job, so the worker should try and run the job after the duration has passed, based on the implementation of the worker itself.

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

The `PerformAt` method enqueues the job, so the worker should try and run the job at (or near) the time specified, based on the implementation of the worker itself.

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

### Starting and Stopping Workers

By default all Buffalo applications created will have a `main.go` that looks something like this:

```go
// cmd/app/main.go

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
```

The [`buffalo.App#Serve`](https://godoc.org/github.com/gobuffalo/buffalo#App.Serve) method will, by default, call the [`worker.Worker#Start`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method for the registered worker. This will also call the [`worker.Worker#Stop`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method when the application is shut down. This is the **recommended** approach for applications.

If you don't want your workers to start automatically, you can set the option [`buffalo.Options#WorkerOff`](https://godoc.org/github.com/gobuffalo/buffalo#Options) to `true` when setting up your application.
