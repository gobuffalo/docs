package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("docs/overview.html"))
}

func ChangeLanguage(c buffalo.Context) error {
	f := struct {
		Language string `form:"language"`
		URL      string `form:"url"`
	}{}
	if err := c.Bind(&f); err != nil {
		return errors.WithStack(err)
	}
	c.Cookies().Set("lang", f.Language, (time.Hour * 24 * 265))

	return c.Redirect(302, f.URL)
}
