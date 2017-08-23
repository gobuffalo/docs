# Testing

<%= title("Running Tests") %>

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

<%= title("Test Suites") %>

Buffalo uses the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package to create test suites.

When running a test that is part of the test suite, the following is available to
the test:

* The application, `as.App`.
* The database, `as.DB` (if using Pop).
* The session, as `as.Session`.
* The [`github.com/stretchr/testify/require`](https://github.com/stretchr/testify) test assertions.
* The [`github.com/markbates/willie`](https://github.com/markbates/willie) HTTP testing library.

<%= title("Test Example") %>

```go
func (as *ActionSuite) Test_WidgetsResource_Create() {
  // setup a Widget model
  w := &models.Widget{Name: "My Widget"} // make a POST /widgets request
  res := as.HTML("/widgets").Post(w)
  // assert that the response status code was 302 as.Equal(302, res.Code)
  // retreive the first Widget from the database
  err := as.DB.First(w)
  as.NoError(err)
  as.NotZero(w.ID)
  // assert the Widget title was saved correctly
  as.Equal("My Widget", w.Name)
  // assert the redirect was sent to the place expected
  as.Equal(fmt.Sprintf("/widgets/%s", w.ID), res.Location())
}
```

<%= title("Accessing the Session") %>

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
