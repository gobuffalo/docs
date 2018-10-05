<%= title("Example Resource Generation") %>

In this example Buffalo will generate the code needed to CRUD a resource named `widget` (Go: `Widget`) that has the following attributes:

* `title` - Model Attribute: `Title`; Go Type: `string`; DB Type: `varchar`; Form Type: `text`
* `description` - Model Attribute: `Description`; Go Type [`nulls.String`](https://godoc.org/github.com/gobuffalo/pop/nulls#String); DB Type: `varchar (nullable)`; Form Type: `textarea`

```bash
$ buffalo generate resource widget title description:nulls.Text
```

<%= exampleDir("docs/resources/_example/standard") %>

### Video Presentation

<%= vimeo("212302823") %>
