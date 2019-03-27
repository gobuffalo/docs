<% seoDescription("Tasks") %>
<% seoKeywords(["buffalo", "go", "golang", "tasks", "scripts", "grift"]) %>

<% contentFor("docsLanding") { %>
  <%= partial("docs/landing.html", {title: "Tasks"}) %>
<% } %>

Tasks are small scripts that are often needed when writing an application. These tasks might be along the lines of seeding a database, parsing a log file, or even a release script. Buffalo uses the [grift](https://github.com/markbates/grift) package to make writing these tasks simple.

<%= vimeo("213096302") %>

## Writing Tasks

Tasks must all be in the `grifts` package. A simple task would look like following:

```go
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
```

<%= partial("docs/generators/tasks.md") %>

## Listing Available Tasks

```bash
$ buffalo task list

Available grifts
================
buffalo task middleware    # Prints out your middleware stack
buffalo task routes        # Print out all defined routes
buffalo task secret        # Generate a cryptographically secure secret key
```

## Running Tasks

### Development

Tasks can be run in development using the `buffalo task` command.

```bash
$ buffalo task hello
```

### From a Built Binary

After a binary has been [built](/docs/building), the tasks can be run with the `task` subcommand:

```bash
$ myapp task hello
```
