package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	err := c.Render(200, r.HTML("docs/getting-started.md"))
	if err != nil {
		return c.Error(404, err)
	}
	return nil
}
