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

func StartSearch(app *buffalo.App) {
	search.AddIndex(site.Indexer(app, r))
	search.AddIndex(blog.Indexer(app))
	search.AddIndex(vimeo.Indexer(app))
	search.AddIndex(godoc.Indexer(app))
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
