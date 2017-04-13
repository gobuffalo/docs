# Tasks


<%= vimeo("213096302") %>

<%= title("What are Tasks?") %>

Tasks are small scripts that are often needed when writing an application. These tasks might be along the lines of seeding a database, parsing a log file, or even a release script. Buffalo uses the [grift](https://github.com/markbates/grift) package to make writing these tasks simple.

<%= title("Writing Tasks") %>

Tasks must all be in the `grifts` package. A simple task would look like following:

<%= code("go") { %>
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
<% } %>

<%= title("Listing Available Tasks") %>

<div class="code-tabs">
<%= code("text") { %>
$ buffalo task list
<% } %>
<%= code("text", {file: "output"}) { %>
buffalo task db:seed
buffalo task db:seed:pets
buffalo task db:seed:users
buffalo task hello
buffalo task routes
<% } %>
</div>

<%= title("Running Tasks") %>

### Development

Tasks can be run in development using the `buffalo task` command.

<%= code("text") { %>
$ buffalo task hello
<% } %>

### From a Built Binary

After a binary has been [built](/docs/building) the tasks can be run with the `task` subcommand:

<%= code("text") { %>
$ myapp task hello
<% } %>
