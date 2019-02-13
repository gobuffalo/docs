package actions

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

const docsRepoBase = "https://github.com/gobuffalo/gobuffalo/edit/master/templates/"

// Docs serves documentation pages
func Docs(c buffalo.Context) error {
	name := strings.TrimRight(c.Param("name"), "/")
	for _, ext := range []string{"md", "html"} {
		f := fmt.Sprintf("docs/%s.%s", name, ext)
		if r.TemplatesBox.Has(f) {
			c.Set("sourceRoot", docsRepoBase)
			l := "docs-layout.html"
			if c.Param("_indexing") == "true" {
				l = "search-layout.html"
			}
			return c.Render(200, r.HTML(f, l))
		}
	}
	return c.Error(404, errors.Errorf("could not find %s", name))
}
