---
Name: "Formularios"
weight: 9
aliases:
  - /docs/forms
  - /es/docs/forms
---

# Forms

Buffalo usa el paquete [`github.com/gobuffalo/tags`](https://github.com/gobuffalo/tags) para facilitar la construcción de formulario.

Plush incluye dos helpers de este paquete que genera un formulario con estulo de Boostrap v3. Estos helpers son `form` y `form_for`.

Ambos tipos de helpers de formulario tienen las siguientes caracteristicas en común:


* Configuración automática del token de autenticidad CSRF
* Soporte para todos los métodos HTTP (PUT, POST, DELETE, etc...)
* [Manejo de errores](#error-handling)
* Soporte de formularios multipart
* Tipos de entrada personalizables
* Paso de atributos de etiquetas HTML

## Formularios básicos

El helper `form` se puede usar para generar formularios HTML. Dado que este tipo de formulario no está amarrado a ningún "modelo" en particular, toda la información se debe pasar como opciones al formulario y sus métodos.

{{<codetabs>}}
{{<tab "templates/talks/edit.html" >}}
```erb
// templates/talks/edit.html

<%= form({action: talkPath({id: 3}), method: "PUT"}) { %>
  <div class="row">
    <div class="col-md-12">
      <%= f.InputTag({name:"Title", value: talk.Title }) %>
    </div>

    <div class="col-md-6">
      <%= f.TextArea({value: talk.Abstract, hide_label: true }) %>
    </div>

    <div class="col-md-6">
      <%= f.SelectTag({name: "TalkFormatID", value: talk.TalkFormatID, options: talk_formats}) %>
      <%= f.SelectTag({name: "AudienceLevel", value: talk.AudienceLevel, options: audience_levels }) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea({name: "Description", value: talk.Description, rows: 10}) %>
    </div>
    <div class="col-md-12">
      <%= f.TextArea({notes:"Notes", value: talk.Notes, rows: 10 }) %>
    </div>

  </div>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```html
// OUTPUT
<form action="/talks/3" method="POST">
  <input name="authenticity_token" type="hidden" value="e0c536b7a1a7d752066727b771f1e5d02220ceff5143f6c77b">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <input class=" form-control" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class=" form-control">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
        <select class=" form-control" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <select class=" form-control" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" notes="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
{{< /tab >}}
{{<  /codetabs  >}}



## Formularios de modelos

El helper `form_for` se puede usar para generar formularios HTML para un modelo específico. Esto hace que el código sea más fácil de escribir, y mantiene un nivel de "consistencia" a través de tu aplicación.

El helper `form_for` se comporta similar al helper `form`, con algunas diferencias claves.

La primera diferencia es que el `form_for` toma un "modelo" como el primer argumento. Este "model" sólo necesita ser un `struct`, no tiene que estar respaldado por una base de datos.

La segunda diferencia está en la etiqueta que llama al modelo directamente. Estas etiquetas, como `InputTag`, toma el nombre del atributo del modelo para el que deseas contruir un campo, luego toma un grupo opcional de opciones como segundo argumento.

{{< codetabs >}}
{{< tab "models/talk.go" >}}
```go
// models/talk.go
type Talk struct {
  ID            int          `json:"id" db:"id"`
  CreatedAt     time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
  UserID        int          `json:"user_id" db:"user_id"`
  Title         string       `json:"title" db:"title"`
  Description   nulls.String `json:"description" db:"description"`
  Notes         nulls.String `json:"notes" db:"notes"`
  ParentID      nulls.Int    `json:"parent_id" db:"parent_id"`
  Abstract      string       `json:"abstract" db:"abstract"`
  AudienceLevel string       `json:"audience_level" db:"audience_level"`
  IsPublic      nulls.Bool   `json:"is_public" db:"is_public"`
  TalkFormatID  int          `json:"talk_format_id" db:"talk_format_id"`
}
```

{{< /tab>}}
{{< tab "templates/talks/edit.html" >}}
```erb
<%= form_for( talk, {action: talkPath({id: 3}), method: "PUT"}) { %>
  <div class="row">
    <div class="col-md-12">
      <%= f.InputTag("Title") %>
    </div>
    <div class="col-md-6">
      <%= f.TextArea("Abstract", {hide_label: true}) %>
    </div>


    <div class="col-md-6">
      <%= f.SelectTag("TalkFormatID", {options: talk_formats}) %>
      <%= f.SelectTag("AudienceLevel", , {options: audience_levels}) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea("Description", {rows: 10}) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea("Notes", {rows: 10}) %>
    </div>
  </div>
<% } %>
```
{{< /tab>}}
{{< tab "OUTPUT">}}
```html
// OUTPUT
<form action="/talks/3" id="talk-form" method="POST">
  <input name="authenticity_token" type="hidden" value="cd998be98a99b452481c43fd3e4715e4e85333a45b982ac999">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <label>Title</label>
        <input class="form-control" id="talk-Title" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class="form-control" id="talk-Abstract" name="Abstract">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
      <label>TalkFormatID</label>
        <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <label>AudienceLevel</label>
        <select class=" form-control" id="talk-AudienceLevel" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Description</label>
        <textarea class=" form-control" id="talk-Description" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Notes</label>
        <textarea class=" form-control" id="talk-Notes" name="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## Etiqueta Select


Para construir tus etiquetas `<select>` dentro de los formularios `Tags` proporciona 3 formas convenientes de agregar tus opciones `<select>`: `form.SelectOptions`, `map[string]interface{}` o `[]string`, todas ellas pasando un campo `options` a las opciones de `form.SelectTag` como:


```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats}) %>
```

o

```erb
<%= f.SelectTag("TalkFormatID", {options: ["one", "two"]}) %>
```

La cual usará el mismo valor del atributo `value` y el cuerpo de la opción, o:

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}}) %>
```

La cual te permite definir el mapa de opciones dentro de la vista.

### Interfaz Selectable

Otra alternativa para las opciones de select es pasar una lista de estructuras que implementen la interfaz `form.Selectable`.

Que consiste en 2 métodos:

```go
//Selectable allows any struct to become an option in the select tag.
type Selectable interface {
  SelectValue() interface{}
  SelectLabel() string
}
```

Al implementar esta interfaz, tags llamará a `SelectValue` y `SelectLabel` para obtener la opción Value y Lavel del implementador.

### Selected

Tags agregará el atributo `selected` a la opcion que tiene el mismo valor al que se recibe en la opción `value` del `form.SelectTag`, asi no tienes que morar la opción que tienen el mismo valor que la seleccionada manualmente, p.e:

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}, value: 2}) %>
```

Produce:

```html
<div class="form-group">
  <label>TalkFormatID</label>
  <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
    <option value="1">one</option>
    <option value="2" selected>two</option>
  </select>
</div>
```

Y similarmente con el slice `form.SelectOptions`:

```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats, value: 2}) %>
```

## Etiqueta Checkbox

Tags proporciona una forma conveniente de construir un elemento HTML `input` con `type="checkbox"`:

```erb
<%= f.CheckboxTag("IsPublic") %>
```

Que produce:

```html
<div class="form-group">
  <label>
    <input class="" id="talk-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
    IsPublic
  </label>
</div>
```

Puedes cambiar facilmente el contenido del label con:

```erb
<%= f.CheckboxTag("IsPublic", {label: "Is the talk public?"}) %>
```

Que produce:

```html
<div class="form-group">
  <label>
    <input class="" id="post-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
     Is the Talk public?
  </label>
</div>
```

### Valores de Checkbox no marcado

Por defecto, cuando un checkbox no está "marcado" no se enviará ningun valor al servidor. A menudo, es útil enviar un valor indicando que el checkbox no está marcado. Esto se puede establecer pasando un valor `unchecked`.

```erb
<%= f.CheckboxTag("IsPublic", {unchecked: false}) %>
```

```html
<div class="form-group">
  <label>
    <input id="widget-IsPublic" name="IsPublic" type="checkbox" value="true">
    <input name="IsPublic" type="hidden" value="false"> IsPublic
  </label>
</div>
```

Cuando se envie el formulario, se enviará la etiqueta `hidden` y el servidor verá el valor `falso`.

## Manejando Errores

Tanto los ayudantes `form` como `form_for` tienen soporte para manejar los errores del paquete [`github.com/gobuffalo/validate`](https://github.com/gobuffalo/validate).

En una acción simplemente establece un valor de tipo `*validate.Errors` en el contexto como `errors` y los helpers de formulario lo recogerán y añadirán mensajes de error a las etiquetas de formulario correspondientes.

{{< codetabs >}}
{{< tab "actions/widgets.go" >}}
```go
// actions/widgets.go
func (v WidgetsResource) Create(c buffalo.Context) error {
  tx := c.Value("tx").(*pop.Connection)
  widget := &models.Widget{}
  if err := c.Bind(widget); err != nil {
    return err
  }
  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(widget)
  if err != nil {
    return errors.WithStack(err)
  }
  if verrs.HasAny() {
    c.Set("widget", widget)
    // Make the errors available inside the html template
    c.Set("errors", verrs)
    return c.Render(422, r.HTML("widgets/new.html"))
  }
  c.Flash().Add("success", "Widget was created successfully")
  return c.Redirect(302, "/widgets/%s", widget.ID)
}
```
{{< /tab >}}
{{< tab "templates/widgets/new.html" >}}
```erb
// templates/widgets/new.html
<%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= f.InputTag("Name") %>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="<%= widgetsPath() %>" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}

```html
// OUTPUT
<form action="/widgets" id="widget-form" method="POST">
  <input name="authenticity_token" type="hidden" value="AI0pb5YFBw2xU/EfcS6FaEOwTLWaGv58Y+w0ArfJoknfqu7l/j6tRLWybbcm+YZqXbBmi7f80l3Sf0WfnR7COA==">
  <div class="form-group has-error">
    <label>Widget</label>
    <input class=" form-control" id="widget-Widget" name="Widget" type="text" value="">
    <span class="help-block">Widget can not be blank.</span>
  </div>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## Uso de helpers de formulario que no son de Bootstrap

Los helpers por defecto, `form` y `form_for` generan formularios compatibles con Boostrap 3. Si esto no es para ti, puedes usar fácilmente la version no-Boostrap de estos helpers.

{{<note>}}
Requiere la versión de Plush `v3.4.8` o superior
{{</note>}}

{{< codetabs >}}
{{< tab "actions/render.go" >}}
```go
// actions/render.go
func init() {
  r = render.New(render.Options{
    // ...
    // Add template helpers here:
    Helpers: render.Helpers{
      "form":     plush.FormHelper,
      "form_for": plush.FormForHelper,
    },
    // ...
  })
}
```
{{< /tab >}}
{{< tab "templates/widgets/new.html" >}}
```erb
// templates/widgets/new.html
<%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= f.InputTag("Name") %>
  <%= f.InputTag("Body") %>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="<%= widgetsPath() %>" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
<% } %>
```
{{< /tab >}}
{{< tab "OUTPUT" >}}
```html
// OUTPUT
<form action="/widgets" id="widget-form" method="POST">
  <input name="authenticity_token" type="hidden" value="jN3nYOhCTqxZvmYnO9v1maso2VMs8fslj3rmKg1TS281W6JKpMd6Uezqp1dd3VBu2su41nKRBkd5AWDyCM4BzQ==">
  <input id="widget-Name" name="Name" type="text" value="">
  <input id="widget-Body" name="Body" type="text" value="">
  <button class="btn btn-success" role="submit">Save</button>
  <a href="/widgets" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
</form>
```
{{< /tab >}}
{{< /codetabs >}}


## Preguntas frecuentes

### ¿Cómo asignar un formulario a un modelo/estructura?

Consulta la página [Vinculación de Peticiones](/es/documentation/request_handling/bind) para mas información sobre la vinculación de peticiones.

### ¿Puedo cambiar el nombre de la variable `f` en mi plantilla?

Por defecto, el valor del formulario dentro del bloque tiene por nombre `f`, sin embargo, este lo puedes cambiar cuando creas el formulario y pasas la opción `var`.

```erb
<%= form({var: "xyz"}) { %>
  <%= xyz.InputTag({name: "Foo"}) %>
<% } %>
```

### Como puedo crear un formulario Multipart?

```erb
<%= form({multipart: true}) { %>
<% } %>
```

```html
<form enctype="multipart/form-data" method="POST">
</form>
```

### Puedo solo usar mi propio formulario (Sin usar el helper de formulario)?

Si, puedes crear tu propio formulario. Los formularios son proporcionados para qeu Buffalo genere tus recursos son simplemente un marcador de pocision para que puedas ponerte en marcha rápidamente. Sin embargo, es importante tener en cuenta que pedir a Buffalo que genere tus recursos, usando generadores suministrados, tambien generará las rutas relacionadas con los CRUD's. Esto es importante ya que la ruta asociada a la acción UPDATE hace uso del método PUT y no es un valor válido para un método de formulario HTML según el [HTML Standard](https://www.w3.org/TR/html5/forms.html#association-of-controls-and-forms). Dicho esto, debes asegurarte de estructurar tu formulario (para editar un recurso) para usar el método POST para conectar con el método HTTP, mientras usas una entrada oculta para indicar tu intención de hacer uso del método PUT del lado del servidor. Un ejemplo de esto sería el siguiente:

```html
<form method="POST" ...>
  <input type="hidden" name="_method" value="PUT" />
...
```


#### Cómo puedo manejar los tokens CSRF si utilizo mi propio formulario?

Si decides usar tus propios formularios, necesitarás una forma de proporcionarle al formulario el token de autenticidad. Hay dos formas de resolver este problema.

La primera forma es usar el `authenticity_token` directamente en el formulario, puesto que ya está en el contexto.

```html
<form method="POST" ...>
  <input name="authenticity_token" type="hidden" value="<%= authenticity_token %>">
</form>
```

Otra forma es escribir el helper que genere esa linea de codigo por ti.

```go
"csrf": func(ctx plush.HelperContext) (template.HTML, error) {
  tok, ok := ctx.Value("authenticity_token").(string)
  if !ok {
    return "", fmt.Errorf("expected CSRF token got %T", ctx.Value("authenticity_token"))
  }
  t := tags.New("input", tags.Options{
    "value": tok,
    "type":  "hidden",
    "name":  "authenticity_token",
  })
  return t.HTML(), nil
},
```

Ahora que has definido un helper para usarlo en tus plantillas puedes usar tu helper dentro de tu formulario con `<%= csrf() %>`. Así que tu formulario personalizado debería terminar viéndose así:

```erb
<form method="POST" ...>
  <%= csrf() %>
</form>
```
