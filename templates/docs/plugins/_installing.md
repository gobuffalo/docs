<%= title("Installing Plugins") %>

<%= sinceVersion("1.1.0", {pkg: "github.com/gobuffalo/buffalo-plugins"}) %>

To add support for the plugin manager, one can either manually edit `./config/buffalo-plugins.toml` or let `buffalo plugins install` create it for you.

<%= codeTabs() { %>
```bash
// $ buffalo plugins install

go get github.com/gobuffalo/buffalo-pop
./config/buffalo-plugins.toml
```

``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"
```

```bash
// $ buffalo plugins list

Bin         |Command               |Description
---         |---                   |---
buffalo-pop |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop |buffalo destroy model |Destroys model files.
buffalo-pop |buffalo pop           |A tasty treat for all your database needs
```
<% } %>

The `buffalo-pop` plugin was automatically added because the application in this example is a Buffalo application that uses Pop.

New plugins can be install in bulk with the `install` command

<%= codeTabs() { %>
```bash
// $ buffalo plugins install
$ buffalo plugins install github.com/markbates/buffalo-trash github.com/gobuffalo/buffalo-heroku

go get github.com/gobuffalo/buffalo-heroku
go get github.com/gobuffalo/buffalo-pop
go get github.com/markbates/buffalo-trash
./config/buffalo-plugins.toml
```

``` bash
// ./config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-heroku"
  go_get = "github.com/gobuffalo/buffalo-heroku"

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"

[[plugin]]
  binary = "buffalo-trash"
  go_get = "github.com/markbates/buffalo-trash"
```

```bash
// $ buffalo plugins list

Bin            |Command               |Description
---            |---                   |---
buffalo-pop    |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-heroku |buffalo heroku        |helps with heroku setup and deployment for buffalo applications
buffalo-pop    |buffalo destroy model |Destroys model files.
buffalo-pop    |buffalo pop           |A tasty treat for all your database needs
buffalo-trash  |buffalo trash         |destroys and recreates a buffalo app
```

<% } %>
