---
name: Generating a New Project
seoDescription: "Generate a new Buffalo project"
seoKeywords: 
  - "buffalo"
  - "go"
  - "golang"
  - "new project"
  - "generator"
  - "framework"
  - "web"
weight: 2
aliases:
  - /docs/getting-started/new-project
  - /en/docs/getting-started/new-project
---

# Generating a New Project

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

```bash
$ buffalo new coke
Buffalo version {{< latestclirelease >}}

      create  .buffalo.dev.yml
      create  assets/images/logo.svg
      create  assets/css/application.scss
      create  assets/images/favicon.ico
      create  assets/js/application.js
      create  .babelrc
      create  package.json
      create  public/assets/.keep
      create  webpack.config.js
         run  yarn install --no-progress --save
yarn install v0.27.5
info No lockfile found.
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...
success Saved lockfile.
Done in 11.71s.
      create  models/models.go
      create  models/models_test.go
      create  grifts/db.go
         run  go get github.com/gobuffalo/pop/...
      create  ./database.yml
         run  goimports -w coke/grifts/db.go coke/models/models.go coke/models/models_test.go
      create  Dockerfile
      create  .dockerignore
         run  go get -u golang.org/x/tools/cmd/goimports
      create  README.md
      create  actions/actions_test.go
      create  actions/app.go
      create  actions/home.go
      create  actions/home_test.go
      create  actions/render.go
      create  .codeclimate.yml
      create  .env
      create  grifts/init.go
      create  inflections.json
      create  locales/all.en-us.yaml
      create  main.go
      create  public/robots.txt
      create  templates/_flash.html
      create  templates/application.html
      create  templates/index.html
         run  go get -t ./...
         run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go
      create  .gitignore
         run  git init
Initialized empty Git repository in /Users/markbates/Dropbox/development/gocode/src/github.com/markbates/coke/.git/
         run  git add .
         run  git commit -q -m Initial Commit
INFO[0055] Congratulations! Your application, coke, has been successfully built!

 
INFO[0055] You can find your new application at:
/Users/markbates/Dropbox/development/gocode/src/github.com/markbates/coke 
INFO[0055] 
Please read the README.md file in your new application for next steps on running your application.
```


## Create a customized app

The default setup is great, but maybe it doesn't fit you. Buffalo provides several options as flags for the `new` command.

You can get the available flags list using the `help` command: 

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci, circleci] (default "none")
      --config string        config file (default is $HOME/.buffalo.yaml)
      --db-type string       specify the type of database you want to use [cockroach, mariadb, mysql, postgres] (default "postgres")
      --docker string        specify the type of Docker file to generate [none, multi, standard] (default "multi")
  -d, --dry-run              dry run
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --module string        specify the root module (package) name. [defaults to 'automatic']
      --skip-config          skips using the config file
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
```


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

{{< note >}}
Before starting Buffalo for the first time, please head over to the [Database](/documentation/database/pop) docs and read a little bit about setting up your databases.
{{< /note >}}

One of the downsides to Go development is the lack of code “reloading”. This means as you change your code **you need to manually stop** your application, rebuild it, then restart it. Buffalo finds this annoying, and wants to make life better for you.

```bash
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files and the [asset](/documentation/frontend-layer/assets) folder by default. It will **rebuild and restart your binary for you** automatically, so you don't have to worry about such things.

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

You can also take a look at the [Env Variables](/documentation/getting_started/configuration) chapter for further information on Buffalo configuration.

## Next Steps

* [Directory Structure](/documentation/getting_started/directory-structure) - Learn more about Buffalo structure.
* [Configuration](/documentation/getting_started/configuration) - Manage your app configuration.
