---
name: Mutaciones
seoDescription: "Maneja las mutaciones con Pop"
weight: 20
aliases:
  - /docs/db/mutations
  - /es/docs/db/mutations
---

# Mutaciones

En este capitulo, aprenderas como crear, actualizar y eliminar objetos de tu base de datos con Pop.

## Crear

### Sin Validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{}

// Crea `fruit` sin ejecutar validaciones
err := c.Create(&fruit)
```

### Con Validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{}

// Ejecuta las validaciones y crea si pasaron todas las validaciones
vErrors, err := c.ValidateAndCreate(&fruit)
```

## Actualizar

### Sin Validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{}

// Actualiza `fruit` sin ejecutar validaciones
err := c.Update(&fruit)
```

### Con Validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{}

// Ejecuta las validaciones y actualiza si pasaron todas las validaciones
vErrors, err := c.ValidateAndUpdate(&fruit)
```

## Guardar

`Save` revisa el valor del campo ID del modelo: Si el ID es el valor cero del tipo (por ejemplo, si es de tipo `int` y su valor es `0`), `Save` llama al metodo `Create`; de lo contrario, llama al metodo `Update`.

### Sin Validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{ID: 0}

// Crea `fruit` sin ejecutar validaciones
err := c.Save(&fruit)
```

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{ID: 1}

// Actualiza `fruit` sin ejecutar validaciones
err := c.Save(&fruit)
```

### Con validación

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{ID: 0}

// Ejecuta las validaciones y crea si pasaron todas las validaciones
vErrors, err := c.ValidateAndSave(&fruit)
```

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{ID: 1}

// Ejecuta las validaciones y actualiza si pasaron todas las validaciones
vErrors, err := c.ValidateAndSave(&fruit)
```

## Eliminar

```go
c, err := pop.Connect("development")
// ...
fruit := models.Fruit{ID: 1}

// Borrar fruit
err := c.Destroy(&fruit)
```

## Siguientes pasos

* [Consultas](/es/documentation/database/querying) - Consulta datos de tu base de datos.
* [Asociaciones y Relaciones](/es/documentation/database/relations) - Manejar relaciones entre modelos.
