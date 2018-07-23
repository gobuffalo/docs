<% seoDescription("Soda CLI") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "utils", "pop", "toolbox", "CLI", "soda"]) %>

<%= h1("Soda CLI") %>

Pop helps you to manage database connections, but it also provides `soda`, a small CLI toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `db` namespace. So everytime you want to use a command from `soda`, just execute `buffalo db` instead.
<% } %>

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