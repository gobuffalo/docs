package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/events"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/gobuffalo/gobuffalo/search/blog"
	"github.com/gobuffalo/gobuffalo/search/godoc"
	"github.com/gobuffalo/gobuffalo/search/site"
	"github.com/gobuffalo/gobuffalo/search/vimeo"
	"github.com/pkg/errors"
)

func init() {
	StartSearch()
}

func StartSearch() {
	search.AddIndex(site.Indexer(App(), r))
	search.AddIndex(blog.Indexer(App()))
	search.AddIndex(vimeo.Indexer(App()))
	search.AddIndex(godoc.Indexer(App()))

	// Start indexing routine on app start
	events.Listen(func(e events.Event) {
		if e.Kind != buffalo.EvtAppStart {
			return
		}
		go func() {
			events.EmitPayload(search.E_INDEX, events.Payload{})
			for {
				select {
				case <-App().Context.Done():
					return
				default:
					time.Sleep(60 * time.Minute)
				}
			}
		}()
	})
}

func init() {
	StartSearch()
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
		c.Set("results", res)
	}

	return c.Render(200, r.HTML("search.html"))
}
