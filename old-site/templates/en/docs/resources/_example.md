## Example Resource Generation

In this example Buffalo will generate the code needed to CRUD a resource named `widget` (Go: `Widget`) that has the following attributes:

|                | Model Attribute | Go Type                                                                   | DB type                  | Form Type                |
|----------------|-----------------|---------------------------------------------------------------------------|--------------------------|--------------------------|
| `title`        | `Title`         | `string`                                                                  | `varchar`                | `text`                   |
| `description`  | `Description`   | [`nulls.String`](https://godoc.org/github.com/gobuffalo/pop/nulls#String) | `varchar (nullable)`     | `textarea`               |

```bash
$ buffalo generate resource widget title description:nulls.Text
```

<%= exampleDir("en/docs/resources/_example/standard") %>

