---
name: Generando Nuevo Proyecto
seoDescription: "Generando un Nuevo Proyecto Buffalo"
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

# Generando un Nuevo Proyecto

Ahora que ya tienes instalado el framework Buffalo, en esta sección, tu vas a aprender como crear **Una nueva Aplicación web** usando el comando `buffalo`.

## Creando un Nuevo Proyecto

El objetivo de Buffalo es que la construcción de una aplicación web en Go sea lo mas **rapido y  sencilla** posible. Que podria ser mas sencillo que un generador de un *nueva aplicación*?

Comienza por ir a tu carpeta favorita donde quieras crear tu proyecto, luego:

```bash
$ buffalo new coke
```

Esto generará aplicacion de Buffalo llamada **Coke**, lista para funcionar:
* El **Diseño del framework de Buffalo** y su configuracion por defecto ([pop/soda](https://github.com/gobuffalo/pop) con soporte a PostgreSQL).
* Toda las **Dependencias de Go** necesarias para correr la aplicación.
* **Dependencias del lado Frontend** y la configuración con [webpack](https://webpack.js.org/).
* Y un **Repositorio de Git** inicial.

```bash
$ buffalo new coke
DEBU[2022-05-25T11:06:33-05:00] Step: 435aea40
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] Exec: go mod init coke
go: creating new go.mod: module coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/README.md
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/actions_test.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/app.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/home.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/home_test.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/render.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/cmd/app/main.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.codeclimate.yml
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.env
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/fixtures/sample.toml
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/grifts/init.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/inflections.json
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/config/buffalo-app.toml
DEBU[2022-05-25T11:06:33-05:00] Step: 638bde0d
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/Dockerfile
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.dockerignore
DEBU[2022-05-25T11:06:33-05:00] Step: 7065092d
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/grifts/db.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/models/models.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/models/models_test.go
DEBU[2022-05-25T11:06:33-05:00] Step: 916dfca0
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/database.yml
DEBU[2022-05-25T11:06:33-05:00] Step: ff1c6d38
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.buffalo.dev.yml
DEBU[2022-05-25T11:06:33-05:00] Step: 103396c6
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] Exec: go install github.com/gobuffalo/buffalo-pop/v3@latest
DEBU[2022-05-25T11:06:34-05:00] Step: ae1260b1
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/config/buffalo-plugins.toml
DEBU[2022-05-25T11:06:34-05:00] Step: 4992ff46
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/app.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/home.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/home_test.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/render.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/locales/all.en-us.yaml
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/locales/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/robots.txt
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/_flash.plush.html
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/application.plush.html
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/home/index.plush.html
DEBU[2022-05-25T11:06:34-05:00] Step: 69d71878
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] LookPath: yarn
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/css/_buffalo.scss
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/css/application.scss
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/images/favicon.ico
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/images/logo.svg
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/js/application.js
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/.babelrc
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/package.json
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/postcss.config.js
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/assets/keep
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/webpack.config.js
DEBU[2022-05-25T11:06:34-05:00] LookPath: yarn
DEBU[2022-05-25T11:06:34-05:00] Exec: yarn --version
1.22.17
DEBU[2022-05-25T11:06:34-05:00] Exec: yarn set version berry
➤ YN0000: Retrieving https://repo.yarnpkg.com/3.2.1/packages/yarnpkg-cli/bin/yarn.js
➤ YN0000: Saving the new release in .yarn/releases/yarn-3.2.1.cjs
➤ YN0000: Done in 0s 637ms
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn config set enableGlobalCache true
➤ YN0000: Successfully set enableGlobalCache to true
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn config set logFilters --json [{"code":"YN0013","level":"discard"}]
➤ YN0000: Successfully set logFilters to [
  {
    code: 'YN0013',
    text: undefined,
    pattern: undefined,
    level: 'discard'
  }
]
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn --version
3.2.1
DEBU[2022-05-25T11:06:38-05:00] Exec: yarn install
DEBU[2022-05-25T11:06:38-05:00] ➤ YN0000: ┌ Resolution step

DEBU[2022-05-25T11:06:39-05:00] ➤ YN0032: │ fsevents@npm:2.3.2: Implicit dependencies on node-gyp are discouraged

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: └ Completed in 5s 401ms

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: ┌ Fetch step

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: └ Completed

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: ┌ Link step

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: │ ESM support for PnP uses the experimental loader API and is therefore experimental

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0007: │ @fortawesome/fontawesome-free@npm:5.15.4 must be built because it never has been before or the last one failed

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: └ Completed in 0s 906ms

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: Done with warnings in 6s 462ms

DEBU[2022-05-25T11:06:44-05:00] Step: bb3c28ed
DEBU[2022-05-25T11:06:44-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:44-05:00] Exec: go mod tidy
go: finding module for package github.com/gobuffalo/buffalo
go: finding module for package github.com/gobuffalo/mw-paramlogger
go: finding module for package github.com/gobuffalo/buffalo-pop/v3/pop/popmw
go: finding module for package github.com/gobuffalo/mw-i18n/v2
go: finding module for package github.com/gobuffalo/buffalo/render
go: finding module for package github.com/gobuffalo/mw-forcessl
go: finding module for package github.com/gobuffalo/envy
go: finding module for package github.com/gobuffalo/mw-csrf
go: finding module for package github.com/unrolled/secure
go: finding module for package github.com/markbates/grift/grift
go: finding module for package github.com/gobuffalo/pop/v6
go: finding module for package github.com/gobuffalo/suite/v4
go: found github.com/gobuffalo/buffalo in github.com/gobuffalo/buffalo v0.18.7
go: found github.com/gobuffalo/buffalo-pop/v3/pop/popmw in github.com/gobuffalo/buffalo-pop/v3 v3.0.4
go: found github.com/gobuffalo/buffalo/render in github.com/gobuffalo/buffalo v0.18.7
go: found github.com/gobuffalo/envy in github.com/gobuffalo/envy v1.10.1
go: found github.com/gobuffalo/mw-csrf in github.com/gobuffalo/mw-csrf v1.0.0
go: found github.com/gobuffalo/mw-forcessl in github.com/gobuffalo/mw-forcessl v0.0.0-20220514125302-be60179938a4
go: found github.com/gobuffalo/mw-i18n/v2 in github.com/gobuffalo/mw-i18n/v2 v2.0.1
go: found github.com/gobuffalo/mw-paramlogger in github.com/gobuffalo/mw-paramlogger v1.0.0
go: found github.com/unrolled/secure in github.com/unrolled/secure v1.10.0
go: found github.com/markbates/grift/grift in github.com/markbates/grift v1.5.0
go: found github.com/gobuffalo/pop/v6 in github.com/gobuffalo/pop/v6 v6.0.4
go: found github.com/gobuffalo/suite/v4 in github.com/gobuffalo/suite/v4 v4.0.2
DEBU[2022-05-25T11:06:45-05:00] Exec: go mod download
DEBU[2022-05-25T11:06:46-05:00] Step: 218a906c
DEBU[2022-05-25T11:06:46-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:46-05:00] Step: a3cee09e
DEBU[2022-05-25T11:06:46-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:46-05:00] File: /your/path/coke/.gitignore
DEBU[2022-05-25T11:06:46-05:00] Exec: git init
hint: Using 'master' as the name for the initial branch. This default branch name
hint: is subject to change. To configure the initial branch name to use in all
hint: of your new repositories, which will suppress this warning, call:
hint:
hint: 	git config --global init.defaultBranch <name>
hint:
hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
hint: 'development'. The just-created branch can be renamed via this command:
hint:
hint: 	git branch -m <name>
Initialized empty Git repository in /your/path/coke/.git/
DEBU[2022-05-25T11:06:46-05:00] Exec: git add .
DEBU[2022-05-25T11:06:46-05:00] Exec: git commit -q -m Initial Commit
INFO[2022-05-25T11:06:46-05:00] Congratulations! Your application, coke, has been successfully generated!
INFO[2022-05-25T11:06:46-05:00] You can find your new application at: /your/path/coke
INFO[2022-05-25T11:06:46-05:00] Please read the README.md file in your new application for next steps on running your application.
```

## Crear una App Personalizada
La configuración por defecto esta perfecta, pero tal vez eso no se adapte a ti. Buffalo provee algunas opciones como banderas para el comando `new`

Tu puede obtener las banderas disponibles usando el comando `help`

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci, circleci] (default "none")
      --config string        config file (default is $HOME/.buffalo.yaml)
      --db-type string       specify the type of database you want to use [cockroach, mariadb, mysql, postgres, sqlite3] (default "postgres")
  -d, --dry-run              dry run
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --module string        specify the root module (package) name. [defaults to 'automatic']
      --skip-config          skips using the config file
      --skip-docker          skips generating the Dockerfile
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
```

Tu puede elegir generar una aplicacion API, saltando el material frontend. ¿Tal vez tu quieres configurar un CI para construir tu aplicación en tu sistema favorito? ¿O incluso usar tu propio paquete para manejar la base de datos? ¡Solo usa la bandera!

## Sobreescribir Configuracion por Defecto

Por defecto el comando `buffalo new` va a mirar el archivo de configuracion `$HOME/.buffalo.yml` y si existe tratará de cargarlo. Tu puedes sobreescribir las banderas encontrandolas en ese archivo pasando lo correcto en la linea de comandos o utilizar la bandera  `--config` para especificar un archivo YAML diferente. Si la bandera `--skip-config` es usado. El comando `buffalo new` no va a cargar ningun archivo de configuracion y solo va a usar las banderas pasadas por la linea de comandos.

Un ejemplo de `.buffalo.yml` puede ser:

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

## Corriendo Tu Aplicacion en Desarrollo

{{< note >}}
Antes de comenzar con buffalo por primera vez, por favor dirijase a documentacion de [Base de datos](/documentation/database/pop) y lea un poco sobre la configuracion de tus bases de datos.
{{< /note >}}

Una de las desventajas de usar Go en desarrollo es la falta de “recarga” de codigo. Esto quiere decir que a medida que cambias tu codigo **tu necesitas parar manualmente** tu aplicacion, reconstruirla y luego reiniciarla. Buffalo encuentra esto muy molesto y quiere hacer la vida mas facil para ti.

```bash
$ buffalo dev
```

El comando `dev` va a mirar tus archivos `.go` and `.html` y la carpeta [asset](/documentation/frontend-layer/assets) por defecto. Eso va a **reconstruir y reestablecer tu binario para ti** automaticamente, asi que no te preocupes por esas cosas.

Solo corre el comando `buffalo dev` y ve a [localhost:3000/](http://localhost:3000/) para todos los cambios.

<figure>
  <img src="/assets/images/new-coke.png" title="screenshot">
  <figcaption>The brand new Coke app.</figcaption>
</figure>

#### Corre el servidor de desarrollo en un puerto personalizado

Algunas veces tu vas a tener una aplicación en un puerto 3000. Tu puedes configurar el puesto del servidor de desarrollo proporcionando la variable de entorno `PORT`.

```bash
$ PORT=3001 buffalo dev 
```

Tu puedes tambien mirar el capitulo de [Env Variables](/documentation/getting_started/configuration) para mas information de la configuracion de Buffalo.

## Siguiente Paso

* [Estructura de Directorios](/documentation/getting_started/directory-structure) - Aprende mas acerca de la estructura de Buffalo.
