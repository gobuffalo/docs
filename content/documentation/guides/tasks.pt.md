---
name: Tasks
weight: 14
seoDescription: Tasks
seoKeywords: ["buffalo", "go", "golang", "tasks", "scripts", "grift"]
aliases:
  - /docs/tasks
  - /pt/docs/tasks
---
# Tasks

Tasks are small scripts that are often needed when writing an application. These tasks might be along the lines of seeding a database, parsing a log file, or even a release script. Buffalo uses the [grift](https://github.com/markbates/grift) package to make writing these tasks simple.

{{< vimeo 213096302>}}

## Writing Tasks

Tasks must all be in the `grifts` package. A simple task would look like following:

```go
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
```

## Tasks Generator

```bash
$ buffalo g task foo:bar

--> grifts/bar.go
```

```go
// grifts/bar.go
package grifts

import (
  . "github.com/markbates/grift/grift"
)

var _ = Namespace("foo", func() {

  Desc("bar", "TODO")
  Add("bar", func(c *Context) error {
    return nil
  })

})
```


## Listing Available Tasks

```bash
$ buffalo task list

Available grifts
================
buffalo task db:seed       # Seeds a database
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

After a binary has been [built](/documentation/deploy/packing), the tasks can be run with the `task` subcommand:

```bash
$ myapp task hello
```
