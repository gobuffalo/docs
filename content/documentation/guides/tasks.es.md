---
name: Tareas
weight: 14
seoDescription: Tareas
seoKeywords: ["buffalo", "go", "golang", "tasks", "scripts", "grift"]
aliases:
  - /docs/tasks
  - /es/docs/tasks
---
# Tareas

Las tareas son pequeños scripts que a menudo se necesitan al escribir una aplicación. Estos tasks pueden ser similares a la inicialización de una base de datos, analizar un archivo de registro, o incluso el lanzamiento de un script. Buffalo usa la librería [grift](https://github.com/gobuffalo/grift) para simplificar la escritura de estas tareas.

{{< vimeo 213096302>}}

## Escribiendo tareas

Las tareas deben estar en el paquete `grift`. Una tarea simple ser;ia la siguiente:

```go
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
```

## Generador de tareas

```bash
$ buffalo g task foo:bar

--> grifts/bar.go
```

```go
// grifts/bar.go
package grifts

import (
	. "github.com/gobuffalo/grift/grift"
)

var _ = Namespace("foo", func() {

	Desc("bar", "Task Description")
	Add("bar", func(c *Context) error {
		return nil
	})

})

```


## Listing Available Tasks

```bash
$ buffalo task list

Available grifts
================
buffalo task db:seed       # Seeds a database
buffalo task middleware    # Prints out your middleware stack
buffalo task routes        # Print out all defined routes
buffalo task secret        # Generate a cryptographically secure secret key
```

## Ejecutando Tareas

### Development

Las tareas pueden ser ejecutadas en el entorno de `development` usando el comando `buffalo task`.

```bash
$ buffalo task hello
```

### Desde un binario construido

Después que un binario sea [construido](/documentation/deploy/packing), las tareas se pueden ejecutar con el subcomando `task`:

```bash
$ myapp task hello
```
