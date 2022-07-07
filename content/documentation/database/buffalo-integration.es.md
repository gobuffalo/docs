---
name: Integración con Buffalo
seoDescription: "Como usar Pop con Buffalo?"
seoKeywords: ["buffalo", "go", "golang", "pop", "buffalo", "integration"]
weight: 4
aliases:
  - /docs/db/buffalo-ingetration
  - /es/docs/db/buffalo-ingetration
---

# Integración con Buffalo

## Generando una nueva App

Cuando generas una nueva aplicación de Buffalo puedes escoger la base de datos de destino con el flag `--db-type`. Por ejemplo, para generar una nueva app con soporte para la base de datos MySQl, puedes escribir lo siguiente:

```bash
$ buffalo new coke --db-type mysql
```

**Por defecto, buffalo generará una aplicacion con PostgreSQL como la base de datos de respaldo.**

### Saltar soporte de la base de datos

Si deseas manejar la base de datos sin usar Pop. o si estas construyento una app sin base de datos, es posible saltar la generacion de todos los componentes de base de datos con el flag `--skip-pop`.

```bash
$ buffalo new coke --skip-pop
```

## El middleware de transacción de Pop

Buffalo proporciona un middleware para facilitar el uso de la base de datos dentro de Buffalo: [buffalo-pop](https://github.com/gobuffalo/buffalo-pop)

### Configuración

Este middleware es configurado para ti por defecto, si escoges usar Pop cuando creas un nuevo proyecto.

```go
// actions/app.go

import "github.com/gobuffalo/buffalo-pop/v3/pop/popmw"

func App() *buffalo.App {
  // ...
  app.Use(popmw.Transaction(models.DB))
  // ...
  app.GET("/", HomeHandler)
  // ...
}
```

`popmw.Transaction(models.DB)` usa la conexión de la base de datos configurada para crear un nuevo middleware `Transaction`. Este middleware hace lo siguiente:

* Registra la duración total gastada durante la solicitud de realizar llamadas a la base de datos.
* envuelve **cada petición HTTP** en una transacción de base de datos.
* Hace commit **si no hay ningun error** ejecutando los middlewares y la accion; **y el estado de la respuesta es un 2xx o 3xx**
* De lo contrario, hace `Rollback`.

### Manejar la transacción a mano

Si necesitas manejar tu transacción a mano, tu puedas saltar el middleware para una ruta dada:

```go
// actions/app.go

import "github.com/gobuffalo/buffalo-pop/v3/pop/popmw"

func App() *buffalo.App {
  // ...
  txm := popmw.Transaction(models.DB)

  app.Use(txm)
  app.Middleware.Skip(txm, HomeHandler)
  // ...
  app.POST("/form", FormHandler)
  app.GET("/", HomeHandler)
  // ...
}
```
