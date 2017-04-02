# Getting Started

<%= title("Installation", {}) %>

<%= code("text") { %>
$ go get -u github.com/gobuffalo/buffalo/...
<% } %>

<%= title("Generating a New Project", {name: "new-project"}) %>

Buffalo aims to make building new web applications in Go as quick and simple as possible, and what could be more simple than a *new application* generator? Start by going to your `$GOPATH` and create your new application!

<%= code("text") { %>
$ cd $GOPATH/src/github.com/$USER/
$ # Make sure $GOPATH/bin is in your $PATH, then:
$ buffalo new <name>
<% } %>

That will generate a whole new Buffalo application that is ready to go. It'll even run `go get` for you to make sure you have all of the necessary dependencies needed to run your application.

<%= code("text") { %>
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
<% } %>

To see a list of available flags for the `new` command, just check out its help.

<%= code("text") { %>
$ buffalo help new
<% } %>

Note: by default, Buffalo generates a database.yml targeted for postgres. If you wish to change this behavior, you can pass in a `--db-type` flag into the `new` command.

<%= code("text") { %>
$ buffalo new coke --db-type sqlite3
<% } %>

If your app doesn't need a database, or if you want to handle it by yourself, you can use the `--skip-pop` flag.

<%= code("text") { %>
$ buffalo new coke --skip-pop
<% } %>

<%= partial("docs/dev.md") %>

<%= title("Building Your Application", {name: "building"}) %>

Buffalo features a command, `build`, that will build a full binary of your application including, but not limited to; assets, migrations, templates, etc... If you buy into the "Buffalo Way" things just work. It's a wonderful experience. :)

<%= code("text") { %>
$ buffalo build
<% } %>

<%= code("text") { %>
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

<% } %>

See [Building](/docs/building) for more options on the `build` command.