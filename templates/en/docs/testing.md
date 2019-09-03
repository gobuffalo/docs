# Testing

## Running Tests

The Buffalo test runner will make sure your test environment is correct, and run your tests.

For example, if using Pop (database), it will first try to setup your test database by using the schema from the development database. If that doesn't exist (CI, for example), it will run the migrations against the test database.

The test runner will also make sure to ignore the dreaded `vendor` directory.

```bash
$ buffalo test

dropped database authrecipe_test
created database authrecipe_test
dumped schema for authrecipe_development
loaded schema for authrecipe_test
go test -p 1 github.com/gobuffalo/authrecipe github.com/gobuffalo/authrecipe/actions github.com/gobuffalo/authrecipe/grifts github.com/gobuffalo/authrecipe/models
?   	github.com/gobuffalo/authrecipe	[no test files]
ok  	github.com/gobuffalo/authrecipe/actions	0.640s
?   	github.com/gobuffalo/authrecipe/grifts	[no test files]
ok  	github.com/gobuffalo/authrecipe/models	0.327s
```

### Execute a single test
<%= sinceVersion("0.10.2") %>

Debugging a specific test is a difficult task, if you must execute all existing tests. You can use the `-m` flag to execute a single test method:

```bash
$ buffalo test -m "FooMethod"
```

This will iterate through all packages and run any test that matches `FooMethod` in any package. 

<%= sinceVersion("0.14.10") %>
To limit to one package, specify the package name:

```bash
$ buffalo test models -m "FooMethod"
```

## Test Suites

Buffalo uses the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package to create test suites.

When running a test that is part of the test suite, the following is available to
the test:

* The application, `as.App`.
* The database, `as.DB` (if using Pop).
* The session, as `as.Session`.
* The [`github.com/stretchr/testify/require`](https://github.com/stretchr/testify) test assertions.
* The [`github.com/markbates/willie`](https://github.com/markbates/willie) HTTP testing library.

## Test Example

```go
func (as *ActionSuite) Test_WidgetsResource_Create() {
  // setup a Widget model
  w := &models.Widget{Name: "My Widget"} // make a POST /widgets request
  res := as.HTML("/widgets").Post(w)
  // assert that the response status code was 302 as.Equal(302, res.Code)
  // retrieve the first Widget from the database
  err := as.DB.First(w)
  as.NoError(err)
  as.NotZero(w.ID)
  // assert the Widget title was saved correctly
  as.Equal("My Widget", w.Name)
  // assert the redirect was sent to the place expected
  as.Equal(fmt.Sprintf("/widgets/%s", w.ID), res.Location())
}
```

## Fixtures

<%= sinceVersion("0.12.0") %>

Often it is useful to load sample data into the database at the start of a test. For example, you need to have a user in the database to log a person into the application, or you need some data in the database to test that a route renders responses correctly. Fixtures help us solve these problems easily.

Fixtures are `toml` files that are placed in your `fixtures` directory. You can have as many fixture files as you like and they can also be named anyway that you like.

For example, if we have a simple `Widget` model that renders to a page that looks like this:

```go
type Widget struct {
  ID          uuid.UUID
  CreatedAt   time.Time
  UpdatedAt   time.Time
  Name        string
  Description string
}
```

We can create a fixture file like this:

```toml
[[scenario]]
name = "lots of widgets"

  [[scenario.table]]
    name = "widgets"

    [[scenario.table.row]]
      id = "\<%= uuidNamed("widget") %>"
      name = "This is widget #1"
      description = "some widget body #1"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"

    [[scenario.table.row]]
      id = "\<%= uuid() %>"
      name = "This is widget #2"
      description = "some widget body #2"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"
```

When we run our suite, these two records will be created in our test database and we can then test against these records. 
All you need to do to load the fixture is to reference it by its name with `ActionSuite.LoadFixture`.

```go
func (as *ActionSuite) Test_WidgetsResource_List() {
  as.LoadFixture("lots of widgets")
  res := as.HTML("/widgets").Get()

  as.Equal(200, res.Code)
  body := res.Body.String()
  as.Contains(body, "This is widget #1")
  as.Contains(body, "This is widget #2")
}
```

You can find more detailed information on fixtures in the [gobuffalo/suite repository README](https://github.com/gobuffalo/suite#fixtures-test-data).

## Accessing the Session

Being able to manipulate the session for testing is very important. Test suites in Buffalo give you access to a testing session that you can use.

See [https://github.com/gobuffalo/authrecipe](https://github.com/gobuffalo/authrecipe) for a more in-depth example.

```go
func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
  // get a user from the DB

  // set the user ID onto the session
  as.Session.Set("current_user_id", user.ID)

  res := as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged in"
  as.Contains(res.Body.String(), "Sign Out")

  // clear the session
  as.Session.Clear()
  res = as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged out"
  as.Contains(res.Body.String(), "Sign In")
}
```

## Coverage Reports

<%= note() { %>
The following feature requires the use of **Go 1.10** or a more recent version.
Go cover does not support the `./...` operator in older versions, and trying to use it will generate an error.
<% } %>
It is possible to generate test coverage reports with buffalo by specifying the `-coverprofile` flag as follows:

```bash
$ buffalo test -coverprofile=c.out ./...
created database authrecipe_test
loaded schema for authrecipe_test
INFO[0010] go test -p 1 -coverprofile=c.out ./...
?       github.com/gobuffalo/authrecipe [no test files]
ok      github.com/gobuffalo/authrecipe/actions 2.770s  coverage: 76.9% of statements
?       github.com/gobuffalo/authrecipe/grifts  [no test files]
ok      github.com/gobuffalo/authrecipe/models  2.609s  coverage: 71.4% of statements

```
