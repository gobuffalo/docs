---
name: Instalar Buffalo
seoDescription: "Cómo instalar el framework Buffalo"
seoKeywords: 
  - "buffalo"
  - "go"
  - "golang"
  - "installation"
  - "framework"
  - "web"
  - "mac"
  - "windows"
  - "linux"
weight: 1
aliases:
  - /docs/getting-started/installation
  - /en/docs/getting-started/installation
  - /es/docs/getting-started/installation
---

# Instalar Buffalo

En este capítulo, aprenderá a instalar Buffalo, ya sea con binarios preconstruidos o desde el codigo fuente.

Buffalo proporciona **dos componentes principales**:
* La herramienta `buffalo`, es una herramienta potente para ayudarle a desarrollar en una manera rápida y eficiente.
* El framework `buffalo`, es una colección de piezas para construir su aplicación.

Buffalo está actualmente disponible y probado en los siguientes sistemas operativos:
* GNU/Linux
* Mac OSX
* Windows

## Requisitos

Antes de instalar, asegúrese de que tiene instalado los requisitos necesarios:
* [Un entorno de trabajo de Go](http://gopherguides.com/before-you-come-to-class)
* [Go](https://golang.org) version `{{< mingoversion >}}`.

##### Requisitos de frontend

La instalación de Buffalo requiere que tenga instalado el siguiente paquete de frontend:
Los siguientes requisitos son opcionales. Tu no necesitas tener esto si tu quieres construir una API o si tu prefieres construir tu aplicacion en una manera mas antigua.

* [Node](https://github.com/nodejs/node) version `8` o superior
* Ya sea [Yarn](https://yarnpkg.com/en/) o [npm](https://github.com/npm/npm) para la [Asset Pipeline](/documentation/frontend-layer/assets) construido sobre [Webpack](https://github.com/webpack/webpack).

##### Requisitos de backend

Nuevamente, si tu no necesitas una base de datos, tu no necesitas tener esto.

* **SQLite 3**: GCC, o un compilador de c equivalente para [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

## Instalacion desde un archivo de descarga - 64 bits

{{< note >}}
El paquete de descarga contiene Buffalo sin Soporte para SQLite.
{{< /note >}}

Desde `v0.10.3`, Los archivos pre-compilados se proporcionan con cada lanzamiento de Buffalo. Si no necesitas tener la ultima version, probablemente prefieras instalar esta versión.

### GNU / Linux

```sh
$ wget https://github.com/gobuffalo/cli/releases/download/{{< latestclirelease >}}/buffalo_{{< latestclibinaryversion >}}_Linux_x86_64.tar.gz
$ tar -xvzf buffalo_{{< latestclibinaryversion >}}_Linux_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### MacOS

```sh
$ curl -OL https://github.com/gobuffalo/cli/releases/download/{{< latestclirelease >}}/buffalo_{{< latestclibinaryversion >}}_Darwin_x86_64.tar.gz
$ tar -xvzf buffalo_{{< latestclibinaryversion >}}_Darwin_x86_64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
# or if you have ~/bin folder setup in the environment PATH variable
$ mv buffalo ~/bin/buffalo
```

## Scoop (Windows)
Buffalo puede ser instalado usando el manejador de paquetes [scoop](http://scoop.sh/)

```powershell
PS C:\> scoop install buffalo
```

## Chocolatey (Windows)
Buffalo puede ser instalado usando el manejador de paquetes [Chocolatey](https://chocolatey.org/packages/buffalo). Las versiones de Chocolatey se publican con un potencial retraso deben pasar por una revisión antes de usarlo.

```powershell
PS C:\> choco install buffalo
```

## Homebrew (macOS)

En macOS, Buffalo puede ser instalado con [Homebrew](https://brew.sh/). Despues de que tengas homebrew [Instalado](https://docs.brew.sh/Installation), tu facilmente puedes instalar buffalo:

```sh
brew install gobuffalo/tap/buffalo
```

## GoFish (Cross-Platforms)

[GoFish](https://gofi.sh/index.html) es un sistema multiplataforma de manejador de paquetes que trabaja en Windows, MacOSX and Linux.

Despues que tengas GoFish [Instalado](https://docs.gofi.sh/installation/), tu puedes instalar Buffalo:

```sh
$ gofish install buffalo
==> Installing buffalo...
🐠  buffalo {{< latestclirelease >}}: installed in 3.223672926s
```

## Instalacion personalizada **con** soporte SQLite3

**SQLite 3** Requiere un GCC o un compilador de c equivalente para [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) para compilar. Tu **deberias** tener un GCC instalado **primero** antes de instalar buffalo.

```sh
$ go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@{{< latestclirelease >}}
```

**Usuarios de Windows**: Sigue la guia de instalacion en [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) para instalar un GCC para Windows 10. En otra alternativa puede ser instlado con el manejador de paquetes [Scoop](http://scoop.sh/):

```powershell
PS C:\> scoop install gcc
```

{{< note >}}
Estas instrucciones pueder ser usadas tambien para actualizar a una nueva version de Buffalo.
{{< /note >}}

## Instalacion personalizada **sin** soporte SQLite3

```sh
$ go install github.com/gobuffalo/cli/cmd/buffalo@{{< latestclirelease >}}
```

{{< note >}}
Estas instrucciones pueder ser usadas tambien para actualizar a una nueva version de Buffalo.
{{< /note >}}

## Verifica tu instalación

Tu puedes comprobar si tu instalacion esta funcionando ejecutando `buffalo` en la terminal:

```sh
$ buffalo
Build Buffalo applications with ease

Usage:
  buffalo [command]

Available Commands:
  build       Build the application binary, including bundling of webpack assets
  completion  Generate the autocompletion script for the specified shell
  db          [PLUGIN] [DEPRECATED] please use `buffalo pop` instead.
  destroy     Destroy generated components
  dev         Run the Buffalo app in 'development' mode
  fix         Attempt to fix a Buffalo applications API to match version v0.18.6
  generate    Generate application components
  help        Help about any command
  info        Print diagnostic information (useful for debugging)
  new         Creates a new Buffalo application
  plugins     tools for working with buffalo plugins
  pop         [PLUGIN] A tasty treat for all your database needs
  routes      Print all defined routes
  setup       Setup a newly created, or recently checked out application.
  task        Run grift tasks
  test        Run the tests for the Buffalo app. Use --force-migrations to skip schema load.
  version     Print the version information

Flags:
  -h, --help   help for buffalo

Use "buffalo [command] --help" for more information about a command.
```
Si tu tienes un resultado similar, tu instalacion de Buffalo esta funcionando.

## Siguiente paso

* [Genera un nuevo proyecto](/es/documentation/getting_started/new-project) - Create tu primero proyecto de Buffalo!
