<% seoDescription("How to install the Buffalo framework") %>
<% seoKeywords(["buffalo", "go", "golang", "installation", "framework", "web"]) %>

<%= h1("Installation") %>

Buffalo provides **two major components**:
* The `buffalo` tool, a powerful toolbox to help you develop in a fast and efficient way.
* The buffalo framework, a collection of pieces to construct your app.

<%= title("Requirements") %>

Before installing make sure you have the required dependencies installed:

* [Guide to Setting a Go Developer Environment](http://gopherguides.com/before-you-come-to-class)
* [Go](https://golang.org) version `<%= goMinVersion %>` or greater.

##### Frontend Requirements

* [node](https://github.com/nodejs/node) and [npm](https://github.com/npm/npm) for the [asset pipeline](/docs/assets) built upon [webpack](https://github.com/webpack/webpack).

##### Database Specific Requirements

* **SQLite 3**: GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

<%= title("Installation **with** SQLite3 Support", {name: "with-SQLite3", title: "Installation with SQLite3 Support"}) %>

**SQLite 3** requires a GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) to compile. You **must** have a GCC installed **first** before installing Buffalo.

To install Buffalo, ensure your `GOPATH` is defined, then:

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

**Windows Users**: Follow the installation guide at [https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3](https://blog.gobuffalo.io/install-buffalo-on-windows-10-e08b3aa304a3) to install a GCC for Windows 10.

<%= note() { %>
These instructions can also be used for upgrading to a newer version of Buffalo.
<% } %>

<%= title("Installation **without** SQLite3 Support", {name: "without-SQLite3", title: "Installation without SQLite3 Support"}) %>

```bash
$ go get -u -v -tags nosqlite github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
These instructions can also be used for upgrading to a newer version of Buffalo.
<% } %>
