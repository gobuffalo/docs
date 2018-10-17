package site_test

import (
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/gobuffalo/gobuffalo/search"
	"github.com/gobuffalo/gobuffalo/search/site"
	"github.com/stretchr/testify/require"
)

func Test_indexDocs(t *testing.T) {
	r := require.New(t)

	in := site.Indexer(actions.App(), actions.Renderer())
	r.NoError(in())

	res, err := search.Fetch(search.Query{
		Text: "error",
	})
	r.NoError(err)
	r.NotEqual(res.Hits, 0)
}
