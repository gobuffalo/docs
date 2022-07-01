---
name: Estructura de Directorio
seoDescription: "Buffalo project directory structure"
seoKeywords: 
  - "buffalo"
  - "go"
  - "golang"
  - "structure"
  - "directory"
  - "framework"
  - "web"
weight: 3
aliases:
  - /docs/getting-started/directory-structure
  - /en/docs/getting-started/directory-structure
---

# Estructura de Directorio

Buffalo te proporciona a ti **una estructura de directorio minimo** para trabajar en tu proyecto. Esta estructura **mantiene el proyecto limpio**, y permite el generador trabajar. No trates de reinventar la rueda, y **deja que buffalo te compre mas tiempo** para hacer la parte interesante de tu app! :)

Ahora que tu tienes un proyecto minimo, vamos a su contenido:

## Directorio Raiz

Aqui esta la estructura de un Proyecto Buffalo:

``` erb
├── .yarn/
├── actions/
│	├── app.go
│	└── render.go
├── assets/
├── cmd/
│	└── app/
│		└── main.go
├── config/
├── fixtures/
├── grifts/
├── locales/
├── models/
├── public/
├── templates/
├── .babelrc
├── .buffalo.dev.yml
├── .codeclimate.yml
├── .docketignore
├── .env
├── .gitignore
├── .pnp.loader.mjs
├── .yarnrc.yml
├── database.yml
├── Dockerfile
├── go.mod
├── go.sum
├── inflections.json
├── package.json
├── postcss.config.js
├── README.md
├── webpack.config.js
└── yarn.lock
```

### Actions

Este directorio maneja la parte de **Controlador** del patron MVC. Eso contiene los controladores para tus URLs, un plus:

* El `app.go` es el archivo para configurar tu app y rutas.
* El `render.go` es el archivo para configurar el motor de las plantillas.

### Assets

{{< note >}}
Este directorio es opcional. Si tu no necesitas usar una configuracion para el frontend (Solo API, por ejemplo), puedes quitarlo sin problemas.

{{< /note >}}

Este directorio contiene *raw* assets el cual sera compilado/comprimido y puestos en el directorio [`public`](#public).

### Cmd

Esta carpeta contiene el `main.go` es el archivo el cual arranca tu app y la inicia.

### Grifts

{{< note >}}
Este directorio es opcional. Si tu no necesitas usarlo [tasks](/documentation/guides/tasks), puedes quitarlo sin problemas.

{{< /note >}}


Este directorio contiene el [tasks](/documentation/guides/tasks) impulsado por [grift](https://github.com/markbates/grift).

### Locales

{{< note >}}
Este directorio es opcional. Si tu usas un solo lenguaje, tu puedes remover esto y el modulo i18n desde el archivo `app.go` en el directorio `actions`
{{< /note >}}

Este directorio is usado por el sistema de <abbr title="internationalization">i18n</abbr>. Eso va a traer la traduccion de cadenas desde aqui.

### Models

{{< note >}}
Si tu usas pop/soda con el generador integrado, eso generará los archivos del modelo aqui.
{{< /note >}}

{{< note >}}
Este directorio es opcional. Si tu no necesitas usar una base de datos, puedes removerlo.
{{< /note >}}

Este directorio se encarga de la parte del **Model** de el patrón MVC. Eso contiene el archivo `models.go` para inicializar la conexion con la fuente de datos y el modelo para reflejar objetos de la base de datos.

### Public

{{< note >}}
El contenido de este directorio se genera automaticamente.
{{< /note >}}

Este directorio contiene asset publico (Compilado/Comprimido). Si tu usas webpack, los assets se van a colocar en este directorio.

### Templates

{{< note >}}
Este directorio es opcional. Si tu no necesitas usar una configuracion de frontend (Solo API, por ejemplo), puedes quitarlo sin problemas.
{{< /note >}}

Este directorio maneja la **Vista** de el patron MVC. Eso contiene los templates de el proyecto, usado para renderizar las vistas.

### Tmp

{{< note >}}
El contenido de este directorio es autogenerado.
{{< /note >}}

Este directorio es usado por el comando `buffalo dev` para reconstruir tu proyecto con cada cambio. Los archivos temporales de Buffalo se van a colocar aqui.

### Database.yml

{{< note >}}
Este archivo es opcional. Si tu no necesitas una base de datos, o si tu quieres manejar la base de datos sin pop/soda puedes removerlo.
{{< /note >}}

Este archivo contiene la configuracion de la base de datos para [pop/soda](https://github.com/gobuffalo/pop).

## Siguiente Paso

* [Configuración](/documentation/getting_started/configuration) -  Maneja la configuracion de tu app.
* [Integracion de Herramientas](/documentation/getting_started/integrations) - Trabaja con Buffalo usando herramientas existentes.
