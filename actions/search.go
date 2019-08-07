package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/gobuffalo/gobuffalo/search/blog"
	"github.com/gobuffalo/gobuffalo/search/godoc"
	"github.com/gobuffalo/gobuffalo/search/site"
	"github.com/gobuffalo/gobuffalo/search/vimeo"
	"github.com/pkg/errors"
)

func StartSearch() {
	search.AddIndex(site.Indexer(App(), r))
	search.AddIndex(blog.Indexer(App()))
	search.AddIndex(vimeo.Indexer(App()))
	search.AddIndex(godoc.Indexer(App()))
}

// Search handles the search queries.
func Search(c buffalo.Context) error {
	if c.Param("query") != "" {
		res, err := search.Fetch(search.Query{
			Text: c.Param("query"),
		})
		if err != nil {
			return errors.WithStack(err)
		}

		c.Set("sourceRoot", docsRepoBase)
		c.Set("results", res)
	}

	return c.Render(200, r.HTML("search.html", "docs-layout.html"))
}
