---
name: Events
seoDescription: Listening for events in a Buffalo application
seoKeywords: ["buffalo", "go", "golang", "events", "plugins"]
aliases:
  - /docs/events
  - /pt/docs/events
---

# Events

{{< since "0.13.0-beta.2" >}}

The {{< doclink href="github.com/gobuffalo/events" message="events" >}} package allows for Go applications, including Buffalo applications, to listen, and emit, global event messages.


## Listening for Events

To start listening for events a {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}} must first be registered with the {{< doclink message="events" href="github.com/gobuffalo/events" >}} package.

```go
func init() {
  _, err := events.Listen(func(e events.Event) {
    // do work
  })
}
```

Once registered this new listener function will be sent all events emitted through the {{< doclink message="events" href="github.com/gobuffalo/events" >}} package.
## Emitting Events

When emitting events the `Kind` attribute should be a unique, but constant, string. It is this attribute that users will use to determine how to respond to events they receive.

It is recommended to namespace this attribute like such, with error events being suffixed with `:err`.

```plain
"<package-name>:<additional-names>:<optional-error>"
"myapp:foo:start"
"myapp:foo:stop"
"mypkg:workers:bar:err"
```

This naming pattern makes it easier for users to filter events to only those that they care about. See [Filtering Events](#filtering-events) for more details.

---

There are multiple ways to emit an {{< doclink message="events#Event" href="github.com/gobuffalo/events#Event" >}} in your Go code. The {{< doclink message="events#EmitError" href="github.com/gobuffalo/events#EmitError" >}} and {{< doclink message="events#EmitPayload" href="github.com/gobuffalo/events#EmitPayload" >}} functions both accept a `payload interface{}` argument. It is recommended to use {{< doclink message="events#Payload" href="github.com/gobuffalo/events#Payload" >}} for payloads; any other type passed in will get converted into a {{< doclink message="events#Payload" href="github.com/gobuffalo/events#Payload" >}} with the argument set in the payload with the key, `data`.

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

## Filtering Events

In the [Emitting Events](#emitting-events) section the naming convention for {{< doclink message="events#Event.Kind" href="github.com/gobuffalo/events#Event.Kind" >}} is described. By the checking the value of {{< doclink message="events#Event.Kind" href="github.com/gobuffalo/events#Event.Kind" >}}.

{{< codetabs >}}
{{< tab "direct match" >}}
```go
// direct match
events.Listen(func(e events.Event) {
  if e.Kind != buffalo.EvtRouteStarted {
    // do nothing
    return
  }
  // do work on the route event
})
```
{{< /tab >}}

{{< tab "matching with a switch statement" >}}
```go
// matching with a switch statement
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

{{< tab "matching error events" >}}
```go
// matching error events
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

{{< tab "matching on prefix" >}}
```go
// matching on prefix
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

## Stop Listening for Events

When registering a new {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}} a {{< doclink message="events#DeleteFn" href="github.com/gobuffalo/events#DeleteFn" >}} is returned. This function should be held on to and used when you want to remove the added listener.

```go
deleteFn, err := events.Listen(func(e events.Event) {
  // do work
})
if err != nil {
  return err
}
defer deleteFn()
```
## Listening with Plugins

To enable a plugin to a receive a JSON version of emitted events, the plugin can set the {{< doclink message="events#Command.BuffaloCommand" href="github.com/gobuffalo/buffalo-plugins/plugins#Command.BuffaloCommand" >}} value to `events` when listing the `available` commands for the plugin.

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

## Integrating a Messaging Queue

It is often desirable to take events emitted and send them to a message queue, such as Kafka or Redis, to be processed externally. The {{< doclink message="events" href="github.com/gobuffalo/events" >}} package does not have a directhook for this sort of functionality, the most direct way of enabling this behavior is to register a {{< doclink message="events#Listener" href="github.com/gobuffalo/events#Listener" >}} that can then hand the event over to the appropriate message queue.

```go
events.Listen(func(e events.Event) {
  myMessageQ.DoWork(e)
})
```
## Known Events

### Application Events

The following events are known to be emitted by Buffalo during the application lifecyle.

| Constant                    | String                       | Emitted When                                                                               | Payload                                                                                                                                                                           |
| ---                         | ---                          | ---                                                                                        | ---                                                                                                                                                                               |
| `buffalo.EvtAppStart`       | `"buffalo:app:start"`        | {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}} is called                         | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtAppStartErr`    | `"buffalo:app:start:err"`    | an error occurs calling {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}}           | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtAppStop`        | `"buffalo:app:stop"`         | {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}} is called                          | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtAppStopErr`     | `"buffalo:app:stop:err"`     | an error occurs calling {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}}            | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtRouteStarted`   | `"buffalo:route:started"`    | a requested route is being processed                                                       | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtRouteFinished`  | `"buffalo:route:finished"`   | a requested route is completed                                                             | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtRouteErr`       | `"buffalo:route:err"`        | there is a problem handling processing a route                                             | `route`: {{< doclink message="buffalo#RouteInfo" href="github.com/gobuffalo/buffalo#RouteInfo" >}}<br> `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}} |
| `buffalo.EvtWorkerStart`    | `"buffalo:worker:start"`     | {{< doclink message="buffalo#App.Serve" href="github.com/gobuffalo/buffalo#App.Serve" >}} is called and workers are started | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtWorkerStartErr` | `"buffalo:worker:start:err"` | an error occurs when starting workers                                                      | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtWorkerStop`     | `"buffalo:worker:stop"`      | {{< doclink message="buffalo#App.Stop" href="github.com/gobuffalo/buffalo#App.Stop" >}} is called and workers are stopped  | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtWorkerStopErr`  | `"buffalo:worker:stop:err"`  | an error occurs when stopping workers                                                      | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}                                                                                                                        |
| `buffalo.EvtFailureErr`     | `"buffalo:failure:err"`      | something can't be processed at all. it is a bad thing                                     | `app`: {{< doclink message="buffalo#App" href="*github.com/gobuffalo/buffalo#App" >}}<br> `context`: {{< doclink message="buffalo#Context" href="github.com/gobuffalo/buffalo#Context" >}}                                                   |


### Buffalo Dev Events

The following events are known to be emitted by the `buffalo dev` during the development lifecyle.

| String                         | Emitted When                   | Payload                                                                                                                                            |
| ---                            | ---                            | ---                                                                                                                                                |
| `"buffalo:dev:raw"`            | an applicable file is modified | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}}                                                                                      |
| `"buffalo:dev:build:started"`  | a build has started            | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}}<br> `cmd`: string of the `go build` command (example: `"go build foo"`)               |
| `"buffalo:dev:build:finished"` | a build has completed          | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}}<br> `pid`: PID of the newly running binary<br> `build_time`: the duration of the build |
| `"buffalo:dev:build:err"`      | a build error has occurred     | `event`: {{< doclink message="fsnotify#Event" href="github.com/fsnotify/fsnotify#Event" >}}<br> `cmd`: string of the `go build` command (example: `"go build foo"`)               |
