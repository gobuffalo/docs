---
name: Solución de problemas
icon: "images/troubleshoot.svg"
aliases:
  - /es/docs/troubleshooting
---

# Solución de problemas

{{< faq "La App falla con `securecookie: hash key is not set`" "securecookie-hash-key-not-set">}}
Después de un cambio reciente en [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions), las aplicacionesde Buffalo con el error `securecookie: hash key is not set` no van a poder iniciarse.

Para solucionar esto, debes definir una variable de entorno llamada `SESSION_SECRET`.

Para más información: [github.com/gobuffalo/buffalo/issues/1067](https://github.com/gobuffalo/buffalo/issues/1067)
{{< /faq >}}

{{< faq "Linea de comandos lenta" "slow-commands">}}
Si al ejecutar `buffalo --help` o cualquier otro comando en la terminal toma mas tiempo de lo esperado, define `export BUFFALO_PLUGIN_PATH=$GOPATH/bin` en la configuracion de tu terminal (p.e. `.bash_profile`).
{{< /faq >}}

{{< faq "No de encuentra el binario de `buffalo`" "binary-not-found">}}
Si no puedes encontrar el binario de `buffalo` después de una instalación exitosa, probablemente es causado porque el `$GOPATH/bin` o `%GOPATH\bin` (Windows), no están en tu variable `$PATH`. Cuando se instala un binario de Go se coloca en `$GOPATH/bin`. Añadiendo esto a tu `$PATH` global, te permitirá encontrar *cualquier* binario de Go en todas partes de tu sistema.

Ve a [golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH) para obtener más detalles.
{{< /faq >}}

{{< faq "`buffalo new` falla al generar un proyecto completo" "failed-to-gen">}}
Esto ocurre porque el comando `bufalo new` no puede encontrar las plantillas que necesita para general una nueva aplicación.

Hay un par de cosas que pueden causar este error.

* Usando múltiples `$GOPATH`s. Esto puede suceder cuando instala Buffalo en un `$GOPATH` y luego crea un nuevo `$GOPATH` temporal e intenta crear una nueva aplicación allí. Debido a que las plantillas están en el primer `$GOPATH` original, el instalador no las encuentra y, posteriormente, genera una aplicación incompleta. Para arreglar esto, define un solo `$GOPATH`.

* Usando un único `$GOPATH`: Si no estas usando múltiples `$GOPATH`s y continuas viento este problema, es mas probable que se deba a una mala instalación, Ejecuta de nuevo `$ go get -u -v github.com/gobuffalo/buffalo` y debería , con suerte, reparar la instalación por ti.

El ticket original para este problema puede ser encontrado en:
[github.com/gobuffalo/buffalo/issues/629](https://github.com/gobuffalo/buffalo/issues/629).
{{< /faq >}}

{{< faq "`buffalo new` falla por problemas de permisos de NPM" "npm-permissions">}}
Esto es causado por una incorrecta instalación de Node/NPM.

Ve a [docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally) para informarte como arreglar este problema.
{{< /faq >}}


{{< faq "la reconstrucción automática de `buffalo dev` no funciona con NFS." "nfs-rebuild">}}
Esto se debe a que el paquete `fsnotify` no es compatible con NFS.

Ve a [github.com/gobuffalo/buffalo/issues/620](https://github.com/gobuffalo/buffalo/issues/620) para mas detalles y una solución.
{{< /faq >}}

{{< faq "`buffalo new` fails with `import path does not begin with hostname`" "import-begins-hostname">}}
Esto es causado porque el `$GOPATH` y la carpeta del sistema no coinciden.


```text
GOPATH: /Users/foobar/Documents/Programming/Go
Actual: /Users/foobar/Documents/programming/go
```

Esos no son lo mismo y causan problemas con muchas herramientas de Go. Arregla `$GOPATH` para que coincida con el sistema de archivos e intenta nuevamente.
{{< /faq >}}

{{< faq "`buffalo new` falla al buscar `golang.org/x/tools/go/gcimporter`" "gcimporter">}}
Esto es causado por una copia obsoleta del paquete `github.com/motemen/gore`.

Para arreglar esto, simplemente actualiza `gore`.

```text
$ go get -u github.com/motemen/gore
```

Para mas información ve a [github.com/gobuffalo/buffalo/issues/108](https://github.com/gobuffalo/buffalo/issues/108) y [github.com/motemen/gore/issues/63](https://github.com/motemen/gore/issues/63).
{{< /faq >}}

{{< faq "`buffalo dev` falla al iniciar con `Unknown`" "fails-unknown">}}
cuando ejecutas `$ buffalo dev`, y te encuentras con este error:

```bash
There was a problem starting the dev server: Unknown, Please review the troubleshooting docs.
```

Esto se puede deberse a que a tu sistema le falta NodeJS/NPM, Asegurate que NodeJS/NPM esten instalados en tu `$PATH`. Si NodeJS/NPM estan en tu `$PATH`, intenta renombrar el archivo `webpack.config.js`.

Si continuas teniendo estos problemas despues de intentar los pasos de arriba, por favor contacta a la comunidad de [Gophers Slack #buffalo channel](https://gophers.slack.com/messages/buffalo/)
{{< /faq >}}

{{< faq "`package context: unrecognized import path \"context\" (import path does not begin with hostname)`" "unrecognized-context-import">}}
Cuando intentas instar Buffalo `go get` retorna este error:

```bash
package context: unrecognized import path "context" (import path does not begin with hostname)
```

Esto se debe a una version obsoleta de Go. Buffalo requiere Go de `{{< mingoversion >}}` o superior, Verifica tu instalación de Go y asegúrate de ejecutar la última versión.
{{< /faq >}}

{{< faq "Error: `unexpected directory layout:` durante `go get`" "unexpected-dir-layout">}}
Ocasionalmente cuando ejecutar  `go get` en buffalo obtendrás el siguiente error:

```bash
unexpected directory layout:
import path: github.com/mattn/go-colorable
dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
expand dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
separator: /
```

Este problema ha sido reportado anteriormente por el equipo de Go [github.com/golang/go/issues/17597](https://github.com/golang/go/issues/17597).

La mejor opción para resolver este problema es ejecutar nuevamente `go get`, y parece arreglarse solo.
{{< /faq >}}

{{< faq "Error: en `application.js` de UglifyJs" "appjs-uglify">}}
Si recives este error cuando ejecutas `buffalo build`, necesitas actualizar tu `webpack.config.js` para que funcione con el transpilador ES6 [github.com/gobuffalo/buffalo/pull/350/files](https://github.com/gobuffalo/buffalo/pull/350/files).
{{< /faq >}}

{{< faq "Error: `Killed 9` cuando ejecutas `buffalo` en Mac OS X con Go 1.8.0" "mac-killed9">}}
Este es un problema conocido con Go [github.com/golang/go/issues/19734](https://github.com/golang/go/issues/19734), no con Buffalo.

La mejor solucion es actualizar a Go 1.8.1, o superior, y reconstruir tus binarios de Go.
{{< /faq >}}

{{< faq "Mac OS X: error `Too many open files in system`" "mac-too-many-files">}}
Si recibe este error cuando ejecuta `buffalo dev`, significa que estas "viendo" demasiados archivos, ya sean archivos **.go** u otro tipo de archivos. 

Para corregir esto, puedes [cambiar el número máximo de archivos abiertos](http://blog.mact.me/2014/10/22/yosemite-upgrade-changes-open-file-limit) en tu sistema.
{{< /faq >}}

{{< faq "`buffalo new` falla tratando de ejecutar `goimports`" "new-goimports">}}
El error completo puede aparecer como el siguiente, y parece ser el resultado de herramientas obsoletas de Go. Para resolverlo, ejecuta `rm -r $GOPATH/src/golang.org/`, luego ejecuta nuevamente `go get`.

```bash
$ buffalo new myapp

--> go get -u golang.org/x/tools/cmd/goimports
package golang.org/x/tools/cmd/goimports: golang.org/x/tools is a custom import path for https://go.googlesource.com/tools, but /Users/foo/go/src/golang.org/x/tools is checked out from https://code.google.com/p/go.tools

Error: exit status 1
```
{{< /faq >}}

{{< faq "`buffalo g goth` falla al generar `auth.go`" "goth-auth-fails">}}
Es posible que veas errores similares a este cuando construyas:

```bash
buffalo: 2018/01/19 20:58:47 === Error! ===
buffalo: 2018/01/19 20:58:47 === exit status 2
path/path/models
path/path/actions
# path/path/actions
actions/app.go:17:2: gothic redeclared as imported package name
    previous declaration at actions/app.go:15:2
actions/app.go:66:36: undefined: AuthCallback
actions/app.go:67:11: undefined: SetCurrentUser
actions/app.go:68:11: undefined: Authorize
actions/app.go:69:23: undefined: Authorize
```

Esto podría deberse a que el complemento `goth` no puede encontrar las plantillas para los diferentes proveedores. Esto puede suceder cuando el plugin `goth` está disponible en `$PATH`, pero el proyecto no está en tu `$GOPATH` actual.

Para solucionarlo, puedes ejecutar `go get -u github.com/gobuffalo/buffalo-goth` en el `$GOPATH` de tu proyecto o usar `dep`.
{{< /faq >}}
