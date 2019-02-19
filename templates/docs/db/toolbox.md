<% seoDescription("Soda CLI") %>
<% seoKeywords(["buffalo", "go", "golang", "database", "utils", "pop", "toolbox", "CLI", "soda"]) %>

<%= h1("Soda CLI") %>

Pop helps you to manage database connections, but it also provides `soda`, a small CLI toolbox to manage your database. It can help you to create a new database, drop existing ones, and so on.

<%= note() { %>
**Note for Buffalo users**: `soda` commands are embedded into the `buffalo` command, behind the `pop` namespace. So every time you want to use a command from `soda`, just execute `buffalo pop` instead. You don't need to install `soda` CLI.
<% } %>

## Installing CLI Support

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

## Creating Databases

Once the `database.yml` has been configured with the appropriate settings, and the database server is running, Soda can create all of the databases in the `database.yml` file with a simple command:

```bash
$ soda create -a
```

You can also create just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ soda create -e test
```

## Dropping Databases

Soda can drop all of your databases, should you want to, with one command:

```bash
$ soda drop -a
```

You can also drop just one of the configured databases by using the `-e` flag and the name of the database:

```bash
$ soda drop -e test
```

## Generating Models

The available types for use in the generator are:

| Base type             | Nullable        | Slice/Array |
|-----------------------|:---------------:|------------:|
|int                    |nulls.Int        |slices.Int   |
|int32                  |nulls.Int32      | ------      |
|int64                  |nulls.Int64      | ------      |
|uint32                 |nulls.UInt32     | ------      |
|float32                |nulls.Float32    | ------      |
|float, float64         |nulls.Float64    |slices.Float |
|bool                   |nulls.Bool       | ------      |
|[]byte                 |nulls.ByteSlice  | ------      |
|string                 |nulls.String     |slices.String|
|uuid.UUID              |nulls.UUID       |slices.UUID  |
|time.Time              |nulls.Time       | ------      |
|map[string]interface{} | ---------       |slices.Map   |


For example:

```bash
soda generate model user id:uuid name:string email:string password:string
soda generate model post user_id:uuid title:string text:blob tags:slices.String
```
