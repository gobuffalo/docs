# Moteurs de templates

<%= title("Configurer un moteur de templates par extension de fichier") %>
<%= sinceVersion("0.10.0") %>

Avant la version 0.10.0, vous n'étiez capable de définir qu'une seule implémentation du [`render.TemplateEngine`](https://godoc.org/github.com/gobuffalo/buffalo/render#TemplateEngine) par [`render.Engine`](https://godoc.org/github.com/gobuffalo/buffalo/render#Engine). Ce n'est plus le cas grâce au nouveau tableau de `render.TemplateEngine`. Vous pouvez désormais faire correspondre l'extension d'un fichier avec une implémentation de `render.TemplateEngine`. Cela signifie que vous pouvez non-seulement utiliser plusieurs moteurs de templates dans une seule application, mais que vous pouvez également des utiliser en chaîne.

Par exemple, si un fichier s'appelle `foo.tmpl.html`, il sera par défaut traité comme un template Go, puis le résultat sera envoyé au moteur de templates Plush.

Voici la liste des implémentations par défaut :

* `.html` - traité comme un template Plush, inchangé depuis les versions précédentes.
* `.md` - traité d'abord comme du Markdown, puis comme un template Plush, inchangé depuis les versions précédentes.
* `.tmpl` - traité comme un template Go.
* `.js` - traité comme un template Plush.

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
  // comme les templates Go n'ont pas de concept de map optionnelle d'arguments comme Plush,
  // on ajoute cette map "null" pour qu'elle soit utilisée dans les templates de la façon suivante :
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
