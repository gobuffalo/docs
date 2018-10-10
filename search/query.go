package search

import (
	"github.com/blevesearch/bleve"
	"github.com/pkg/errors"
)

type Query struct {
	Language int
	Text     string
}

func Fetch(q Query) (*bleve.SearchResult, error) {
	if q.Language == UNKNOWN {
		q.Language = L_EN
	}
	query := bleve.NewMatchPhraseQuery(q.Text)
	// TODO: figure out how to filter by language
	// q1 := bleve.NewMatchQuery(q.Text)
	// q2 := bleve.NewTermQuery(strconv.Itoa(q.Language))
	// query := bleve.NewConjunctionQuery(q1, q2)
	req := bleve.NewSearchRequest(query)
	// req.Fields = []string{"*"}
	req.Size = 100
	req.Highlight = bleve.NewHighlight()
	req.SortBy([]string{"source", "-_score"})
	// lf := bleve.NewFacetRequest("language", q.Language)
	// req.AddFacet("language", lf)

	i, err := index()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := i.Search(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}
