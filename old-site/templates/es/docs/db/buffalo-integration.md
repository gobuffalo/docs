<% seoDescription("¿Cómo usar Pop con Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "pop", "buffalo", "integración"]) %>

# Integración de Buffalo

## Generar una nueva aplicación

Cuando genera una nueva aplicación Buffalo, puede elegir la base de datos de destino con el indicador `--db-type`. Por ejemplo, para generar una nueva aplicación con soporte de base de datos MySQL, puede escribir lo siguiente:

```bash
$ buffalo new coke --db-type mysql
```

**Por defecto, Buffalo generará una aplicación con PostgreSQL como base de datos.**

### Proceder sin soporte de base de datos

Si desea manejar la base de datos sin utilizar Pop, o si está creando una aplicación sin base de datos, también es posible omitir la generación de todos los componentes de la base de datos con el indicador `--skip-pop`.

```bash
$ buffalo new coke --skip-pop
```

## El Middleware de Transacción Pop

Buffalo proporciona un middleware Pop para facilitar el uso de la base de datos dentro de Buffalo: https://github.com/gobuffalo/buffalo-pop

### Configuración

Este middleware está configurado para usted de manera predeterminada, si elige usar Pop al crear un nuevo proyecto.

**actions/app.go**

```go
func App() *buffalo.App {
	if app == nil {
        // [...]

        app.Use(poptx.PopTransaction(models.DB))

        // [...]

        app.GET("/", HomeHandler)
    }

    return app
}
```

`poptx.PopTransaction(models.DB)` utiliza la conexión a la base de datos configurada para crear un nuevo middleware `PopTransaction`. Este middleware hace lo siguiente:

* Registre la duración total gastada durante la solicitud de llamadas a la base de datos.
* Ajustar **cada solicitud HTTP** en una transacción de base de datos.
* Confirme el cambio **si no hubo error** ejecutando los middlewares y la acción; **y el estado de respuesta es un 2xx o 3xx**.
* Realizar un _rollback_ de lo contrario.

### Manejar la transacción a mano

Si necesita manejar una transacción a mano, puede omitir el middleware para una ruta determinada:

```go
func App() *buffalo.App {
	if app == nil {
        // [...]
        txm := poptx.PopTransaction(models.DB)
        app.Use(txm)
        a.Middleware.Skip(txm, HomeHandler)

        // [...]

        app.POST("/form", FormHandler)
        app.GET("/", HomeHandler)
    }

    return app
}
```