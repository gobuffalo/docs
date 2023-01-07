---
name: Template Engines
seoDescription: Template Engines for Buffalo application
seoKeywords: ["buffalo", "go", "golang", "template engine"]
aliases:
  - /docs/template-engines
  - /pt/docs/template-engines
---

# Template Engines

## Map Template Engines by Extension
{{< since "0.10.0" >}}

Previously you were able to define a single implementation of [`render.TemplateEngine`](https://godoc.org/github.com/gobuffalo/buffalo/render#TemplateEngine) per [`render.Engine`](https://godoc.org/github.com/gobuffalo/buffalo/render#Engine). This has been removed in favor of a map of `render.TemplateEngine`. Now you can map a file extension to an implementation of `render.TemplateEngine`. This means, not only, can you now use multiple template engines in one application, but you can also chain them together.

For example, if the file was `foo.tmpl.html` it would, by default, first be processed as a Go template, then that result would be sent to the Plush engine.

Here is a list of default implementations:

* `.html` - processed as a Plush template, unchanged from previous releases.
* `.md` - processed first as Markdown, then as a Plush template, unchanged from previous releases.
* `.tmpl` - processed as a Go template.
* `.js` - processed as a Plush template.

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
