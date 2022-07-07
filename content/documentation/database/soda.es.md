---
name: Soda CLI
seoDescription: "Soda CLI"
seoKeywords: ["buffalo", "go", "golang", "database", "utils", "pop", "toolbox", "CLI", "soda"]
weight: 2
aliases:
  - /documentation/database/toolbox
  - /docs/db/toolbox
  - /es/docs/db/toolbox
---

# Soda CLI

Pop te ayuda a administrar la conexión a la base de datos, pero tambien provee `soda`, una pequeña caja de herramientas de CLI para administrar tu base de datos. Puede ayudarte a crear una nueva, borrar existentes, etc.

{{<note>}}
**Nota para usuarios de Buffalo**: Los comandos de `Soda` estan adheridos de los comandos de `Buffalo`, detrás del espacio de nombres de `pop`. Asi que cada vez que desees usar un comando de `soda`, solo ejecuta `buffalo pop` en su lugar. No necesitas instalar la CLI de `soda`.
{{</note>}}

## Instalando soporte CLI

### Desde un archivo desplegado

{{<note>}}
Los archivos pre-compilados contienen Soda **con soporte SQLite**
{{</note>}}

Descarga la version apropiada para tu plataforma en [Pop releases](https://github.com/gobuffalo/pop/releases).

Ubicalo en cualquier lugar de tu `PATH`, y asegurate que el binario de `soda` se pueda ejecutar.

### Homebrew (MacOS)

```bash
$ brew install gobuffalo/tap/pop
```

### Desde la fuente

{{<note>}}
Para Go `{{< mingoversion >}}` y superiores
{{</note>}}

**SIN** soporte de SQLite 3:

```bash
$ go install github.com/gobuffalo/pop/v6/soda@latest
```

**CON** soporte de SQLite 3 (requiere GCC o un compilador de C equivalente):

```bash
$ go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```
Si no estas compilando tu codigo con `buffalo build`, tendras que pasar tambien `-tag sqlite` a `go build` al construir tu programa.

## Creando bases de datos

Una vez que `database.yml` haya sido configurada con los ajustes apropiados y el servidor de la base de datos se esté ejecutando, `soda` puede crear todas las bases de datos que esten en el archivo `database.yml` con un simple comando:

```bash
$ soda create -a
```

Tambien puedes crear solo una base de datos configurada usando el flag `-e` y el nombre de la base de datos:

```bash
$ soda create -e test
```

## Borrando bases de datos

`soda` puede borrar todas las bases de datos si lo deseas, con un comando:

```bash
$ soda drop -a
```

Tambien puedes borrar solo una de las bases de datos configuradas usando el flag `-e` y el nombre de la base de datos:

```bash
$ soda drop -e test
```

## Siguientes pasos

* [Configuración](/es/documentation/database/configuration) - Configura tu conexión de la base de datos.
