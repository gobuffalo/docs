---
name: Motores de Plantillas
seoDescription: Motores de plantillas para aplicaciones de Buffalo
seoKeywords: ["buffalo", "go", "golang", "template engine"]
aliases:
  - /docs/template-engines
  - /es/docs/template-engines
---

# Motores de Plantillas

## Mapeo de motores de plantillas por extensión

{{< since "0.10.0" >}}

Anteriormente, podías definir una única implementacion de [`render.TemplateEngine`](https://godoc.org/github.com/gobuffalo/buffalo/render#TemplateEngine) por [`render.Engine`](https://godoc.org/github.com/gobuffalo/buffalo/render#Engine). Esti se eliminó en favor de un mapa de tipo `render.TemplateEngine`. Ahora puedes mapear una extensión de archivo a una implementación de `render.TemplateEngine`. Esto significa, no solo, que puedes usar multiples motores plantillas en una aplicacion, sino que puedes encadenarlos.


Por ejemplo, si el archivo fuera `foo.tmpl.html`, por defecto, se procesaría como una plantilla de Go y luego el resultado se enviará al motor de Plush.

Esta es una lista de implementaciones predeterminadas:

* `.html` - Procesados como una plantilla Plush, sin cambios con respecto a versiones anteriores.
* `.md` - Procesado como un Markdown, despues como una plantilla de Plush, sin cambios con respecto a versiones anteriores.
* `.tmpl` - Procesado como una plantilla de Go.
* `.js` - Procesado como una plantilla de Plush.

```go
func init() {
  r = render.New(render.Options{
    // ...
    TemplateEngines: map[string]render.TemplateEngine{
      ".tmpl": GoTemplateEngine,
    },
    // ...
  })
}

func GoTemplateEngine(input string, data map[string]interface{}, helpers map[string]interface{}) (string, error) {
  // since go templates don't have the concept of an optional map argument like Plush does
  // add this "null" map so it can be used in templates like this:
  // {{ partial "flash.html" .nilOpts }}
  data["nilOpts"] = map[string]interface{}{}

  t := template.New(input)
  if helpers != nil {
    t = t.Funcs(helpers)
  }

  t, err := t.Parse(input)
  if err != nil {
    return "", err
  }

  bb := &bytes.Buffer{}
  err = t.Execute(bb, data)
  return bb.String(), err
}
```
