{{#panel title="Resources"}}

```text
$ buffalo g resource --help

Generates a new actions/resource file

Usage:
  buffalo generate resource [name] [flags]

Aliases:
  resource, r
```

```text
$ buffalo g resource users

Buffalo version 0.4.6

--> actions/users.go
--> actions/users_test.go
--> goimports -w .
```

<table width="100%">
<tr>
<td width="50%" valign="top">
```go
// actions/users.go
package actions

import (
	"errors"

	"github.com/markbates/buffalo"
)

type UsersResource struct{}

// List default implementation. Returns a 404
func (v *UsersResource) List(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// Show default implementation. Returns a 404
func (v *UsersResource) Show(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// New default implementation. Returns a 404
func (v *UsersResource) New(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// Create default implementation. Returns a 404
func (v *UsersResource) Create(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// Edit default implementation. Returns a 404
func (v *UsersResource) Edit(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// Update default implementation. Returns a 404
func (v *UsersResource) Update(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}

// Destroy default implementation. Returns a 404
func (v *UsersResource) Destroy(c buffalo.Context) error {
	return c.Error(404, errors.New("resource not implemented"))
}
```
</td>
<td width="50%" valign="top">
```go
// actions/users_test.go

package actions_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UsersResource_List(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_Show(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_New(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_Create(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_Edit(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_Update(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}

func Test_UsersResource_Destroy(t *testing.T) {
	r := require.New(t)
	r.Fail("Not Implemented!")
}
```
</td>
</tr>
</table>
{{/panel}}
