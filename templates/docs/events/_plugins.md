## Listening with Plugins

To enable a plugin to a receive a JSON version of emitted events, the plugin can set the <%= doclink("github.com/gobuffalo/buffalo-plugins/plugins#Command.BuffaloCommand") %> value to `events` when listing the `available` commands for the plugin.

<%= codeTabs() { %>
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
<% } %>
