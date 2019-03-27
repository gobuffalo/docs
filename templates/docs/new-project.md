<% seoDescription("Generate a new Buffalo project") %>
<% seoKeywords(["buffalo", "go", "golang", "new project", "generator", "framework", "web"]) %>

<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Generating a New Project"}) %>
<% } %>

You now have a working Buffalo installation. In this section, you will learn how to create **a brand new web application**, using the `buffalo` command. 

## Create a new project

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

## Create a customized app

The default setup is great, but maybe it doesn't fit you. Buffalo provides several options as flags for the `new` command.

You can get the available flags list using the `help` command: 

<%= partial("docs/new-project/help.md") %>

You can choose to generate an API application, skipping the frontend stuff. Maybe you want to setup a CI to build your app on your favorite system? Or even use your own package to handle the database? Just use the flags!

## Override Default Config

By default `buffalo new` command will look for a configuration file at `$HOME/.buffalo.yml` and if it exists will try to load it. You can override the flags found in that file by passing the right ones in the command line or use the `--config` flag to specify a different YAML file. If the `--skip-config` flag is used `buffalo new` command will not load any config file and will use only the flags passed by the command line.

An example of a `.buffalo.yml` config file can be:

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

## Running Your Application in Development

<%= note() { %>
Before starting Buffalo for the first time, please head over to the [Database](/docs/db) docs and read a little bit about setting up your databases.
<% } %>

One of the downsides to Go development is the lack of code “reloading”. This means as you change your code **you need to manually stop** your application, rebuild it, then restart it. Buffalo finds this is annoying, and wants to make life better for you.

```bash
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files and the [asset](/docs/assets) folder by default. It will **rebuild and restart your binary for you** automatically, so you don't have to worry about such things.

Just run the `buffalo dev` command and go to [localhost:3000/](http://localhost:3000/) to see all changes live!

<figure>
  <img src="/assets/images/new-coke.png" title="screenshot">
  <figcaption>The brand new Coke app.</figcaption>
</figure>

#### Run the dev server on a custom port

Sometimes you will already have an app working on the 3000 port. You can configure the dev server port by providing the `PORT` environment variable:

```bash
$ PORT=3001 buffalo dev
```

You can also take a look at the [Env Variables](/en/docs/config-vars) chapter for further information on Buffalo configuration.

## Next Steps

* [Directory Structure](/en/docs/directory-structure) - Learn more about Buffalo structure.
* [Configuration](/en/docs/config-vars) - Manage your app configuration.