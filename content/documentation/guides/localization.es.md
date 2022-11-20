---
name: Localización
seoDescription: i18n y l10n para las aplicaciones de Buffalo
seoKeywords: ["buffalo", "go", "golang", "i18n", "l10n", "internationalization", "localization"]
aliases:
  - /docs/localization
  - /es/docs/localization
---

# Localización

Traducir tu aplicación es una forma efectiva de **hacerla comprensible para muchas personas** al rededor del mundo!. Buffalo usa el proyecto [go-i18n](https://github.com/nicksnyder/go-i18n) para proporcionar {{< abbr title="Internacionalización" message="i18n">}} (adaptando el software para hacerlo traducible sin cambiar código) y {{< abbr title="Localización" message="l10n">}} (proporciona parablas de traducción y formatos específicos) en tu aplicación.

## Marquar las palabras traducibles

{{< note >}}
Este documento solo aplica cuando se usa [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render).
Consulta [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) para mas detalles sobre el paquete de plantillas.
{{< /note >}}

Las palabras traducibles deben tener una **marca específica** para permitir que el motor *l10* obtenga las traducciones.

En una plantilla de Plush, puedes usar el helper `t`:

```plain
<%= t("greetings") %>
```

### Contexto de traducción
{{< since "0.10.2" >}}

Puedes usar un contexto con el helper `t`, para proporcionar variables para la palabra a traducir:

```plain
<%= t("name-format", {name: "Mark"}) %>
```

### Manejando Plurales

{{< since "0.10.2" >}}

Puedes usar este helper como un segundo argumento numerico para manejar el caso singular/plural:

```plain
<%= t("messages-notification", notificationsLen) %>
```

Proporciona un contexto usando un tercer argumento:

```plain
<%= t("messages-notification", notificationsLen, ctx) %>
```

{{< note >}}
Se puede acceder al segundo argumento como "Count" en las palabras a traducir.
{{< /note >}}

## Provide Translations

Las traducciones se guardan en la carpeta `locales`, Por defecto, se guardan en el archivo `all.en-us.yaml` para las palabras en Ingles Americano.

Puedes proporcionar traducciones para otro lenguaje en un nuevo archivo `all.my-language-code.yaml`. Si deseas separar tus palabras en modulos locales, incluso puedes crear múltiples archivos. Por ejemplo: `users.en-us.yaml` para las cosas reacionadas con usuarios, y `all.en-us.yaml` para cosas globales.

El formato de localización que usa [go-i18n](https://github.com/nicksnyder/go-i18n) es el siguiente:

```yaml
- id: greetings
  translation: "Welcome to Buffalo (EN)"

- id: messages-notification
  translation:
    one: "You have {{.Count}} notification"
    other: "You have {{.Count}} notifications"
```

## Define un lenguaje por defecto

Para definir un lenguaje por defecto en tu aplicación, necesitas editar el archivo `app.go` dentro de la carpeta `actions`:

```go
// Setup and use translations:
var err error
if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
  app.Stop(err)
}
app.Use(T.Middleware())
```

Cambiar `"en-US"` por otro código de languaje cambiará el lenguaje por defecto.

## Vistas Localizadas

{{< since "0.10.2" >}}

A veces, tienes que **traducir una página completa** y marcar cada parte de la página toma mucho tiempo. En algunos otros casos, desearás localizar la página de una manera diferente para una configuración regional específica. Las vistas localizadas son una forma complementaria de manejar tus traducciones.

Las vistas localizadas están **incluidas en el middleware i18n**, por lo que no neceitas configurar nada más para usarlas.

### Crear versiones con sufijos de las plantillas

Primero, crea una versión para la configuración regional por defecto, sin un sufijo de lenguaje:

**page.html**:
```html
<p>This is my default language page.</p>
```

Luego, crea una version con sufijo para cada lenguaje que deseeas admitir:

**page.en-us.html**:
```html
<p>This is my en-US version.</p>
```

**page.fr-fr.html**:
```html
<p>This is my fr-FR version.</p>
```

El middleware detectará el lenguaje del usuario y seleccionará pla plantilla correcta para ti! Tambien funciona con usuarios invitaods, usando el header HTTP `Accept-Language`.

## Usando i18n en las acciones

Deberás usar las funciones de i18n en acciones, por ejemplo, para traducir mensajes flash. Está la forma de usarlo:

``` go
func Login(c buffalo.Context) error {
  // [...]
  // Set a translated flash message
  c.Flash().Add("success", T.Translate(c, "users.login-success"))
  return c.Redirect(303, "/users/signin")
}
```

`T.Translate` toma `buffalo.Context` como primer argumento, luego los siguientes argumentos son los mismos que los del helper `t` (`t` llama a `T.Translate` con el contexto, detrás de escena).

## Actualizar contexto de traducción

{{< since "0.12.0" >}}

Si proporcionas versiones traducidas de tu aplicación, probablemente tendrás una función de cambio de idioma. De esta forma, los usuarios pueden escoger el idioma correcto.
Buffalo no puede detectar cuándo cambia el idioma en una acción, ya que extraerá los idiomas del usuario una vez por solicitud. Después tendrás que redirigir a otra página para ver los cambios. Pero incluso con ese truco, si usas un mensaje flash dentro de la acción, el idioma utilizado será el anterior.

Para solucionar este problema. puedes usar el método `T.Refresh` y actualizar el idioma utilizado por las traducciones, dentro de una ación.

```go
func SwitchLanguage(c buffalo.Context) error {
  f := struct {
    Language string `form:"lang"`
    URL      string `form:"url"`
  }{}
  if err := c.Bind(&f); err != nil {
    return errors.WithStack(err)
  }

  // Set new current language using a cookie, for instance
  cookie := http.Cookie{
    Name:   "lang",
    Value:  f.Language,
    MaxAge: int((time.Hour * 24 * 265).Seconds()),
    Path:   "/",
  }
  http.SetCookie(c.Response(), &cookie)

  // Update language for the flash message
  T.Refresh(c, f.Language)

  c.Flash().Add("success", T.Translate(c, "users.language-changed", f))

  return c.Redirect(302, f.URL)
}
```

## Personalizar Nombres Generados

{{< since "0.10.2" >}}

Muchos generadores de Buffalo usan [gobuffalo/flect](https://github.com/gobuffalo/flect) para generar una versión normalizada de un nombre. Por ejemplo, cuando desees generar un nuevo modelo, el nombre que le das a la línea de comando se normaliza en plural, mayúsculas, etc.

A veces, las reglas utilizadas por **flect** no son correctas (en este caso, sientete libre de abrir un PR en el repo!). A veces, una regla no es correcta para tu caso de uso, pero sigue siendo correcta en una regla general. En este caso, puedes proporcionar reglas personalizadas usando el archivo `inflections.json` en la raíz de tu proyecto.

**inflections.json:**
```json
{
  "singular form": "plural form"
}
```

## Recursos Relacionados

* [Traducir una aplicación de Buffalo](https://blog.gobuffalo.io/translating-a-buffalo-app-1b4f32e6cb57) - An article about using Buffalo i18n tools.
