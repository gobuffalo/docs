package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

const docsRepoBase = "https://github.com/gobuffalo/gobuffalo/edit/master/templates/"

// Docs serves documentation pages
func Docs(c buffalo.Context) error {
	for _, ext := range []string{"md", "html"} {
		f := fmt.Sprintf("docs/%s.%s", c.Param("name"), ext)
		if r.TemplatesBox.Has(f) {
			c.Set("sourceRoot", docsRepoBase)
			return c.Render(200, r.HTML(f, "docs-layout.html"))
		}
	}
	return c.Error(404, errors.Errorf("could not find %s", c.Param("name")))
}
