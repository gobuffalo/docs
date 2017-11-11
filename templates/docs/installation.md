<%= h1("Installation") %>

<%= title("Requirements") %>

Before installing make sure you have the required dependencies installed:

* [Guide to Setting a Go Developer Environment](http://gopherguides.com/before-you-come-to-class)
* [Go](https://golang.org) version `<%= goMinVersion %>` or greater.

##### Frontend requirements

* [node](https://github.com/nodejs/node) and [npm](https://github.com/npm/npm) for the [asset pipeline](/docs/assets) built upon [webpack](https://github.com/webpack/webpack).

##### Database specific requirements

* **SQLite 3**: GCC, or equivalent C compiler for [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3).

<%= title("Basic Installation") %>

Buffalo provides **two major components**:
* The `buffalo` tool, a powerful toolbox to help you develop in a fast and efficient way.
* The buffalo framework, a collection of pieces to construct your app.

```bash
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
```

<%= note() { %>
These instructions can also be used for upgrading to a newer version of Buffalo.
<% } %>
