## Exemple de génération de ressource

Dans cet exemple, Buffalo va générer le code nécessaire pour une ressource `widget` (en Go : `Widget`) qui possède les attributs suivants :

|                | Modèle        | Type Go                                                                   | Type en base de données  | Type dans le formulaire  |
|----------------|---------------|---------------------------------------------------------------------------|--------------------------|--------------------------|
| `title`        | `Title`       | `string`                                                                  | `varchar`                | `text`                   |
| `description`  | `Description` | [`nulls.String`](https://godoc.org/github.com/gobuffalo/pop/nulls#String) | `varchar (nullable)`     | `textarea`               |

```bash
$ buffalo generate resource widget title description:nulls.Text
```

<%= exampleDir("fr/docs/resources/_example/standard") %>

