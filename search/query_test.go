package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Fetch(t *testing.T) {
	r := require.New(t)
	q := Query{
		Text: "foo",
	}

	res, err := Fetch(q)
	r.NoError(err)
	r.Len(res.Hits, 3)

	var names []string
	for _, s := range res.Hits {
		names = append(names, string(s.IndexInternalID))
	}
	r.Equal([]string{"/en/foo", "/fr/foo", "godoc.org/foo"}, names)
}
