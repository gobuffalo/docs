<% seoDescription("¿Cómo usar Pop con Buffalo?") %>
<% seoKeywords(["buffalo", "go", "golang", "pop", "buffalo", "integración"]) %>

<%= h1("Integración de Buffalo") %>

Buffalo proporciona un middleware Pop para facilitar el uso de la base de datos dentro de Buffalo: https://github.com/gobuffalo/buffalo-pop

<%= title("El Middleware de Transacción Pop") %>

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