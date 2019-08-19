package actions

import (
	"sync"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/docs/search"
	"github.com/gobuffalo/docs/search/blog"
	"github.com/gobuffalo/docs/search/godoc"
	"github.com/gobuffalo/docs/search/site"
	"github.com/gobuffalo/docs/search/vimeo"
	"github.com/pkg/errors"
)

var searchOnce = &sync.Once{}

func StartSearch(app *buffalo.App) {
	searchOnce.Do(func() {
		for {
			select {
			case <-app.Context.Done():
				break
			default:
				go site.Indexer(app, r)()
				go blog.Indexer(app)()
				go vimeo.Indexer(app)()
				go godoc.Indexer(app)()
				time.Sleep(60 * time.Minute)
			}
		}
	})
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
