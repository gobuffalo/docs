<% seoDescription("Generate a new Buffalo project") %>
<% seoKeywords(["buffalo", "go", "golang", "new project", "generator", "framework", "web"]) %>

<%= h1("Generating a New Project") %>

You now have a working Buffalo installation. In this section, you will learn how to create **a brand new web application**, using the `buffalo` command. 

<%= title("Create a new project") %>

Buffalo aims to make building new web applications in Go as **quick and simple** as possible. What could be more simple than a *new application* generator?

Start by going to your `$GOPATH` and create your new application!

```bash
$ cd $GOPATH/src/github.com/$USER/
```

Make sure `$GOPATH/bin` is in your `$PATH`, then:

```bash
$ buffalo new coke
```

That will generate a whole new Buffalo application called **coke**, all ready to go:
* the **Buffalo framework layout** and default configuration ([pop/soda](https://github.com/gobuffalo/pop) with PostgreSQL support),
* all necessary **Go dependencies** needed to run the current application,
* **frontend dependencies** and working setup with [webpack](https://webpack.js.org/)
* and an initial **Git repository**.

<%= partial("docs/new-project/new.md") %>

<%= title("Create a customized app") %>

The default setup is great, but maybe it doesn't fit you. Buffalo provides several options as flags for the `new` command.

You can get the available flags list using the `help` command: 

<%= partial("docs/new-project/help.md") %>

You can choose to generate an API application, skipping the frontend stuff. Maybe you want to setup a CI to build your app on your favourite system? Or even use your own package to handle the database? Just use the flags!

<%= title("Override Default Config") %>

By default `buffalo new` command will look for a configuration file at `$HOME/.buffalo.yml` and if it exists will try to load it. You can override the flags found in that file by passing the right ones in the command line or use the `--config` flag to specify a different YAML file. If the `--skip-config` flag is used `buffalo new` command will not load any config file and will use only the flags passed by the command line.

An example of a `.buffalo.yml` config file can be:

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

<%= partial("docs/dev.md") %>

<%= title("Next Steps") %>

* [Directory Structure](/en/docs/directory-structure) - Learn more about Buffalo structure.
* [Configuration](/en/docs/config-vars) - Manage your app configuration.