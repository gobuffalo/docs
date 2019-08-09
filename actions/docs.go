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
	languages, ok := c.Value("languages").([]string)
	if !ok {
		return errors.New("could not get user languages")
	}

	// try to get the docs for language in URL, then fallback to default one (en)
	name := strings.TrimRight(c.Param("name"), "/")
	for i, lang := range []string{languages[0], "en"} {
		for _, ext := range []string{"md", "html"} {
			f := fmt.Sprintf("%s/docs/%s.%s", lang, name, ext)
			if r.TemplatesBox.Has(f) {
				c.Set("sourceRoot", docsRepoBase)
				l := "docs-layout.html"
				if c.Param("_indexing") == "true" {
					l = "search-layout.html"
				}
				c.Set("is_fallback_version", i > 0)
				return c.Render(200, r.HTML(f, l))
			}
		}
	}
	return c.Error(404, errors.Errorf("could not find %s", name))
}
