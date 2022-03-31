<% seoDescription("Empacar una aplicación Buffalo") %>
<% seoKeywords(["buffalo", "go", "golang", "paquete", "binario", "compilar", "compilacio"]) %>

<%= h1("Empacar") %>

Ahora, su proyecto está listo para ser implementado. En esta sección, aprenderá a empaquetar una versión de su aplicación para implementarla en un servidor.

## El comando de compilación

Buffalo tiene un comando, `build`, que creará un **binario completo** de su aplicación, que incluye, entre otros; activos, migraciones, plantillas, etc. Si sigue el _"Buffalo Way"_, las cosas simplemente funcionan. Es una experiencia maravillosa :)

<%= partial("en/docs/building/build_trace.md") %>

Cuando finaliza la compilación, tiene un binario recién creado en la carpeta `bin`. También tendrá el **tiempo de compilación** y **chit commit SHA** grabado, por lo que los binarios serán "versionados".

## Personalizar la compilación

Para obtener la lista de opciones disponibles, use el comando de ayuda:

<%= partial("en/docs/building/build_options.md") %>

### Nombre / ubicación de los binarios

Por defecto, su aplicación se compilará en el directorio `bin` de su proyecto, y el nombre del ejecutable será el nombre que utilizó para crear el proyecto con el comando` new`.

Puede cambiar este nombre predeterminado utilizando el indicador `-o` o `-output`:

<%= partial("en/docs/building/output_flag.md") %>

De hecho, también puede cambiar el directorio de destino:

<%= partial("en/docs/building/output_dir.md") %>

### Extraer activos en un archivo Zip

De forma predeterminada, toda su aplicación está empaquetada en un único ejecutable, incluidos los activos. En las configuraciones de producción, es posible que desee servir estos activos con un servidor proxy (como Apache o NGINX) para reducir la carga de la aplicación. Incluso puede usar un *CDN* para manejar sus activos.

Buffalo proporciona una forma de extraer los recursos compilados de la aplicación en un único archivo, utilizando el indicador `-e` o `-extract-assets`:

<%= partial("en/docs/building/extract_assets.md") %>

Por defecto, el archivo de activos se coloca en el directorio *bin*, pero si cambia el directorio de salida ejecutable con el indicador `-o`, los activos se colocarán en el mismo directorio.

<%= partial("en/docs/building/extract_assets_layout.md") %>

## Opciones avanzadas

### Creación de binarios "estáticos" / CGO

Crear binarios enlazados estáticamente que contengan CGO, como SQLite3, puede ser complicado. Al usar el indicador `--estático` con `buffalo build`, los indicadores `--ldflags '-linkmode externos -extldflags" -static "'` se agregarán al comando `go build`.

### Etiquetas de _build_

Al construir un binario de Buffalo utilizando el comando `buffalo build`, puede pasar `--tags` y `--ldflags` al binario construido; como lo haría normalmente cuando usas las herramientas `go build`.

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

## Comandos binarios

### Modos

Los binarios, por defecto, se ejecutan en modo `desarrollo`, lo que significa que todos los subcomandos se ejecutarán también en ese modo. Para cambiar el modo, debe usar la variable de entorno `GO_ENV`.

```bash
$ GO_ENV=production ./coke
```

### Comandos disponibles

Una vez que se ha construido un binario, hay varios subcomandos que se pueden ejecutar en ese binario:

#### _Default_

El comando _default_, si solo ejecuta el binario, iniciará la aplicación.

#### migrate

El sub-comando `migrate` ejecutará las migraciones para la aplicación.

#### version

El sub-comando `version` dará salida a la información de la versión para el binario, incluyendo el nombre, el git commit SHA usado para construir el binario, y la hora en que se construyó el binario.

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task

El sub-comando `task` ejecuta tareas.

```bash
$ ./coke task greet

Hello World!
```

## Siguientes pasos

* [Usando un Proxy](/es/docs/deploy/proxy) - Integrar su aplicación con un servidor como NGINX.
* [Servicio de systemd](/es/docs/deploy/systemd) - Ejecutar su aplicación como un servicio systemd.
* [Proveedores de la nube](/es/docs/deploy/providers) - Despeglar su aplicación en un proveedor de la nube.
