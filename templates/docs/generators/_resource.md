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

import "github.com/gobuffalo/buffalo"

type UsersResource struct {
	buffalo.Resource
}

func init() {
	App().Resource("/users", &UsersResource{&buffalo.BaseResource{}})
}

// List default implementation.
func (v *UsersResource) List(c buffalo.Context) error {
	return c.Render(200, r.String("Users#List"))
}

// Show default implementation.
func (v *UsersResource) Show(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Show"))
}

// New default implementation.
func (v *UsersResource) New(c buffalo.Context) error {
	return c.Render(200, r.String("Users#New"))
}

// Create default implementation.
func (v *UsersResource) Create(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Create"))
}

// Edit default implementation.
func (v *UsersResource) Edit(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Edit"))
}

// Update default implementation.
func (v *UsersResource) Update(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Update"))
}

// Destroy default implementation.
func (v *UsersResource) Destroy(c buffalo.Context) error {
	return c.Render(200, r.String("Users#Destroy"))
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
