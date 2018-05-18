<% seoDescription("Database toolbox") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "ORM", "pop", "toolbox"]) %>

<%= h1("Database toolbox") %>

Pop helps you to manage database connections, but it also provides `soda`, a small toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

Soda commands are embedded into the `buffalo` command, in the `db` namespace. So everytime you want to use a command from `soda`, just execute `buffalo db` instead.

<%= title("Creating Databases") %>

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Buffalo can create all of the databases in the `database.yml` file with a simple command:

```bash
$ buffalo db create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ buffalo db create -e test
```

<%= title("Dropping Databases") %>

Buffalo can drop all of your databases, should you want to, with one command:

```bash
$ buffalo db drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ buffalo db drop -e test
```