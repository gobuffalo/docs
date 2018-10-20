<%= title("Removing Plugins") %>

<%= sinceVersion("1.1.0", {pkg: "github.com/gobuffalo/buffalo-plugins"}) %>

Plugins can be removed with the `remove` command. This only removes them from the config file, not from the users system.

<%= codeTabs() { %>
```bash
// $ buffalo plugins remove
$ buffalo plugins remove github.com/gobuffalo/buffalo-heroku

./config/buffalo-plugins.toml
```

``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"

[[plugin]]
  binary = "buffalo-trash"
  go_get = "github.com/markbates/buffalo-trash"
```

```bash
// $ buffalo plugins list

Bin           |Command               |Description
---           |---                   |---
buffalo-pop   |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop   |buffalo destroy model |Destroys model files.
buffalo-pop   |buffalo pop           |A tasty treat for all your database needs
buffalo-trash |buffalo trash         |destroys and recreates a buffalo app
```
<% } %>
