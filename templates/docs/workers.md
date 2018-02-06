# Background Job Workers

When building complex applications it is often nice to be able to run things in the “background”. While Go provides beautiful concurrency features, like the famed Goroutine, often one wants to run these on different machines, persist them using Redis, or any number of potential reasons why a simple Goroutine isn’t sufficient.

<%= title("The Worker Interface") %>

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

Current there are two implementations of this interface. The first is `worker.Simple`. This uses Goroutines to implement the interface. This is great for simple applications, but since the queues are not persisted any jobs enqueued will be lost if the server was to be shut down. This implementation is turned on by default.

The other implementation is the [`github.com/gobuffalo/gocraft-work-adapter`](https://github.com/gobuffalo/gocraft-work-adapter) package which implements the [`github.com/gocraft/work`](https://github.com/gocraft/work) package using Redis as the backing store.

<%= title("Setting Up a Worker Adapter") %>

When setting up your application you can assign a worker implementation to the `Worker` option.

```go
import "github.com/gobuffalo/gocraft-work-adapter"
import "github.com/garyburd/redigo/redis"

buffalo.New(buffalo.Options{
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

<%= title("The Job type") %>

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

<%= title("Registering a Worker Handler") %>

Handler are functions that will be run to process jobs in the queue. These handlers have to be first registered with the worker that will be running them.

```go
// Handler function that will be run by the worker and given
// a slice of arguments
type Handler func(Args) error
```

```go
import "github.com/gobuffalo/buffalo/worker"

var w worker.Worker

func init() {
  w = worker.NewSimple()
  w.Register("send_email", func(args worker.Args) error {
    // do work to send an email
    return nil
  })
}
```

<%= title("Enqueueing a Job") %>

It is recommended to only use basic types when enqueueing a job. For example, use the ID of a model, and not the whole model itself.

### `worker.Perform`

The `Perform` method should enqueue the job and the worker should try and run the job as soon as possible, based on the implementation of the worker itself.

```go
func doWork() {
  w.Perform(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  })
}
```

### `worker.PerformIn`

The `PerformIn` method should enqueue the job and the worker should try and run the job after the duration has passed, based on the implementation of the worker itself.

```go
func doWork() {
  w.PerformIn(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, 5 * time.Second)
}
```

### `worker.PerformAt`

The `PerformIn` method should enqueue the job and the worker should try and run the job at (or near) the time specified, based on the implementation of the worker itself.

```go
func doWork() {
  w.PerformIn(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, time.Now().Add(5 * time.Second))
}
```

<%= title("Starting and Stopping Workers") %>

By default all Buffalo applications created will have a `main.go` that looks something like this:

```go
func main() {
  app := actions.App()
  if err := app.Serve(); err != nil {
    log.Fatal(err)
  }
}
```

The [`buffalo.App#Serve`](https://godoc.org/github.com/gobuffalo/buffalo#App.Serve) method will, by default, call the [`worker.Worker#Start`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method for the registered worker. This will also call the [`worker.Worker#Stop`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method when the application is shut down. This is the **recommended** approach for applications.

If you don't want your workers to start automatically, you can set the option [`buffalo.Options#WorkerOff`](https://godoc.org/github.com/gobuffalo/buffalo#Options) to `true` when setting up your application.
