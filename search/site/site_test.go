package site_test

import (
	"sync"
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/events"
	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/stretchr/testify/require"
)

func Test_indexDocs(t *testing.T) {
	r := require.New(t)
	wg := &sync.WaitGroup{}
	go func() {
		events.Listen(func(e events.Event) {
			if e.Kind != "gobuffalo:site:search:finished" {
				return
			}
			wg.Done()
		})
	}()
	actions.StartSearch()
	events.EmitPayload(buffalo.EvtAppStart, nil)
	wg.Add(1)
	wg.Wait()

	res, err := search.Fetch(search.Query{
		Text: "error",
	})
	r.NoError(err)
	r.NotEqual(res.Hits, 0)
}
