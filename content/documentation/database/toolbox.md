---
seoDescription: "Soda CLI"
seoKeywords: ["buffalo", "go", "golang", "database", "utils", "pop", "toolbox", "CLI", "soda"]
name: Soda CLI
--- 

# Soda CLI

Pop helps you to manage database connections, but it also provides `soda`, a small CLI toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

{{< note >}}
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead. You don't need to install `soda` CLI.
{{< /note >}}

## Installing CLI Support

### From a release archive

Pre-compiled archives contain Soda **with SQLite support**.

Download the appropriate version for your platform from [Pop releases](https://github.com/gobuffalo/pop/releases).

Place it somewhere in your `PATH`, and ensure the `soda` binary is executable.

### Homebrew (macOS)

```console
$ brew install gobuffalo/tap/pop
```

### From source

For go version 1.16 and later,

**Without** sqlite 3 support:

```console
$ go install github.com/gobuffalo/pop/v6/soda@latest
```

**With** sqlite 3 support (requires GCC or equivalent C compiler):

```console
$ go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

If you're not building your code with `buffalo build`, you'll also have to pass `-tags sqlite` to `go build` when building your program.

## Creating Databases

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Soda can create all of the databases in the `database.yml` file with a simple command:

```console
$ soda create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

```console
$ soda create -e test
```

## Dropping Databases

Soda can drop all of your databases, should you want to, with one command:

```console
$ soda drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

```console
$ soda drop -e test
```
