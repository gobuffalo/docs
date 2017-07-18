# Test Suites

Buffalo uses the `github.com/gobuffalo/suite` package to create test suites.

When running a test that is part of the test suite, the following is available to
the test:

* The application, `as.App`.
* The database, `as.DB` (if using Pop).
* The `github.com/stretchr/testify/require` test assertions.
* The `github.com/markbates/willie HTTP` testing library.

## Test Example

<%= code("text") { %>
func (as *ActionSuite) Test_PostsResource_Create() {
// setup a Post model
p := &models.Post{Title: "My Post", Body: "The Body"} // make a POST /posts request
res := as.HTML("/posts").Post(p)
// assert that the response status code was 302 as.Equal(302, res.Code)
   // retreive the first Post from the database
   err := as.DB.First(p)
   as.NoError(err)
   as.NotZero(p.ID)
   // assert the Post title was saved correctly
   as.Equal("My Post", p.Title)
   // assert the redirect was sent to the place expected
   as.Equal(fmt.Sprintf("/posts/%s", p.ID), res.Location())
 }
 <% } %>