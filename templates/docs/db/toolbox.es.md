<% seoDescription("Soda CLI") %>
<% seoKeywords(["buffalo", "go", "golang", "base de datos", "utils", "pop", "toolbox", "CLI", "soda"]) %>

<%= h1("Soda CLI") %>

Pop le ayuda a administrar las conexiones a la base de datos, pero también proporciona `soda`, una pequeña caja de herramientas CLI para administrar su base de datos. Puede ayudarlo a crear una nueva base de datos, descartar las existentes, y así sucesivamente.

<%= note() { %>
**Nota para los usuarios de Buffalo**: los comandos `soda` están incrustados en el comando `buffalo`, detrás del espacio de nombres `db`. Entonces, cada vez que quiera usar un comando de `soda`, simplemente ejecute `buffalo db` en su lugar. No es necesario instalar la CLI `soda`.
<% } %>

<%= title("Instalar soporte CLI") %>

**Sin** compatibilidad con sqlite 3:

```bash
$ go get github.com/gobuffalo/pop/...
$ go install github.com/gobuffalo/pop/soda
```

**Con** compatibilidad con sqlite 3 (requiere GCC o compilador de C equivalente):

```bash
$ go get -u -v -tags sqlite github.com/gobuffalo/pop/...
$ go install -tags sqlite github.com/gobuffalo/pop/soda
```

Si no está construyendo su código con `buffalo build`, también tendrás que pasar `-tags sqlite` a `go build` cuando construya su programa.

<%= title("Creación de bases de datos") %>

Una vez que el `database.yml` se ha configurado con las configuraciones apropiadas, y el servidor de la base de datos se está ejecutando, Soda puede crear todas las bases de datos en el archivo `database.yml` con un simple comando:

```bash
$ soda create -a
```

También se puede crear solo una de las bases de datos configuradas utilizando el indicador `-e` y el nombre de la base de datos:

```bash
$ soda create -e test
```

<%= title("Borrar base de datos") %>

Soda puede borrar todas sus bases de datos, si lo desea, con un solo comando:

```bash
$ soda drop -a
```

También puede borrar solo una de las bases de datos configuradas utilizando el indicador `-e` y el nombre de la base de datos:

```bash
$ soda drop -e test
```
