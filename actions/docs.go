package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

func Docs(c buffalo.Context) error {
	for _, ext := range []string{"md", "html"} {
		f := fmt.Sprintf("docs/%s.%s", c.Param("name"), ext)
		if r.TemplatesBox.Has(f) {
			err := c.Render(200, r.HTML(f))
			if err != nil {
				return c.Error(404, err)
			}
			return nil
		}
	}
	return c.Error(404, errors.Errorf("could not find %s", c.Param("name")))
}
