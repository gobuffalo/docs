---
name: Directory structure
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

# Directory Structure

Buffalo provides you **a minimal directory structure** to work on your project. This structure **keeps the project clean**, and allow the generators to work. Don't try to reinvent the wheel, and **let Buffalo buy you more time** to do the interesting part of your app! :)

Now that you have a minimal new project, let's go through its contents.

## The Root Directory

Here is the structure of a Buffalo project:

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

### actions

This directory handles the **Controller** part of the MVC pattern. It contains the handlers for your URLs, plus:

* The `app.go` file to setup your app & routes,
* The `render.go` file to setup the template engine(s).

### assets

{{< note >}}
This directory is optional. If you don't need to use a frontend setup (API only, for instance), it can be removed.
{{< /note >}}

This directory contains **raw** assets which will be compiled/compressed & put in the [`public`](#public) directory.

### cmd

This folder contains the `main.go` file which bootstraps your app and starts it.

### grifts

{{< note >}}
This directory is optional. If you don't need to use [tasks](/documentation/guides/tasks), you can remove it.
{{< /note >}}

This directory contains the [tasks](/documentation/guides/tasks) powered by [grift](https://github.com/markbates/grift).

### locales

{{< note >}}
This directory is optional. If you use only one language, you can remove it and the i18n module from the `app.go` file in the `actions` directory.
{{< /note >}}

This directory is used by the <abbr title="internationalization">i18n</abbr> system. It will fetch the translation strings from here.

### models

{{< note >}}
If you use pop/soda with the integrated generator, it will generate the model files here.
{{< /note >}}

{{< note >}}
This directory is optional. If you don't need to use a database, it can be removed.
{{< /note >}}

This directory handles the **Model** part of the MVC pattern. It contains the `models.go` file to initialize the datasource connection, and the model files to reflect objects from the datasource.

### public

{{< note >}}
The contents of this directory are auto-generated.
{{< /note >}}

This directory contains the public (compiled/compressed) assets. If you use webpack, it will put its assets in this directory.

### templates

{{< note >}}
This directory is optional. If you don't need to use a frontend setup (API only, for instance), it can be removed.
{{< /note >}}

This directory handles the **View** part of the MVC pattern. It contains the project templates, used to render the views.

### tmp

{{< note >}}
The contents of this directory are auto-generated.
{{< /note >}}

This directory is used by the `buffalo dev` command to rebuild your project on every change. The temporary files that Buffalo works with are put here.

### database.yml

{{< note >}}
This file is optional. If you don't need a database, or if you want to handle the database without pop/soda, you can remove it.
{{< /note >}}

This file contains the database configuration for [pop/soda](https://github.com/gobuffalo/pop).

## Next Steps

* [Configuration](/documentation/getting_started/configuration) - Manage your app configuration.
* [Tooling Integration](/documentation/getting_started/integrations) - Work with Buffalo, using existing tools.
