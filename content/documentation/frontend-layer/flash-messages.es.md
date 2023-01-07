---
Name: "Mensajes Flash"
weight: 8
aliases:
  - /docs/flash-messages
  - /es/docs/flash-messages
---

# Mensajes Flash

## Qué son los mensaje Flash?

Los mensajes flash son una forma de comunicar mensajes al usuario final desde el interior de la aplicacón. Estos mensajes pueden ser de tipo error, advertencia o éxito.

Algunos ejemplos de mensajes flash son:

* "Se ha cerrado la sesión con éxito."
* "Su widget no pudo ser actualizado".
* "Hubo un problema de acceso a su cuenta".

Poder establecer estos mensajes en un handler de Buffalo y luego pasarlos a las vistas es increíblemente útil.

## Estableciendo mensajes Flahs

La creación de mensajes flash puede hacerse fácilmente utilizando la función `c.Flash()` proporcionada en el [`buffalo.Context`](/documentation/request_handling/context).

```go
func WidgetsCreate(c buffalo.Context) error {
  // do some work
  c.Flash().Add("success", "Widget was successfully created!")
  // do more work and return
}
```

El nombre de la "llave", en este ejemplo, "success", se dejan a discreción de tu aplicación para que las uses como mejor te convenga. No hay llaves "especiales" o "predefinidas".

## Accediendo a los mensajes Flash en plantillas

{{<note>}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).

Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para más detalles sobre el paquete de plantillas.
{{</note>}}

### Recorriendo sobre todos los mensajes Flash

```html
<div class="row">
  <div class="col-md-12">
    <%= for (k, messages) in flash { %>
      <%= for (msg) in messages { %>
        <div class="alert alert-<%= k %>" role="alert">
          <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
          <%= msg %>
        </div>
      <% } %>
    <% } %>
  </div>
</div>
```

### Recorriendo sobre una llave de mensaje Flash

```html
<div class="row">
  <div class="col-md-12">
    <%= for (message) in flash["success"] { %>
      <div class="alert alert-success" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <%= message %>
      </div>
    <% } %>
  </div>
</div>
```
