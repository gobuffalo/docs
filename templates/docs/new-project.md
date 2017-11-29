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
* the **Buffalo framework layout** and default configuration ([pop/soda](https://github.com/markbates/pop) with PostgreSQL support),
* all necessary **Go dependencies** needed to run the current application,
* **frontend dependencies** and working setup with [webpack](https://webpack.js.org/)
* and an initial **Git repository**.

<%= partial("docs/new-project/new.md") %>

<%= title("Create a customized app") %>

The default setup is great, but maybe it doesn't fit you. Buffalo provides several options as flags for the `new` command.

You can get the available flags list using the `help` command: 

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci] (default "none")
      --db-type string       specify the type of database you want to use [postgres, mysql, sqlite3] (default "postgres")
      --docker string        specify the type of Docker file to generate [none, multi, standard] (default "multi")
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
      --with-dep             adds github.com/golang/dep to your app
```

You can choose to generate an API application, skipping the frontend stuff. Maybe you want to setup a CI to build your app on your favourite system? Or even use your own package to handle the database? Just use the flags!

<%= partial("docs/dev.md") %>
