# Test Suites

Buffalo uses the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package to create test suites.

When running a test that is part of the test suite, the following is available to
the test:

* The application, `as.App`.
* The database, `as.DB` (if using Pop).
* The [`github.com/stretchr/testify/require`](https://github.com/stretchr/testify) test assertions.
* The [`github.com/markbates/willie`](https://github.com/markbates/willie) HTTP testing library.

<%= title("Test Example") %>

<%= code("go") { %>
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
<% } %>
