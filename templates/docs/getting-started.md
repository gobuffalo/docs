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
          Buffalo isn't just a framework, it's a hollistic web development environment and project structure that let's developers get straight to the business of, well, building their business.
        </p>


      </div>
    </td>
  </tr>
</table>


{{#panel title="Installation"}}

```
$ go get -u github.com/gobuffalo/buffalo/buffalo
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
Buffalo version 0.5.0

--> ../coke/main.go
--> ../coke/Procfile
--> ../coke/Procfile.development
--> ../coke/.buffalo.dev.yml
--> ../coke/actions/app.go
--> ../coke/actions/home.go
--> ../coke/actions/home_test.go
--> ../coke/actions/render.go
--> ../coke/grifts/routes.go
--> ../coke/templates/index.html
--> ../coke/templates/application.html
--> ../coke/public/images/logo.svg
--> ../coke/.gitignore
--> go get github.com/markbates/refresh/...
--> go install github.com/markbates/refresh
--> go get github.com/markbates/grift/...
--> go install github.com/markbates/grift
--> go get github.com/motemen/gore
--> go install github.com/motemen/gore
--> ../coke/public/assets/application.js
--> ../coke/public/assets/application.css
--> ../coke/webpack.config.js
--> ../coke/assets/js/application.js
--> ../coke/assets/css/application.scss
--> npm install webpack -g
--> npm init -y
--> npm install --save webpack sass-loader css-loader style-loader node-sass babel-loader extract-text-webpack-plugin babel babel-core url-loader file-loader jquery bootstrap path font-awesome npm-install-webpack-plugin jquery-ujs
--> ../coke/models/models.go
--> go get github.com/markbates/pop/...
--> go install github.com/markbates/pop/soda
--> database.yml
--> go get -t ./...
--> goimports -w .
```

To see a list of available flags for the `new` command, just check out its help.

```
$ buffalo help new
```

Note: by default, Buffalo generates a database.yml targeted for postgres. If you wish to change this behavior, you can pass in a `--db-type` flag into the `new` command.

```
$ buffalo new coke --db-type sqlite3
```

{{/panel}}

{{ partial "docs/dev.md" }}

{{#panel title="Building Your Application" name="building"}}

Buffalo features a command, `build`, that will build a full binary of your application including, but not limited to; assets, migrations, templates, etc... If you buy into the "Buffalo Way" things just work. It's a wonderful experience. :)

```
$ buffalo build
```

```text
--> running webpack
Hash: 47a7dc2dd9d5da7eb169
Version: webpack 1.14.0
Time: 2961ms
                                 Asset     Size  Chunks             Chunk Names
                             .DS_Store  6.15 kB          [emitted]
  f4769f9bdb7466be65088239c12046d1.eot  20.1 kB          [emitted]
 fa2772327f55d8198301fdb8bcfc8158.woff  23.4 kB          [emitted]
  e18bbf611f2a2e43afc071aa2f4e1512.ttf  45.4 kB          [emitted]
  89889688147bd7575d6327160d64e760.svg   109 kB          [emitted]
  674f50d287a8c48dc19ba404d20fe713.eot   166 kB          [emitted]
af7ae505a9eed503f8b8e6982036873e.woff2  77.2 kB          [emitted]
 fee66e712a8a08eef5805a46892932ad.woff    98 kB          [emitted]
  b06871f281fee6b241d60582ae9369b9.ttf   166 kB          [emitted]
  912ec66d7572ff821749319396470bde.svg   444 kB          [emitted]
                        application.js   372 kB       0  [emitted]  main
                       application.css   335 kB       0  [emitted]  main
448c34a56d699c29117adc64c43affeb.woff2    18 kB          [emitted]
        fonts/fontawesome-webfont.woff    98 kB          [emitted]
         fonts/fontawesome-webfont.eot   166 kB          [emitted]
       fonts/fontawesome-webfont.woff2  77.2 kB          [emitted]
         fonts/fontawesome-webfont.ttf   166 kB          [emitted]
                 fonts/FontAwesome.otf   135 kB          [emitted]
                   images/logo_med.png   157 kB          [emitted]
         fonts/fontawesome-webfont.svg   444 kB          [emitted]
                  images/uncle_sam.jpg  20.5 kB          [emitted]
                       images/logo.svg    66 kB          [emitted]
                       images/logo.png   366 kB          [emitted]
                       images/repl.gif  7.15 MB          [emitted]
   [0] multi main 52 bytes {0} [built]
    + 23 hidden modules
Child extract-text-webpack-plugin:
        + 17 hidden modules
--> built rice box ../github.com/gobuffalo/gobuffalo/rice-box.go
--> built rice box ../github.com/gobuffalo/gobuffalo/actions/rice-box.go
--> running go build -v -o bin/gobuffalo -ldflags -X main.version=db92753 -X main.buildTime="2017-01-03T11:20:30-05:00"
--> cleaning up build
```

If you deploying to a machine that has `zip` installed, it is recommended that you use the `-z` flag with `buffalo build`, it will result in faster build times, and smaller binaries.

{{/panel}}

