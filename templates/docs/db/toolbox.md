<% seoDescription("Soda CLI") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "utils", "pop", "toolbox", "CLI", "soda"]) %>

<%= h1("Soda CLI") %>

Pop helps you to manage database connections, but it also provides `soda`, a small CLI toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead. You don't need to install `soda` CLI.
<% } %>

<%= title("Installing CLI Support") %>

**Without** sqlite 3 support:

```bash
$ go get github.com/gobuffalo/pop/...
$ go install github.com/gobuffalo/pop/soda
```

**With** sqlite 3 support (requires GCC or equivalent C compiler):

```bash
$ go get -u -v -tags sqlite github.com/gobuffalo/pop/...
$ go install -tags sqlite github.com/gobuffalo/pop/soda
```

If you're not building your code with `buffalo build`, you'll also have to pass `-tags sqlite` to `go build` when building your program.

<%= title("Creating Databases") %>

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Soda can create all of the databases in the `database.yml` file with a simple command:

```bash
$ soda create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ soda create -e test
```

<%= title("Dropping Databases") %>

Soda can drop all of your databases, should you want to, with one command:

```bash
$ soda drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ soda drop -e test
```