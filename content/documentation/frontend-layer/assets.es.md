---
Name: "Recursos"
aliases:
  - /docs/assets
  - /ess/docs/assets
---
# Recursos

Por defecto, cuando se genera una aplicación nueva de Buffalo a través del comando `buffalo new`, se genera un archivo de configuracion de [Webpack](https://webpack.github.io), y la aplicación se configura para usar Webpack la cadena de recursos de la aplicaión.

Si no se encuentra [`npm`](https://www.npmjs.com) en la maquina que genera la nueva aplicación de Buffalo, entonces Webpack no se configurará y se omitirá la cadena de recursos.

La cadena de recursos también puede ser omitido durante la generación de la aplicación con el flag `--skip-webpack`.

## JavaScript

La cadena de recursos está configurada inicialmente para soportar archivos JavaScript ES6, siendo `/assets/js/application.js` el principal punto de entrada.

Los siguientes elementos se instalan y configuran automáticamente durante la configuración de la cadena de recursos:

* [jQuery](https://jquery.com)
* [Bootstrap 4](http://getbootstrap.com)
* [jQuery UJS](https://github.com/rails/jquery-ujs)

Ninguno de los paquetes instalados son necesarios y se pueden remover. Se incluyen por comodidad.

## CSS

Por defecto, la cadena de recursos se configura para usar archivos [.scss](http://sass-lang.com), siendo `/assets/css/application.scss` el principal punto de entrada. Esto, por supuesto, se puede cambiar.

Los siguientes elementos se instalan y configuran automáticamente durante la configuración de la cadena de recursos:

* [Bootstrap 4](http://getbootstrap.com)
* [Font Awesome](http://fontawesome.io)


Ninguno de los paquetes instalados son necesarios y se pueden remover. Se incluyen por comodidad.

## Otros recursos

Cualquier recurso colocado en la carpeta `/assets` se copiará en la "distribución" automáticamente, y se puede encontrar en `/assets/path/to/asset`.

## Recursos de Huellas digitales

En la version `v0.9.5` se introdujo el recurso de huella digital en la configuración por defecto de Webpack para las nuevas aplicaciones. La huella digital de los recursos funciona generando un hash del contenido del archivo y añadiéndolo al nombre del archivo. Así, por ejemplo, `application.js` podría pasar a ser `application.a8adff90f4c6d47529c4.js`. La ventaja de esto es que permite que los recursos se almacenen en la caché, pero también permite que esa caché se rompa cuando se ha hecho un cambio en el contenido de este archivo.

Ten en cuenta que para que esto funcione, Buffalo espera que haya un archivo `public/assets/manifest.json` que contenga los mapeos entre los archivos a los que haces referencia en los helpers (por ejemplo, `javascriptTag("application.js")`) y sus homólogos con hash. Esto no es algo de lo que tengas que preocuparte si utilizas la configuración por defecto de Webpack. Sin embargo, si eliges usar el flag `--skip-webpack` cuando generas el proyecto, ten en cuenta que tendrás que manejarlo tú mismo. Aunque tener el archivo de manifiesto presente no es estrictamente necesario para que tu aplicación se ejecute, puedes experimentar problemas de caché sin él durante el desarrollo.

**NOTA:** Las aplicaciones escritas antes de la versión `v0.9.5` pueden necesitar establecer un `AssetsBox` en tu archivo `actions/render.go` en el `render.Options`, si los recursos no se están renderizando correctamente. Se recomienda moverlo del archivo `actions/app.go` a su lugar. Esto no configurará la la huella digital de los recursos, pero se asegurará de que los recursos se rendericen correctamente. Mira los cambios [https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbddf205da1687721a8acd29043cR43](https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbddf205da1687721a8acd29043cR43) y [https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403fbR13](https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403fbR13) para ejemplos.

Por defecto, las nuevas aplicaciones están configuradas para tomar huellas digitales sólo de los archivos JavaScript y CSS.

## Helpers de Recursos

Con la introducción de la huella digital de recursos en la versión `v0.9.5` se hizo difícil encontrar los archivos de recursos porque el nombre del archivo cambiaba constantemente. Para ayudar con esto, se introdujeron tres nuevos helpers.

1. `assetPath` - Este helper devolverá la ruta del recurso solicitado. Por ejemplo, `<%= assetPath("application.js") %>` devolvería algo como `/assets/application.a8adff90f4c6d47529c4.js`.

2. `javascriptTag` - Este helper generará una etiqueta script `<script src="xxx"></script>` para el archivo de Javascript solicitado. Ejemplo: `<%= javascriptTag("application.js") %>` devolvería algo como `<script src="/assets/application.bd76587ded82386f388f.js" type="text/javascript"></script>`.

3. `stylesheetTag` - Este helper generará una etiqueta de estilo `<link href="xxx">` para el archivo CSS solicitado. Ejemplo: `<%= stylesheetTag("application.css") %>` devolcería algo como `<link href="/assets/application.bd76587ded82386f388f.css" media="screen" rel="stylesheet" />`.

## Construir recursos en Desarrollo

El comando `buffalo dev`, además de vigilar y reconstruir el binario Go de la aplicación, vigilará y reconstruirá también la cadena de recursos. No es necesario ejecutar nada especial.

## Creación de recursos para el despliegue

El comando `buffalo build` construirá la cadena de recursos, y lo adjuntará correctamente al binario Go generado. Un binario para ejecutarlos todos! Consilta [Packing](/es/documentation/deploy/packing) para más opciones sobre la construcción de recursos para el despliegue.
