# Getting Started

<table>
  <tr>
    <td width="30%">
      <img src="/assets/images/logo.svg">
    </td>
    <td>
      <div class="jumbotron">
        <p>
        Buffalo is a Go web framework. Yeah, I hate the word "framework" too! Buffalo is different though. Buffalo doesn't want to re-invent wheels like routing and templating. Buffalo is glue that wraps all of the best packages available and makes them all play nicely together.
        </p>

        <p>
        Buffalo is "idiomatic", for whatever that is worth. The purpose of a framework is not to bend you to it's will, but to rather get out of your way and make your job of building your application easy. That is the goal of Buffalo.
        </p>

        <p>
        If you were to look through the Buffalo code base you'll find little code, just enough to assemble the amazing packages that other's have written into one coherent system.
        </p>

      </div>
    </td>
  </tr>
</table>


{{#panel title="Installation"}}

```
$ go get -u github.com/markbates/buffalo/buffalo
```
{{/panel}}

{{#panel title="Generating a New Project" name="new-project"}}

Buffalo aims to make building new web applications in Go as simple as possible, and what could be more simple that a new application generator?

```
$ buffalo new <name>
```

That will generate a whole new Buffalo application that is ready to go. It'll even run `go get` for you to make sure you have all of the necessary dependencies needed to run your application.

```text
$ buffalo new coke
Buffalo version 0.4.6
--> .../coke/main.go
--> .../coke/Procfile
--> .../coke/Procfile.development
--> .../coke/.buffalo.dev.yml
--> .../coke/actions/app.go
--> .../coke/actions/home.go
--> .../coke/actions/home_test.go
--> .../coke/actions/render.go
--> .../coke/grifts/routes.go
--> .../coke/templates/index.html
--> .../coke/templates/application.html
--> .../coke/assets/js/application.js
--> .../coke/assets/css/application.css
--> .../coke/.gitignore
--> go get github.com/markbates/refresh/...
--> go install github.com/markbates/refresh
--> go get github.com/markbates/grift/...
--> go install github.com/markbates/grift
--> .../coke/assets/js/jquery.js
--> .../coke/assets/js/jquery.map
--> .../coke/assets/css/bootstrap.css
--> .../coke/assets/js/bootstrap.js
--> .../coke/assets/js/jquery.js
--> .../coke/assets/js/jquery.map
--> .../coke/models/models.go
--> go get github.com/markbates/pop/...
--> go install github.com/markbates/pop/soda
--> database.yml
--> go get -t ./...
--> goimports -w .
```

To see a list of available flags for the `new` command, just check out it's help.

```
$ buffalo help new
```

{{/panel}}

{{#panel title="Running Your Application" name="running"}}

Buffalo is Go "standards" compliant, that means you can just build your binary and run it. It's that simple.

{{/panel}}

{{#panel title="Running Your Application in Development" name="running-in-dev"}}

One of the downsides to Go development is the lack of code "reloading". This means as you change your code you need to manually stop your application, rebuild it, and then restart it. Buffalo finds this is annoying, and wants to make life better for you.

```
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files by default and rebuild, and restart, your binary for you so you don't have to worry about such things. Just run the `dev` command and start coding.

{{/panel}}
