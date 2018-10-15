<%= title("Known Events") %>

### Application Events

The following events are known to be emitted by Buffalo during the application lifecyle.

| Constant                    | String                       | Emitted When                                                                               | Payload                                                                                                                                                                           |
| ---                         | ---                          | ---                                                                                        | ---                                                                                                                                                                               |
| `buffalo.EvtAppStart`       | `"buffalo:app:start"`        | <%= doclink("github.com/gobuffalo/buffalo#App.Serve") %> is called                         | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtAppStartErr`    | `"buffalo:app:start:err"`    | an error occurs calling <%= doclink("github.com/gobuffalo/buffalo#App.Serve") %>           | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtAppStop`        | `"buffalo:app:stop"`         | <%= doclink("github.com/gobuffalo/buffalo#App.Stop") %> is called                          | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtAppStopErr`     | `"buffalo:app:stop:err"`     | an error occurs calling <%= doclink("github.com/gobuffalo/buffalo#App.Stop") %>            | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtRouteStarted`   | `"buffalo:route:started"`    | a requested route is being processed                                                       | `route`: <%= doclink("github.com/gobuffalo/buffalo#RouteInfo") %><br>`app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %><br>`context`: <%= doclink("github.com/gobuffalo/buffalo#Context") %> |
| `buffalo.EvtRouteFinished`  | `"buffalo:route:finished"`   | a requested route is completed                                                             | `route`: <%= doclink("github.com/gobuffalo/buffalo#RouteInfo") %><br>`app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %><br>`context`: <%= doclink("github.com/gobuffalo/buffalo#Context") %> |
| `buffalo.EvtRouteErr`       | `"buffalo:route:err"`        | there is a problem handling processing a route                                             | `route`: <%= doclink("github.com/gobuffalo/buffalo#RouteInfo") %><br>`app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %><br>`context`: <%= doclink("github.com/gobuffalo/buffalo#Context") %> |
| `buffalo.EvtWorkerStart`    | `"buffalo:worker:start"`     | <%= doclink("github.com/gobuffalo/buffalo#App.Serve") %> is called and workers are started | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtWorkerStartErr` | `"buffalo:worker:start:err"` | an error occurs when starting workers                                                      | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtWorkerStop`     | `"buffalo:worker:stop"`      | <%= doclink("github.com/gobuffalo/buffalo#App.Stop") %> is called and workers are stopped  | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtWorkerStopErr`  | `"buffalo:worker:stop:err"`  | an error occurs when stopping workers                                                      | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %>                                                                                                                        |
| `buffalo.EvtFailureErr`     | `"buffalo:failure:err"`      | something can't be processed at all. it is a bad thing                                     | `app`: <%= doclink("*github.com/gobuffalo/buffalo#App") %><br>`context`: <%= doclink("github.com/gobuffalo/buffalo#Context") %>                                                   |


### Buffalo Dev Events

The following events are known to be emitted by the `buffalo dev` during the development lifecyle.

| String                         | Emitted When                   | Payload                                                                                                                                            |
| ---                            | ---                            | ---                                                                                                                                                |
| `"buffalo:dev:raw"`            | an applicable file is modified | `event`: <%= doclink("github.com/fsnotify/fsnotify#Event") %>                                                                                      |
| `"buffalo:dev:build:started"`  | a build has started            | `event`: <%= doclink("github.com/fsnotify/fsnotify#Event") %><br>`cmd`: string of the `go build` command (example: `"go build foo"`)               |
| `"buffalo:dev:build:finished"` | a build has completed          | `event`: <%= doclink("github.com/fsnotify/fsnotify#Event") %><br>`pid`: PID of the newly running binary<br>`build_time`: the duration of the build |
| `"buffalo:dev:build:err"`      | a build error has occurred     | `event`: <%= doclink("github.com/fsnotify/fsnotify#Event") %><br>`cmd`: string of the `go build` command (example: `"go build foo"`)               |
