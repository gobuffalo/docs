# Getting Started

<table>
  <tr>
    <td width="30%">
      <img src="/assets/images/logo.svg">
    </td>
    <td>
      <div class="jumbotron">
        <p>
        Buffalo is a Go web development eco-system. Designed to make the life of a Go web developer easier.
        </p>

        <p>
        Buffalo starts by generating a web project for you that already has everything from front-end (JavaScript, SCSS, etc...) to back-end (Database, Routing, etc...) already hooked up and ready to run. From there it provides easy APIs to build your web application quickly in Go.
        </p>

        <p>
          Buffalo isn't just a framework, it's a holistic web development environment and project structure that let's developers get straight to the business of, well, building their business.
        </p>


      </div>
    </td>
  </tr>
</table>

<%= panel("Installation", {}) { %>

```
$ go get -u github.com/gobuffalo/buffalo/...
```
<% } %>

<%= panel("Generating a New Project", {name: "new-project"}) { %>

Buffalo aims to make building new web applications in Go as quick and simple as possible, and what could be more simple than a *new application* generator? Start by going to your `$GOPATH` and create your new application!

```
$ cd $GOPATH/src/github.com/$USER/
$ buffalo new <name>
```

That will generate a whole new Buffalo application that is ready to go. It'll even run `go get` for you to make sure you have all of the necessary dependencies needed to run your application.

```text
$ buffalo new coke
Buffalo version <%= version %>

--> .../coke/README.md
--> .../coke/actions/actions_test.go
--> .../coke/actions/app.go
--> .../coke/actions/home.go
--> .../coke/actions/home_test.go
--> .../coke/actions/render.go
--> .../coke/.codeclimate.yml
--> .../coke/.gitignore
--> .../coke/grifts/routes.go
--> .../coke/main.go
--> .../coke/templates/application.html
--> .../coke/templates/index.html
--> .../coke/.buffalo.dev.yml
--> go get github.com/markbates/refresh/...
--> go install github.com/markbates/refresh
--> go get github.com/markbates/grift/...
--> go install github.com/markbates/grift
--> go get github.com/motemen/gore
--> go install github.com/motemen/gore
--> .../coke/assets/images/logo.svg
--> .../coke/assets/css/application.scss
--> .../coke/assets/js/application.js
--> .../coke/public/assets/.gitignore
--> .../coke/webpack.config.js
--> npm init -y
--> .../coke/models/models.go
--> .../coke/models/models_test.go
--> go get github.com/markbates/pop/...
--> go install github.com/markbates/pop/soda
--> database.yml
--> go get -t ./...
--> goimports -w .
Congratulations! Your application, coke, has been successfully built!

You can find your new application at:
.../coke

Please read the README.md file in your new application for next steps on running your application.
```

To see a list of available flags for the `new` command, just check out its help.

```
$ buffalo help new
```

Note: by default, Buffalo generates a database.yml targeted for postgres. If you wish to change this behavior, you can pass in a `--db-type` flag into the `new` command.

```
$ buffalo new coke --db-type sqlite3
```

<% } %>

<%= partial("docs/dev.md") %>

<%= panel("Building Your Application", {name: "building"}) { %>

Buffalo features a command, `build`, that will build a full binary of your application including, but not limited to; assets, migrations, templates, etc... If you buy into the "Buffalo Way" things just work. It's a wonderful experience. :)

```
$ buffalo build
```

```text
Buffalo version 0.8.0.dev

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/gobuffalo -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

<% } %>

